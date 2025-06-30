package analyzer

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/smiller333/dockerutils/src/dockerclient"
)

// AnalyzeImage pulls and analyzes the specified Docker image
func AnalyzeImage(imageName string) (*AnalysisResult, error) {
	if imageName == "" {
		return nil, fmt.Errorf("image name cannot be empty")
	}

	// Initialize result structure for image analysis
	result := &AnalysisResult{
		ImageTag:        imageName,
		IsImageAnalysis: true,
		BuildSuccess:    false, // For images, this indicates successful inspection
	}

	// Create Docker client
	dockerClient, err := dockerclient.NewDefaultClient()
	if err != nil {
		return result, fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer dockerClient.Close()

	// Test connection to Docker daemon
	ctx := context.Background()
	if !dockerClient.IsConnected(ctx) {
		return result, fmt.Errorf("cannot connect to Docker daemon")
	}

	// Pull the image
	startTime := time.Now()
	pullReader, err := dockerClient.PullImage(ctx, imageName, nil)
	if err != nil {
		result.BuildTime = time.Since(startTime).Seconds()
		return result, fmt.Errorf("failed to pull image %s: %w", imageName, err)
	}
	defer pullReader.Close()

	// Read and capture pull output
	_, err = io.ReadAll(pullReader)
	if err != nil {
		result.BuildTime = time.Since(startTime).Seconds()
		return result, fmt.Errorf("failed to read pull output: %w", err)
	}

	result.BuildTime = time.Since(startTime).Seconds()

	// Inspect the image to get detailed information
	imageInfo, err := dockerClient.InspectImage(ctx, imageName)
	if err != nil {
		result.BuildOutput = fmt.Sprintf("Failed to inspect image %s: %v", imageName, err)
		return result, fmt.Errorf("failed to inspect image %s: %w", imageName, err)
	}

	// Extract image information
	result.BuildSuccess = true
	result.ImageSize = imageInfo.Size
	result.Architecture = imageInfo.Architecture
	result.OS = imageInfo.Os
	result.Created = imageInfo.Created
	result.Author = imageInfo.Author

	// Count layers (RootFS layers)
	if imageInfo.RootFS.Type == "layers" {
		result.LayerCount = len(imageInfo.RootFS.Layers)

		// for i, layer := range imageInfo.RootFS.Layers {
		// 	fmt.Printf("Layer %d: %s\n", i+1, layer)
		// }
	}

	// Save the image to a tar file
	err = saveImageToTar(ctx, dockerClient, imageName, result)
	if err != nil {
		fmt.Printf("Failed to save image %s: %v", imageName, err)
		// Continue even if save fails - we still have the analysis results
	}

	// Extract the image if save was successful
	if result.SaveSuccess {
		err = extractImageTar(result.SavedTarPath, result)
		if err != nil {
			fmt.Printf("Failed to extract image %s: %v", imageName, err)
			// Continue even if extraction fails
		}
	}

	// Extract file systems from layer tar files if extraction was successful
	if result.ExtractSuccess {
		err = extractLayerFileSystems(result.ExtractedPath)
		if err != nil {
			fmt.Printf("Failed to extract layer file systems: %v", err)
		}
	}

	// Create a container from the image without starting it
	err = createContainerFromImage(ctx, dockerClient, imageName, result)
	if err != nil {
		fmt.Printf("Failed to create container from image %s: %v", imageName, err)
		// Continue even if container creation fails
	}

	// Copy the container's root filesystem if container creation was successful
	if result.ContainerSuccess {
		err = copyContainerFilesystem(ctx, dockerClient, result)
		if err != nil {
			fmt.Printf("Failed to copy container filesystem: %v", err)
			// Continue even if filesystem copy fails
		}
	}

	// Clean up temporary files and directories, keeping only container_contents and layer_contents
	err = cleanupTemporaryFiles(result)
	if err != nil {
		fmt.Printf("Failed to cleanup temporary files: %v", err)
		// Continue even if cleanup fails - analysis is still successful
	}

	// Clean up the created container if it exists
	if result.ContainerSuccess && result.ContainerID != "" {
		err = dockerClient.RemoveContainer(ctx, result.ContainerID, true)
		if err != nil {
			fmt.Printf("Failed to remove container %s: %v", result.ContainerID, err)
			// Continue even if container removal fails
		}
	}

	return result, nil
}

// saveImageToTar saves the Docker image to a tar file in the tmp directory
func saveImageToTar(ctx context.Context, dockerClient *dockerclient.DockerClient, imageName string, result *AnalysisResult) error {
	// Create tmp directory if it doesn't exist
	tmpDir := "tmp"
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return fmt.Errorf("failed to create tmp directory: %w", err)
	}

	// Generate tar filename from image name and tag
	// Replace problematic characters with underscores
	safeName := strings.ReplaceAll(imageName, ":", "_")
	safeName = strings.ReplaceAll(safeName, "/", "_")
	tarFileName := fmt.Sprintf("%s.tar", safeName)
	tarPath := filepath.Join(tmpDir, tarFileName)

	// Save the image
	saveReader, err := dockerClient.SaveImage(ctx, []string{imageName})
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}
	defer saveReader.Close()

	// Create the tar file
	tarFile, err := os.Create(tarPath)
	if err != nil {
		return fmt.Errorf("failed to create tar file %s: %w", tarPath, err)
	}
	defer tarFile.Close()

	// Copy the image data to the tar file
	_, err = io.Copy(tarFile, saveReader)
	if err != nil {
		return fmt.Errorf("failed to write image data to tar file: %w", err)
	}

	// Update result with successful save information
	result.SavedTarPath = tarPath
	result.SaveSuccess = true

	return nil
}

// extractImageTar extracts the contents of a Docker image tar file to a subdirectory
func extractImageTar(tarPath string, result *AnalysisResult) error {
	// Generate extraction directory name based on tar filename
	tarDir := filepath.Dir(tarPath)
	tarBaseName := filepath.Base(tarPath)
	// Remove .tar extension
	extractDirName := strings.TrimSuffix(tarBaseName, ".tar")
	extractPath := filepath.Join(tarDir, extractDirName)

	// Create extraction directory
	if err := os.MkdirAll(extractPath, 0755); err != nil {
		return fmt.Errorf("failed to create extraction directory %s: %w", extractPath, err)
	}

	// Extract using tar command for simplicity and reliability
	// This handles the Docker image tar format properly
	// Use absolute path for tar file to avoid path issues
	absTarPath, err := filepath.Abs(tarPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for tar file: %w", err)
	}

	// Execute tar extraction directly without cd command for better reliability
	execCmd := exec.Command("tar", "-xf", absTarPath)
	execCmd.Dir = extractPath

	output, err := execCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to extract tar file: %w, output: %s", err, string(output))
	}

	// Update result with successful extraction information
	result.ExtractedPath = extractPath
	result.ExtractSuccess = true

	return nil
}

// extractLayerFileSystems extracts tar files from blobs/sha256 directory into layer_contents subdirectories
func extractLayerFileSystems(extractedImagePath string) error {
	blobsPath := filepath.Join(extractedImagePath, "blobs", "sha256")

	// Check if blobs/sha256 directory exists
	if _, err := os.Stat(blobsPath); os.IsNotExist(err) {
		return fmt.Errorf("blobs/sha256 directory not found in %s", extractedImagePath)
	}

	// Create layer_contents directory
	fileSystemsPath := filepath.Join(extractedImagePath, "layer_contents")
	if err := os.MkdirAll(fileSystemsPath, 0755); err != nil {
		return fmt.Errorf("failed to create layer_contents directory: %w", err)
	}

	// Read all files in the blobs/sha256 directory
	blobFiles, err := os.ReadDir(blobsPath)
	if err != nil {
		return fmt.Errorf("failed to read blobs/sha256 directory: %w", err)
	}

	for _, blobFile := range blobFiles {
		if blobFile.IsDir() {
			continue // Skip directories
		}

		blobName := blobFile.Name()
		blobFilePath := filepath.Join(blobsPath, blobName)

		// Check if this file might be a tar file by trying to extract it
		// Create subdirectory for this layer
		layerDir := filepath.Join(fileSystemsPath, blobName)
		if err := os.MkdirAll(layerDir, 0755); err != nil {
			continue // Skip this layer if we can't create the directory
		}

		// Try to extract the blob as a tar file
		err = extractBlobTar(blobFilePath, layerDir)
		if err != nil {
			// If extraction fails, this might not be a tar file - that's okay
			// Remove the empty directory we created
			os.RemoveAll(layerDir)
			continue
		}
	}

	return nil
}

// extractBlobTar extracts a blob tar file to the specified directory
func extractBlobTar(blobPath, extractDir string) error {
	// Get absolute path for the blob file
	absBlobPath, err := filepath.Abs(blobPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for blob file: %w", err)
	}

	// Execute tar extraction
	execCmd := exec.Command("tar", "-xf", absBlobPath)
	execCmd.Dir = extractDir

	output, err := execCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to extract blob tar file: %w, output: %s", err, string(output))
	}

	return nil
}

// createContainerFromImage creates a container from the specified image without starting it
func createContainerFromImage(ctx context.Context, dockerClient *dockerclient.DockerClient, imageName string, result *AnalysisResult) error {
	// Generate container name based on image name
	// Replace problematic characters with underscores
	safeName := strings.ReplaceAll(imageName, ":", "_")
	safeName = strings.ReplaceAll(safeName, "/", "_")
	containerName := fmt.Sprintf("analysis_%s_%d", safeName, time.Now().Unix())

	// Create basic container configuration
	// Use minimal configuration to just create the container without starting it
	config := &container.Config{
		Image: imageName,
		// Set a simple command that won't interfere with analysis
		Cmd: []string{"true"}, // 'true' command that does nothing and exits successfully
		// Disable networking for security during analysis
		NetworkDisabled: true,
		// Set working directory to root
		WorkingDir: "/",
		// Disable stdin/stdout/stderr attachment
		AttachStdin:  false,
		AttachStdout: false,
		AttachStderr: false,
		Tty:          false,
		OpenStdin:    false,
		StdinOnce:    false,
	}

	// Create host configuration with minimal privileges
	hostConfig := &container.HostConfig{
		// Set restart policy to never restart
		RestartPolicy: container.RestartPolicy{
			Name: "no",
		},
		// Disable auto-removal to allow inspection
		AutoRemove: false,
		// Use default resource limits
		Resources: container.Resources{},
		// Disable privileged mode for security
		Privileged: false,
		// Set read-only root filesystem for analysis safety
		ReadonlyRootfs: true,
	}

	// Create empty networking configuration
	networkingConfig := &network.NetworkingConfig{}

	// Set platform to nil to use default
	var platform *ocispec.Platform = nil

	// Create the container
	createResp, err := dockerClient.CreateContainer(ctx, config, hostConfig, networkingConfig, platform, containerName)
	if err != nil {
		return fmt.Errorf("failed to create container from image %s: %w", imageName, err)
	}

	// Update result with container information
	result.ContainerID = createResp.ID
	result.ContainerName = containerName
	result.ContainerSuccess = true
	result.ContainerWarnings = createResp.Warnings

	return nil
}

// copyContainerFilesystem copies the entire root filesystem from a container to a subdirectory
func copyContainerFilesystem(ctx context.Context, dockerClient *dockerclient.DockerClient, result *AnalysisResult) error {
	if result.ContainerID == "" {
		return fmt.Errorf("container ID is empty")
	}

	// Determine the target directory path
	// If we have an extracted path, use it; otherwise create based on image name
	var baseDir string
	if result.ExtractedPath != "" {
		baseDir = result.ExtractedPath
	} else {
		// Create a directory based on the image name in tmp/
		tmpDir := "tmp"
		safeName := strings.ReplaceAll(result.ImageTag, ":", "_")
		safeName = strings.ReplaceAll(safeName, "/", "_")
		baseDir = filepath.Join(tmpDir, safeName)
		if err := os.MkdirAll(baseDir, 0755); err != nil {
			return fmt.Errorf("failed to create base directory %s: %w", baseDir, err)
		}
	}

	// Create container_contents subdirectory
	containerFSPath := filepath.Join(baseDir, "container_contents")
	if err := os.MkdirAll(containerFSPath, 0755); err != nil {
		return fmt.Errorf("failed to create container_contents directory %s: %w", containerFSPath, err)
	}

	// Copy the entire root filesystem from the container
	// Use "/" as the source path to copy everything from the root
	reader, _, err := dockerClient.CopyFromContainer(ctx, result.ContainerID, "/")
	if err != nil {
		return fmt.Errorf("failed to copy filesystem from container %s: %w", result.ContainerID, err)
	}
	defer reader.Close()

	// Extract the tar archive directly to the container_contents directory
	err = extractTarReader(reader, containerFSPath)
	if err != nil {
		return fmt.Errorf("failed to extract container filesystem tar: %w", err)
	}

	// Update result with successful filesystem copy information
	result.ContainerFSPath = containerFSPath
	result.ContainerFSSuccess = true

	return nil
}

// extractTarReader extracts a tar archive from a reader to the specified directory
func extractTarReader(reader io.Reader, destDir string) error {
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return fmt.Errorf("error reading tar header: %w", err)
		}

		// Build the full path for extraction
		destPath := filepath.Join(destDir, header.Name)

		// Ensure the destination path is within the target directory (security check)
		if !strings.HasPrefix(destPath, filepath.Clean(destDir)+string(os.PathSeparator)) &&
			destPath != filepath.Clean(destDir) {
			return fmt.Errorf("invalid path in tar archive: %s", header.Name)
		}

		// Create the directory structure if needed
		if header.Typeflag == tar.TypeDir {
			if err := os.MkdirAll(destPath, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", destPath, err)
			}
			continue
		}

		// Ensure parent directory exists
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return fmt.Errorf("failed to create parent directory for %s: %w", destPath, err)
		}

		// Handle regular files
		if header.Typeflag == tar.TypeReg {
			outFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", destPath, err)
			}

			_, err = io.Copy(outFile, tarReader)
			outFile.Close()
			if err != nil {
				return fmt.Errorf("failed to write file %s: %w", destPath, err)
			}
		}

		// Handle symbolic links
		if header.Typeflag == tar.TypeSymlink {
			err := os.Symlink(header.Linkname, destPath)
			if err != nil {
				// On some systems, symlink creation might fail, but we can continue
				fmt.Printf("Warning: failed to create symlink %s -> %s: %v\n", destPath, header.Linkname, err)
			}
		}

		// Handle hard links
		if header.Typeflag == tar.TypeLink {
			linkTarget := filepath.Join(destDir, header.Linkname)
			err := os.Link(linkTarget, destPath)
			if err != nil {
				// Hard link creation might fail, but we can continue
				fmt.Printf("Warning: failed to create hard link %s -> %s: %v\n", destPath, linkTarget, err)
			}
		}
	}

	return nil
}

// cleanupTemporaryFiles removes temporary files and directories created during analysis,
// keeping only the container_contents and layer_contents directories.
func cleanupTemporaryFiles(result *AnalysisResult) error {
	var errors []error

	// Remove the original tar file if it exists
	if result.SavedTarPath != "" {
		if err := os.Remove(result.SavedTarPath); err != nil && !os.IsNotExist(err) {
			errors = append(errors, fmt.Errorf("failed to remove tar file %s: %w", result.SavedTarPath, err))
		}
	}

	// Clean up extracted directory contents, keeping only container_contents and layer_contents
	if result.ExtractedPath != "" {
		err := cleanupExtractedDirectory(result.ExtractedPath)
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Return combined errors if any occurred
	if len(errors) > 0 {
		var errorStrings []string
		for _, err := range errors {
			errorStrings = append(errorStrings, err.Error())
		}
		return fmt.Errorf("cleanup errors: %s", strings.Join(errorStrings, "; "))
	}

	return nil
}

// cleanupExtractedDirectory removes all files and directories in the extracted path
// except for container_contents and layer_contents directories.
func cleanupExtractedDirectory(extractedPath string) error {
	// Read all entries in the extracted directory
	entries, err := os.ReadDir(extractedPath)
	if err != nil {
		return fmt.Errorf("failed to read extracted directory %s: %w", extractedPath, err)
	}

	// Directories to keep
	keepDirs := map[string]bool{
		"container_contents": true,
		"layer_contents":     true,
	}

	var errors []error

	// Remove entries that are not in the keep list
	for _, entry := range entries {
		entryName := entry.Name()
		entryPath := filepath.Join(extractedPath, entryName)

		// Skip directories we want to keep
		if entry.IsDir() && keepDirs[entryName] {
			continue
		}

		// Remove the file or directory
		if err := os.RemoveAll(entryPath); err != nil {
			errors = append(errors, fmt.Errorf("failed to remove %s: %w", entryPath, err))
		}
	}

	// Return combined errors if any occurred
	if len(errors) > 0 {
		var errorStrings []string
		for _, err := range errors {
			errorStrings = append(errorStrings, err.Error())
		}
		return fmt.Errorf("cleanup errors in %s: %s", extractedPath, strings.Join(errorStrings, "; "))
	}

	return nil
}

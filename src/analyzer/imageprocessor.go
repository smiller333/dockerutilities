// Copyright (c) 2025 Docker Utils Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

package analyzer

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/smiller333/dockerutilities/src/dockerclient"
)

// SafeTarExtraction extracts tar archive using Go native library
func SafeTarExtraction(tarPath string, destDir string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return fmt.Errorf("failed to open tar file: %w", err)
	}
	defer file.Close()

	var reader io.Reader = file

	// Handle gzipped tar files
	if strings.HasSuffix(tarPath, ".tar.gz") || strings.HasSuffix(tarPath, ".tgz") {
		gzReader, err := gzip.NewReader(file)
		if err != nil {
			return fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzReader.Close()
		reader = gzReader
	}

	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading tar: %w", err)
		}

		// Security: Validate file path
		if err := validateTarPath(header.Name); err != nil {
			return fmt.Errorf("invalid tar path %s: %w", header.Name, err)
		}

		targetPath := filepath.Join(destDir, header.Name)

		// Security: Ensure target path is within destination directory
		if !strings.HasPrefix(targetPath, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return fmt.Errorf("path traversal attempt: %s", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		case tar.TypeReg:
			if err := extractSecureFile(tarReader, targetPath, header.Mode); err != nil {
				return fmt.Errorf("failed to extract file %s: %w", header.Name, err)
			}
		case tar.TypeSymlink:
			// Handle symbolic links with security checks
			if err := extractSecureSymlink(header.Linkname, targetPath, destDir); err != nil {
				// Log warning but continue - symlink creation is not critical
				continue
			}
		case tar.TypeLink:
			// Handle hard links with security checks
			if err := extractSecureHardlink(header.Linkname, targetPath, destDir); err != nil {
				// Log warning but continue - hardlink creation is not critical
				continue
			}
		}
	}

	return nil
}

// validateTarPath validates tar file paths for security
func validateTarPath(path string) error {
	// Check for path traversal attempts
	if strings.Contains(path, "..") {
		return fmt.Errorf("path contains '..' directory traversal")
	}

	// Check for absolute paths
	if filepath.IsAbs(path) {
		return fmt.Errorf("absolute paths not allowed")
	}

	// Check for suspicious paths - allow them but warn
	suspiciousPaths := []string{
		"etc/", "/etc/", "var/", "/var/", "usr/", "/usr/",
		".ssh/", "/.ssh/", ".env", "/.env",
	}

	for _, suspicious := range suspiciousPaths {
		if strings.HasPrefix(path, suspicious) {
			// Don't block, just log - these might be legitimate in container images
			// This is intentionally checking without action for awareness
			break
		}
	}

	return nil
}

// extractSecureFile safely extracts a single file from tar
func extractSecureFile(reader io.Reader, targetPath string, mode int64) error {
	// Create parent directory
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}

	// Create file with restricted permissions
	file, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(mode)&0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy with size limit to prevent zip bombs
	const maxFileSize = 100 * 1024 * 1024 // 100MB limit
	_, err = io.CopyN(file, reader, maxFileSize)
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

// extractSecureSymlink safely extracts a symbolic link with security checks
func extractSecureSymlink(linkTarget, linkPath, baseDir string) error {
	// Security: ensure link target is relative and within base directory
	if filepath.IsAbs(linkTarget) {
		return fmt.Errorf("absolute symlink targets not allowed: %s", linkTarget)
	}

	// Resolve the target path
	resolvedTarget := filepath.Join(filepath.Dir(linkPath), linkTarget)
	cleanTarget := filepath.Clean(resolvedTarget)
	cleanBase := filepath.Clean(baseDir)

	// Ensure target is within base directory
	if !strings.HasPrefix(cleanTarget, cleanBase+string(os.PathSeparator)) && cleanTarget != cleanBase {
		return fmt.Errorf("symlink target outside base directory: %s", linkTarget)
	}

	// Check if the symlink already exists
	if _, err := os.Lstat(linkPath); err == nil {
		// File already exists, skip symlink creation to avoid warnings
		return nil
	}

	return os.Symlink(linkTarget, linkPath)
}

// extractSecureHardlink safely extracts a hard link with security checks
func extractSecureHardlink(linkTarget, linkPath, baseDir string) error {
	// Security: ensure link target is within base directory
	cleanTarget := filepath.Clean(filepath.Join(baseDir, linkTarget))
	cleanBase := filepath.Clean(baseDir)

	if !strings.HasPrefix(cleanTarget, cleanBase+string(os.PathSeparator)) && cleanTarget != cleanBase {
		return fmt.Errorf("hardlink target outside base directory: %s", linkTarget)
	}

	// Check if source exists
	if _, err := os.Stat(cleanTarget); err != nil {
		return fmt.Errorf("hardlink source does not exist: %s", linkTarget)
	}

	return os.Link(cleanTarget, linkPath)
}

// secureFileCreate creates files with restricted permissions
func secureFileCreate(path string) (*os.File, error) {
	// Create file with permissions 0600 (rw-------)
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
}

// ImageInfo represents the JSON info of an analyzed Docker image
type ImageInfo struct {
	ImageID            string                    `json:"image_id"` // ID of the Docker image
	ImageTag           string                    `json:"image_tag"`
	ImageSource        string                    `json:"image_source,omitempty"` // Source registry for non-DockerHub images
	LayerCount         int                       `json:"layer_count"`
	Layers             []string                  `json:"layers"`     // Layer hashes without "blobs/sha256/" prefix
	ImageSize          int64                     `json:"image_size"` // Size in bytes
	Architecture       string                    `json:"architecture"`
	OS                 string                    `json:"os"`
	Created            string                    `json:"created"`
	Author             string                    `json:"author"`
	AnalyzedAt         string                    `json:"analyzed_at"`                   // Timestamp when analysis was performed
	ContainerDirectory *DirectoryInfo            `json:"container_directory,omitempty"` // Container filesystem analysis
	LayerDirectories   map[string]*DirectoryInfo `json:"layer_directories,omitempty"`   // Layer filesystem analysis by layer hash
}

// FileInfo represents information about a file
type FileInfo struct {
	Name string `json:"name"` // File name
	Size int64  `json:"size"` // File size in bytes
}

// DirectoryInfo represents information about a directory and its contents
type DirectoryInfo struct {
	Path        string                    `json:"path"`        // Relative path from analysis root
	Size        int64                     `json:"size"`        // Total size in bytes (including subdirectories)
	FileCount   int                       `json:"file_count"`  // Number of files in this directory
	DirCount    int                       `json:"dir_count"`   // Number of subdirectories
	Files       []FileInfo                `json:"files"`       // List of files with names and sizes in this directory
	Directories map[string]*DirectoryInfo `json:"directories"` // Subdirectories mapped by name
}

// DockerManifest represents the structure of Docker image manifest.json
type DockerManifest struct {
	Config       string                 `json:"Config"`
	RepoTags     []string               `json:"RepoTags"`
	Layers       []string               `json:"Layers"`
	LayerSources map[string]interface{} `json:"LayerSources,omitempty"`
}

// parseImageNameAndSource parses a full image name and returns the image tag and source
// For DockerHub images (e.g., "nginx:latest"), returns tag="nginx:latest", source=""
// For external registries (e.g., "registry.gitlab.com/user/repo/image:tag"),
// returns tag="image:tag", source="registry.gitlab.com/user/repo"
func parseImageNameAndSource(fullImageName string) (imageTag, imageSource string) {
	// Remove any potential docker.io prefix for DockerHub images
	fullImageName = strings.TrimPrefix(fullImageName, "docker.io/")

	// Check if this looks like an external registry (contains domain-like pattern)
	// Split by '/' to analyze the path structure
	parts := strings.Split(fullImageName, "/")

	// If there's only one part or two parts without dots, it's likely DockerHub
	if len(parts) <= 2 && !strings.Contains(parts[0], ".") {
		return fullImageName, ""
	}

	// If the first part contains a dot, it's likely a registry domain
	if strings.Contains(parts[0], ".") {
		// Find the image name (last part before tag)
		imagePart := parts[len(parts)-1]

		// Extract tag if present
		var tag string
		if colonIndex := strings.LastIndex(imagePart, ":"); colonIndex != -1 {
			tag = imagePart[colonIndex:]
			imagePart = imagePart[:colonIndex]
		}

		// Build the image tag (image name + tag)
		imageTag = imagePart + tag

		// Build the source (everything except the last path component)
		if len(parts) > 1 {
			imageSource = strings.Join(parts[:len(parts)-1], "/")
		}

		return imageTag, imageSource
	}

	// Fallback: treat as DockerHub image
	return fullImageName, ""
}

// AnalyzeImageWithTmpDir pulls and analyzes the specified Docker image with a custom tmp directory
// If forcePull is false, it will only pull the image if it doesn't already exist locally
func AnalyzeImageWithTmpDir(imageName string, keepTempFiles bool, forcePull bool, tmpDir string) (*AnalysisResult, error) {
	if imageName == "" {
		return nil, fmt.Errorf("image name cannot be empty")
	}

	// Validate image name format
	if err := dockerclient.ValidateImageName(imageName); err != nil {
		return nil, fmt.Errorf("invalid image name: %w", err)
	}

	// Warn about untrusted images
	dockerclient.WarnUntrustedImage(imageName)

	// Use provided tmp directory and ensure it exists with secure permissions
	tmpBaseDir := tmpDir
	if err := os.MkdirAll(tmpBaseDir, 0700); err != nil {
		return nil, fmt.Errorf("failed to create tmp directory: %w", err)
	}

	// Parse the image name to extract tag and source
	imageTag, imageSource := parseImageNameAndSource(imageName)

	// Initialize result structure for image analysis
	result := &AnalysisResult{
		ImageTag:        imageTag,
		ImageSource:     imageSource,
		IsImageAnalysis: true,
		BuildSuccess:    false,      // For images, this indicates successful inspection
		ExtractedPath:   tmpBaseDir, // Set the base directory for all operations
	}

	// Validate Docker access and warn about security implications
	if err := dockerclient.ValidateDockerAccess(); err != nil {
		return result, fmt.Errorf("docker access validation failed: %w", err)
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

	// Check if image exists locally unless force pull is requested
	var needsPull bool = forcePull
	if !forcePull {
		// Try to inspect the image to see if it exists locally
		_, err := dockerClient.InspectImage(ctx, imageName)
		if err != nil {
			// Image doesn't exist locally, need to pull it
			needsPull = true
		}
	}

	// Pull the image if needed
	startTime := time.Now()
	if needsPull {
		pullReader, err := dockerClient.PullImage(ctx, imageName, nil)
		if err != nil {
			result.BuildTime = time.Since(startTime).Seconds()
			if errors.Is(err, dockerclient.ErrImageNotFound) {
				return result, err
			}

			return result, fmt.Errorf("failed to pull image %s: %w", imageName, err)

		}
		defer pullReader.Close()

		// Read and capture pull output
		_, err = io.ReadAll(pullReader)
		if err != nil {
			result.BuildTime = time.Since(startTime).Seconds()
			return result, fmt.Errorf("failed to read pull output: %w", err)
		}

		result.Pulled = true
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
	result.ImageID = imageInfo.ID
	result.ImageSize = imageInfo.Size
	result.Architecture = imageInfo.Architecture
	result.OS = imageInfo.Os
	result.Created = imageInfo.Created
	result.Author = imageInfo.Author

	// Count layers (RootFS layers)
	if imageInfo.RootFS.Type == "layers" {
		result.LayerCount = len(imageInfo.RootFS.Layers)
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

	// Create image summary JSON file after filesystem operations are complete
	err = createImageInfo(result)
	if err != nil {
		fmt.Printf("Failed to create image summary: %v", err)
		// Continue even if summary creation fails - analysis is still successful
	}

	// Clean up temporary files and directories, keeping only container_contents and layer_contents
	if !keepTempFiles {
		err = cleanupTemporaryFiles(result)
		if err != nil {
			fmt.Printf("Failed to cleanup temporary files: %v", err)
			// Continue even if cleanup fails - analysis is still successful
		}
	} else {
		fmt.Printf("Temporary files preserved at: %s\n", result.ExtractedPath)
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

// saveImageToTar saves the Docker image to a tar file in a temporary directory
func saveImageToTar(ctx context.Context, dockerClient *dockerclient.DockerClient, imageName string, result *AnalysisResult) error {
	// Use the pre-created temporary directory from the result
	tmpDir := result.ExtractedPath
	if tmpDir == "" {
		return fmt.Errorf("temporary directory not set in result")
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

	// Create the tar file with secure permissions
	tarFile, err := secureFileCreate(tarPath)
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

	// Create extraction directory with secure permissions
	if err := os.MkdirAll(extractPath, 0700); err != nil {
		return fmt.Errorf("failed to create extraction directory %s: %w", extractPath, err)
	}

	// Extract using secure Go native tar extraction
	err := SafeTarExtraction(tarPath, extractPath)
	if err != nil {
		return fmt.Errorf("failed to extract tar file: %w", err)
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

	// Create layer_contents directory with secure permissions
	fileSystemsPath := filepath.Join(extractedImagePath, "layer_contents")
	if err := os.MkdirAll(fileSystemsPath, 0700); err != nil {
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
		// Create subdirectory for this layer with secure permissions
		layerDir := filepath.Join(fileSystemsPath, blobName)
		if err := os.MkdirAll(layerDir, 0700); err != nil {
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
	// Extract using secure Go native tar extraction
	err := SafeTarExtraction(blobPath, extractDir)
	if err != nil {
		return fmt.Errorf("failed to extract blob tar file: %w", err)
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
	// Use the pre-created temporary directory from the result
	var baseDir string
	if result.ExtractedPath != "" {
		baseDir = result.ExtractedPath
	} else {
		return fmt.Errorf("temporary directory not set in result")
	}

	// Create container_contents subdirectory with secure permissions
	containerFSPath := filepath.Join(baseDir, "container_contents")
	if err := os.MkdirAll(containerFSPath, 0700); err != nil {
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
			// Check if the symlink already exists
			if _, err := os.Lstat(destPath); err == nil {
				// File already exists, skip symlink creation to avoid warnings
				continue
			}

			err := os.Symlink(header.Linkname, destPath)
			if err != nil {
				// Symlink creation failed, but we can continue analysis without it
				// Only log detailed errors in debug mode to reduce noise
				continue
			}
		}

		// Handle hard links
		if header.Typeflag == tar.TypeLink {
			// Check if the target file already exists
			if _, err := os.Lstat(destPath); err == nil {
				// File already exists, skip hard link creation
				continue
			}

			linkTarget := filepath.Join(destDir, header.Linkname)
			err := os.Link(linkTarget, destPath)
			if err != nil {
				// Hard link creation failed, but we can continue analysis without it
				continue
			}
		}
	}

	return nil
}

// cleanupTemporaryFiles removes temporary files and directories created during analysis,
// keeping only the summary JSON files.
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
// except for the summary JSON files.
func cleanupExtractedDirectory(extractedPath string) error {
	// Read all entries in the extracted directory
	entries, err := os.ReadDir(extractedPath)
	if err != nil {
		return fmt.Errorf("failed to read extracted directory %s: %w", extractedPath, err)
	}

	var errors []error

	// Remove entries that are not in the keep list
	for _, entry := range entries {
		entryName := entry.Name()
		entryPath := filepath.Join(extractedPath, entryName)

		// Keep JSON files that match the summary pattern: "info.{12 hex chars}.json"
		if !entry.IsDir() && strings.HasPrefix(entryName, "info.") && strings.HasSuffix(entryName, ".json") {
			continue // Keep summary JSON files
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

// analyzeDirectory recursively analyzes a directory and returns DirectoryInfo
func analyzeDirectory(dirPath string, relativePath string) (*DirectoryInfo, error) {
	// Get directory info
	dirStat, err := os.Stat(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to stat directory %s: %w", dirPath, err)
	}

	if !dirStat.IsDir() {
		return nil, fmt.Errorf("path %s is not a directory", dirPath)
	}

	dirInfo := &DirectoryInfo{
		Path:        relativePath,
		Size:        0,
		FileCount:   0,
		DirCount:    0,
		Files:       make([]FileInfo, 0),
		Directories: make(map[string]*DirectoryInfo),
	}

	// Read directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	// Process each entry
	for _, entry := range entries {
		entryPath := filepath.Join(dirPath, entry.Name())
		entryRelativePath := filepath.Join(relativePath, entry.Name())

		if entry.IsDir() {
			// Recursively analyze subdirectory
			subDirInfo, err := analyzeDirectory(entryPath, entryRelativePath)
			if err != nil {
				// Log error but continue with other directories
				fmt.Printf("Warning: failed to analyze subdirectory %s: %v\n", entryPath, err)
				continue
			}

			// Add to directories map and update counts
			dirInfo.Directories[entry.Name()] = subDirInfo
			dirInfo.DirCount++
			dirInfo.Size += subDirInfo.Size
		} else {
			// Process file - collect size and name information
			fileInfo, err := entry.Info()
			if err != nil {
				// Log error but continue with other files
				fmt.Printf("Warning: failed to get info for file %s: %v\n", entryPath, err)
				continue
			}

			// Create FileInfo struct with name and size
			fileData := FileInfo{
				Name: entry.Name(),
				Size: fileInfo.Size(),
			}

			// Track file count, add size to directory total, and add file info to files list
			dirInfo.FileCount++
			dirInfo.Size += fileInfo.Size()
			dirInfo.Files = append(dirInfo.Files, fileData)
		}
	}

	return dirInfo, nil
}

// analyzeLayerContents analyzes all layer directories in layer_contents folder
func analyzeLayerContents(extractedPath string) (map[string]*DirectoryInfo, error) {
	layerContentsPath := filepath.Join(extractedPath, "layer_contents")

	// Check if layer_contents directory exists
	if _, err := os.Stat(layerContentsPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("layer_contents directory not found in %s", extractedPath)
	}

	layerDirectories := make(map[string]*DirectoryInfo)

	// Read all layer directories
	layerDirs, err := os.ReadDir(layerContentsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read layer_contents directory: %w", err)
	}

	for _, layerDir := range layerDirs {
		if !layerDir.IsDir() {
			continue // Skip non-directories
		}

		layerName := layerDir.Name()
		layerPath := filepath.Join(layerContentsPath, layerName)

		// Analyze this layer directory
		layerInfo, err := analyzeDirectory(layerPath, layerName)
		if err != nil {
			fmt.Printf("Warning: failed to analyze layer %s: %v\n", layerName, err)
			continue
		}

		layerDirectories[layerName] = layerInfo
	}

	return layerDirectories, nil
}

// analyzeContainerContents analyzes the container_contents directory
func analyzeContainerContents(extractedPath string) (*DirectoryInfo, error) {
	containerContentsPath := filepath.Join(extractedPath, "container_contents")

	// Check if container_contents directory exists
	if _, err := os.Stat(containerContentsPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("container_contents directory not found in %s", extractedPath)
	}

	// Analyze the container contents directory
	containerInfo, err := analyzeDirectory(containerContentsPath, "container_contents")
	if err != nil {
		return nil, fmt.Errorf("failed to analyze container_contents: %w", err)
	}

	return containerInfo, nil
}

// createImageInfo creates a JSON info file with key information about the analyzed image
func createImageInfo(result *AnalysisResult) error {
	if result.ExtractedPath == "" {
		return fmt.Errorf("extracted path not available in analysis result")
	}

	// Read and parse the manifest.json file
	manifestPath := filepath.Join(result.ExtractedPath, "manifest.json")
	layers, err := extractLayersFromManifest(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to extract layers from manifest: %w", err)
	}

	// Analyze container filesystem if available
	var containerDirInfo *DirectoryInfo
	if result.ContainerFSSuccess && result.ContainerFSPath != "" {
		containerDirInfo, err = analyzeContainerContents(result.ExtractedPath)
		if err != nil {
			fmt.Printf("Warning: failed to analyze container contents: %v\n", err)
		}
	}

	// Analyze layer filesystems if available
	var layerDirInfos map[string]*DirectoryInfo
	if result.ExtractSuccess {
		layerDirInfos, err = analyzeLayerContents(result.ExtractedPath)
		if err != nil {
			fmt.Printf("Warning: failed to analyze layer contents: %v\n", err)
		}
	}

	// Create the image info
	info := ImageInfo{
		ImageID:            result.ImageID,
		ImageTag:           result.ImageTag,
		ImageSource:        result.ImageSource,
		LayerCount:         result.LayerCount,
		Layers:             layers,
		ImageSize:          result.ImageSize,
		Architecture:       result.Architecture,
		OS:                 result.OS,
		Created:            result.Created,
		Author:             result.Author,
		AnalyzedAt:         time.Now().UTC().Format(time.RFC3339),
		ContainerDirectory: containerDirInfo,
		LayerDirectories:   layerDirInfos,
	}

	// Generate filename based on first 12 characters of image ID (without sha256: prefix)
	imageIDShort := strings.TrimPrefix(result.ImageID, "sha256:")
	if len(imageIDShort) > 12 {
		imageIDShort = imageIDShort[:12]
	}
	infoFileName := fmt.Sprintf("info.%s.json", imageIDShort)

	// Write the summary to a JSON file
	infoPath := filepath.Join(result.ExtractedPath, infoFileName)
	err = writeImageIntoToFile(info, infoPath)
	if err != nil {
		return fmt.Errorf("failed to write summary to file: %w", err)
	}

	return nil
}

// extractLayersFromManifest reads the manifest.json file and extracts layer hashes without the "blobs/sha256/" prefix
func extractLayersFromManifest(manifestPath string) ([]string, error) {
	// Check if manifest file exists
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("manifest.json not found at %s", manifestPath)
	}

	// Read the manifest file
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest file: %w", err)
	}

	// Parse the manifest JSON
	var manifests []DockerManifest
	err = json.Unmarshal(manifestData, &manifests)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest JSON: %w", err)
	}

	// Check if we have at least one manifest entry
	if len(manifests) == 0 {
		return nil, fmt.Errorf("no manifest entries found")
	}

	// Extract layers from the first manifest entry
	manifest := manifests[0]
	var cleanLayers []string

	for _, layer := range manifest.Layers {
		// Remove the "blobs/sha256/" prefix from each layer
		cleanLayer := strings.TrimPrefix(layer, "blobs/sha256/")
		cleanLayers = append(cleanLayers, cleanLayer)
	}

	return cleanLayers, nil
}

// writeImageIntoToFile marshals the ImageInfo to JSON and writes it to the specified file
func writeImageIntoToFile(info ImageInfo, filePath string) error {
	// Marshal the summary to JSON with proper indentation
	jsonData, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal info to JSON: %w", err)
	}

	// Write the JSON data to the file
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file %s: %w", filePath, err)
	}

	return nil
}

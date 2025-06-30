package analyzer

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

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

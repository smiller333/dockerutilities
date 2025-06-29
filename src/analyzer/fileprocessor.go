package analyzer

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/smiller333/dockerutils/src/dockerclient"
)

// AnalyzeDockerfile reads and analyzes a Dockerfile at the specified path, then builds the image
func AnalyzeDockerfile(dockerfilePath string) (*AnalysisResult, error) {
	// Check if the file exists
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("dockerfile not found: %s", dockerfilePath)
	}

	// Get absolute path for better error reporting
	absPath, err := filepath.Abs(dockerfilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Read the Dockerfile
	content, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dockerfile %s: %w", absPath, err)
	}

	// Generate image tag based on the dockerfile path
	imageTag := generateImageTag(absPath)

	result := &AnalysisResult{
		Path:         dockerfilePath,
		AbsolutePath: absPath,
		Size:         len(content),
		Content:      string(content),
		ImageTag:     imageTag,
		BuildSuccess: false,
	}

	// Create Docker client and build the image
	dockerClient, err := dockerclient.NewDefaultClient()
	if err != nil {
		result.BuildOutput = fmt.Sprintf("Failed to create Docker client: %v", err)
		return result, nil // Return result with build failure, don't fail the analysis
	}
	defer dockerClient.Close()

	// Test Docker connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Start timing from Docker connection test
	var buildStartTime time.Time
	buildStartTime = time.Now()

	if err := dockerClient.Ping(ctx); err != nil {
		result.BuildTime = time.Since(buildStartTime).Seconds()
		result.BuildOutput = fmt.Sprintf("Failed to connect to Docker daemon: %v", err)
		return result, nil
	}

	// Build the image
	buildCtx, buildCancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer buildCancel()

	// Continue timing from connection test (buildStartTime already declared above)
	buildStartTime = time.Now()

	buildResponse, err := dockerClient.BuildImage(buildCtx, absPath, imageTag)
	if err != nil {
		result.BuildTime = time.Since(buildStartTime).Seconds()
		result.BuildOutput = fmt.Sprintf("Build failed: %v", err)
		return result, nil
	}
	defer buildResponse.Close()

	// Read the build output
	buildOutput, err := io.ReadAll(buildResponse)
	if err != nil {
		result.BuildTime = time.Since(buildStartTime).Seconds()
		result.BuildOutput = fmt.Sprintf("Failed to read build output: %v", err)
		return result, nil
	}

	// Calculate total build time
	result.BuildTime = time.Since(buildStartTime).Seconds()
	result.BuildOutput = string(buildOutput)
	result.BuildSuccess = true

	// Get layer count from the built image
	layerCount, err := dockerClient.GetImageLayerCount(buildCtx, imageTag)
	if err != nil {
		// Don't fail the analysis if we can't get layer count, just log it in build output
		result.BuildOutput += fmt.Sprintf("\nWarning: Failed to get layer count: %v", err)
	} else {
		result.LayerCount = layerCount
	}

	return result, nil
}

// generateImageTag creates a unique image tag based on the dockerfile path
func generateImageTag(dockerfilePath string) string {
	// Get the directory name and use it as part of the tag
	dir := filepath.Dir(dockerfilePath)
	baseName := filepath.Base(dir)

	// If we're in the root or the directory name is not useful, use "dockerutils"
	if baseName == "." || baseName == "/" || baseName == "" {
		baseName = "dockerutils"
	}

	// Clean the name to be Docker-compatible (lowercase, no special chars except -, _)
	baseName = strings.ToLower(baseName)
	baseName = strings.ReplaceAll(baseName, " ", "-")

	return fmt.Sprintf("%s:test", baseName)
}

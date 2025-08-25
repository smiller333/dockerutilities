// Copyright (c) 2025 Docker Utils Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

package analyzer

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/smiller333/dockerutilities/src/dockerclient"
)

// closeWithErrorCheck is a helper function to close resources and log any errors
func closeWithErrorCheck(closer io.Closer, resourceName string) {
	if err := closer.Close(); err != nil {
		// Log the error but don't fail the operation
		fmt.Printf("Warning: failed to close %s: %v\n", resourceName, err)
	}
}

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
		DFSize:       len(content),
		Content:      string(content),
		ImageTag:     imageTag,
		BuildSuccess: false,
	}

	// Validate Docker access and warn about security implications
	if err := dockerclient.ValidateDockerAccess(); err != nil {
		result.BuildOutput = fmt.Sprintf("Docker access validation failed: %v", err)
		return result, nil // Return result with build failure, don't fail the analysis
	}

	// Create Docker client and build the image
	dockerClient, err := dockerclient.NewDefaultClient()
	if err != nil {
		result.BuildOutput = fmt.Sprintf("Failed to create Docker client: %v", err)
		return result, nil // Return result with build failure, don't fail the analysis
	}
	defer closeWithErrorCheck(dockerClient, "Docker client")

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
	defer closeWithErrorCheck(buildResponse, "Build response")

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

	// Get the image information
	imageInfo, err := dockerClient.InspectImage(buildCtx, imageTag)
	if err != nil {
		result.BuildOutput += fmt.Sprintf("\nWarning: Failed to inspect image: %v", err)
	} else {
		result.LayerCount = len(imageInfo.RootFS.Layers)
		result.ImageSize = imageInfo.Size

		if len(imageInfo.RootFS.Layers) > 0 {
			result.BuildOutput += fmt.Sprintln("\n---")
			result.BuildOutput += fmt.Sprintln("\nRootFS layers:")
			for _, layer := range imageInfo.RootFS.Layers {
				result.BuildOutput += fmt.Sprintf("- %s\n", layer)
			}
		} else {
			result.BuildOutput += "\nNo RootFS layers found in the image."
		}
	}

	return result, nil
}

// generateImageTag creates a unique image tag based on the dockerfile path
func generateImageTag(dockerfilePath string) string {
	// Get the directory name and use it as part of the tag
	dir := filepath.Dir(dockerfilePath)
	baseName := filepath.Base(dir)

	// If we're in the root or the directory name is not useful, use "dockerutilities"
	if baseName == "." || baseName == "/" || baseName == "" {
		baseName = "dockerutilities"
	}

	// Clean the name to be Docker-compatible (lowercase, no special chars except -, _)
	baseName = strings.ToLower(baseName)
	baseName = strings.ReplaceAll(baseName, " ", "-")

	return fmt.Sprintf("%s:test", baseName)
}

// Package analyzer provides functionality for analyzing Dockerfiles and Docker images.
package analyzer

import (
	"fmt"
)

// AnalysisResult contains the results of analyzing a Dockerfile
type AnalysisResult struct {
	Path         string
	AbsolutePath string
	DFSize       int   // Dockerfile size in bytes
	ImageSize    int64 // Size of the built Docker image in bytes
	Content      string
	ImageTag     string
	BuildOutput  string
	BuildSuccess bool
	LayerCount   int
	BuildTime    float64 // Time taken to build the image in seconds
}

// PrintAnalysisResult prints a formatted summary of the analysis result
func PrintAnalysisResult(result *AnalysisResult, showBuildOutput bool) {
	fmt.Printf("Successfully read Dockerfile: %s\n", result.AbsolutePath)
	fmt.Printf("File size: %d bytes\n", result.DFSize)

	if result.ImageTag != "" {
		if result.BuildSuccess {
			fmt.Printf("Image layers: %d\n", result.LayerCount)
			fmt.Printf("Image size: %d bytes\n", result.ImageSize)
			fmt.Printf("Docker image: %s\n", result.ImageTag)
			fmt.Printf("Build time: %.3f seconds\n\n", result.BuildTime)

			fmt.Println("Build status: SUCCESS")
		} else {
			fmt.Printf("Build time: %.3f seconds\n", result.BuildTime)
			fmt.Println("Build status: FAILED")
		}
		fmt.Println("")
	}

	if showBuildOutput && result.BuildOutput != "" {
		fmt.Println("Build Output:")
		fmt.Println("---")
		fmt.Println(result.BuildOutput)
		fmt.Println("---")
	}
}

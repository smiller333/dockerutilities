// Package analyzer provides functionality for analyzing Dockerfiles and Docker images.
package analyzer

import (
	"fmt"
)

// AnalysisResult contains the results of analyzing a Dockerfile or Docker image
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

	// Image-specific fields
	IsImageAnalysis bool   // Indicates if this is an image analysis (vs Dockerfile)
	Architecture    string // Image architecture (e.g., amd64, arm64)
	OS              string // Image operating system (e.g., linux, windows)
	Created         string // Image creation timestamp
	Author          string // Image author
	SavedTarPath    string // Path to the saved tar file
	SaveSuccess     bool   // Whether the image save was successful
	ExtractedPath   string // Path to the extracted image contents
	ExtractSuccess  bool   // Whether the image extraction was successful
}

// PrintAnalysisResult prints a formatted summary of the analysis result
func PrintAnalysisResult(result *AnalysisResult, showBuildOutput bool) {
	if result.IsImageAnalysis {
		// Print image analysis results
		fmt.Printf("Successfully analyzed Docker image: %s\n", result.ImageTag)

		if result.BuildSuccess {
			fmt.Printf("Image size: %d bytes\n", result.ImageSize)
			fmt.Printf("Image layers: %d\n", result.LayerCount)
			fmt.Printf("Architecture: %s\n", result.Architecture)
			fmt.Printf("OS: %s\n", result.OS)
			if result.Created != "" {
				fmt.Printf("Created: %s\n", result.Created)
			}
			if result.Author != "" {
				fmt.Printf("Author: %s\n", result.Author)
			}
			fmt.Println("Inspection status: SUCCESS")

			// Print save tar information
			if result.SaveSuccess {
				fmt.Printf("Saved tar: %s\n", result.SavedTarPath)
				fmt.Println("Save status: SUCCESS")
			} else if result.SavedTarPath != "" {
				fmt.Println("Save status: FAILED")
			}

			// Print extraction information
			if result.ExtractSuccess {
				fmt.Printf("Extracted to: %s\n", result.ExtractedPath)
				fmt.Println("Extract status: SUCCESS")
			} else if result.ExtractedPath != "" {
				fmt.Println("Extract status: FAILED")
			}
		} else {
			fmt.Println("Inspection status: FAILED")
		}
	} else {
		// Original Dockerfile analysis output
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
}

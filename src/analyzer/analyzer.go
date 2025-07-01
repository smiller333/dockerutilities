// Package analyzer provides functionality for analyzing Dockerfiles and Docker images.
package analyzer

import (
	"fmt"
)

// AnalysisResult contains the results of analyzing a Dockerfile or Docker image
type AnalysisResult struct {
	ImageID      string // ID of the Docker image
	Pulled       bool   // Whether the image was pulled
	Path         string
	AbsolutePath string
	DFSize       int   // Dockerfile size in bytes
	ImageSize    int64 // Size of the built Docker image in bytes
	Content      string
	ImageTag     string
	ImageSource  string // Source registry for non-DockerHub images
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

	// Container-specific fields
	ContainerID       string   // ID of the created container
	ContainerName     string   // Name of the created container
	ContainerSuccess  bool     // Whether container creation was successful
	ContainerWarnings []string // Warnings from container creation

	// Container filesystem fields
	ContainerFSPath    string // Path to the extracted container filesystem
	ContainerFSSuccess bool   // Whether container filesystem extraction was successful
}

// PrintAnalysisResult prints a formatted summary of the analysis result
func PrintAnalysisResult(result *AnalysisResult, showBuildOutput bool) {
	if result.IsImageAnalysis {
		// Print image analysis results
		fmt.Printf("Successfully analyzed Docker image: %s\n", result.ImageTag)
		fmt.Printf("Image Pulled: %t\n", result.Pulled)
		fmt.Printf("Image Tag: %s\n", result.ImageTag)
		fmt.Printf("Image Source: %s\n", result.ImageSource)

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

			// Print container creation information
			if result.ContainerSuccess {
				fmt.Printf("Container ID: %s\n", result.ContainerID)
				if result.ContainerName != "" {
					fmt.Printf("Container Name: %s\n", result.ContainerName)
				}
				fmt.Println("Container creation status: SUCCESS")
				if len(result.ContainerWarnings) > 0 {
					fmt.Println("Container warnings:")
					for _, warning := range result.ContainerWarnings {
						fmt.Printf("  - %s\n", warning)
					}
				}

				// Print container filesystem copy information
				if result.ContainerFSSuccess {
					fmt.Printf("Container filesystem copied to: %s\n", result.ContainerFSPath)
					fmt.Println("Container filesystem copy status: SUCCESS")
				} else if result.ContainerFSPath != "" {
					fmt.Println("Container filesystem copy status: FAILED")
				}
			} else if result.ContainerID != "" {
				fmt.Println("Container creation status: FAILED")
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

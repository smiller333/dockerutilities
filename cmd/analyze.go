// Package cmd provides command-line interface functionality for dockerutils.
package cmd

import (
	"fmt"

	"github.com/smiller333/dockerutils/src/analyzer"
	"github.com/spf13/cobra"
)

var (
	// Flags for the analyze command
	dockerfilePath  string
	imageTag        string
	showBuildOutput bool
	keepTempFiles   bool
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a Dockerfile or Docker image",
	Long: `Analyze a Dockerfile or Docker image to understand its structure and contents.

You can analyze either:
- A Dockerfile by specifying the --dockerfile flag with the path to the file
- A Docker image by specifying the --image flag with the image name and tag

Examples:
  dockerutils analyze --dockerfile ./Dockerfile
  dockerutils analyze --image nginx:latest
  dockerutils analyze --image myregistry.com/myapp:v1.0.0`,
	Args: cobra.NoArgs,
	RunE: runAnalyze,
}

func init() {
	// Add the analyze command to the root command
	rootCmd.AddCommand(analyzeCmd)

	// Add mutually exclusive flags
	analyzeCmd.Flags().StringVar(&dockerfilePath, "dockerfile", "", "Path to the Dockerfile to analyze")
	analyzeCmd.Flags().StringVar(&imageTag, "image", "", "Docker image name and tag to analyze (e.g., nginx:latest)")
	analyzeCmd.Flags().BoolVar(&showBuildOutput, "build-output", false, "Show Docker build output in analysis results")
	analyzeCmd.Flags().BoolVar(&keepTempFiles, "keep-temp", false, "Keep temporary files after analysis (useful for debugging)")

	// Mark the flags as mutually exclusive
	analyzeCmd.MarkFlagsMutuallyExclusive("dockerfile", "image")
}

// runAnalyze executes the analyze command with the provided flags
func runAnalyze(cmd *cobra.Command, args []string) error {
	// Validate that exactly one flag is provided
	if dockerfilePath == "" && imageTag == "" {
		return fmt.Errorf("you must specify either --dockerfile or --image flag")
	}

	if dockerfilePath != "" {
		return analyzeDockerfile(dockerfilePath, showBuildOutput)
	}

	if imageTag != "" {
		return analyzeImage(imageTag, keepTempFiles)
	}

	return nil
}

// analyzeDockerfile reads and analyzes the specified Dockerfile
func analyzeDockerfile(path string, showBuildOutput bool) error {
	// Use the analyzer package to perform the analysis
	result, err := analyzer.AnalyzeDockerfile(path)
	if err != nil {
		return err
	}

	// Print the analysis result
	analyzer.PrintAnalysisResult(result, showBuildOutput)

	return nil
}

// analyzeImage analyzes the specified Docker image
func analyzeImage(image string, keepTempFiles bool) error {
	// Use the analyzer package to perform the image analysis
	result, err := analyzer.AnalyzeImage(image, keepTempFiles)
	if err != nil {
		return err
	}

	// Print the analysis result
	analyzer.PrintAnalysisResult(result, showBuildOutput)

	return nil
}

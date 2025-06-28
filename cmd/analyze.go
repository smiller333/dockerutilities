// Package cmd provides command-line interface functionality for dockerutils.
package cmd

import (
	"github.com/smiller333/dockerutils/src/analyzer"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze [dockerfile-path]",
	Short: "Analyze a Dockerfile",
	Long: `Analyze a Dockerfile to understand its structure and contents.

This command reads a Dockerfile and performs analysis on its contents.
You can specify the path to the Dockerfile as an argument.`,
	Args: cobra.ExactArgs(1),
	RunE: analyzeDockerfile,
}

func init() {
	// Add the analyze command to the root command
	rootCmd.AddCommand(analyzeCmd)
}

// analyzeDockerfile reads and analyzes the specified Dockerfile
func analyzeDockerfile(cmd *cobra.Command, args []string) error {
	dockerfilePath := args[0]

	// Use the analyzer package to perform the analysis
	result, err := analyzer.AnalyzeDockerfile(dockerfilePath)
	if err != nil {
		return err
	}

	// Print the analysis result
	analyzer.PrintAnalysisResult(result)

	return nil
}

// Package cmd provides command-line interface functionality for dockerutils.
// Copyright (c) 2025 Docker Utils Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

package cmd

import (
	"fmt"

	"github.com/smiller333/dockerutils/src/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dockerutils",
	Short: "Docker analysis and management utilities",
	Long: `dockerutils is a comprehensive CLI tool for Docker image analysis and management.

This tool provides:
- Interactive web interface for Docker image analysis
- Command-line utilities for Docker operations
- REST API for programmatic access to analysis data
- Docker build context preview and validation
- Image filesystem exploration and metadata extraction

Key Features:
- Analyze local and remote Docker images
- Extract and browse image filesystem contents  
- View layer information and build history
- Generate comprehensive analysis reports
- Docker build context preview and .dockerignore validation

Examples:
  dockerutils --version                    # Show version information
  dockerutils server --port 8080          # Start web analysis server
  dockerutils completion bash             # Generate bash completion script

For more information about a command, use 'dockerutils [command] --help'`,
	Version: version.GetVersionString(),
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information of dockerutils",
	Long:  `Print detailed version information of dockerutils including build time, git commit, and Go version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(version.GetFullVersionString())
	},
}

func init() {
	// Add the version command to the root command
	rootCmd.AddCommand(versionCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

// Package cmd provides command-line interface functionality for dockerutilities.
// Copyright (c) 2025 Docker Utils Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

package cmd

import (
	"fmt"

	"github.com/smiller333/dockerutilities/src/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dockerutilities",
	Short: "Docker image analysis and build context utilities",
	Long: `dockerutilities is a CLI tool for analyzing Docker images and managing Docker build contexts.

This tool provides:
- Web interface for Docker image analysis with REST API
- Docker build context preview and .dockerignore validation
- Image filesystem exploration and metadata extraction
- Local development utilities (not for production use)

Key Features:
- Analyze local and remote Docker images via web interface
- Extract and browse image filesystem contents (with size limits)
- View layer information and build history
- Preview Docker build contexts with .dockerignore validation
- REST API for programmatic access to analysis data

Examples:
  	dockerutilities --version                    # Show version information
	dockerutilities server --port 8080          # Start web analysis server

For more information about a command, use 'dockerutilities [command] --help'`,
	Version: version.GetVersionString(),
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information of dockerutilities",
	Long:  `Print detailed version information of dockerutilities including build time, git commit, and Go version.`,
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

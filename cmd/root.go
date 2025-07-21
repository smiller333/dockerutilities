// Package cmd provides command-line interface functionality for dockerutils.
package cmd

import (
	"fmt"

	"github.com/smiller333/dockerutils/src/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dockerutils",
	Short: "A collection of Docker utilities",
	Long: `dockerutils is a CLI tool that provides various utilities for working with Docker containers, images, and other Docker-related tasks.

This tool aims to simplify common Docker operations and provide additional functionality for Docker workflows.`,
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

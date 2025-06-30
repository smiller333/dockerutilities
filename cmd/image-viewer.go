// Package cmd provides command-line interface functionality for dockerutils.
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/smiller333/dockerutils/src/webserver"
	"github.com/spf13/cobra"
)

var (
	// Flags for the image-viewer command
	serverPort string
	webRoot    string
	host       string
)

// imageViewerCmd represents the image-viewer command
var imageViewerCmd = &cobra.Command{
	Use:   "image-viewer",
	Short: "Start a web server for viewing Docker image analysis results",
	Long: `Start a local web server that provides a web interface for viewing Docker image analysis results.

The server provides:
- A static web interface for browsing image analysis data
- REST API endpoints for accessing image summaries and filesystem data
- Interactive visualization of Docker image layers and contents

Examples:
  dockerutils image-viewer                           # Start server on default port 8080
  dockerutils image-viewer --port 3000              # Start server on port 3000
  dockerutils image-viewer --port 8080 --host 0.0.0.0 # Bind to all interfaces
  dockerutils image-viewer --web-root ./custom-ui   # Use custom web root directory`,
	Args: cobra.NoArgs,
	RunE: runImageViewer,
}

func init() {
	// Add the image-viewer command to the root command
	rootCmd.AddCommand(imageViewerCmd)

	// Add command flags
	imageViewerCmd.Flags().StringVar(&serverPort, "port", "8080", "Port number for the web server")
	imageViewerCmd.Flags().StringVar(&host, "host", "localhost", "Host/IP address to bind the server to")
	imageViewerCmd.Flags().StringVar(&webRoot, "web-root", "", "Custom path to web root directory (optional)")
}

// runImageViewer starts the web server for viewing Docker image analysis results
func runImageViewer(cmd *cobra.Command, args []string) error {
	fmt.Printf("Starting Docker Image Viewer server...\n")
	fmt.Printf("Server will be available at: http://%s:%s\n", host, serverPort)

	// Create server configuration
	config := &webserver.Config{
		Host:    host,
		Port:    serverPort,
		WebRoot: webRoot,
	}

	// Create and start the web server
	server, err := webserver.New(config)
	if err != nil {
		return fmt.Errorf("failed to create web server: %w", err)
	}

	// Set up graceful shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil {
			serverErrChan <- err
		}
	}()

	// Wait for shutdown signal or server error
	select {
	case sig := <-sigChan:
		fmt.Printf("\nReceived signal %v, shutting down server...\n", sig)
		return server.Shutdown()
	case err := <-serverErrChan:
		return fmt.Errorf("server error: %w", err)
	}
}

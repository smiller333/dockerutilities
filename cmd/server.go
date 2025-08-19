// Package cmd provides command-line interface functionality for dockerutils.
// Copyright (c) 2025 Docker Utils Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/smiller333/dockerutils/src/webserver"
	"github.com/spf13/cobra"
)

var (
	// Flags for the server command
	serverPort string
	webRoot    string
	host       string
	tmpDir     string
	noBrowser  bool
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Docker analysis web server",
	Long: `Start a local web server that provides Docker image analysis tools through both a web interface and REST API.

The server provides:
- Interactive web interface for analyzing Docker images in real-time
- Live Docker image analysis with detailed metadata extraction  
- REST API endpoints for programmatic access to analysis data
- Asynchronous image analysis operations with progress tracking
- File management and cleanup of analysis results
- Embedded modern UI with responsive design

Key Features:
- Analyze local and remote Docker images
- Extract and browse image filesystem contents
- View layer information and build history
- Generate comprehensive analysis reports
- Docker build context preview and .dockerignore validation

Examples:
  dockerutils server                                       # Start server on default port 8080 (opens browser automatically)
  dockerutils server --port 3000                          # Start server on port 3000
  dockerutils server --port 8080 --host 0.0.0.0           # Bind to all interfaces
  dockerutils server --web-root ./custom-ui               # Use custom web root directory
  dockerutils server --tmp-dir /app/data                  # Use custom tmp directory for analysis data
  dockerutils server --no-browser                         # Start server without opening browser automatically

The server will automatically open your default web browser to the analysis tools interface.
Use --no-browser to disable automatic browser opening.`,
	Args: cobra.NoArgs,
	RunE: runServer,
}

func init() {
	// Add the server command to the root command
	rootCmd.AddCommand(serverCmd)

	// Add command flags
	serverCmd.Flags().StringVar(&serverPort, "port", "8080", "Port number for the web server (default: 8080)")
	serverCmd.Flags().StringVar(&host, "host", "localhost", "Host/IP address to bind the server to (default: localhost)")
	serverCmd.Flags().StringVar(&webRoot, "web-root", "", "Custom path to web root directory for custom UI (optional)")
	serverCmd.Flags().StringVar(&tmpDir, "tmp-dir", "", "Custom path to temporary directory for analysis data (optional)")
	serverCmd.Flags().BoolVar(&noBrowser, "no-browser", false, "Disable automatic browser opening when server starts")
}

// runServer starts the web server for viewing Docker image analysis results
func runServer(cmd *cobra.Command, args []string) error {
	fmt.Printf("🚀 Starting Docker Utils Server...\n")
	fmt.Printf("📍 Server will be available at: http://%s:%s\n", host, serverPort)

	if webRoot != "" {
		fmt.Printf("📁 Using custom web root: %s\n", webRoot)
	}
	if tmpDir != "" {
		fmt.Printf("📁 Using custom temp directory: %s\n", tmpDir)
	}

	// Create server configuration
	config := &webserver.Config{
		Host:    host,
		Port:    serverPort,
		WebRoot: webRoot,
		TmpDir:  tmpDir,
	}

	// Create and start the web server
	server, err := webserver.New(config)
	if err != nil {
		return fmt.Errorf(`❌ Failed to create web server: %w

Troubleshooting steps:
1. Check if port %s is already in use: lsof -i :%s
2. Verify you have permission to bind to %s:%s
3. Try a different port: dockerutils server --port 8081
4. Check system resources and available memory

For more help, see: https://github.com/smiller333/dockerutils/docs/TROUBLESHOOTING.md`, err, serverPort, serverPort, host, serverPort)
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

	// Open browser automatically after a short delay (unless disabled)
	if !noBrowser {
		go func() {
			// Wait a moment for the server to start
			time.Sleep(1 * time.Second)
			url := fmt.Sprintf("http://%s:%s", host, serverPort)
			if err := openBrowser(url); err != nil {
				fmt.Printf("⚠️  Could not automatically open browser: %v\n", err)
				fmt.Printf("🌐 Please manually open: %s\n", url)
			} else {
				fmt.Printf("🌐 Opening browser to: %s\n", url)
			}
		}()
	} else {
		fmt.Printf("🌐 Browser auto-open disabled. Open manually: http://%s:%s\n", host, serverPort)
	}

	// Wait for shutdown signal or server error
	select {
	case sig := <-sigChan:
		fmt.Printf("\n🛑 Received signal %v, shutting down server gracefully...\n", sig)
		if err := server.Shutdown(); err != nil {
			return fmt.Errorf("❌ Error during server shutdown: %w", err)
		}
		fmt.Printf("✅ Server shutdown complete\n")
		return nil
	case err := <-serverErrChan:
		return fmt.Errorf("❌ Server error: %w", err)
	}
}

// openBrowser opens the specified URL in the default web browser
func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return cmd.Start()
}

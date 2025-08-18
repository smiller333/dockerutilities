# Docker Utils API Documentation

This document provides comprehensive API documentation for all packages in the Docker Utils project.

## Table of Contents

1. [Package Overview](#package-overview)
2. [Web Server API](#web-server-api)
3. [Docker Client Package](#docker-client-package)
4. [Analyzer Package](#analyzer-package)
5. [Build Context Package](#build-context-package)
6. [Version Package](#version-package)
7. [Error Handling](#error-handling)
8. [Examples](#examples)

## Package Overview

Docker Utils is organized into several focused packages:

| Package | Purpose | Key Features |
|---------|---------|--------------|
| `webserver` | Web interface and REST API | HTTP server, embedded UI, analysis endpoints |
| `dockerclient` | Docker SDK wrapper | Enhanced client with timeouts, error handling |
| `analyzer` | Image and file analysis | Image extraction, filesystem analysis, reporting |
| `buildcontext` | Build context management | Dockerfile validation, context analysis |
| `version` | Version management | Build-time version embedding, version strings |

## Web Server API

### Overview

The `webserver` package provides both a modern web interface and comprehensive REST API for Docker image analysis.

### Configuration

```go
type Config struct {
    Host    string // Host/IP address to bind to (default: "localhost")
    Port    string // Port number to listen on (default: "8080")
    WebRoot string // Custom web root directory (optional)
    TmpDir  string // Custom tmp directory for analysis data (optional)
}
```

### REST API Endpoints

#### Health Check
- **Endpoint**: `GET /api/health`
- **Description**: Returns server health status and build information
- **Response**:
```json
{
  "status": "healthy",
  "timestamp": "2025-01-21T14:47:10Z",
  "version": "v1.0.0",
  "git_commit": "a1b2c3d",
  "build_time": "2025-01-21T14:47:10Z",
  "go_version": "go1.24.2"
}
```

#### Image Summaries
- **Endpoint**: `GET /api/summaries`
- **Description**: Lists all analyzed images with metadata
- **Response**:
```json
[
  {
    "image_id": "sha256:7aab056cecc6...",
    "image_tag": "alpine:3.20",
    "image_source": "registry.example.com",
    "image_size": 7738912,
    "architecture": "amd64",
    "analyzed_at": "2025-01-21T14:47:10Z",
    "status": "completed",
    "request_id": "7aab056cecc6"
  }
]
```

#### Image Analysis Details
- **Endpoint**: `GET /api/info/{id}`
- **Description**: Returns complete analysis data for a specific image
- **Parameters**: 
  - `id` (string): Image ID or analysis request ID
- **Response**: Complete JSON analysis data including filesystem, layers, metadata

#### Delete Analysis Results
- **Endpoint**: `DELETE /api/info/{id}`
- **Description**: Removes analysis results and associated files
- **Parameters**:
  - `id` (string): Image ID or analysis request ID
- **Response**:
```json
{
  "success": true,
  "message": "Info 7aab056cecc6 deleted successfully"
}
```

#### Synchronous Image Analysis
- **Endpoint**: `POST /api/analyze`
- **Description**: Analyzes a Docker image synchronously (blocks until complete)
- **Request Body**:
```json
{
  "image_name": "nginx:latest",
  "keep_temp_files": true,
  "force_pull": false
}
```
- **Response**:
```json
{
  "success": true,
  "image_id": "7aab056cecc6",
  "message": "Image analysis completed successfully"
}
```

#### Asynchronous Image Analysis
- **Endpoint**: `POST /api/analyze-async`
- **Description**: Starts asynchronous image analysis (returns immediately)
- **Request Body**:
```json
{
  "image_name": "nginx:latest",
  "keep_temp_files": true,
  "force_pull": false
}
```
- **Response**:
```json
{
  "success": true,
  "request_id": "7aab056cecc6",
  "message": "Image analysis started"
}
```

### Server Management

#### Creating a Server
```go
config := &webserver.Config{
    Host: "localhost",
    Port: "8080",
    TmpDir: "./tmp",
}

server, err := webserver.New(config)
if err != nil {
    log.Fatal(err)
}
```

#### Starting the Server
```go
// Start server (blocking)
if err := server.Start(); err != nil {
    log.Fatal(err)
}

// Or start with graceful shutdown
go func() {
    if err := server.Start(); err != nil {
        log.Printf("Server error: %v", err)
    }
}()

// Handle shutdown signals
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
<-sigChan

server.Shutdown()
```

## Docker Client Package

### Overview

The `dockerclient` package provides an enhanced wrapper around the Docker SDK client with improved error handling, configuration management, and convenience methods.

### Configuration

```go
type Config struct {
    Host       string        // Docker daemon host
    APIVersion string        // Docker API version to use
    HTTPClient *http.Client  // Custom HTTP client
    Timeout    time.Duration // Default timeout for operations
    TLSVerify  bool         // Enable TLS verification
    CertPath   string       // Path to TLS certificates
}
```

### Client Creation

#### Default Client
```go
client, err := dockerclient.NewDefaultClient()
if err != nil {
    log.Fatal(err)
}
defer client.Close()
```

#### Custom Configuration
```go
config := &dockerclient.Config{
    Host:       "unix:///var/run/docker.sock",
    APIVersion: "1.41",
    Timeout:    30 * time.Second,
}

client, err := dockerclient.NewDockerClient(config)
if err != nil {
    log.Fatal(err)
}
defer client.Close()
```

### Core Operations

#### Connection Management
```go
// Test connection
ctx := context.Background()
if err := client.Ping(ctx); err != nil {
    log.Fatal("Cannot connect to Docker daemon:", err)
}

// Check connection status
isConnected := client.IsConnected(ctx)

// Get system information
info, err := client.GetInfo(ctx)
if err != nil {
    log.Fatal(err)
}
```

#### Image Operations
```go
// Pull an image
pullReader, err := client.PullImage(ctx, "nginx:latest", nil)
if err != nil {
    log.Fatal(err)
}
defer pullReader.Close()

// Inspect an image
imageInfo, err := client.InspectImage(ctx, "nginx:latest")
if err != nil {
    log.Fatal(err)
}

// Save an image to tar archive
saveReader, err := client.SaveImage(ctx, []string{"nginx:latest"})
if err != nil {
    log.Fatal(err)
}
defer saveReader.Close()

// Build an image from Dockerfile
buildReader, err := client.BuildImage(ctx, "/path/to/Dockerfile", "myapp:latest")
if err != nil {
    log.Fatal(err)
}
defer buildReader.Close()
```

#### Container Operations
```go
// Create a container
containerConfig := &container.Config{
    Image: "nginx:latest",
    Cmd:   []string{"nginx", "-g", "daemon off;"},
}

resp, err := client.CreateContainer(ctx, containerConfig, nil, nil, nil, "my-nginx")
if err != nil {
    log.Fatal(err)
}

// Copy files from container
reader, stat, err := client.CopyFromContainer(ctx, resp.ID, "/etc/nginx")
if err != nil {
    log.Fatal(err)
}
defer reader.Close()

// Remove container
err = client.RemoveContainer(ctx, resp.ID, true)
if err != nil {
    log.Fatal(err)
}
```

## Analyzer Package

### Overview

The `analyzer` package provides comprehensive Docker image and filesystem analysis capabilities with secure extraction and reporting.

### Core Types

#### AnalysisResult
```go
type AnalysisResult struct {
    ImageID       string  // ID of the Docker image
    ImageTag      string  // Tag of the analyzed image
    ImageSource   string  // Source registry
    ImageSize     int64   // Size of the image in bytes
    Architecture  string  // Image architecture (amd64, arm64, etc.)
    OS            string  // Operating system (linux, windows, etc.)
    Created       string  // Image creation timestamp
    Author        string  // Image author
    LayerCount    int     // Number of layers in the image
    
    // Analysis paths and status
    SavedTarPath    string // Path to saved tar file
    ExtractedPath   string // Path to extracted contents
    SaveSuccess     bool   // Tar save operation success
    ExtractSuccess  bool   // Extraction operation success
    
    // Container analysis
    ContainerID       string   // Created container ID
    ContainerName     string   // Container name
    ContainerSuccess  bool     // Container creation success
    ContainerWarnings []string // Container creation warnings
    ContainerFSPath   string   // Container filesystem path
    ContainerFSSuccess bool    // Container FS extraction success
    
    // Build information (for Dockerfile analysis)
    BuildOutput  string  // Docker build output
    BuildSuccess bool    // Build operation success
    BuildTime    float64 // Build time in seconds
}
```

### Key Functions

#### Safe Tar Extraction
```go
// SafeTarExtraction extracts tar archives using Go's native library
func SafeTarExtraction(tarPath string, destDir string) error
```

**Features:**
- Secure extraction with path traversal protection
- Support for gzipped tar files (.tar.gz, .tgz)
- Proper file permission handling
- Error handling for malformed archives

**Example:**
```go
err := analyzer.SafeTarExtraction("/path/to/image.tar", "/tmp/extracted")
if err != nil {
    log.Printf("Extraction failed: %v", err)
}
```

#### Image Processing
The package provides comprehensive image analysis functions for:
- Docker image inspection and metadata extraction
- Layer analysis and filesystem extraction
- Container creation and filesystem copying
- Build context analysis and validation

### Security Features

The analyzer package implements several security measures:
- **Path Traversal Protection**: Validates all file paths during extraction
- **Native Tar Handling**: Uses Go's archive/tar instead of external commands
- **File Permission Control**: Proper handling of file permissions and ownership
- **Resource Limits**: Configurable limits for extraction operations

## Build Context Package

### Overview

The `buildcontext` package provides functionality for analyzing and validating Docker build contexts, including Dockerfile parsing and .dockerignore handling.

### Key Features

- Dockerfile syntax validation
- Build context size analysis
- .dockerignore file processing
- Security scanning for build contexts
- Path validation and normalization

### Usage Example

```go
import "github.com/smiller333/dockerutils/src/buildcontext"

// Analyze a build context
result, err := buildcontext.Analyze("/path/to/build/context")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Build context size: %d bytes\n", result.Size)
fmt.Printf("Number of files: %d\n", result.FileCount)
```

## Version Package

### Overview

The `version` package manages build-time version information embedding and provides version string formatting.

### Build-time Variables

The following variables are set during build using ldflags:

```go
var (
    Version   = "dev"     // Set via -ldflags "-X path.Version=v1.0.0"
    GitCommit = "unknown" // Set via -ldflags "-X path.GitCommit=abc123"
    BuildTime = "unknown" // Set via -ldflags "-X path.BuildTime=2025-01-21T14:47:10Z"
)
```

### Functions

#### GetVersionString
```go
func GetVersionString() string
```
Returns a short version string for CLI usage.

**Example output:** `v1.0.0`

#### GetFullVersionString
```go
func GetFullVersionString() string
```
Returns detailed version information including git commit, build time, and Go version.

**Example output:**
```
dockerutils v1.0.0
Git Commit: a1b2c3d
Build Time: 2025-01-21T14:47:10Z
Go Version: go1.24.2
OS/Arch: darwin/arm64
```

### Build Integration

To embed version information during build:

```bash
VERSION=$(git describe --tags --always --dirty)
GIT_COMMIT=$(git rev-parse --short HEAD)
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

go build -ldflags "
  -X github.com/smiller333/dockerutils/src/version.Version=$VERSION
  -X github.com/smiller333/dockerutils/src/version.GitCommit=$GIT_COMMIT
  -X github.com/smiller333/dockerutils/src/version.BuildTime=$BUILD_TIME
" -o dockerutils
```

## Error Handling

### Standard Error Types

All packages follow Go's standard error handling conventions:

```go
// Docker client errors
var (
    ErrImageNotFound = errors.New("image not found")
    ErrConnectionFailed = errors.New("connection to Docker daemon failed")
)

// Analysis errors
var (
    ErrExtractionFailed = errors.New("tar extraction failed")
    ErrInvalidPath = errors.New("invalid file path detected")
)
```

### Error Wrapping

Errors are properly wrapped with context:

```go
if err != nil {
    return fmt.Errorf("failed to analyze image %s: %w", imageName, err)
}
```

### HTTP API Errors

The web server returns structured error responses:

```json
{
  "error": "Image not found",
  "message": "The specified image 'nonexistent:latest' could not be found",
  "code": 404
}
```

## Examples

### Complete Image Analysis Workflow

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/smiller333/dockerutils/src/analyzer"
    "github.com/smiller333/dockerutils/src/dockerclient"
)

func main() {
    // Create Docker client
    client, err := dockerclient.NewDefaultClient()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    
    // Test connection
    ctx := context.Background()
    if err := client.Ping(ctx); err != nil {
        log.Fatal("Cannot connect to Docker:", err)
    }
    
    // Pull and analyze an image
    imageName := "alpine:latest"
    
    // Pull the image
    pullReader, err := client.PullImage(ctx, imageName, nil)
    if err != nil {
        log.Fatal(err)
    }
    pullReader.Close()
    
    // Save image to tar
    saveReader, err := client.SaveImage(ctx, []string{imageName})
    if err != nil {
        log.Fatal(err)
    }
    
    // Extract and analyze
    tarPath := "/tmp/alpine.tar"
    destDir := "/tmp/alpine_extracted"
    
    // Save tar to file (implementation would go here)
    // ...
    
    // Extract safely
    err = analyzer.SafeTarExtraction(tarPath, destDir)
    if err != nil {
        log.Fatal("Extraction failed:", err)
    }
    
    fmt.Printf("Image %s analyzed successfully\n", imageName)
}
```

### Web Server Integration

```go
package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    
    "github.com/smiller333/dockerutils/src/webserver"
)

func main() {
    // Configure server
    config := &webserver.Config{
        Host:   "localhost",
        Port:   "8080",
        TmpDir: "./tmp",
    }
    
    // Create server
    server, err := webserver.New(config)
    if err != nil {
        log.Fatal(err)
    }
    
    // Start server in goroutine
    go func() {
        if err := server.Start(); err != nil {
            log.Printf("Server error: %v", err)
        }
    }()
    
    // Wait for shutdown signal
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan
    
    // Graceful shutdown
    if err := server.Shutdown(); err != nil {
        log.Printf("Shutdown error: %v", err)
    }
}
```

---

For more detailed examples and use cases, see the individual package README files and the main project documentation.

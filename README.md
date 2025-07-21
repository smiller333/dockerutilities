# Docker Utilities

A collection of Docker utilities providing a command-line interface for Docker analysis and management tasks.

## Overview

`dockerutils` is a CLI tool built in Go that provides comprehensive Docker image analysis capabilities through an interactive web interface. The tool uses the official Docker SDK to interact with Docker Engine and offers detailed insights into Docker images through a modern web-based interface.

### Key Features

- **Web Interface**: Interactive web server for viewing Docker image analysis results
- **Docker SDK Integration**: Built on the official Docker client library (v28.3.0+incompatible)
- **Image Analysis**: Inspect Docker images and extract their contents through the web interface
- **Comprehensive Reporting**: Generate detailed analysis reports with build metrics and image metadata

## Prerequisites

- Go 1.24.2 or later
- Git
- Docker Engine running locally
- Access to Docker socket (typically `/var/run/docker.sock` on Unix systems)
- On Unix systems, ensure your user is in the `docker` group or has appropriate Docker permissions

## Quick Start

```bash
# Clone the repository
git clone https://github.com/smiller333/dockerutils.git
cd dockerutils

# Quick build
./build.sh dev

# Start web interface for viewing and analyzing Docker images
./bin/dockerutils image-viewer --port 8080
```

## Installation

### Building from Source

This project uses a sophisticated build system that embeds version information at compile time.

#### Quick Build (Recommended)

```bash
# Clone the repository
git clone https://github.com/smiller333/dockerutils.git
cd dockerutils

# Build development version
./build.sh dev

# Build optimized release version  
./build.sh release

# Build with specific version
./build.sh -v v1.0.0 release

# Show version information
./build.sh version
```

#### Using Makefile

```bash
# Development build (with debug info)
make build-dev

# Optimized release build
make build-release

# Build for all platforms
make build-all

# Run tests
make test

# Display version information
make version

# Clean build artifacts
make clean
```

#### Manual Build

```bash
# Basic build (no version embedding)
go build -o dockerutils

# Build with version information
VERSION=$(git describe --tags --always --dirty)
GIT_COMMIT=$(git rev-parse --short HEAD)
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

go build -ldflags "-X github.com/smiller333/dockerutils/src/version.Version=$VERSION -X github.com/smiller333/dockerutils/src/version.GitCommit=$GIT_COMMIT -X github.com/smiller333/dockerutils/src/version.BuildTime=$BUILD_TIME" -o dockerutils
```

#### Installation

```bash
# Install to GOPATH/bin or GOBIN
make install

# Or manually copy to PATH
cp bin/dockerutils /usr/local/bin/
```

### Version Information

The built binary includes comprehensive build information:

```bash
# Short version
dockerutils -v
# Output: dockerutils v1.0.0 (built 2025-07-21 14:47:10 UTC)

# Detailed version information
dockerutils version
# Output:
# dockerutils v1.0.0
# Git Commit: a1b2c3d
# Build Time: 2025-07-21 14:47:10 UTC
# Go Version: go1.24.2
# OS/Arch: darwin/arm64
```

## Output and Data Management

### Temporary Files

The web server can generate temporary files and directories during image analysis:
- **Image analysis**: Creates temporary directories under `tmp/` with extracted image contents
- **Analysis reports**: Generates JSON summary files with detailed analysis results
- **Cleanup**: Temporary files can be managed through the web interface

### Docker Permissions

The tool requires access to the Docker daemon to perform analysis operations through the web interface. Ensure that:
- Docker daemon is running
- Your user has permission to access the Docker socket
- On Unix systems, you may need to add your user to the `docker` group or run with appropriate permissions

## Usage

### Available Commands

- `version` - Print the version number of dockerutils
- `image-viewer` - Start a web server for viewing Docker image analysis results
- `completion` - Generate autocompletion scripts for various shells

### Basic Commands

```bash
# Display help
dockerutils --help

# Show version
dockerutils version

# Start web server for viewing analysis results
dockerutils image-viewer --port 8080
```

### Image Viewer Command

The `image-viewer` command starts a local web server that provides comprehensive Docker image analysis capabilities through an interactive web interface:

**Features:**
- Interactive web interface for analyzing Docker images in real-time
- Live Docker image analysis with detailed metadata extraction
- Static file serving for analysis reports and extracted contents
- REST API endpoints for accessing image summaries and filesystem data
- Real-time visualization of Docker image layers and contents
- Capability to analyze both local and remote Docker images

**Options:**
- `--port <port>` - Port to run the web server on (default: 8080)
- `--host <host>` - Host to bind the server to (default: localhost)
- `--web-root <path>` - Root directory for serving web files

**Web Interface Usage:**
1. Start the web server with `dockerutils image-viewer --port 8080`
2. Open your browser to `http://localhost:8080`
3. Use the web interface to analyze Docker images interactively
4. Browse existing analysis results if any are available in the temp directory

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

## Project Structure

```
dockerutils/
├── .github/             # GitHub configuration and documentation
│   └── copilot-instructions.md  # GitHub Copilot configuration
├── cmd/                 # Command definitions
│   ├── root.go         # Root command and CLI setup
│   ├── completion.go   # Shell completion command
│   └── image-viewer.go # Web server command for viewing results
├── src/                 # Implementation logic (separated by concern)
│   ├── analyzer/       # Dockerfile and image analysis logic
│   │   ├── analyzer.go
│   │   ├── fileprocessor.go
│   │   ├── fileprocessor_test.go
│   │   ├── imageprocessor.go
│   │   └── imageprocessor_test.go
│   ├── dockerclient/   # Docker SDK client wrapper
│   │   ├── client.go
│   │   ├── client_test.go
│   │   └── README.md
│   ├── version/        # Version management
│   │   ├── version.go  # Version constants and functions
│   │   └── version_test.go
│   └── webserver/      # Web server for viewing analysis results
│       ├── server.go
│       ├── server_test.go
│       ├── README.md
│       └── webpages/   # HTML templates and static files
├── docs/               # Documentation
│   ├── apis/           # API documentation
│   │   └── dockersdk/  # Docker SDK API references
│   └── examples/       # Example Dockerfiles
├── tmp/                # Temporary analysis outputs
├── bin/                # Built binaries
├── .gitignore          # Git ignore patterns
├── build.sh            # Build script with version injection
├── Makefile            # Build automation
├── main.go             # Application entry point
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
└── README.md
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) v1.9.1 - CLI framework for Go
- [Docker SDK for Go](https://github.com/docker/docker) v28.3.0+incompatible - Official Docker client library
- [OpenContainers Image Spec](https://github.com/opencontainers/image-spec) v1.1.1 - OCI image specification support
- [golang.org/x/text](https://golang.org/x/text) v0.26.0 - Additional text processing utilities

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

### Development Guidelines

1. Follow Go best practices and conventions
2. Write unit tests for new functionality
3. Update documentation as needed
4. Ensure all tests pass before submitting PRs
5. Use the provided build scripts for consistent builds

## License

This project is open source.

## Version Information

Current version: Based on git commit hash (e.g., aeaea31)

The version is automatically determined from git information during build time. Use `dockerutils version` to see the current build details.
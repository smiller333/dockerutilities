# Docker Utilities

A collection of Docker utilities providing a command-line interface for various Docker-related tasks.

## Overview

`dockerutils` is a CLI tool built in Go that provides comprehensive analysis capabilities for Docker containers and images. The tool uses the official Docker SDK to interact with Docker Engine and offers detailed insights into Dockerfile builds and existing Docker images.

Key features:
- **Dockerfile Analysis**: Parse, build, and analyze Dockerfiles with detailed metrics
- **Image Analysis**: Inspect existing Docker images and extract their contents
- **Docker SDK Integration**: Built on the official Docker client library (v28.3.0+incompatible)
- **Comprehensive Reporting**: Generate detailed analysis reports with build metrics and image metadata

## Installation

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/smiller333/dockerutils.git
   cd dockerutils
   ```

2. Build the binary:
   ```bash
   go build -o dockerutils
   ```

3. (Optional) Move to your PATH:
   ```bash
   mv dockerutils /usr/local/bin/
   ```

## Usage

### Basic Commands

```bash
# Display help
dockerutils --help

# Show version
dockerutils version

# Analyze a Dockerfile
dockerutils analyze --dockerfile ./Dockerfile

# Analyze a Docker image
dockerutils analyze --image nginx:latest

# Analyze with build output (for Dockerfiles)
dockerutils analyze --dockerfile ./Dockerfile --build-output

# Analyze and keep temporary files (for images)
dockerutils analyze --image alpine:3.20 --keep-temp
```

### Available Commands

- `version` - Print the version number of dockerutils
- `analyze` - Analyze a Dockerfile or Docker image to understand its structure and contents

#### Analyze Command

The `analyze` command provides comprehensive analysis of Dockerfiles and Docker images:

**Dockerfile Analysis:**
- Parses and validates Dockerfile syntax
- Builds the Docker image from the Dockerfile
- Reports image size, layer count, and build time
- Optionally displays build output for debugging

**Image Analysis:**
- Inspects existing Docker images
- Extracts image metadata (architecture, OS, creation date)
- Saves and extracts image contents for examination
- Creates temporary containers for filesystem analysis
- Generates detailed analysis reports

**Options:**
- `--dockerfile <path>` - Analyze a Dockerfile at the specified path
- `--image <name:tag>` - Analyze an existing Docker image
- `--build-output` - Show Docker build output (Dockerfile analysis only)
- `--keep-temp` - Keep temporary files after analysis (useful for debugging)

## Development

### Prerequisites

- Go 1.24.2 or later
- Git
- Docker Engine running locally
- Access to Docker socket (typically `/var/run/docker.sock` on Unix systems)

### Running Tests

```bash
go test ./...
```

### Output and Temporary Files

The `analyze` command generates temporary files and directories during analysis:
- **Image analysis**: Creates temporary directories under `tmp/` with extracted image contents
- **Analysis reports**: Generates JSON summary files with detailed analysis results
- **Cleanup**: Temporary files are automatically cleaned up unless `--keep-temp` flag is used

## Project Structure

```
dockerutils/
├── .github/             # GitHub configuration and documentation
│   └── copilot-instructions.md  # GitHub Copilot configuration
├── cmd/                 # Command definitions
│   ├── root.go         # Root command and CLI setup
│   └── analyze.go      # Analyze command for Dockerfiles and images
├── src/                 # Implementation logic (separated by concern)
│   ├── analyzer/       # Dockerfile and image analysis logic
│   │   ├── analyzer.go
│   │   ├── fileprocessor.go
│   │   ├── fileprocessor_test.go
│   │   └── imageprocessor.go
│   ├── dockerclient/   # Docker SDK client wrapper
│   │   ├── client.go
│   │   ├── client_test.go
│   │   └── README.md
│   └── version/        # Version management
│       ├── version.go  # Version constants and functions
│       └── version_test.go
├── docs/               # Documentation
│   ├── apis/           # API documentation
│   │   └── dockersdk/  # Docker SDK API references
│   └── examples/       # Example Dockerfiles
├── tmp/                # Temporary analysis outputs
├── .gitignore          # Git ignore patterns
├── main.go             # Application entry point
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
└── README.md
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) v1.9.1 - CLI framework for Go
- [Docker SDK for Go](https://github.com/docker/docker) v28.3.0+incompatible - Official Docker client library
- [OpenContainers Image Spec](https://github.com/opencontainers/image-spec) v1.1.1 - OCI image specification support

## Docker Permissions

The tool requires access to the Docker daemon to perform analysis operations. Ensure that:
- Docker daemon is running
- Your user has permission to access the Docker socket
- On Unix systems, you may need to add your user to the `docker` group or run with appropriate permissions

## License

This project is open source.

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## Version

Current version: v0.0.1
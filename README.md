# Docker Utilities

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/smiller333/dockerutils)](https://goreportcard.com/report/github.com/smiller333/dockerutils)
[![Go Version](https://img.shields.io/badge/go-1.24.2+-blue.svg)](https://golang.org/)

A powerful CLI tool and web interface for comprehensive Docker image analysis and management. Built with Go and the official Docker SDK, it provides detailed insights into Docker images through both command-line tools and an interactive web interface.

## 🚀 Quick Start

Get up and running in under 30 seconds:

```bash
# Using Docker (recommended)
git clone https://github.com/smiller333/dockerutils.git
cd dockerutils
./scripts/docker-run.sh run-persistent

# Access the web interface at http://localhost:8080
# Try analyzing: alpine:latest, nginx:latest, or node:18-alpine
```

> **New to Docker Utils?** Check out our [Installation Guide](INSTALLATION.md) for platform-specific instructions, or try our [hands-on examples](EXAMPLES.md).

## ✨ Key Features

- **🖥️ Interactive Web Interface**: Modern, responsive UI for Docker image analysis
- **🔍 Deep Image Analysis**: Extract and analyze Docker image contents, layers, and metadata
- **⚡ Live Analysis**: Real-time Docker image inspection with detailed reports
- **📊 Comprehensive Reporting**: Generate detailed analysis reports with build metrics
- **🐳 Docker SDK Integration**: Built on the official Docker client library (v28.3.0+)
- **🛠️ CLI Commands**: Command-line interface for automation and scripting
- **📁 File System Explorer**: Browse extracted image contents through the web interface
- **🔄 Asynchronous Operations**: Non-blocking image analysis with progress tracking

## Prerequisites

- Go 1.24.2 or later
- Git
- Docker Engine running locally
- Access to Docker socket (typically `/var/run/docker.sock` on Unix systems)
- On Unix systems, ensure your user is in the `docker` group or has appropriate Docker permissions
- **Rancher Desktop users**: Enable "Administrative access" in Preferences → General to ensure proper Docker socket access

## Quick Start

### Using Docker (Recommended)

The easiest way to run dockerutils is using Docker:

```bash
# Clone the repository
git clone https://github.com/smiller333/dockerutils.git
cd dockerutils

# Build and run with our convenience script
./scripts/docker-run.sh run-persistent

# Or build and run manually
docker build -t dockerutils-viewer .
docker run -d --name dockerutils-viewer \
  -p 8080:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(pwd)/data:/app/data \
  dockerutils-viewer

# Access the web interface at http://localhost:8080
```

For more Docker deployment options, see [DOCKER.md](DOCKER.md).

> **💡 Need help?** Our [Troubleshooting Guide](TROUBLESHOOTING.md) covers common issues and solutions.

### Building from Source

This project uses a sophisticated build system that embeds version information at compile time.

#### Quick Build (Recommended)

```bash
# Clone the repository
git clone https://github.com/smiller333/dockerutils.git
cd dockerutils

# Build development version
./scripts/build.sh dev

# Build optimized release version  
./scripts/build.sh release

# Build with specific version
./scripts/build.sh -v v1.0.0 release

# Show version information
./scripts/build.sh version
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

## 🔧 Usage

### Basic Commands

```bash
# Display help and available commands
dockerutils --help

# Show detailed version information
dockerutils version

# Start the web interface (automatically opens browser)
dockerutils server

# Start on custom port
dockerutils server --port 3000

# Start without opening browser automatically
dockerutils server --no-browser

# Start with custom configuration
dockerutils server --port 8080 --host 0.0.0.0 --tmp-dir /app/data
```

### Web Interface Features

Once the server is running, you can:

1. **🔍 Analyze Images**: Enter any Docker image name (e.g., `nginx:latest`, `alpine:3.20`)
2. **📊 View Analysis Results**: Browse detailed image metadata, layers, and filesystem contents
3. **📁 Explore File Systems**: Navigate through extracted image contents
4. **🗂️ Manage Results**: Delete old analysis results to free up space
5. **⚡ Live Analysis**: Perform real-time image analysis through the web UI

### Common Use Cases

#### Analyzing a Public Image
```bash
# Start the web interface
dockerutils server

# Then in the web interface:
# 1. Enter "nginx:latest" in the image name field
# 2. Click "Analyze Image"
# 3. Browse the results including filesystem, layers, and metadata
```

#### Analyzing a Private Registry Image
```bash
# Ensure you're logged into your registry first
docker login your-registry.com

# Start dockerutils
dockerutils server

# In the web interface, analyze:
# your-registry.com/your-org/your-app:v1.0.0
```

#### Batch Analysis via API
```bash
# Start the server
dockerutils server &

# Use the REST API for automation
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{"image_name": "alpine:latest"}'

# Check analysis results
curl http://localhost:8080/api/summaries
```

### API Reference

The web server provides a comprehensive REST API for programmatic access:

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/health` | GET | Server health status and version info |
| `/api/summaries` | GET | List all analyzed images with metadata |
| `/api/info/{id}` | GET | Get full analysis data for specific image |
| `/api/info/{id}` | DELETE | Remove analysis results for specific image |
| `/api/analyze` | POST | Analyze image synchronously (blocks until complete) |
| `/api/analyze-async` | POST | Start asynchronous image analysis |

#### Example API Usage

```bash
# Check server health
curl http://localhost:8080/api/health

# List all analyzed images
curl http://localhost:8080/api/summaries

# Analyze an image (synchronous)
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "image_name": "alpine:latest",
    "keep_temp_files": true,
    "force_pull": false
  }'

# Start async analysis
curl -X POST http://localhost:8080/api/analyze-async \
  -H "Content-Type: application/json" \
  -d '{"image_name": "nginx:latest"}'

# Get detailed analysis results
curl http://localhost:8080/api/info/7aab056cecc6
```

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
│   └── server.go       # Web server command for analysis tools
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
│   ├── USER_GUIDE.md   # Comprehensive user guide
│   ├── API.md          # REST API documentation
│   ├── CLI_REFERENCE.md # Command-line reference
│   ├── apis/           # API documentation
│   │   └── dockersdk/  # Docker SDK API references
│   └── examples/       # Example Dockerfiles and tutorials
├── INSTALLATION.md     # Platform-specific installation guide
├── TROUBLESHOOTING.md  # Common issues and solutions
├── EXAMPLES.md         # Hands-on examples and tutorials
├── DOCKER.md           # Docker deployment guide
├── tmp/                # Temporary analysis outputs
├── bin/                # Built binaries
├── .gitignore          # Git ignore patterns
├── scripts/            # Build and utility scripts
│   ├── build.sh       # Build script with version injection
│   └── docker-run.sh  # Docker container management script
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

## 🔧 Troubleshooting

### Quick Fixes

#### Docker Permission Denied
```bash
# Add your user to the docker group (Linux/macOS)
sudo usermod -aG docker $USER
# Then logout and login again
```

#### Port Already in Use
```bash
# Use a different port
dockerutils server --port 3000
```

#### Analysis Fails
```bash
# Check Docker is running
docker info

# Verify image exists
docker pull alpine:latest
```

> **📋 For comprehensive troubleshooting:** See our detailed [Troubleshooting Guide](TROUBLESHOOTING.md) which covers installation issues, Docker problems, API errors, and performance optimization.

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Quick Development Setup

```bash
# Clone and setup
git clone https://github.com/smiller333/dockerutils.git
cd dockerutils

# Install dependencies
go mod download

# Run tests
make test

# Build development version
make build-dev

# Run locally
./bin/dockerutils server
```

### Development Guidelines

1. **Code Quality**: Follow Go best practices and run `golangci-lint`
2. **Testing**: Write unit tests for new functionality (`make test`)
3. **Documentation**: Update documentation for any API changes
4. **Commits**: Use conventional commit messages
5. **Pull Requests**: Ensure all CI checks pass before submitting

## 📚 Documentation

Our comprehensive documentation covers everything from installation to advanced usage:

### Getting Started
- **[Installation Guide](INSTALLATION.md)** - Platform-specific installation instructions for Linux, macOS, Windows
- **[User Guide](docs/USER_GUIDE.md)** - Complete guide to using Docker Utils effectively
- **[Examples](EXAMPLES.md)** - Hands-on tutorials and real-world usage scenarios

### Reference Documentation
- **[API Reference](docs/API.md)** - Complete REST API documentation with examples
- **[CLI Reference](docs/CLI_REFERENCE.md)** - Command-line interface documentation
- **[Docker Deployment](DOCKER.md)** - Containerized deployment options

### Troubleshooting & Support
- **[Troubleshooting Guide](TROUBLESHOOTING.md)** - Common issues, solutions, and performance tips
- **[Contributing Guide](CONTRIBUTING.md)** - How to contribute to the project
- **[Code of Conduct](CODE_OF_CONDUCT.md)** - Community guidelines

### Quick Links
| Task | Documentation |
|------|---------------|
| First-time setup | [Installation Guide](INSTALLATION.md) |
| Learn the basics | [User Guide](docs/USER_GUIDE.md) |
| Try examples | [Examples](EXAMPLES.md) |
| Deploy with Docker | [Docker Guide](DOCKER.md) |
| Use the API | [API Reference](docs/API.md) |
| Troubleshoot issues | [Troubleshooting](TROUBLESHOOTING.md) |
| Contribute | [Contributing](CONTRIBUTING.md) |

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🏷️ Version Information

Version information is automatically embedded during build:

```bash
# Quick version check
dockerutils -v
# Output: dockerutils v1.0.0 (built 2025-01-21 14:47:10 UTC)

# Detailed version information
dockerutils version
# Output:
# dockerutils v1.0.0
# Git Commit: a1b2c3d
# Build Time: 2025-01-21 14:47:10 UTC
# Go Version: go1.24.2
# OS/Arch: darwin/arm64
```

## 🌟 Star History

If you find this project helpful, please consider giving it a star! ⭐

## 🔗 Related Projects

- [Docker Official Documentation](https://docs.docker.com/)
- [Docker SDK for Go](https://github.com/docker/docker)
- [Dive](https://github.com/wagoodman/dive) - Another Docker image explorer

---

**Made with ❤️ by the Docker Utils community**
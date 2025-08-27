# Docker Utilities

> **⚠️ Hobby Project Notice**: This is a personal hobby project created for learning and experimentation purposes. The code is provided "as is" and may not receive regular updates or maintenance. Use at your own discretion.

A Go-based CLI tool and web interface for Docker image analysis and management. Built with the official Docker SDK, it provides insights into Docker images through both command-line tools and an interactive web interface.

## ⚠️ Security Warning

**This application should ONLY be run on local development systems and trusted networks.**

- The web interface binds to `localhost` by default but can be configured to bind to other interfaces
- The application requires access to the Docker socket (`/var/run/docker.sock`) which provides full access to the Docker daemon
- Running this application on public networks or production systems without proper security measures could expose your system to unauthorized access
- This tool is intended for development and analysis purposes only

**Do not deploy this application to production environments or expose it to untrusted networks.**

## 🚀 Quick Start

### Option 1: Download Prebuilt Binaries (Recommended)

Prebuilt binaries are available for multiple platforms on the [GitHub Releases page](https://github.com/smiller333/dockerutilities/releases).

#### Installation Steps:

1. **Download the appropriate binary** for your platform from the [releases page](https://github.com/smiller333/dockerutilities/releases)
2. **Set up your environment**:
   ```bash
   # Create bin directory (if it doesn't exist)
   mkdir -p ~/bin
   
   # Move binary to bin directory
   mv dockerutilities-* ~/bin/
   
   # Make executable (Linux/macOS)
   chmod +x ~/bin/dockerutilities-*
   
   # Add to PATH (add to your shell profile if needed)
   export PATH="$HOME/bin:$PATH"
   ```
3. **Start the web interface**: `dockerutilities server`
4. **Access at**: http://localhost:8080

#### Platform Support:
- **Linux**: x86_64, ARM64
- **macOS**: Intel (x86_64), Apple Silicon (ARM64)  
- **Windows**: x86_64, ARM64 (ZIP format)

### Option 2: Build from Source

```bash
# Clone and build
git clone https://github.com/smiller333/dockerutilities.git
cd dockerutilities

# Build the application (development)
make dev

# Start the web interface
./bin/dockerutilities server

# Access at http://localhost:8080
```

## ✨ Features

- **🖥️ Web Interface**: Interactive UI for Docker image analysis
- **🔍 Image Analysis**: Extract and analyze Docker image contents and layers
- **📊 Detailed Reports**: Generate analysis reports with build metrics
- **🐳 Docker SDK Integration**: Built on the official Docker client library
- **🛠️ CLI Commands**: Command-line interface for automation
- **🔒 Security Features**: File size limits, path traversal protection, secure extraction
- **📦 Automated Releases**: GoReleaser integration for cross-platform builds
- **🔄 CI/CD Pipeline**: Automated testing, linting, and validation
- **🐳 Docker Support**: Multi-stage Docker builds with security hardening

## Prerequisites

- Go 1.24.2 or later
- Docker Engine running locally
- Access to Docker socket (`/var/run/docker.sock`)
- **Local development environment only** - Do not run on production systems

## Usage

### Web Interface

```bash
# Start the server
./bin/dockerutilities server

# Start on custom port
./bin/dockerutilities server --port 3000

# Set file size limit (default: 100MB)
./bin/dockerutilities server --max-file-size 200MB

# Bind to all interfaces (use with caution)
./bin/dockerutilities server --host 0.0.0.0

# Use custom temporary directory
./bin/dockerutilities server --tmp-dir /app/data
```

Once running, you can:
1. Enter any Docker image name (e.g., `nginx:latest`, `alpine:3.20`)
2. View detailed analysis results including metadata, layers, and filesystem
3. Browse extracted image contents (with file size limits)
4. Manage analysis results
5. Check Docker connectivity status via the "System Status" button

### CLI Commands

```bash
# Show help
./bin/dockerutilities --help

# Show version
./bin/dockerutilities version

# Start web server
./bin/dockerutilities server
```

### API Endpoints

The web server provides a REST API:

- `GET /api/health` - Server health status
- `GET /api/summaries` - List analyzed images
- `GET /api/info/{id}` - Get analysis data for specific image
- `POST /api/analyze` - Analyze image synchronously
- `POST /api/analyze-async` - Start asynchronous analysis

## Project Structure

```
dockerutilities/
├── cmd/                 # CLI command definitions
├── src/                 # Implementation logic
│   ├── analyzer/       # Image analysis logic
│   ├── buildcontext/   # Docker build context processing
│   ├── dockerclient/   # Docker SDK wrapper
│   ├── version/        # Version management
│   └── webserver/      # Web server implementation
├── scripts/            # Build and utility scripts
├── docs/               # Documentation
├── .github/workflows/  # CI/CD workflows
├── bin/                # Built binaries
└── tmp/                # Temporary files
```

## Development

### Building

```bash
# Development build (with debug info)
make dev

# Production build (optimized)
make release

# Local development build (verbose)
make local

# Clean build artifacts
make clean
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run smoke tests
./scripts/smoke-tests.sh

# Verify binary integrity
./scripts/verify-binaries.sh
```

### Code Quality

```bash
# Run linter
golangci-lint run

# Format code
go fmt ./...

# Vet code
go vet ./...
```

### Docker Development

```bash
# Build Docker image
make docker-build

# Run Docker container
make docker-run

# Use Docker run script with options
./scripts/docker-run.sh --persistent
```

## CI/CD Pipeline

The project includes automated CI/CD workflows:

### Continuous Integration (.github/workflows/ci.yml)
- **Testing**: Comprehensive test suite with coverage reporting
- **Linting**: Code quality checks with golangci-lint
- **Validation**: Binary integrity and functionality verification
- **Security**: Dependency scanning and vulnerability checks

### Release Automation (.github/workflows/release.yml)
- **GoReleaser Integration**: Automated releases on git tags
- **Cross-Platform Builds**: Linux, macOS, Windows (amd64/arm64)
- **GitHub Releases**: Automatic release creation with artifacts
- **Changelog Generation**: Automated changelog from commit messages

### Dependabot Integration
- **Automated Updates**: Weekly dependency updates
- **Security Focus**: Prioritizes security updates
- **Version Limits**: Patch and minor version updates only

## GoReleaser Integration

The project uses GoReleaser for automated release management:

### Features
- **Cross-Platform Builds**: 6 platforms (Linux/macOS/Windows amd64/arm64)
- **Version Injection**: Embeds git tag version information
- **Automated Changelog**: Generates changelogs from commit messages
- **GitHub Releases**: Creates releases with downloadable artifacts
- **Checksums**: Provides SHA256 checksums for all binaries

### Testing GoReleaser Locally
```bash
# Test GoReleaser configuration
./scripts/test-goreleaser.sh

# Test with development tag
./scripts/test-goreleaser.sh --dev

# Test with production tag
./scripts/test-goreleaser.sh --release
```

## Security Features

### File Extraction Security
- **File Size Limits**: Configurable maximum file size (default: 100MB)
- **Path Traversal Protection**: Prevents directory traversal attacks
- **Secure Extraction**: Validates file paths and permissions
- **Resource Limits**: Prevents resource exhaustion attacks

### Docker Security
- **Non-Root User**: Docker containers run as non-root user
- **Socket Permissions**: Dynamic Docker socket group assignment
- **Minimal Runtime**: Alpine-based minimal runtime environment
- **Security Scanning**: Automated vulnerability scanning in CI

## Documentation

- **[Docker Build Process](docs/docker-build-process.md)**: Detailed Docker build documentation
- **[Release Notes Template](docs/release-notes-template.md)**: Guidelines for release notes
- **[DOCKER.md](DOCKER.md)**: Docker-specific usage and configuration

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Docker SDK for Go](https://github.com/docker/docker) - Official Docker client
- [OpenContainers Image Spec](https://github.com/opencontainers/image-spec) - OCI support
- [GoReleaser](https://goreleaser.com/) - Release automation
- [golangci-lint](https://golangci-lint.run/) - Code quality and linting

## License

MIT License - see [LICENSE](LICENSE) file for details.
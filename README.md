# Docker Utilities

A Go-based CLI tool and web interface for Docker image analysis and management. Built with the official Docker SDK, it provides insights into Docker images through both command-line tools and an interactive web interface.

## 🚀 Quick Start

```bash
# Clone and build
git clone https://github.com/smiller333/dockerutilities.git
cd dockerutilities

# Make scripts executable (required for first-time setup)
chmod +x scripts/*.sh

# Build the application
./scripts/build.sh dev

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

## Prerequisites

- Go 1.24.2 or later
- Docker Engine running locally
- Access to Docker socket (`/var/run/docker.sock`)

## Usage

### Web Interface

```bash
# Start the server
./bin/dockerutilities server

# Start on custom port
./bin/dockerutilities server --port 3000
```

Once running, you can:
1. Enter any Docker image name (e.g., `nginx:latest`, `alpine:3.20`)
2. View detailed analysis results including metadata, layers, and filesystem
3. Browse extracted image contents
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
│   ├── dockerclient/   # Docker SDK wrapper
│   ├── version/        # Version management
│   └── webserver/      # Web server implementation
├── scripts/            # Build scripts
└── bin/                # Built binaries
```

## Development

```bash
# Run tests
go test ./...

# Build development version
./scripts/build.sh dev

# Build release version
./scripts/build.sh release
```

> **Note**: If you're setting up the project for the first time, make sure to run `chmod +x scripts/*.sh` to make the build scripts executable.

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Docker SDK for Go](https://github.com/docker/docker) - Official Docker client
- [OpenContainers Image Spec](https://github.com/opencontainers/image-spec) - OCI support

## License

MIT License - see [LICENSE](LICENSE) file for details.
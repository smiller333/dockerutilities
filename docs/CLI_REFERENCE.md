# Docker Utils CLI Reference

Complete command-line interface reference for Docker Utils.

## Table of Contents

1. [Overview](#overview)
2. [Global Options](#global-options)
3. [Commands](#commands)
4. [Shell Completion](#shell-completion)
5. [Examples](#examples)
6. [Environment Variables](#environment-variables)

## Overview

Docker Utils provides a command-line interface for Docker image analysis and management through a web-based interface and REST API.

### Basic Syntax

```
dockerutils [global-options] <command> [command-options] [arguments]
```

### Getting Help

```bash
# General help
dockerutils --help

# Command-specific help
dockerutils <command> --help

# List all available commands
dockerutils help
```

## Global Options

These options are available for all commands:

| Option | Short | Description |
|--------|-------|-------------|
| `--help` | `-h` | Show help information |
| `--version` | `-v` | Show version information |

### Version Information

```bash
# Quick version
dockerutils -v
# Output: dockerutils v1.0.0 (built 2025-01-21 14:47:10 UTC)

# Detailed version information
dockerutils version
# Output:
# dockerutils v1.0.0
# Git Commit: a1b2c3d
# Build Time: 2025-01-21T14:47:10 UTC
# Go Version: go1.24.2
# OS/Arch: darwin/arm64
```

## Commands

### server

Start the Docker analysis web server with an interactive interface and REST API.

#### Syntax
```bash
dockerutils server [flags]
```

#### Description
Starts a local web server that provides comprehensive Docker image analysis capabilities through both a web interface and REST API. The server includes:

- Interactive web interface for real-time Docker image analysis
- Live Docker image analysis with detailed metadata extraction
- REST API endpoints for programmatic access to analysis data
- Asynchronous image analysis operations with progress tracking
- File management and cleanup of analysis results
- Embedded modern UI with responsive design

#### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--host` | string | `localhost` | Host/IP address to bind the server to |
| `--port` | string | `8080` | Port number for the web server |
| `--tmp-dir` | string | `./tmp` | Custom path to tmp directory for analysis data |
| `--web-root` | string | (empty) | Custom path to web root directory (optional) |
| `--no-browser` | bool | `false` | Don't automatically open browser when server starts |

#### Examples

```bash
# Start server with default settings (localhost:8080)
dockerutils server

# Start server on custom port
dockerutils server --port 3000

# Bind to all interfaces
dockerutils server --host 0.0.0.0 --port 8080

# Use custom directories
dockerutils server --tmp-dir /app/data --web-root ./custom-ui

# Start without opening browser automatically
dockerutils server --no-browser

# Production-like setup
dockerutils server --host 0.0.0.0 --port 8080 --tmp-dir /var/lib/dockerutils --no-browser
```

#### Key Features

- **Analyze Images**: Support for local and remote Docker images
- **Extract Contents**: Browse image filesystem contents interactively
- **Layer Analysis**: View detailed layer information and build history
- **Comprehensive Reports**: Generate detailed analysis reports with metadata
- **Build Context**: Preview Docker build context and validate .dockerignore files

#### Server Behavior

1. **Automatic Browser Opening**: By default, opens your default web browser to the analysis interface
2. **Graceful Shutdown**: Handles SIGINT (Ctrl+C) and SIGTERM signals for clean shutdown
3. **Request Logging**: All HTTP requests are logged with timing information
4. **Auto-discovery**: Automatically finds and serves existing analysis results from the tmp directory

### version

Print detailed version information for Docker Utils.

#### Syntax
```bash
dockerutils version
```

#### Description
Displays comprehensive build and version information including:
- Version number (from git tags or "dev" for development builds)
- Git commit hash
- Build timestamp
- Go compiler version
- Operating system and architecture

#### Examples

```bash
dockerutils version
```

**Sample Output:**
```
dockerutils v1.0.0
Git Commit: a1b2c3d
Build Time: 2025-01-21T14:47:10Z
Go Version: go1.24.2
OS/Arch: darwin/arm64
```

### completion

Generate shell autocompletion scripts for Docker Utils.

#### Syntax
```bash
dockerutils completion [shell]
```

#### Description
Generates autocompletion scripts for the specified shell. Supports bash, zsh, fish, and PowerShell. The completion script enables tab completion for commands, flags, and arguments.

#### Arguments

| Shell | Description |
|-------|-------------|
| `bash` | Generate autocompletion script for bash |
| `zsh` | Generate autocompletion script for zsh |
| `fish` | Generate autocompletion script for fish |
| `powershell` | Generate autocompletion script for PowerShell |

#### Installation Examples

**Bash (Linux):**
```bash
dockerutils completion bash | sudo tee /etc/bash_completion.d/dockerutils
```

**Bash (macOS with Homebrew):**
```bash
dockerutils completion bash | sudo tee /usr/local/etc/bash_completion.d/dockerutils
```

**Zsh:**
```bash
# Add to ~/.zshrc first:
autoload -U compinit; compinit

# Generate completion file:
dockerutils completion zsh > "${fpath[1]}/_dockerutils"
```

**Fish:**
```bash
dockerutils completion fish > ~/.config/fish/completions/dockerutils.fish
```

**PowerShell:**
```powershell
# Add to your PowerShell profile:
dockerutils completion powershell | Out-String | Invoke-Expression
```

#### Manual Installation

You can also save the completion script to a file and source it manually:

```bash
# Generate and save
dockerutils completion bash > dockerutils-completion.bash

# Source in your shell profile (~/.bashrc, ~/.zshrc, etc.)
source /path/to/dockerutils-completion.bash
```

## Shell Completion

### Features

Shell completion provides:
- **Command Completion**: Tab completion for all available commands
- **Flag Completion**: Tab completion for command flags and options
- **Value Completion**: Context-aware completion for flag values where applicable
- **Help Integration**: Completion suggestions include brief descriptions

### Supported Shells

| Shell | Status | Notes |
|-------|--------|-------|
| Bash | ✅ Full support | Requires bash-completion package |
| Zsh | ✅ Full support | Built-in support |
| Fish | ✅ Full support | Built-in support |
| PowerShell | ✅ Full support | Windows and cross-platform |

### Testing Completion

After installation, test completion by typing:

```bash
dockerutils <TAB>
# Should show: completion, help, server, version

dockerutils server --<TAB>
# Should show: --help, --host, --no-browser, --port, --tmp-dir, --web-root
```

## Examples

### Basic Usage

```bash
# Show help
dockerutils --help

# Check version
dockerutils -v

# Start web interface
dockerutils server
```

### Web Server Configuration

```bash
# Development setup
dockerutils server --port 3000

# Production setup
dockerutils server \
  --host 0.0.0.0 \
  --port 8080 \
  --tmp-dir /var/lib/dockerutils \
  --no-browser

# Custom UI development
dockerutils server \
  --web-root ./my-custom-ui \
  --tmp-dir ./dev-data \
  --port 3000
```

### Automation and Scripting

```bash
#!/bin/bash
# start-analysis-server.sh

# Check if dockerutils is available
if ! command -v dockerutils &> /dev/null; then
    echo "dockerutils not found in PATH"
    exit 1
fi

# Start server in background
dockerutils server \
  --host 0.0.0.0 \
  --port 8080 \
  --no-browser \
  --tmp-dir /app/analysis-data &

SERVER_PID=$!
echo "Docker Utils server started with PID: $SERVER_PID"

# Save PID for cleanup
echo $SERVER_PID > /tmp/dockerutils.pid

# Wait for server to start
sleep 3

# Test server health
if curl -f -s http://localhost:8080/api/health > /dev/null; then
    echo "✅ Server is healthy and ready"
else
    echo "❌ Server health check failed"
    kill $SERVER_PID
    exit 1
fi
```

### CI/CD Integration

```bash
#!/bin/bash
# ci-analysis.sh

# Start dockerutils for CI analysis
dockerutils server --no-browser --port 8080 &
DOCKERUTILS_PID=$!

# Cleanup function
cleanup() {
    kill $DOCKERUTILS_PID 2>/dev/null
    exit $1
}

# Setup cleanup on script exit
trap 'cleanup $?' EXIT

# Wait for server to start
sleep 5

# Analyze test image
IMAGE_NAME="$1"
if [ -z "$IMAGE_NAME" ]; then
    echo "Usage: $0 <image-name>"
    cleanup 1
fi

echo "Analyzing image: $IMAGE_NAME"

# Perform analysis via API
RESULT=$(curl -s -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d "{\"image_name\": \"$IMAGE_NAME\"}")

# Check if analysis succeeded
if echo "$RESULT" | jq -r '.success' | grep -q true; then
    echo "✅ Analysis completed successfully"
    
    # Extract key metrics
    IMAGE_ID=$(echo "$RESULT" | jq -r '.image_id')
    echo "Image ID: $IMAGE_ID"
    
    # You can add additional checks here
    # e.g., size limits, security scans, etc.
    
    cleanup 0
else
    echo "❌ Analysis failed"
    echo "$RESULT" | jq -r '.message'
    cleanup 1
fi
```

## Environment Variables

Docker Utils supports the following environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `DOCKER_HOST` | Docker daemon socket to connect to | Platform default |
| `DOCKER_API_VERSION` | Docker API version to use | Latest supported |
| `DOCKER_CERT_PATH` | Path to TLS certificates | (empty) |
| `DOCKER_TLS_VERIFY` | Enable TLS verification | `false` |
| `DOCKER_UTILS_DEBUG` | Enable debug logging | `false` |

### Example with Environment Variables

```bash
# Connect to remote Docker daemon
export DOCKER_HOST=tcp://remote-docker:2376
export DOCKER_TLS_VERIFY=1
export DOCKER_CERT_PATH=/path/to/certs

dockerutils server

# Enable debug logging
export DOCKER_UTILS_DEBUG=1
dockerutils server
```

### Docker Socket Access

Docker Utils requires access to the Docker daemon socket. Common configurations:

#### Linux/macOS (Unix Socket)
```bash
# Default socket location
export DOCKER_HOST=unix:///var/run/docker.sock

# Ensure user has permission
sudo usermod -aG docker $USER
# Then logout and login again
```

#### Windows (Named Pipe)
```cmd
# Default Windows configuration
set DOCKER_HOST=npipe:////./pipe/docker_engine
```

#### Remote Docker Daemon
```bash
# TCP connection with TLS
export DOCKER_HOST=tcp://docker-host:2376
export DOCKER_TLS_VERIFY=1
export DOCKER_CERT_PATH=/path/to/certs

# TCP connection without TLS (not recommended)
export DOCKER_HOST=tcp://docker-host:2375
```

---

For more information about specific features and advanced usage, see:
- [User Guide](USER_GUIDE.md) - Comprehensive tutorials and use cases
- [API Documentation](API.md) - REST API reference and integration examples
- [README](../README.md) - Project overview and quick start guide

# Development Guide

This guide provides comprehensive instructions for setting up a development environment and contributing to Docker Utils.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Development Setup](#development-setup)
3. [Project Structure](#project-structure)
4. [Development Workflow](#development-workflow)
5. [Testing](#testing)
6. [Building](#building)
7. [Debugging](#debugging)
8. [Code Standards](#code-standards)
9. [Development Tools](#development-tools)

## Prerequisites

### Required Software

- **Go 1.24.2 or later**
  ```bash
  # Check Go version
  go version
  # Should output: go version go1.24.2 or later
  ```

- **Docker Engine**
  ```bash
  # Check Docker installation
  docker --version
  docker ps
  ```

- **Git**
  ```bash
  git --version
  ```

- **Make** (recommended for build automation)
  ```bash
  make --version
  ```

### Optional Development Tools

- **golangci-lint** (code linting)
  ```bash
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

- **gofumpt** (enhanced code formatting)
  ```bash
  go install mvdan.cc/gofumpt@latest
  ```

- **govulncheck** (vulnerability scanning)
  ```bash
  go install golang.org/x/vuln/cmd/govulncheck@latest
  ```

### Platform-Specific Requirements

#### Linux
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install docker.io git make build-essential

# Add user to docker group
sudo usermod -aG docker $USER
newgrp docker

# RHEL/CentOS/Fedora
sudo dnf install docker git make gcc
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker $USER
```

#### macOS
```bash
# Install Docker Desktop
brew install --cask docker

# Install development tools
brew install git make
```

#### Windows
1. Install [Docker Desktop for Windows](https://docker.com)
2. Install [Git for Windows](https://git-scm.com/download/win)
3. Install [Go for Windows](https://golang.org/dl/)
4. Use PowerShell or WSL2 for development

## Development Setup

### 1. Fork and Clone Repository

```bash
# Fork the repository on GitHub, then clone your fork
git clone https://github.com/YOUR-USERNAME/dockerutils.git
cd dockerutils

# Add upstream remote
git remote add upstream https://github.com/smiller333/dockerutils.git

# Verify remotes
git remote -v
```

### 2. Environment Configuration

```bash
# Set up Go workspace (if using GOPATH mode)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Verify Go environment
go env GOPATH
go env GOROOT
```

### 3. Install Dependencies

```bash
# Download Go modules
go mod download

# Verify dependencies
go mod verify

# Install development tools
make install-tools  # or manual installation as shown above
```

### 4. Verify Setup

```bash
# Build the project
make build-dev

# Run tests
make test

# Start development server
./bin/dockerutils server --port 8080

# In another terminal, verify it's working
curl http://localhost:8080/api/health
```

## Project Structure

```
dockerutils/
├── cmd/                     # CLI command implementations
│   ├── root.go             # Root command and global flags
│   ├── tools.go            # Web server command
│   └── completion.go       # Shell completion
├── src/                     # Core implementation packages
│   ├── analyzer/           # Docker image analysis
│   │   ├── analyzer.go     # Analysis types and utilities
│   │   ├── fileprocessor.go# File processing logic
│   │   ├── imageprocessor.go# Image processing logic
│   │   └── *_test.go       # Unit tests
│   ├── buildcontext/       # Build context analysis
│   │   ├── buildcontext.go # Main implementation
│   │   └── *_test.go       # Unit tests
│   ├── dockerclient/       # Docker API wrapper
│   │   ├── client.go       # Enhanced Docker client
│   │   ├── README.md       # Package documentation
│   │   └── *_test.go       # Unit tests
│   ├── version/            # Version information
│   │   ├── version.go      # Version types and functions
│   │   └── *_test.go       # Unit tests
│   └── webserver/          # HTTP server and API
│       ├── server.go       # Server implementation
│       ├── webpages/       # Embedded static files
│       ├── README.md       # Package documentation
│       └── *_test.go       # Unit tests
├── docs/                   # Documentation
├── examples/               # Example files and usage
├── scripts/                # Build and utility scripts
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Module checksums
├── Makefile               # Build automation
└── Dockerfile             # Container image definition
```

### Package Organization Principles

1. **Separation of Concerns**: Each package has a single, well-defined responsibility
2. **Dependency Direction**: Packages import from lower-level to higher-level (no circular imports)
3. **Interface-Based Design**: Use interfaces for testability and modularity
4. **Minimal External Dependencies**: Prefer standard library when possible

## Development Workflow

### 1. Branch Management

```bash
# Create feature branch
git checkout -b feature/your-feature-name

# Keep branch updated
git fetch upstream
git rebase upstream/main

# Push changes
git push origin feature/your-feature-name
```

### 2. Development Cycle

```bash
# 1. Make changes to code
vim src/analyzer/analyzer.go

# 2. Run tests frequently
make test-package PKG=analyzer

# 3. Check code formatting and linting
make fmt
make lint

# 4. Build and test manually
make build-dev
./bin/dockerutils server

# 5. Run full test suite
make test
make test-integration

# 6. Commit changes
git add .
git commit -m "feat(analyzer): add new analysis feature"
```

### 3. Code Review Preparation

```bash
# Run all checks before creating PR
make check           # Runs fmt, lint, test
make build-release   # Verify release build works
make clean          # Clean up artifacts
```

## Testing

### Unit Tests

```bash
# Run all tests
make test

# Run specific package tests
go test -v ./src/analyzer/

# Run tests with coverage
make test-coverage

# Run tests with race detection
go test -race ./...

# Run specific test function
go test -v -run TestAnalyzeImage ./src/analyzer/
```

### Integration Tests

```bash
# Run integration tests (requires Docker)
go test -v -tags=integration ./...

# Run integration tests with cleanup
make test-integration
```

### Benchmark Tests

```bash
# Run benchmarks
go test -bench=. ./src/analyzer/

# Profile benchmarks
go test -bench=. -cpuprofile=cpu.prof ./src/analyzer/
go tool pprof cpu.prof
```

### Test Guidelines

1. **Test Coverage**: Aim for >80% coverage on new code
2. **Table-Driven Tests**: Use for multiple test cases
3. **Mocking**: Mock external dependencies using interfaces
4. **Test Data**: Use test fixtures in `testdata/` directories
5. **Integration Tests**: Tag with `// +build integration`

### Example Test Structure

```go
func TestAnalyzeImage(t *testing.T) {
    tests := []struct {
        name        string
        imageName   string
        expectError bool
        want        *AnalysisResult
    }{
        {
            name:        "valid_nginx_image",
            imageName:   "nginx:alpine",
            expectError: false,
            want: &AnalysisResult{
                ImageTag: "nginx:alpine",
                // ... expected fields
            },
        },
        {
            name:        "invalid_image_name",
            imageName:   "nonexistent:latest",
            expectError: true,
            want:        nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
            result, err := AnalyzeImage(tt.imageName)
            
            if tt.expectError {
                if err == nil {
                    t.Error("expected error but got none")
                }
                return
            }
            
            if err != nil {
                t.Errorf("unexpected error: %v", err)
            }
            
            // Validate result...
        })
    }
}
```

## Building

### Development Build

```bash
# Quick development build
make build-dev

# Build with specific version
make VERSION=v1.0.0 build-dev

# Build for current platform
go build -o bin/dockerutils .
```

### Release Build

```bash
# Build release version with optimizations
make build-release

# Build for all platforms
make build-all

# Build with custom ldflags
go build -ldflags "-X github.com/smiller333/dockerutils/src/version.Version=v1.0.0" -o bin/dockerutils .
```

### Cross-Platform Building

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 make build-release

# Build for Windows
GOOS=windows GOARCH=amd64 make build-release

# Build for macOS (Intel)
GOOS=darwin GOARCH=amd64 make build-release

# Build for macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 make build-release
```

### Container Building

```bash
# Build Docker image
make docker-build

# Build with specific tag
make docker-build TAG=v1.0.0

# Multi-platform build
docker buildx build --platform linux/amd64,linux/arm64 -t dockerutils:latest .
```

## Debugging

### Local Debugging

```bash
# Enable debug logging
export DEBUG=1
./bin/dockerutils server

# Use delve debugger
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug main application
dlv debug . -- tools --port 8080

# Debug specific test
dlv test ./src/analyzer/ -- -test.run TestAnalyzeImage
```

### Web Server Debugging

```bash
# Start with verbose logging
./bin/dockerutils server --port 8080 --verbose

# Test API endpoints
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{"image_name": "alpine:latest"}'

# Monitor server logs
tail -f /tmp/dockerutils.log
```

### Docker Integration Debugging

```bash
# Check Docker connectivity
docker ps
docker version

# Test Docker client
go run -tags debug ./cmd/test-docker-client.go

# Debug Docker socket permissions
ls -la /var/run/docker.sock
groups $USER
```

### Performance Profiling

```bash
# Start server with profiling
go run . tools --port 8080 --profile

# CPU profiling
go tool pprof http://localhost:8080/debug/pprof/profile

# Memory profiling
go tool pprof http://localhost:8080/debug/pprof/heap

# Goroutine profiling
go tool pprof http://localhost:8080/debug/pprof/goroutine
```

## Code Standards

### Go Code Standards

Follow the project's [Go coding standards](../.cursor/rules/go-coding-standards.mdc):

1. **Formatting**: Use `gofmt` and `gofumpt`
2. **Linting**: Pass `golangci-lint` checks
3. **Documentation**: Include godoc comments for all exported items
4. **Error Handling**: Always handle errors explicitly
5. **Testing**: Write comprehensive tests for all new code

### Commit Message Format

Follow [Conventional Commits](https://conventionalcommits.org/) format:

```
type(scope): brief description

Longer description if needed

Fixes #123
```

**Types**: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

### Code Review Checklist

- [ ] Code follows Go best practices
- [ ] All tests pass locally
- [ ] Code is properly documented
- [ ] No linting errors
- [ ] Appropriate error handling
- [ ] Security considerations addressed
- [ ] Performance implications considered

## Development Tools

### Makefile Targets

```bash
# Development
make build-dev      # Build development version
make test           # Run all tests
make test-coverage  # Run tests with coverage
make lint          # Run linter
make fmt           # Format code

# Release
make build-release  # Build optimized release
make build-all     # Build for all platforms
make docker-build  # Build Docker image

# Utilities
make clean         # Clean build artifacts
make install-tools # Install development tools
make check         # Run all checks (fmt, lint, test)
```

### IDE Configuration

#### VS Code

Create `.vscode/settings.json`:
```json
{
    "go.useLanguageServer": true,
    "go.formatTool": "gofumpt",
    "go.lintTool": "golangci-lint",
    "go.lintOnSave": "package",
    "go.testFlags": ["-v"],
    "go.buildTags": "integration"
}
```

#### GoLand/IntelliJ

1. Enable Go modules support
2. Configure Go linter to use golangci-lint
3. Set code style to use gofumpt
4. Enable automatic imports organization

### Git Hooks

Set up pre-commit hooks:

```bash
# Install pre-commit hook
cat > .git/hooks/pre-commit << 'EOF'
#!/bin/bash
make fmt
make lint
make test
EOF

chmod +x .git/hooks/pre-commit
```

---

**Next Steps:**
- Review the [Contributing Guide](../CONTRIBUTING.md) for contribution workflow
- Check [Testing Guidelines](internal/TESTING_GUIDE.md) for detailed testing practices
- See [Architecture Overview](ARCHITECTURE.md) for system design understanding

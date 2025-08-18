# Contributing to Docker Utils

Thank you for your interest in contributing to Docker Utils! This document provides comprehensive guidelines for contributors, from first-time contributors to experienced developers.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Development Setup](#development-setup)
3. [Code Standards](#code-standards)
4. [Testing Guidelines](#testing-guidelines)
5. [Documentation](#documentation)
6. [Submitting Changes](#submitting-changes)
7. [Project Structure](#project-structure)
8. [Security Guidelines](#security-guidelines)
9. [Release Process](#release-process)
10. [Community](#community)

## Getting Started

### Prerequisites

- **Go 1.24.2 or later**: [Installation Guide](https://golang.org/doc/install)
- **Git**: For version control
- **Docker**: For testing Docker integrations
- **Make**: For build automation (optional but recommended)

### Quick Start

```bash
# 1. Fork and clone the repository
git clone https://github.com/YOUR_USERNAME/dockerutils.git
cd dockerutils

# 2. Install dependencies
go mod download

# 3. Run tests to ensure everything works
make test

# 4. Build the project
make build-dev

# 5. Test the binary
./bin/dockerutils --help
```

### Your First Contribution

Looking for a way to contribute? Here are some great starting points:

- 🐛 **Bug Reports**: Help us identify and fix issues
- 📚 **Documentation**: Improve existing docs or add missing information
- 🧪 **Tests**: Add test coverage for untested functionality
- ✨ **Small Features**: Implement small, well-defined features
- 🎨 **UI Improvements**: Enhance the web interface

Check our [Issues](https://github.com/smiller333/dockerutils/issues) for tasks labeled `good first issue` or `help wanted`.

## Development Setup

### Project Layout

```
dockerutils/
├── cmd/                 # CLI command definitions (Cobra-based)
├── src/                 # Implementation logic (organized by package)
│   ├── analyzer/       # Image and file analysis
│   ├── buildcontext/   # Build context management
│   ├── dockerclient/   # Docker SDK wrapper
│   ├── version/        # Version management
│   └── webserver/      # Web server and API
├── docs/               # Documentation
├── examples/           # Example files and usage
├── scripts/            # Build and utility scripts
├── tmp/                # Temporary analysis outputs
└── bin/                # Built binaries
```

### Development Environment

#### IDE Setup

**VS Code (Recommended)**
```json
// .vscode/settings.json
{
    "go.testFlags": ["-v"],
    "go.testOnSave": true,
    "go.lintOnSave": "package",
    "go.formatOnSave": true,
    "go.useLanguageServer": true
}
```

**GoLand/IntelliJ IDEA**
- Enable Go modules support
- Set GOROOT to your Go installation
- Configure gofmt on save

#### Environment Variables

```bash
# Development environment
export DOCKER_UTILS_DEBUG=1  # Enable debug logging
export DOCKER_HOST=unix:///var/run/docker.sock  # Docker socket
```

#### Git Configuration

```bash
# Set up commit message template
git config commit.template .gitmessage

# Set up useful aliases
git config alias.co checkout
git config alias.br branch
git config alias.ci commit
git config alias.st status
```

### Build System

#### Make Targets

```bash
# Development
make build-dev          # Build with debug info
make test              # Run all tests
make test-short        # Run unit tests only
make lint              # Run linters
make fmt               # Format code

# Production
make build-release     # Optimized build
make build-all         # Multi-platform builds
make clean             # Clean build artifacts

# Development workflow
make dev               # build-dev + test + lint
```

#### Manual Build Commands

```bash
# Basic development build
go build -o bin/dockerutils

# Build with version information
scripts/build.sh dev

# Cross-compilation for different platforms
GOOS=linux GOARCH=amd64 go build -o bin/dockerutils-linux-amd64
GOOS=windows GOARCH=amd64 go build -o bin/dockerutils-windows-amd64.exe
```

## Code Standards

### Go Coding Standards

Docker Utils follows strict Go best practices:

#### Package Organization
- **All implementation logic** goes in `src/` directory
- **CLI commands** are defined in `cmd/` directory using Cobra
- **Each package** should have a single, well-defined responsibility
- **Delegate business logic** from commands to packages in `src/`

#### Code Style

```go
// ✅ Good: Clear package documentation
// Package analyzer provides functionality for analyzing Dockerfiles and Docker images.
// It includes secure tar extraction, image processing, and comprehensive reporting capabilities.
package analyzer

// ✅ Good: Exported function with comprehensive documentation
// SafeTarExtraction extracts tar archive using Go native library with security controls.
// It validates file paths to prevent directory traversal attacks and handles gzipped archives.
//
// Parameters:
//   - tarPath: Path to the tar archive file
//   - destDir: Destination directory for extraction
//
// Returns:
//   - error: nil on success, error describing the failure otherwise
//
// Security: This function includes path traversal protection and validates all file paths.
func SafeTarExtraction(tarPath string, destDir string) error {
    // Implementation...
}

// ✅ Good: Proper error handling
func processImage(imageName string) (*AnalysisResult, error) {
    if imageName == "" {
        return nil, fmt.Errorf("image name cannot be empty")
    }
    
    result, err := analyzeImage(imageName)
    if err != nil {
        return nil, fmt.Errorf("failed to analyze image %s: %w", imageName, err)
    }
    
    return result, nil
}
```

#### Naming Conventions

```go
// ✅ Good naming examples
type DockerClient struct {}        // PascalCase for exported types
var ErrImageNotFound = errors.New("image not found")  // Exported errors with Err prefix

func NewDockerClient() *DockerClient {}  // Constructor with New prefix
func (c *DockerClient) InspectImage() {} // Methods in PascalCase

var imageID string                 // camelCase for variables
const defaultTimeout = 30         // camelCase for constants
```

#### Error Handling

```go
// ✅ Good: Explicit error handling
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file %s: %w", filename, err)
    }
    defer file.Close()
    
    // Process file...
    
    return nil
}

// ✅ Good: Custom error types
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}
```

### Documentation Standards

#### Package Documentation

```go
// Package webserver provides a web server with both a modern web UI and REST API 
// for viewing and managing Docker image analysis results.
//
// The server supports both viewing existing analysis results and triggering new 
// image analysis operations through a comprehensive REST API and interactive web interface.
//
// Key Features:
//   - Embedded web interface with modern UI
//   - Complete REST API for programmatic access
//   - Live Docker image analysis capabilities
//   - Asynchronous operations with progress tracking
//   - Automatic file discovery and management
//
// Example usage:
//   config := &webserver.Config{
//       Host: "localhost",
//       Port: "8080",
//   }
//   
//   server, err := webserver.New(config)
//   if err != nil {
//       log.Fatal(err)
//   }
//   
//   if err := server.Start(); err != nil {
//       log.Fatal(err)
//   }
package webserver
```

#### Function Documentation

```go
// InspectImage retrieves detailed information about a Docker image.
//
// This method uses the Docker API to fetch comprehensive metadata about the specified image,
// including configuration, layer information, and creation details.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - nameOrID: Image name (e.g., "nginx:latest") or image ID
//
// Returns:
//   - *image.InspectResponse: Detailed image information from Docker API
//   - error: ErrImageNotFound if image doesn't exist, or other API errors
//
// Example:
//   info, err := client.InspectImage(ctx, "nginx:latest")
//   if err != nil {
//       if errors.Is(err, dockerclient.ErrImageNotFound) {
//           log.Println("Image not found")
//           return
//       }
//       log.Fatal(err)
//   }
//   fmt.Printf("Image ID: %s\n", info.ID)
func (c *DockerClient) InspectImage(ctx context.Context, nameOrID string) (*image.InspectResponse, error) {
    // Implementation...
}
```

### Code Quality Tools

#### Linting

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linting
make lint

# Or manually
golangci-lint run
```

#### Formatting

```bash
# Format all Go files
make fmt

# Or manually
gofmt -w .
go mod tidy
```

#### Static Analysis

```bash
# Run Go vet
go vet ./...

# Run staticcheck
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...
```

## Testing Guidelines

### Test Structure

Docker Utils uses Go's built-in testing package with table-driven tests:

```go
func TestSafeTarExtraction(t *testing.T) {
    tests := []struct {
        name        string
        tarPath     string
        destDir     string
        setupFunc   func(t *testing.T) string  // Returns temp dir
        wantErr     bool
        errContains string
    }{
        {
            name:    "valid tar file",
            tarPath: "testdata/valid.tar",
            destDir: "",  // Will be set by setupFunc
            setupFunc: func(t *testing.T) string {
                return t.TempDir()
            },
            wantErr: false,
        },
        {
            name:        "nonexistent tar file",
            tarPath:     "testdata/nonexistent.tar",
            destDir:     "",
            setupFunc:   func(t *testing.T) string { return t.TempDir() },
            wantErr:     true,
            errContains: "failed to open tar file",
        },
        // Add more test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.setupFunc != nil {
                tt.destDir = tt.setupFunc(t)
            }
            
            err := SafeTarExtraction(tt.tarPath, tt.destDir)
            
            if tt.wantErr {
                if err == nil {
                    t.Errorf("SafeTarExtraction() expected error, got nil")
                    return
                }
                if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
                    t.Errorf("SafeTarExtraction() error = %v, want error containing %v", err, tt.errContains)
                }
                return
            }
            
            if err != nil {
                t.Errorf("SafeTarExtraction() unexpected error = %v", err)
            }
        })
    }
}
```

### Test Categories

#### Unit Tests
```bash
# Run unit tests only (no Docker daemon required)
go test -short ./...

# Run specific package tests
go test ./src/analyzer/

# Run with verbose output
go test -v ./src/analyzer/
```

#### Integration Tests
```bash
# Run all tests including integration tests (requires Docker daemon)
go test ./...

# Run integration tests only
go test -run Integration ./...
```

#### Test Coverage
```bash
# Generate coverage report
go test -cover ./...

# Detailed coverage with HTML report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Data Management

#### Test Files Structure
```
testdata/
├── dockerfiles/
│   ├── valid.Dockerfile
│   ├── invalid.Dockerfile
│   └── complex.Dockerfile
├── tar-archives/
│   ├── simple.tar
│   ├── gzipped.tar.gz
│   └── malicious.tar
└── expected-outputs/
    ├── analysis-result.json
    └── summary.json
```

#### Creating Test Data
```go
func createTestTarFile(t *testing.T, files map[string]string) string {
    t.Helper()
    
    tmpFile, err := os.CreateTemp("", "test-*.tar")
    if err != nil {
        t.Fatal(err)
    }
    defer tmpFile.Close()
    
    tw := tar.NewWriter(tmpFile)
    defer tw.Close()
    
    for name, content := range files {
        hdr := &tar.Header{
            Name: name,
            Mode: 0644,
            Size: int64(len(content)),
        }
        
        if err := tw.WriteHeader(hdr); err != nil {
            t.Fatal(err)
        }
        
        if _, err := tw.Write([]byte(content)); err != nil {
            t.Fatal(err)
        }
    }
    
    return tmpFile.Name()
}
```

### Mocking

#### Docker Client Mocking
```go
// Interface for dependency injection
type DockerClientInterface interface {
    InspectImage(ctx context.Context, imageID string) (*image.InspectResponse, error)
    SaveImage(ctx context.Context, imageNames []string) (io.ReadCloser, error)
    // ... other methods
}

// Mock implementation for testing
type MockDockerClient struct {
    InspectImageFunc func(ctx context.Context, imageID string) (*image.InspectResponse, error)
    SaveImageFunc    func(ctx context.Context, imageNames []string) (io.ReadCloser, error)
}

func (m *MockDockerClient) InspectImage(ctx context.Context, imageID string) (*image.InspectResponse, error) {
    if m.InspectImageFunc != nil {
        return m.InspectImageFunc(ctx, imageID)
    }
    return nil, errors.New("not implemented")
}
```

## Documentation

### Documentation Requirements

All changes must include appropriate documentation updates:

#### Code Documentation
- **Package-level**: Comprehensive overview of package purpose and usage
- **Exported functions**: Detailed godoc comments with examples
- **Complex logic**: Inline comments explaining algorithms and decisions
- **Security considerations**: Document security implications of functions

#### User Documentation
- **API changes**: Update API documentation
- **New features**: Add to user guide with examples
- **CLI changes**: Update CLI reference
- **Breaking changes**: Clearly document in upgrade guide

### Documentation Tools

#### Generating API Documentation
```bash
# Generate godoc locally
godoc -http=:6060
# Then visit http://localhost:6060/pkg/github.com/smiller333/dockerutils/

# Generate markdown documentation
go install github.com/davecheney/godoc2md@latest
godoc2md github.com/smiller333/dockerutils/src/analyzer > docs/analyzer-api.md
```

#### Documentation Testing
```bash
# Test example code in documentation
go test -run Example ./...

# Validate all links in markdown files
markdown-link-check docs/*.md
```

## Submitting Changes

### Git Workflow

#### Branch Naming
```bash
# Feature branches
feature/add-image-scanning
feature/improve-web-ui

# Bug fix branches
bugfix/fix-tar-extraction
bugfix/resolve-memory-leak

# Documentation branches
docs/update-api-reference
docs/add-contributing-guide
```

#### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```bash
# Format: type(scope): description
feat(analyzer): add secure tar extraction with path validation
fix(webserver): resolve memory leak in image analysis
docs(api): update REST API documentation with new endpoints
test(dockerclient): add integration tests for image operations
refactor(version): simplify version string generation
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `test`: Adding or updating tests
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `chore`: Maintenance tasks

### Pull Request Process

#### Before Submitting

1. **Run Tests**: Ensure all tests pass
   ```bash
   make test
   ```

2. **Run Linting**: Fix any linting issues
   ```bash
   make lint
   ```

3. **Update Documentation**: Include relevant documentation updates

4. **Test Build**: Verify the build works
   ```bash
   make build-release
   ```

#### Pull Request Template

```markdown
## Description
Brief description of changes and motivation.

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Performance improvement
- [ ] Code refactoring

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass (if applicable)
- [ ] Manual testing completed
- [ ] New tests added for new functionality

## Documentation
- [ ] Code documentation updated
- [ ] User documentation updated (if applicable)
- [ ] API documentation updated (if applicable)

## Checklist
- [ ] My code follows the project's style guidelines
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes

## Related Issues
Fixes #(issue_number)
Related to #(issue_number)
```

#### Review Process

1. **Automated Checks**: CI/CD pipeline runs tests and linting
2. **Code Review**: At least one maintainer reviews the code
3. **Documentation Review**: Documentation changes are reviewed
4. **Testing**: Changes are tested in different environments
5. **Approval**: Maintainer approves and merges the PR

## Security Guidelines

### Security Considerations

#### Input Validation
```go
// ✅ Good: Validate and sanitize file paths
func validatePath(path string) error {
    // Check for path traversal attempts
    if strings.Contains(path, "..") {
        return fmt.Errorf("path traversal detected: %s", path)
    }
    
    // Ensure path is within expected directory
    absPath, err := filepath.Abs(path)
    if err != nil {
        return fmt.Errorf("failed to resolve absolute path: %w", err)
    }
    
    if !strings.HasPrefix(absPath, allowedBasePath) {
        return fmt.Errorf("path outside allowed directory: %s", path)
    }
    
    return nil
}
```

#### Secure Defaults
```go
// ✅ Good: Secure configuration defaults
type Config struct {
    Host           string        `json:"host"`           // Default: "localhost" (not 0.0.0.0)
    Port           string        `json:"port"`           // Default: "8080"
    Timeout        time.Duration `json:"timeout"`        // Default: 30s
    MaxFileSize    int64         `json:"max_file_size"`  // Default: 100MB
    AllowedOrigins []string      `json:"allowed_origins"` // Default: ["localhost"]
}

func DefaultConfig() *Config {
    return &Config{
        Host:           "localhost",  // Bind to localhost only
        Port:           "8080",
        Timeout:        30 * time.Second,
        MaxFileSize:    100 * 1024 * 1024,  // 100MB
        AllowedOrigins: []string{"localhost", "127.0.0.1"},
    }
}
```

#### Sensitive Data Handling
```go
// ✅ Good: Avoid logging sensitive information
func logImageAnalysis(imageName string, envVars map[string]string) {
    log.Printf("Analyzing image: %s", imageName)
    
    // Don't log environment variables that might contain secrets
    sensitiveKeys := []string{"password", "token", "key", "secret"}
    filteredEnv := make(map[string]string)
    
    for k, v := range envVars {
        if containsSensitiveKey(k, sensitiveKeys) {
            filteredEnv[k] = "[REDACTED]"
        } else {
            filteredEnv[k] = v
        }
    }
    
    log.Printf("Environment variables: %v", filteredEnv)
}
```

### Security Reporting

If you discover a security vulnerability, please:

1. **DO NOT** open a public issue
2. **Email**: security@example.com (or create a private security advisory)
3. **Include**: Detailed description, steps to reproduce, and potential impact
4. **Wait**: For confirmation and resolution before public disclosure

## Project Structure

### Architecture Overview

```
Docker Utils Architecture

┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   CLI Commands  │───▶│  Business Logic  │───▶│ External APIs   │
│   (cmd/)        │    │  (src/)          │    │ (Docker SDK)    │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       
         │                       ▼                       
         │              ┌──────────────────┐             
         │              │   Web Server     │             
         │              │   & REST API     │             
         │              │   (src/webserver)│             
         └──────────────▶└──────────────────┘             
                                 │                        
                                 ▼                        
                        ┌──────────────────┐             
                        │   Web Interface  │             
                        │   (embedded)     │             
                        └──────────────────┘             
```

### Package Responsibilities

| Package | Responsibility | Dependencies |
|---------|---------------|--------------|
| `cmd/` | CLI interface, command definitions | `src/*`, `cobra` |
| `src/analyzer/` | Image analysis, tar extraction | `dockerclient`, Docker SDK |
| `src/buildcontext/` | Build context validation | Standard library |
| `src/dockerclient/` | Docker API wrapper | Docker SDK |
| `src/version/` | Version management | Standard library |
| `src/webserver/` | Web server, REST API, UI | `analyzer`, `dockerclient` |

### Dependency Management

#### Adding Dependencies
```bash
# Add a new dependency
go get github.com/new/dependency@v1.2.3

# Update go.mod and go.sum
go mod tidy

# Verify dependencies
go mod verify
```

#### Dependency Guidelines
- **Minimize dependencies**: Only add dependencies that provide significant value
- **Use stable versions**: Prefer tagged releases over commit SHAs
- **Security**: Regularly update dependencies for security patches
- **License compatibility**: Ensure licenses are compatible with MIT license

## Release Process

### Version Strategy

Docker Utils follows [Semantic Versioning](https://semver.org/):

- **MAJOR** (1.0.0): Breaking changes, major architectural changes
- **MINOR** (1.1.0): New features, backwards compatible changes  
- **PATCH** (1.1.1): Bug fixes, security updates

### Release Workflow

#### 1. Prepare Release
```bash
# Create release branch
git checkout -b release/v1.2.0

# Update version information
# Update CHANGELOG.md
# Update documentation

# Final testing
make test
make build-all
```

#### 2. Create Release
```bash
# Tag the release
git tag -a v1.2.0 -m "Release version 1.2.0"

# Push tag
git push origin v1.2.0

# GitHub Actions will automatically:
# - Build binaries for all platforms
# - Create GitHub release
# - Upload release assets
```

#### 3. Post-Release
```bash
# Merge release branch back to main
git checkout main
git merge release/v1.2.0

# Delete release branch
git branch -d release/v1.2.0
git push origin --delete release/v1.2.0

# Update documentation site
# Announce release in community channels
```

### Release Checklist

- [ ] All tests pass
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
- [ ] Security review completed
- [ ] Performance benchmarks run
- [ ] Cross-platform builds tested
- [ ] Migration guide written (for breaking changes)

## Community

### Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). Please read and follow it in all interactions.

### Communication Channels

- **GitHub Issues**: Bug reports, feature requests
- **GitHub Discussions**: General questions, ideas, announcements
- **Pull Requests**: Code contributions, documentation improvements

### Recognition

Contributors are recognized in several ways:

- **CONTRIBUTORS.md**: Listed in the contributors file
- **Release Notes**: Mentioned in release announcements
- **GitHub**: Contributor status and commit history

### Getting Help

- **Documentation**: Check the docs/ directory first
- **Issues**: Search existing issues before creating new ones
- **Discussions**: Use GitHub Discussions for questions
- **Code Review**: Maintainers provide helpful feedback on PRs

---

## Thank You!

Thank you for contributing to Docker Utils! Your contributions help make Docker image analysis more accessible and secure for developers worldwide.

For questions about contributing, please open a GitHub Discussion or reach out to the maintainers.

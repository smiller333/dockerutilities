# Docker Utils Architecture

This document provides a comprehensive overview of the Docker Utils system architecture, component structure, and technical design decisions.

## Table of Contents

1. [System Overview](#system-overview)
2. [Architecture Principles](#architecture-principles)
3. [Component Architecture](#component-architecture)
4. [Data Flow](#data-flow)
5. [Technology Stack](#technology-stack)
6. [Security Architecture](#security-architecture)
7. [Performance Considerations](#performance-considerations)
8. [Deployment Architecture](#deployment-architecture)

## System Overview

Docker Utils is a CLI tool and web interface for comprehensive Docker image analysis. The system is built on a modular, service-oriented architecture that separates concerns between command-line operations, web services, and analysis engines.

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Docker Utils                             │
├─────────────────────────────────────────────────────────────┤
│  CLI Interface (Cobra)     │     Web Interface (HTTP)       │
│  ┌─────────────────────────┼─────────────────────────────┐   │
│  │ Commands                │ REST API + Static Files    │   │
│  │ - server                │ - /api/analyze             │   │
│  │ - version               │ - /api/summaries           │   │
│  │ - completion            │ - /api/buildcontext        │   │
│  │                         │ - Static HTML/CSS/JS       │   │
│  └─────────────────────────┼─────────────────────────────┘   │
├─────────────────────────────────────────────────────────────┤
│                  Core Services Layer                        │
│  ┌──────────────┬──────────────┬──────────────┬─────────────┐│
│  │   Analyzer   │  Web Server  │Build Context │  Version    ││
│  │   Service    │   Service    │   Service    │  Service    ││
│  └──────────────┴──────────────┴──────────────┴─────────────┘│
├─────────────────────────────────────────────────────────────┤
│                 Infrastructure Layer                        │
│  ┌──────────────┬──────────────┬──────────────┬─────────────┐│
│  │Docker Client │ File System  │   Temporary  │   Pattern   ││
│  │   Wrapper    │   Handlers   │   Storage    │  Matching   ││
│  └──────────────┴──────────────┴──────────────┴─────────────┘│
├─────────────────────────────────────────────────────────────┤
│                   External Dependencies                     │
│  ┌──────────────┬──────────────┬──────────────┬─────────────┐│
│  │Docker Engine │ File System  │   Network    │   Browser   ││
│  │     API      │              │              │             ││
│  └──────────────┴──────────────┴──────────────┴─────────────┘│
└─────────────────────────────────────────────────────────────┘
```

## Architecture Principles

### 1. Separation of Concerns
- **CLI Layer**: Command-line interface and argument parsing
- **Service Layer**: Business logic and core functionality
- **Infrastructure Layer**: External integrations and low-level operations

### 2. Dependency Injection
- Configurable components through interfaces
- Easy testing and mocking
- Flexible configuration management

### 3. Error Handling
- Explicit error returns throughout the system
- Contextual error wrapping
- Graceful degradation where possible

### 4. Security-First Design
- Docker socket access validation
- Image source verification
- Temporary file cleanup
- Input validation and sanitization

## Component Architecture

### CLI Layer (`cmd/`)

Built on the [Cobra](https://github.com/spf13/cobra) library for command-line interface management.

```go
// Command Structure
rootCmd (dockerutils)
├── serverCmd    // Web server management
├── versionCmd   // Version information
└── completion   // Shell completion
```

**Key Components:**
- `root.go` - Root command and global configuration
- `server.go` - Web server startup and management
- `completion.go` - Shell completion generation

### Service Layer (`src/`)

#### Analyzer Service (`src/analyzer/`)
Provides Docker image analysis capabilities with comprehensive metadata extraction.

**Core Types:**
```go
type AnalysisResult struct {
    ImageID      string // Docker image identifier
    ImageTag     string // Image tag (e.g., nginx:latest)
    ImageSource  string // Registry source
    ImageSize    int64  // Size in bytes
    Architecture string // Target architecture
    OS           string // Target operating system
    
    // Analysis results
    ExtractedPath   string // Filesystem extraction path
    ContainerFSPath string // Container filesystem path
    SavedTarPath    string // Saved image tar path
    
    // Status flags
    BuildSuccess    bool
    SaveSuccess     bool
    ExtractSuccess  bool
    ContainerFSSuccess bool
}
```

**Responsibilities:**
- Docker image inspection and metadata extraction
- Image layer analysis and history
- Filesystem extraction and container creation
- Analysis result formatting and reporting

#### Web Server Service (`src/webserver/`)
HTTP server providing REST API and web interface for Docker image analysis.

**Architecture:**
```go
type Server struct {
    config       *Config
    httpServer   *http.Server
    webRoot      string        // Static file root
    tmpDir       string        // Analysis data storage
    dockerClient *dockerclient.DockerClient
}
```

**API Endpoints:**
- `GET /` - Web interface (embedded HTML/CSS/JS)
- `POST /api/analyze` - Start image analysis
- `GET /api/summaries` - List analysis results
- `GET /api/buildcontext` - Build context analysis
- `DELETE /api/analysis/:id` - Cleanup analysis data

**Features:**
- Embedded static files using Go's `embed` package
- Asynchronous analysis operations
- Temporary file management and cleanup
- Cross-origin resource sharing (CORS) support

#### Docker Client Service (`src/dockerclient/`)
Enhanced wrapper around the official Docker SDK with additional security and validation.

**Key Features:**
```go
type DockerClient struct {
    client  client.APIClient  // Official Docker SDK client
    config  *Config          // Client configuration
    timeout time.Duration    // Operation timeout
}
```

**Security Features:**
- Docker socket access validation
- Image source verification against trusted registries
- Input validation and sanitization
- Security warnings for untrusted images

**Operations:**
- Image pulling and pushing
- Image inspection and metadata extraction
- Container creation and filesystem copying
- Image saving and tar file generation

#### Build Context Service (`src/buildcontext/`)
Analyzes Docker build contexts with `.dockerignore` pattern matching.

**Core Types:**
```go
type BuildDirectoryInfo struct {
    Path        string                         // Directory path
    Size        int64                          // Total size
    FileCount   int                            // File count
    Files       []BuildFileInfo                // Direct files
    Directories map[string]*BuildDirectoryInfo // Subdirectories
}
```

**Features:**
- Docker-compatible `.dockerignore` pattern matching
- Recursive directory tree analysis
- Size calculations and file counting
- Pattern exclusion with the official `moby/patternmatcher` library

#### Version Service (`src/version/`)
Provides build information and version management.

**Build Information:**
```go
type BuildInfo struct {
    Version   string // Version string (e.g., v1.0.0)
    GitCommit string // Git commit hash
    BuildTime string // Build timestamp
    GoVersion string // Go version used
    GOOS      string // Target OS
    GOARCH    string // Target architecture
}
```

## Data Flow

### 1. CLI Command Execution
```
User Input → Cobra Parser → Command Handler → Service Layer → Result Output
```

### 2. Web Interface Analysis
```
Browser Request → HTTP Router → API Handler → Analysis Service → Docker Client → Result Storage → JSON Response
```

### 3. Image Analysis Workflow
```
Image Name → Validation → Docker Pull → Inspection → Container Creation → Filesystem Copy → Tar Extraction → Analysis Result
```

### 4. Build Context Analysis
```
Directory Path → .dockerignore Reading → Pattern Parsing → File Walking → Exclusion Logic → Context Structure
```

## Technology Stack

### Core Technologies
- **Language**: Go 1.24.2+ (leveraging modern Go features)
- **CLI Framework**: Cobra v1.9.1+ (industry-standard CLI library)
- **Docker Integration**: Docker SDK v28.3.0+ (official Docker client library)
- **Pattern Matching**: Moby PatternMatcher v0.6.0+ (Docker-compatible patterns)

### Key Dependencies
```go
require (
    github.com/docker/docker v28.3.0+incompatible      // Docker SDK
    github.com/moby/patternmatcher v0.6.0               // .dockerignore patterns
    github.com/opencontainers/image-spec v1.1.1        // OCI image specification
    github.com/spf13/cobra v1.9.1                      // CLI framework
    golang.org/x/text v0.26.0                          // Text processing
)
```

### Web Technologies
- **Frontend**: Vanilla HTML5, CSS3, JavaScript (no framework dependencies)
- **Backend**: Go standard library HTTP server
- **Static Assets**: Embedded using Go's `embed` package
- **API**: RESTful JSON APIs with standard HTTP methods

## Security Architecture

### 1. Docker Socket Security
```go
func ValidateDockerAccess() error {
    // Platform-specific socket validation
    // Security warnings and risk assessment
    // Permission verification
}
```

### 2. Image Source Validation
```go
var trustedRegistries = []string{
    "docker.io",
    "gcr.io", 
    "quay.io",
    "registry.k8s.io",
    "ghcr.io",
}
```

### 3. Input Validation
- Image name format validation with regex patterns
- Path traversal prevention
- Size limitations and resource controls
- Timeout enforcement for all operations

### 4. Temporary File Security
- Isolated temporary directories
- Automatic cleanup on shutdown
- Permission restrictions
- Resource usage monitoring

## Performance Considerations

### 1. Asynchronous Operations
- Non-blocking image analysis
- Background processing with status tracking
- Timeout management for long-running operations

### 2. Resource Management
- Temporary file cleanup
- Memory usage optimization
- Docker container lifecycle management
- Efficient tar file processing

### 3. Caching Strategy
- Analysis result caching
- Static file serving optimization
- Docker layer reuse where possible

### 4. Scalability Design
- Stateless operation design
- Configurable resource limits
- Horizontal scaling capability through containerization

## Deployment Architecture

### 1. Standalone Binary
```bash
# Single executable with embedded static files
dockerutils server --port 8080
```

### 2. Container Deployment
```dockerfile
# Multi-stage build with security hardening
FROM golang:1.24-alpine AS builder
# ... build process
FROM alpine:latest
# ... runtime configuration with non-root user
```

### 3. Docker Compose
```yaml
services:
  dockerutils:
    build: .
    ports: ["8080:8080"]
    volumes: ["/var/run/docker.sock:/var/run/docker.sock"]
    environment: ["DOCKERUTILS_PORT=8080"]
```

### 4. Kubernetes Deployment
- StatefulSet for persistent analysis storage
- ConfigMap for configuration management
- Service for network exposure
- PersistentVolume for data storage

### Security Considerations in Deployment
- Non-root container execution
- Docker socket mounting with appropriate permissions
- Network security and firewall configuration
- Resource quotas and limits

---

**Related Documentation:**
- [Development Guide](DEVELOPMENT.md) - Development environment setup
- [Deployment Guide](DEPLOYMENT.md) - Production deployment instructions
- [Security Policy](SECURITY.md) - Security guidelines and procedures
- [API Reference](API_REFERENCE.md) - Complete API documentation

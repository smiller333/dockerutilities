# Docker Build Process Documentation

## Overview

This document describes the current Docker build process for the `dockerutilities` project. The Docker build process is currently separate from the main build system and will be integrated with GoReleaser in Phase 4 of the automated build planning.

## Current Docker Build Process

### Dockerfile Structure

The current Dockerfile uses a multi-stage build approach with security hardening:

```dockerfile
# Build stage
FROM golang:1.24-alpine AS builder
# ... build process

# Runtime stage  
FROM alpine:latest
# ... runtime setup
```

### Key Features

#### 1. Multi-Stage Build
- **Build Stage**: Uses `golang:1.24-alpine` for compilation
- **Runtime Stage**: Uses `alpine:latest` for minimal runtime footprint
- **Security**: Non-root user with proper permissions

#### 2. Security Hardening
- **Non-Root User**: Creates `dockerutilities` user (UID 1001)
- **Docker Socket Access**: Dynamic group assignment for Docker socket permissions
- **Minimal Dependencies**: Only essential runtime packages

#### 3. Dynamic Docker Socket Handling
- **Entrypoint Script**: Complex logic for handling Docker socket permissions
- **Group Detection**: Automatically detects Docker socket group
- **User Assignment**: Adds user to appropriate group for socket access

### Current Limitations

#### 1. Version Injection Missing
- **No Version Embedding**: Docker builds don't include version information
- **No Git Tag Support**: Doesn't use git tags for versioning
- **Separate Process**: Not integrated with main build system

#### 2. Platform Limitations
- **Linux Only**: Only builds for Linux platform
- **No Multi-Architecture**: No ARM64 support
- **No Cross-Platform**: Limited to single platform builds

#### 3. Build Process Isolation
- **Separate from Makefile**: Not integrated with main build process
- **Manual Process**: Requires manual Docker build commands
- **No Automation**: Not part of automated build pipeline

## Build Commands

### Current Docker Build Commands

```bash
# Build Docker image
make docker-build

# Build and run Docker container
make docker-run

# Manual Docker commands
docker build -t dockerutilities:latest .
docker run --rm -it dockerutilities:latest
```

### Docker Run Script

The `scripts/docker-run.sh` script provides additional functionality:

```bash
# Run with auto-detected Docker socket
./scripts/docker-run.sh

# Run with persistent data
./scripts/docker-run.sh --persistent

# Run with custom socket
./scripts/docker-run.sh --socket /path/to/docker.sock
```

## Future Integration Plans (Phase 4)

### GoReleaser Integration

The Docker build process will be integrated with GoReleaser in Phase 4:

#### Planned Configuration
```yaml
# .goreleaser.yml (Phase 4)
dockers:
  - image_templates:
      - "ghcr.io/smiller333/dockerutilities:{{ .Version }}"
      - "ghcr.io/smiller333/dockerutilities:latest"
    dockerfile: Dockerfile
    use: buildx
    build_flag_templates:
      - "--platform=linux/amd64"
      - "--platform=linux/arm64"
```

#### Benefits of Integration
1. **Multi-Platform Support**: Automatic ARM64 and AMD64 builds
2. **Version Injection**: Proper version embedding from git tags
3. **Automated Publishing**: Integration with GitHub Container Registry
4. **Consistent Builds**: Unified build process with GoReleaser

### Planned Improvements

#### 1. Version Integration
- **Git Tag Support**: Use git tags for Docker image versioning
- **Version Embedding**: Include version information in Docker builds
- **Consistent Naming**: Align with GoReleaser naming conventions

#### 2. Multi-Platform Support
- **ARM64 Support**: Add ARM64 builds for modern architectures
- **Buildx Integration**: Use Docker Buildx for multi-platform builds
- **Platform Matrix**: Support Linux AMD64 and ARM64

#### 3. Registry Integration
- **GitHub Container Registry**: Primary registry for Docker images
- **Docker Hub**: Secondary registry for broader distribution
- **Automated Publishing**: Automatic image publishing on releases

## Current Docker Configuration

### Entrypoint Script

The current entrypoint script handles dynamic Docker socket permissions:

```bash
#!/bin/sh
# Handle Docker socket permissions dynamically
if [ -S /var/run/docker.sock ]; then
    DOCKER_SOCKET_GID=$(stat -c '%g' /var/run/docker.sock)
    # ... group creation and user assignment
fi
```

### Health Check

```dockerfile
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/health || exit 1
```

### Volume Mounts

```dockerfile
# Create volume mount point for persistent data
VOLUME ["/app/data"]
```

## Security Considerations

### Current Security Features

1. **Non-Root User**: Runs as `dockerutilities` user (UID 1001)
2. **Minimal Base Image**: Uses Alpine Linux for smaller attack surface
3. **Dynamic Permissions**: Handles Docker socket permissions securely
4. **Health Checks**: Monitors application health

### Planned Security Enhancements

1. **Image Scanning**: Integrate vulnerability scanning
2. **Multi-Stage Optimization**: Further reduce image size
3. **Security Policies**: Implement security policies for registry
4. **Signed Images**: Add image signing for authenticity

## Testing Docker Builds

### Local Testing

```bash
# Test Docker build
make docker-build

# Test Docker run
make docker-run

# Test with custom arguments
docker run --rm -it dockerutilities:latest --help
```

### Integration Testing

```bash
# Test Docker socket access
docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock dockerutilities:latest

# Test persistent data
docker run --rm -it -v ./data:/app/data dockerutilities:latest
```

## Maintenance Notes

### Current Maintenance Tasks

1. **Base Image Updates**: Keep Alpine and Go base images updated
2. **Security Patches**: Monitor for security vulnerabilities
3. **Dependency Updates**: Update runtime dependencies as needed
4. **Entrypoint Script**: Maintain Docker socket handling logic

### Future Maintenance (Phase 4)

1. **GoReleaser Integration**: Maintain GoReleaser Docker configuration
2. **Multi-Platform Testing**: Test builds across all supported platforms
3. **Registry Management**: Manage multiple registry configurations
4. **Automated Updates**: Integrate with automated dependency updates

## Conclusion

The current Docker build process provides a solid foundation with security hardening and dynamic Docker socket handling. While it has limitations in version injection and platform support, these will be addressed in Phase 4 through GoReleaser integration.

The process is well-documented and ready for future enhancement without disrupting the current build system simplification efforts in Phase 0.

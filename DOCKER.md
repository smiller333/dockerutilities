# Docker Deployment for Docker Utils

This guide covers deploying and running Docker Utils in containerized environments using Docker, Docker Compose, and container orchestration platforms.

## Quick Start

### Option 1: Using the Build Script (Recommended)

```bash
# Build and run with ephemeral storage
./scripts/docker-run.sh run

# Build and run with persistent data storage
./scripts/docker-run.sh run --persistent
# or use the legacy command:
./scripts/docker-run.sh run-persistent

# Use a custom Docker socket location (e.g., Docker Desktop on macOS)
./scripts/docker-run.sh run --socket ~/.docker/desktop/docker.sock

# Combine options for persistent storage with custom socket
./scripts/docker-run.sh run --persistent --socket ~/.docker/desktop/docker.sock

# Use custom port and data directory
./scripts/docker-run.sh run --persistent --port 3000 --data-dir ./my-data

# Available commands:
./scripts/docker-run.sh help          # Show help and usage
./scripts/docker-run.sh build         # Build Docker image only
./scripts/docker-run.sh run           # Build and run container
./scripts/docker-run.sh logs          # View container logs
./scripts/docker-run.sh shell         # Open shell in running container
./scripts/docker-run.sh stop          # Stop the container
./scripts/docker-run.sh clean         # Remove container and image
```

**Environment Variables:**
```bash
# Set custom Docker socket via environment variable
export DOCKER_SOCKET=~/.docker/desktop/docker.sock
./scripts/docker-run.sh run

# Common Docker socket locations:
# Linux (default):           /var/run/docker.sock
# Docker Desktop (macOS):    ~/.docker/desktop/docker.sock  
# Docker Desktop (Windows):  //./pipe/docker_engine
# Podman:                    /run/user/$(id -u)/podman/podman.sock
```

### Option 2: Using Docker Compose

> **Note**: Docker Compose files are not included in the current version. Use the build script (Option 1) or manual Docker commands (Option 3).

For a basic Docker Compose setup, create `docker-compose.yml`:

```yaml
version: '3.8'
services:
  dockerutils:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - ./data:/app/data  # Optional: for persistent storage
    environment:
      - HOST=0.0.0.0
      - PORT=8080
    restart: unless-stopped
```

Then run:
```bash
docker-compose up -d
```

### Option 3: Manual Docker Commands

```bash
# Build the image
docker build -t dockerutils-viewer .

# Run with ephemeral storage
docker run -d \
  --name dockerutils-viewer \
  -p 8080:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  dockerutils-viewer

# Run with persistent storage
docker run -d \
  --name dockerutils-viewer \
  -p 8080:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(pwd)/data:/app/data \
  dockerutils-viewer
```

## Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `HOST` | `0.0.0.0` | Host/IP address to bind the server to |
| `PORT` | `8080` | Port number for the web server |
| `DOCKER_HOST` | `unix:///var/run/docker.sock` | Docker daemon endpoint |
| `DOCKER_SOCKET` | `/var/run/docker.sock` | Docker socket path for container mount |

**Note:** `DOCKER_SOCKET` is used by the `scripts/docker-run.sh` script to determine which socket to mount into the container, while `DOCKER_HOST` is used by the Docker client itself.

### Volume Mounts

| Host Path | Container Path | Purpose |
|-----------|----------------|---------|
| `/var/run/docker.sock` | `/var/run/docker.sock` | Docker socket for Docker API access |
| `./data` | `/app/data` | Persistent storage for analysis data (optional) |

### Ports

| Container Port | Host Port | Description |
|----------------|-----------|-------------|
| `8080` | `8080` | Web interface and API |

## Data Persistence

### Ephemeral Storage (Default)
- Analysis data is stored inside the container
- Data is lost when the container is removed
- Suitable for testing and development

### Persistent Storage
- Analysis data is stored in a host directory (`./data`)
- Data persists across container restarts and updates
- Recommended for production use

To enable persistent storage:

1. **Using the build script:**
   ```bash
   ./scripts/docker-run.sh run-persistent
   ```

2. **Using Docker Compose:**
   ```bash
   # Edit docker-compose.yml and uncomment the data volume line
   # - ./data:/app/data
   docker-compose up -d
   ```

3. **Using manual Docker commands:**
   ```bash
   mkdir -p ./data
   docker run -d \
     --name dockerutils-viewer \
     -p 8080:8080 \
     -v /var/run/docker.sock:/var/run/docker.sock \
     -v $(pwd)/data:/app/data \
     dockerutils-viewer
   ```

## Access

Once the container is running, you can access:

- **Web Interface:** http://localhost:8080
- **API Health Check:** http://localhost:8080/api/health
- **API Documentation:** Available through the web interface

## Troubleshooting

### Container Won't Start

1. **Check if port 8080 is available:**
   ```bash
   lsof -i :8080
   ```

2. **Check Docker daemon is running:**
   ```bash
   docker version
   ```

3. **View container logs:**
   ```bash
   ./scripts/docker-run.sh logs
   # or
   docker logs dockerutils-viewer
   ```

### Cannot Access Docker Images

1. **Verify Docker socket mount:**
   ```bash
   docker exec dockerutils-viewer ls -la /var/run/docker.sock
   ```

2. **Check Docker socket permissions:**
   ```bash
   ls -la /var/run/docker.sock
   ```

3. **For non-standard Docker socket locations:**
   ```bash
   # Check if your Docker socket is in a different location
   # Docker Desktop on macOS:
   ls -la ~/.docker/desktop/docker.sock
   
   # Podman:
   ls -la /run/user/$(id -u)/podman/podman.sock
   
   # Use the --socket option to specify the correct path:
   ./scripts/docker-run.sh run --socket ~/.docker/desktop/docker.sock
   ```

4. **Test Docker connectivity:**
   ```bash
   # Test if Docker is accessible from the host
   docker version
   
   # If using a custom socket, test with DOCKER_HOST
   DOCKER_HOST=unix://~/.docker/desktop/docker.sock docker version
   ```

### Data Not Persisting

1. **Check volume mount:**
   ```bash
   docker inspect dockerutils-viewer | grep -A 10 "Mounts"
   ```

2. **Verify data directory permissions:**
   ```bash
   ls -la ./data
   ```

### Performance Issues

1. **Check container resources:**
   ```bash
   docker stats dockerutils-viewer
   ```

2. **Increase memory limits in docker-compose.prod.yml if needed**

## Security Considerations

1. **Docker Socket Access:** The container requires access to the Docker socket to analyze images. This grants significant privileges.

2. **Network Binding:** By default, the container binds to all interfaces (`0.0.0.0`). For production, consider:
   - Using a reverse proxy (nginx, traefik)
   - Binding to specific interfaces only
   - Implementing authentication

3. **User Permissions:** The container runs as a non-root user (`dockerutils:dockerutils`) for security.

## Development

### Building During Development

```bash
# Build image
./scripts/docker-run.sh build

# For development, use local mode instead of Docker:
go run main.go tools --port 8080
```

### Debugging

```bash
# Open shell in running container
./scripts/docker-run.sh shell

# Run container with shell access
docker run -it --rm \
  -v /var/run/docker.sock:/var/run/docker.sock \
  dockerutils-viewer /bin/sh
```

## Production Deployment

For production deployment, consider:

1. **Use Docker Compose for production:**
   ```bash
   # Create production docker-compose.yml with proper resource limits
   docker-compose up -d
   ```

2. **Enable persistent storage**
3. **Set up log rotation**
4. **Configure monitoring and health checks**
5. **Use a reverse proxy for HTTPS**
6. **Implement backup strategies for analysis data**

## Related Files

- `Dockerfile` - Multi-stage Docker build configuration
- `scripts/docker-run.sh` - Convenience script for building and running
- `.dockerignore` - Files to exclude from Docker build context

## Additional Resources

- [Installation Guide](INSTALLATION.md) - Comprehensive installation instructions
- [User Guide](docs/USER_GUIDE.md) - Complete usage documentation
- [Troubleshooting](TROUBLESHOOTING.md) - Common Docker issues and solutions
- [Examples](EXAMPLES.md) - Docker deployment examples

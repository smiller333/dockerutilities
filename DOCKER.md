# Docker Deployment

This guide covers running Docker Utils in a containerized environment.

## Quick Start

### Using the Build Script

```bash
# Build and run with persistent storage
./scripts/docker-run.sh run-persistent

# Build and run with ephemeral storage
./scripts/docker-run.sh run

# Use custom port
./scripts/docker-run.sh run --port 3000

# Available commands:
./scripts/docker-run.sh help          # Show help
./scripts/docker-run.sh build         # Build Docker image
./scripts/docker-run.sh run           # Build and run container
./scripts/docker-run.sh logs          # View container logs
./scripts/docker-run.sh stop          # Stop the container
./scripts/docker-run.sh clean         # Remove container and image
```

### Manual Docker Commands

```bash
# Build the image
docker build -t dockerutilities .

# Run with persistent storage
docker run -d \
  --name dockerutilities \
  -p 8080:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(pwd)/data:/app/data \
  dockerutilities

# Run with ephemeral storage
docker run -d \
  --name dockerutilities \
  -p 8080:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  dockerutilities
```

## Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `HOST` | `0.0.0.0` | Host/IP address to bind the server to |
| `PORT` | `8080` | Port number for the web server |

### Volume Mounts

| Host Path | Container Path | Purpose |
|-----------|----------------|---------|
| `/var/run/docker.sock` | `/var/run/docker.sock` | Docker socket for API access |
| `./data` | `/app/data` | Persistent storage (optional) |

## Access

Once the container is running, access the web interface at:
- **Web Interface:** http://localhost:8080

You can check Docker connectivity status by clicking the "System Status" button in the web interface, which will show:
- Docker daemon connection status
- Docker version and system information
- Container and image counts
- System architecture and resources

## Troubleshooting

### Common Issues

1. **Port already in use:**
   ```bash
   # Use a different port
   ./scripts/docker-run.sh run --port 3000
   ```

2. **Docker socket not found:**
   ```bash
   # Check if Docker is running
   docker version
   
   # Verify socket exists
   ls -la /var/run/docker.sock
   ```

3. **View container logs:**
   ```bash
   ./scripts/docker-run.sh logs
   ```

## Development

For development, you can run locally instead of using Docker:

```bash
go run main.go server --port 8080
```

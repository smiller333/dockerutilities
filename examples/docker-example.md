# Example: Running DockerUtils Image Viewer in Docker

This example demonstrates how to run the dockerutils image-viewer server in a Docker container with different storage configurations.

## Quick Demo

### 1. Build and Run with Ephemeral Storage

```bash
# Build the Docker image and run with ephemeral storage
./docker-run.sh run

# The server will be available at http://localhost:8080
# Analysis data will be lost when the container stops
```

### 2. Build and Run with Persistent Storage

```bash
# Build the Docker image and run with persistent storage
./docker-run.sh run-persistent

# The server will be available at http://localhost:8080
# Analysis data will be preserved in ./data directory
```

### 3. Using Docker Compose

```bash
# Start with docker-compose (ephemeral storage)
docker-compose up -d

# Start with docker-compose (persistent storage)
docker-compose -f docker-compose.prod.yml up -d
```

### 4. Manual Docker Commands

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
mkdir -p ./data
docker run -d \
  --name dockerutils-viewer \
  -p 8080:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(pwd)/data:/app/data \
  dockerutils-viewer
```

## Testing the Server

1. **Access the web interface**:
   - Open http://localhost:8080 in your browser
   - Check the health endpoint: http://localhost:8080/api/health

2. **Check container logs**:
   ```bash
   ./docker-run.sh logs
   # or
   docker logs dockerutils-viewer
   ```

3. **Open shell in container** (for debugging):
   ```bash
   ./docker-run.sh shell
   # or
   docker exec -it dockerutils-viewer /bin/sh
   ```

## Storage Comparison

### Ephemeral Storage
- **Use case**: Testing, development, temporary analysis
- **Data persistence**: Data is lost when container stops/restarts
- **Command**: `./docker-run.sh run`

### Persistent Storage
- **Use case**: Production, long-term analysis storage
- **Data persistence**: Data persists across container restarts
- **Command**: `./docker-run.sh run-persistent`
- **Data location**: `./data` directory on host

## Stopping and Cleanup

```bash
# Stop the container
./docker-run.sh stop

# Remove container and image
./docker-run.sh clean

# Stop docker-compose
docker-compose down
```

## Configuration Options

### Environment Variables

You can customize the server behavior with environment variables:

```bash
# Example with custom configuration
docker run -d \
  --name dockerutils-viewer \
  -p 3000:3000 \
  -e HOST=0.0.0.0 \
  -e PORT=3000 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(pwd)/data:/app/data \
  dockerutils-viewer
```

### Volume Mounts

The container supports these volume mounts:

- `/var/run/docker.sock`: Required for Docker API access
- `/app/data`: Optional persistent storage for analysis data

## Troubleshooting

### Container Won't Start

```bash
# Check Docker daemon is running
docker version

# Check if port is available
lsof -i :8080

# View container logs
docker logs dockerutils-viewer
```

### Can't Access Docker Images

```bash
# Verify Docker socket mount
docker exec dockerutils-viewer ls -la /var/run/docker.sock

# Check Docker socket permissions on host
ls -la /var/run/docker.sock
```

### Data Not Persisting

```bash
# Check volume mount
docker inspect dockerutils-viewer | grep -A 10 "Mounts"

# Verify data directory permissions
ls -la ./data
```

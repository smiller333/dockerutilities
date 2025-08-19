# Docker Utils User Guide

This comprehensive guide will walk you through all the features and capabilities of Docker Utils, from basic image analysis to advanced API integration.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Web Interface Tutorial](#web-interface-tutorial)
3. [Command Line Usage](#command-line-usage)
4. [Common Use Cases](#common-use-cases)
5. [Advanced Features](#advanced-features)
6. [API Integration](#api-integration)
7. [Troubleshooting](#troubleshooting)
8. [Tips and Best Practices](#tips-and-best-practices)

## Getting Started

### Prerequisites

Before starting, ensure you have:
- Docker Engine running on your system
- Access to Docker socket (usually automatic if Docker is properly installed)
- Sufficient disk space for image analysis (varies by image size)

### Quick Setup

Choose your preferred method:

#### Option 1: Docker (Recommended)
```bash
git clone https://github.com/smiller333/dockerutils.git
cd dockerutils
./scripts/docker-run.sh run-persistent
```

#### Option 2: Build from Source
```bash
git clone https://github.com/smiller333/dockerutils.git
cd dockerutils
make build-release
./bin/dockerutils server
```

#### Option 3: Quick Development Build
```bash
go build -o dockerutils
./dockerutils server
```

Once running, open http://localhost:8080 in your browser.

## Web Interface Tutorial

### First Steps

1. **Start the Server**: Run `dockerutils server` and wait for the browser to open
2. **Check Connection**: The interface will show if Docker is accessible
3. **View Dashboard**: See any existing analysis results (if any)

### Analyzing Your First Image

Let's analyze the popular Alpine Linux image:

1. **Enter Image Name**: Type `alpine:latest` in the image input field
2. **Click Analyze**: The analysis will start immediately
3. **Monitor Progress**: Watch the progress indicator
4. **View Results**: Once complete, click on the result card to explore

#### What You'll See

The analysis provides:
- **Basic Information**: Image size, architecture, creation date
- **Layer Details**: Each layer with size and commands
- **File System**: Browse the complete extracted filesystem
- **Environment**: All environment variables and configurations
- **Metadata**: Labels, exposed ports, working directory, etc.

### Exploring Analysis Results

#### Image Gallery View
- Each analyzed image appears as a card
- Cards show key metadata: size, architecture, analysis date
- Color coding indicates analysis status (green = complete, yellow = in progress, red = failed)

#### Detailed View
Click any image card to open the detailed analysis:

1. **Overview Tab**: Essential image information and statistics
2. **Layers Tab**: Detailed layer breakdown with commands and sizes
3. **FileSystem Tab**: Interactive file browser with download capability
4. **Environment Tab**: All environment variables, labels, and configurations
5. **Raw Data Tab**: Complete JSON analysis data

#### File System Browser
- Navigate through directories like a file manager
- Click files to view content (text files only)
- Download individual files or entire directories
- View file permissions and ownership

### Managing Analysis Results

#### Cleanup
- Click the trash icon on any image card to delete results
- This frees up disk space and removes temporary files
- Deleted analyses can be re-run at any time

#### Bulk Operations
- Use the "Clear All" button to remove all analysis results
- Individual file cleanup through the API (see API section)

## Command Line Usage

### Basic Commands

#### Version Information
```bash
# Quick version
dockerutils -v

# Detailed version info
dockerutils version
```

#### Help and Documentation
```bash
# General help
dockerutils --help

# Command-specific help
dockerutils server --help
dockerutils completion --help
```

#### Starting the Web Server
```bash
# Default settings (localhost:8080)
dockerutils server

# Custom port
dockerutils server --port 3000

# Bind to all interfaces
dockerutils server --host 0.0.0.0 --port 8080

# Custom directories
dockerutils server --tmp-dir /app/data --web-root /custom/ui

# Don't open browser automatically
dockerutils server --no-browser
```

### Command Options Reference

#### Global Options
- `-h, --help`: Show help information
- `-v, --version`: Show version information

#### Server Command Options
- `--port string`: Port number for web server (default: "8080")
- `--host string`: Host/IP to bind server to (default: "localhost")
- `--tmp-dir string`: Directory for analysis data (default: "./tmp")
- `--web-root string`: Custom web UI directory (optional)
- `--no-browser`: Don't automatically open browser

### Shell Completion

Generate shell completion scripts:

```bash
# Bash
dockerutils completion bash > /etc/bash_completion.d/dockerutils

# Zsh
dockerutils completion zsh > "${fpath[1]}/_dockerutils"

# Fish
dockerutils completion fish > ~/.config/fish/completions/dockerutils.fish

# PowerShell
dockerutils completion powershell > dockerutils.ps1
```

## Common Use Cases

### Security Analysis

#### Scanning for Vulnerabilities
```bash
# Start dockerutils
dockerutils server

# Analyze a potentially vulnerable image
# Navigate to: http://localhost:8080
# Enter: nginx:1.14  # Older version for demonstration
# Examine the results for:
# - Outdated packages in filesystem
# - Exposed ports and services
# - File permissions and ownership
# - Environment variables with sensitive data
```

#### Comparing Image Versions
```bash
# Analyze multiple versions
# In the web interface, analyze:
# - alpine:3.15
# - alpine:3.16
# - alpine:latest
# Compare file systems, packages, and configurations
```

### Development and Debugging

#### Understanding Build Layers
When your Docker build is slow or large:

1. Analyze your built image
2. Review each layer in the "Layers" tab
3. Identify layers contributing most to size
4. Examine the commands that created large layers
5. Optimize your Dockerfile based on findings

#### Debugging Runtime Issues
When containers behave unexpectedly:

1. Analyze the problematic image
2. Check environment variables
3. Examine the filesystem for missing files
4. Verify file permissions
5. Check the working directory and entry point

### CI/CD Integration

#### Automated Security Scanning
```bash
#!/bin/bash
# ci-security-scan.sh

IMAGE_NAME="$1"
REPORT_DIR="./security-reports"

# Start dockerutils in background
dockerutils server --no-browser --port 8080 &
SERVER_PID=$!
sleep 5  # Wait for server to start

# Analyze the image via API
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d "{\"image_name\": \"$IMAGE_NAME\"}" \
  -o "$REPORT_DIR/analysis.json"

# Check if analysis succeeded
if jq -r '.success' "$REPORT_DIR/analysis.json" | grep -q true; then
  echo "✅ Analysis completed for $IMAGE_NAME"
  # Extract image ID for detailed results
  IMAGE_ID=$(jq -r '.image_id' "$REPORT_DIR/analysis.json")
  
  # Get detailed analysis
  curl "http://localhost:8080/api/info/$IMAGE_ID" \
    -o "$REPORT_DIR/detailed-$IMAGE_ID.json"
  
  # Process results (example: check for specific packages)
  if jq '.filesystem' "$REPORT_DIR/detailed-$IMAGE_ID.json" | grep -q "vulnerable-package"; then
    echo "❌ Vulnerable package detected"
    exit 1
  fi
else
  echo "❌ Analysis failed for $IMAGE_NAME"
  exit 1
fi

# Cleanup
kill $SERVER_PID
```

#### Size Optimization Reports
```bash
#!/bin/bash
# size-report.sh

IMAGES=("myapp:v1.0" "myapp:v1.1" "myapp:latest")
REPORT_FILE="size-comparison.md"

echo "# Image Size Comparison Report" > $REPORT_FILE
echo "Generated: $(date)" >> $REPORT_FILE
echo "" >> $REPORT_FILE

for image in "${IMAGES[@]}"; do
  # Analyze each image and extract size information
  # Implementation would call the API and format results
  echo "## $image" >> $REPORT_FILE
  echo "Analysis results..." >> $REPORT_FILE
done
```

### Container Optimization

#### Finding Unnecessary Files
1. Analyze your production image
2. Browse the filesystem to identify:
   - Package manager caches (`/var/cache/`, `/tmp/`)
   - Development tools (compilers, build tools)
   - Documentation files (`/usr/share/doc/`, `/usr/share/man/`)
   - Log files and temporary data

#### Multi-stage Build Verification
1. Analyze your final image
2. Verify only necessary files are present
3. Check that build artifacts are not included
4. Confirm proper file ownership and permissions

## Advanced Features

### API Automation

#### Batch Image Analysis
```python
#!/usr/bin/env python3
import requests
import json
import time

# List of images to analyze
images = [
    "nginx:latest",
    "alpine:latest", 
    "ubuntu:20.04",
    "node:16-alpine"
]

base_url = "http://localhost:8080/api"

for image in images:
    print(f"Analyzing {image}...")
    
    # Start async analysis
    response = requests.post(f"{base_url}/analyze-async", 
                           json={"image_name": image})
    
    if response.status_code == 200:
        result = response.json()
        request_id = result["request_id"]
        print(f"Started analysis with ID: {request_id}")
    else:
        print(f"Failed to start analysis: {response.text}")
        continue

# Wait for all analyses to complete
print("Waiting for analyses to complete...")
time.sleep(30)

# Check results
summaries = requests.get(f"{base_url}/summaries").json()
for summary in summaries:
    print(f"Image: {summary['image_tag']}")
    print(f"Size: {summary['image_size']} bytes")
    print(f"Architecture: {summary['architecture']}")
    print("---")
```

#### Custom Analysis Scripts
```bash
#!/bin/bash
# custom-analysis.sh

API_BASE="http://localhost:8080/api"
IMAGE_NAME="$1"

# Analyze image
echo "Analyzing $IMAGE_NAME..."
RESULT=$(curl -s -X POST "$API_BASE/analyze" \
  -H "Content-Type: application/json" \
  -d "{\"image_name\": \"$IMAGE_NAME\"}")

IMAGE_ID=$(echo "$RESULT" | jq -r '.image_id')

if [ "$IMAGE_ID" != "null" ]; then
  # Get detailed analysis
  DETAILS=$(curl -s "$API_BASE/info/$IMAGE_ID")
  
  # Extract specific information
  echo "=== Analysis Results ==="
  echo "Image Size: $(echo "$DETAILS" | jq -r '.image_size') bytes"
  echo "Architecture: $(echo "$DETAILS" | jq -r '.architecture')"
  echo "Layers: $(echo "$DETAILS" | jq -r '.layers | length')"
  echo "Created: $(echo "$DETAILS" | jq -r '.created')"
  
  # Check for specific packages or files
  echo "=== Security Check ==="
  if echo "$DETAILS" | jq -r '.filesystem' | grep -q "ssh"; then
    echo "⚠️  SSH found in image"
  fi
  
  if echo "$DETAILS" | jq -r '.environment.env[]' | grep -q "PASSWORD"; then
    echo "⚠️  Password environment variable detected"
  fi
  
  echo "=== Size Breakdown ==="
  echo "$DETAILS" | jq -r '.layers[] | "\(.size) bytes - \(.command)"' | head -10
else
  echo "❌ Analysis failed"
  exit 1
fi
```

### Integration with Other Tools

#### Docker Compose Analysis
```yaml
# docker-compose.analysis.yml
version: '3.8'
services:
  dockerutils:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - ./analysis-data:/app/tmp
    command: server --host 0.0.0.0
    
  analyzer:
    image: curlimages/curl
    depends_on:
      - dockerutils
    volumes:
      - ./scripts:/scripts
    command: /scripts/batch-analyze.sh
```

#### Kubernetes Integration
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dockerutils
spec:
  replicas: 1
  selector:
    matchLabels:
      app: dockerutils
  template:
    metadata:
      labels:
        app: dockerutils
    spec:
      containers:
      - name: dockerutils
        image: dockerutils:latest
        ports:
        - containerPort: 8080
        volumeMounts:
        - name: docker-socket
          mountPath: /var/run/docker.sock
        command: ["dockerutils", "server", "--host", "0.0.0.0"]
      volumes:
      - name: docker-socket
        hostPath:
          path: /var/run/docker.sock
---
apiVersion: v1
kind: Service
metadata:
  name: dockerutils-service
spec:
  selector:
    app: dockerutils
  ports:
  - port: 80
    targetPort: 8080
  type: LoadBalancer
```

## API Integration

### REST API Overview

The Docker Utils API provides programmatic access to all image analysis functionality:

- **Base URL**: `http://localhost:8080/api`
- **Content-Type**: `application/json`
- **Authentication**: None (local access only)

### Core Endpoints

#### Health Check
```bash
curl http://localhost:8080/api/health
```

#### List All Analyses
```bash
curl http://localhost:8080/api/summaries
```

#### Get Detailed Analysis
```bash
curl http://localhost:8080/api/info/{image_id}
```

#### Analyze Image (Synchronous)
```bash
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "image_name": "nginx:latest",
    "keep_temp_files": true,
    "force_pull": false
  }'
```

#### Analyze Image (Asynchronous)
```bash
curl -X POST http://localhost:8080/api/analyze-async \
  -H "Content-Type: application/json" \
  -d '{
    "image_name": "nginx:latest",
    "keep_temp_files": true,
    "force_pull": false
  }'
```

#### Delete Analysis Results
```bash
curl -X DELETE http://localhost:8080/api/info/{image_id}
```

### API Response Examples

#### Successful Analysis Response
```json
{
  "success": true,
  "image_id": "sha256:7aab056cecc6",
  "message": "Image analysis completed successfully",
  "analysis_time": 15.432,
  "size_bytes": 7738912
}
```

#### Error Response
```json
{
  "success": false,
  "error": "Image not found",
  "message": "The specified image 'nonexistent:latest' could not be found in any registry",
  "code": 404
}
```

#### Summary List Response
```json
[
  {
    "image_id": "sha256:7aab056cecc6",
    "image_tag": "nginx:latest",
    "image_source": "docker.io",
    "image_size": 142984192,
    "architecture": "amd64",
    "analyzed_at": "2025-01-21T14:47:10Z",
    "status": "completed",
    "request_id": "7aab056cecc6"
  }
]
```

## Troubleshooting

### Common Issues and Solutions

#### Docker Connection Problems

**Problem**: "Cannot connect to Docker daemon"
```bash
# Check if Docker is running
docker info

# Check Docker socket permissions (Linux/macOS)
ls -la /var/run/docker.sock

# Add user to docker group (Linux)
sudo usermod -aG docker $USER
# Then logout and login again
```

**Problem**: "Permission denied on Docker socket"
```bash
# Temporary solution (not recommended for production)
sudo chmod 666 /var/run/docker.sock

# Better solution: ensure user is in docker group
groups $USER
# Should show 'docker' in the output
```

#### Port and Network Issues

**Problem**: "Port 8080 already in use"
```bash
# Find what's using the port
lsof -i :8080
# Or on Linux
netstat -tulpn | grep 8080

# Use a different port
dockerutils server --port 3000

# Or kill the process using the port
sudo kill -9 $(lsof -ti:8080)
```

**Problem**: "Cannot access from remote machine"
```bash
# Bind to all interfaces instead of localhost
dockerutils server --host 0.0.0.0 --port 8080

# Check firewall settings
# Linux (ufw)
sudo ufw allow 8080

# macOS (built-in firewall)
# Go to System Preferences > Security & Privacy > Firewall
```

#### Analysis Failures

**Problem**: "Image analysis failed"
```bash
# Check if image exists and is accessible
docker pull <image-name>

# Check available disk space
df -h

# Check tmp directory permissions
ls -la ./tmp

# Try with a smaller, simpler image first
# Use alpine:latest as a test
```

**Problem**: "Extraction failed"
```bash
# Check tmp directory has write permissions
chmod 755 ./tmp

# Ensure sufficient disk space (at least 2x image size)
df -h

# Try cleaning up old analysis results
# Delete files in ./tmp/ directory
```

#### Performance Issues

**Problem**: "Analysis is very slow"
```bash
# Use faster storage for tmp directory
dockerutils server --tmp-dir /path/to/fast/storage

# Analyze smaller images first
# Avoid very large images (>1GB) for testing

# Check system resources
top
iostat 1
```

**Problem**: "High memory usage"
```bash
# Monitor memory usage
htop

# Close other applications
# Consider using smaller images for analysis

# Restart dockerutils periodically for long-running analyses
```

### Debug Mode

Enable verbose logging:
```bash
# Set environment variable for detailed logging
export DOCKER_UTILS_DEBUG=1
dockerutils server

# Or check Docker daemon logs
journalctl -u docker.service -f  # Linux systemd
# Or
tail -f /var/log/docker.log      # Some systems
```

### Getting Help

1. **Check Logs**: Look at console output when running dockerutils
2. **Verify Prerequisites**: Ensure Docker is running and accessible
3. **Test with Simple Image**: Try analyzing `alpine:latest` first
4. **Check Documentation**: Review this guide and the API documentation
5. **Report Issues**: Use GitHub issues for bugs and feature requests

## Tips and Best Practices

### Performance Optimization

#### Choose the Right Analysis Method
- **Synchronous** (`/api/analyze`): For small images (<100MB) or when you need immediate results
- **Asynchronous** (`/api/analyze-async`): For large images or batch processing

#### Manage Disk Space
```bash
# Regular cleanup of analysis results
curl -X DELETE http://localhost:8080/api/info/{old_image_id}

# Monitor tmp directory size
du -sh ./tmp

# Use dedicated storage for tmp directory
dockerutils server --tmp-dir /dedicated/fast/storage
```

#### Network Optimization
```bash
# Pre-pull images to avoid network delays during analysis
docker pull nginx:latest
docker pull alpine:latest

# Use local registry for private images
# Reduces network latency and external dependencies
```

### Security Best Practices

#### Local Development Only
- Docker Utils is designed for local development use
- Do not expose the web interface to public networks
- Always bind to localhost in production environments

#### Image Source Verification
```bash
# Verify image signatures when possible
docker trust inspect nginx:latest

# Use specific tags instead of 'latest'
# Analyze: nginx:1.21.6 instead of nginx:latest

# Be cautious with unknown or unofficial images
```

#### Sensitive Data Detection
When analyzing images, check for:
- Hardcoded passwords in environment variables
- Private keys or certificates in the filesystem
- Database credentials in configuration files
- API keys or tokens in application files

### Development Workflow

#### Pre-commit Analysis
```bash
#!/bin/bash
# .git/hooks/pre-commit

# Analyze the image before committing Dockerfile changes
if [ -f Dockerfile ]; then
  echo "Analyzing Docker image..."
  
  # Build the image
  IMAGE_TAG="myapp:$(git rev-parse --short HEAD)"
  docker build -t "$IMAGE_TAG" .
  
  # Analyze with dockerutils
  dockerutils server --no-browser &
  SERVER_PID=$!
  sleep 5
  
  # Quick size check
  RESULT=$(curl -s -X POST http://localhost:8080/api/analyze \
    -H "Content-Type: application/json" \
    -d "{\"image_name\": \"$IMAGE_TAG\"}")
  
  SIZE=$(echo "$RESULT" | jq -r '.size_bytes')
  if [ "$SIZE" -gt 1000000000 ]; then  # 1GB limit
    echo "❌ Image too large: ${SIZE} bytes"
    kill $SERVER_PID
    exit 1
  fi
  
  echo "✅ Image analysis passed"
  kill $SERVER_PID
fi
```

#### Continuous Integration
```yaml
# .github/workflows/docker-analysis.yml
name: Docker Image Analysis

on:
  push:
    paths:
      - 'Dockerfile'
      - '.dockerignore'

jobs:
  analyze:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    
    - name: Build Docker image
      run: docker build -t test-image .
      
    - name: Run Docker Utils analysis
      run: |
        # Setup dockerutils
        wget https://github.com/smiller333/dockerutils/releases/latest/download/dockerutils-linux-amd64
        chmod +x dockerutils-linux-amd64
        
        # Start server
        ./dockerutils-linux-amd64 server --no-browser &
        sleep 10
        
        # Analyze image
        curl -X POST http://localhost:8080/api/analyze \
          -H "Content-Type: application/json" \
          -d '{"image_name": "test-image"}' > analysis.json
        
        # Check results
        if jq -r '.success' analysis.json | grep -q false; then
          echo "Analysis failed"
          exit 1
        fi
        
        # Size check (example: fail if >500MB)
        SIZE=$(jq -r '.size_bytes' analysis.json)
        if [ "$SIZE" -gt 500000000 ]; then
          echo "Image too large: $SIZE bytes"
          exit 1
        fi
```

### Monitoring and Alerting

#### Health Monitoring
```bash
#!/bin/bash
# health-check.sh

HEALTH_URL="http://localhost:8080/api/health"

while true; do
  if ! curl -f -s "$HEALTH_URL" > /dev/null; then
    echo "$(date): Docker Utils health check failed"
    # Send alert or restart service
  fi
  sleep 30
done
```

#### Analysis Metrics
```python
#!/usr/bin/env python3
# metrics-collector.py

import requests
import json
import time
from datetime import datetime

def collect_metrics():
    api_base = "http://localhost:8080/api"
    
    # Get all summaries
    summaries = requests.get(f"{api_base}/summaries").json()
    
    metrics = {
        "timestamp": datetime.now().isoformat(),
        "total_analyses": len(summaries),
        "total_size": sum(s["image_size"] for s in summaries),
        "architectures": list(set(s["architecture"] for s in summaries)),
        "avg_size": sum(s["image_size"] for s in summaries) / len(summaries) if summaries else 0
    }
    
    print(json.dumps(metrics, indent=2))

if __name__ == "__main__":
    collect_metrics()
```

---

This completes the comprehensive user guide for Docker Utils. For technical API details, see the [API Documentation](API.md). For contributing to the project, see the [Contributing Guide](CONTRIBUTING.md).

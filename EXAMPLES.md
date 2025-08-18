# Docker Utils Examples

Comprehensive hands-on examples and tutorials for using Docker Utils effectively. These examples range from basic usage to advanced automation scenarios.

## Table of Contents

1. [Quick Start Examples](#quick-start-examples)
2. [Web Interface Examples](#web-interface-examples)
3. [API Usage Examples](#api-usage-examples)
4. [Common Use Cases](#common-use-cases)
5. [Automation Examples](#automation-examples)
6. [Advanced Scenarios](#advanced-scenarios)
7. [Integration Examples](#integration-examples)

## Quick Start Examples

### Example 1: First Time Setup

```bash
# Step 1: Clone and build (or download binary)
git clone https://github.com/your-org/dockerutils.git
cd dockerutils
make build-dev

# Step 2: Start the web interface
./dockerutils tools --port 8080

# Step 3: Open browser and navigate to http://localhost:8080

# Step 4: Analyze your first image
# In the web interface, enter: alpine:latest
# Click "Analyze Image" and explore the results
```

### Example 2: Using Docker Container

```bash
# Quick start with Docker
./scripts/docker-run.sh run-persistent

# The web interface will be available at http://localhost:8080
# Analysis data will persist in the ./data directory
```

### Example 3: Command Line Analysis (Future Feature)

```bash
# Quick analysis via command line
dockerutils analyze nginx:latest --output json

# Save analysis to specific directory
dockerutils analyze ubuntu:20.04 --output-dir ./my-analysis

# Multiple images at once
dockerutils analyze alpine:latest nginx:latest node:18-alpine
```

## Web Interface Examples

### Example 1: Analyzing a Simple Image

**Scenario**: Analyze a lightweight Alpine Linux image

**Steps**:
1. Start dockerutils: `dockerutils tools`
2. Open http://localhost:8080
3. Enter image name: `alpine:latest`
4. Click "Analyze Image"
5. Explore the results:
   - **Layers tab**: See the 1-2 layers that make up Alpine
   - **Filesystem tab**: Browse the minimal filesystem
   - **Metadata tab**: View image configuration

**Expected Results**:
- Small image size (~5MB)
- Minimal layer count (1-2 layers)
- Basic Alpine Linux filesystem structure
- Simple configuration with basic shell

### Example 2: Analyzing a Complex Application Image

**Scenario**: Analyze an NGINX web server image

**Steps**:
1. Enter image name: `nginx:latest`
2. Enable all analysis options:
   - ✅ Include Layers
   - ✅ Include Filesystem
   - ⚠️ Security Scan (optional)
3. Start analysis
4. Review results:
   - **Layers**: Multiple layers showing OS base, package installation, NGINX setup
   - **Filesystem**: Complete web server filesystem
   - **Configuration**: Exposed ports (80, 443), startup commands

**Expected Results**:
- Larger image size (~100MB+)
- Multiple layers (5-10 layers)
- Web server configuration files
- NGINX-specific directory structure

### Example 3: Comparing Images

**Scenario**: Compare different versions of the same image

**Steps**:
1. Analyze `nginx:1.20-alpine`
2. Then analyze `nginx:1.21-alpine`
3. Use the interface to switch between analyses
4. Compare:
   - Layer differences
   - Size differences
   - Configuration changes
   - Filesystem structure

**Key Observations**:
- Version differences in installed packages
- Security updates between versions
- Changes in default configuration

## API Usage Examples

### Example 1: Basic API Health Check

```bash
# Check if the server is running
curl http://localhost:8080/api/health

# Expected response:
{
  "status": "healthy",
  "timestamp": "2025-01-21T15:30:00Z",
  "version": "v1.0.0",
  "git_commit": "abc123",
  "build_time": "2025-01-21T14:00:00Z",
  "go_version": "go1.24.2"
}
```

### Example 2: List All Analyzed Images

```bash
# Get list of all analyzed images
curl http://localhost:8080/api/summaries

# Example response:
[
  {
    "image_id": "sha256:7aab056cecc6",
    "image_tag": "alpine:latest",
    "image_source": "docker.io",
    "image_size": 7738912,
    "architecture": "amd64",
    "analyzed_at": "2025-01-21T15:30:00Z",
    "status": "completed",
    "request_id": "7aab056cecc6"
  }
]
```

### Example 3: Analyze Image via API

```bash
# Synchronous analysis (blocks until complete)
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "image_name": "hello-world:latest",
    "keep_temp_files": true,
    "force_pull": false
  }'

# Expected response:
{
  "success": true,
  "message": "Analysis completed successfully",
  "image_id": "sha256:feb5d9fea6a5",
  "request_id": "feb5d9fea6a5",
  "analysis_time": "2.3s"
}
```

### Example 4: Asynchronous Analysis

```bash
# Start analysis without waiting
curl -X POST http://localhost:8080/api/analyze-async \
  -H "Content-Type: application/json" \
  -d '{
    "image_name": "nginx:latest"
  }'

# Response:
{
  "success": true,
  "message": "Analysis started",
  "request_id": "abc123def456",
  "estimated_time": "30s"
}

# Check status later
curl http://localhost:8080/api/summaries | jq '.[] | select(.request_id == "abc123def456")'
```

### Example 5: Get Detailed Analysis Results

```bash
# Get full analysis data for a specific image
curl http://localhost:8080/api/info/7aab056cecc6 > alpine-analysis.json

# Pretty print the JSON
curl http://localhost:8080/api/info/7aab056cecc6 | jq . > alpine-analysis-formatted.json

# Extract specific information
curl http://localhost:8080/api/info/7aab056cecc6 | jq '.layers[] | {id: .id, size: .size, command: .command}'
```

### Example 6: Delete Analysis Results

```bash
# Remove analysis results
curl -X DELETE http://localhost:8080/api/info/7aab056cecc6

# Response:
{
  "success": true,
  "message": "Info 7aab056cecc6 deleted successfully"
}
```

## Common Use Cases

### Use Case 1: Security Audit

**Scenario**: Audit a production image for security issues

```bash
# 1. Start dockerutils
dockerutils tools

# 2. Analyze the production image
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "image_name": "myapp:production",
    "security_scan": true
  }'

# 3. Get security information
curl http://localhost:8080/api/info/$(curl -s http://localhost:8080/api/summaries | jq -r '.[0].request_id') | \
  jq '.security // "No security data available"'

# 4. Check for common security issues
curl http://localhost:8080/api/info/$(curl -s http://localhost:8080/api/summaries | jq -r '.[0].request_id') | \
  jq '.layers[] | select(.command | contains("root")) | {layer: .id, command: .command}'
```

**What to Look For**:
- Processes running as root
- Exposed secrets in environment variables
- Outdated packages
- Unnecessary network ports
- World-writable files

### Use Case 2: Image Size Optimization

**Scenario**: Find opportunities to reduce image size

```bash
# 1. Analyze your image
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{"image_name": "myapp:latest"}'

# 2. Get layer size information
IMAGE_ID=$(curl -s http://localhost:8080/api/summaries | jq -r '.[0].request_id')
curl http://localhost:8080/api/info/$IMAGE_ID | \
  jq '.layers[] | {size: .size, command: .command}' | \
  jq -s 'sort_by(.size) | reverse'

# 3. Look for large layers
curl http://localhost:8080/api/info/$IMAGE_ID | \
  jq '.layers[] | select(.size > 100000000) | {size_mb: (.size / 1024 / 1024), command: .command}'

# 4. Check for duplicate files
curl http://localhost:8080/api/info/$IMAGE_ID | \
  jq '.filesystem.files | group_by(.name) | map(select(length > 1))'
```

**Optimization Strategies**:
- Combine RUN commands to reduce layers
- Remove package managers and caches
- Use multi-stage builds
- Choose smaller base images

### Use Case 3: Debugging Application Issues

**Scenario**: Debug why an application isn't working in a container

```bash
# 1. Analyze the problematic image
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{"image_name": "myapp:broken"}'

# 2. Check environment variables
IMAGE_ID=$(curl -s http://localhost:8080/api/summaries | jq -r '.[0].request_id')
curl http://localhost:8080/api/info/$IMAGE_ID | jq '.config.env'

# 3. Check executable permissions
curl http://localhost:8080/api/info/$IMAGE_ID | \
  jq '.filesystem.files[] | select(.name | endswith("myapp")) | {name: .name, permissions: .permissions}'

# 4. Check working directory and entrypoint
curl http://localhost:8080/api/info/$IMAGE_ID | jq '.config | {workdir: .workdir, entrypoint: .entrypoint, cmd: .cmd}'

# 5. Look for missing dependencies
curl http://localhost:8080/api/info/$IMAGE_ID | \
  jq '.filesystem.files[] | select(.name | contains("lib")) | .name'
```

### Use Case 4: Compliance Checking

**Scenario**: Ensure images meet organizational standards

```bash
# Check script to validate image compliance
#!/bin/bash
IMAGE_NAME="$1"
IMAGE_ID=""

# Analyze image
echo "Analyzing $IMAGE_NAME..."
RESPONSE=$(curl -s -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d "{\"image_name\": \"$IMAGE_NAME\"}")

# Get image ID
IMAGE_ID=$(echo $RESPONSE | jq -r '.request_id')

# Wait for analysis to complete
sleep 5

echo "Running compliance checks..."

# Check 1: No root user
ROOT_PROCESSES=$(curl -s http://localhost:8080/api/info/$IMAGE_ID | \
  jq '.config.user // "root"')
if [ "$ROOT_PROCESSES" = "root" ]; then
  echo "❌ FAIL: Container runs as root user"
else
  echo "✅ PASS: Container runs as non-root user"
fi

# Check 2: No secrets in environment
SECRETS=$(curl -s http://localhost:8080/api/info/$IMAGE_ID | \
  jq '.config.env[]? | select(. | ascii_downcase | contains("password") or contains("secret") or contains("key"))')
if [ -n "$SECRETS" ]; then
  echo "❌ FAIL: Potential secrets found in environment variables"
  echo "$SECRETS"
else
  echo "✅ PASS: No obvious secrets in environment"
fi

# Check 3: Image size under limit (100MB)
SIZE=$(curl -s http://localhost:8080/api/info/$IMAGE_ID | jq '.image_size')
if [ "$SIZE" -gt 104857600 ]; then
  echo "❌ FAIL: Image size $(($SIZE / 1024 / 1024))MB exceeds 100MB limit"
else
  echo "✅ PASS: Image size $(($SIZE / 1024 / 1024))MB is within limits"
fi

echo "Compliance check complete for $IMAGE_NAME"
```

## Automation Examples

### Example 1: Automated Image Analysis Pipeline

```bash
#!/bin/bash
# analyze-images.sh - Batch analyze multiple images

IMAGES=(
  "alpine:latest"
  "nginx:latest"
  "node:18-alpine"
  "python:3.11-slim"
)

# Start dockerutils if not running
if ! curl -s http://localhost:8080/api/health >/dev/null 2>&1; then
  echo "Starting dockerutils..."
  dockerutils tools --port 8080 &
  sleep 5
fi

echo "Starting batch analysis of ${#IMAGES[@]} images..."

for IMAGE in "${IMAGES[@]}"; do
  echo "Analyzing $IMAGE..."
  
  # Start async analysis
  RESPONSE=$(curl -s -X POST http://localhost:8080/api/analyze-async \
    -H "Content-Type: application/json" \
    -d "{\"image_name\": \"$IMAGE\"}")
  
  REQUEST_ID=$(echo $RESPONSE | jq -r '.request_id')
  echo "  Request ID: $REQUEST_ID"
  
  # Optional: wait for completion
  while true; do
    STATUS=$(curl -s http://localhost:8080/api/summaries | \
      jq -r ".[] | select(.request_id == \"$REQUEST_ID\") | .status")
    
    if [ "$STATUS" = "completed" ]; then
      echo "  ✅ Analysis complete"
      break
    elif [ "$STATUS" = "failed" ]; then
      echo "  ❌ Analysis failed"
      break
    else
      echo "  ⏳ Status: $STATUS"
      sleep 10
    fi
  done
done

echo "Batch analysis complete. View results at http://localhost:8080"
```

### Example 2: Continuous Integration Integration

```yaml
# .github/workflows/docker-analysis.yml
name: Docker Image Analysis

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  analyze:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Build Docker image
      run: docker build -t myapp:${{ github.sha }} .
    
    - name: Download dockerutils
      run: |
        curl -L https://github.com/your-org/dockerutils/releases/latest/download/dockerutils-linux-amd64 -o dockerutils
        chmod +x dockerutils
    
    - name: Start dockerutils
      run: |
        ./dockerutils tools --port 8080 &
        sleep 10
    
    - name: Analyze image
      run: |
        curl -X POST http://localhost:8080/api/analyze \
          -H "Content-Type: application/json" \
          -d '{"image_name": "myapp:${{ github.sha }}"}'
    
    - name: Get analysis results
      run: |
        IMAGE_ID=$(curl -s http://localhost:8080/api/summaries | jq -r '.[0].request_id')
        curl http://localhost:8080/api/info/$IMAGE_ID > analysis-results.json
    
    - name: Check image size
      run: |
        SIZE=$(jq '.image_size' analysis-results.json)
        if [ "$SIZE" -gt 104857600 ]; then
          echo "❌ Image size $(($SIZE / 1024 / 1024))MB exceeds 100MB limit"
          exit 1
        fi
        echo "✅ Image size $(($SIZE / 1024 / 1024))MB is acceptable"
    
    - name: Upload analysis results
      uses: actions/upload-artifact@v3
      with:
        name: docker-analysis-results
        path: analysis-results.json
```

### Example 3: Monitoring Script

```bash
#!/bin/bash
# monitor-images.sh - Monitor image updates and re-analyze

WATCH_IMAGES=(
  "nginx:latest"
  "alpine:latest"
  "node:18-alpine"
)

LAST_ANALYSIS_FILE="./last-analysis.txt"

# Function to get image digest
get_image_digest() {
  docker inspect "$1" --format='{{index .RepoDigests 0}}' 2>/dev/null || echo "not-found"
}

# Function to analyze image
analyze_image() {
  local image="$1"
  echo "Analyzing updated image: $image"
  
  curl -X POST http://localhost:8080/api/analyze \
    -H "Content-Type: application/json" \
    -d "{\"image_name\": \"$image\"}" \
    -s | jq '.message'
}

# Load last analysis data
if [ -f "$LAST_ANALYSIS_FILE" ]; then
  source "$LAST_ANALYSIS_FILE"
fi

echo "Checking for image updates..."

for IMAGE in "${IMAGES[@]}"; do
  # Pull latest version
  docker pull "$IMAGE" >/dev/null 2>&1
  
  # Get current digest
  CURRENT_DIGEST=$(get_image_digest "$IMAGE")
  
  # Get stored digest
  LAST_VAR="LAST_DIGEST_$(echo "$IMAGE" | tr ':/-' '___')"
  LAST_DIGEST="${!LAST_VAR}"
  
  if [ "$CURRENT_DIGEST" != "$LAST_DIGEST" ]; then
    echo "🔄 Update detected for $IMAGE"
    echo "  Previous: $LAST_DIGEST"
    echo "  Current:  $CURRENT_DIGEST"
    
    # Re-analyze the image
    analyze_image "$IMAGE"
    
    # Update stored digest
    echo "$LAST_VAR=\"$CURRENT_DIGEST\"" >> "$LAST_ANALYSIS_FILE.new"
  else
    echo "✅ No update for $IMAGE"
    echo "$LAST_VAR=\"$CURRENT_DIGEST\"" >> "$LAST_ANALYSIS_FILE.new"
  fi
done

# Replace old analysis file
mv "$LAST_ANALYSIS_FILE.new" "$LAST_ANALYSIS_FILE"

echo "Image monitoring complete"
```

## Advanced Scenarios

### Scenario 1: Multi-Registry Analysis

```bash
#!/bin/bash
# analyze-multi-registry.sh

REGISTRIES=(
  "docker.io/library/alpine:latest"
  "quay.io/prometheus/prometheus:latest"
  "gcr.io/distroless/java:11"
  "registry.redhat.io/ubi8/ubi:latest"
)

echo "Analyzing images from multiple registries..."

for FULL_IMAGE in "${REGISTRIES[@]}"; do
  echo "Processing: $FULL_IMAGE"
  
  # Extract registry and image
  REGISTRY=$(echo "$FULL_IMAGE" | cut -d'/' -f1)
  IMAGE=$(echo "$FULL_IMAGE" | cut -d'/' -f2-)
  
  echo "  Registry: $REGISTRY"
  echo "  Image: $IMAGE"
  
  # Login might be required for some registries
  case "$REGISTRY" in
    "registry.redhat.io")
      echo "  Note: Red Hat registry may require authentication"
      ;;
    "gcr.io")
      echo "  Note: GCR may require gcloud authentication"
      ;;
  esac
  
  # Analyze the image
  curl -X POST http://localhost:8080/api/analyze-async \
    -H "Content-Type: application/json" \
    -d "{\"image_name\": \"$FULL_IMAGE\"}" \
    -s | jq '.message'
  
  echo "  Analysis started for $FULL_IMAGE"
  echo ""
done

echo "All analyses started. Check http://localhost:8080 for results."
```

### Scenario 2: Image Vulnerability Scanning

```bash
#!/bin/bash
# security-scan.sh - Enhanced security analysis

IMAGE_NAME="$1"

if [ -z "$IMAGE_NAME" ]; then
  echo "Usage: $0 <image-name>"
  exit 1
fi

echo "Security analysis for: $IMAGE_NAME"

# Analyze image
echo "Starting analysis..."
RESPONSE=$(curl -s -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d "{\"image_name\": \"$IMAGE_NAME\"}")

IMAGE_ID=$(echo $RESPONSE | jq -r '.request_id')

# Wait for completion
sleep 10

echo "Generating security report..."

# Get analysis data
ANALYSIS=$(curl -s http://localhost:8080/api/info/$IMAGE_ID)

echo "=== SECURITY ANALYSIS REPORT ==="
echo "Image: $IMAGE_NAME"
echo "Analysis Date: $(date)"
echo "Image ID: $IMAGE_ID"
echo ""

# Check 1: User configuration
echo "1. USER CONFIGURATION:"
USER=$(echo "$ANALYSIS" | jq -r '.config.user // "root"')
if [ "$USER" = "root" ]; then
  echo "   ❌ HIGH: Container runs as root user"
  echo "   Recommendation: Add USER directive to Dockerfile"
else
  echo "   ✅ GOOD: Container runs as user: $USER"
fi
echo ""

# Check 2: Environment variables
echo "2. ENVIRONMENT VARIABLES:"
SECRETS=$(echo "$ANALYSIS" | jq -r '.config.env[]? | select(. | ascii_downcase | test("password|secret|key|token"))')
if [ -n "$SECRETS" ]; then
  echo "   ❌ HIGH: Potential secrets in environment:"
  echo "$SECRETS" | sed 's/^/     /'
else
  echo "   ✅ GOOD: No obvious secrets in environment"
fi
echo ""

# Check 3: Network ports
echo "3. EXPOSED PORTS:"
PORTS=$(echo "$ANALYSIS" | jq -r '.config.exposed_ports[]? // empty')
if [ -n "$PORTS" ]; then
  echo "   ℹ️  INFO: Exposed ports:"
  echo "$PORTS" | sed 's/^/     /'
  
  # Check for common insecure ports
  if echo "$PORTS" | grep -q "22/tcp"; then
    echo "   ⚠️  WARNING: SSH port 22 exposed"
  fi
  if echo "$PORTS" | grep -q "23/tcp"; then
    echo "   ❌ HIGH: Telnet port 23 exposed (insecure)"
  fi
else
  echo "   ✅ GOOD: No exposed ports"
fi
echo ""

# Check 4: File permissions
echo "4. FILE PERMISSIONS:"
WORLD_WRITABLE=$(echo "$ANALYSIS" | jq -r '.filesystem.files[]? | select(.permissions | test(".*w.$")) | .name' | head -5)
if [ -n "$WORLD_WRITABLE" ]; then
  echo "   ⚠️  WARNING: World-writable files found:"
  echo "$WORLD_WRITABLE" | sed 's/^/     /'
else
  echo "   ✅ GOOD: No world-writable files in sample"
fi
echo ""

# Check 5: Image size
echo "5. IMAGE SIZE:"
SIZE=$(echo "$ANALYSIS" | jq -r '.image_size')
SIZE_MB=$((SIZE / 1024 / 1024))
if [ "$SIZE" -gt 1073741824 ]; then  # 1GB
  echo "   ⚠️  WARNING: Large image size: ${SIZE_MB}MB"
  echo "   Recommendation: Consider using smaller base images"
elif [ "$SIZE" -gt 524288000 ]; then  # 500MB
  echo "   ℹ️  INFO: Medium image size: ${SIZE_MB}MB"
else
  echo "   ✅ GOOD: Compact image size: ${SIZE_MB}MB"
fi
echo ""

# Check 6: Base image
echo "6. BASE IMAGE ANALYSIS:"
LAYERS=$(echo "$ANALYSIS" | jq -r '.layers | length')
echo "   Layer count: $LAYERS"

FIRST_LAYER=$(echo "$ANALYSIS" | jq -r '.layers[0].command // "unknown"')
if echo "$FIRST_LAYER" | grep -qi "alpine"; then
  echo "   ✅ GOOD: Uses Alpine Linux (security-focused)"
elif echo "$FIRST_LAYER" | grep -qi "distroless"; then
  echo "   ✅ EXCELLENT: Uses distroless image"
elif echo "$FIRST_LAYER" | grep -qi "ubuntu\|debian"; then
  echo "   ℹ️  INFO: Uses Ubuntu/Debian base"
else
  echo "   ℹ️  INFO: Base image: $FIRST_LAYER"
fi
echo ""

echo "=== END SECURITY REPORT ==="
echo ""
echo "For detailed results, visit: http://localhost:8080"
```

### Scenario 3: Image Comparison Tool

```bash
#!/bin/bash
# compare-images.sh - Compare two Docker images

IMAGE1="$1"
IMAGE2="$2"

if [ -z "$IMAGE1" ] || [ -z "$IMAGE2" ]; then
  echo "Usage: $0 <image1> <image2>"
  echo "Example: $0 nginx:1.20-alpine nginx:1.21-alpine"
  exit 1
fi

echo "Comparing images: $IMAGE1 vs $IMAGE2"

# Analyze both images
echo "Analyzing $IMAGE1..."
RESPONSE1=$(curl -s -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d "{\"image_name\": \"$IMAGE1\"}")
ID1=$(echo $RESPONSE1 | jq -r '.request_id')

echo "Analyzing $IMAGE2..."
RESPONSE2=$(curl -s -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d "{\"image_name\": \"$IMAGE2\"}")
ID2=$(echo $RESPONSE2 | jq -r '.request_id')

# Wait for analyses to complete
echo "Waiting for analyses to complete..."
sleep 15

# Get analysis data
ANALYSIS1=$(curl -s http://localhost:8080/api/info/$ID1)
ANALYSIS2=$(curl -s http://localhost:8080/api/info/$ID2)

echo ""
echo "=== IMAGE COMPARISON REPORT ==="
echo "Image 1: $IMAGE1"
echo "Image 2: $IMAGE2"
echo ""

# Size comparison
SIZE1=$(echo "$ANALYSIS1" | jq -r '.image_size')
SIZE2=$(echo "$ANALYSIS2" | jq -r '.image_size')
SIZE1_MB=$((SIZE1 / 1024 / 1024))
SIZE2_MB=$((SIZE2 / 1024 / 1024))
SIZE_DIFF=$((SIZE2 - SIZE1))
SIZE_DIFF_MB=$((SIZE_DIFF / 1024 / 1024))

echo "SIZE COMPARISON:"
echo "  $IMAGE1: ${SIZE1_MB}MB"
echo "  $IMAGE2: ${SIZE2_MB}MB"
if [ $SIZE_DIFF -gt 0 ]; then
  echo "  Difference: +${SIZE_DIFF_MB}MB (${IMAGE2} is larger)"
elif [ $SIZE_DIFF -lt 0 ]; then
  echo "  Difference: ${SIZE_DIFF_MB}MB (${IMAGE2} is smaller)"
else
  echo "  Difference: Same size"
fi
echo ""

# Layer comparison
LAYERS1=$(echo "$ANALYSIS1" | jq -r '.layers | length')
LAYERS2=$(echo "$ANALYSIS2" | jq -r '.layers | length')

echo "LAYER COMPARISON:"
echo "  $IMAGE1: $LAYERS1 layers"
echo "  $IMAGE2: $LAYERS2 layers"
echo ""

# Environment comparison
echo "ENVIRONMENT DIFFERENCES:"
ENV1=$(echo "$ANALYSIS1" | jq -r '.config.env[]?' | sort)
ENV2=$(echo "$ANALYSIS2" | jq -r '.config.env[]?' | sort)

echo "$ENV1" > /tmp/env1.txt
echo "$ENV2" > /tmp/env2.txt

if ! diff -q /tmp/env1.txt /tmp/env2.txt >/dev/null; then
  echo "  Environment variables differ:"
  diff -u /tmp/env1.txt /tmp/env2.txt | grep "^[+-]" | sed 's/^+/  Added:   /' | sed 's/^-/  Removed: /'
else
  echo "  Environment variables are identical"
fi
echo ""

# Port comparison
echo "PORT DIFFERENCES:"
PORTS1=$(echo "$ANALYSIS1" | jq -r '.config.exposed_ports[]?' | sort)
PORTS2=$(echo "$ANALYSIS2" | jq -r '.config.exposed_ports[]?' | sort)

echo "$PORTS1" > /tmp/ports1.txt
echo "$PORTS2" > /tmp/ports2.txt

if ! diff -q /tmp/ports1.txt /tmp/ports2.txt >/dev/null; then
  echo "  Exposed ports differ:"
  diff -u /tmp/ports1.txt /tmp/ports2.txt | grep "^[+-]" | sed 's/^+/  Added:   /' | sed 's/^-/  Removed: /'
else
  echo "  Exposed ports are identical"
fi
echo ""

# Cleanup
rm -f /tmp/env1.txt /tmp/env2.txt /tmp/ports1.txt /tmp/ports2.txt

echo "=== END COMPARISON REPORT ==="
echo ""
echo "For detailed analysis, visit: http://localhost:8080"
```

## Integration Examples

### Example 1: Jenkins Pipeline Integration

```groovy
// Jenkinsfile
pipeline {
    agent any
    
    environment {
        DOCKERUTILS_URL = 'http://dockerutils:8080'
        IMAGE_NAME = "${env.JOB_NAME}:${env.BUILD_NUMBER}"
    }
    
    stages {
        stage('Build') {
            steps {
                script {
                    docker.build("${IMAGE_NAME}")
                }
            }
        }
        
        stage('Analyze') {
            steps {
                script {
                    // Start analysis
                    def response = sh(
                        script: """
                            curl -X POST ${DOCKERUTILS_URL}/api/analyze \
                              -H "Content-Type: application/json" \
                              -d '{"image_name": "${IMAGE_NAME}"}' \
                              -s
                        """,
                        returnStdout: true
                    )
                    
                    def result = readJSON text: response
                    env.IMAGE_ID = result.request_id
                    
                    echo "Analysis started with ID: ${env.IMAGE_ID}"
                }
            }
        }
        
        stage('Security Check') {
            steps {
                script {
                    // Get analysis results
                    def analysis = sh(
                        script: "curl -s ${DOCKERUTILS_URL}/api/info/${env.IMAGE_ID}",
                        returnStdout: true
                    )
                    
                    def result = readJSON text: analysis
                    
                    // Check image size
                    def sizeLimit = 104857600 // 100MB
                    if (result.image_size > sizeLimit) {
                        error("Image size ${result.image_size / 1024 / 1024}MB exceeds limit")
                    }
                    
                    // Check for root user
                    if (result.config.user == null || result.config.user == "root") {
                        error("Image runs as root user - security violation")
                    }
                    
                    echo "Security checks passed"
                }
            }
        }
        
        stage('Generate Report') {
            steps {
                script {
                    sh """
                        curl -s ${DOCKERUTILS_URL}/api/info/${env.IMAGE_ID} > analysis-report.json
                    """
                    
                    archiveArtifacts artifacts: 'analysis-report.json'
                }
            }
        }
    }
    
    post {
        always {
            // Clean up analysis data
            script {
                sh """
                    curl -X DELETE ${DOCKERUTILS_URL}/api/info/${env.IMAGE_ID} || true
                """
            }
        }
    }
}
```

### Example 2: GitLab CI Integration

```yaml
# .gitlab-ci.yml
stages:
  - build
  - analyze
  - security
  - deploy

variables:
  DOCKERUTILS_URL: "http://dockerutils:8080"
  IMAGE_NAME: "$CI_REGISTRY_IMAGE:$CI_COMMIT_SHA"

build:
  stage: build
  script:
    - docker build -t $IMAGE_NAME .
    - docker push $IMAGE_NAME

analyze:
  stage: analyze
  script:
    - |
      RESPONSE=$(curl -X POST $DOCKERUTILS_URL/api/analyze \
        -H "Content-Type: application/json" \
        -d "{\"image_name\": \"$IMAGE_NAME\"}" \
        -s)
      export IMAGE_ID=$(echo $RESPONSE | jq -r '.request_id')
      echo "IMAGE_ID=$IMAGE_ID" >> analyze.env
  artifacts:
    reports:
      dotenv: analyze.env

security:
  stage: security
  dependencies:
    - analyze
  script:
    - |
      # Wait for analysis to complete
      sleep 30
      
      # Get analysis results
      curl -s $DOCKERUTILS_URL/api/info/$IMAGE_ID > analysis.json
      
      # Security checks
      SIZE=$(jq '.image_size' analysis.json)
      USER=$(jq -r '.config.user // "root"' analysis.json)
      
      if [ $SIZE -gt 104857600 ]; then
        echo "❌ Image size exceeds 100MB limit"
        exit 1
      fi
      
      if [ "$USER" = "root" ]; then
        echo "❌ Image runs as root user"
        exit 1
      fi
      
      echo "✅ Security checks passed"
  artifacts:
    reports:
      junit: analysis.json
    paths:
      - analysis.json

deploy:
  stage: deploy
  dependencies:
    - security
  script:
    - kubectl set image deployment/myapp myapp=$IMAGE_NAME
  only:
    - main
```

### Example 3: Terraform Integration

```hcl
# terraform/dockerutils.tf
resource "docker_container" "dockerutils" {
  name  = "dockerutils"
  image = "dockerutils:latest"
  
  ports {
    internal = 8080
    external = 8080
  }
  
  volumes {
    host_path      = "/var/run/docker.sock"
    container_path = "/var/run/docker.sock"
  }
  
  volumes {
    host_path      = abspath("./data")
    container_path = "/app/data"
  }
  
  env = [
    "DOCKERUTILS_HOST=0.0.0.0",
    "DOCKERUTILS_PORT=8080"
  ]
  
  restart = "unless-stopped"
}

resource "null_resource" "wait_for_dockerutils" {
  depends_on = [docker_container.dockerutils]
  
  provisioner "local-exec" {
    command = <<-EOT
      echo "Waiting for dockerutils to start..."
      for i in {1..30}; do
        if curl -s http://localhost:8080/api/health >/dev/null; then
          echo "Dockerutils is ready"
          break
        fi
        echo "Attempt $i/30: Waiting..."
        sleep 10
      done
    EOT
  }
}

output "dockerutils_url" {
  value = "http://localhost:8080"
}
```

### Example 4: Kubernetes CronJob for Regular Analysis

```yaml
# k8s/scheduled-analysis.yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: docker-image-analysis
spec:
  schedule: "0 2 * * *"  # Daily at 2 AM
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: analyzer
            image: curlimages/curl:latest
            command:
            - /bin/sh
            - -c
            - |
              # List of images to analyze
              IMAGES="nginx:latest alpine:latest node:18-alpine python:3.11-slim"
              
              # Dockerutils API endpoint
              API_URL="http://dockerutils-service:8080/api"
              
              echo "Starting scheduled analysis..."
              
              for IMAGE in $IMAGES; do
                echo "Analyzing $IMAGE..."
                
                # Start async analysis
                RESPONSE=$(curl -X POST $API_URL/analyze-async \
                  -H "Content-Type: application/json" \
                  -d "{\"image_name\": \"$IMAGE\"}" \
                  -s)
                
                REQUEST_ID=$(echo $RESPONSE | jq -r '.request_id')
                echo "Started analysis: $REQUEST_ID"
              done
              
              echo "Scheduled analysis jobs started"
          restartPolicy: OnFailure
```

## Conclusion

These examples demonstrate the versatility and power of Docker Utils across various scenarios:

- **Development**: Quick image inspection and debugging
- **Security**: Comprehensive security auditing and compliance checking
- **Operations**: Automated monitoring and analysis pipelines
- **CI/CD**: Integration with build and deployment workflows

### Next Steps

1. **Start Simple**: Begin with the quick start examples
2. **Explore Web Interface**: Use the browser interface for interactive exploration
3. **Automate**: Implement automation scripts for your specific use cases
4. **Integrate**: Add Docker Utils to your existing CI/CD pipelines
5. **Customize**: Adapt the examples to your organization's needs

### Additional Resources

- [User Guide](docs/USER_GUIDE.md) - Comprehensive usage documentation
- [API Reference](docs/API.md) - Complete API documentation
- [Troubleshooting](TROUBLESHOOTING.md) - Common issues and solutions
- [Installation Guide](INSTALLATION.md) - Setup instructions

**Happy analyzing!** 🐳

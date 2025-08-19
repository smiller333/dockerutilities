# Installation Guide

Comprehensive installation instructions for Docker Utils across different platforms and use cases.

## Quick Install (Recommended)

### Using Docker (Easiest)

```bash
# Clone repository
git clone https://github.com/your-org/dockerutils.git
cd dockerutils

# Run with Docker
./scripts/docker-run.sh run-persistent

# Access at http://localhost:8080
```

### Using Pre-built Binaries

```bash
# Download latest release
curl -L https://github.com/your-org/dockerutils/releases/latest/download/dockerutils-linux-amd64 -o dockerutils

# Make executable
chmod +x dockerutils

# Move to PATH
sudo mv dockerutils /usr/local/bin/

# Verify installation
dockerutils version
```

## Platform-Specific Installation

### Linux

#### Ubuntu/Debian
```bash
# Install dependencies
sudo apt update
sudo apt install docker.io git

# Add user to docker group
sudo usermod -aG docker $USER
newgrp docker

# Install dockerutils
wget https://github.com/your-org/dockerutils/releases/latest/download/dockerutils-linux-amd64
chmod +x dockerutils-linux-amd64
sudo mv dockerutils-linux-amd64 /usr/local/bin/dockerutils
```

#### RHEL/CentOS/Fedora
```bash
# Install dependencies
sudo dnf install docker git
# or for older versions: sudo yum install docker git

# Start Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Add user to docker group
sudo usermod -aG docker $USER
newgrp docker

# Install dockerutils
curl -L https://github.com/your-org/dockerutils/releases/latest/download/dockerutils-linux-amd64 -o dockerutils
chmod +x dockerutils
sudo mv dockerutils /usr/local/bin/
```

#### Arch Linux
```bash
# Install dependencies
sudo pacman -S docker git

# Start Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Add user to docker group
sudo usermod -aG docker $USER
newgrp docker

# Install dockerutils from AUR (if available) or manually
yay -S dockerutils
# or manually:
curl -L https://github.com/your-org/dockerutils/releases/latest/download/dockerutils-linux-amd64 -o dockerutils
chmod +x dockerutils
sudo mv dockerutils /usr/local/bin/
```

### macOS

#### Using Homebrew (Recommended)
```bash
# Install Docker Desktop
brew install --cask docker

# Start Docker Desktop
open /Applications/Docker.app

# Install dockerutils (if published to Homebrew)
brew install your-org/tap/dockerutils
```

#### Manual Installation
```bash
# Install Docker Desktop from https://docker.com

# Download dockerutils
curl -L https://github.com/your-org/dockerutils/releases/latest/download/dockerutils-darwin-amd64 -o dockerutils

# For Apple Silicon Macs
curl -L https://github.com/your-org/dockerutils/releases/latest/download/dockerutils-darwin-arm64 -o dockerutils

# Make executable and install
chmod +x dockerutils
sudo mv dockerutils /usr/local/bin/
```

#### Using MacPorts
```bash
# Install Docker Desktop first
# Then install dockerutils
sudo port install dockerutils
```

### Windows

#### Using Chocolatey
```powershell
# Install Docker Desktop
choco install docker-desktop

# Install dockerutils (if published to Chocolatey)
choco install dockerutils
```

#### Using Scoop
```powershell
# Add bucket (if exists)
scoop bucket add dockerutils https://github.com/your-org/scoop-dockerutils.git

# Install Docker Desktop
scoop install docker

# Install dockerutils
scoop install dockerutils
```

#### Manual Installation
1. Install [Docker Desktop for Windows](https://docker.com)
2. Download `dockerutils-windows-amd64.exe` from [releases](https://github.com/your-org/dockerutils/releases)
3. Rename to `dockerutils.exe` and place in a directory in your PATH
4. Open PowerShell/Command Prompt and verify: `dockerutils version`

#### Using Windows Package Manager (winget)
```powershell
# Install Docker Desktop
winget install Docker.DockerDesktop

# Install dockerutils (if published)
winget install YourOrg.DockerUtils
```

## Building from Source

### Prerequisites

- Go 1.24.2 or later
- Git
- Docker Engine
- Make (optional but recommended)

### Build Steps

```bash
# Clone repository
git clone https://github.com/your-org/dockerutils.git
cd dockerutils

# Install dependencies
go mod download

# Build development version
make build-dev

# Or build release version
make build-release

# Install to GOPATH/bin
make install

# Verify installation
dockerutils version
```

### Custom Build Options

```bash
# Build with specific version
make VERSION=v1.0.0 build-release

# Build for specific platform
GOOS=linux GOARCH=amd64 make build-release

# Build for multiple platforms
make build-all

# Build with custom flags
go build -ldflags "-X main.version=custom" -o dockerutils main.go
```

### Development Build

```bash
# Quick development build
go build -o dockerutils main.go

# Run tests
go test ./...

# Run with race detection
go run -race main.go server
```

## Container Deployment

### Docker Compose

Create `docker-compose.yml`:
```yaml
version: '3.8'
services:
  dockerutils:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - ./data:/app/data
    environment:
      - DOCKERUTILS_PORT=8080
      - DOCKERUTILS_HOST=0.0.0.0
    restart: unless-stopped
```

Run with:
```bash
docker-compose up -d
```

### Kubernetes

Create `kubernetes.yml`:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dockerutils
  labels:
    app: dockerutils
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
        image: your-org/dockerutils:latest
        ports:
        - containerPort: 8080
        volumeMounts:
        - name: docker-sock
          mountPath: /var/run/docker.sock
        - name: data-storage
          mountPath: /app/data
        env:
        - name: DOCKERUTILS_PORT
          value: "8080"
        - name: DOCKERUTILS_HOST
          value: "0.0.0.0"
        resources:
          requests:
            memory: "128Mi"
            cpu: "100m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /api/health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /api/health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
      volumes:
      - name: docker-sock
        hostPath:
          path: /var/run/docker.sock
          type: Socket
      - name: data-storage
        persistentVolumeClaim:
          claimName: dockerutils-data
---
apiVersion: v1
kind: Service
metadata:
  name: dockerutils-service
  labels:
    app: dockerutils
spec:
  selector:
    app: dockerutils
  ports:
  - port: 8080
    targetPort: 8080
    name: http
  type: ClusterIP
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: dockerutils-data
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
```

Deploy with:
```bash
kubectl apply -f kubernetes.yml
```

### Helm Chart

Create a basic Helm chart structure:
```bash
helm create dockerutils
# Edit values.yaml and templates as needed
helm install dockerutils ./dockerutils
```

## Verification

### Test Installation

```bash
# Check version
dockerutils version

# Test Docker connection
docker ps

# Start server
dockerutils server --port 8080

# Test API (in another terminal)
curl http://localhost:8080/api/health
```

### Complete Installation Test

```bash
# 1. Verify Docker access
docker run hello-world

# 2. Start dockerutils
dockerutils server &

# 3. Test image analysis
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{"image_name": "hello-world:latest"}'

# 4. Check results
curl http://localhost:8080/api/summaries

# 5. Clean up
pkill dockerutils
```

## Troubleshooting Installation

### Permission Issues

#### Docker Permissions
```bash
# Add user to docker group
sudo usermod -aG docker $USER
newgrp docker

# Verify group membership
groups $USER

# Test Docker access
docker ps
```

#### Binary Permissions
```bash
# Fix executable permissions
chmod +x dockerutils

# Check file permissions
ls -la dockerutils
```

#### Directory Permissions
```bash
# Create data directory with proper permissions
mkdir -p ~/.dockerutils/data
chmod 755 ~/.dockerutils/data
```

### Path Issues

#### Add to PATH (Linux/macOS)
```bash
# Check current PATH
echo $PATH

# Add to PATH (bash/zsh)
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
source ~/.bashrc

# For zsh users
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.zshrc
source ~/.zshrc

# Verify PATH includes dockerutils location
which dockerutils
```

#### Add to PATH (Windows)
```powershell
# Check current PATH
$env:PATH

# Add directory to PATH (temporary)
$env:PATH += ";C:\path\to\dockerutils"

# Add permanently via System Properties > Environment Variables
# Or using PowerShell (requires admin):
[Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";C:\path\to\dockerutils", "Machine")
```

### Docker Issues

#### Docker Daemon Not Running
```bash
# Linux - start Docker daemon
sudo systemctl start docker
sudo systemctl enable docker

# Check Docker status
sudo systemctl status docker

# macOS - start Docker Desktop
open /Applications/Docker.app

# Windows - start Docker Desktop
# Use Start menu or desktop shortcut
```

#### Docker Socket Issues
```bash
# Check Docker socket exists
ls -la /var/run/docker.sock

# Fix socket permissions (if needed)
sudo chmod 666 /var/run/docker.sock

# Alternative: use TCP connection
export DOCKER_HOST=tcp://localhost:2376
```

#### Test Docker Connection
```bash
# Basic Docker test
docker version
docker info

# Test with simple container
docker run hello-world

# Check Docker socket permissions
ls -la /var/run/docker.sock

# Test with custom socket location (macOS Docker Desktop)
DOCKER_HOST=unix://~/.docker/desktop/docker.sock docker version
```

### Network Issues

#### Port Already in Use
```bash
# Check what's using port 8080
lsof -i :8080
netstat -tlnp | grep 8080

# Use different port
dockerutils server --port 8081
```

#### Firewall Issues
```bash
# Linux - allow port through firewall
sudo ufw allow 8080

# macOS - check System Preferences > Security & Privacy
# Windows - check Windows Defender Firewall
```

### Dependency Issues

#### Missing Go (for source builds)
```bash
# Install Go
wget https://go.dev/dl/go1.24.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

#### Missing Git
```bash
# Ubuntu/Debian
sudo apt install git

# RHEL/CentOS/Fedora
sudo dnf install git

# macOS
brew install git
# or install Xcode Command Line Tools
xcode-select --install

# Windows
# Download from https://git-scm.com/
# or use package manager
choco install git
```

#### Missing Make
```bash
# Ubuntu/Debian
sudo apt install build-essential

# RHEL/CentOS/Fedora
sudo dnf groupinstall "Development Tools"

# macOS
xcode-select --install
# or
brew install make

# Windows
choco install make
```

## Uninstallation

### Remove Binary
```bash
# Remove from system PATH
sudo rm /usr/local/bin/dockerutils

# Remove from user PATH
rm ~/bin/dockerutils

# Windows
del "C:\path\to\dockerutils.exe"
```

### Clean Data
```bash
# Remove analysis data
rm -rf ~/.dockerutils
rm -rf ./tmp/dockerutils-*
rm -rf ./data

# Remove Docker images (optional)
docker rmi dockerutils:latest
docker rmi your-org/dockerutils:latest
```

### Remove Configuration
```bash
# Remove any config files
rm -rf ~/.config/dockerutils

# Remove from PATH (Linux/macOS)
# Edit ~/.bashrc or ~/.zshrc and remove dockerutils entries
```

### Complete Cleanup
```bash
# Remove all dockerutils-related files
find /usr/local/bin -name "*dockerutils*" -delete
find ~ -name "*dockerutils*" -type d -exec rm -rf {} +
docker rmi $(docker images --filter=reference="*dockerutils*" -q) 2>/dev/null || true
```

## Verification Steps

After installation, verify everything works:

1. **Basic functionality:**
   ```bash
   dockerutils version
   dockerutils --help
   ```

2. **Docker connectivity:**
   ```bash
   docker ps
   ```

3. **Web server:**
   ```bash
   dockerutils server --port 8080 &
   curl http://localhost:8080/api/health
   ```

4. **Image analysis:**
   ```bash
   # Via web interface: http://localhost:8080
   # Via API:
   curl -X POST http://localhost:8080/api/analyze \
     -H "Content-Type: application/json" \
     -d '{"image_name": "alpine:latest"}'
   ```

---

**Next Steps:** [User Guide](docs/USER_GUIDE.md) for using Docker Utils

**Need Help?** Check [Troubleshooting](TROUBLESHOOTING.md) for common issues

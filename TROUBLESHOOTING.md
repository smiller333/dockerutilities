# Troubleshooting Guide

Common issues and solutions for Docker Utils. If you encounter problems not covered here, please check our [GitHub Issues](https://github.com/your-org/dockerutils/issues) or create a new issue.

## Table of Contents

1. [Installation Issues](#installation-issues)
2. [Docker Connection Issues](#docker-connection-issues)
3. [Analysis Issues](#analysis-issues)
4. [Web Interface Issues](#web-interface-issues)
5. [Performance Issues](#performance-issues)
6. [Container Issues](#container-issues)
7. [API Issues](#api-issues)
8. [Getting Help](#getting-help)

## Installation Issues

### Binary Not Found

**Error:** `dockerutils: command not found`

**Solutions:**
1. **Check if dockerutils is in PATH:**
   ```bash
   which dockerutils
   echo $PATH
   ```

2. **Add to PATH:**
   ```bash
   # Linux/macOS
   export PATH=$PATH:/usr/local/bin
   echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
   source ~/.bashrc
   
   # Windows PowerShell
   $env:PATH += ";C:\path\to\dockerutils"
   ```

3. **Use full path:**
   ```bash
   /usr/local/bin/dockerutils version
   ```

4. **Reinstall dockerutils:**
   ```bash
   # Download and install again
   curl -L https://github.com/your-org/dockerutils/releases/latest/download/dockerutils-linux-amd64 -o dockerutils
   chmod +x dockerutils
   sudo mv dockerutils /usr/local/bin/
   ```

### Permission Denied

**Error:** `permission denied: dockerutils`

**Solutions:**
1. **Make binary executable:**
   ```bash
   chmod +x dockerutils
   ```

2. **Check file permissions:**
   ```bash
   ls -la dockerutils
   # Should show: -rwxr-xr-x
   ```

3. **Move to system directory:**
   ```bash
   sudo mv dockerutils /usr/local/bin/
   ```

### Version Mismatch

**Error:** `incompatible version` or unexpected behavior

**Solutions:**
1. **Check version:**
   ```bash
   dockerutils version
   ```

2. **Download latest release:**
   ```bash
   curl -L https://github.com/your-org/dockerutils/releases/latest/download/dockerutils-linux-amd64 -o dockerutils
   ```

3. **Clear old binaries:**
   ```bash
   which -a dockerutils  # Find all dockerutils binaries
   sudo rm /path/to/old/dockerutils
   ```

## Docker Connection Issues

### Docker Daemon Not Running

**Error:** `Cannot connect to the Docker daemon`

**Solutions:**
1. **Start Docker daemon:**
   ```bash
   # Linux
   sudo systemctl start docker
   sudo systemctl enable docker
   
   # macOS
   open /Applications/Docker.app
   
   # Windows
   # Start Docker Desktop from Start menu
   ```

2. **Check Docker status:**
   ```bash
   # Linux
   sudo systemctl status docker
   
   # All platforms
   docker version
   docker info
   ```

3. **Test Docker connection:**
   ```bash
   docker run hello-world
   ```

### Docker Permission Issues

**Error:** `permission denied while trying to connect to the Docker daemon socket`

**Solutions:**
1. **Add user to docker group:**
   ```bash
   sudo usermod -aG docker $USER
   newgrp docker
   ```

2. **Verify group membership:**
   ```bash
   groups $USER
   id $USER
   ```

3. **Restart shell session:**
   ```bash
   # Log out and log back in, or:
   exec su -l $USER
   ```

4. **Test Docker access:**
   ```bash
   docker ps
   ```

### Docker Socket Issues

**Error:** `dial unix /var/run/docker.sock: connect: no such file or directory`

**Solutions:**
1. **Check Docker socket exists:**
   ```bash
   ls -la /var/run/docker.sock
   ```

2. **Alternative socket locations:**
   ```bash
   # Docker Desktop on macOS
   ls -la ~/.docker/desktop/docker.sock
   export DOCKER_HOST=unix://~/.docker/desktop/docker.sock
   
   # Podman
   ls -la /run/user/$(id -u)/podman/podman.sock
   export DOCKER_HOST=unix:///run/user/$(id -u)/podman/podman.sock
   ```

3. **Fix socket permissions:**
   ```bash
   sudo chmod 666 /var/run/docker.sock
   ```

4. **Use TCP connection:**
   ```bash
   export DOCKER_HOST=tcp://localhost:2376
   ```

### Docker API Version Issues

**Error:** `client version X.X is too new. Maximum supported API version is Y.Y`

**Solutions:**
1. **Check API versions:**
   ```bash
   docker version
   ```

2. **Set specific API version:**
   ```bash
   export DOCKER_API_VERSION=1.40
   dockerutils tools
   ```

3. **Update Docker daemon:**
   ```bash
   # Follow Docker installation guide for your platform
   ```

## Analysis Issues

### Image Not Found

**Error:** `Error response from daemon: pull access denied` or `image not found`

**Solutions:**
1. **Pull image first:**
   ```bash
   docker pull nginx:latest
   docker images
   ```

2. **Check image name spelling:**
   ```bash
   # Correct formats:
   # nginx:latest
   # docker.io/library/nginx:latest
   # registry.example.com/myorg/myapp:v1.0.0
   ```

3. **Login to registry:**
   ```bash
   docker login
   # or for private registries:
   docker login registry.example.com
   ```

4. **Use full image path:**
   ```bash
   # Instead of: nginx:latest
   # Use: docker.io/library/nginx:latest
   ```

### Analysis Hangs or Fails

**Symptoms:** Analysis progress stops or fails with errors

**Solutions:**
1. **Check Docker daemon resources:**
   ```bash
   docker system df
   docker system events
   ```

2. **Try with smaller image:**
   ```bash
   # Start with a small image
   dockerutils tools
   # Then analyze: alpine:latest (5MB)
   ```

3. **Check available disk space:**
   ```bash
   df -h
   # Ensure sufficient space in /tmp and working directory
   ```

4. **Restart dockerutils:**
   ```bash
   pkill dockerutils
   dockerutils tools
   ```

5. **Clean up Docker system:**
   ```bash
   docker system prune -f
   docker volume prune -f
   ```

6. **Check logs for specific errors:**
   ```bash
   dockerutils tools --verbose
   ```

### Large Image Timeouts

**Issue:** Large images (>1GB) fail to analyze or timeout

**Solutions:**
1. **Increase timeout (if running via API):**
   ```bash
   curl -X POST http://localhost:8080/api/analyze \
     -H "Content-Type: application/json" \
     -d '{"image_name": "large-image:latest", "timeout": "10m"}'
   ```

2. **Use asynchronous analysis:**
   ```bash
   curl -X POST http://localhost:8080/api/analyze-async \
     -H "Content-Type: application/json" \
     -d '{"image_name": "large-image:latest"}'
   ```

3. **Ensure sufficient resources:**
   ```bash
   # Check available memory and disk space
   free -h
   df -h
   ```

4. **Analyze in chunks (if possible):**
   ```bash
   # Analyze base layers first, then specific layers
   ```

### Analysis Results Missing

**Issue:** Analysis completes but results are not visible

**Solutions:**
1. **Check tmp directory:**
   ```bash
   ls -la ./tmp/
   ls -la ./tmp/summaries.json
   ```

2. **Check file permissions:**
   ```bash
   ls -la ./tmp/*/info.*.json
   ```

3. **Check analysis status:**
   ```bash
   curl http://localhost:8080/api/summaries
   ```

4. **Clear cache and retry:**
   ```bash
   rm -rf ./tmp/
   # Restart dockerutils and try again
   ```

## Web Interface Issues

### Cannot Access Web Interface

**Error:** `This site can't be reached` at http://localhost:8080

**Solutions:**
1. **Check if server is running:**
   ```bash
   ps aux | grep dockerutils
   netstat -tlnp | grep 8080
   ```

2. **Start the server:**
   ```bash
   dockerutils tools --port 8080
   ```

3. **Try different port:**
   ```bash
   dockerutils tools --port 8081
   open http://localhost:8081
   ```

4. **Check firewall settings:**
   ```bash
   # Linux
   sudo ufw status
   sudo ufw allow 8080
   
   # macOS
   # Check System Preferences > Security & Privacy > Firewall
   ```

5. **Use different interface:**
   ```bash
   dockerutils tools --host 127.0.0.1 --port 8080
   ```

### Web Interface Loads But No Images

**Issue:** Interface loads but shows "No images analyzed yet"

**Solutions:**
1. **Check if any analyses exist:**
   ```bash
   curl http://localhost:8080/api/summaries
   ls ./tmp/
   ```

2. **Analyze an image:**
   ```bash
   # Via web interface: enter "alpine:latest"
   # Via API:
   curl -X POST http://localhost:8080/api/analyze \
     -H "Content-Type: application/json" \
     -d '{"image_name": "alpine:latest"}'
   ```

3. **Check browser console for errors:**
   ```
   Press F12 > Console tab > Look for JavaScript errors
   ```

4. **Try refreshing the page:**
   ```bash
   # Force refresh: Ctrl+F5 or Cmd+Shift+R
   ```

### Images Not Loading in Browser

**Issue:** Analysis exists but images don't load in web interface

**Solutions:**
1. **Check API connectivity:**
   ```bash
   curl http://localhost:8080/api/health
   curl http://localhost:8080/api/summaries
   ```

2. **Check browser developer tools:**
   ```
   F12 > Network tab > Look for failed requests
   F12 > Console tab > Look for JavaScript errors
   ```

3. **Clear browser cache:**
   ```
   Ctrl+Shift+Delete (Windows/Linux)
   Cmd+Shift+Delete (macOS)
   ```

4. **Try different browser:**
   ```bash
   # Test with curl first
   curl http://localhost:8080/api/summaries
   ```

### CORS Issues

**Error:** `Access to fetch at 'http://localhost:8080/api/...' from origin '...' has been blocked by CORS policy`

**Solutions:**
1. **Access directly from same origin:**
   ```
   Open http://localhost:8080 directly (not via file://)
   ```

2. **Check server CORS settings:**
   ```bash
   # Server should handle CORS automatically
   curl -H "Origin: http://localhost:3000" \
        -H "Access-Control-Request-Method: GET" \
        -H "Access-Control-Request-Headers: Content-Type" \
        -X OPTIONS \
        http://localhost:8080/api/health
   ```

## Performance Issues

### Slow Analysis

**Issue:** Image analysis takes very long time

**Solutions:**
1. **Check system resources:**
   ```bash
   top
   htop
   docker stats
   ```

2. **Close other applications:**
   ```bash
   # Free up memory and CPU
   ```

3. **Use SSD storage:**
   ```bash
   # Move tmp directory to SSD
   mkdir /path/to/ssd/dockerutils-tmp
   dockerutils tools --tmp-dir /path/to/ssd/dockerutils-tmp
   ```

4. **Increase Docker resources:**
   ```bash
   # Docker Desktop: Settings > Resources > Advanced
   # Increase Memory and CPU allocation
   ```

5. **Analyze smaller images first:**
   ```bash
   # Test with: alpine:latest, busybox:latest
   ```

### High Memory Usage

**Issue:** dockerutils consumes too much memory

**Solutions:**
1. **Monitor memory usage:**
   ```bash
   top -p $(pgrep dockerutils)
   ps aux | grep dockerutils
   ```

2. **Limit concurrent analyses:**
   ```bash
   # Analyze one image at a time
   # Wait for completion before starting next
   ```

3. **Clean up old analyses:**
   ```bash
   # Remove old analysis data
   find ./tmp -name "*.json" -mtime +7 -delete
   ```

4. **Restart dockerutils periodically:**
   ```bash
   # For long-running sessions
   pkill dockerutils
   dockerutils tools
   ```

### Disk Space Issues

**Issue:** Running out of disk space during analysis

**Solutions:**
1. **Check disk usage:**
   ```bash
   df -h
   du -sh ./tmp/
   ```

2. **Clean up Docker:**
   ```bash
   docker system prune -af
   docker volume prune -f
   ```

3. **Remove old analyses:**
   ```bash
   rm -rf ./tmp/old-analysis-*
   ```

4. **Use different tmp directory:**
   ```bash
   dockerutils tools --tmp-dir /path/to/large/disk/tmp
   ```

5. **Enable auto-cleanup:**
   ```bash
   # Clean up analyses older than 7 days
   find ./tmp -name "info.*.json" -mtime +7 -delete
   ```

## Container Issues

### Container Won't Start

**Error:** Docker container fails to start

**Solutions:**
1. **Check port availability:**
   ```bash
   lsof -i :8080
   netstat -tlnp | grep 8080
   ```

2. **Use different port:**
   ```bash
   ./scripts/docker-run.sh run --port 3000
   ```

3. **Check Docker socket mount:**
   ```bash
   ls -la /var/run/docker.sock
   # For Docker Desktop on macOS:
   ls -la ~/.docker/desktop/docker.sock
   ```

4. **View container logs:**
   ```bash
   ./scripts/docker-run.sh logs
   # or
   docker logs dockerutils-viewer
   ```

5. **Remove existing container:**
   ```bash
   docker rm dockerutils-viewer
   ./scripts/docker-run.sh run
   ```

### Container Cannot Access Docker

**Issue:** Container runs but cannot analyze images

**Solutions:**
1. **Verify Docker socket mount:**
   ```bash
   docker exec dockerutils-viewer ls -la /var/run/docker.sock
   ```

2. **Check socket permissions:**
   ```bash
   ls -la /var/run/docker.sock
   ```

3. **Use correct socket path:**
   ```bash
   # Docker Desktop on macOS
   ./scripts/docker-run.sh run --socket ~/.docker/desktop/docker.sock
   ```

4. **Test Docker access from container:**
   ```bash
   docker exec dockerutils-viewer docker ps
   ```

### Data Not Persisting

**Issue:** Analysis data is lost when container restarts

**Solutions:**
1. **Use persistent storage:**
   ```bash
   ./scripts/docker-run.sh run-persistent
   ```

2. **Check volume mount:**
   ```bash
   docker inspect dockerutils-viewer | grep -A 10 "Mounts"
   ```

3. **Verify data directory:**
   ```bash
   ls -la ./data
   ```

4. **Fix permissions:**
   ```bash
   chmod 755 ./data
   ```

## API Issues

### API Returns 404

**Error:** `404 Not Found` for API endpoints

**Solutions:**
1. **Check correct URL:**
   ```bash
   # Correct endpoints:
   curl http://localhost:8080/api/health
   curl http://localhost:8080/api/summaries
   ```

2. **Verify server is running:**
   ```bash
   curl http://localhost:8080/
   ```

3. **Check server logs:**
   ```bash
   # Look for routing errors in dockerutils output
   ```

### API Returns 500 Error

**Error:** `500 Internal Server Error`

**Solutions:**
1. **Check server logs:**
   ```bash
   dockerutils tools --verbose
   ```

2. **Verify request format:**
   ```bash
   curl -X POST http://localhost:8080/api/analyze \
     -H "Content-Type: application/json" \
     -d '{"image_name": "nginx:latest"}'
   ```

3. **Test with simple request:**
   ```bash
   curl http://localhost:8080/api/health
   ```

4. **Restart server:**
   ```bash
   pkill dockerutils
   dockerutils tools
   ```

### JSON Parse Errors

**Error:** `invalid character` or JSON parsing errors

**Solutions:**
1. **Validate JSON:**
   ```bash
   echo '{"image_name": "nginx:latest"}' | jq .
   ```

2. **Use proper Content-Type:**
   ```bash
   curl -X POST http://localhost:8080/api/analyze \
     -H "Content-Type: application/json" \
     -d '{"image_name": "nginx:latest"}'
   ```

3. **Escape special characters:**
   ```bash
   # For image names with special characters
   curl -X POST http://localhost:8080/api/analyze \
     -H "Content-Type: application/json" \
     -d '{"image_name": "registry.example.com/org/app:v1.0.0"}'
   ```

## Getting Help

### Diagnostic Information

When reporting issues, include:

```bash
# System information
uname -a
docker version
dockerutils version

# Error details
dockerutils tools --verbose 2>&1 | head -50

# Docker status
docker info
docker ps -a

# Resource usage
df -h
free -h
```

### Log Collection

```bash
# Collect logs
dockerutils tools --verbose > dockerutils.log 2>&1 &
# Reproduce the issue
# Then:
pkill dockerutils

# Docker logs (if using container)
./scripts/docker-run.sh logs > container.log
```

### Contact and Support

- **GitHub Issues**: [Report bugs and request features](https://github.com/your-org/dockerutils/issues)
- **GitHub Discussions**: [Ask questions and get help](https://github.com/your-org/dockerutils/discussions)
- **Documentation**: [Read the full documentation](docs/)
- **Examples**: [Check usage examples](EXAMPLES.md)

### Creating Good Bug Reports

Include:
1. **Steps to reproduce**
2. **Expected behavior**
3. **Actual behavior**
4. **System information** (OS, Docker version, dockerutils version)
5. **Log output** (with `--verbose` flag)
6. **Screenshots** (for UI issues)

### Before Reporting

1. **Search existing issues**: Check if the problem is already reported
2. **Try latest version**: Update to the latest release
3. **Test with minimal case**: Use simple images like `alpine:latest`
4. **Check documentation**: Review this troubleshooting guide and user documentation

---

**Still having issues?** Please [create an issue](https://github.com/your-org/dockerutils/issues/new) with detailed information about your problem.

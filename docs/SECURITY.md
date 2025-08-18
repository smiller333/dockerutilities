# Security Policy

This document outlines the security policies, procedures, and best practices for Docker Utils.

## Table of Contents

1. [Security Overview](#security-overview)
2. [Supported Versions](#supported-versions)
3. [Reporting Security Vulnerabilities](#reporting-security-vulnerabilities)
4. [Security Architecture](#security-architecture)
5. [Docker Socket Security](#docker-socket-security)
6. [Image Security](#image-security)
7. [Network Security](#network-security)
8. [Data Security](#data-security)
9. [Deployment Security](#deployment-security)
10. [Security Best Practices](#security-best-practices)
11. [Security Monitoring](#security-monitoring)
12. [Incident Response](#incident-response)

## Security Overview

Docker Utils requires elevated privileges to access the Docker daemon, which introduces significant security considerations. This document provides comprehensive guidance for secure deployment and operation.

### Security Risk Assessment

| Risk Category | Risk Level | Mitigation |
|---------------|------------|------------|
| **Docker Socket Access** | 🔴 HIGH | Restricted access, socket proxy, non-root containers |
| **Untrusted Images** | 🟡 MEDIUM | Registry validation, security scanning, sandboxing |
| **Network Exposure** | 🟡 MEDIUM | TLS encryption, access controls, rate limiting |
| **Data Persistence** | 🟢 LOW | Encryption at rest, secure cleanup, access controls |
| **Code Injection** | 🟡 MEDIUM | Input validation, sandboxed execution, security scanning |

## Supported Versions

We provide security updates for the following versions:

| Version | Supported | Status |
|---------|-----------|--------|
| 1.x.x   | ✅ Yes    | Active development |
| 0.9.x   | ✅ Yes    | Security updates only |
| 0.8.x   | ❌ No     | End of life |
| < 0.8   | ❌ No     | End of life |

## Reporting Security Vulnerabilities

### How to Report

**DO NOT** create public GitHub issues for security vulnerabilities.

Instead, please report security vulnerabilities privately:

1. **Email**: Send details to `security@dockerutils.dev`
2. **Encrypted Email**: Use our PGP key for sensitive reports
3. **GitHub Security Advisories**: Use GitHub's private vulnerability reporting

### What to Include

Please include the following information:

- Description of the vulnerability
- Steps to reproduce the issue
- Affected versions
- Potential impact assessment
- Suggested mitigation (if any)
- Your contact information

### Response Timeline

| Timeframe | Action |
|-----------|--------|
| **24 hours** | Initial acknowledgment |
| **72 hours** | Initial assessment and triage |
| **7 days** | Detailed analysis and response plan |
| **30 days** | Fix development and testing |
| **45 days** | Public disclosure (coordinated) |

### Acknowledgments

We maintain a security hall of fame for responsible disclosure:

- [Security Contributors](SECURITY_CONTRIBUTORS.md)

## Security Architecture

### Threat Model

```
┌─────────────────────────────────────────────────────────┐
│                    Threat Landscape                     │
├─────────────────────────────────────────────────────────┤
│  External Threats          │  Internal Threats          │
│  ┌─────────────────────────┼─────────────────────────┐   │
│  │ • Malicious Images      │ • Privilege Escalation │   │
│  │ • Network Attacks       │ • Data Exfiltration    │   │
│  │ • Code Injection        │ • Resource Abuse       │   │
│  │ • Registry Poisoning    │ • Configuration Drift  │   │
│  └─────────────────────────┼─────────────────────────┘   │
├─────────────────────────────────────────────────────────┤
│                  Security Controls                      │
│  ┌─────────────────────────────────────────────────────┐ │
│  │ Authentication │ Authorization │ Audit & Monitoring │ │
│  │      +         │       +       │         +          │ │
│  │ Input Valid.   │  Network Sec. │   Secure Storage   │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────┘
```

### Security Boundaries

1. **Host System Boundary**: Docker socket access grants host-level privileges
2. **Container Boundary**: Analysis containers should be isolated and ephemeral
3. **Network Boundary**: API exposure requires authentication and encryption
4. **Data Boundary**: Sensitive data must be encrypted and access-controlled

## Docker Socket Security

### Critical Security Warning

⚠️ **DANGER**: Docker socket access is equivalent to root access on the host system.

### Risk Mitigation Strategies

#### 1. Docker Socket Proxy (Recommended)

```bash
# Deploy socket proxy with restricted permissions
docker run -d \
  --name docker-socket-proxy \
  --restart unless-stopped \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -p 127.0.0.1:2375:2375 \
  -e CONTAINERS=1 \
  -e IMAGES=1 \
  -e BUILD=1 \
  -e SERVICES=0 \
  -e VOLUMES=0 \
  -e NETWORKS=0 \
  -e SECRETS=0 \
  --security-opt no-new-privileges:true \
  --cap-drop ALL \
  --read-only \
  tecnativa/docker-socket-proxy

# Connect Docker Utils to proxy instead of socket
docker run -d \
  --name dockerutils \
  -e DOCKER_HOST=tcp://127.0.0.1:2375 \
  dockerutils:latest
```

#### 2. Socket Permissions

```bash
# Restrict socket permissions
sudo chmod 660 /var/run/docker.sock
sudo chown root:docker /var/run/docker.sock

# Create dedicated group
sudo groupadd dockerutils
sudo usermod -aG dockerutils dockerutils-user
sudo chown root:dockerutils /var/run/docker.sock
```

#### 3. Docker Rootless Mode

```bash
# Install Docker in rootless mode
curl -fsSL https://get.docker.com/rootless | sh

# Configure dockerutils for rootless
export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock
dockerutils server
```

### Socket Access Monitoring

```bash
#!/bin/bash
# monitor-socket-access.sh

# Monitor Docker socket access
sudo auditctl -w /var/run/docker.sock -p rwxa -k docker_socket

# Create alerting rule
cat << EOF > /etc/audit/rules.d/docker-socket.rules
-w /var/run/docker.sock -p rwxa -k docker_socket
-a always,exit -F arch=b64 -S connect -F key=docker_connect
EOF

# Monitor logs
sudo ausearch -k docker_socket | auditctl -l
```

## Image Security

### Trusted Registry Validation

```go
// Built-in trusted registries
var trustedRegistries = []string{
    "docker.io",           // Docker Hub
    "gcr.io",             // Google Container Registry
    "quay.io",            // Red Hat Quay
    "registry.k8s.io",    // Kubernetes Registry
    "ghcr.io",            // GitHub Container Registry
}

// Custom registry validation
func ValidateImageSource(imageName string) (bool, string) {
    registry := extractRegistry(imageName)
    
    for _, trusted := range trustedRegistries {
        if registry == trusted {
            return true, "Trusted registry"
        }
    }
    
    return false, fmt.Sprintf("Untrusted registry: %s", registry)
}
```

### Image Scanning Integration

```yaml
# docker-compose with Trivy scanning
version: '3.8'
services:
  trivy:
    image: aquasec/trivy:latest
    command: server --listen 0.0.0.0:4954
    ports:
      - "4954:4954"
    volumes:
      - trivy-cache:/root/.cache/trivy

  dockerutils:
    image: dockerutils:latest
    environment:
      - TRIVY_SERVER_URL=http://trivy:4954
      - SECURITY_SCAN_ENABLED=true
    depends_on:
      - trivy

volumes:
  trivy-cache:
```

### Image Signature Verification

```bash
# Install Cosign for image verification
go install github.com/sigstore/cosign/cmd/cosign@latest

# Verify image signatures
cosign verify --key cosign.pub dockerutils:latest

# Integration in deployment
cat << EOF > verify-image.sh
#!/bin/bash
IMAGE=$1
if ! cosign verify --key /etc/cosign/cosign.pub "$IMAGE"; then
    echo "ERROR: Image signature verification failed"
    exit 1
fi
echo "Image signature verified successfully"
EOF
```

### Content Trust

```bash
# Enable Docker Content Trust
export DOCKER_CONTENT_TRUST=1

# Sign images during build
docker trust key generate dockerutils
docker trust signer add --key dockerutils.pub dockerutils dockerutils:latest
docker push dockerutils:latest
```

## Network Security

### TLS Configuration

```yaml
# nginx SSL configuration
server {
    listen 443 ssl http2;
    server_name dockerutils.example.com;

    # SSL/TLS Configuration
    ssl_certificate /etc/ssl/certs/dockerutils.crt;
    ssl_certificate_key /etc/ssl/private/dockerutils.key;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384;
    ssl_prefer_server_ciphers off;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;

    # HSTS and Security Headers
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains; preload" always;
    add_header X-Frame-Options "DENY" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header Referrer-Policy "strict-origin-when-cross-origin" always;
    add_header Content-Security-Policy "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; connect-src 'self';" always;

    # Rate Limiting
    limit_req_zone $binary_remote_addr zone=dockerutils:10m rate=10r/s;
    limit_req zone=dockerutils burst=20 nodelay;
}
```

### Authentication and Authorization

```go
// JWT-based authentication middleware
func AuthMiddleware() gin.HandlerFunc {
    return gin.HandlerFunc(func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(401, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil || !token.Valid {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    })
}

// RBAC implementation
type Permission struct {
    Resource string `json:"resource"`
    Action   string `json:"action"`
}

type Role struct {
    Name        string       `json:"name"`
    Permissions []Permission `json:"permissions"`
}

var roles = map[string]Role{
    "viewer": {
        Name: "viewer",
        Permissions: []Permission{
            {Resource: "images", Action: "read"},
            {Resource: "analyses", Action: "read"},
        },
    },
    "analyst": {
        Name: "analyst", 
        Permissions: []Permission{
            {Resource: "images", Action: "read"},
            {Resource: "images", Action: "analyze"},
            {Resource: "analyses", Action: "read"},
            {Resource: "analyses", Action: "create"},
        },
    },
    "admin": {
        Name: "admin",
        Permissions: []Permission{
            {Resource: "*", Action: "*"},
        },
    },
}
```

### Network Isolation

```yaml
# Docker Compose with network isolation
version: '3.8'
services:
  dockerutils:
    networks:
      - frontend
      - backend
    
  nginx:
    networks:
      - frontend
      
  database:
    networks:
      - backend

networks:
  frontend:
    driver: bridge
    ipam:
      config:
        - subnet: 172.20.0.0/24
  backend:
    driver: bridge
    internal: true
    ipam:
      config:
        - subnet: 172.21.0.0/24
```

### Firewall Configuration

```bash
# UFW configuration
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow ssh
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow from 172.20.0.0/24 to any port 8080
sudo ufw enable

# iptables rules
sudo iptables -A INPUT -i lo -j ACCEPT
sudo iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT
sudo iptables -A INPUT -p tcp --dport 22 -j ACCEPT
sudo iptables -A INPUT -p tcp --dport 80 -j ACCEPT
sudo iptables -A INPUT -p tcp --dport 443 -j ACCEPT
sudo iptables -A INPUT -s 172.20.0.0/24 -p tcp --dport 8080 -j ACCEPT
sudo iptables -P INPUT DROP
```

## Data Security

### Encryption at Rest

```bash
# LUKS encryption for data volumes
sudo cryptsetup luksFormat /dev/sdb1
sudo cryptsetup luksOpen /dev/sdb1 dockerutils_data
sudo mkfs.ext4 /dev/mapper/dockerutils_data
sudo mount /dev/mapper/dockerutils_data /opt/dockerutils/data

# Automated mounting with key file
sudo dd if=/dev/urandom of=/etc/dockerutils/luks.key bs=512 count=4
sudo chmod 600 /etc/dockerutils/luks.key
sudo cryptsetup luksAddKey /dev/sdb1 /etc/dockerutils/luks.key

# fstab entry
echo "/dev/mapper/dockerutils_data /opt/dockerutils/data ext4 defaults 0 2" >> /etc/fstab
```

### Encryption in Transit

```go
// TLS configuration for API server
func createTLSConfig() *tls.Config {
    return &tls.Config{
        MinVersion:               tls.VersionTLS12,
        CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
            tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
        },
    }
}
```

### Secure Data Cleanup

```go
// Secure deletion of sensitive data
func secureDelete(path string) error {
    file, err := os.OpenFile(path, os.O_WRONLY, 0)
    if err != nil {
        return err
    }
    defer file.Close()

    info, err := file.Stat()
    if err != nil {
        return err
    }

    // Overwrite with random data multiple times
    for i := 0; i < 3; i++ {
        file.Seek(0, 0)
        randomData := make([]byte, info.Size())
        rand.Read(randomData)
        file.Write(randomData)
        file.Sync()
    }

    return os.Remove(path)
}
```

## Deployment Security

### Container Security

```dockerfile
# Security-hardened Dockerfile
FROM scratch

# Copy CA certificates
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Use specific non-root user
COPY --from=alpine:latest /etc/passwd /etc/passwd
USER 65534:65534

# Add security labels
LABEL security.contact="security@dockerutils.dev" \
      security.scan="enabled" \
      security.updates="auto"

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD ["/app/dockerutils", "health"]

# Run as non-root
ENTRYPOINT ["/app/dockerutils"]
```

### Kubernetes Security

```yaml
# SecurityContext configuration
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dockerutils
spec:
  template:
    spec:
      securityContext:
        runAsNonRoot: true
        runAsUser: 65534
        runAsGroup: 65534
        fsGroup: 65534
        seccompProfile:
          type: RuntimeDefault
      containers:
      - name: dockerutils
        securityContext:
          allowPrivilegeEscalation: false
          readOnlyRootFilesystem: true
          capabilities:
            drop:
            - ALL
            add:
            - DAC_OVERRIDE  # Only if Docker socket access required
        resources:
          limits:
            memory: "1Gi"
            cpu: "1000m"
          requests:
            memory: "512Mi"
            cpu: "500m"
```

### Pod Security Standards

```yaml
# Pod Security Policy
apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: dockerutils-psp
spec:
  privileged: false
  allowPrivilegeEscalation: false
  requiredDropCapabilities:
    - ALL
  allowedCapabilities:
    - DAC_OVERRIDE
  volumes:
    - 'configMap'
    - 'emptyDir'
    - 'projected'
    - 'secret'
    - 'downwardAPI'
    - 'persistentVolumeClaim'
    - 'hostPath'  # Only for Docker socket
  allowedHostPaths:
    - pathPrefix: "/var/run/docker.sock"
      readOnly: true
  runAsUser:
    rule: 'MustRunAsNonRoot'
  seLinux:
    rule: 'RunAsAny'
  fsGroup:
    rule: 'RunAsAny'
```

### Network Policies

```yaml
# Kubernetes Network Policy
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: dockerutils-netpol
spec:
  podSelector:
    matchLabels:
      app: dockerutils
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          name: ingress-nginx
    ports:
    - protocol: TCP
      port: 8080
  egress:
  - to: []
    ports:
    - protocol: TCP
      port: 53
    - protocol: UDP
      port: 53
  - to:
    - namespaceSelector:
        matchLabels:
          name: kube-system
    ports:
    - protocol: TCP
      port: 443
```

## Security Best Practices

### Input Validation

```go
// Comprehensive input validation
func validateImageName(imageName string) error {
    if len(imageName) == 0 {
        return errors.New("image name cannot be empty")
    }
    
    if len(imageName) > 255 {
        return errors.New("image name too long")
    }
    
    // Docker image name regex
    pattern := regexp.MustCompile(`^[a-z0-9]+(?:[._-][a-z0-9]+)*(?:/[a-z0-9]+(?:[._-][a-z0-9]+)*)*(?::[a-zA-Z0-9._-]+)?$`)
    if !pattern.MatchString(imageName) {
        return errors.New("invalid image name format")
    }
    
    // Check for suspicious patterns
    suspicious := []string{"../", "../../", "\\", "$", "`", ";", "&", "|"}
    for _, pattern := range suspicious {
        if strings.Contains(imageName, pattern) {
            return errors.New("potentially malicious image name")
        }
    }
    
    return nil
}
```

### Rate Limiting

```go
// Rate limiting implementation
import "golang.org/x/time/rate"

type RateLimiter struct {
    limiters map[string]*rate.Limiter
    mu       sync.RWMutex
}

func (rl *RateLimiter) GetLimiter(clientID string) *rate.Limiter {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    limiter, exists := rl.limiters[clientID]
    if !exists {
        limiter = rate.NewLimiter(rate.Limit(10), 20) // 10 req/sec, burst 20
        rl.limiters[clientID] = limiter
    }
    
    return limiter
}
```

### Resource Limits

```yaml
# Docker Compose resource limits
version: '3.8'
services:
  dockerutils:
    deploy:
      resources:
        limits:
          memory: 1G
          cpus: '1.0'
        reservations:
          memory: 512M
          cpus: '0.5'
      restart_policy:
        condition: on-failure
        delay: 5s
        max_attempts: 3
        window: 120s
    ulimits:
      nproc: 100
      nofile:
        soft: 1024
        hard: 2048
```

## Security Monitoring

### Audit Logging

```go
// Audit logging implementation
type AuditLog struct {
    Timestamp string `json:"timestamp"`
    UserID    string `json:"user_id"`
    Action    string `json:"action"`
    Resource  string `json:"resource"`
    Result    string `json:"result"`
    IPAddress string `json:"ip_address"`
    UserAgent string `json:"user_agent"`
}

func logAuditEvent(userID, action, resource, result string, c *gin.Context) {
    event := AuditLog{
        Timestamp: time.Now().UTC().Format(time.RFC3339),
        UserID:    userID,
        Action:    action,
        Resource:  resource,
        Result:    result,
        IPAddress: c.ClientIP(),
        UserAgent: c.GetHeader("User-Agent"),
    }
    
    logJSON, _ := json.Marshal(event)
    log.Printf("AUDIT: %s", logJSON)
}
```

### Security Metrics

```yaml
# Prometheus alerting rules
groups:
- name: dockerutils-security
  rules:
  - alert: HighFailedAuthRate
    expr: rate(auth_failures_total[5m]) > 0.1
    for: 2m
    labels:
      severity: warning
    annotations:
      summary: "High authentication failure rate"
      
  - alert: SuspiciousImageAnalysis
    expr: increase(untrusted_image_analysis_total[1h]) > 10
    for: 0m
    labels:
      severity: critical
    annotations:
      summary: "Multiple untrusted image analyses detected"
      
  - alert: ResourceExhaustion
    expr: container_memory_usage_bytes / container_spec_memory_limit_bytes > 0.9
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "Container memory usage above 90%"
```

### Intrusion Detection

```bash
#!/bin/bash
# security-monitor.sh

# Monitor for suspicious activities
while true; do
    # Check for privilege escalation attempts
    if grep -q "sudo\|su -" /var/log/auth.log; then
        echo "$(date): Privilege escalation detected" >> /var/log/security-alerts.log
    fi
    
    # Monitor Docker socket access
    if ss -lx | grep -q "/var/run/docker.sock"; then
        CONNECTIONS=$(ss -lx | grep "/var/run/docker.sock" | wc -l)
        if [ "$CONNECTIONS" -gt 5 ]; then
            echo "$(date): Unusual Docker socket connections: $CONNECTIONS" >> /var/log/security-alerts.log
        fi
    fi
    
    # Check for unusual network activity
    CONNECTIONS=$(netstat -tn | grep :8080 | wc -l)
    if [ "$CONNECTIONS" -gt 100 ]; then
        echo "$(date): High number of connections: $CONNECTIONS" >> /var/log/security-alerts.log
    fi
    
    sleep 60
done
```

## Incident Response

### Response Plan

1. **Detection**: Automated monitoring and alerting
2. **Assessment**: Determine severity and impact
3. **Containment**: Isolate affected systems
4. **Eradication**: Remove threat and vulnerabilities
5. **Recovery**: Restore normal operations
6. **Lessons Learned**: Post-incident analysis

### Emergency Contacts

| Role | Contact | Escalation |
|------|---------|------------|
| Security Team | security@dockerutils.dev | 24/7 |
| DevOps Team | devops@dockerutils.dev | Business hours |
| Legal Team | legal@dockerutils.dev | As needed |

### Incident Classification

| Severity | Definition | Response Time |
|----------|------------|---------------|
| **Critical** | Active exploitation, data breach | 15 minutes |
| **High** | Potential exploitation, system compromise | 1 hour |
| **Medium** | Security weakness, policy violation | 4 hours |
| **Low** | Security advisory, best practice | 24 hours |

### Recovery Procedures

```bash
#!/bin/bash
# incident-response.sh

case "$1" in
    "compromise")
        echo "System compromise detected - initiating containment"
        # Isolate containers
        docker pause $(docker ps -q)
        # Block network access
        iptables -A INPUT -j DROP
        iptables -A OUTPUT -j DROP
        # Preserve evidence
        docker logs --details $(docker ps -aq) > /incident/docker-logs.txt
        # Alert team
        curl -X POST "$SLACK_WEBHOOK" -d '{"text":"SECURITY INCIDENT: System compromise detected"}'
        ;;
    "restore")
        echo "Restoring from secure backup"
        # Restore from clean backup
        docker stop $(docker ps -q)
        docker system prune -af
        docker pull dockerutils:latest-secure
        # Restore data from encrypted backup
        mount /dev/mapper/backup /mnt/backup
        rsync -av /mnt/backup/ /opt/dockerutils/data/
        ;;
esac
```

---

**Security Contacts:**
- **Emergency**: security@dockerutils.dev
- **General Security**: security@dockerutils.dev
- **Bug Bounty**: security+bounty@dockerutils.dev

**Related Documentation:**
- [Deployment Guide](DEPLOYMENT.md) - Secure deployment practices
- [Architecture Overview](ARCHITECTURE.md) - Security architecture details
- [Contributing Guide](../CONTRIBUTING.md) - Secure development practices

# Deployment Guide

This guide provides comprehensive instructions for deploying Docker Utils in production environments across various platforms and deployment scenarios.

## Table of Contents

1. [Deployment Overview](#deployment-overview)
2. [Container Deployment](#container-deployment)
3. [Kubernetes Deployment](#kubernetes-deployment)
4. [Cloud Platform Deployment](#cloud-platform-deployment)
5. [Bare Metal Deployment](#bare-metal-deployment)
6. [Configuration Management](#configuration-management)
7. [Security Considerations](#security-considerations)
8. [Monitoring and Logging](#monitoring-and-logging)
9. [High Availability](#high-availability)
10. [Troubleshooting](#troubleshooting)

## Deployment Overview

Docker Utils can be deployed in several configurations depending on your requirements:

- **Standalone Container** - Single container with Docker socket access
- **Kubernetes Pod** - Managed container orchestration
- **Cloud Instances** - Virtual machines with container runtime
- **Bare Metal** - Direct installation on physical servers

### Architecture Considerations

```
┌─────────────────────────────────────────────────────────┐
│                  Production Deployment                  │
├─────────────────────────────────────────────────────────┤
│  Load Balancer (optional)                              │
│  ┌─────────────────────────────────────────────────────┐ │
│  │                Reverse Proxy                        │ │
│  │            (nginx/traefik/envoy)                    │ │
│  └─────────────────────────────────────────────────────┘ │
├─────────────────────────────────────────────────────────┤
│  Docker Utils Application Layer                         │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐    │ │
│  │  │ Container 1 │ │ Container 2 │ │ Container N │    │ │
│  │  │  Port 8080  │ │  Port 8080  │ │  Port 8080  │    │ │
│  │  └─────────────┘ └─────────────┘ └─────────────┘    │ │
│  └─────────────────────────────────────────────────────┘ │
├─────────────────────────────────────────────────────────┤
│  Storage Layer                                          │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  Persistent Storage  │  Temporary Storage           │ │
│  │  - Analysis Results  │  - Image Extractions        │ │
│  │  - Configuration     │  - Build Contexts           │ │
│  └─────────────────────────────────────────────────────┘ │
├─────────────────────────────────────────────────────────┤
│  Docker Engine                                          │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  Docker Socket: /var/run/docker.sock               │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────┘
```

## Container Deployment

### Docker Compose (Recommended)

Create a production-ready `docker-compose.yml`:

```yaml
version: '3.8'

services:
  dockerutils:
    image: dockerutils:latest
    restart: unless-stopped
    ports:
      - "8080:8080"
    volumes:
      # Docker socket access (SECURITY RISK - see security section)
      - /var/run/docker.sock:/var/run/docker.sock:ro
      # Persistent storage for analysis data
      - dockerutils_data:/app/data
      # Optional: custom configuration
      - ./config:/app/config:ro
    environment:
      - DOCKERUTILS_PORT=8080
      - DOCKERUTILS_HOST=0.0.0.0
      - DOCKERUTILS_TMP_DIR=/app/data/tmp
      - DOCKERUTILS_LOG_LEVEL=info
    healthcheck:
      test: ["CMD", "wget", "--quiet", "--tries=1", "--spider", "http://localhost:8080/api/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
    deploy:
      resources:
        limits:
          memory: 1G
          cpus: '1.0'
        reservations:
          memory: 512M
          cpus: '0.5'
    security_opt:
      - no-new-privileges:true
    cap_drop:
      - ALL
    cap_add:
      - DAC_OVERRIDE  # Required for Docker socket access
    read_only: true
    tmpfs:
      - /tmp:noexec,nosuid,size=100m

  # Optional: Reverse proxy
  nginx:
    image: nginx:alpine
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro
      - ./ssl:/etc/nginx/ssl:ro
    depends_on:
      - dockerutils

volumes:
  dockerutils_data:
    driver: local
    driver_opts:
      type: none
      o: bind
      device: /opt/dockerutils/data

networks:
  default:
    name: dockerutils_network
```

### Standalone Docker Container

```bash
# Pull the image
docker pull dockerutils:latest

# Run with minimal configuration
docker run -d \
  --name dockerutils \
  --restart unless-stopped \
  -p 8080:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -v dockerutils_data:/app/data \
  -e DOCKERUTILS_PORT=8080 \
  -e DOCKERUTILS_HOST=0.0.0.0 \
  dockerutils:latest

# Run with enhanced security
docker run -d \
  --name dockerutils \
  --restart unless-stopped \
  -p 8080:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -v dockerutils_data:/app/data \
  --read-only \
  --tmpfs /tmp:noexec,nosuid,size=100m \
  --security-opt no-new-privileges:true \
  --cap-drop ALL \
  --cap-add DAC_OVERRIDE \
  --memory 1g \
  --cpus 1.0 \
  -e DOCKERUTILS_PORT=8080 \
  -e DOCKERUTILS_HOST=0.0.0.0 \
  dockerutils:latest
```

### Building Production Image

Create an optimized Dockerfile for production:

```dockerfile
# Multi-stage build for production
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Create appuser for security
RUN adduser -D -g '' dockerutils

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o dockerutils .

# Production stage
FROM scratch

# Import from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd

# Copy binary
COPY --from=builder /build/dockerutils /app/dockerutils

# Use non-root user
USER dockerutils

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD ["/app/dockerutils", "version"]

# Run the binary
ENTRYPOINT ["/app/dockerutils"]
CMD ["tools", "--port", "8080", "--host", "0.0.0.0"]
```

## Kubernetes Deployment

### Namespace and ConfigMap

```yaml
# namespace.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: dockerutils
  labels:
    name: dockerutils

---
# configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: dockerutils-config
  namespace: dockerutils
data:
  PORT: "8080"
  HOST: "0.0.0.0"
  LOG_LEVEL: "info"
  TMP_DIR: "/app/data/tmp"
```

### Deployment

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dockerutils
  namespace: dockerutils
  labels:
    app: dockerutils
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 1
  selector:
    matchLabels:
      app: dockerutils
  template:
    metadata:
      labels:
        app: dockerutils
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "8080"
        prometheus.io/path: "/metrics"
    spec:
      serviceAccountName: dockerutils
      securityContext:
        runAsNonRoot: true
        runAsUser: 1001
        runAsGroup: 1001
        fsGroup: 1001
      containers:
      - name: dockerutils
        image: dockerutils:v1.0.0
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8080
          name: http
          protocol: TCP
        env:
        - name: DOCKERUTILS_PORT
          valueFrom:
            configMapKeyRef:
              name: dockerutils-config
              key: PORT
        - name: DOCKERUTILS_HOST
          valueFrom:
            configMapKeyRef:
              name: dockerutils-config
              key: HOST
        - name: DOCKERUTILS_LOG_LEVEL
          valueFrom:
            configMapKeyRef:
              name: dockerutils-config
              key: LOG_LEVEL
        resources:
          limits:
            memory: "1Gi"
            cpu: "1000m"
          requests:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /api/health
            port: http
          initialDelaySeconds: 30
          periodSeconds: 30
          timeoutSeconds: 10
          failureThreshold: 3
        readinessProbe:
          httpGet:
            path: /api/health
            port: http
          initialDelaySeconds: 10
          periodSeconds: 5
          timeoutSeconds: 5
          failureThreshold: 3
        volumeMounts:
        - name: docker-sock
          mountPath: /var/run/docker.sock
          readOnly: true
        - name: data
          mountPath: /app/data
        - name: tmp
          mountPath: /tmp
        securityContext:
          allowPrivilegeEscalation: false
          readOnlyRootFilesystem: true
          capabilities:
            drop:
            - ALL
            add:
            - DAC_OVERRIDE
      volumes:
      - name: docker-sock
        hostPath:
          path: /var/run/docker.sock
          type: Socket
      - name: data
        persistentVolumeClaim:
          claimName: dockerutils-data
      - name: tmp
        emptyDir:
          sizeLimit: 1Gi
      nodeSelector:
        kubernetes.io/os: linux
      tolerations:
      - key: "node-role.kubernetes.io/master"
        operator: "Exists"
        effect: "NoSchedule"
```

### Service and Ingress

```yaml
# service.yaml
apiVersion: v1
kind: Service
metadata:
  name: dockerutils-service
  namespace: dockerutils
  labels:
    app: dockerutils
spec:
  type: ClusterIP
  ports:
  - port: 80
    targetPort: 8080
    protocol: TCP
    name: http
  selector:
    app: dockerutils

---
# ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: dockerutils-ingress
  namespace: dockerutils
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
    nginx.ingress.kubernetes.io/rate-limit: "100"
    nginx.ingress.kubernetes.io/proxy-body-size: "50m"
spec:
  tls:
  - hosts:
    - dockerutils.example.com
    secretName: dockerutils-tls
  rules:
  - host: dockerutils.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: dockerutils-service
            port:
              number: 80
```

### Persistent Storage

```yaml
# pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: dockerutils-data
  namespace: dockerutils
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
  storageClassName: fast-ssd

---
# Optional: Storage class for high performance
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: fast-ssd
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-ssd
  zones: us-central1-a,us-central1-b
reclaimPolicy: Retain
allowVolumeExpansion: true
```

### RBAC Configuration

```yaml
# rbac.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: dockerutils
  namespace: dockerutils

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: dockerutils
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "list"]
- apiGroups: ["metrics.k8s.io"]
  resources: ["nodes", "pods"]
  verbs: ["get", "list"]

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: dockerutils
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: dockerutils
subjects:
- kind: ServiceAccount
  name: dockerutils
  namespace: dockerutils
```

## Cloud Platform Deployment

### AWS ECS

```yaml
# ecs-task-definition.json
{
  "family": "dockerutils",
  "networkMode": "awsvpc",
  "requiresCompatibilities": ["FARGATE"],
  "cpu": "1024",
  "memory": "2048",
  "executionRoleArn": "arn:aws:iam::123456789012:role/ecsTaskExecutionRole",
  "taskRoleArn": "arn:aws:iam::123456789012:role/dockerutilsTaskRole",
  "containerDefinitions": [
    {
      "name": "dockerutils",
      "image": "dockerutils:latest",
      "portMappings": [
        {
          "containerPort": 8080,
          "protocol": "tcp"
        }
      ],
      "environment": [
        {"name": "DOCKERUTILS_PORT", "value": "8080"},
        {"name": "DOCKERUTILS_HOST", "value": "0.0.0.0"}
      ],
      "healthCheck": {
        "command": ["CMD-SHELL", "wget --quiet --tries=1 --spider http://localhost:8080/api/health || exit 1"],
        "interval": 30,
        "timeout": 5,
        "retries": 3,
        "startPeriod": 60
      },
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/dockerutils",
          "awslogs-region": "us-west-2",
          "awslogs-stream-prefix": "ecs"
        }
      }
    }
  ]
}
```

### Google Cloud Run

```yaml
# cloud-run.yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: dockerutils
  annotations:
    run.googleapis.com/ingress: all
    run.googleapis.com/execution-environment: gen2
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/maxScale: "10"
        autoscaling.knative.dev/minScale: "1"
        run.googleapis.com/execution-environment: gen2
        run.googleapis.com/cpu-throttling: "false"
    spec:
      containerConcurrency: 80
      timeoutSeconds: 300
      containers:
      - image: gcr.io/PROJECT_ID/dockerutils:latest
        ports:
        - containerPort: 8080
        env:
        - name: DOCKERUTILS_PORT
          value: "8080"
        - name: DOCKERUTILS_HOST
          value: "0.0.0.0"
        resources:
          limits:
            cpu: "2000m"
            memory: "2Gi"
          requests:
            cpu: "1000m"
            memory: "1Gi"
        livenessProbe:
          httpGet:
            path: /api/health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 30
        startupProbe:
          httpGet:
            path: /api/health
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 5
          failureThreshold: 12
```

### Azure Container Instances

```yaml
# azure-container-instance.yaml
apiVersion: 2021-09-01
location: eastus
name: dockerutils
properties:
  containers:
  - name: dockerutils
    properties:
      image: dockerutils:latest
      resources:
        requests:
          cpu: 1
          memoryInGb: 2
      ports:
      - port: 8080
        protocol: TCP
      environmentVariables:
      - name: DOCKERUTILS_PORT
        value: 8080
      - name: DOCKERUTILS_HOST
        value: 0.0.0.0
      livenessProbe:
        httpGet:
          path: /api/health
          port: 8080
        initialDelaySeconds: 30
        periodSeconds: 30
      readinessProbe:
        httpGet:
          path: /api/health
          port: 8080
        initialDelaySeconds: 10
        periodSeconds: 5
  osType: Linux
  restartPolicy: Always
  ipAddress:
    type: Public
    ports:
    - protocol: TCP
      port: 8080
    dnsNameLabel: dockerutils-demo
type: Microsoft.ContainerInstance/containerGroups
```

## Bare Metal Deployment

### System Service (systemd)

```ini
# /etc/systemd/system/dockerutils.service
[Unit]
Description=Docker Utils Analysis Service
After=network.target docker.service
Requires=docker.service

[Service]
Type=exec
User=dockerutils
Group=docker
ExecStart=/usr/local/bin/dockerutils server --port 8080 --host 0.0.0.0
ExecReload=/bin/kill -HUP $MAINPID
KillMode=mixed
Restart=always
RestartSec=5
TimeoutStopSec=30

# Security settings
NoNewPrivileges=yes
PrivateTmp=yes
ProtectSystem=strict
ProtectHome=yes
ReadWritePaths=/opt/dockerutils/data
EnvironmentFile=/etc/dockerutils/config

# Resource limits
LimitNOFILE=65536
MemoryMax=2G
CPUQuota=100%

# Logging
StandardOutput=journal
StandardError=journal
SyslogIdentifier=dockerutils

[Install]
WantedBy=multi-user.target
```

### Installation Script

```bash
#!/bin/bash
# install-dockerutils.sh

set -euo pipefail

# Configuration
DOCKERUTILS_VERSION="v1.0.0"
INSTALL_DIR="/usr/local/bin"
CONFIG_DIR="/etc/dockerutils"
DATA_DIR="/opt/dockerutils/data"
USER="dockerutils"
GROUP="docker"

# Create user and directories
sudo useradd -r -s /bin/false -d /opt/dockerutils "$USER" || true
sudo usermod -aG "$GROUP" "$USER"
sudo mkdir -p "$CONFIG_DIR" "$DATA_DIR"
sudo chown "$USER:$GROUP" "$DATA_DIR"

# Download and install binary
wget -O "/tmp/dockerutils" \
  "https://github.com/smiller333/dockerutils/releases/download/${DOCKERUTILS_VERSION}/dockerutils-linux-amd64"
sudo mv "/tmp/dockerutils" "$INSTALL_DIR/"
sudo chmod +x "${INSTALL_DIR}/dockerutils"

# Create configuration
cat << EOF | sudo tee "${CONFIG_DIR}/config"
DOCKERUTILS_PORT=8080
DOCKERUTILS_HOST=0.0.0.0
DOCKERUTILS_TMP_DIR=${DATA_DIR}/tmp
DOCKERUTILS_LOG_LEVEL=info
EOF

# Install systemd service
sudo systemctl daemon-reload
sudo systemctl enable dockerutils
sudo systemctl start dockerutils

echo "Docker Utils installed successfully!"
echo "Service status: $(sudo systemctl is-active dockerutils)"
echo "Access at: http://localhost:8080"
```

## Configuration Management

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DOCKERUTILS_PORT` | `8080` | HTTP server port |
| `DOCKERUTILS_HOST` | `localhost` | Bind address |
| `DOCKERUTILS_TMP_DIR` | `/tmp` | Temporary storage path |
| `DOCKERUTILS_LOG_LEVEL` | `info` | Logging level |
| `DOCKERUTILS_MAX_ANALYSES` | `100` | Maximum concurrent analyses |
| `DOCKERUTILS_CLEANUP_INTERVAL` | `1h` | Cleanup interval |
| `DOCKERUTILS_MAX_STORAGE` | `10GB` | Maximum storage usage |

### Configuration File

```yaml
# /etc/dockerutils/config.yaml
server:
  port: 8080
  host: "0.0.0.0"
  read_timeout: "30s"
  write_timeout: "30s"
  idle_timeout: "60s"

storage:
  data_dir: "/opt/dockerutils/data"
  tmp_dir: "/opt/dockerutils/data/tmp"
  max_size: "10GB"
  cleanup_interval: "1h"

analysis:
  max_concurrent: 5
  timeout: "10m"
  memory_limit: "1GB"

logging:
  level: "info"
  format: "json"
  output: "stdout"

security:
  trusted_registries:
    - "docker.io"
    - "gcr.io"
    - "quay.io"
  max_image_size: "5GB"
  enable_security_scan: false
```

## Security Considerations

### Docker Socket Security

⚠️ **CRITICAL**: Docker socket access grants significant system privileges.

```bash
# Minimize socket permissions
sudo chmod 660 /var/run/docker.sock

# Use Docker socket proxy (recommended)
docker run -d \
  --name docker-socket-proxy \
  --restart unless-stopped \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -p 127.0.0.1:2375:2375 \
  -e CONTAINERS=1 \
  -e IMAGES=1 \
  -e BUILD=1 \
  tecnativa/docker-socket-proxy
```

### Network Security

```yaml
# nginx.conf - Reverse proxy configuration
upstream dockerutils {
    server 127.0.0.1:8080;
}

server {
    listen 80;
    server_name dockerutils.example.com;
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name dockerutils.example.com;

    ssl_certificate /etc/ssl/certs/dockerutils.crt;
    ssl_certificate_key /etc/ssl/private/dockerutils.key;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES256-GCM-SHA512:DHE-RSA-AES256-GCM-SHA512;

    # Security headers
    add_header X-Frame-Options DENY;
    add_header X-Content-Type-Options nosniff;
    add_header X-XSS-Protection "1; mode=block";
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains";

    # Rate limiting
    limit_req_zone $binary_remote_addr zone=dockerutils:10m rate=10r/s;
    limit_req zone=dockerutils burst=20 nodelay;

    # File size limits
    client_max_body_size 50M;

    location / {
        proxy_pass http://dockerutils;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        # Timeouts
        proxy_connect_timeout 5s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }
}
```

### Container Security

```dockerfile
# Security-hardened Dockerfile
FROM scratch

# Use non-root user
USER 65534:65534

# Read-only filesystem
VOLUME ["/tmp"]

# Drop all capabilities
# Add only required capabilities via docker run --cap-add

# Security labels
LABEL security.scan="enabled" \
      security.updates="auto" \
      security.contact="security@example.com"
```

## Monitoring and Logging

### Prometheus Metrics

```yaml
# prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
- job_name: 'dockerutils'
  static_configs:
  - targets: ['dockerutils:8080']
  metrics_path: /metrics
  scrape_interval: 30s
```

### Grafana Dashboard

```json
{
  "dashboard": {
    "title": "Docker Utils Monitoring",
    "panels": [
      {
        "title": "Request Rate",
        "type": "graph",
        "targets": [
          {
            "expr": "rate(http_requests_total[5m])"
          }
        ]
      },
      {
        "title": "Response Time",
        "type": "graph", 
        "targets": [
          {
            "expr": "histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))"
          }
        ]
      },
      {
        "title": "Active Analyses",
        "type": "singlestat",
        "targets": [
          {
            "expr": "dockerutils_active_analyses"
          }
        ]
      }
    ]
  }
}
```

### Log Management

```yaml
# docker-compose with logging
version: '3.8'
services:
  dockerutils:
    image: dockerutils:latest
    logging:
      driver: "json-file"
      options:
        max-size: "100m"
        max-file: "5"
        labels: "service,environment"
    labels:
      - "service=dockerutils"
      - "environment=production"

  # ELK Stack for log aggregation
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:8.11.0
    environment:
      - discovery.type=single-node
      - xpack.security.enabled=false
    volumes:
      - esdata:/usr/share/elasticsearch/data

  logstash:
    image: docker.elastic.co/logstash/logstash:8.11.0
    volumes:
      - ./logstash.conf:/usr/share/logstash/pipeline/logstash.conf

  kibana:
    image: docker.elastic.co/kibana/kibana:8.11.0
    ports:
      - "5601:5601"

volumes:
  esdata:
```

## High Availability

### Load Balancer Configuration

```yaml
# haproxy.cfg
global
    daemon
    maxconn 4096

defaults
    mode http
    timeout connect 5000ms
    timeout client 50000ms
    timeout server 50000ms
    option httplog

frontend dockerutils_frontend
    bind *:80
    bind *:443 ssl crt /etc/ssl/certs/dockerutils.pem
    redirect scheme https if !{ ssl_fc }
    default_backend dockerutils_backend

backend dockerutils_backend
    balance roundrobin
    option httpchk GET /api/health
    server web1 10.0.1.10:8080 check
    server web2 10.0.1.11:8080 check
    server web3 10.0.1.12:8080 check
```

### Database Clustering (if using external storage)

```yaml
# docker-compose.ha.yml
version: '3.8'
services:
  postgres-primary:
    image: postgres:15
    environment:
      POSTGRES_DB: dockerutils
      POSTGRES_USER: dockerutils
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_REPLICATION_USER: replicator
      POSTGRES_REPLICATION_PASSWORD: ${REPL_PASSWORD}
    volumes:
      - postgres_primary:/var/lib/postgresql/data
      - ./postgres-primary.conf:/etc/postgresql/postgresql.conf
    command: postgres -c config_file=/etc/postgresql/postgresql.conf

  postgres-replica:
    image: postgres:15
    environment:
      POSTGRES_USER: dockerutils
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      PGUSER: postgres
    volumes:
      - postgres_replica:/var/lib/postgresql/data
    depends_on:
      - postgres-primary
    command: |
      bash -c "
      until pg_basebackup -h postgres-primary -D /var/lib/postgresql/data -U replicator -v -P -W; do
        echo 'Waiting for primary to connect...'
        sleep 1s
      done
      echo 'standby_mode = on' >> /var/lib/postgresql/data/recovery.conf
      echo 'primary_conninfo = \"host=postgres-primary port=5432 user=replicator\"' >> /var/lib/postgresql/data/recovery.conf
      postgres"
```

## Troubleshooting

### Common Issues

#### 1. Docker Socket Permission Denied

```bash
# Check socket permissions
ls -la /var/run/docker.sock

# Fix permissions
sudo chmod 666 /var/run/docker.sock

# Add user to docker group (preferred)
sudo usermod -aG docker $USER
newgrp docker
```

#### 2. Out of Storage Space

```bash
# Check disk usage
df -h /opt/dockerutils/data

# Clean up old analyses
docker exec dockerutils-container /app/dockerutils cleanup --older-than 7d

# Configure automatic cleanup
echo "DOCKERUTILS_CLEANUP_INTERVAL=1h" >> /etc/dockerutils/config
```

#### 3. Memory Issues

```bash
# Check memory usage
docker stats dockerutils-container

# Increase memory limits
docker update --memory 2g dockerutils-container

# Or in docker-compose
version: '3.8'
services:
  dockerutils:
    deploy:
      resources:
        limits:
          memory: 2G
```

#### 4. Network Connectivity

```bash
# Test Docker daemon connectivity
docker ps

# Test container networking
docker exec dockerutils-container wget -qO- http://localhost:8080/api/health

# Check firewall rules
sudo ufw status
sudo iptables -L
```

### Health Checks

```bash
#!/bin/bash
# health-check.sh

# API health check
if ! curl -f http://localhost:8080/api/health; then
    echo "ERROR: API health check failed"
    exit 1
fi

# Docker connectivity check
if ! docker ps >/dev/null 2>&1; then
    echo "ERROR: Docker daemon not accessible"
    exit 1
fi

# Disk space check
USAGE=$(df /opt/dockerutils/data | awk 'NR==2 {print $5}' | sed 's/%//')
if [ "$USAGE" -gt 90 ]; then
    echo "ERROR: Disk usage above 90%"
    exit 1
fi

echo "All checks passed"
```

### Log Analysis

```bash
# View container logs
docker logs dockerutils-container

# Follow logs in real-time
docker logs -f dockerutils-container

# Filter error logs
docker logs dockerutils-container 2>&1 | grep ERROR

# System service logs
sudo journalctl -u dockerutils -f

# Analyze performance
docker stats dockerutils-container
```

---

**Related Documentation:**
- [Security Policy](SECURITY.md) - Security guidelines and best practices
- [Architecture Overview](ARCHITECTURE.md) - System architecture details
- [Troubleshooting Guide](../TROUBLESHOOTING.md) - Detailed troubleshooting steps

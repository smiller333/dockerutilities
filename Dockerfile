# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dockerutilities .

# Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates wget

# Create docker group with common GID and non-root user
# Handle case where docker group might already exist
RUN (addgroup -g 999 -S docker 2>/dev/null || addgroup -S docker) && \
    addgroup -g 1001 -S dockerutilities && \
adduser -u 1001 -S dockerutilities -G dockerutilities && \
adduser dockerutilities docker

# Create directories and handle potential Docker socket group
RUN mkdir -p /app/data && \
    chown -R dockerutilities:dockerutilities /app

# Create an entrypoint script to handle dynamic Docker socket permissions
COPY <<'EOF' /entrypoint.sh
#!/bin/sh
# Handle Docker socket permissions dynamically
if [ -S /var/run/docker.sock ]; then
    DOCKER_SOCKET_GID=$(stat -c '%g' /var/run/docker.sock)
    echo "Docker socket GID: $DOCKER_SOCKET_GID"
    
    # Check if group exists, if not create it
    if ! getent group "$DOCKER_SOCKET_GID" >/dev/null 2>&1; then
        addgroup -g "$DOCKER_SOCKET_GID" -S dockerhost
        echo "Created group dockerhost with GID $DOCKER_SOCKET_GID"
    fi
    
    # Get the group name for this GID
    GROUP_NAME=$(getent group "$DOCKER_SOCKET_GID" | cut -d: -f1)
    echo "Adding dockerutilities user to group: $GROUP_NAME"
    
    # Add dockerutilities user to the docker socket group
adduser dockerutilities "$GROUP_NAME"
fi

# Install su-exec if not already available
if ! command -v su-exec >/dev/null 2>&1; then
    apk add --no-cache su-exec
fi

# Switch to dockerutilities user and execute the original command
exec su-exec dockerutilities "$@"
EOF

RUN chmod +x /entrypoint.sh

# Install su-exec for privilege dropping
RUN apk add --no-cache su-exec

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/dockerutilities .

# Change ownership of the binary
RUN chown dockerutilities:dockerutilities dockerutilities

# Create volume mount point for persistent data
VOLUME ["/app/data"]

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/health || exit 1

# Command to run the application
ENTRYPOINT ["/entrypoint.sh"]
CMD ["./dockerutilities", "server", "--host", "0.0.0.0", "--port", "8080", "--tmp-dir", "/app/data"]

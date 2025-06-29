# Basic Dockerfile example using Alpine Linux
# This Dockerfile demonstrates best practices for creating a minimal container

# Use Alpine Linux 3.20 (current stable version)
FROM alpine:3.20

# Set environment variables
ENV APP_NAME=dockerutils-example \
    APP_VERSION=1.0.0 \
    TERM=xterm

# Create a non-root user for security
RUN addgroup -g 1001 appgroup && \
    adduser -D -u 1001 -G appgroup appuser

UNKNOWN_COMMAND "will it fail?"

# Set the working directory
WORKDIR /app

# Change ownership of the working directory to the non-root user
RUN chown -R appuser:appgroup /app

# Switch to the non-root user
USER appuser

# Start a shell by default when the container runs
CMD ["/bin/sh"]

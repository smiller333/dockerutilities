#!/bin/bash

# Build and run script for dockerutilities in Docker
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
IMAGE_NAME="dockerutilities"
CONTAINER_NAME="dockerutilities"
PORT="8080"
DATA_DIR="./data"

# Docker socket configuration - auto-detect or use environment variable/command line
detect_docker_socket() {
    # If DOCKER_SOCKET is already set via environment variable, use it
    if [ -n "$DOCKER_SOCKET" ]; then
        return
    fi
    
    # Try to get the current Docker context endpoint
    local context_endpoint=$(docker context inspect --format '{{.Endpoints.docker.Host}}' 2>/dev/null || echo "")
    
    if [[ "$context_endpoint" == unix://* ]]; then
        # Extract socket path from unix:// URL
        DOCKER_SOCKET="${context_endpoint#unix://}"
        log_info "Auto-detected Docker socket from context: $DOCKER_SOCKET"
    else
        # Fall back to common locations in order of preference
        local common_sockets=(
            "/var/run/docker.sock"                    # Linux default
            "$HOME/.docker/desktop/docker.sock"      # Docker Desktop macOS
            "/run/user/$(id -u)/podman/podman.sock"  # Podman
        )
        
        for socket in "${common_sockets[@]}"; do
            if [ -e "$socket" ]; then
                DOCKER_SOCKET="$socket"
                log_info "Auto-detected Docker socket: $DOCKER_SOCKET"
                return
            fi
        done
        
        # Default fallback
        DOCKER_SOCKET="/var/run/docker.sock"
    fi
}

# Initialize socket detection later after functions are defined

# Functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Initialize Docker socket detection now that functions are defined
detect_docker_socket

# Parse command line arguments
COMMAND=${1:-"run"}

# Parse additional arguments
shift || true  # Remove first argument (command)
PERSISTENT_DATA="false"

while [[ $# -gt 0 ]]; do
    case $1 in
        --persistent|--persist|-p)
            PERSISTENT_DATA="true"
            shift
            ;;
        --socket|-s)
            DOCKER_SOCKET="$2"
            shift 2
            ;;
        --port)
            PORT="$2"
            shift 2
            ;;
        --data-dir)
            DATA_DIR="$2"
            shift 2
            ;;
        --help|-h)
            show_help
            exit 0
            ;;
        *)
            # Handle legacy second argument for persistent data
            if [[ "$1" == "true" ]] && [[ "$PERSISTENT_DATA" == "false" ]]; then
                PERSISTENT_DATA="true"
            fi
            shift
            ;;
    esac
done

show_help() {
    echo "Usage: $0 [COMMAND] [OPTIONS]"
    echo ""
    echo "Commands:"
    echo "  build          Build the Docker image"
    echo "  run            Build and run the container (default)"
    echo "  run-persistent Build and run with persistent data storage"
    echo "  stop           Stop the running container"
    echo "  clean          Stop and remove container and image"
    echo "  logs           Show container logs"
    echo "  shell          Open shell in running container"
    echo "  help           Show this help message"
    echo ""
    echo "Options:"
    echo "  --persistent, -p           Enable persistent data storage"
    echo "  --socket PATH, -s PATH     Docker socket path (default: /var/run/docker.sock)"
    echo "  --port PORT                Host port to bind to (default: 8080)"
    echo "  --data-dir DIR            Data directory for persistent storage (default: ./data)"
    echo "  --help, -h                Show this help message"
    echo ""
    echo "Environment Variables:"
    echo "  DOCKER_SOCKET             Docker socket path (overridden by --socket)"
    echo ""
    echo "Examples:"
    echo "  $0 build                                    # Build the image only"
    echo "  $0 run                                      # Build and run with ephemeral storage"
    echo "  $0 run --persistent                         # Build and run with persistent storage"
    echo "  $0 run -s /custom/docker.sock              # Use custom Docker socket"
    echo "  $0 run --socket ~/.docker/desktop/docker.sock --persistent  # Docker Desktop on macOS"
    echo "  $0 stop                                     # Stop the container"
    echo "  $0 clean                                    # Clean up everything"
    echo ""
    echo "Common Docker Socket Locations:"
    echo "  Linux (default):           /var/run/docker.sock"
    echo "  Docker Desktop (macOS):    ~/.docker/desktop/docker.sock"
    echo "  Docker Desktop (Windows):  //./pipe/docker_engine"
    echo "  Podman:                    /run/user/$(id -u)/podman/podman.sock"
    echo ""
    echo "Note: The script automatically detects the Docker socket group ID and"
    echo "      grants the container user access to the socket for proper permissions."
}

build_image() {
    log_info "Building Docker image: $IMAGE_NAME"
    docker build -t $IMAGE_NAME .
    log_success "Image built successfully"
}

create_data_dir() {
    if [ ! -d "$DATA_DIR" ]; then
        log_info "Creating data directory: $DATA_DIR"
        mkdir -p "$DATA_DIR"
        log_success "Data directory created"
    fi
}

validate_docker_socket() {
    if [ ! -e "$DOCKER_SOCKET" ]; then
        log_error "Docker socket not found at: $DOCKER_SOCKET"
        log_info "Common Docker socket locations:"
        log_info "  Linux (default):           /var/run/docker.sock"
        log_info "  Docker Desktop (macOS):    ~/.docker/desktop/docker.sock"
        log_info "  Docker Desktop (Windows):  //./pipe/docker_engine"
        log_info "  Podman:                    /run/user/$(id -u)/podman/podman.sock"
        log_info ""
        log_info "Use --socket option to specify a custom location:"
        log_info "  $0 run --socket /path/to/docker.sock"
        exit 1
    fi
    
    if [ ! -r "$DOCKER_SOCKET" ]; then
        log_error "Docker socket is not readable: $DOCKER_SOCKET"
        log_info "Check permissions and ensure Docker daemon is running"
        exit 1
    fi
    
    log_info "Using Docker socket: $DOCKER_SOCKET"
}

run_container() {
    local persistent=$1
    
    # Validate Docker socket before starting container
    validate_docker_socket
    
    # Stop existing container if running
    if docker ps -q -f name=$CONTAINER_NAME | grep -q .; then
        log_warning "Stopping existing container"
        docker stop $CONTAINER_NAME > /dev/null 2>&1
    fi
    
    # Remove existing container if exists
    if docker ps -aq -f name=$CONTAINER_NAME | grep -q .; then
        log_info "Removing existing container"
        docker rm $CONTAINER_NAME > /dev/null 2>&1
    fi
    
    log_info "Starting container: $CONTAINER_NAME"
    log_info "Port mapping: $PORT:8080"
    log_info "Docker socket permissions will be handled automatically"
    
    if [ "$persistent" = "true" ]; then
        create_data_dir
        log_info "Running with persistent data storage"
        log_info "Data directory: $(pwd)/$DATA_DIR"
        docker run -d \
            --name $CONTAINER_NAME \
            -p $PORT:8080 \
            -v "$DOCKER_SOCKET:/var/run/docker.sock" \
            -v "$(pwd)/$DATA_DIR:/app/data" \
            $IMAGE_NAME
    else
        log_info "Running with ephemeral storage"
        docker run -d \
            --name $CONTAINER_NAME \
            -p $PORT:8080 \
            -v "$DOCKER_SOCKET:/var/run/docker.sock" \
            $IMAGE_NAME
    fi
    
    log_success "Container started successfully"
    log_info "Web interface available at: http://localhost:$PORT"
    log_info "API health check: http://localhost:$PORT/api/health"
}

stop_container() {
    if docker ps -q -f name=$CONTAINER_NAME | grep -q .; then
        log_info "Stopping container: $CONTAINER_NAME"
        docker stop $CONTAINER_NAME
        log_success "Container stopped"
    else
        log_warning "Container is not running"
    fi
}

clean_up() {
    log_info "Cleaning up Docker resources"
    
    # Stop container if running
    if docker ps -q -f name=$CONTAINER_NAME | grep -q .; then
        docker stop $CONTAINER_NAME > /dev/null 2>&1
    fi
    
    # Remove container if exists
    if docker ps -aq -f name=$CONTAINER_NAME | grep -q .; then
        docker rm $CONTAINER_NAME > /dev/null 2>&1
        log_success "Container removed"
    fi
    
    # Remove image if exists
    if docker images -q $IMAGE_NAME | grep -q .; then
        docker rmi $IMAGE_NAME > /dev/null 2>&1
        log_success "Image removed"
    fi
    
    log_success "Cleanup completed"
}

show_logs() {
    if docker ps -q -f name=$CONTAINER_NAME | grep -q .; then
        log_info "Showing logs for container: $CONTAINER_NAME"
        docker logs -f $CONTAINER_NAME
    else
        log_error "Container is not running"
        exit 1
    fi
}

open_shell() {
    if docker ps -q -f name=$CONTAINER_NAME | grep -q .; then
        log_info "Opening shell in container: $CONTAINER_NAME"
        docker exec -it $CONTAINER_NAME /bin/sh
    else
        log_error "Container is not running"
        exit 1
    fi
}

# Main execution
case $COMMAND in
    "build")
        build_image
        ;;
    "run")
        build_image
        run_container "$PERSISTENT_DATA"
        ;;
    "run-persistent")
        # Legacy command - set persistent to true
        PERSISTENT_DATA="true"
        build_image
        run_container "$PERSISTENT_DATA"
        ;;
    "stop")
        stop_container
        ;;
    "clean")
        clean_up
        ;;
    "logs")
        show_logs
        ;;
    "shell")
        open_shell
        ;;
    "help"|"--help"|"-h")
        show_help
        ;;
    *)
        log_error "Unknown command: $COMMAND"
        show_help
        exit 1
        ;;
esac

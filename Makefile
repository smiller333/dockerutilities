# Makefile for dockerutils - Go-based Docker command-line utilities
# This Makefile builds the dockerutils binary with embedded build-time information

.PHONY: build build-dev clean test version help install deps lint fmt check dev watch

# Project information
PROJECT_NAME := dockerutils
MODULE_NAME := github.com/smiller333/dockerutils
CMD_PATH := ./cmd/dockerutils
MAIN_PATH := .

# Build output directory
BUILD_DIR := bin
BINARY_NAME := $(PROJECT_NAME)
BINARY_PATH := $(BUILD_DIR)/$(BINARY_NAME)

# Version information - these will be injected at build time
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GO_VERSION := $(shell go version | awk '{print $$3}')

# Go build flags for embedding version information
VERSION_PKG := $(MODULE_NAME)/src/version
LDFLAGS := -X $(VERSION_PKG).Version=$(VERSION)
LDFLAGS += -X $(VERSION_PKG).GitCommit=$(GIT_COMMIT)
LDFLAGS += -X $(VERSION_PKG).BuildTime=$(BUILD_TIME)

# Go build options
GO_BUILD_FLAGS := -ldflags "$(LDFLAGS)"
GO_BUILD_FLAGS_RELEASE := -ldflags "$(LDFLAGS) -s -w" -trimpath

# Default target
all: clean build

# Development build (faster, includes debug info)
build: deps
	@echo "Building $(BINARY_NAME) (development)..."
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)" 
	@echo "Build Time: $(BUILD_TIME)"
	@mkdir -p $(BUILD_DIR)
	go build $(GO_BUILD_FLAGS) -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Binary created at: $(BINARY_PATH)"

# Development build with more verbose output
build-dev: deps
	@echo "Building $(BINARY_NAME) (development with debug)..."
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Go Version: $(GO_VERSION)"
	@mkdir -p $(BUILD_DIR)
	go build -v $(GO_BUILD_FLAGS) -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Binary created at: $(BINARY_PATH)"
	@echo "Running version check:"
	@$(BINARY_PATH) version 2>/dev/null || echo "Note: version command may not be implemented yet"

# Production build (optimized, stripped)
build-release: deps
	@echo "Building $(BINARY_NAME) (release)..."
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@mkdir -p $(BUILD_DIR)
	go build $(GO_BUILD_FLAGS_RELEASE) -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Optimized binary created at: $(BINARY_PATH)"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Run tests
test: deps
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage: deps
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	@echo "Clean complete"

# Display version information
version:
	@echo "Project: $(PROJECT_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Go Version: $(GO_VERSION)"
	@echo "Module: $(MODULE_NAME)"

# Install the binary to GOPATH/bin or GOBIN
install: build-release
	@echo "Installing $(BINARY_NAME)..."
	go install $(GO_BUILD_FLAGS_RELEASE) $(MAIN_PATH)
	@echo "$(BINARY_NAME) installed successfully"

# Development helpers
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

lint: deps
	@echo "Running golangci-lint..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi

# Check code quality (format, lint, test)
check: fmt lint test
	@echo "All checks passed!"

# Development mode - build and run with common flags
dev: build
	@echo "Running $(BINARY_NAME) in development mode..."
	@echo "Use 'make dev ARGS=\"your-args-here\"' to pass arguments"
	$(BINARY_PATH) $(ARGS)

# Watch mode for development (requires fswatch or inotify-tools)
watch:
	@echo "Starting watch mode (requires fswatch)..."
	@if command -v fswatch >/dev/null 2>&1; then \
		fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make build; \
	else \
		echo "fswatch not found. Install with: brew install fswatch (macOS) or apt-get install fswatch (Linux)"; \
		exit 1; \
	fi

# Docker targets (if you want to containerize your application)
docker-build:
	@echo "Building Docker image..."
	docker build -t $(PROJECT_NAME):$(VERSION) -t $(PROJECT_NAME):latest .

docker-run: docker-build
	@echo "Running Docker container..."
	docker run --rm -it $(PROJECT_NAME):latest $(ARGS)

# Cross-compilation targets
build-linux: deps
	@echo "Building for Linux amd64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(GO_BUILD_FLAGS_RELEASE) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)

build-windows: deps
	@echo "Building for Windows amd64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build $(GO_BUILD_FLAGS_RELEASE) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

build-darwin: deps
	@echo "Building for macOS amd64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build $(GO_BUILD_FLAGS_RELEASE) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)

build-darwin-arm64: deps
	@echo "Building for macOS arm64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 go build $(GO_BUILD_FLAGS_RELEASE) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)

# Build all cross-platform binaries
build-all: build-linux build-windows build-darwin build-darwin-arm64
	@echo "All platform binaries built in $(BUILD_DIR)/"

# Help target
help:
	@echo "Available targets:"
	@echo "  build          - Build development binary with debug info"
	@echo "  build-dev      - Build development binary with verbose output"
	@echo "  build-release  - Build optimized release binary"
	@echo "  build-all      - Build binaries for all supported platforms"
	@echo "  clean          - Remove build artifacts"
	@echo "  test           - Run all tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  deps           - Install/update dependencies"
	@echo "  install        - Install binary to GOPATH/bin"
	@echo "  fmt            - Format Go code"
	@echo "  lint           - Run linter (requires golangci-lint)"
	@echo "  check          - Run fmt, lint, and test"
	@echo "  dev            - Build and run with ARGS"
	@echo "  watch          - Watch for changes and rebuild (requires fswatch)"
	@echo "  version        - Display version information"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Build and run Docker container"
	@echo "  help           - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make build                    # Build development binary"
	@echo "  make build-release           # Build optimized binary"
	@echo "  make dev ARGS=\"--help\"       # Build and run with --help"
	@echo "  make test                    # Run tests"
	@echo "  make VERSION=v1.2.3 build    # Build with specific version"

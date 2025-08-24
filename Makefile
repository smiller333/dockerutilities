# Makefile for dockerutilities - Go-based Docker command-line utilities
# This Makefile builds the dockerutilities binary with embedded build-time information
# Simplified for local development - GoReleaser handles cross-platform builds

.PHONY: dev release local clean test version help install deps lint fmt check

# Project information
PROJECT_NAME := dockerutilities
MODULE_NAME := github.com/smiller333/dockerutilities
MAIN_PATH := .

# Build output directory
BUILD_DIR := bin
BINARY_NAME := $(PROJECT_NAME)
BINARY_PATH := $(BUILD_DIR)/$(BINARY_NAME)

# Version information - these will be injected at build time
# Uses git tags for version detection with support for development tags
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GO_VERSION := $(shell go version | awk '{print $$3}')

# Go build flags for embedding version information
# Compatible with GoReleaser template variables for future integration
VERSION_PKG := $(MODULE_NAME)/src/version
LDFLAGS := -X $(VERSION_PKG).Version=$(VERSION)
LDFLAGS += -X $(VERSION_PKG).GitCommit=$(GIT_COMMIT)
LDFLAGS += -X $(VERSION_PKG).BuildTime=$(BUILD_TIME)

# Go build options
GO_BUILD_FLAGS := -ldflags "$(LDFLAGS)"
GO_BUILD_FLAGS_RELEASE := -ldflags "$(LDFLAGS) -s -w" -trimpath

# Default target
all: clean dev

# Development build (faster, includes debug info)
dev: deps
	@echo "Building $(BINARY_NAME) (development)..."
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)" 
	@echo "Build Time: $(BUILD_TIME)"
	@mkdir -p $(BUILD_DIR)
	go build $(GO_BUILD_FLAGS) -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Binary created at: $(BINARY_PATH)"

# Production build (optimized, stripped)
release: deps
	@echo "Building $(BINARY_NAME) (release)..."
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@mkdir -p $(BUILD_DIR)
	go build $(GO_BUILD_FLAGS_RELEASE) -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Optimized binary created at: $(BINARY_PATH)"

# Local development build with verbose output
local: deps
	@echo "Building $(BINARY_NAME) (local development)..."
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Go Version: $(GO_VERSION)"
	@mkdir -p $(BUILD_DIR)
	go build -v $(GO_BUILD_FLAGS) -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Binary created at: $(BINARY_PATH)"
	@echo "Running version check:"
	@$(BINARY_PATH) version 2>/dev/null || echo "Note: version command may not be implemented yet"

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





# Docker targets (if you want to containerize your application)
docker-build:
	@echo "Building Docker image..."
	docker build -t $(PROJECT_NAME):$(VERSION) -t $(PROJECT_NAME):latest .

docker-run: docker-build
	@echo "Running Docker container..."
	docker run --rm -it $(PROJECT_NAME):latest $(ARGS)

# Note: Cross-platform builds are handled by GoReleaser
# Use 'goreleaser release --snapshot' for local cross-platform testing

# Help target
help:
	@echo "Available targets:"
	@echo "  dev            - Build development binary (default)"
	@echo "  release        - Build optimized release binary"
	@echo "  local          - Build with verbose output for local development"
	@echo "  clean          - Remove build artifacts"
	@echo "  test           - Run all tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  deps           - Install/update dependencies"
	@echo "  install        - Install binary to GOPATH/bin"
	@echo "  fmt            - Format Go code"
	@echo "  lint           - Run linter (requires golangci-lint)"
	@echo "  check          - Run fmt, lint, and test"
	@echo "  version        - Display version information"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Build and run Docker container"
	@echo "  help           - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make dev                      # Build development binary"
	@echo "  make release                  # Build optimized binary"
	@echo "  make local                    # Build with verbose output"
	@echo "  make test                     # Run tests"
	@echo "  make VERSION=v1.2.3 dev       # Build with specific version"
	@echo ""
	@echo "Note: Cross-platform builds are handled by GoReleaser"
	@echo "Use 'goreleaser release --snapshot' for local cross-platform testing"

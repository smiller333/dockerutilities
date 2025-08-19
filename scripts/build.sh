#!/bin/bash

# build.sh - Simple build script for dockerutilities with version injection
# This script wraps the Makefile for easier usage

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print colored output
print_info() {
    echo -e "${BLUE}ℹ ${1}${NC}"
}

print_success() {
    echo -e "${GREEN}✓ ${1}${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠ ${1}${NC}"
}

print_error() {
    echo -e "${RED}✗ ${1}${NC}"
}

# Show usage
show_usage() {
    echo "Usage: $0 [OPTIONS] [TARGET]"
    echo ""
    echo "Build dockerutilities with embedded version information"
    echo ""
    echo "TARGETS:"
    echo "  dev          Build development version (default)"
    echo "  release      Build optimized release version"
    echo "  all          Build all platform binaries"
    echo "  test         Run tests"
    echo "  clean        Clean build artifacts"
    echo "  version      Show version information"
    echo ""
    echo "OPTIONS:"
    echo "  -v VERSION   Override version (auto-detected from git by default)"
    echo "  -h, --help   Show this help message"
    echo ""
    echo "EXAMPLES:"
    echo "  $0                    # Build development version"
    echo "  $0 release           # Build release version"
    echo "  $0 -v v1.2.3 release # Build release with specific version"
    echo "  $0 test              # Run tests"
    echo ""
}

# Parse command line arguments
VERSION=""
TARGET="dev"

while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--version)
            VERSION="$2"
            shift 2
            ;;
        -h|--help)
            show_usage
            exit 0
            ;;
        dev|release|all|test|clean|version)
            TARGET="$1"
            shift
            ;;
        *)
            print_error "Unknown option: $1"
            show_usage
            exit 1
            ;;
    esac
done

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    print_warning "Not in a git repository. Version detection may be limited."
fi

# Detect version if not provided
if [ -z "$VERSION" ]; then
    if command -v git >/dev/null 2>&1; then
        VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
        print_info "Auto-detected version: $VERSION"
    else
        VERSION="dev"
        print_warning "Git not found. Using default version: $VERSION"
    fi
fi

# Check if make is available
if ! command -v make >/dev/null 2>&1; then
    print_error "make command not found. Please install make to use this script."
    exit 1
fi

# Check if go is available
if ! command -v go >/dev/null 2>&1; then
    print_error "go command not found. Please install Go to build this project."
    exit 1
fi

# Execute the target
print_info "Building dockerutilities..."
print_info "Target: $TARGET"

case $TARGET in
    dev)
        print_info "Building development version..."
        if [ -n "$VERSION" ]; then
            make build-dev VERSION="$VERSION"
        else
            make build-dev
        fi
        print_success "Development build completed!"
        ;;
    release)
        print_info "Building release version..."
        if [ -n "$VERSION" ]; then
            make build-release VERSION="$VERSION"
        else
            make build-release
        fi
        print_success "Release build completed!"
        ;;
    all)
        print_info "Building all platform binaries..."
        if [ -n "$VERSION" ]; then
            make build-all VERSION="$VERSION"
        else
            make build-all
        fi
        print_success "All platform builds completed!"
        ;;
    test)
        print_info "Running tests..."
        make test
        print_success "Tests completed!"
        ;;
    clean)
        print_info "Cleaning build artifacts..."
        make clean
        print_success "Clean completed!"
        ;;
    version)
        print_info "Version information:"
        make version
        ;;
    *)
        print_error "Unknown target: $TARGET"
        show_usage
        exit 1
        ;;
esac

# Show binary information if build was successful
if [[ "$TARGET" == "dev" || "$TARGET" == "release" ]]; then
    	if [ -f "bin/dockerutilities" ]; then
		print_info "Binary information:"
		ls -lh bin/dockerutilities
		
		print_info "Testing version output:"
		./bin/dockerutilities version 2>/dev/null || echo "Note: version command may not be fully implemented yet"
	fi
fi

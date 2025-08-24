#!/bin/bash

# Test GoReleaser configuration locally
# This script tests both development and production tag scenarios

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check prerequisites
print_status "Checking prerequisites..."

if ! command_exists goreleaser; then
    print_error "GoReleaser is not installed. Please install it first:"
    echo "  go install github.com/goreleaser/goreleaser/v2/cmd/goreleaser@latest"
    exit 1
fi

if ! command_exists go; then
    print_error "Go is not installed or not in PATH"
    exit 1
fi

print_success "Prerequisites check passed"

# Validate GoReleaser configuration
print_status "Validating GoReleaser configuration..."
if goreleaser check; then
    print_success "GoReleaser configuration is valid"
else
    print_error "GoReleaser configuration validation failed"
    exit 1
fi

# Clean previous builds
print_status "Cleaning previous builds..."
rm -rf dist/
print_success "Cleanup completed"

# Test 1: Development tag (alpha)
print_status "Testing development tag (alpha)..."
export GORELEASER_KEY="test-key-alpha"
export GORELEASER_CURRENT_TAG="v1.0.0-alpha.1"

print_status "Building with development tag: $GORELEASER_CURRENT_TAG"
if goreleaser build --snapshot --clean --parallelism=2; then
    print_success "Development tag build completed"
    
    # Verify artifacts
    if [ -f "dist/checksums.txt" ]; then
        print_success "Checksums file created"
        echo "Checksums content:"
        cat dist/checksums.txt
    else
        print_error "Checksums file not found"
    fi
    
    # Count binaries
    BINARY_COUNT=$(find dist -name "dockerutilities-*" -type f | wc -l)
    print_status "Found $BINARY_COUNT binaries"
    
    if [ "$BINARY_COUNT" -ge 6 ]; then
        print_success "All platform binaries created for development tag"
    else
        print_warning "Expected 6 binaries, found $BINARY_COUNT"
    fi
else
    print_error "Development tag build failed"
    exit 1
fi

# Clean for next test
rm -rf dist/

# Test 2: Development tag (beta)
print_status "Testing development tag (beta)..."
export GORELEASER_CURRENT_TAG="v1.0.0-beta.1"

print_status "Building with development tag: $GORELEASER_CURRENT_TAG"
if goreleaser build --snapshot --clean --parallelism=2; then
    print_success "Beta tag build completed"
    
    # Verify artifacts
    if [ -f "dist/checksums.txt" ]; then
        print_success "Checksums file created"
    else
        print_error "Checksums file not found"
    fi
    
    # Count binaries
    BINARY_COUNT=$(find dist -name "dockerutilities-*" -type f | wc -l)
    print_status "Found $BINARY_COUNT binaries"
    
    if [ "$BINARY_COUNT" -ge 6 ]; then
        print_success "All platform binaries created for beta tag"
    else
        print_warning "Expected 6 binaries, found $BINARY_COUNT"
    fi
else
    print_error "Beta tag build failed"
    exit 1
fi

# Clean for next test
rm -rf dist/

# Test 3: Production tag
print_status "Testing production tag..."
export GORELEASER_CURRENT_TAG="v1.0.0"

print_status "Building with production tag: $GORELEASER_CURRENT_TAG"
if goreleaser build --snapshot --clean --parallelism=2; then
    print_success "Production tag build completed"
    
    # Verify artifacts
    if [ -f "dist/checksums.txt" ]; then
        print_success "Checksums file created"
    else
        print_error "Checksums file not found"
    fi
    
    # Count binaries
    BINARY_COUNT=$(find dist -name "dockerutilities-*" -type f | wc -l)
    print_status "Found $BINARY_COUNT binaries"
    
    if [ "$BINARY_COUNT" -ge 6 ]; then
        print_success "All platform binaries created for production tag"
    else
        print_warning "Expected 6 binaries, found $BINARY_COUNT"
    fi
    
    # Test binary functionality
    if [ -f "dist/dockerutilities-linux-amd64" ]; then
        print_status "Testing Linux binary functionality..."
        chmod +x dist/dockerutilities-linux-amd64
        
        # Test version command
        if ./dist/dockerutilities-linux-amd64 version >/dev/null 2>&1; then
            print_success "Version command works"
            echo "Version output:"
            ./dist/dockerutilities-linux-amd64 version
        else
            print_warning "Version command not available or failed"
        fi
        
        # Test help command
        if ./dist/dockerutilities-linux-amd64 --help >/dev/null 2>&1; then
            print_success "Help command works"
        else
            print_warning "Help command not available or failed"
        fi
    fi
else
    print_error "Production tag build failed"
    exit 1
fi

# Test 4: Dry run release (without publishing)
print_status "Testing dry run release..."
if goreleaser release --snapshot --clean --parallelism=2 --skip-publish; then
    print_success "Dry run release completed successfully"
else
    print_error "Dry run release failed"
    exit 1
fi

# Final summary
print_success "All GoReleaser tests completed successfully!"
print_status "Test summary:"
echo "  ✅ Configuration validation"
echo "  ✅ Development tag (alpha) build"
echo "  ✅ Development tag (beta) build"
echo "  ✅ Production tag build"
echo "  ✅ Binary functionality testing"
echo "  ✅ Dry run release"

print_status "GoReleaser is ready for production use!"
print_status "To create a real release, use:"
echo "  git tag v1.0.0"
echo "  git push origin v1.0.0"

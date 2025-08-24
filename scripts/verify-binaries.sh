#!/bin/bash

# Binary verification script for dockerutilities
# This script verifies built binaries for all platforms
# Part of Milestone 2.2: Artifact Management

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

# Function to get file size in human readable format
get_file_size() {
    local file="$1"
    if [ -f "$file" ]; then
        stat -c%s "$file" 2>/dev/null || stat -f%z "$file" 2>/dev/null || echo "unknown"
    else
        echo "0"
    fi
}

# Function to test binary functionality
test_binary() {
    local binary="$1"
    local platform="$2"
    local arch="$3"
    
    print_status "Testing $platform/$arch binary: $binary"
    
    if [ ! -f "$binary" ]; then
        print_error "Binary not found: $binary"
        return 1
    fi
    
    # Get file size
    local size=$(get_file_size "$binary")
    print_status "Binary size: $size bytes"
    
    # Make executable if needed
    if [ "$platform" != "windows" ]; then
        chmod +x "$binary"
    fi
    
    # Test version command
    print_status "Testing version command..."
    local version_output
    if [ "$platform" = "windows" ]; then
        version_output=$("$binary" version 2>/dev/null || echo "Version command failed")
    else
        version_output=$("$binary" version 2>/dev/null || echo "Version command failed")
    fi
    echo "Version output: $version_output"
    
    # Test help command
    print_status "Testing help command..."
    local help_output
    if [ "$platform" = "windows" ]; then
        help_output=$("$binary" --help 2>/dev/null || echo "Help command failed")
    else
        help_output=$("$binary" --help 2>/dev/null || echo "Help command failed")
    fi
    echo "Help output length: ${#help_output} characters"
    
    # Basic functionality test
    print_status "Testing basic functionality..."
    if [ "$platform" = "windows" ]; then
        if "$binary" version >/dev/null 2>&1; then
            print_success "$platform/$arch binary basic test passed"
        else
            print_warning "$platform/$arch binary basic test failed"
        fi
    else
        if "$binary" version >/dev/null 2>&1; then
            print_success "$platform/$arch binary basic test passed"
        else
            print_warning "$platform/$arch binary basic test failed"
        fi
    fi
    
    echo ""
}

# Function to verify checksums
verify_checksums() {
    local checksums_file="$1"
    
    print_status "Verifying checksums from: $checksums_file"
    
    if [ ! -f "$checksums_file" ]; then
        print_error "Checksums file not found: $checksums_file"
        return 1
    fi
    
    # Change to directory containing checksums file
    local checksums_dir=$(dirname "$checksums_file")
    local checksums_name=$(basename "$checksums_file")
    
    cd "$checksums_dir"
    
    if command_exists sha256sum; then
        if sha256sum -c "$checksums_name"; then
            print_success "All checksums verified successfully"
        else
            print_error "Checksum verification failed"
            return 1
        fi
    elif command_exists shasum; then
        if shasum -a 256 -c "$checksums_name"; then
            print_success "All checksums verified successfully"
        else
            print_error "Checksum verification failed"
            return 1
        fi
    else
        print_warning "No checksum verification tool available (sha256sum or shasum)"
    fi
    
    cd - >/dev/null
}

# Main verification function
verify_binaries() {
    local dist_dir="${1:-dist}"
    
    print_status "Starting binary verification for directory: $dist_dir"
    
    if [ ! -d "$dist_dir" ]; then
        print_error "Distribution directory not found: $dist_dir"
        exit 1
    fi
    
    # Check for checksums file
    local checksums_file="$dist_dir/checksums.txt"
    if [ -f "$checksums_file" ]; then
        print_success "Found checksums file: $checksums_file"
        echo "Checksums content:"
        cat "$checksums_file"
        echo ""
        
        # Verify checksums
        verify_checksums "$checksums_file"
    else
        print_warning "No checksums file found: $checksums_file"
    fi
    
    # Expected binaries for all platforms
    local expected_binaries=(
        "dockerutilities-linux-amd64:linux:amd64"
        "dockerutilities-linux-arm64:linux:arm64"
        "dockerutilities-darwin-amd64:darwin:amd64"
        "dockerutilities-darwin-arm64:darwin:arm64"
        "dockerutilities-windows-amd64.exe:windows:amd64"
        "dockerutilities-windows-arm64.exe:windows:arm64"
    )
    
    local found_count=0
    local total_count=${#expected_binaries[@]}
    
    print_status "Testing $total_count expected binaries..."
    
    # Test each expected binary
    for binary_info in "${expected_binaries[@]}"; do
        IFS=':' read -r binary_name platform arch <<< "$binary_info"
        local binary_path="$dist_dir/$binary_name"
        
        if [ -f "$binary_path" ]; then
            test_binary "$binary_path" "$platform" "$arch"
            ((found_count++))
        else
            print_warning "Expected binary not found: $binary_path"
        fi
    done
    
    # Summary
    echo ""
    print_status "Verification Summary:"
    echo "========================"
    echo "Expected binaries: $total_count"
    echo "Found binaries: $found_count"
    echo "Missing binaries: $((total_count - found_count))"
    
    if [ "$found_count" -eq "$total_count" ]; then
        print_success "All expected binaries found and tested!"
    else
        print_warning "Some expected binaries are missing"
    fi
    
    # Check for unexpected files
    local unexpected_files=$(find "$dist_dir" -name "dockerutilities-*" -type f | grep -v -E "(linux-amd64|linux-arm64|darwin-amd64|darwin-arm64|windows-amd64\.exe|windows-arm64\.exe)$" || true)
    
    if [ -n "$unexpected_files" ]; then
        print_warning "Unexpected files found:"
        echo "$unexpected_files"
    fi
}

# Function to show usage
show_usage() {
    echo "Usage: $0 [OPTIONS] [DIST_DIR]"
    echo ""
    echo "Options:"
    echo "  -h, --help     Show this help message"
    echo "  -v, --verbose  Enable verbose output"
    echo ""
    echo "Arguments:"
    echo "  DIST_DIR       Distribution directory (default: dist)"
    echo ""
    echo "Examples:"
    echo "  $0                    # Verify binaries in ./dist"
    echo "  $0 /path/to/dist      # Verify binaries in specified directory"
    echo "  $0 -v                 # Verbose verification"
}

# Parse command line arguments
VERBOSE=false
DIST_DIR="dist"

while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_usage
            exit 0
            ;;
        -v|--verbose)
            VERBOSE=true
            shift
            ;;
        -*)
            print_error "Unknown option: $1"
            show_usage
            exit 1
            ;;
        *)
            DIST_DIR="$1"
            shift
            ;;
    esac
done

# Main execution
echo "🔍 Binary Verification Script"
echo "============================="
echo "Distribution directory: $DIST_DIR"
echo "Verbose mode: $VERBOSE"
echo ""

# Check prerequisites
print_status "Checking prerequisites..."

if ! command_exists stat; then
    print_error "stat command not available"
    exit 1
fi

if ! command_exists find; then
    print_error "find command not available"
    exit 1
fi

print_success "Prerequisites check passed"
echo ""

# Run verification
verify_binaries "$DIST_DIR"

echo ""
print_success "Binary verification completed!"

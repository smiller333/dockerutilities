#!/bin/bash
set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
VERIFY_MODE="${1:-ci}"  # ci or release
BINARY_PATH="${2:-./bin/dockerutilities}"
DIST_DIR="${3:-./dist}"

echo -e "${BLUE}🔍 Binary Verification Script${NC}"
echo "Mode: $VERIFY_MODE"
echo "Binary Path: $BINARY_PATH"
echo "Dist Directory: $DIST_DIR"
echo ""

# Function to run comprehensive smoke tests on a binary
run_smoke_tests() {
    local binary_path="$1"
    local platform_name="$2"
    
    echo -e "${BLUE}🧪 Testing $platform_name binary: $binary_path${NC}"
    
    if [[ ! -f "$binary_path" ]]; then
        echo -e "${RED}❌ Binary not found: $binary_path${NC}"
        return 1
    fi
    
    # Make executable
    chmod +x "$binary_path"
    
    # Test 1: Version command
    echo "  📋 Test 1: Version command"
    VERSION_OUTPUT=$("$binary_path" version 2>&1 || echo "VERSION_COMMAND_FAILED")
    if [[ "$VERSION_OUTPUT" == "VERSION_COMMAND_FAILED" ]]; then
        echo -e "    ${RED}❌ Version command failed${NC}"
        return 1
    fi
    
    # Validate version output contains expected patterns
    REQUIRED_PATTERNS=("dockerutilities" "version")
    for pattern in "${REQUIRED_PATTERNS[@]}"; do
        if ! echo "$VERSION_OUTPUT" | grep -qi "$pattern"; then
            echo -e "    ${RED}❌ Version output missing '$pattern'${NC}"
            return 1
        fi
    done
    echo -e "    ${GREEN}✅ Version command passed${NC}"
    
    # Test 2: Help command
    echo "  📋 Test 2: Help command"
    HELP_OUTPUT=$("$binary_path" --help 2>&1 || echo "HELP_COMMAND_FAILED")
    if [[ "$HELP_OUTPUT" == "HELP_COMMAND_FAILED" ]]; then
        echo -e "    ${RED}❌ Help command failed${NC}"
        return 1
    fi
    
    # Check for required help information
    if ! echo "$HELP_OUTPUT" | grep -q "Docker image analysis and management"; then
        echo -e "    ${RED}❌ Help output missing expected description${NC}"
        return 1
    fi
    
    if ! echo "$HELP_OUTPUT" | grep -q "Available Commands"; then
        echo -e "    ${RED}❌ Help output missing 'Available Commands' section${NC}"
        return 1
    fi
    echo -e "    ${GREEN}✅ Help command passed${NC}"
    
    # Test 3: Server startup (basic test)
    echo "  📋 Test 3: Server startup"
    TEMP_LOG=$(mktemp)
    trap 'rm -f "$TEMP_LOG"' EXIT
    
    # Start server in background
    "$binary_path" server --no-browser --port 0 > "$TEMP_LOG" 2>&1 &
    SERVER_PID=$!
    
    # Wait a moment for server to start
    sleep 3
    
    # Check if server process is still running
    if kill -0 "$SERVER_PID" 2>/dev/null; then
        echo -e "    ${GREEN}✅ Server startup passed${NC}"
        # Clean up server process
        kill "$SERVER_PID" 2>/dev/null || true
        wait "$SERVER_PID" 2>/dev/null || true
    else
        echo -e "    ${RED}❌ Server process died unexpectedly${NC}"
        cat "$TEMP_LOG" | head -3 | sed 's/^/      /'
        return 1
    fi
    
    # Test 4: Binary integrity check
    echo "  📋 Test 4: Binary integrity"
    # Cross-platform file size check
    if [[ "$OSTYPE" == "darwin"* ]]; then
        BINARY_SIZE=$(stat -f%z "$binary_path")
    else
        BINARY_SIZE=$(stat -c%s "$binary_path")
    fi
    if [[ $BINARY_SIZE -lt 1000000 ]]; then
        echo -e "    ${YELLOW}⚠️ Binary size seems small ($BINARY_SIZE bytes)${NC}"
    elif [[ $BINARY_SIZE -gt 50000000 ]]; then
        echo -e "    ${YELLOW}⚠️ Binary size seems large ($BINARY_SIZE bytes)${NC}"
    else
        echo -e "    ${GREEN}✅ Binary size is reasonable ($BINARY_SIZE bytes)${NC}"
    fi
    
    # Test 5: File type verification
    echo "  📋 Test 5: File type verification"
    FILE_TYPE=$(file "$binary_path")
    echo "    File type: $FILE_TYPE"
    
    # Check for required sections (ELF binary analysis)
    if command -v readelf >/dev/null 2>&1 && [[ "$FILE_TYPE" == *"ELF"* ]]; then
        echo "    Analyzing ELF binary sections..."
        readelf -S "$binary_path" | grep -E "(text|data|bss)" >/dev/null && echo -e "    ${GREEN}✅ ELF sections verified${NC}" || echo -e "    ${YELLOW}⚠️ ELF section analysis incomplete${NC}"
    fi
    
    echo -e "${GREEN}✅ $platform_name binary smoke tests passed${NC}"
    return 0
}

# Function to verify binary integrity
verify_binary_integrity() {
    local binary_path="$1"
    local platform_name="$2"
    
    echo -e "${BLUE}🔐 Verifying $platform_name binary integrity${NC}"
    
    if [[ ! -f "$binary_path" ]]; then
        echo -e "${RED}❌ Binary not found: $binary_path${NC}"
        return 1
    fi
    
    # Check file permissions
    if [[ ! -x "$binary_path" ]]; then
        echo -e "${RED}❌ Binary not executable: $binary_path${NC}"
        echo "   File permissions: $(ls -la "$binary_path")"
        return 1
    fi
    
    # Check file size
    if [[ "$OSTYPE" == "darwin"* ]]; then
        BINARY_SIZE=$(stat -f%z "$binary_path")
    else
        BINARY_SIZE=$(stat -c%s "$binary_path")
    fi
    echo "   Binary size: $BINARY_SIZE bytes"
    
    # Check file type
    FILE_TYPE=$(file "$binary_path")
    echo "   File type: $FILE_TYPE"
    
    # Generate checksum
    CHECKSUM=$(sha256sum "$binary_path" | awk '{print $1}')
    echo "   SHA256: $CHECKSUM"
    
    echo -e "${GREEN}✅ $platform_name binary integrity verified${NC}"
    return 0
}

# Main verification logic
if [[ "$VERIFY_MODE" == "ci" ]]; then
    echo -e "${BLUE}🔍 CI Mode: Single Binary Verification${NC}"
    echo ""
    
    # Run smoke tests on single binary
    if run_smoke_tests "$BINARY_PATH" "Development"; then
        echo -e "${GREEN}🎉 CI binary verification passed!${NC}"
    else
        echo -e "${RED}❌ CI binary verification failed!${NC}"
        exit 1
    fi
    
elif [[ "$VERIFY_MODE" == "release" ]]; then
    echo -e "${BLUE}🔍 Release Mode: Multi-Platform Binary Verification${NC}"
    echo ""
    
    # Define all platform binaries
    PLATFORMS=(
        "dist/dockerutilities-linux-amd64:Linux AMD64"
        "dist/dockerutilities-linux-arm64:Linux ARM64"
        "dist/dockerutilities-darwin-amd64:macOS AMD64"
        "dist/dockerutilities-darwin-arm64:macOS ARM64"
        "dist/dockerutilities-windows-amd64.exe:Windows AMD64"
        "dist/dockerutilities-windows-arm64.exe:Windows ARM64"
    )
    
    FAILED_TESTS=()
    FAILED_INTEGRITY=()
    
    # Verify all platform binaries
    for platform_info in "${PLATFORMS[@]}"; do
        IFS=':' read -r binary_path platform_name <<< "$platform_info"
        
        echo -e "${BLUE}📦 Verifying $platform_name${NC}"
        
        # Run smoke tests
        if run_smoke_tests "$binary_path" "$platform_name"; then
            echo -e "${GREEN}✅ $platform_name smoke tests passed${NC}"
        else
            echo -e "${RED}❌ $platform_name smoke tests failed${NC}"
            FAILED_TESTS+=("$platform_name")
        fi
        
        # Verify binary integrity
        if verify_binary_integrity "$binary_path" "$platform_name"; then
            echo -e "${GREEN}✅ $platform_name integrity verified${NC}"
        else
            echo -e "${RED}❌ $platform_name integrity check failed${NC}"
            FAILED_INTEGRITY+=("$platform_name")
        fi
        
        echo ""
    done
    
    # Report results
    echo -e "${BLUE}📊 Verification Summary${NC}"
    echo "========================"
    
    if [[ ${#FAILED_TESTS[@]} -gt 0 ]]; then
        echo -e "${RED}❌ Smoke tests failed for platforms: ${FAILED_TESTS[*]}${NC}"
    else
        echo -e "${GREEN}✅ All platform smoke tests passed${NC}"
    fi
    
    if [[ ${#FAILED_INTEGRITY[@]} -gt 0 ]]; then
        echo -e "${RED}❌ Integrity checks failed for platforms: ${FAILED_INTEGRITY[*]}${NC}"
    else
        echo -e "${GREEN}✅ All platform integrity checks passed${NC}"
    fi
    
    # Overall result
    if [[ ${#FAILED_TESTS[@]} -eq 0 && ${#FAILED_INTEGRITY[@]} -eq 0 ]]; then
        echo -e "${GREEN}🎉 All release binary verification passed!${NC}"
        echo ""
        echo -e "${BLUE}📋 Release Artifacts Summary:${NC}"
        echo "   ✅ 6 platform binaries created"
        echo "   ✅ All binaries executable and functional"
        echo "   ✅ Version commands working"
        echo "   ✅ Help commands working"
        echo "   ✅ Server startup functional"
        echo "   ✅ Binary integrity verified"
        echo "   ✅ SHA256 checksums generated"
    else
        echo -e "${RED}❌ Release binary verification failed!${NC}"
        exit 1
    fi
    
else
    echo -e "${RED}❌ Invalid verification mode: $VERIFY_MODE${NC}"
    echo "Usage: $0 [ci|release] [binary_path] [dist_dir]"
    echo "  ci: Verify single development binary"
    echo "  release: Verify all platform binaries in dist directory"
    exit 1
fi

echo ""
echo -e "${GREEN}🚀 Binary verification completed successfully!${NC}"

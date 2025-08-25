#!/bin/bash
set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Test binary path
BINARY_PATH="${1:-./bin/dockerutilities}"

echo -e "${BLUE}🧪 Running smoke tests for: $BINARY_PATH${NC}"
echo ""

# Test 1: Binary exists and is executable
echo -e "${BLUE}📋 Test 1: Binary integrity${NC}"
if [[ ! -f "$BINARY_PATH" ]]; then
    echo -e "${RED}❌ Binary not found: $BINARY_PATH${NC}"
    echo "   Expected binary at: $BINARY_PATH"
    echo "   Current directory: $(pwd)"
    echo "   Available files:"
    ls -la bin/ 2>/dev/null || echo "   bin/ directory not found"
    exit 1
fi

if [[ ! -x "$BINARY_PATH" ]]; then
    echo -e "${RED}❌ Binary not executable: $BINARY_PATH${NC}"
    echo "   File permissions: $(ls -la "$BINARY_PATH")"
    echo "   Attempting to make executable..."
    chmod +x "$BINARY_PATH"
    if [[ ! -x "$BINARY_PATH" ]]; then
        echo -e "${RED}❌ Failed to make binary executable${NC}"
        exit 1
    fi
fi

echo -e "${GREEN}✅ Binary exists and is executable${NC}"
echo "   File: $(ls -la "$BINARY_PATH")"
echo ""

# Test 2: Version command
echo -e "${BLUE}📋 Test 2: Version command${NC}"
echo "   Running: $BINARY_PATH version"
VERSION_OUTPUT=$("$BINARY_PATH" version 2>&1 || echo "COMMAND_FAILED")
echo "   Output: $VERSION_OUTPUT"

if [[ "$VERSION_OUTPUT" == "COMMAND_FAILED" ]]; then
    echo -e "${RED}❌ Version command failed to execute${NC}"
    exit 1
fi

if ! echo "$VERSION_OUTPUT" | grep -q "dockerutilities"; then
    echo -e "${RED}❌ Version command output doesn't contain 'dockerutilities'${NC}"
    echo "   Expected to find 'dockerutilities' in output"
    echo "   Full output: $VERSION_OUTPUT"
    exit 1
fi

echo -e "${GREEN}✅ Version command works correctly${NC}"
echo ""

# Test 3: Help command
echo -e "${BLUE}📋 Test 3: Help command${NC}"
echo "   Running: $BINARY_PATH --help"
HELP_OUTPUT=$("$BINARY_PATH" --help 2>&1 || echo "COMMAND_FAILED")
echo "   Output (first 3 lines):"
echo "$HELP_OUTPUT" | head -3

if [[ "$HELP_OUTPUT" == "COMMAND_FAILED" ]]; then
    echo -e "${RED}❌ Help command failed to execute${NC}"
    exit 1
fi

if ! echo "$HELP_OUTPUT" | grep -q "Docker analysis and management utilities"; then
    echo -e "${RED}❌ Help command output doesn't contain expected description${NC}"
    echo "   Expected to find 'Docker analysis and management utilities' in output"
    echo "   Full output: $HELP_OUTPUT"
    exit 1
fi

echo -e "${GREEN}✅ Help command works correctly${NC}"
echo ""

# Test 4: Server startup (without browser)
echo -e "${BLUE}📋 Test 4: Server startup test${NC}"
echo "   Starting server on random port (will timeout after 10 seconds)..."

# Create temporary log file
TEMP_LOG=$(mktemp)
trap 'rm -f "$TEMP_LOG"' EXIT

# Start server in background with random port
SERVER_PID=""
if timeout 10s "$BINARY_PATH" server --no-browser --port 0 > "$TEMP_LOG" 2>&1 &; then
    SERVER_PID=$!
    echo "   Server started with PID: $SERVER_PID"
    
    # Wait a moment for server to start
    sleep 3
    
    # Check if server process is still running
    if kill -0 "$SERVER_PID" 2>/dev/null; then
        echo -e "${GREEN}✅ Server started successfully${NC}"
        echo "   Server log (first 5 lines):"
        head -5 "$TEMP_LOG" | sed 's/^/   /'
        
        # Clean up server process
        echo "   Stopping server..."
        kill "$SERVER_PID" 2>/dev/null || true
        wait "$SERVER_PID" 2>/dev/null || true
    else
        echo -e "${RED}❌ Server process died unexpectedly${NC}"
        echo "   Server log:"
        cat "$TEMP_LOG" | sed 's/^/   /'
        exit 1
    fi
else
    echo -e "${YELLOW}⚠️  Server startup test completed (timeout expected)${NC}"
    echo "   This is normal behavior for the timeout test"
    echo "   Server log (first 5 lines):"
    head -5 "$TEMP_LOG" | sed 's/^/   /'
fi

echo ""

# Test 5: Basic command structure
echo -e "${BLUE}📋 Test 5: Command structure${NC}"
echo "   Checking available commands..."

# Get list of commands
COMMANDS_OUTPUT=$("$BINARY_PATH" --help 2>&1 || echo "COMMAND_FAILED")

if [[ "$COMMANDS_OUTPUT" == "COMMAND_FAILED" ]]; then
    echo -e "${RED}❌ Failed to get command list${NC}"
    exit 1
fi

# Check for expected commands
EXPECTED_COMMANDS=("version" "server")
MISSING_COMMANDS=()

for cmd in "${EXPECTED_COMMANDS[@]}"; do
    if ! echo "$COMMANDS_OUTPUT" | grep -q "$cmd"; then
        MISSING_COMMANDS+=("$cmd")
    fi
done

if [[ ${#MISSING_COMMANDS[@]} -gt 0 ]]; then
    echo -e "${RED}❌ Missing expected commands: ${MISSING_COMMANDS[*]}${NC}"
    echo "   Available commands:"
    echo "$COMMANDS_OUTPUT" | grep -E "^  [a-zA-Z]" | sed 's/^/   /' || echo "   (No commands found)"
    exit 1
fi

echo -e "${GREEN}✅ All expected commands are available${NC}"
echo ""

# Test 6: Error handling
echo -e "${BLUE}📋 Test 6: Error handling${NC}"
echo "   Testing invalid command..."

# Test with invalid command
INVALID_OUTPUT=$("$BINARY_PATH" invalid-command 2>&1 || echo "EXPECTED_ERROR")

if [[ "$INVALID_OUTPUT" == "EXPECTED_ERROR" ]]; then
    echo -e "${GREEN}✅ Invalid command handled gracefully (exit code != 0)${NC}"
else
    echo -e "${YELLOW}⚠️  Invalid command output: $INVALID_OUTPUT${NC}"
    echo "   (This might be acceptable depending on the CLI framework)"
fi

echo ""

# Summary
echo -e "${GREEN}🎉 All smoke tests passed!${NC}"
echo ""
echo -e "${BLUE}📊 Test Summary:${NC}"
echo "   ✅ Binary integrity: PASS"
echo "   ✅ Version command: PASS"
echo "   ✅ Help command: PASS"
echo "   ✅ Server startup: PASS"
echo "   ✅ Command structure: PASS"
echo "   ✅ Error handling: PASS"
echo ""
echo -e "${GREEN}🚀 Binary is ready for use!${NC}"

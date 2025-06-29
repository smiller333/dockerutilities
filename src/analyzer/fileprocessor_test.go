package analyzer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAnalyzeDockerfile(t *testing.T) {
	// Create a temporary test Dockerfile
	tmpDir := t.TempDir()
	testDockerfile := filepath.Join(tmpDir, "Dockerfile")

	testContent := `FROM ubuntu:20.04
RUN apt-get update && apt-get install -y curl
COPY . /app
WORKDIR /app
CMD ["./app"]`

	err := os.WriteFile(testDockerfile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test Dockerfile: %v", err)
	}

	// Test analyzing the Dockerfile
	result, err := AnalyzeDockerfile(testDockerfile)
	if err != nil {
		t.Fatalf("AnalyzeDockerfile failed: %v", err)
	}

	// Verify the result
	if result.Path != testDockerfile {
		t.Errorf("Expected path %s, got %s", testDockerfile, result.Path)
	}

	if result.DFSize != len(testContent) {
		t.Errorf("Expected size %d, got %d", len(testContent), result.DFSize)
	}

	if result.Content != testContent {
		t.Errorf("Content mismatch")
	}

	expectedAbsPath, _ := filepath.Abs(testDockerfile)
	if result.AbsolutePath != expectedAbsPath {
		t.Errorf("Expected absolute path %s, got %s", expectedAbsPath, result.AbsolutePath)
	}
}

func TestAnalyzeDockerfile_NonexistentFile(t *testing.T) {
	_, err := AnalyzeDockerfile("/nonexistent/Dockerfile")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

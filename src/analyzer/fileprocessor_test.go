package analyzer

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// MockCloser implements io.Closer for testing
type MockCloser struct {
	closeError error
	closed     bool
}

func (m *MockCloser) Close() error {
	m.closed = true
	return m.closeError
}

func TestCloseWithErrorCheck(t *testing.T) {
	tests := []struct {
		name         string
		closer       io.Closer
		resourceName string
		expectOutput string
	}{
		{
			name:         "successful close",
			closer:       &MockCloser{closeError: nil},
			resourceName: "test resource",
			expectOutput: "",
		},
		{
			name:         "close with error",
			closer:       &MockCloser{closeError: io.ErrClosedPipe},
			resourceName: "test resource",
			expectOutput: "Warning: failed to close test resource: io: read/write on closed pipe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function
			closeWithErrorCheck(tt.closer, tt.resourceName)

			// Restore stdout
			if err := w.Close(); err != nil {
				t.Errorf("Failed to close pipe writer: %v", err)
			}
			os.Stdout = oldStdout

			// Read captured output
			var buf bytes.Buffer
			if _, err := buf.ReadFrom(r); err != nil {
				t.Errorf("Failed to read from pipe: %v", err)
			}
			output := buf.String()

			// Check output
			if tt.expectOutput == "" {
				if output != "" {
					t.Errorf("Expected no output, got: %s", output)
				}
			} else {
				if !strings.Contains(output, tt.expectOutput) {
					t.Errorf("Expected output to contain '%s', got: %s", tt.expectOutput, output)
				}
			}

			// Verify closer was called
			if mockCloser, ok := tt.closer.(*MockCloser); ok {
				if !mockCloser.closed {
					t.Error("Expected closer to be closed")
				}
			}
		})
	}
}

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

func TestAnalyzeDockerfile_UnreadableFile(t *testing.T) {
	// Create a directory with the name of a Dockerfile to make it unreadable
	tmpDir := t.TempDir()
	unreadableDockerfile := filepath.Join(tmpDir, "Dockerfile")

	// Create a directory instead of a file
	err := os.Mkdir(unreadableDockerfile, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	_, err = AnalyzeDockerfile(unreadableDockerfile)
	if err == nil {
		t.Error("Expected error for unreadable file, got nil")
	}
}

func TestAnalyzeDockerfile_RelativePath(t *testing.T) {
	// Create a temporary test Dockerfile
	tmpDir := t.TempDir()
	testDockerfile := filepath.Join(tmpDir, "Dockerfile")

	testContent := `FROM ubuntu:20.04
CMD ["echo", "hello"]`

	err := os.WriteFile(testDockerfile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test Dockerfile: %v", err)
	}

	// Change to the temp directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer func() {
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Test with relative path
	result, err := AnalyzeDockerfile("Dockerfile")
	if err != nil {
		t.Fatalf("AnalyzeDockerfile failed: %v", err)
	}

	// Verify the result
	if result.Path != "Dockerfile" {
		t.Errorf("Expected path Dockerfile, got %s", result.Path)
	}

	expectedAbsPath, _ := filepath.Abs("Dockerfile")
	if result.AbsolutePath != expectedAbsPath {
		t.Errorf("Expected absolute path %s, got %s", expectedAbsPath, result.AbsolutePath)
	}
}

func TestGenerateImageTag(t *testing.T) {
	tests := []struct {
		name           string
		dockerfilePath string
		expectedTag    string
	}{
		{
			name:           "simple directory name",
			dockerfilePath: "/path/to/myapp/Dockerfile",
			expectedTag:    "myapp:test",
		},
		{
			name:           "directory with spaces",
			dockerfilePath: "/path/to/my app/Dockerfile",
			expectedTag:    "my-app:test",
		},
		{
			name:           "directory with special characters",
			dockerfilePath: "/path/to/my_app-v1.0/Dockerfile",
			expectedTag:    "my_app-v1.0:test",
		},
		{
			name:           "root directory",
			dockerfilePath: "/Dockerfile",
			expectedTag:    "dockerutilities:test",
		},
		{
			name:           "current directory",
			dockerfilePath: "./Dockerfile",
			expectedTag:    "dockerutilities:test",
		},
		{
			name:           "empty directory name",
			dockerfilePath: "//Dockerfile",
			expectedTag:    "dockerutilities:test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateImageTag(tt.dockerfilePath)
			if result != tt.expectedTag {
				t.Errorf("Expected tag %s, got %s", tt.expectedTag, result)
			}
		})
	}
}

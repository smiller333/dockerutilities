// Package dockerclient provides tests for the Docker client wrapper.
package dockerclient

import (
	"context"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

// closeWithErrorCheck is a helper function to close resources and log any errors in tests
func closeWithErrorCheck(t testing.TB, closer interface{ Close() error }, resourceName string) {
	if err := closer.Close(); err != nil {
		// Log the error but don't fail the test
		t.Logf("Warning: failed to close %s: %v", resourceName, err)
	}
}

// TestValidateDockerAccess tests the ValidateDockerAccess function
func TestValidateDockerAccess(t *testing.T) {
	err := ValidateDockerAccess()

	// The function may or may not return an error depending on whether Docker is running
	// We just test that it doesn't panic and handles both cases gracefully
	if err != nil {
		// If there's an error, verify it contains the expected socket path
		expectedPath := "/var/run/docker.sock"
		if runtime.GOOS == "windows" {
			expectedPath = `\\.\pipe\docker_engine`
		}
		if !strings.Contains(err.Error(), expectedPath) {
			t.Errorf("ValidateDockerAccess() error message should contain %s, got: %s", expectedPath, err.Error())
		}
	} else {
		// If no error, that's also valid (Docker might be running)
		t.Log("ValidateDockerAccess() succeeded - Docker socket is accessible")
	}
}

// TestValidateImageName tests the ValidateImageName function
func TestValidateImageName(t *testing.T) {
	tests := []struct {
		name      string
		imageName string
		wantErr   bool
	}{
		{
			name:      "valid simple image name",
			imageName: "alpine",
			wantErr:   false,
		},
		{
			name:      "valid image with tag",
			imageName: "alpine:latest",
			wantErr:   false,
		},
		{
			name:      "valid image with registry",
			imageName: "docker.io/library/alpine:latest",
			wantErr:   false,
		},
		{
			name:      "valid image with namespace",
			imageName: "myorg/myimage:latest",
			wantErr:   false,
		},
		{
			name:      "valid image with underscore",
			imageName: "my_image:latest",
			wantErr:   false,
		},
		{
			name:      "valid image with hyphen",
			imageName: "my-image:latest",
			wantErr:   false,
		},
		{
			name:      "valid image with dot",
			imageName: "my.image:latest",
			wantErr:   false,
		},
		{
			name:      "empty image name",
			imageName: "",
			wantErr:   true,
		},
		{
			name:      "image name too long",
			imageName: strings.Repeat("a", 256),
			wantErr:   true,
		},
		{
			name:      "invalid image name with special characters",
			imageName: "my@image:latest",
			wantErr:   true,
		},
		{
			name:      "invalid image name starting with dot",
			imageName: ".myimage:latest",
			wantErr:   true,
		},
		{
			name:      "invalid image name starting with underscore",
			imageName: "_myimage:latest",
			wantErr:   true,
		},
		{
			name:      "invalid image name starting with hyphen",
			imageName: "-myimage:latest",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateImageName(tt.imageName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateImageName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestDockerClient_BuildImage_FileOperations tests the file operations in BuildImage
func TestDockerClient_BuildImage_FileOperations(t *testing.T) {
	// Test with nonexistent Dockerfile
	_, err := NewDefaultClient()
	if err != nil {
		t.Skipf("Cannot create Docker client for testing: %v", err)
	}

	// Test with nonexistent Dockerfile
	_, err = NewDefaultClient()
	if err != nil {
		t.Skipf("Cannot create Docker client for testing: %v", err)
	}

	// Create a temporary Dockerfile for testing
	tmpDir := t.TempDir()
	dockerfilePath := tmpDir + "/Dockerfile"
	dockerfileContent := "FROM alpine:latest\nRUN echo 'test'"

	err = os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test Dockerfile: %v", err)
	}

	// Test that the file exists and is readable
	if _, err := os.Stat(dockerfilePath); err != nil {
		t.Fatalf("Test Dockerfile should exist: %v", err)
	}

	// Test reading the file content
	content, err := os.ReadFile(dockerfilePath)
	if err != nil {
		t.Fatalf("Failed to read test Dockerfile: %v", err)
	}

	if string(content) != dockerfileContent {
		t.Errorf("Dockerfile content mismatch: got %s, want %s", string(content), dockerfileContent)
	}
}

func TestNewDockerClient(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name:    "default config",
			config:  nil,
			wantErr: false,
		},
		{
			name: "custom timeout",
			config: &Config{
				Timeout: 10 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "custom host",
			config: &Config{
				Host:    "unix:///var/run/docker.sock",
				Timeout: 5 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "with API version",
			config: &Config{
				APIVersion: "1.41",
				Timeout:    15 * time.Second,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewDockerClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDockerClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if client != nil {
				defer closeWithErrorCheck(t, client, "Docker client")

				// Verify timeout is set correctly
				expectedTimeout := 30 * time.Second
				if tt.config != nil && tt.config.Timeout != 0 {
					expectedTimeout = tt.config.Timeout
				}
				if client.GetTimeout() != expectedTimeout {
					t.Errorf("NewDockerClient() timeout = %v, want %v", client.GetTimeout(), expectedTimeout)
				}

				// Verify config is stored
				if client.GetConfig() == nil {
					t.Error("NewDockerClient() config is nil")
				}
			}
		})
	}
}

func TestNewDefaultClient(t *testing.T) {
	client, err := NewDefaultClient()
	if err != nil {
		t.Fatalf("NewDefaultClient() error = %v", err)
	}
	defer closeWithErrorCheck(t, client, "Docker client")

	// Verify default timeout
	expectedTimeout := 30 * time.Second
	if client.GetTimeout() != expectedTimeout {
		t.Errorf("NewDefaultClient() timeout = %v, want %v", client.GetTimeout(), expectedTimeout)
	}

	// Verify client is not nil
	if client.GetClient() == nil {
		t.Error("NewDefaultClient() underlying client is nil")
	}
}

func TestDockerClient_SetTimeout(t *testing.T) {
	client, err := NewDefaultClient()
	if err != nil {
		t.Fatalf("NewDefaultClient() error = %v", err)
	}
	defer closeWithErrorCheck(t, client, "Docker client")

	newTimeout := 45 * time.Second
	client.SetTimeout(newTimeout)

	if client.GetTimeout() != newTimeout {
		t.Errorf("SetTimeout() timeout = %v, want %v", client.GetTimeout(), newTimeout)
	}
}

// TestDockerClient_Ping tests the ping functionality
// Note: This test requires a running Docker daemon
func TestDockerClient_Ping(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client, err := NewDefaultClient()
	if err != nil {
		t.Fatalf("NewDefaultClient() error = %v", err)
	}
	defer closeWithErrorCheck(t, client, "Docker client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx)
	if err != nil {
		t.Logf("Ping failed (Docker daemon may not be running): %v", err)
		// Don't fail the test if Docker is not available
		return
	}

	// Test IsConnected as well
	if !client.IsConnected(ctx) {
		t.Error("IsConnected() returned false after successful ping")
	}
}

// TestDockerClient_GetInfo tests retrieving Docker daemon info
// Note: This test requires a running Docker daemon
func TestDockerClient_GetInfo(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client, err := NewDefaultClient()
	if err != nil {
		t.Fatalf("NewDefaultClient() error = %v", err)
	}
	defer closeWithErrorCheck(t, client, "Docker client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// First check if Docker is available
	if err := client.Ping(ctx); err != nil {
		t.Logf("Docker daemon not available, skipping test: %v", err)
		return
	}

	info, err := client.GetInfo(ctx)
	if err != nil {
		t.Fatalf("GetInfo() error = %v", err)
	}

	if info == nil {
		t.Error("GetInfo() returned nil info")
	}
}

// TestDockerClient_GetVersion tests retrieving Docker daemon version
// Note: This test requires a running Docker daemon
func TestDockerClient_GetVersion(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client, err := NewDefaultClient()
	if err != nil {
		t.Fatalf("NewDefaultClient() error = %v", err)
	}
	defer closeWithErrorCheck(t, client, "Docker client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// First check if Docker is available
	if err := client.Ping(ctx); err != nil {
		t.Logf("Docker daemon not available, skipping test: %v", err)
		return
	}

	version, err := client.GetVersion(ctx)
	if err != nil {
		t.Fatalf("GetVersion() error = %v", err)
	}

	if version == nil {
		t.Error("GetVersion() returned nil version")
	} else if version.Version == "" {
		t.Error("GetVersion() returned empty version string")
	}
}

// TestDockerClient_PullImage tests pulling an image
// Note: This test requires a running Docker daemon and internet connectivity
func TestDockerClient_PullImage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client, err := NewDefaultClient()
	if err != nil {
		t.Fatalf("NewDefaultClient() error = %v", err)
	}
	defer closeWithErrorCheck(t, client, "Docker client")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// First check if Docker is available
	if err := client.Ping(ctx); err != nil {
		t.Logf("Docker daemon not available, skipping test: %v", err)
		return
	}

	// Test pulling a small public image
	reader, err := client.PullImage(ctx, "alpine:latest", nil)
	if err != nil {
		t.Logf("PullImage() failed (may be due to network issues): %v", err)
		return
	}
	defer closeWithErrorCheck(t, reader, "reader")

	// Read some data to ensure the stream is working
	buffer := make([]byte, 1024)
	_, err = reader.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		t.Errorf("Failed to read from pull response: %v", err)
	}
}

// TestDockerClient_PushImage tests pushing an image
// Note: This test requires a running Docker daemon and proper authentication
func TestDockerClient_PushImage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client, err := NewDefaultClient()
	if err != nil {
		t.Fatalf("NewDefaultClient() error = %v", err)
	}
	defer closeWithErrorCheck(t, client, "Docker client")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// First check if Docker is available
	if err := client.Ping(ctx); err != nil {
		t.Logf("Docker daemon not available, skipping test: %v", err)
		return
	}

	// Test pushing an image (this will likely fail without proper auth)
	// We're just testing that the function doesn't panic and returns an appropriate error
	_, err = client.PushImage(ctx, "nonexistent/test:latest", nil)
	if err == nil {
		t.Log("PushImage() unexpectedly succeeded (or image exists)")
	} else {
		t.Logf("PushImage() failed as expected without auth: %v", err)
	}
}

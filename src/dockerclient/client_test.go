// Package dockerclient provides tests for the Docker client wrapper.
package dockerclient

import (
	"context"
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

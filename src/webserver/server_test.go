package webserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smiller333/dockerutilities/src/analyzer"
)

// closeWithErrorCheck is a helper function to close resources and log any errors in tests
func closeWithErrorCheck(t testing.TB, closer interface{ Close() error }, resourceName string) {
	if err := closer.Close(); err != nil {
		// Log the error but don't fail the test
		t.Logf("Warning: failed to close %s: %v", resourceName, err)
	}
}

// closeWithErrorCheckNoReturn is a helper function to close resources that don't return errors
func closeWithErrorCheckNoReturn(closer interface{ Close() }, resourceName string) {
	closer.Close()
	// resourceName is used for consistency with closeWithErrorCheck function signature
	_ = resourceName // explicitly ignore to satisfy linter
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name:    "nil config",
			config:  nil,
			wantErr: true,
		},
		{
			name: "valid config",
			config: &Config{
				Host: "localhost",
				Port: "8080",
			},
			wantErr: false,
		},
		{
			name:    "empty config with defaults",
			config:  &Config{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, err := New(tt.config)
			if tt.wantErr {
				if err == nil {
					t.Errorf("New() expected error, got nil")
				}
				if server != nil {
					t.Errorf("New() expected nil server, got %v", server)
				}
			} else {
				if err != nil {
					t.Errorf("New() unexpected error: %v", err)
				}
				if server == nil {
					t.Errorf("New() expected server, got nil")
				}

				// Check defaults are applied
				if tt.config != nil && server != nil {
					if tt.config.Host == "" && server.config.Host != "localhost" {
						t.Errorf("Expected default host 'localhost', got %s", server.config.Host)
					}
					if tt.config.Port == "" && server.config.Port != "8080" {
						t.Errorf("Expected default port '8080', got %s", server.config.Port)
					}
				}
			}
		})
	}
}

func TestHealthEndpoint(t *testing.T) {
	// Create a test server
	config := &Config{
		Host:    "localhost",
		Port:    "8080",
		WebRoot: ".", // Use current directory for testing
	}

	server, err := New(config)
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	// Create a test HTTP server
	mux := http.NewServeMux()
	server.registerRoutes(mux)
	testServer := httptest.NewServer(mux)
	defer closeWithErrorCheckNoReturn(testServer, "testServer")

	// Test health endpoint
	resp, err := http.Get(testServer.URL + "/api/health")
	if err != nil {
		t.Fatalf("http.Get() unexpected error: %v", err)
	}
	defer closeWithErrorCheck(t, resp.Body, "response body")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expectedContentType := "application/json"
	if contentType := resp.Header.Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}

	// Parse response
	var health map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&health)
	if err != nil {
		t.Fatalf("json.Decode() unexpected error: %v", err)
	}

	if status, ok := health["status"]; !ok || status != "healthy" {
		t.Errorf("Expected status 'healthy', got %v", status)
	}
	if _, ok := health["timestamp"]; !ok {
		t.Errorf("Expected 'timestamp' field in response")
	}
	if _, ok := health["version"]; !ok {
		t.Errorf("Expected 'version' field in response")
	}
}

func TestSummariesEndpoint(t *testing.T) {
	// Create a test server
	config := &Config{
		Host:    "localhost",
		Port:    "8080",
		WebRoot: ".", // Use current directory for testing
	}

	server, err := New(config)
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	// Create a test HTTP server
	mux := http.NewServeMux()
	server.registerRoutes(mux)
	testServer := httptest.NewServer(mux)
	defer closeWithErrorCheckNoReturn(testServer, "testServer")

	// Test summaries endpoint
	resp, err := http.Get(testServer.URL + "/api/summaries")
	if err != nil {
		t.Fatalf("http.Get() unexpected error: %v", err)
	}
	defer closeWithErrorCheck(t, resp.Body, "response body")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expectedContentType := "application/json"
	if contentType := resp.Header.Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}

	// Parse response - should be an array (empty if no summaries exist)
	var infos []analyzer.ImageInfo
	err = json.NewDecoder(resp.Body).Decode(&infos)
	if err != nil {
		t.Fatalf("json.Decode() unexpected error: %v", err)
	}

	// Should be a valid array (summaries can be empty if no files exist)
	// In Go, an empty slice is not nil, so this check verifies proper initialization
	if infos == nil {
		t.Errorf("Expected initialized summaries slice, got nil")
	}
}

func TestMethodNotAllowed(t *testing.T) {
	// Create a test server
	config := &Config{
		Host:    "localhost",
		Port:    "8080",
		WebRoot: ".", // Use current directory for testing
	}

	server, err := New(config)
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	// Create a test HTTP server
	mux := http.NewServeMux()
	server.registerRoutes(mux)
	testServer := httptest.NewServer(mux)
	defer closeWithErrorCheckNoReturn(testServer, "testServer")

	// Test POST to health endpoint (should return 404 with Go 1.22+ pattern matching)
	resp, err := http.Post(testServer.URL+"/api/health", "application/json", nil)
	if err != nil {
		t.Fatalf("http.Post() unexpected error: %v", err)
	}
	defer closeWithErrorCheck(t, resp.Body, "response body")

	// With Go 1.22+ pattern matching, unmatched methods return 404, not 405
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}

func TestSummaryInfo(t *testing.T) {
	// Test SummaryInfo struct
	info := analyzer.ImageInfo{
		ImageTag:     "nginx:latest",
		ImageID:      "sha256:abc123",
		AnalyzedAt:   "2025-06-30T12:00:00Z",
		Architecture: "amd64",
		OS:           "linux",
		LayerCount:   5,
		ImageSize:    123456789,
	}

	// Marshal to JSON and back
	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("json.Marshal() unexpected error: %v", err)
	}

	var parsed analyzer.ImageInfo
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		t.Fatalf("json.Unmarshal() unexpected error: %v", err)
	}

	// Compare fields individually for better error messages
	if parsed.ImageTag != info.ImageTag {
		t.Errorf("ImageTag mismatch: expected %s, got %s", info.ImageTag, parsed.ImageTag)
	}
	if parsed.ImageID != info.ImageID {
		t.Errorf("ImageID mismatch: expected %s, got %s", info.ImageID, parsed.ImageID)
	}
	if parsed.AnalyzedAt != info.AnalyzedAt {
		t.Errorf("AnalyzedAt mismatch: expected %s, got %s", info.AnalyzedAt, parsed.AnalyzedAt)
	}
	if parsed.Architecture != info.Architecture {
		t.Errorf("Architecture mismatch: expected %s, got %s", info.Architecture, parsed.Architecture)
	}
	if parsed.OS != info.OS {
		t.Errorf("OS mismatch: expected %s, got %s", info.OS, parsed.OS)
	}
	if parsed.LayerCount != info.LayerCount {
		t.Errorf("LayerCount mismatch: expected %d, got %d", info.LayerCount, parsed.LayerCount)
	}
	if parsed.ImageSize != info.ImageSize {
		t.Errorf("ImageSize mismatch: expected %d, got %d", info.ImageSize, parsed.ImageSize)
	}
}

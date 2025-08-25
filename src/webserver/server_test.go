package webserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

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

// createTestServer creates a test server with a temporary directory
func createTestServer(t *testing.T) (*Server, string) {
	// Create a temp directory in the current working directory to avoid /var restrictions
	tmpDir := filepath.Join(".", "test-tmp-"+t.Name())
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		t.Fatalf("Failed to create test tmp dir: %v", err)
	}

	// Clean up after test
	t.Cleanup(func() {
		os.RemoveAll(tmpDir)
	})

	config := &Config{
		Host:    "localhost",
		Port:    "8080",
		TmpDir:  tmpDir,
		WebRoot: ".",
	}

	server, err := New(config)
	if err != nil {
		t.Fatalf("Failed to create test server: %v", err)
	}

	return server, tmpDir
}

// createTestImageInfo creates a test ImageInfo for testing
func createTestImageInfo() analyzer.ImageInfo {
	return analyzer.ImageInfo{
		ImageTag:     "nginx:latest",
		ImageID:      "sha256:abc123def456",
		ImageSource:  "docker.io/library/nginx",
		ImageSize:    123456789,
		Architecture: "amd64",
		OS:           "linux",
		LayerCount:   5,
		AnalyzedAt:   time.Now().Format(time.RFC3339),
		Layers:       []string{"sha256:layer1"},
	}
}

// createTestImageSummary creates a test ImageSummary for testing
func createTestImageSummary() ImageSummary {
	return ImageSummary{
		ImageID:      "sha256:abc123def456",
		ImageTag:     "nginx:latest",
		ImageSource:  "docker.io/library/nginx",
		ImageSize:    123456789,
		Architecture: "amd64",
		AnalyzedAt:   time.Now().Format(time.RFC3339),
		Status:       "completed",
		RequestID:    "req-123",
	}
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

func TestNewWithDockerClient(t *testing.T) {
	tmpDir := t.TempDir()
	config := &Config{
		Host:    "localhost",
		Port:    "8080",
		TmpDir:  tmpDir,
		WebRoot: ".",
	}

	server, err := New(config)
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if server == nil {
		t.Fatal("New() expected server, got nil")
	}

	if server.tmpDir != tmpDir {
		t.Errorf("Expected tmpDir %s, got %s", tmpDir, server.tmpDir)
	}
}

func TestServerStart(t *testing.T) {
	server, _ := createTestServer(t)

	// Test that Start() returns an error when trying to bind to an invalid port
	server.config.Port = "99999" // Invalid port
	err := server.Start()
	if err == nil {
		t.Error("Start() expected error with invalid port, got nil")
	}
}

func TestServerShutdown(t *testing.T) {
	server, _ := createTestServer(t)

	// Test shutdown when server is nil
	err := server.Shutdown()
	if err != nil {
		t.Errorf("Shutdown() unexpected error: %v", err)
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

func TestHandleGetInfo(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Create a test image info file
	imageInfo := createTestImageInfo()
	infoID := "test-image-123"
	infoPath := filepath.Join(tmpDir, fmt.Sprintf("info.%s.json", infoID))

	infoData, err := json.Marshal(imageInfo)
	if err != nil {
		t.Fatalf("Failed to marshal image info: %v", err)
	}

	err = os.WriteFile(infoPath, infoData, 0644)
	if err != nil {
		t.Fatalf("Failed to write test info file: %v", err)
	}

	// Test successful get info
	req := httptest.NewRequest("GET", "/api/info/"+infoID, nil)
	w := httptest.NewRecorder()
	server.handleGetInfo(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Test missing info ID
	req = httptest.NewRequest("GET", "/api/info/", nil)
	w = httptest.NewRecorder()
	server.handleGetInfo(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Test non-existent info ID
	req = httptest.NewRequest("GET", "/api/info/non-existent", nil)
	w = httptest.NewRecorder()
	server.handleGetInfo(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestHandleDeleteInfo(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Create a test image info file
	imageInfo := createTestImageInfo()
	infoID := "test-image-123"
	infoPath := filepath.Join(tmpDir, fmt.Sprintf("info.%s.json", infoID))

	infoData, err := json.Marshal(imageInfo)
	if err != nil {
		t.Fatalf("Failed to marshal image info: %v", err)
	}

	err = os.WriteFile(infoPath, infoData, 0644)
	if err != nil {
		t.Fatalf("Failed to write test info file: %v", err)
	}

	// Create a summary entry for this info
	summary := server.imageInfoToSummary(imageInfo)
	summary.RequestID = infoID // Use the same ID for testing
	err = server.addSummaryToFile(summary)
	if err != nil {
		t.Fatalf("Failed to add summary: %v", err)
	}

	// Test successful delete
	req := httptest.NewRequest("DELETE", "/api/info/"+infoID, nil)
	w := httptest.NewRecorder()
	server.handleDeleteInfo(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Verify file was deleted
	if _, err := os.Stat(infoPath); !os.IsNotExist(err) {
		t.Error("Expected info file to be deleted")
	}

	// Test missing info ID
	req = httptest.NewRequest("DELETE", "/api/info/", nil)
	w = httptest.NewRecorder()
	server.handleDeleteInfo(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Test non-existent info ID
	req = httptest.NewRequest("DELETE", "/api/info/non-existent", nil)
	w = httptest.NewRecorder()
	server.handleDeleteInfo(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestHandleAnalyzeImage(t *testing.T) {
	server, _ := createTestServer(t)

	// Test invalid request body
	req := httptest.NewRequest("POST", "/api/analyze", strings.NewReader("invalid json"))
	w := httptest.NewRecorder()
	server.handleAnalyzeImage(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Test missing image name
	analyzeReq := AnalyzeRequest{}
	reqBody, _ := json.Marshal(analyzeReq)
	req = httptest.NewRequest("POST", "/api/analyze", bytes.NewReader(reqBody))
	w = httptest.NewRecorder()
	server.handleAnalyzeImage(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Test valid request (will succeed if image already exists or fail gracefully)
	analyzeReq = AnalyzeRequest{ImageName: "nginx:latest"}
	reqBody, _ = json.Marshal(analyzeReq)
	req = httptest.NewRequest("POST", "/api/analyze", bytes.NewReader(reqBody))
	w = httptest.NewRecorder()
	server.handleAnalyzeImage(w, req)

	// Should return either success or error response
	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d or %d, got %d", http.StatusOK, http.StatusInternalServerError, w.Code)
	}
}

func TestHandleAnalyzeImageAsync(t *testing.T) {
	server, _ := createTestServer(t)

	// Test invalid request body
	req := httptest.NewRequest("POST", "/api/analyze-async", strings.NewReader("invalid json"))
	w := httptest.NewRecorder()
	server.handleAnalyzeImageAsync(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Test missing image name
	analyzeReq := AsyncAnalyzeRequest{}
	reqBody, _ := json.Marshal(analyzeReq)
	req = httptest.NewRequest("POST", "/api/analyze-async", bytes.NewReader(reqBody))
	w = httptest.NewRecorder()
	server.handleAnalyzeImageAsync(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Test valid request
	analyzeReq = AsyncAnalyzeRequest{ImageName: "nginx:latest"}
	reqBody, _ = json.Marshal(analyzeReq)
	req = httptest.NewRequest("POST", "/api/analyze-async", bytes.NewReader(reqBody))
	w = httptest.NewRecorder()
	server.handleAnalyzeImageAsync(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Verify response contains request ID
	var response AsyncAnalyzeResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if !response.Success {
		t.Error("Expected success to be true")
	}

	if response.RequestID == "" {
		t.Error("Expected request ID to be set")
	}
}

func TestWriteSummaryFile(t *testing.T) {
	server, tmpDir := createTestServer(t)

	summaries := []ImageSummary{createTestImageSummary()}
	filePath := filepath.Join(tmpDir, "test-summaries.json")

	// Test successful write
	err := server.writeSummaryFile(filePath, summaries)
	if err != nil {
		t.Errorf("writeSummaryFile() unexpected error: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("Expected summary file to be created")
	}

	// Test write to non-existent directory
	invalidPath := filepath.Join(tmpDir, "nonexistent", "test.json")
	err = server.writeSummaryFile(invalidPath, summaries)
	// The function creates directories automatically, so this should succeed
	if err != nil {
		t.Errorf("writeSummaryFile() unexpected error: %v", err)
	}
}

func TestRebuildSummaryFile(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Create test info files
	imageInfo := createTestImageInfo()
	infoID := "test-image-123"
	infoPath := filepath.Join(tmpDir, fmt.Sprintf("info.%s.json", infoID))

	infoData, err := json.Marshal(imageInfo)
	if err != nil {
		t.Fatalf("Failed to marshal image info: %v", err)
	}

	err = os.WriteFile(infoPath, infoData, 0644)
	if err != nil {
		t.Fatalf("Failed to write test info file: %v", err)
	}

	// Create a summary file for this info
	summary := server.imageInfoToSummary(imageInfo)
	summaryPath := filepath.Join(tmpDir, fmt.Sprintf("summary.%s.json", infoID))
	summaryData, _ := json.Marshal(summary)
	os.WriteFile(summaryPath, summaryData, 0644)

	summaryFilePath := filepath.Join(tmpDir, summaryFileName)
	summaries, err := server.rebuildSummaryFile(tmpDir, summaryFilePath)
	if err != nil {
		t.Errorf("rebuildSummaryFile() unexpected error: %v", err)
	}

	if len(summaries) == 0 {
		t.Error("Expected at least one summary to be rebuilt")
	}
}

func TestAddSummaryToFile(t *testing.T) {
	server, _ := createTestServer(t)

	summary := createTestImageSummary()

	// Test adding new summary
	err := server.addSummaryToFile(summary)
	if err != nil {
		t.Errorf("addSummaryToFile() unexpected error: %v", err)
	}

	// Test adding duplicate summary (should update existing)
	err = server.addSummaryToFile(summary)
	if err != nil {
		t.Errorf("addSummaryToFile() unexpected error on duplicate: %v", err)
	}
}

func TestRemoveSummaryFromFile(t *testing.T) {
	server, _ := createTestServer(t)

	// First add a summary
	summary := createTestImageSummary()
	err := server.addSummaryToFile(summary)
	if err != nil {
		t.Fatalf("Failed to add summary: %v", err)
	}

	// Test removing existing summary
	err = server.removeSummaryFromFile(summary.ImageID)
	if err != nil {
		t.Errorf("removeSummaryFromFile() unexpected error: %v", err)
	}

	// Test removing non-existent summary
	err = server.removeSummaryFromFile("non-existent")
	// The function logs a warning but doesn't return an error for non-existent summaries
	if err != nil {
		t.Errorf("removeSummaryFromFile() unexpected error: %v", err)
	}
}

func TestRemoveSummaryByRequestID(t *testing.T) {
	server, _ := createTestServer(t)

	// First add a summary with request ID
	summary := createTestImageSummary()
	summary.RequestID = "test-request-123"
	err := server.addSummaryToFile(summary)
	if err != nil {
		t.Fatalf("Failed to add summary: %v", err)
	}

	// Test removing existing summary by request ID
	err = server.removeSummaryByRequestID(summary.RequestID)
	if err != nil {
		t.Errorf("removeSummaryByRequestID() unexpected error: %v", err)
	}

	// Test removing non-existent request ID
	err = server.removeSummaryByRequestID("non-existent")
	if err == nil {
		t.Error("removeSummaryByRequestID() expected error for non-existent request ID, got nil")
	}
}

func TestParseSummaryFile(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Create a valid summary file
	summary := createTestImageSummary()
	filePath := filepath.Join(tmpDir, "test-summary.json")

	summaryData, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("Failed to marshal summary: %v", err)
	}

	err = os.WriteFile(filePath, summaryData, 0644)
	if err != nil {
		t.Fatalf("Failed to write test summary file: %v", err)
	}

	// Test parsing valid file
	parsedSummary, err := server.parseSummaryFile(filePath)
	if err != nil {
		t.Errorf("parseSummaryFile() unexpected error: %v", err)
	}

	if parsedSummary.ImageID != summary.ImageID {
		t.Errorf("Expected ImageID %s, got %s", summary.ImageID, parsedSummary.ImageID)
	}

	// Test parsing non-existent file
	_, err = server.parseSummaryFile("non-existent.json")
	if err == nil {
		t.Error("parseSummaryFile() expected error for non-existent file, got nil")
	}

	// Test parsing invalid JSON
	invalidPath := filepath.Join(tmpDir, "invalid.json")
	os.WriteFile(invalidPath, []byte("invalid json"), 0644)
	_, err = server.parseSummaryFile(invalidPath)
	if err == nil {
		t.Error("parseSummaryFile() expected error for invalid JSON, got nil")
	}
}

func TestParseInfoFile(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Create a valid info file
	imageInfo := createTestImageInfo()
	filePath := filepath.Join(tmpDir, "test-info.json")

	infoData, err := json.Marshal(imageInfo)
	if err != nil {
		t.Fatalf("Failed to marshal image info: %v", err)
	}

	err = os.WriteFile(filePath, infoData, 0644)
	if err != nil {
		t.Fatalf("Failed to write test info file: %v", err)
	}

	// Test parsing valid file
	parsedInfo, err := server.parseInfoFile(filePath)
	if err != nil {
		t.Errorf("parseInfoFile() unexpected error: %v", err)
	}

	if parsedInfo.ImageID != imageInfo.ImageID {
		t.Errorf("Expected ImageID %s, got %s", imageInfo.ImageID, parsedInfo.ImageID)
	}

	// Test parsing non-existent file
	_, err = server.parseInfoFile("non-existent.json")
	if err == nil {
		t.Error("parseInfoFile() expected error for non-existent file, got nil")
	}

	// Test parsing invalid JSON
	invalidPath := filepath.Join(tmpDir, "invalid-info.json")
	os.WriteFile(invalidPath, []byte("invalid json"), 0644)
	_, err = server.parseInfoFile(invalidPath)
	if err == nil {
		t.Error("parseInfoFile() expected error for invalid JSON, got nil")
	}
}

func TestGetInfoByID(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Create a test image info file
	imageInfo := createTestImageInfo()
	infoID := "test-image-123"
	infoPath := filepath.Join(tmpDir, fmt.Sprintf("info.%s.json", infoID))

	infoData, err := json.Marshal(imageInfo)
	if err != nil {
		t.Fatalf("Failed to marshal image info: %v", err)
	}

	err = os.WriteFile(infoPath, infoData, 0644)
	if err != nil {
		t.Fatalf("Failed to write test info file: %v", err)
	}

	// Test getting existing info
	parsedInfo, err := server.getInfoByID(infoID)
	if err != nil {
		t.Errorf("getInfoByID() unexpected error: %v", err)
	}

	if parsedInfo.ImageID != imageInfo.ImageID {
		t.Errorf("Expected ImageID %s, got %s", imageInfo.ImageID, parsedInfo.ImageID)
	}

	// Test getting non-existent info
	_, err = server.getInfoByID("non-existent")
	if err == nil {
		t.Error("getInfoByID() expected error for non-existent info, got nil")
	}
}

func TestDeleteInfoByID(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Create a test image info file
	imageInfo := createTestImageInfo()
	infoID := "test-image-123"
	infoPath := filepath.Join(tmpDir, fmt.Sprintf("info.%s.json", infoID))

	infoData, err := json.Marshal(imageInfo)
	if err != nil {
		t.Fatalf("Failed to marshal image info: %v", err)
	}

	err = os.WriteFile(infoPath, infoData, 0644)
	if err != nil {
		t.Fatalf("Failed to write test info file: %v", err)
	}

	// Create a summary entry for this info
	summary := server.imageInfoToSummary(imageInfo)
	summary.RequestID = infoID // Use the same ID for testing
	err = server.addSummaryToFile(summary)
	if err != nil {
		t.Fatalf("Failed to add summary: %v", err)
	}

	// Test deleting existing info
	err = server.deleteInfoByID(infoID)
	if err != nil {
		t.Errorf("deleteInfoByID() unexpected error: %v", err)
	}

	// Verify file was deleted
	if _, err := os.Stat(infoPath); !os.IsNotExist(err) {
		t.Error("Expected info file to be deleted")
	}

	// Test deleting non-existent info
	err = server.deleteInfoByID("non-existent")
	if err == nil {
		t.Error("deleteInfoByID() expected error for non-existent info, got nil")
	}
}

func TestUpdateSummaryByRequestID(t *testing.T) {
	server, _ := createTestServer(t)

	// First add a summary with request ID
	summary := createTestImageSummary()
	summary.RequestID = "test-request-123"
	err := server.addSummaryToFile(summary)
	if err != nil {
		t.Fatalf("Failed to add summary: %v", err)
	}

	// Test updating existing summary
	updatedSummary := summary
	updatedSummary.Status = "completed"
	err = server.updateSummaryByRequestID(summary.RequestID, updatedSummary)
	if err != nil {
		t.Errorf("updateSummaryByRequestID() unexpected error: %v", err)
	}

	// Test updating non-existent request ID
	err = server.updateSummaryByRequestID("non-existent", updatedSummary)
	if err == nil {
		t.Error("updateSummaryByRequestID() expected error for non-existent request ID, got nil")
	}
}

func TestImageInfoToSummary(t *testing.T) {
	server, _ := createTestServer(t)

	imageInfo := createTestImageInfo()
	summary := server.imageInfoToSummary(imageInfo)

	if summary.ImageID != imageInfo.ImageID {
		t.Errorf("Expected ImageID %s, got %s", imageInfo.ImageID, summary.ImageID)
	}

	if summary.ImageTag != imageInfo.ImageTag {
		t.Errorf("Expected ImageTag %s, got %s", imageInfo.ImageTag, summary.ImageTag)
	}

	if summary.ImageSize != imageInfo.ImageSize {
		t.Errorf("Expected ImageSize %d, got %d", imageInfo.ImageSize, summary.ImageSize)
	}

	if summary.Architecture != imageInfo.Architecture {
		t.Errorf("Expected Architecture %s, got %s", imageInfo.Architecture, summary.Architecture)
	}
}

func TestValidateContextDir(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Test valid directory (tmpDir should be allowed)
	validPath, err := server.validateContextDir(tmpDir)
	if err != nil {
		t.Errorf("validateContextDir() unexpected error: %v", err)
	}

	// The function returns the absolute path, so we need to compare with the absolute path
	absTmpDir, _ := filepath.Abs(tmpDir)
	if validPath != absTmpDir {
		t.Errorf("Expected path %s, got %s", absTmpDir, validPath)
	}

	// Test non-existent directory
	_, err = server.validateContextDir("non-existent")
	if err == nil {
		t.Error("validateContextDir() expected error for non-existent directory, got nil")
	}

	// Test file instead of directory
	filePath := filepath.Join(tmpDir, "test-file")
	os.WriteFile(filePath, []byte("test"), 0644)
	_, err = server.validateContextDir(filePath)
	if err == nil {
		t.Error("validateContextDir() expected error for file, got nil")
	}

	// Test sensitive directory (should be blocked)
	_, err = server.validateContextDir("/var")
	if err == nil {
		t.Error("validateContextDir() expected error for sensitive directory, got nil")
	}
}

func TestHandleBuildContextPreview(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Create a test Dockerfile
	dockerfilePath := filepath.Join(tmpDir, "Dockerfile")
	os.WriteFile(dockerfilePath, []byte("FROM ubuntu:20.04\nRUN echo 'test'"), 0644)

	// Test valid request
	previewReq := BuildContextPreviewRequest{
		ContextDir:          tmpDir,
		DockerignoreContent: "node_modules\n.git",
	}
	reqBody, _ := json.Marshal(previewReq)
	req := httptest.NewRequest("POST", "/api/buildcontext/preview", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()
	server.handleBuildContextPreview(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Test invalid request body
	req = httptest.NewRequest("POST", "/api/buildcontext/preview", strings.NewReader("invalid json"))
	w = httptest.NewRecorder()
	server.handleBuildContextPreview(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Test invalid context directory
	previewReq.ContextDir = "non-existent"
	reqBody, _ = json.Marshal(previewReq)
	req = httptest.NewRequest("POST", "/api/buildcontext/preview", bytes.NewReader(reqBody))
	w = httptest.NewRecorder()
	server.handleBuildContextPreview(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandleBuildContextRead(t *testing.T) {
	server, tmpDir := createTestServer(t)

	// Create a .dockerignore file
	dockerignorePath := filepath.Join(tmpDir, ".dockerignore")
	os.WriteFile(dockerignorePath, []byte("node_modules\n.git\n*.log"), 0644)

	// Test valid request
	readReq := BuildContextReadRequest{
		ContextDir: tmpDir,
	}
	reqBody, _ := json.Marshal(readReq)
	req := httptest.NewRequest("POST", "/api/buildcontext/read", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()
	server.handleBuildContextRead(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Test invalid request body
	req = httptest.NewRequest("POST", "/api/buildcontext/read", strings.NewReader("invalid json"))
	w = httptest.NewRecorder()
	server.handleBuildContextRead(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Test non-existent directory
	readReq.ContextDir = "non-existent"
	reqBody, _ = json.Marshal(readReq)
	req = httptest.NewRequest("POST", "/api/buildcontext/read", bytes.NewReader(reqBody))
	w = httptest.NewRecorder()
	server.handleBuildContextRead(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
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
		t.Fatalf("json.UnMarshal() unexpected error: %v", err)
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

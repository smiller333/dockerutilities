package analyzer

import (
	"archive/tar"
	"compress/gzip"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseImageNameAndSource(t *testing.T) {
	tests := []struct {
		name           string
		fullImageName  string
		expectedTag    string
		expectedSource string
	}{
		{
			name:           "DockerHub image with tag",
			fullImageName:  "nginx:latest",
			expectedTag:    "nginx:latest",
			expectedSource: "",
		},
		{
			name:           "DockerHub image without tag",
			fullImageName:  "nginx",
			expectedTag:    "nginx",
			expectedSource: "",
		},
		{
			name:           "DockerHub image with docker.io prefix",
			fullImageName:  "docker.io/nginx:latest",
			expectedTag:    "nginx:latest",
			expectedSource: "",
		},
		{
			name:           "DockerHub official library image",
			fullImageName:  "library/nginx:latest",
			expectedTag:    "library/nginx:latest",
			expectedSource: "",
		},
		{
			name:           "DockerHub user image",
			fullImageName:  "user/myapp:v1.0",
			expectedTag:    "user/myapp:v1.0",
			expectedSource: "",
		},
		{
			name:           "GitLab registry image",
			fullImageName:  "registry.gitlab.com/yumbrands/phus/web/web2-frontend/web2-app-image:v2.10.3154",
			expectedTag:    "web2-app-image:v2.10.3154",
			expectedSource: "registry.gitlab.com/yumbrands/phus/web/web2-frontend",
		},
		{
			name:           "GitLab registry image without tag",
			fullImageName:  "registry.gitlab.com/user/project/image",
			expectedTag:    "image",
			expectedSource: "registry.gitlab.com/user/project",
		},
		{
			name:           "AWS ECR image",
			fullImageName:  "123456789012.dkr.ecr.us-east-1.amazonaws.com/my-app:latest",
			expectedTag:    "my-app:latest",
			expectedSource: "123456789012.dkr.ecr.us-east-1.amazonaws.com",
		},
		{
			name:           "Google Container Registry",
			fullImageName:  "gcr.io/project-id/image-name:tag",
			expectedTag:    "image-name:tag",
			expectedSource: "gcr.io/project-id",
		},
		{
			name:           "Harbor registry",
			fullImageName:  "harbor.company.com/project/app:v1.0.0",
			expectedTag:    "app:v1.0.0",
			expectedSource: "harbor.company.com/project",
		},
		{
			name:           "Complex path with multiple levels",
			fullImageName:  "registry.example.com/org/team/project/service/app:v1.2.3",
			expectedTag:    "app:v1.2.3",
			expectedSource: "registry.example.com/org/team/project/service",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTag, gotSource := parseImageNameAndSource(tt.fullImageName)
			if gotTag != tt.expectedTag {
				t.Errorf("parseImageNameAndSource() gotTag = %v, want %v", gotTag, tt.expectedTag)
			}
			if gotSource != tt.expectedSource {
				t.Errorf("parseImageNameAndSource() gotSource = %v, want %v", gotSource, tt.expectedSource)
			}
		})
	}
}

func TestSafeTarExtraction_ValidTar(t *testing.T) {
	// Create a temporary directory for extraction
	tmpDir := t.TempDir()
	extractDir := filepath.Join(tmpDir, "extracted")

	// Create a simple tar file
	tarPath := filepath.Join(tmpDir, "test.tar")
	createTestTar(t, tarPath, []tarEntry{
		{name: "file1.txt", content: "content1", isDir: false},
		{name: "dir1", content: "", isDir: true},
		{name: "dir1/file2.txt", content: "content2", isDir: false},
	})

	// Test extraction
	config := &Config{MaxFileSize: 100 * 1024 * 1024}
	err := SafeTarExtraction(tarPath, extractDir, config)
	if err != nil {
		t.Fatalf("SafeTarExtraction failed: %v", err)
	}

	// Verify extracted files
	expectedFiles := []string{
		filepath.Join(extractDir, "file1.txt"),
		filepath.Join(extractDir, "dir1"),
		filepath.Join(extractDir, "dir1", "file2.txt"),
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist", file)
		}
	}

	// Verify file contents
	content1, err := os.ReadFile(filepath.Join(extractDir, "file1.txt"))
	if err != nil {
		t.Fatalf("Failed to read file1.txt: %v", err)
	}
	if string(content1) != "content1" {
		t.Errorf("Expected content 'content1', got '%s'", string(content1))
	}

	content2, err := os.ReadFile(filepath.Join(extractDir, "dir1", "file2.txt"))
	if err != nil {
		t.Fatalf("Failed to read file2.txt: %v", err)
	}
	if string(content2) != "content2" {
		t.Errorf("Expected content 'content2', got '%s'", string(content2))
	}
}

func TestSafeTarExtraction_GzippedTar(t *testing.T) {
	// Create a temporary directory for extraction
	tmpDir := t.TempDir()
	extractDir := filepath.Join(tmpDir, "extracted")

	// Create a gzipped tar file
	tarPath := filepath.Join(tmpDir, "test.tar.gz")
	createTestGzippedTar(t, tarPath, []tarEntry{
		{name: "file1.txt", content: "content1", isDir: false},
		{name: "dir1", content: "", isDir: true},
		{name: "dir1/file2.txt", content: "content2", isDir: false},
	})

	// Test extraction
	config := &Config{MaxFileSize: 100 * 1024 * 1024}
	err := SafeTarExtraction(tarPath, extractDir, config)
	if err != nil {
		t.Fatalf("SafeTarExtraction failed: %v", err)
	}

	// Verify extracted files
	expectedFiles := []string{
		filepath.Join(extractDir, "file1.txt"),
		filepath.Join(extractDir, "dir1"),
		filepath.Join(extractDir, "dir1", "file2.txt"),
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist", file)
		}
	}
}

func TestSafeTarExtraction_PathTraversal(t *testing.T) {
	// Create a temporary directory for extraction
	tmpDir := t.TempDir()
	extractDir := filepath.Join(tmpDir, "extracted")

	// Create a tar file with path traversal attempt
	tarPath := filepath.Join(tmpDir, "malicious.tar")
	createTestTar(t, tarPath, []tarEntry{
		{name: "../../../etc/passwd", content: "malicious", isDir: false},
	})

	// Test extraction - should fail due to path traversal
	config := &Config{MaxFileSize: 100 * 1024 * 1024}
	err := SafeTarExtraction(tarPath, extractDir, config)
	if err == nil {
		t.Error("Expected error for path traversal attempt, got nil")
	}

	// Verify the malicious file was not created
	maliciousPath := filepath.Join(extractDir, "../../../etc/passwd")
	if _, err := os.Stat(maliciousPath); !os.IsNotExist(err) {
		t.Error("Malicious file should not exist")
	}
}

func TestSafeTarExtraction_NonexistentFile(t *testing.T) {
	tmpDir := t.TempDir()
	extractDir := filepath.Join(tmpDir, "extracted")

	config := &Config{MaxFileSize: 100 * 1024 * 1024}
	err := SafeTarExtraction("/nonexistent/file.tar", extractDir, config)
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestSafeTarExtraction_InvalidGzip(t *testing.T) {
	// Create a temporary directory for extraction
	tmpDir := t.TempDir()
	extractDir := filepath.Join(tmpDir, "extracted")

	// Create a file with .tar.gz extension but invalid gzip content
	tarPath := filepath.Join(tmpDir, "invalid.tar.gz")
	err := os.WriteFile(tarPath, []byte("not a gzip file"), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid gzip file: %v", err)
	}

	config := &Config{MaxFileSize: 100 * 1024 * 1024}
	err = SafeTarExtraction(tarPath, extractDir, config)
	if err == nil {
		t.Error("Expected error for invalid gzip file, got nil")
	}
}

func TestExtractSecureFile(t *testing.T) {
	tmpDir := t.TempDir()
	targetPath := filepath.Join(tmpDir, "test.txt")
	content := "test content"

	reader := strings.NewReader(content)
	err := extractSecureFile(reader, targetPath, 0644, 100*1024*1024)
	if err != nil {
		t.Fatalf("extractSecureFile failed: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		t.Error("Expected file to be created")
	}

	// Verify content
	fileContent, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("Failed to read created file: %v", err)
	}
	if string(fileContent) != content {
		t.Errorf("Expected content '%s', got '%s'", content, string(fileContent))
	}
}

func TestSecureFileCreate(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		path        string
		expectError bool
	}{
		{
			name:        "valid file path",
			path:        filepath.Join(tmpDir, "test.txt"),
			expectError: false,
		},
		{
			name:        "file in subdirectory",
			path:        filepath.Join(tmpDir, "subdir", "test.txt"),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create subdirectory if needed
			if strings.Contains(tt.path, "subdir") {
				subdir := filepath.Dir(tt.path)
				err := os.MkdirAll(subdir, 0755)
				if err != nil {
					t.Fatalf("Failed to create subdirectory: %v", err)
				}
			}

			file, err := secureFileCreate(tt.path)
			if tt.expectError && err == nil {
				t.Errorf("Expected error for path '%s', got nil", tt.path)
			}
			if !tt.expectError {
				if err != nil {
					t.Errorf("Expected no error for path '%s', got: %v", tt.path, err)
				} else {
					if err := file.Close(); err != nil {
						t.Errorf("Failed to close file: %v", err)
					}
					// Verify file was created
					if _, err := os.Stat(tt.path); os.IsNotExist(err) {
						t.Errorf("Expected file '%s' to be created", tt.path)
					}
				}
			}
		})
	}
}

func TestCleanupTemporaryFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create some temporary files
	file1 := filepath.Join(tmpDir, "temp1.txt")
	file2 := filepath.Join(tmpDir, "temp2.txt")
	dir1 := filepath.Join(tmpDir, "tempdir")

	err := os.WriteFile(file1, []byte("content1"), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file 1: %v", err)
	}

	err = os.WriteFile(file2, []byte("content2"), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file 2: %v", err)
	}

	err = os.Mkdir(dir1, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	result := &AnalysisResult{
		SavedTarPath:    file1,
		ExtractedPath:   dir1,
		ContainerFSPath: file2,
	}

	// Test cleanup
	err = cleanupTemporaryFiles(result)
	if err != nil {
		t.Fatalf("cleanupTemporaryFiles failed: %v", err)
	}

	// Verify files were removed
	if _, err := os.Stat(file1); !os.IsNotExist(err) {
		t.Error("Expected temp file 1 to be removed")
	}
	// Note: ContainerFSPath is not cleaned up by cleanupTemporaryFiles
	// Note: ExtractedPath contents are cleaned but directory remains
	if _, err := os.Stat(dir1); os.IsNotExist(err) {
		t.Error("Expected temp directory to remain (only contents cleaned)")
	}
}

func TestCleanupTemporaryFiles_NonexistentFiles(t *testing.T) {
	result := &AnalysisResult{
		SavedTarPath:    "/nonexistent/file1.txt",
		ExtractedPath:   "/nonexistent/dir",
		ContainerFSPath: "/nonexistent/file2.txt",
	}

	// Test cleanup with nonexistent files - should error for extracted path
	err := cleanupTemporaryFiles(result)
	if err == nil {
		t.Error("Expected error for nonexistent extracted directory, got nil")
	}
}

func TestCleanupExtractedDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	extractDir := filepath.Join(tmpDir, "extracted")

	// Create a directory with some files
	err := os.Mkdir(extractDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create extract directory: %v", err)
	}

	file1 := filepath.Join(extractDir, "file1.txt")
	err = os.WriteFile(file1, []byte("content1"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file in extract directory: %v", err)
	}

	subdir := filepath.Join(extractDir, "subdir")
	err = os.Mkdir(subdir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	file2 := filepath.Join(subdir, "file2.txt")
	err = os.WriteFile(file2, []byte("content2"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file in subdirectory: %v", err)
	}

	// Test cleanup
	err = cleanupExtractedDirectory(extractDir)
	if err != nil {
		t.Fatalf("cleanupExtractedDirectory failed: %v", err)
	}

	// Verify directory contents were cleaned but directory remains
	if _, err := os.Stat(extractDir); os.IsNotExist(err) {
		t.Error("Expected extracted directory to remain (only contents cleaned)")
	}

	// Verify files were removed
	if _, err := os.Stat(file1); !os.IsNotExist(err) {
		t.Error("Expected file1 to be removed")
	}
	if _, err := os.Stat(file2); !os.IsNotExist(err) {
		t.Error("Expected file2 to be removed")
	}
	if _, err := os.Stat(subdir); !os.IsNotExist(err) {
		t.Error("Expected subdir to be removed")
	}
}

func TestCleanupExtractedDirectory_Nonexistent(t *testing.T) {
	// Test cleanup with nonexistent directory - should error
	err := cleanupExtractedDirectory("/nonexistent/directory")
	if err == nil {
		t.Error("Expected error for nonexistent directory, got nil")
	}
}

func TestWriteImageIntoToFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "image_info.json")

	// Create test image info
	info := ImageInfo{
		ImageID:      "sha256:abc123",
		ImageTag:     "test:latest",
		Architecture: "amd64",
		OS:           "linux",
		ImageSize:    1024000,
		Layers:       []string{"layer1", "layer2"},
		Created:      "2023-01-01T00:00:00Z",
		Author:       "test author",
	}

	// Test writing to file
	err := writeImageIntoToFile(info, filePath)
	if err != nil {
		t.Fatalf("writeImageIntoToFile failed: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("Expected image info file to be created")
	}

	// Verify file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read image info file: %v", err)
	}

	// Check that the content contains expected fields
	contentStr := string(content)
	expectedFields := []string{
		"image_id",
		"image_tag",
		"architecture",
		"os",
		"image_size",
		"layers",
		"created",
		"author",
	}

	for _, field := range expectedFields {
		if !strings.Contains(contentStr, field) {
			t.Errorf("Expected image info to contain field '%s'", field)
		}
	}
}

func TestWriteImageIntoToFile_InvalidPath(t *testing.T) {
	info := ImageInfo{
		ImageID:  "sha256:abc123",
		ImageTag: "test:latest",
	}

	// Test writing to invalid path
	err := writeImageIntoToFile(info, "/nonexistent/directory/file.json")
	if err == nil {
		t.Error("Expected error for invalid path, got nil")
	}
}

func TestExtractLayersFromManifest(t *testing.T) {
	tmpDir := t.TempDir()
	manifestPath := filepath.Join(tmpDir, "manifest.json")

	// Create a test manifest file
	manifestContent := `[
		{
			"Config": "config.json",
			"RepoTags": ["test:latest"],
			"Layers": [
				"sha256:layer1",
				"sha256:layer2",
				"sha256:layer3"
			]
		}
	]`

	err := os.WriteFile(manifestPath, []byte(manifestContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create manifest file: %v", err)
	}

	// Test extracting layers
	layers, err := extractLayersFromManifest(manifestPath)
	if err != nil {
		t.Fatalf("extractLayersFromManifest failed: %v", err)
	}

	// Verify layers
	expectedLayers := []string{"sha256:layer1", "sha256:layer2", "sha256:layer3"}
	if len(layers) != len(expectedLayers) {
		t.Errorf("Expected %d layers, got %d", len(expectedLayers), len(layers))
	}

	for i, expected := range expectedLayers {
		if i >= len(layers) {
			t.Errorf("Missing layer at index %d", i)
			continue
		}
		if layers[i] != expected {
			t.Errorf("Expected layer %s, got %s", expected, layers[i])
		}
	}
}

func TestExtractLayersFromManifest_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	manifestPath := filepath.Join(tmpDir, "invalid_manifest.json")

	// Create an invalid JSON file
	invalidContent := `{ invalid json }`

	err := os.WriteFile(manifestPath, []byte(invalidContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid manifest file: %v", err)
	}

	// Test extracting layers from invalid JSON
	_, err = extractLayersFromManifest(manifestPath)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestExtractLayersFromManifest_NonexistentFile(t *testing.T) {
	// Test extracting layers from nonexistent file
	_, err := extractLayersFromManifest("/nonexistent/manifest.json")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestExtractLayersFromManifest_EmptyLayers(t *testing.T) {
	tmpDir := t.TempDir()
	manifestPath := filepath.Join(tmpDir, "empty_manifest.json")

	// Create a manifest with empty layers
	manifestContent := `[
		{
			"Config": "config.json",
			"RepoTags": ["test:latest"],
			"Layers": []
		}
	]`

	err := os.WriteFile(manifestPath, []byte(manifestContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create empty manifest file: %v", err)
	}

	// Test extracting layers
	layers, err := extractLayersFromManifest(manifestPath)
	if err != nil {
		t.Fatalf("extractLayersFromManifest failed: %v", err)
	}

	// Verify empty layers
	if len(layers) != 0 {
		t.Errorf("Expected 0 layers, got %d", len(layers))
	}
}

// Helper types and functions for testing

type tarEntry struct {
	name    string
	content string
	isDir   bool
}

func createTestTar(t *testing.T, tarPath string, entries []tarEntry) {
	file, err := os.Create(tarPath)
	if err != nil {
		t.Fatalf("Failed to create tar file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Errorf("Failed to close file: %v", err)
		}
	}()

	tw := tar.NewWriter(file)
	defer func() {
		if err := tw.Close(); err != nil {
			t.Errorf("Failed to close tar writer: %v", err)
		}
	}()

	for _, entry := range entries {
		header := &tar.Header{
			Name: entry.name,
			Mode: 0644,
		}

		if entry.isDir {
			header.Typeflag = tar.TypeDir
			header.Mode = 0755
		} else {
			header.Typeflag = tar.TypeReg
			header.Size = int64(len(entry.content))
		}

		err := tw.WriteHeader(header)
		if err != nil {
			t.Fatalf("Failed to write tar header: %v", err)
		}

		if !entry.isDir {
			_, err = tw.Write([]byte(entry.content))
			if err != nil {
				t.Fatalf("Failed to write tar content: %v", err)
			}
		}
	}
}

func createTestGzippedTar(t *testing.T, tarPath string, entries []tarEntry) {
	file, err := os.Create(tarPath)
	if err != nil {
		t.Fatalf("Failed to create gzipped tar file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Errorf("Failed to close file: %v", err)
		}
	}()

	gw := gzip.NewWriter(file)
	defer func() {
		if err := gw.Close(); err != nil {
			t.Errorf("Failed to close gzip writer: %v", err)
		}
	}()

	tw := tar.NewWriter(gw)
	defer func() {
		if err := tw.Close(); err != nil {
			t.Errorf("Failed to close tar writer: %v", err)
		}
	}()

	for _, entry := range entries {
		header := &tar.Header{
			Name: entry.name,
			Mode: 0644,
		}

		if entry.isDir {
			header.Typeflag = tar.TypeDir
			header.Mode = 0755
		} else {
			header.Typeflag = tar.TypeReg
			header.Size = int64(len(entry.content))
		}

		err := tw.WriteHeader(header)
		if err != nil {
			t.Fatalf("Failed to write tar header: %v", err)
		}

		if !entry.isDir {
			_, err = tw.Write([]byte(entry.content))
			if err != nil {
				t.Fatalf("Failed to write tar content: %v", err)
			}
		}
	}
}

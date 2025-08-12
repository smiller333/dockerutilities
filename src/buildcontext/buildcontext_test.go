package buildcontext

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadDockerignore(t *testing.T) {
	tests := []struct {
		name           string
		dockerignore   string
		expectedError  bool
		expectedResult string
	}{
		{
			name:           "empty dockerignore",
			dockerignore:   "",
			expectedError:  false,
			expectedResult: "",
		},
		{
			name: "simple patterns",
			dockerignore: `*.log
.git
node_modules/
`,
			expectedError:  false,
			expectedResult: "*.log\n.git\nnode_modules/\n",
		},
		{
			name: "with comments and empty lines",
			dockerignore: `# Build files
*.log

# Dependencies  
node_modules/
.git

# Temporary files
*.tmp
`,
			expectedError:  false,
			expectedResult: "# Build files\n*.log\n\n# Dependencies  \nnode_modules/\n.git\n\n# Temporary files\n*.tmp\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary directory
			tempDir := t.TempDir()

			// Create .dockerignore file if content provided
			if tt.dockerignore != "" {
				dockerignorePath := filepath.Join(tempDir, ".dockerignore")
				err := os.WriteFile(dockerignorePath, []byte(tt.dockerignore), 0644)
				if err != nil {
					t.Fatalf("Failed to create .dockerignore file: %v", err)
				}
			}

			result, err := ReadDockerignore(tempDir)

			if tt.expectedError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tt.expectedResult {
				t.Errorf("Expected result %q, got %q", tt.expectedResult, result)
			}
		})
	}
}

func TestParseDockerignorePatterns(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "empty content",
			content:  "",
			expected: nil,
		},
		{
			name:     "simple patterns",
			content:  "*.log\n.git\nnode_modules/",
			expected: []string{"*.log", ".git", "node_modules/"},
		},
		{
			name: "with comments and empty lines",
			content: `# Build files
*.log

# Dependencies
node_modules/
.git

# Temporary files
*.tmp`,
			expected: []string{"*.log", "node_modules/", ".git", "*.tmp"},
		},
		{
			name:     "whitespace handling",
			content:  "  *.log  \n\t.git\t\n  node_modules/  ",
			expected: []string{"*.log", ".git", "node_modules/"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseDockerignorePatterns(tt.content)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestShouldIgnore(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		patterns []string
		expected bool
	}{
		{
			name:     "no patterns",
			path:     "file.txt",
			patterns: []string{},
			expected: false,
		},
		{
			name:     "exact match",
			path:     "file.txt",
			patterns: []string{"file.txt"},
			expected: true,
		},
		{
			name:     "wildcard match",
			path:     "file.log",
			patterns: []string{"*.log"},
			expected: true,
		},
		{
			name:     "directory pattern",
			path:     "node_modules/package/file.js",
			patterns: []string{"node_modules/"},
			expected: true,
		},
		{
			name:     "directory exact match",
			path:     "node_modules",
			patterns: []string{"node_modules/"},
			expected: true,
		},
		{
			name:     "negation pattern",
			path:     "important.log",
			patterns: []string{"*.log", "!important.log"},
			expected: false,
		},
		{
			name:     "complex negation",
			path:     "logs/debug.log",
			patterns: []string{"logs/", "!logs/debug.log"},
			expected: false,
		},
		{
			name:     "no match",
			path:     "file.txt",
			patterns: []string{"*.log", ".git"},
			expected: false,
		},
		{
			name:     "nested directory",
			path:     "src/main/java/App.java",
			patterns: []string{"src/"},
			expected: true,
		},
		{
			name:     "double asterisk pattern",
			path:     "deep/nested/test.tmp",
			patterns: []string{"**/*.tmp"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shouldIgnore(tt.path, tt.patterns)
			if result != tt.expected {
				t.Errorf("Expected %v for path %q with patterns %v, got %v",
					tt.expected, tt.path, tt.patterns, result)
			}
		})
	}
}

func TestMatchesPattern(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		pattern  string
		expected bool
	}{
		{
			name:     "exact match",
			path:     "file.txt",
			pattern:  "file.txt",
			expected: true,
		},
		{
			name:     "wildcard extension",
			path:     "file.log",
			pattern:  "*.log",
			expected: true,
		},
		{
			name:     "directory pattern with trailing slash",
			path:     "node_modules/package.json",
			pattern:  "node_modules/",
			expected: true,
		},
		{
			name:     "directory exact match",
			path:     "node_modules",
			pattern:  "node_modules/",
			expected: true,
		},
		{
			name:     "double asterisk prefix",
			path:     "deep/nested/file.tmp",
			pattern:  "**/*.tmp",
			expected: true,
		},
		{
			name:     "double asterisk middle",
			path:     "src/main/java/App.java",
			pattern:  "src/**/App.java",
			expected: true,
		},
		{
			name:     "no match",
			path:     "file.txt",
			pattern:  "*.log",
			expected: false,
		},
		{
			name:     "partial path match",
			path:     "logs/error.log",
			pattern:  "logs",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchesPattern(tt.path, tt.pattern)
			if result != tt.expected {
				t.Errorf("Expected %v for path %q with pattern %q, got %v",
					tt.expected, tt.path, tt.pattern, result)
			}
		})
	}
}

func TestComputeBuildContext(t *testing.T) {
	tests := []struct {
		name             string
		files            map[string]string // filename -> content
		dockerignore     string
		expectedIncluded []string // paths that should be included
		expectedExcluded []string // paths that should be excluded
	}{
		{
			name: "no dockerignore",
			files: map[string]string{
				"main.go":        "package main",
				"README.md":      "# Test",
				"logs/error.log": "error",
				"src/helper.go":  "package src",
			},
			dockerignore:     "",
			expectedIncluded: []string{"main.go", "README.md", "logs", "logs/error.log", "src", "src/helper.go"},
			expectedExcluded: []string{},
		},
		{
			name: "simple patterns",
			files: map[string]string{
				"main.go":        "package main",
				"README.md":      "# Test",
				"logs/error.log": "error",
				"logs/debug.log": "debug",
				"temp.tmp":       "temp",
			},
			dockerignore:     "*.log\n*.tmp\nREADME.md",
			expectedIncluded: []string{"main.go", "logs"},
			expectedExcluded: []string{"README.md", "temp.tmp", "logs/error.log", "logs/debug.log"},
		},
		{
			name: "directory exclusion",
			files: map[string]string{
				"main.go":               "package main",
				"node_modules/pkg/a.js": "module",
				"src/main.go":           "package src",
				"src/helper.go":         "package src",
			},
			dockerignore:     "node_modules/",
			expectedIncluded: []string{"main.go", "src", "src/main.go", "src/helper.go"},
			expectedExcluded: []string{"node_modules"}, // Only the directory itself, not its contents due to SkipDir
		},
		{
			name: "negation patterns",
			files: map[string]string{
				"app.log":       "app logs",
				"error.log":     "error logs",
				"important.log": "important logs",
				"debug.log":     "debug logs",
			},
			dockerignore:     "*.log\n!important.log",
			expectedIncluded: []string{"important.log"},
			expectedExcluded: []string{"app.log", "error.log", "debug.log"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary directory
			tempDir := t.TempDir()

			// Create test files
			for filename, content := range tt.files {
				filePath := filepath.Join(tempDir, filename)
				dir := filepath.Dir(filePath)

				// Create parent directories
				err := os.MkdirAll(dir, 0755)
				if err != nil {
					t.Fatalf("Failed to create directory %s: %v", dir, err)
				}

				// Create file
				err = os.WriteFile(filePath, []byte(content), 0644)
				if err != nil {
					t.Fatalf("Failed to create file %s: %v", filePath, err)
				}
			}

			// Run ComputeBuildContext
			result, excluded, err := ComputeBuildContext(tempDir, tt.dockerignore)
			if err != nil {
				t.Fatalf("ComputeBuildContext failed: %v", err)
			}

			// Check excluded paths (order may vary, so check contents)
			if !containsAllPaths(excluded, tt.expectedExcluded) || !containsAllPaths(tt.expectedExcluded, excluded) {
				t.Errorf("Expected excluded paths %v, got %v", tt.expectedExcluded, excluded)
			}

			// Check included paths by traversing the result structure
			included := extractIncludedPaths(result)
			if !containsAllPaths(included, tt.expectedIncluded) {
				t.Errorf("Expected included paths %v, got %v", tt.expectedIncluded, included)
			}

			// Verify no excluded paths appear in included
			for _, excludedPath := range excluded {
				if contains(included, excludedPath) {
					t.Errorf("Excluded path %s found in included paths", excludedPath)
				}
			}
		})
	}
}

func TestAddDirectoryToContext(t *testing.T) {
	root := &BuildDirectoryInfo{
		Path:        ".",
		Directories: make(map[string]*BuildDirectoryInfo),
		Files:       []BuildFileInfo{},
	}

	addDirectoryToContext(root, "src/main/java")

	// Check that directories were created
	if root.Directories["src"] == nil {
		t.Error("Expected 'src' directory to be created")
	}
	if root.Directories["src"].Directories["main"] == nil {
		t.Error("Expected 'src/main' directory to be created")
	}
	if root.Directories["src"].Directories["main"].Directories["java"] == nil {
		t.Error("Expected 'src/main/java' directory to be created")
	}

	// Check directory counts
	if root.DirCount != 1 {
		t.Errorf("Expected root DirCount to be 1, got %d", root.DirCount)
	}
}

func TestAddFileToContext(t *testing.T) {
	root := &BuildDirectoryInfo{
		Path:        ".",
		Directories: make(map[string]*BuildDirectoryInfo),
		Files:       []BuildFileInfo{},
	}

	// Add file in subdirectory
	addFileToContext(root, "src/main.go", 100)

	// Check that directory was created
	if root.Directories["src"] == nil {
		t.Error("Expected 'src' directory to be created")
	}

	// Check that file was added
	srcDir := root.Directories["src"]
	if len(srcDir.Files) != 1 {
		t.Errorf("Expected 1 file in src directory, got %d", len(srcDir.Files))
	}
	if srcDir.Files[0].Name != "main.go" {
		t.Errorf("Expected file name 'main.go', got %s", srcDir.Files[0].Name)
	}
	if srcDir.Files[0].Size != 100 {
		t.Errorf("Expected file size 100, got %d", srcDir.Files[0].Size)
	}

	// Check sizes were updated
	if root.Size != 100 {
		t.Errorf("Expected root size 100, got %d", root.Size)
	}
	if srcDir.Size != 100 {
		t.Errorf("Expected src size 100, got %d", srcDir.Size)
	}
}

// Helper functions for tests

func extractIncludedPaths(dir *BuildDirectoryInfo) []string {
	var paths []string

	// Add directories
	for name, subdir := range dir.Directories {
		var dirPath string
		if dir.Path == "." {
			dirPath = name
		} else {
			dirPath = dir.Path + "/" + name
		}
		paths = append(paths, dirPath)

		// Recursively add subdirectories
		subpaths := extractIncludedPaths(subdir)
		paths = append(paths, subpaths...)
	}

	// Add files
	for _, file := range dir.Files {
		var filePath string
		if dir.Path == "." {
			filePath = file.Name
		} else {
			filePath = dir.Path + "/" + file.Name
		}
		paths = append(paths, filePath)
	}

	return paths
}

func containsAllPaths(actual, expected []string) bool {
	for _, exp := range expected {
		if !contains(actual, exp) {
			return false
		}
	}
	return true
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

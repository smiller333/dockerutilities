// Package buildcontext provides functionality for analyzing and managing Docker build contexts,
// including handling .dockerignore patterns to determine which files should be included
// or excluded from a Docker build.
package buildcontext

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// BuildFileInfo represents a file within a build context with its metadata
type BuildFileInfo struct {
	Name string `json:"name"` // File name
	Size int64  `json:"size"` // File size in bytes
}

// BuildDirectoryInfo represents a directory within a build context with its contents and metadata
type BuildDirectoryInfo struct {
	Path        string                         `json:"path"`        // Directory path relative to context root
	Size        int64                          `json:"size"`        // Total size of directory contents in bytes
	FileCount   int                            `json:"file_count"`  // Number of files in this directory
	DirCount    int                            `json:"dir_count"`   // Number of subdirectories in this directory
	Files       []BuildFileInfo                `json:"files"`       // Files directly in this directory
	Directories map[string]*BuildDirectoryInfo `json:"directories"` // Subdirectories mapped by name
}

// ReadDockerignore reads the .dockerignore file from the specified directory and returns its content.
// If the file doesn't exist, it returns an empty string without error.
func ReadDockerignore(contextDir string) (string, error) {
	dockerignorePath := filepath.Join(contextDir, ".dockerignore")

	content, err := os.ReadFile(dockerignorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // Return empty string if .dockerignore doesn't exist
		}
		return "", fmt.Errorf("failed to read .dockerignore file: %w", err)
	}

	return string(content), nil
}

// ComputeBuildContext analyzes a directory and returns the build context structure,
// applying .dockerignore patterns to exclude files and directories.
// Returns the included files/directories structure and a list of excluded paths.
func ComputeBuildContext(contextDir string, dockerignoreContent string) (*BuildDirectoryInfo, []string, error) {
	// Clean and validate context directory
	contextDir = filepath.Clean(contextDir)

	// If dockerignoreContent is empty, try to read from file
	if dockerignoreContent == "" {
		var err error
		dockerignoreContent, err = ReadDockerignore(contextDir)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read .dockerignore: %w", err)
		}
	}

	// Parse .dockerignore patterns
	patterns := parseDockerignorePatterns(dockerignoreContent)

	// Initialize root directory info
	root := &BuildDirectoryInfo{
		Path:        ".",
		Directories: make(map[string]*BuildDirectoryInfo),
		Files:       []BuildFileInfo{},
	}

	var excludedPaths []string

	// Walk the directory tree
	err := filepath.WalkDir(contextDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Get relative path from context directory
		relPath, err := filepath.Rel(contextDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Skip the root directory itself
		if relPath == "." {
			return nil
		}

		// Check if path should be ignored
		if shouldIgnore(relPath, patterns) {
			excludedPaths = append(excludedPaths, relPath)
			if d.IsDir() {
				return filepath.SkipDir // Skip entire directory tree
			}
			return nil
		}

		// Get file info
		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("failed to get file info for %s: %w", relPath, err)
		}

		// Add to build context structure
		if d.IsDir() {
			addDirectoryToContext(root, relPath)
		} else {
			addFileToContext(root, relPath, info.Size())
		}

		return nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("failed to walk directory %s: %w", contextDir, err)
	}

	return root, excludedPaths, nil
}

// parseDockerignorePatterns parses the content of a .dockerignore file and returns
// a slice of patterns, handling comments and empty lines.
func parseDockerignorePatterns(content string) []string {
	var patterns []string
	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		patterns = append(patterns, line)
	}

	return patterns
}

// shouldIgnore determines if a path should be ignored based on .dockerignore patterns.
// This implements a simplified version of Docker's ignore logic.
func shouldIgnore(path string, patterns []string) bool {
	// Normalize path separators to forward slashes for consistent pattern matching
	normalizedPath := filepath.ToSlash(path)

	ignored := false

	for _, pattern := range patterns {
		// Handle negation patterns (starting with !)
		if strings.HasPrefix(pattern, "!") {
			negatePattern := strings.TrimPrefix(pattern, "!")
			if matchesPattern(normalizedPath, negatePattern) {
				ignored = false // Un-ignore this path
			}
			continue
		}

		// Check if path matches the ignore pattern
		if matchesPattern(normalizedPath, pattern) {
			ignored = true
		}
	}

	return ignored
}

// matchesPattern checks if a path matches a dockerignore pattern.
// This is a simplified implementation that handles basic glob patterns.
func matchesPattern(path, pattern string) bool {
	// Normalize pattern separators
	pattern = filepath.ToSlash(pattern)

	// Handle directory patterns (ending with /)
	if strings.HasSuffix(pattern, "/") {
		pattern = strings.TrimSuffix(pattern, "/")
		// For directory patterns, check if path starts with the pattern
		return strings.HasPrefix(path, pattern+"/") || path == pattern
	}

	// Handle patterns with **/ (match any number of directories)
	if strings.Contains(pattern, "**/") {
		parts := strings.Split(pattern, "**/")
		if len(parts) == 2 {
			prefix, suffix := parts[0], parts[1]

			// Remove trailing slash from prefix if present
			prefix = strings.TrimSuffix(prefix, "/")

			// Check if path matches the pattern
			if prefix == "" {
				// Pattern starts with **/
				matched, _ := filepath.Match(suffix, filepath.Base(path))
				if matched {
					return true
				}
				// Also check if any part of the path matches
				return strings.Contains(path, suffix)
			} else {
				// Pattern has prefix/**/suffix
				if strings.HasPrefix(path, prefix+"/") {
					remaining := strings.TrimPrefix(path, prefix+"/")
					matched, _ := filepath.Match(suffix, filepath.Base(remaining))
					if matched {
						return true
					}
					return strings.Contains(remaining, suffix)
				}
			}
		}
	}

	// Handle simple glob patterns
	matched, err := filepath.Match(pattern, filepath.Base(path))
	if err == nil && matched {
		return true
	}

	// Check if the full path matches
	matched, err = filepath.Match(pattern, path)
	if err == nil && matched {
		return true
	}

	// Check if any parent directory matches for directory-style patterns
	if strings.Contains(path, "/") {
		pathParts := strings.Split(path, "/")
		for i := range pathParts {
			partialPath := strings.Join(pathParts[:i+1], "/")
			matched, err := filepath.Match(pattern, partialPath)
			if err == nil && matched {
				return true
			}
		}
	}

	return false
}

// addDirectoryToContext adds a directory to the build context structure
func addDirectoryToContext(root *BuildDirectoryInfo, dirPath string) {
	pathParts := strings.Split(filepath.ToSlash(dirPath), "/")
	current := root

	for _, part := range pathParts {
		if part == "" {
			continue
		}

		if current.Directories[part] == nil {
			var newPath string
			if current.Path == "." {
				newPath = part
			} else {
				newPath = current.Path + "/" + part
			}
			current.Directories[part] = &BuildDirectoryInfo{
				Path:        newPath,
				Directories: make(map[string]*BuildDirectoryInfo),
				Files:       []BuildFileInfo{},
			}
			current.DirCount++
		}
		current = current.Directories[part]
	}
}

// addFileToContext adds a file to the build context structure
func addFileToContext(root *BuildDirectoryInfo, filePath string, size int64) {
	dir := filepath.Dir(filePath)
	filename := filepath.Base(filePath)

	// Navigate to the parent directory
	var current *BuildDirectoryInfo
	if dir == "." {
		current = root
	} else {
		// Ensure parent directory exists
		addDirectoryToContext(root, dir)

		// Navigate to the directory
		pathParts := strings.Split(filepath.ToSlash(dir), "/")
		current = root
		for _, part := range pathParts {
			if part == "" {
				continue
			}
			current = current.Directories[part]
		}
	}

	// Add file to directory
	current.Files = append(current.Files, BuildFileInfo{
		Name: filename,
		Size: size,
	})
	current.FileCount++

	// Update size recursively up the tree
	updateSizeUpTree(root, filePath, size)
}

// updateSizeUpTree updates the size of all parent directories when a file is added
func updateSizeUpTree(root *BuildDirectoryInfo, filePath string, size int64) {
	dir := filepath.Dir(filePath)

	// Update root
	root.Size += size

	if dir == "." {
		return
	}

	// Update parent directories
	pathParts := strings.Split(filepath.ToSlash(dir), "/")
	current := root

	for _, part := range pathParts {
		if part == "" {
			continue
		}
		if current.Directories[part] != nil {
			current.Directories[part].Size += size
			current = current.Directories[part]
		}
	}
}

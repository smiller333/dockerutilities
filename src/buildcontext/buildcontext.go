// Package buildcontext provides functionality for analyzing and managing Docker build contexts,
// including handling .dockerignore patterns to determine which files should be included
// or excluded from a Docker build.
// Copyright (c) 2025 Docker Utils Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

package buildcontext

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/moby/patternmatcher"
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
	return ComputeBuildContextWithOptions(contextDir, dockerignoreContent, false)
}

// ComputeBuildContextWithOptions analyzes a directory with additional options.
// If useEmptyContent is true, empty dockerignoreContent will be treated as explicitly empty
// rather than falling back to reading from the .dockerignore file.
func ComputeBuildContextWithOptions(contextDir string, dockerignoreContent string, useEmptyContent bool) (*BuildDirectoryInfo, []string, error) {
	// Clean and validate context directory
	contextDir = filepath.Clean(contextDir)

	// If dockerignoreContent is empty and useEmptyContent is false, try to read from file
	if dockerignoreContent == "" && !useEmptyContent {
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

	excludedPaths := make([]string, 0) // Initialize as empty slice instead of nil

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
// a slice of patterns, handling comments and empty lines with proper preprocessing
// as specified in Docker documentation.
func parseDockerignorePatterns(content string) []string {
	var patterns []string
	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Apply preprocessing step as mentioned in Docker docs:
		// Use Go's filepath.Clean to trim whitespace and remove . and ..
		line = filepath.Clean(line)

		// Skip the special "." pattern for historical reasons (Docker behavior)
		if line == "." {
			continue
		}

		// Lines that are blank after preprocessing are ignored
		if line != "" {
			patterns = append(patterns, line)
		}
	}

	return patterns
}

// shouldIgnore determines if a path should be ignored based on .dockerignore patterns.
// This uses the official moby/patternmatcher library for full Docker compatibility.
func shouldIgnore(path string, patterns []string) bool {
	// If no patterns, don't ignore anything
	if len(patterns) == 0 {
		return false
	}

	// Create pattern matcher using the official moby implementation
	pm, err := patternmatcher.New(patterns)
	if err != nil {
		// If pattern compilation fails, fall back to not ignoring
		// In a production system, you might want to log this error
		return false
	}

	// Use the pattern matcher to check if the path should be excluded
	matches, err := pm.Matches(path)
	if err != nil {
		// If matching fails, fall back to not ignoring
		return false
	}

	return matches
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

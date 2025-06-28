// Package analyzer provides functionality for analyzing Dockerfiles.
package analyzer

import (
	"fmt"
	"os"
	"path/filepath"
)

// AnalysisResult contains the results of analyzing a Dockerfile
type AnalysisResult struct {
	Path         string
	AbsolutePath string
	Size         int
	Content      string
}

// AnalyzeDockerfile reads and analyzes a Dockerfile at the specified path
func AnalyzeDockerfile(dockerfilePath string) (*AnalysisResult, error) {
	// Check if the file exists
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("dockerfile not found: %s", dockerfilePath)
	}

	// Get absolute path for better error reporting
	absPath, err := filepath.Abs(dockerfilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Read the Dockerfile
	content, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dockerfile %s: %w", absPath, err)
	}

	result := &AnalysisResult{
		Path:         dockerfilePath,
		AbsolutePath: absPath,
		Size:         len(content),
		Content:      string(content),
	}

	return result, nil
}

// PrintAnalysisResult prints a formatted summary of the analysis result
func PrintAnalysisResult(result *AnalysisResult) {
	fmt.Printf("Successfully read Dockerfile: %s\n", result.AbsolutePath)
	fmt.Printf("File size: %d bytes\n", result.Size)
	fmt.Printf("Content preview (first 500 characters):\n")
	fmt.Println("---")

	// Show a preview of the content
	preview := result.Content
	if len(preview) > 500 {
		preview = preview[:500] + "..."
	}
	fmt.Println(preview)
	fmt.Println("---")
}

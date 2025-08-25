package analyzer

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestPrintAnalysisResult_ImageAnalysis(t *testing.T) {
	tests := []struct {
		name            string
		result          *AnalysisResult
		showBuildOutput bool
		expectedOutput  []string
	}{
		{
			name: "successful image analysis with all fields",
			result: &AnalysisResult{
				IsImageAnalysis:    true,
				ImageTag:           "nginx:latest",
				Pulled:             true,
				ImageSource:        "docker.io",
				BuildSuccess:       true,
				ImageSize:          1024000,
				LayerCount:         5,
				Architecture:       "amd64",
				OS:                 "linux",
				Created:            "2023-01-01T00:00:00Z",
				Author:             "nginx team",
				SaveSuccess:        true,
				SavedTarPath:       "/tmp/image.tar",
				ExtractSuccess:     true,
				ExtractedPath:      "/tmp/extracted",
				ContainerSuccess:   true,
				ContainerID:        "abc123",
				ContainerName:      "test-container",
				ContainerWarnings:  []string{"Warning 1", "Warning 2"},
				ContainerFSSuccess: true,
				ContainerFSPath:    "/tmp/container-fs",
			},
			showBuildOutput: false,
			expectedOutput: []string{
				"Successfully analyzed Docker image: nginx:latest",
				"Image Pulled: true",
				"Image Tag: nginx:latest",
				"Image Source: docker.io",
				"Image size: 1024000 bytes",
				"Image layers: 5",
				"Architecture: amd64",
				"OS: linux",
				"Created: 2023-01-01T00:00:00Z",
				"Author: nginx team",
				"Inspection status: SUCCESS",
				"Saved tar: /tmp/image.tar",
				"Save status: SUCCESS",
				"Extracted to: /tmp/extracted",
				"Extract status: SUCCESS",
				"Container ID: abc123",
				"Container Name: test-container",
				"Container creation status: SUCCESS",
				"Container warnings:",
				"  - Warning 1",
				"  - Warning 2",
				"Container filesystem copied to: /tmp/container-fs",
				"Container filesystem copy status: SUCCESS",
			},
		},
		{
			name: "image analysis with build failure",
			result: &AnalysisResult{
				IsImageAnalysis: true,
				ImageTag:        "invalid:image",
				Pulled:          false,
				ImageSource:     "docker.io",
				BuildSuccess:    false,
			},
			showBuildOutput: false,
			expectedOutput: []string{
				"Successfully analyzed Docker image: invalid:image",
				"Image Pulled: false",
				"Image Tag: invalid:image",
				"Image Source: docker.io",
				"Inspection status: FAILED",
			},
		},
		{
			name: "image analysis with save failure",
			result: &AnalysisResult{
				IsImageAnalysis: true,
				ImageTag:        "nginx:latest",
				Pulled:          true,
				ImageSource:     "docker.io",
				BuildSuccess:    true,
				ImageSize:       1024000,
				LayerCount:      5,
				Architecture:    "amd64",
				OS:              "linux",
				SaveSuccess:     false,
				SavedTarPath:    "/tmp/failed.tar",
			},
			showBuildOutput: false,
			expectedOutput: []string{
				"Successfully analyzed Docker image: nginx:latest",
				"Image Pulled: true",
				"Image Tag: nginx:latest",
				"Image Source: docker.io",
				"Image size: 1024000 bytes",
				"Image layers: 5",
				"Architecture: amd64",
				"OS: linux",
				"Inspection status: SUCCESS",
				"Save status: FAILED",
			},
		},
		{
			name: "image analysis with container failure",
			result: &AnalysisResult{
				IsImageAnalysis:  true,
				ImageTag:         "nginx:latest",
				Pulled:           true,
				ImageSource:      "docker.io",
				BuildSuccess:     true,
				ImageSize:        1024000,
				LayerCount:       5,
				Architecture:     "amd64",
				OS:               "linux",
				SaveSuccess:      true,
				SavedTarPath:     "/tmp/image.tar",
				ExtractSuccess:   true,
				ExtractedPath:    "/tmp/extracted",
				ContainerSuccess: false,
				ContainerID:      "failed123",
			},
			showBuildOutput: false,
			expectedOutput: []string{
				"Successfully analyzed Docker image: nginx:latest",
				"Image Pulled: true",
				"Image Tag: nginx:latest",
				"Image Source: docker.io",
				"Image size: 1024000 bytes",
				"Image layers: 5",
				"Architecture: amd64",
				"OS: linux",
				"Inspection status: SUCCESS",
				"Saved tar: /tmp/image.tar",
				"Save status: SUCCESS",
				"Extracted to: /tmp/extracted",
				"Extract status: SUCCESS",
				"Container creation status: FAILED",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function
			PrintAnalysisResult(tt.result, tt.showBuildOutput)

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

			// Check that all expected strings are present
			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain '%s', but got:\n%s", expected, output)
				}
			}
		})
	}
}

func TestPrintAnalysisResult_DockerfileAnalysis(t *testing.T) {
	tests := []struct {
		name            string
		result          *AnalysisResult
		showBuildOutput bool
		expectedOutput  []string
	}{
		{
			name: "successful Dockerfile analysis with build output",
			result: &AnalysisResult{
				IsImageAnalysis: false,
				AbsolutePath:    "/path/to/Dockerfile",
				DFSize:          1024,
				ImageTag:        "test:latest",
				BuildSuccess:    true,
				LayerCount:      3,
				ImageSize:       512000,
				BuildTime:       2.5,
				BuildOutput:     "Step 1/3: FROM ubuntu:20.04\nStep 2/3: RUN apt-get update\nStep 3/3: CMD [\"/bin/bash\"]",
			},
			showBuildOutput: true,
			expectedOutput: []string{
				"Successfully read Dockerfile: /path/to/Dockerfile",
				"File size: 1024 bytes",
				"Image layers: 3",
				"Image size: 512000 bytes",
				"Docker image: test:latest",
				"Build time: 2.500 seconds",
				"Build status: SUCCESS",
				"Build Output:",
				"---",
				"Step 1/3: FROM ubuntu:20.04",
				"Step 2/3: RUN apt-get update",
				"Step 3/3: CMD [\"/bin/bash\"]",
				"---",
			},
		},
		{
			name: "Dockerfile analysis with build failure",
			result: &AnalysisResult{
				IsImageAnalysis: false,
				AbsolutePath:    "/path/to/Dockerfile",
				DFSize:          1024,
				ImageTag:        "test:latest",
				BuildSuccess:    false,
				BuildTime:       1.2,
			},
			showBuildOutput: false,
			expectedOutput: []string{
				"Successfully read Dockerfile: /path/to/Dockerfile",
				"File size: 1024 bytes",
				"Build time: 1.200 seconds",
				"Build status: FAILED",
			},
		},
		{
			name: "Dockerfile analysis without image tag",
			result: &AnalysisResult{
				IsImageAnalysis: false,
				AbsolutePath:    "/path/to/Dockerfile",
				DFSize:          1024,
				ImageTag:        "",
				BuildSuccess:    false,
			},
			showBuildOutput: false,
			expectedOutput: []string{
				"Successfully read Dockerfile: /path/to/Dockerfile",
				"File size: 1024 bytes",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function
			PrintAnalysisResult(tt.result, tt.showBuildOutput)

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

			// Check that all expected strings are present
			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain '%s', but got:\n%s", expected, output)
				}
			}
		})
	}
}

func TestPrintAnalysisResult_EdgeCases(t *testing.T) {
	tests := []struct {
		name            string
		result          *AnalysisResult
		showBuildOutput bool
		expectedOutput  []string
	}{
		{
			name: "image analysis with empty optional fields",
			result: &AnalysisResult{
				IsImageAnalysis:    true,
				ImageTag:           "nginx:latest",
				Pulled:             true,
				ImageSource:        "docker.io",
				BuildSuccess:       true,
				ImageSize:          1024000,
				LayerCount:         5,
				Architecture:       "amd64",
				OS:                 "linux",
				Created:            "",
				Author:             "",
				SaveSuccess:        true,
				SavedTarPath:       "/tmp/image.tar",
				ExtractSuccess:     true,
				ExtractedPath:      "/tmp/extracted",
				ContainerSuccess:   true,
				ContainerID:        "abc123",
				ContainerName:      "",
				ContainerWarnings:  []string{},
				ContainerFSSuccess: true,
				ContainerFSPath:    "/tmp/container-fs",
			},
			showBuildOutput: false,
			expectedOutput: []string{
				"Successfully analyzed Docker image: nginx:latest",
				"Image Pulled: true",
				"Image Tag: nginx:latest",
				"Image Source: docker.io",
				"Image size: 1024000 bytes",
				"Image layers: 5",
				"Architecture: amd64",
				"OS: linux",
				"Inspection status: SUCCESS",
				"Saved tar: /tmp/image.tar",
				"Save status: SUCCESS",
				"Extracted to: /tmp/extracted",
				"Extract status: SUCCESS",
				"Container ID: abc123",
				"Container creation status: SUCCESS",
				"Container filesystem copied to: /tmp/container-fs",
				"Container filesystem copy status: SUCCESS",
			},
		},
		{
			name: "Dockerfile analysis with empty build output",
			result: &AnalysisResult{
				IsImageAnalysis: false,
				AbsolutePath:    "/path/to/Dockerfile",
				DFSize:          1024,
				ImageTag:        "test:latest",
				BuildSuccess:    true,
				LayerCount:      3,
				ImageSize:       512000,
				BuildTime:       2.5,
				BuildOutput:     "",
			},
			showBuildOutput: true,
			expectedOutput: []string{
				"Successfully read Dockerfile: /path/to/Dockerfile",
				"File size: 1024 bytes",
				"Image layers: 3",
				"Image size: 512000 bytes",
				"Docker image: test:latest",
				"Build time: 2.500 seconds",
				"Build status: SUCCESS",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function
			PrintAnalysisResult(tt.result, tt.showBuildOutput)

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

			// Check that all expected strings are present
			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain '%s', but got:\n%s", expected, output)
				}
			}
		})
	}
}

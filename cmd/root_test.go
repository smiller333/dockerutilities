package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/smiller333/dockerutilities/src/version"
	"github.com/spf13/cobra"
)

// TestInit tests that commands are properly initialized
func TestInit(t *testing.T) {
	// Verify that version command was added during package initialization
	commands := rootCmd.Commands()
	if len(commands) == 0 {
		t.Error("Expected commands to be added during initialization")
	}

	// Find version command
	var versionCommand *cobra.Command
	for _, cmd := range commands {
		if cmd.Use == "version" {
			versionCommand = cmd
			break
		}
	}

	if versionCommand == nil {
		t.Error("Expected version command to be added during initialization")
		return
	}

	// Verify version command properties
	if versionCommand.Short != "Print the version information of dockerutilities" {
		t.Errorf("Expected version command short description, got '%s'", versionCommand.Short)
	}
}

// TestExecute tests the Execute() function
func TestExecute(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "no args",
			args:    []string{},
			wantErr: false,
		},
		{
			name:    "version command",
			args:    []string{"version"},
			wantErr: false,
		},
		{
			name:    "help command",
			args:    []string{"--help"},
			wantErr: false,
		},
		{
			name:    "invalid command",
			args:    []string{"invalid"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout to verify output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Set up command with test args
			rootCmd.SetArgs(tt.args)

			// Execute command
			err := Execute()

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

			// Check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}

			// For version command, verify output contains version info
			if tt.name == "version command" && !tt.wantErr {
				output := buf.String()
				// The version command outputs the full version string, so just check it's not empty
				if strings.TrimSpace(output) == "" {
					t.Errorf("Version command output should not be empty, got: %s", output)
				}
			}

			// For help command, verify output contains help text
			if tt.name == "help command" && !tt.wantErr {
				output := buf.String()
				if !strings.Contains(output, "dockerutilities") {
					t.Errorf("Help command output should contain command name, got: %s", output)
				}
			}
		})
	}
}

// TestRootCommandProperties tests the root command properties
func TestRootCommandProperties(t *testing.T) {
	// Verify root command properties
	if rootCmd.Use != "dockerutilities" {
		t.Errorf("Expected root command use 'dockerutilities', got '%s'", rootCmd.Use)
	}

	if rootCmd.Short != "Docker analysis and management utilities" {
		t.Errorf("Expected root command short description, got '%s'", rootCmd.Short)
	}

	if !strings.Contains(rootCmd.Long, "dockerutilities is a comprehensive CLI tool") {
		t.Errorf("Expected root command long description to contain comprehensive description")
	}

	if rootCmd.Version != version.GetVersionString() {
		t.Errorf("Expected root command version to match version package, got '%s'", rootCmd.Version)
	}
}

// TestVersionCommandProperties tests the version command properties
func TestVersionCommandProperties(t *testing.T) {
	// Verify version command properties
	if versionCmd.Use != "version" {
		t.Errorf("Expected version command use 'version', got '%s'", versionCmd.Use)
	}

	if versionCmd.Short != "Print the version information of dockerutilities" {
		t.Errorf("Expected version command short description, got '%s'", versionCmd.Short)
	}

	if !strings.Contains(versionCmd.Long, "Print detailed version information") {
		t.Errorf("Expected version command long description to contain detailed version information")
	}
}

// TestVersionCommandExecution tests the version command execution
func TestVersionCommandExecution(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execute version command
	versionCmd.Run(versionCmd, []string{})

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

	// Verify output contains version information
	if !strings.Contains(output, version.GetFullVersionString()) {
		t.Errorf("Version command output should contain full version string, got: %s", output)
	}
}

// TestRootCommandHelp tests the root command help functionality
func TestRootCommandHelp(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Set up command with help flag
	rootCmd.SetArgs([]string{"--help"})

	// Execute command
	err := rootCmd.Execute()

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

	// Check for no error
	if err != nil {
		t.Errorf("Root command help should not return error, got: %v", err)
	}

	// Verify help output contains expected content
	expectedContent := []string{
		"dockerutilities",
		"Usage:",
		"Available Commands:",
		"Flags:",
		"Use \"dockerutilities [command] --help\"",
	}

	for _, content := range expectedContent {
		if !strings.Contains(output, content) {
			t.Errorf("Help output should contain '%s', got: %s", content, output)
		}
	}
}

// TestCommandStructure tests the overall command structure
func TestCommandStructure(t *testing.T) {
	// Verify root command has subcommands
	commands := rootCmd.Commands()
	if len(commands) == 0 {
		t.Error("Root command should have subcommands")
	}

	// Verify each subcommand has required properties
	for _, cmd := range commands {
		if cmd.Use == "" {
			t.Errorf("Command should have Use property set")
		}
		if cmd.Short == "" {
			t.Errorf("Command %s should have Short property set", cmd.Use)
		}
	}
}

// TestErrorHandling tests error handling scenarios
func TestErrorHandling(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "unknown command",
			args:    []string{"unknown"},
			wantErr: true,
		},
		{
			name:    "invalid flag",
			args:    []string{"--invalid-flag"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()

			if (err != nil) != tt.wantErr {
				t.Errorf("Command execution error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

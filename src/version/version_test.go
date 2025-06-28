package version

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "returns correct version",
			expected: "v0.0.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetVersion()
			if result != tt.expected {
				t.Errorf("GetVersion() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestVersionConstant(t *testing.T) {
	tests := []struct {
		name       string
		expected   string
		checkEmpty bool
	}{
		{
			name:       "version constant is set correctly",
			expected:   "v0.0.1",
			checkEmpty: false,
		},
		{
			name:       "version constant is not empty",
			expected:   "",
			checkEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.checkEmpty {
				if Version == "" {
					t.Errorf("Version constant should not be empty")
				}
			} else {
				if Version != tt.expected {
					t.Errorf("Version = %v, want %v", Version, tt.expected)
				}
			}
		})
	}
}

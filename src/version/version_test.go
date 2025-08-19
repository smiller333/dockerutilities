package version

import (
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestGetBuildInfo(t *testing.T) {
	buildInfo := GetBuildInfo()

	// Test that all required fields are present
	if buildInfo.Version == "" {
		t.Error("Version should not be empty")
	}

	if buildInfo.GitCommit == "" {
		t.Error("GitCommit should not be empty")
	}

	if buildInfo.BuildTime == "" {
		t.Error("BuildTime should not be empty")
	}

	if buildInfo.GoVersion == "" {
		t.Error("GoVersion should not be empty")
	}

	if buildInfo.GOOS == "" {
		t.Error("GOOS should not be empty")
	}

	if buildInfo.GOARCH == "" {
		t.Error("GOARCH should not be empty")
	}

	// Test that runtime values match expected values
	if buildInfo.GOOS != runtime.GOOS {
		t.Errorf("Expected GOOS %s, got %s", runtime.GOOS, buildInfo.GOOS)
	}

	if buildInfo.GOARCH != runtime.GOARCH {
		t.Errorf("Expected GOARCH %s, got %s", runtime.GOARCH, buildInfo.GOARCH)
	}

	if buildInfo.GoVersion != runtime.Version() {
		t.Errorf("Expected GoVersion %s, got %s", runtime.Version(), buildInfo.GoVersion)
	}
}

func TestGetVersion(t *testing.T) {
	version := GetVersion()
	if version == "" {
		t.Error("GetVersion() should not return empty string")
	}

	// Test backward compatibility
	buildInfo := GetBuildInfo()
	if version != buildInfo.Version {
		t.Errorf("GetVersion() should return same as BuildInfo.Version, got %s vs %s", version, buildInfo.Version)
	}
}

func TestGetVersionString(t *testing.T) {
	versionStr := GetVersionString()
	if versionStr == "" {
		t.Error("GetVersionString() should not return empty string")
	}

	// Should contain the version
	if !strings.Contains(versionStr, Version) {
		t.Errorf("Version string should contain version %s, got: %s", Version, versionStr)
	}

	// Test with a valid build time
	originalBuildTime := BuildTime
	defer func() { BuildTime = originalBuildTime }()

	BuildTime = time.Now().UTC().Format(time.RFC3339)
	versionStr = GetVersionString()

	if !strings.Contains(versionStr, "built") {
		t.Error("Version string should contain 'built' when BuildTime is set")
	}

	// Test with unknown build time
	BuildTime = "unknown"
	versionStr = GetVersionString()

	if !strings.Contains(versionStr, "build time unknown") {
		t.Error("Version string should contain 'build time unknown' when BuildTime is unknown")
	}
}

func TestGetFullVersionString(t *testing.T) {
	fullVersionStr := GetFullVersionString()
	if fullVersionStr == "" {
		t.Error("GetFullVersionString() should not return empty string")
	}

	// Should contain various components
	expectedComponents := []string{
		"dockerutilities",
		"Git Commit:",
		"Build Time:",
		"Go Version:",
		"OS/Arch:",
	}

	for _, component := range expectedComponents {
		if !strings.Contains(fullVersionStr, component) {
			t.Errorf("Full version string should contain '%s', got: %s", component, fullVersionStr)
		}
	}
}

func TestBuildTimeFormatting(t *testing.T) {
	// Test with valid RFC3339 time
	originalBuildTime := BuildTime
	defer func() { BuildTime = originalBuildTime }()

	testTime := "2023-12-25T10:30:45Z"
	BuildTime = testTime

	fullVersionStr := GetFullVersionString()

	// Should format the time nicely, not show raw RFC3339
	if strings.Contains(fullVersionStr, testTime) {
		t.Error("Full version string should format build time, not show raw RFC3339")
	}

	// Should contain formatted date components
	if !strings.Contains(fullVersionStr, "2023-12-25") {
		t.Error("Full version string should contain formatted date")
	}
}

// TestDefaultValues ensures the default values are sensible
func TestDefaultValues(t *testing.T) {
	// These are the default values when not built with ldflags
	if Version != "dev" && Version == "" {
		t.Error("Default Version should be 'dev' when not set by build")
	}

	if GitCommit != "unknown" && GitCommit == "" {
		t.Error("Default GitCommit should be 'unknown' when not set by build")
	}

	if BuildTime != "unknown" && BuildTime == "" {
		t.Error("Default BuildTime should be 'unknown' when not set by build")
	}

	// GoVersion should always be set by runtime
	if GoVersion == "" {
		t.Error("GoVersion should never be empty")
	}
}

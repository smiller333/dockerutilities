// Package version provides build information and version details for the dockerutils application.
package version

import (
	"fmt"
	"runtime"
	"time"
)

var (
	// These variables will be set at build time using ldflags
	Version   = "dev"             // Version string (e.g., v1.0.0 or git describe output)
	GitCommit = "unknown"         // Git commit hash (short)
	BuildTime = "unknown"         // Build timestamp in RFC3339 format
	GoVersion = runtime.Version() // Go version used to build
)

// BuildInfo contains version and build information
type BuildInfo struct {
	Version   string `json:"version"`    // Version string
	GitCommit string `json:"git_commit"` // Git commit hash
	BuildTime string `json:"build_time"` // Build timestamp
	GoVersion string `json:"go_version"` // Go version
	GOOS      string `json:"goos"`       // Target operating system
	GOARCH    string `json:"goarch"`     // Target architecture
}

// GetBuildInfo returns structured build information
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version:   Version,
		GitCommit: GitCommit,
		BuildTime: BuildTime,
		GoVersion: GoVersion,
		GOOS:      runtime.GOOS,
		GOARCH:    runtime.GOARCH,
	}
}

// GetVersion returns the current version string (backward compatibility)
func GetVersion() string {
	return Version
}

// GetVersionString returns a formatted version string
func GetVersionString() string {
	if BuildTime != "unknown" {
		if buildTime, err := time.Parse(time.RFC3339, BuildTime); err == nil {
			return fmt.Sprintf("%s (built %s)", Version, buildTime.Format("2006-01-02 15:04:05 MST"))
		}
	}
	return fmt.Sprintf("%s (build time unknown)", Version)
}

// GetFullVersionString returns a detailed version string with all build information
func GetFullVersionString() string {
	buildInfo := GetBuildInfo()
	var buildTimeStr string

	if buildInfo.BuildTime != "unknown" {
		if buildTime, err := time.Parse(time.RFC3339, buildInfo.BuildTime); err == nil {
			buildTimeStr = buildTime.Format("2006-01-02 15:04:05 MST")
		} else {
			buildTimeStr = buildInfo.BuildTime
		}
	} else {
		buildTimeStr = "unknown"
	}

	return fmt.Sprintf("dockerutils %s\nGit Commit: %s\nBuild Time: %s\nGo Version: %s\nOS/Arch: %s/%s",
		buildInfo.Version,
		buildInfo.GitCommit,
		buildTimeStr,
		buildInfo.GoVersion,
		buildInfo.GOOS,
		buildInfo.GOARCH,
	)
}

// Package version provides version information for the dockerutils CLI tool.
package version

// Version represents the current version of the dockerutils tool.
const Version = "v0.0.1"

// GetVersion returns the current version string.
func GetVersion() string {
	return Version
}

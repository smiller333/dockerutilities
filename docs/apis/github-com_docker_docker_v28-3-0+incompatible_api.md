# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:24:55 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/common.go#L4)

```go
const (
	// DefaultVersion of the current REST API.
	DefaultVersion = "1.51"

	// MinSupportedAPIVersion is the minimum API version that can be supported
	// by the API server, specified as "major.minor". Note that the daemon
	// may be configured with a different minimum API version, as returned
	// in [github.com/docker/docker/api/types.Version.MinAPIVersion].
	//
	// API requests for API versions lower than the configured version produce
	// an error.
	MinSupportedAPIVersion = "1.24"

	// NoBaseImageSpecifier is the symbol used by the FROM
	// command to specify that no base image is to be used.
	NoBaseImageSpecifier = "scratch"
)
```

## Variables

This section is empty.

## Functions

This section is empty.

## Types

This section is empty.


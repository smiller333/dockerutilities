# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/dockerversion

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:01 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/dockerversion/version_lib.go#L5)

```go
var (
	GitCommit             = "library-import"
	Version               = "library-import"
	BuildTime             = "library-import"
	PlatformName          = ""
	ProductName           = ""
	DefaultProductLicense = ""
)
```

## Functions

### DockerUserAgent

DockerUserAgent is the User-Agent the Docker client uses to identify itself.
In accordance with RFC 7231 (5.5.3) is of the form:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/dockerversion/useragent.go#L20)  
**Added in:** v1.11.0

```go
func DockerUserAgent(ctx context.Context, extraVersions ...useragent.VersionInfo) string
```

---

## Types

### UAStringKey

UAStringKey is used as key type for user-agent string in net/context struct

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/dockerversion/useragent.go#L14)  

```go
type UAStringKey struct{}
```

---


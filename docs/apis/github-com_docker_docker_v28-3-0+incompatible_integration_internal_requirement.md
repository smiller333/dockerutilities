# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration/internal/requirement

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:06:22 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CgroupNamespacesEnabled

CgroupNamespacesEnabled checks if cgroup namespaces are enabled on this host

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/requirement/requirement_linux.go#L12)  

```go
func CgroupNamespacesEnabled() bool
```

---

### HasHubConnectivity

HasHubConnectivity checks to see if https://hub.docker.com is
accessible from the present environment

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/requirement/requirement.go#L12)  

```go
func HasHubConnectivity(t *testing.T) bool
```

---

### Overlay2Supported

Overlay2Supported returns true if the current system supports overlay2 as graphdriver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/requirement/requirement_linux.go#L29)  

```go
func Overlay2Supported(kernelVersion string) bool
```

---

## Types

This section is empty.


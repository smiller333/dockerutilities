# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/cmd/dockerd/debug

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:23 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Disable

Disable sets the DEBUG env var to false
and makes the logger to log at info level.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/cmd/dockerd/debug/debug.go#L18)  

```go
func Disable()
```

---

### Enable

Enable sets the DEBUG env var to true
and makes the logger to log at debug level.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/cmd/dockerd/debug/debug.go#L11)  

```go
func Enable()
```

---

### IsEnabled

IsEnabled checks whether the debug flag is set or not.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/cmd/dockerd/debug/debug.go#L24)  

```go
func IsEnabled() bool
```

---

## Types

This section is empty.


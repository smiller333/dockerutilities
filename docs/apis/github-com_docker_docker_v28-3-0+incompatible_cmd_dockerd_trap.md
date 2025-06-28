# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/cmd/dockerd/trap

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:25 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Trap

Trap sets up a simplified signal "trap", appropriate for common
behavior expected from a vanilla unix command-line tool in general
(and the Docker engine in particular).

The first time a SIGINT or SIGTERM signal is received, `cleanup` is called in
a new goroutine.

If SIGINT or SIGTERM are received 3 times, the process is terminated
immediately with an exit code of 128 + the signal number.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/cmd/dockerd/trap/trap.go#L27)  

```go
func Trap(cleanup func())
```

---

## Types

This section is empty.


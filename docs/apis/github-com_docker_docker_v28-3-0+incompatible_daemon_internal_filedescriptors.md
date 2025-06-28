# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/internal/filedescriptors

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:28:43 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GetTotalUsedFds

GetTotalUsedFds Returns the number of used File Descriptors by
reading it via /proc filesystem.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/filedescriptors/filiedescriptors_linux.go#L16)  

```go
func GetTotalUsedFds(ctx context.Context) int
```

---

## Types

This section is empty.


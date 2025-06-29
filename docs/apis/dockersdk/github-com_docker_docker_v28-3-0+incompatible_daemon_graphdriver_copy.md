# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/graphdriver/copy

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:11 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### DirCopy

DirCopy copies or hardlinks the contents of one directory to another, properly
handling soft links, "security.capability" and (optionally) "trusted.overlay.opaque"
xattrs.

The copyOpaqueXattrs controls if "trusted.overlay.opaque" xattrs are copied.
Passing false disables copying "trusted.overlay.opaque" xattrs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/copy/copy.go#L124)  

```go
func DirCopy(srcDir, dstDir string, copyMode Mode, copyOpaqueXattrs bool) error
```

---

## Types

### Mode

Mode indicates whether to use hardlink or copy content

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/copy/copy.go#L22)  

```go
type Mode int
```

---


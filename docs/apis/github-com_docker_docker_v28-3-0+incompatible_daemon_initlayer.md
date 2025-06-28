# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/initlayer

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:35 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Setup

Setup populates a directory with mountpoints suitable
for bind-mounting things into the container.

This extra layer is used by all containers as the top-most ro layer. It protects
the container from unwanted side-effects on the rw layer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/initlayer/setup_unix.go#L19)  

```go
func Setup(initLayerFs string, uid int, gid int) error
```

---

## Types

This section is empty.


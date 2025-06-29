# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/mounttree

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:22 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### SwitchRoot

SwitchRoot changes path to be the root of the mount tree and changes the
current working directory to the new root.

This function bind-mounts onto path; it is the caller's responsibility to set
the desired propagation mode of path's parent mount beforehand to prevent
unwanted propagation into different mount namespaces.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/mounttree/switchroot_linux.go#L19)  

```go
func SwitchRoot(path string) error
```

---

## Types

This section is empty.


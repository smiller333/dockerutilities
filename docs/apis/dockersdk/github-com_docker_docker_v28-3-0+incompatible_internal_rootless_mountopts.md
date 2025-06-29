# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/rootless/mountopts

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:37 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### UnprivilegedMountFlags

UnprivilegedMountFlags gets the set of mount flags that are set on the mount that contains the given
path and are locked by CL_UNPRIVILEGED. This is necessary to ensure that
bind-mounting "with options" will not fail with user namespaces, due to
kernel restrictions that require user namespace mounts to preserve
CL_UNPRIVILEGED locked flags.

TODO: Move to github.com/moby/sys/mount, and update BuildKit copy of this code as well (https://github.com/moby/buildkit/blob/v0.13.0/util/rootless/mountopts/mountopts_linux.go#L11-L18)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/rootless/mountopts/mountopts_linux.go#L14)  

```go
func UnprivilegedMountFlags(path string) ([]string, error)
```

---

## Types

This section is empty.


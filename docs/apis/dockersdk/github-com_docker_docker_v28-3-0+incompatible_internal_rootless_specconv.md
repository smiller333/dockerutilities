# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/rootless/specconv

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:39 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ToRootfulInRootless

ToRootfulInRootless is used for "rootful-in-rootless" dind;
the daemon is running in UserNS but has no access to RootlessKit API socket, host filesystem, etc.

This function does:
* Fix up OOMScoreAdj (needed since systemd v250: https://github.com/moby/moby/issues/46563)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/rootless/specconv/specconv_linux.go#L21)  

```go
func ToRootfulInRootless(spec *specs.Spec)
```

---

### ToRootless

ToRootless converts spec to be compatible with "rootless" runc.
* Remove non-supported cgroups
* Fix up OOMScoreAdj
* Fix up /proc if --pid=host
* Fix up /dev/shm and /dev/mqueue if --ipc=host

v2Controllers should be non-nil only if running with v2 and systemd.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/rootless/specconv/specconv_linux.go#L37)  

```go
func ToRootless(spec *specs.Spec, v2Controllers []string) error
```

---

## Types

This section is empty.


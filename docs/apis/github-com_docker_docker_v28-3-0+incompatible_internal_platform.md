# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/platform

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:34 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Architecture

Architecture returns the runtime architecture of the process.

Unlike runtime.GOARCH (which refers to the compiler platform),
Architecture refers to the running platform.

For example, Architecture reports "x86_64" as architecture, even
when running a "linux/386" compiled binary on "linux/amd64" hardware.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/platform/platform.go#L23)  

```go
func Architecture() string
```

---

### NumProcs ⚠️ **DEPRECATED**

NumProcs returns the number of processors on the system

Deprecated: temporary stub for non-Windows to provide an alias for the deprecated github.com/docker/docker/pkg/platform package.

FIXME(thaJeztah): remove once we remove  github.com/docker/docker/pkg/platform

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/platform/platform_unix.go#L23)  

```go
func NumProcs() uint32
```

---

### PossibleCPU

PossibleCPU returns the set of possible CPUs on the host (which is equal or
larger to the number of CPUs currently online). The returned set may be a
single CPU number ({0}), or a continuous range of CPU numbers ({0,1,2,3}), or
a non-continuous range of CPU numbers ({0,1,2,3,12,13,14,15}).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/platform/platform.go#L38)  

```go
func PossibleCPU() []int
```

---

## Types

This section is empty.


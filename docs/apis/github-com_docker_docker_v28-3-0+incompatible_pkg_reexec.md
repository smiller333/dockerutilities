# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/reexec

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:23 UTC

## Overview

Package reexec facilitates the busybox style reexec of a binary.

Deprecated: this package is deprecated and moved to a separate module. Use github.com/moby/sys/reexec instead.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Command

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/reexec/reexec_deprecated.go#L32)  

```go
func Command(args ...string) *exec.Cmd
```

---

### Init

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/reexec/reexec_deprecated.go#L24)  

```go
func Init() bool
```

---

### Register

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/reexec/reexec_deprecated.go#L16)  

```go
func Register(name string, initializer func())
```

---

### Self

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/reexec/reexec_deprecated.go#L39)  

```go
func Self() string
```

---

## Types

This section is empty.


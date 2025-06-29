# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/reexec

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:02 UTC

## Overview

Package reexec facilitates the busybox style reexec of a binary.

Deprecated: this package is deprecated and moved to a separate module. Use github.com/moby/sys/reexec instead.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Command ⚠️ **DEPRECATED**

Command returns an *exec.Cmd with its Path set to the path of the current
binary using the result of Self.

Deprecated: use reexec.Command instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/reexec/reexec_deprecated.go#L32)  

```go
func Command(args ...string) *exec.Cmd
```

---

### Init ⚠️ **DEPRECATED**

Init is called as the first part of the exec process and returns true if an
initialization function was called.

Deprecated: use reexec.Init instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/reexec/reexec_deprecated.go#L24)  

```go
func Init() bool
```

---

### Register ⚠️ **DEPRECATED**

Register adds an initialization func under the specified name. It panics
if the given name is already registered.

Deprecated: use reexec.Register instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/reexec/reexec_deprecated.go#L16)  

```go
func Register(name string, initializer func())
```

---

### Self ⚠️ **DEPRECATED**

Self returns the path to the current process's binary.

Deprecated: use reexec.Self instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/reexec/reexec_deprecated.go#L39)  

```go
func Self() string
```

---

## Types

This section is empty.


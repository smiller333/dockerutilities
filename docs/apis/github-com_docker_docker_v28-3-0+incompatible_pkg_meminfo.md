# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/meminfo

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:43 UTC

## Overview

Package meminfo provides utilities to retrieve memory statistics of
the host system.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Memory

Memory contains memory statistics of the host system.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/meminfo/meminfo.go#L13)  

```go
type Memory struct {
	// Total usable RAM (i.e. physical RAM minus a few reserved bits and the
	// kernel binary code).
	MemTotal int64

	// Amount of free memory.
	MemFree int64

	// Total amount of swap space available.
	SwapTotal int64

	// Amount of swap space that is currently unused.
	SwapFree int64
}
```

#### Functions

##### Read

Read retrieves memory statistics of the host system and returns a
Memory type. It is only supported on Linux and Windows, and returns an
error on other platforms.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/meminfo/meminfo.go#L8)  

```go
func Read() (*Memory, error)
```

---


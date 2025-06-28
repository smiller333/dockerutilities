# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/parsers/kernel

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:51 UTC

## Overview

Package kernel provides helper function to get, parse and compare kernel
versions for different platforms.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CheckKernelVersion

CheckKernelVersion checks if current kernel is newer than (or equal to)
the given version.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/kernel/kernel_unix.go#L25)  
**Added in:** v1.13.0

```go
func CheckKernelVersion(k, major, minor int) bool
```

---

### CompareKernelVersion

CompareKernelVersion compares two kernel.VersionInfo structs.
Returns -1 if a < b, 0 if a == b, 1 it a > b

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/kernel/kernel.go#L26)  

```go
func CompareKernelVersion(a, b VersionInfo) int
```

---

## Types

### VersionInfo

VersionInfo holds information about the kernel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/kernel/kernel.go#L13)  
**Added in:** v1.9.0

```go
type VersionInfo struct {
	Kernel int    // Version of the kernel (e.g. 4.1.2-generic -> 4)
	Major  int    // Major part of the kernel version (e.g. 4.1.2-generic -> 1)
	Minor  int    // Minor part of the kernel version (e.g. 4.1.2-generic -> 2)
	Flavor string // Flavor of the kernel version (e.g. 4.1.2-generic -> generic)
}
```

#### Functions

##### GetKernelVersion

GetKernelVersion gets the current kernel version.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/kernel/kernel_unix.go#L13)  

```go
func GetKernelVersion() (*VersionInfo, error)
```

##### ParseRelease

ParseRelease parses a string and creates a VersionInfo based on it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/kernel/kernel.go#L49)  

```go
func ParseRelease(release string) (*VersionInfo, error)
```

#### Methods

##### VersionInfo.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/kernel/kernel.go#L20)  
**Added in:** v1.9.0

```go
func (k *VersionInfo) String() string
```

---


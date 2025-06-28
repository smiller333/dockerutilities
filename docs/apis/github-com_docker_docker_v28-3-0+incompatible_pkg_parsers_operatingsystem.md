# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/parsers/operatingsystem

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:53 UTC

## Overview

Package operatingsystem provides helper function to get the operating system
name for different platforms.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GetOperatingSystem

GetOperatingSystem gets the name of the current operating system.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/operatingsystem/operatingsystem_linux.go#L24)  

```go
func GetOperatingSystem() (string, error)
```

---

### GetOperatingSystemVersion

GetOperatingSystemVersion gets the version of the current operating system, as a string.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/operatingsystem/operatingsystem_linux.go#L37)  

```go
func GetOperatingSystemVersion() (string, error)
```

---

### IsContainerized

IsContainerized returns true if we are running inside a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/operatingsystem/operatingsystem_linux.go#L69)  

```go
func IsContainerized() (bool, error)
```

---

## Types

This section is empty.


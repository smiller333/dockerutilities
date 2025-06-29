# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/longpath

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:34 UTC

## Overview

Package longpath introduces some constants and helper functions for handling
long paths in Windows.

Long paths are expected to be prepended with "\\?\" and followed by either a
drive letter, a UNC server\share, or a volume identifier.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AddPrefix

AddPrefix adds the Windows long path prefix to the path provided if
it does not already have it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/longpath/longpath.go#L19)  

```go
func AddPrefix(path string) string
```

---

### MkdirTemp

MkdirTemp is the equivalent of os.MkdirTemp, except that on Windows
the result is in Windows longpath format. On Unix systems it is
equivalent to os.MkdirTemp.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/longpath/longpath.go#L33)  

```go
func MkdirTemp(dir, prefix string) (string, error)
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/fileutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:23 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CopyFile

CopyFile copies from src to dst until either EOF is reached
on src or an error occurs. It verifies src exists and removes
the dst if it exists.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/fileutils/fileutils.go#L13)  
**Added in:** v1.7.0

```go
func CopyFile(src, dst string) (int64, error)
```

---

### CreateIfNotExists

CreateIfNotExists creates a file or a directory only if it does not already exist.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/fileutils/fileutils.go#L58)  
**Added in:** v1.7.1

```go
func CreateIfNotExists(path string, isDir bool) error
```

---

### ReadSymlinkedDirectory

ReadSymlinkedDirectory returns the target directory of a symlink.
The target of the symbolic link may not be a file.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/fileutils/fileutils.go#L37)  
**Added in:** v1.7.0

```go
func ReadSymlinkedDirectory(path string) (realPath string, _ error)
```

---

## Types

This section is empty.


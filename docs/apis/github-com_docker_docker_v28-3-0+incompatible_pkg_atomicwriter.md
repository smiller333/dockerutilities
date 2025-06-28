# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/atomicwriter

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:21 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/atomicwriter/atomicwriter_deprecated.go#L23)  

```go
func New(filename string, perm os.FileMode) (io.WriteCloser, error)
```

---

### NewWriteSet

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/atomicwriter/atomicwriter_deprecated.go#L54)  

```go
func NewWriteSet(tmpDir string) (*atomicwriter.WriteSet, error)
```

---

### WriteFile

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/atomicwriter/atomicwriter_deprecated.go#L36)  

```go
func WriteFile(filename string, data []byte, perm os.FileMode) error
```

---

## Types

### WriteSet

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/atomicwriter/atomicwriter_deprecated.go#L45)  

```go
type WriteSet = atomicwriter.WriteSet
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/chrootarchive

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:26 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ApplyLayer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/diff_deprecated.go#L14)  

```go
func ApplyLayer(dest string, layer io.Reader) (size int64, err error)
```

---

### ApplyUncompressedLayer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/diff_deprecated.go#L22)  
**Added in:** v1.8.0

```go
func ApplyUncompressedLayer(dest string, layer io.Reader, options *archive.TarOptions) (int64, error)
```

---

### NewArchiver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L14)  

```go
func NewArchiver(idMapping idtools.IdentityMapping) *archive.Archiver
```

---

### Tar

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L47)  

```go
func Tar(srcPath string, options *archive.TarOptions, root string) (io.ReadCloser, error)
```

---

### Untar

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L25)  

```go
func Untar(tarArchive io.Reader, dest string, options *archive.TarOptions) error
```

---

### UntarUncompressed

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L40)  
**Added in:** v1.8.0

```go
func UntarUncompressed(tarArchive io.Reader, dest string, options *archive.TarOptions) error
```

---

### UntarWithRoot

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L32)  

```go
func UntarWithRoot(tarArchive io.Reader, dest string, options *archive.TarOptions, root string) error
```

---

## Types

This section is empty.


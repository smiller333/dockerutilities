# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/chrootarchive

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:20 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ApplyLayer ⚠️ **DEPRECATED**

ApplyLayer parses a diff in the standard layer format from `layer`,
and applies it to the directory `dest`.

Deprecated: use chrootarchive.ApplyLayer insteead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/diff_deprecated.go#L14)  

```go
func ApplyLayer(dest string, layer io.Reader) (size int64, err error)
```

---

### ApplyUncompressedLayer ⚠️ **DEPRECATED**

ApplyUncompressedLayer parses a diff in the standard layer format from
`layer`, and applies it to the directory `dest`.

Deprecated: use chrootarchive.ApplyUncompressedLayer insteead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/diff_deprecated.go#L22)  
**Added in:** v1.8.0

```go
func ApplyUncompressedLayer(dest string, layer io.Reader, options *archive.TarOptions) (int64, error)
```

---

### NewArchiver ⚠️ **DEPRECATED**

NewArchiver returns a new Archiver which uses chrootarchive.Untar

Deprecated: use chrootarchive.NewArchiver instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L14)  

```go
func NewArchiver(idMapping idtools.IdentityMapping) *archive.Archiver
```

---

### Tar ⚠️ **DEPRECATED**

Tar tars the requested path while chrooted to the specified root.

Deprecated: use chrootarchive.Tar instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L47)  

```go
func Tar(srcPath string, options *archive.TarOptions, root string) (io.ReadCloser, error)
```

---

### Untar ⚠️ **DEPRECATED**

Untar reads a stream of bytes from `archive`, parses it as a tar archive,
and unpacks it into the directory at `dest`.

Deprecated: use chrootarchive.Untar instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L25)  

```go
func Untar(tarArchive io.Reader, dest string, options *archive.TarOptions) error
```

---

### UntarUncompressed ⚠️ **DEPRECATED**

UntarUncompressed reads a stream of bytes from tarArchive, parses it as a tar archive,
and unpacks it into the directory at dest.

Deprecated: use chrootarchive.UntarUncompressed instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L40)  
**Added in:** v1.8.0

```go
func UntarUncompressed(tarArchive io.Reader, dest string, options *archive.TarOptions) error
```

---

### UntarWithRoot ⚠️ **DEPRECATED**

UntarWithRoot is the same as Untar, but allows you to pass in a root directory.

Deprecated: use chrootarchive.UntarWithRoot instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/chrootarchive/archive_deprecated.go#L32)  

```go
func UntarWithRoot(tarArchive io.Reader, dest string, options *archive.TarOptions, root string) error
```

---

## Types

This section is empty.


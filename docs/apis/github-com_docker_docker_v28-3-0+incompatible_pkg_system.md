# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/system

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:18 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/errors.go#L6)

```go
var ErrNotSupportedPlatform = errors.New("platform and architecture is not supported")
```

## Functions

### Chtimes

Chtimes changes the access time and modified time of a file at the given path.
If the modified time is prior to the Unix Epoch (unixMinTime), or after the
end of Unix Time (unixEpochTime), os.Chtimes has undefined behavior. In this
case, Chtimes defaults to Unix Epoch, just in case.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/chtimes.go#L33)  
**Added in:** v1.9.0

```go
func Chtimes(name string, atime time.Time, mtime time.Time) error
```

---

### IsAbs

IsAbs is a platform-agnostic wrapper for filepath.IsAbs.

On Windows, golang filepath.IsAbs does not consider a path \windows\system32
as absolute as it doesn't start with a drive-letter/colon combination. However,
in docker we need to verify things such as WORKDIR /windows/system32 in
a Dockerfile (which gets translated to \windows\system32 when being processed
by the daemon). This SHOULD be treated as absolute from a docker processing
perspective.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/filesys.go#L17)  
**Added in:** v1.9.0

```go
func IsAbs(path string) bool
```

---

### LUtimesNano

LUtimesNano is used to change access and modification time of the specified path.
It's used for symbol link file because unix.UtimesNano doesn't support a NOFOLLOW flag atm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/utimes_unix.go#L14)  

```go
func LUtimesNano(path string, ts []syscall.Timespec) error
```

---

### Lgetxattr

Lgetxattr retrieves the value of the extended attribute identified by attr
and associated with the given path in the file system.
It returns a nil slice and nil error if the xattr is not set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/xattrs_linux.go#L12)  

```go
func Lgetxattr(path string, attr string) ([]byte, error)
```

---

### Lsetxattr

Lsetxattr sets the value of the extended attribute identified by attr
and associated with the given path in the file system.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/xattrs_linux.go#L43)  

```go
func Lsetxattr(path string, attr string, data []byte, flags int) error
```

---

### MkdirAll ⚠️ **DEPRECATED**

MkdirAll creates a directory named path along with any necessary parents,
with permission specified by attribute perm for all dir created.

Deprecated: os.MkdirAll now natively supports Windows GUID volume paths, and should be used instead. This alias will be removed in the next release.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/filesys.go#L25)  
**Added in:** v1.7.0

```go
func MkdirAll(path string, perm os.FileMode) error
```

---

### MkdirAllWithACL

MkdirAllWithACL is a wrapper for os.MkdirAll on unix systems.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/filesys_unix.go#L8)  
**Added in:** v1.13.0

```go
func MkdirAllWithACL(path string, perm os.FileMode, _ string) error
```

---

## Types

### XattrError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/xattrs.go#L3)  

```go
type XattrError struct {
	Op   string
	Attr string
	Path string
	Err  error
}
```

#### Methods

##### XattrError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/xattrs.go#L10)  

```go
func (e *XattrError) Error() string
```

##### XattrError.Timeout

Timeout reports whether this error represents a timeout.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/xattrs.go#L15)  

```go
func (e *XattrError) Timeout() bool
```

##### XattrError.Unwrap

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/system/xattrs.go#L12)  

```go
func (e *XattrError) Unwrap() error
```

---


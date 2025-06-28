# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/atomicwriter

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:16 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New ⚠️ **DEPRECATED**

New returns a WriteCloser so that writing to it writes to a
temporary file and closing it atomically changes the temporary file to
destination path. Writing and closing concurrently is not allowed.
NOTE: umask is not considered for the file's permissions.

New uses [sequential.CreateTemp] to use sequential file access on Windows,
avoiding depleting the standby list un-necessarily. On Linux, this equates to
a regular os.CreateTemp. Refer to the Win32 API documentation for details
on sequential file access.

Deprecated: use atomicwriter.New instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/atomicwriter/atomicwriter_deprecated.go#L23)  

```go
func New(filename string, perm os.FileMode) (io.WriteCloser, error)
```

---

### NewWriteSet ⚠️ **DEPRECATED**

NewWriteSet creates a new atomic write set to
atomically create a set of files. The given directory
is used as the base directory for storing files before
commit. If no temporary directory is given the system
default is used.

Deprecated: use atomicwriter.NewWriteSet instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/atomicwriter/atomicwriter_deprecated.go#L54)  

```go
func NewWriteSet(tmpDir string) (*atomicwriter.WriteSet, error)
```

---

### WriteFile ⚠️ **DEPRECATED**

WriteFile atomically writes data to a file named by filename and with the
specified permission bits. The given filename is created if it does not exist,
but the destination directory must exist. It can be used as a drop-in replacement
for os.WriteFile, but currently does not allow the destination path to be
a symlink. WriteFile is implemented using New for its implementation.

NOTE: umask is not considered for the file's permissions.

Deprecated: use atomicwriter.WriteFile instead.

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


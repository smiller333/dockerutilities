# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/ioutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:30 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AtomicWriteFile ⚠️ **DEPRECATED**

AtomicWriteFile atomically writes data to a file named by filename and with the specified permission bits.
NOTE: umask is not considered for the file's permissions.

Deprecated: use atomicwriter.WriteFile instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/fswriters_deprecated.go#L24)  
**Added in:** v1.12.0

```go
func AtomicWriteFile(filename string, data []byte, perm os.FileMode) error
```

---

### NewAtomicFileWriter ⚠️ **DEPRECATED**

NewAtomicFileWriter returns WriteCloser so that writing to it writes to a
temporary file and closing it atomically changes the temporary file to
destination path. Writing and closing concurrently is not allowed.
NOTE: umask is not considered for the file's permissions.

Deprecated: use atomicwriter.New instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/fswriters_deprecated.go#L16)  
**Added in:** v1.12.0

```go
func NewAtomicFileWriter(filename string, perm os.FileMode) (io.WriteCloser, error)
```

---

### NewAtomicWriteSet ⚠️ **DEPRECATED**

NewAtomicWriteSet creates a new atomic write set to
atomically create a set of files. The given directory
is used as the base directory for storing files before
commit. If no temporary directory is given the system
default is used.

Deprecated: use atomicwriter.NewWriteSet instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/fswriters_deprecated.go#L42)  
**Added in:** v1.13.0

```go
func NewAtomicWriteSet(tmpDir string) (*atomicwriter.WriteSet, error)
```

---

### NewCancelReadCloser

NewCancelReadCloser creates a wrapper that closes the ReadCloser when the
context is cancelled. The returned io.ReadCloser must be closed when it is
no longer needed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/readers.go#L51)  
**Added in:** v1.10.0

```go
func NewCancelReadCloser(ctx context.Context, in io.ReadCloser) io.ReadCloser
```

---

### NewReadCloserWrapper

NewReadCloserWrapper wraps an io.Reader, and implements an io.ReadCloser.
It calls the given callback function when closed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/readers.go#L32)  

```go
func NewReadCloserWrapper(r io.Reader, closer func() error) io.ReadCloser
```

---

### NewWriteCloserWrapper

NewWriteCloserWrapper returns a new io.WriteCloser.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/writers.go#L23)  

```go
func NewWriteCloserWrapper(r io.Writer, closer func() error) io.WriteCloser
```

---

## Types

### AtomicWriteSet

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/fswriters_deprecated.go#L33)  
**Added in:** v1.13.0

```go
type AtomicWriteSet = atomicwriter.WriteSet
```

---

### WriteFlusher

WriteFlusher wraps the Write and Flush operation ensuring that every write
is a flush. In addition, the Close method can be called to intercept
Read/Write calls if the targets lifecycle has already ended.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/writeflusher.go#L11)  
**Added in:** v1.7.0

```go
type WriteFlusher struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewWriteFlusher

NewWriteFlusher returns a new WriteFlusher.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/writeflusher.go#L88)  
**Added in:** v1.7.0

```go
func NewWriteFlusher(w io.Writer) *WriteFlusher
```

#### Methods

##### WriteFlusher.Close

Close closes the write flusher, disallowing any further writes to the
target. After the flusher is closed, all calls to write or flush will
result in an error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/writeflusher.go#L68)  
**Added in:** v1.9.1

```go
func (wf *WriteFlusher) Close() error
```

##### WriteFlusher.Flush

Flush the stream immediately.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/writeflusher.go#L37)  
**Added in:** v1.7.0

```go
func (wf *WriteFlusher) Flush()
```

##### WriteFlusher.Flushed

Flushed returns the state of flushed.
If it's flushed, return true, or else it return false.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/writeflusher.go#L52)  
**Added in:** v1.7.0

```go
func (wf *WriteFlusher) Flushed() bool
```

##### WriteFlusher.Write

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/ioutils/writeflusher.go#L24)  
**Added in:** v1.7.0

```go
func (wf *WriteFlusher) Write(b []byte) (int, error)
```

---

## Notes

**Section structure:**
- Documentation-note

## Bugs


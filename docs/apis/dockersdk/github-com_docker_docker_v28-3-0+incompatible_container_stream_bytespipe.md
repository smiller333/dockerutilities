# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/container/stream/bytespipe

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:33 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/bytespipe/bytespipe.go#L19)

```go
var (
	// ErrClosed is returned when Write is called on a closed BytesPipe.
	ErrClosed = errors.New("write to closed BytesPipe")
)
```

## Functions

This section is empty.

## Types

### BytesPipe

BytesPipe is io.ReadWriteCloser which works similarly to pipe(queue).
All written data may be read at most once. Also, BytesPipe allocates
and releases new byte slices to adjust to current needs, so the buffer
won't be overgrown after peak loads.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/bytespipe/bytespipe.go#L31)  

```go
type BytesPipe struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates new BytesPipe, initialized by specified slice.
If buf is nil, then it will be initialized with slice which cap is 64.
buf will be adjusted in a way that len(buf) == 0, cap(buf) == cap(buf).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/bytespipe/bytespipe.go#L43)  

```go
func New() *BytesPipe
```

#### Methods

##### BytesPipe.Close

Close causes further reads from a BytesPipe to return immediately.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/bytespipe/bytespipe.go#L123)  

```go
func (bp *BytesPipe) Close() error
```

##### BytesPipe.CloseWithError

CloseWithError causes further reads from a BytesPipe to return immediately.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/bytespipe/bytespipe.go#L110)  

```go
func (bp *BytesPipe) CloseWithError(err error) error
```

##### BytesPipe.Read

Read reads bytes from BytesPipe.
Data could be read only once.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/bytespipe/bytespipe.go#L129)  

```go
func (bp *BytesPipe) Read(p []byte) (int, error)
```

##### BytesPipe.Write

Write writes p to BytesPipe.
It can allocate new []byte slices in a process of writing.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/bytespipe/bytespipe.go#L52)  

```go
func (bp *BytesPipe) Write(p []byte) (int, error)
```

---


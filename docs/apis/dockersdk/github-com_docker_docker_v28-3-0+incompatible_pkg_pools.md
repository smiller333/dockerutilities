# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/pools

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:14:55 UTC

## Overview

Package pools provides a collection of pools which provide various
data types with buffers. These can be used to lower the number of
memory allocations and reuse buffers.

New pools should be added to this package to allow them to be
shared across packages.

Utility functions which operate on pools should be added to this
package to allow them to be reused.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L22)

```go
var (
	// BufioReader32KPool is a pool which returns bufio.Reader with a 32K buffer.
	BufioReader32KPool = newBufioReaderPoolWithSize(buffer32K)
	// BufioWriter32KPool is a pool which returns bufio.Writer with a 32K buffer.
	BufioWriter32KPool = newBufioWriterPoolWithSize(buffer32K)
)
```

## Functions

### Copy

Copy is a convenience wrapper which uses a buffer to avoid allocation in io.Copy.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L79)  
**Added in:** v1.8.0

```go
func Copy(dst io.Writer, src io.Reader) (written int64, _ error)
```

---

## Types

### BufioReaderPool

BufioReaderPool is a bufio reader that uses sync.Pool.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L31)  

```go
type BufioReaderPool struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### BufioReaderPool.Get

Get returns a bufio.Reader which reads from r. The buffer size is that of the pool.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L46)  

```go
func (bufPool *BufioReaderPool) Get(r io.Reader) *bufio.Reader
```

##### BufioReaderPool.NewReadCloserWrapper

NewReadCloserWrapper returns a wrapper which puts the bufio.Reader back
into the pool and closes the reader if it's an io.ReadCloser.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L88)  

```go
func (bufPool *BufioReaderPool) NewReadCloserWrapper(buf *bufio.Reader, r io.Reader) io.ReadCloser
```

##### BufioReaderPool.Put

Put puts the bufio.Reader back into the pool.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L53)  

```go
func (bufPool *BufioReaderPool) Put(b *bufio.Reader)
```

---

### BufioWriterPool

BufioWriterPool is a bufio writer that uses sync.Pool.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L99)  

```go
type BufioWriterPool struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### BufioWriterPool.Get

Get returns a bufio.Writer which writes to w. The buffer size is that of the pool.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L114)  

```go
func (bufPool *BufioWriterPool) Get(w io.Writer) *bufio.Writer
```

##### BufioWriterPool.NewWriteCloserWrapper

NewWriteCloserWrapper returns a wrapper which puts the bufio.Writer back
into the pool and closes the writer if it's an io.WriteCloser.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L128)  

```go
func (bufPool *BufioWriterPool) NewWriteCloserWrapper(buf *bufio.Writer, w io.Writer) io.WriteCloser
```

##### BufioWriterPool.Put

Put puts the bufio.Writer back into the pool.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pools/pools.go#L121)  

```go
func (bufPool *BufioWriterPool) Put(b *bufio.Writer)
```

---


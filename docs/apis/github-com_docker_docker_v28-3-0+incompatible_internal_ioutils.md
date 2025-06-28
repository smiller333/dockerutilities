# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/ioutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:07 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CopyCtx

CopyCtx copies from src to dst until either EOF is reached on src or a context is cancelled.
The writer is not closed when the context is cancelled.

After CopyCtx exits due to context cancellation, the goroutine that performed
the copy may still be running if either the reader or writer blocks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/ioutils/copy.go#L13)  

```go
func CopyCtx(ctx context.Context, dst io.Writer, src io.Reader) (n int64, err error)
```

---

### NewCtxReader

NewCtxReader wraps the given reader with a reader that doesn't proceed with
reading if the context is done.

Note: Read will still block if the underlying reader blocks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/ioutils/copy.go#L41)  

```go
func NewCtxReader(ctx context.Context, r io.Reader) io.Reader
```

---

## Types

This section is empty.


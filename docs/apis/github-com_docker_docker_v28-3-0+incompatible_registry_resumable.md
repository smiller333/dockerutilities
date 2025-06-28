# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/registry/resumable

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:47 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRequestReader

NewRequestReader makes it possible to resume reading a request's body transparently
maxfail is the number of times we retry to make requests again (not resumes)
totalsize is the total length of the body; auto detect if not provided

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/resumable/resumablerequestreader.go#L28)  

```go
func NewRequestReader(c *http.Client, r *http.Request, maxfail uint32, totalsize int64) io.ReadCloser
```

---

### NewRequestReaderWithInitialResponse

NewRequestReaderWithInitialResponse makes it possible to resume
reading the body of an already initiated request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/resumable/resumablerequestreader.go#L34)  

```go
func NewRequestReaderWithInitialResponse(c *http.Client, r *http.Request, maxfail uint32, totalsize int64, initialResponse *http.Response) io.ReadCloser
```

---

## Types

This section is empty.


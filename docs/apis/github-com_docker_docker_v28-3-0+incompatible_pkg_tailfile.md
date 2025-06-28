# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/tailfile

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:20 UTC

## Overview

Package tailfile provides helper functions to read the nth lines of any
ReadSeeker.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tailfile/tailfile.go#L19)

```go
var ErrNonPositiveLinesNumber = errors.New("The number of lines to extract from the file must be positive")
```

## Functions

### NewTailReader

NewTailReader scopes the passed in reader to just the last N lines passed in

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tailfile/tailfile.go#L51)  

```go
func NewTailReader(ctx context.Context, r SizeReaderAt, reqLines int) (*io.SectionReader, int, error)
```

---

### NewTailReaderWithDelimiter

NewTailReaderWithDelimiter scopes the passed in reader to just the last N lines passed in
In this case a "line" is defined by the passed in delimiter.

Delimiter lengths should be generally small, no more than 12 bytes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tailfile/tailfile.go#L59)  

```go
func NewTailReaderWithDelimiter(ctx context.Context, r SizeReaderAt, reqLines int, delimiter []byte) (*io.SectionReader, int, error)
```

---

### TailFile

TailFile returns last n lines of the passed in file.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tailfile/tailfile.go#L22)  

```go
func TailFile(f *os.File, n int) ([][]byte, error)
```

---

## Types

### SizeReaderAt

SizeReaderAt is an interface used to get a ReaderAt as well as the size of the underlying reader.
Note that the size of the underlying reader should not change when using this interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tailfile/tailfile.go#L45)  

```go
type SizeReaderAt interface {
	io.ReaderAt
	Size() int64
}
```

---


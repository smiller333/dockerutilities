# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/stdcopy

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:30 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewStdWriter

NewStdWriter instantiates a new Writer.
Everything written to it will be encapsulated using a custom format,
and written to the underlying `w` stream.
This allows multiple write streams (e.g. stdout and stderr) to be muxed into a single connection.
`t` indicates the id of the stream to encapsulate.
It can be stdcopy.Stdin, stdcopy.Stdout, stdcopy.Stderr.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/stdcopy/stdcopy.go#L77)  

```go
func NewStdWriter(w io.Writer, t StdType) io.Writer
```

---

### StdCopy

StdCopy is a modified version of io.Copy.

StdCopy will demultiplex `src`, assuming that it contains two streams,
previously multiplexed together using a StdWriter instance.
As it reads from `src`, StdCopy will write to `dstout` and `dsterr`.

StdCopy will read until it hits EOF on `src`. It will then return a nil error.
In other words: if `err` is non nil, it indicates a real underlying error.

`written` will hold the total number of bytes written to `dstout` and `dsterr`.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/stdcopy/stdcopy.go#L94)  

```go
func StdCopy(dstout, dsterr io.Writer, src io.Reader) (written int64, _ error)
```

---

## Types

### StdType

StdType is the type of standard stream
a writer can multiplex to.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/stdcopy/stdcopy.go#L14)  

```go
type StdType byte
```

---


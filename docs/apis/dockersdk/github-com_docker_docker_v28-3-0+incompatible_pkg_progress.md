# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/progress

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:14:59 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Aux

Aux sends auxiliary information over a progress interface, which will not be
formatted for the UI. This is used for things such as push signing.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progress.go#L91)  

```go
func Aux(out Output, a interface{})
```

---

### Message

Message is a convenience function to write a progress message to the channel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progress.go#L79)  

```go
func Message(out Output, id, message string)
```

---

### Messagef

Messagef is a convenience function to write a printf-formatted progress
message to the channel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progress.go#L85)  

```go
func Messagef(out Output, id, format string, a ...interface{})
```

---

### Update

Update is a convenience function to write a progress update to the channel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progress.go#L68)  

```go
func Update(out Output, id, action string)
```

---

### Updatef

Updatef is a convenience function to write a printf-formatted progress update
to the channel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progress.go#L74)  

```go
func Updatef(out Output, id, format string, a ...interface{})
```

---

## Types

### Output

Output is an interface for writing progress information. It's
like a writer for progress, but we don't call it Writer because
that would be confusing next to ProgressReader (also, because it
doesn't implement the io.Writer interface).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progress.go#L35)  

```go
type Output interface {
	WriteProgress(Progress) error
}
```

#### Functions

##### ChanOutput

ChanOutput returns an Output that writes progress updates to the
supplied channel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progress.go#L52)  

```go
func ChanOutput(progressChan chan<- Progress) Output
```

##### DiscardOutput

DiscardOutput returns an Output that discards progress

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progress.go#L63)  
**Added in:** v1.13.0

```go
func DiscardOutput() Output
```

---

### Progress

Progress represents the progress of a transfer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progress.go#L8)  

```go
type Progress struct {
	ID string

	// Progress contains a Message or...
	Message string

	// ...progress of an action
	Action  string
	Current int64
	Total   int64

	// If true, don't show xB/yB
	HideCounts bool
	// If not empty, use units instead of bytes for counts
	Units string

	// Aux contains extra information not presented to the user, such as
	// digests for push signing.
	Aux interface{}

	LastUpdate bool
}
```

---

### Reader

Reader is a Reader with progress bar.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progressreader.go#L11)  

```go
type Reader struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewProgressReader

NewProgressReader creates a new ProgressReader.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progressreader.go#L23)  

```go
func NewProgressReader(in io.ReadCloser, out Output, size int64, id, action string) *Reader
```

#### Methods

##### Reader.Close

Close closes the progress reader and its underlying reader.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progressreader.go#L53)  

```go
func (p *Reader) Close() error
```

##### Reader.Read

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/progress/progressreader.go#L34)  

```go
func (p *Reader) Read(buf []byte) (int, error)
```

---


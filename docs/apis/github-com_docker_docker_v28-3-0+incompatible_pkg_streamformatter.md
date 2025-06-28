# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/streamformatter

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:33 UTC

## Overview

Package streamformatter provides helper functions to format a stream.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### FormatError

FormatError formats the error as a JSON object

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/streamformatter/streamformatter.go#L33)  

```go
func FormatError(err error) []byte
```

---

### FormatStatus

FormatStatus formats the specified objects according to the specified format (and id).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/streamformatter/streamformatter.go#L23)  

```go
func FormatStatus(id, format string, a ...interface{}) []byte
```

---

### NewJSONProgressOutput

NewJSONProgressOutput returns a progress.Output that formats output
using JSON objects

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/streamformatter/streamformatter.go#L100)  

```go
func NewJSONProgressOutput(out io.Writer, newLines bool) progress.Output
```

---

### NewProgressOutput

NewProgressOutput returns a progress.Output object that can be passed to
progress.NewProgressReader.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/streamformatter/streamformatter.go#L94)  

```go
func NewProgressOutput(out io.Writer) progress.Output
```

---

### NewStderrWriter

NewStderrWriter returns a writer which formats the output as json message
representing stderr lines

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/streamformatter/streamwriter.go#L43)  

```go
func NewStderrWriter(out io.Writer) io.Writer
```

---

### NewStdoutWriter

NewStdoutWriter returns a writer which formats the output as json message
representing stdout lines

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/streamformatter/streamwriter.go#L35)  

```go
func NewStdoutWriter(out io.Writer) io.Writer
```

---

## Types

### AuxFormatter

AuxFormatter is a streamFormatter that writes aux progress messages

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/streamformatter/streamformatter.go#L142)  

```go
type AuxFormatter struct {
	io.Writer
}
```

#### Methods

##### AuxFormatter.Emit

Emit emits the given interface as an aux progress message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/streamformatter/streamformatter.go#L147)  

```go
func (sf *AuxFormatter) Emit(id string, aux interface{}) error
```

---


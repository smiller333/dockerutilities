# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/jsonfilelog/jsonlog

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:18 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### JSONLog

JSONLog is a log message, typically a single entry from a given log stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonlog/jsonlog.go#L8)  

```go
type JSONLog struct {
	// Log is the log message
	Log string `json:"log,omitempty"`
	// Stream is the log source
	Stream string `json:"stream,omitempty"`
	// Created is the created timestamp of log
	Created time.Time `json:"time"`
	// Attrs is the list of extra attributes provided by the user
	Attrs map[string]string `json:"attrs,omitempty"`
}
```

#### Methods

##### JSONLog.Reset

Reset all fields to their zero value.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonlog/jsonlog.go#L20)  

```go
func (jl *JSONLog) Reset()
```

---

### JSONLogs

JSONLogs marshals encoded JSONLog objects

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonlog/jsonlogbytes.go#L11)  

```go
type JSONLogs struct {
	Log     []byte    `json:"log,omitempty"`
	Stream  string    `json:"stream,omitempty"`
	Created time.Time `json:"time"`

	// json-encoded bytes
	RawAttrs json.RawMessage `json:"attrs,omitempty"`
}
```

#### Methods

##### JSONLogs.MarshalJSONBuf

MarshalJSONBuf is an optimized JSON marshaller that avoids reflection
and unnecessary allocation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonlog/jsonlogbytes.go#L22)  

```go
func (mj *JSONLogs) MarshalJSONBuf(buf *bytes.Buffer) error
```

---


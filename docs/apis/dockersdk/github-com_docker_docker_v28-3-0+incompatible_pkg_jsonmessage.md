# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/jsonmessage

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:32 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L17)

```go
const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"
```

## Variables

This section is empty.

## Functions

### DisplayJSONMessagesStream

DisplayJSONMessagesStream reads a JSON message stream from in, and writes
each JSONMessage to out. It returns an error if an invalid JSONMessage
is received, or if a JSONMessage containers a non-zero [JSONMessage.Error].

Presentation of the JSONMessage depends on whether a terminal is attached,
and on the terminal width. Progress bars (JSONProgress) are suppressed
on narrower terminals (< 110 characters).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L235)  

```go
func DisplayJSONMessagesStream(in io.Reader, out io.Writer, terminalFd uintptr, isTerminal bool, auxCallback func(JSONMessage)) error
```

---

### DisplayJSONMessagesToStream

DisplayJSONMessagesToStream prints json messages to the output Stream. It is
used by the Docker CLI to print JSONMessage streams.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L312)  
**Added in:** v1.13.0

```go
func DisplayJSONMessagesToStream(in io.Reader, stream Stream, auxCallback func(JSONMessage)) error
```

---

## Types

### JSONError

JSONError wraps a concrete Code and Message, Code is
an integer error code, Message is the error message.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L21)  

```go
type JSONError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
```

#### Methods

##### JSONError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L26)  

```go
func (e *JSONError) Error() string
```

---

### JSONMessage

JSONMessage defines a message struct. It describes
the created time, where it from, status, ID of the
message. It's used for docker events.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L144)  

```go
type JSONMessage struct {
	Stream   string        `json:"stream,omitempty"`
	Status   string        `json:"status,omitempty"`
	Progress *JSONProgress `json:"progressDetail,omitempty"`

	// ProgressMessage is a pre-formatted presentation of [Progress].
	//
	// Deprecated: this field is deprecated since docker v0.7.1 / API v1.8. Use the information in [Progress] instead. This field will be omitted in a future release.
	ProgressMessage string     `json:"progress,omitempty"`
	ID              string     `json:"id,omitempty"`
	From            string     `json:"from,omitempty"`
	Time            int64      `json:"time,omitempty"`
	TimeNano        int64      `json:"timeNano,omitempty"`
	Error           *JSONError `json:"errorDetail,omitempty"`

	// ErrorMessage contains errors encountered during the operation.
	//
	// Deprecated: this field is deprecated since docker v0.6.0 / API v1.4. Use [Error.Message] instead. This field will be omitted in a future release.
	ErrorMessage string `json:"error,omitempty"` // deprecated
	// Aux contains out-of-band data, such as digests for push signing and image id after building.
	Aux *json.RawMessage `json:"aux,omitempty"`
}
```

#### Methods

##### JSONMessage.Display

Display prints the JSONMessage to out. If isTerminal is true, it erases
the entire current line when displaying the progressbar. It returns an
error if the [JSONMessage.Error] field is non-nil.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L184)  

```go
func (jm *JSONMessage) Display(out io.Writer, isTerminal bool) error
```

---

### JSONProgress

JSONProgress describes a progress message in a JSON stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L31)  

```go
type JSONProgress struct {
	// Current is the current status and value of the progress made towards Total.
	Current int64 `json:"current,omitempty"`
	// Total is the end value describing when we made 100% progress for an operation.
	Total int64 `json:"total,omitempty"`
	// Start is the initial value for the operation.
	Start int64 `json:"start,omitempty"`
	// HideCounts. if true, hides the progress count indicator (xB/yB).
	HideCounts bool `json:"hidecounts,omitempty"`
	// Units is the unit to print for progress. It defaults to "bytes" if empty.
	Units string `json:"units,omitempty"`
	// contains filtered or unexported fields
}
```

#### Methods

##### JSONProgress.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L54)  

```go
func (p *JSONProgress) String() string
```

---

### Stream

Stream is an io.Writer for output with utilities to get the output's file
descriptor and to detect whether it's a terminal.

it is subset of the streams.Out type in
https://pkg.go.dev/github.com/docker/cli@v20.10.17+incompatible/cli/streams#Out

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/jsonmessage/jsonmessage.go#L304)  

```go
type Stream interface {
	io.Writer
	FD() uintptr
	IsTerminal() bool
}
```

---


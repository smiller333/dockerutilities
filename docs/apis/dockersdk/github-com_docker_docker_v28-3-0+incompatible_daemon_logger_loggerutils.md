# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/loggerutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:25 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/log_tag.go#L11)

```go
const DefaultTemplate = "{{.ID}}"
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/queue.go#L67)

```go
var ErrQueueClosed = errors.New("queue is closed")
```

## Functions

### ParseLogTag

ParseLogTag generates a context aware tag for consistency across different
log drivers based on the context of the running container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/log_tag.go#L15)  

```go
func ParseLogTag(info logger.Info, defaultTemplate string) (string, error)
```

---

## Types

### Decoder

Decoder is for reading logs
It is created by the log reader by calling the `MakeDecoderFunc`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L98)  

```go
type Decoder interface {
	// Reset resets the decoder
	// Reset is called for certain events, such as log rotations
	Reset(io.Reader)
	// Decode decodes the next log message from the stream
	Decode() (*logger.Message, error)
	// Close signals to the decoder that it can release whatever resources it was using.
	Close()
}
```

---

### GetTailReaderFunc

GetTailReaderFunc is used to truncate a reader to only read as much as is required
in order to get the passed in number of log lines.
It returns the sectioned reader, the number of lines that the section reader
contains, and any error that occurs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L120)  

```go
type GetTailReaderFunc func(ctx context.Context, f SizeReaderAt, nLogLines int) (rdr SizeReaderAt, nLines int, err error)
```

---

### LogFile

LogFile is Logger implementation for default Docker logging.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L34)  

```go
type LogFile struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewLogFile

NewLogFile creates new LogFile

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L123)  

```go
func NewLogFile(logPath string, capacity int64, maxFiles int, compress bool, decodeFunc MakeDecoderFn, perms os.FileMode, getTailReader GetTailReaderFunc) (*LogFile, error)
```

#### Methods

##### LogFile.Close

Close closes underlying file and signals all readers to stop.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L360)  

```go
func (w *LogFile) Close() error
```

##### LogFile.MaxFiles

MaxFiles return maximum number of files

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L355)  

```go
func (w *LogFile) MaxFiles() int
```

##### LogFile.ReadLogs

ReadLogs decodes entries from log files.

It is the caller's responsibility to call ConsumerGone on the LogWatcher.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L381)  

```go
func (w *LogFile) ReadLogs(ctx context.Context, config logger.ReadConfig) *logger.LogWatcher
```

##### LogFile.WriteLogEntry

WriteLogEntry writes the provided log message to the current log file.
This may trigger a rotation event if the max file/capacity limits are hit.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L160)  

```go
func (w *LogFile) WriteLogEntry(timestamp time.Time, marshalled []byte) error
```

---

### MakeDecoderFn

MakeDecoderFn creates a decoder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L94)  

```go
type MakeDecoderFn func(rdr io.Reader) Decoder
```

---

### MessageQueue

MessageQueue is a queue for log messages.

MessageQueue.Enqueue will block when the queue is full.
To dequeue messages call MessageQueue.Receiver and pull messages off the
returned channel.

Closing only prevents new messages from being added to the queue.
The queue can still be drained after close.

The zero value of MessageQueue is safe to use, but does not do any internal
buffering (queue size is 0).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/queue.go#L22)  

```go
type MessageQueue struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewMessageQueue

NewMessageQueue creates a new queue with the specified size.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/queue.go#L42)  

```go
func NewMessageQueue(maxSize int) *MessageQueue
```

#### Methods

##### MessageQueue.Close

Close prevents any new messages from being added to the queue.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/queue.go#L116)  

```go
func (q *MessageQueue) Close()
```

##### MessageQueue.Enqueue

Enqueue adds the provided message to the queue.
Enqueue blocks if the queue is full.

The two possible error cases are:
1. The provided context is cancelled
2. ErrQueueClosed when the queue has been closed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/queue.go#L75)  

```go
func (q *MessageQueue) Enqueue(ctx context.Context, m *logger.Message) error
```

##### MessageQueue.Receiver

Receiver returns a channel that can be used to dequeue messages
The channel will be closed when the message queue is closed but may have
messages buffered.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/queue.go#L149)  

```go
func (q *MessageQueue) Receiver() <-chan *logger.Message
```

---

### SizeReaderAt

SizeReaderAt defines a ReaderAt that also reports its size.
This is used for tailing log files.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/logfile.go#L110)  

```go
type SizeReaderAt interface {
	io.Reader
	io.ReaderAt
	Size() int64
}
```

---


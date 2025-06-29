# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/container/stream

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:30 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### AttachConfig

AttachConfig is the config struct used to attach a client to a stream's stdio

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/attach.go#L17)  

```go
type AttachConfig struct {
	// Tells the attach copier that the stream's stdin is a TTY and to look for
	// escape sequences in stdin to detach from the stream.
	// When true the escape sequence is not passed to the underlying stream
	TTY bool
	// Specifies the detach keys the client will be using
	// Only useful when `TTY` is true
	DetachKeys []byte

	// CloseStdin signals that once done, stdin for the attached stream should be closed
	// For example, this would close the attached container's stdin.
	CloseStdin bool

	// UseStd* indicate whether the client has requested to be connected to the
	// given stream or not.  These flags are used instead of checking Std* != nil
	// at points before the client streams Std* are wired up.
	UseStdin, UseStdout, UseStderr bool

	// CStd* are the streams directly connected to the container
	CStdin           io.WriteCloser
	CStdout, CStderr io.ReadCloser

	// Provide client streams to wire up to
	Stdin          io.ReadCloser
	Stdout, Stderr io.Writer
}
```

---

### Config

Config holds information about I/O streams managed together.

config.StdinPipe returns a WriteCloser which can be used to feed data
to the standard input of the streamConfig's active process.
config.StdoutPipe and streamConfig.StderrPipe each return a ReadCloser
which can be used to retrieve the standard output (and error) generated
by the container's active process. The output (and error) are actually
copied and delivered to all StdoutPipe and StderrPipe consumers, using
a kind of "broadcaster".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L26)  

```go
type Config struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewConfig

NewConfig creates a stream config and initializes
the standard err and standard out to new unbuffered broadcasters.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L39)  

```go
func NewConfig() *Config
```

#### Methods

##### Config.AttachStreams

AttachStreams attaches the container's streams to the AttachConfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/attach.go#L45)  

```go
func (c *Config) AttachStreams(cfg *AttachConfig)
```

##### Config.CloseStreams

CloseStreams ensures that the configured streams are properly closed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L101)  

```go
func (c *Config) CloseStreams() error
```

##### Config.CopyStreams

CopyStreams starts goroutines to copy data in and out to/from the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/attach.go#L60)  

```go
func (c *Config) CopyStreams(ctx context.Context, cfg *AttachConfig) <-chan error
```

##### Config.CopyToPipe

CopyToPipe connects streamconfig with a libcontainerd.IOPipe

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L124)  

```go
func (c *Config) CopyToPipe(iop *cio.DirectIO)
```

##### Config.NewInputPipes

NewInputPipes creates new pipes for both standard inputs, Stdin and StdinPipe.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L85)  

```go
func (c *Config) NewInputPipes()
```

##### Config.NewNopInputPipe

NewNopInputPipe creates a new input pipe that will silently drop all messages in the input.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L90)  

```go
func (c *Config) NewNopInputPipe()
```

##### Config.Stderr

Stderr returns the standard error in the configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L52)  

```go
func (c *Config) Stderr() io.Writer
```

##### Config.StderrPipe

StderrPipe creates a new io.ReadCloser with an empty bytes pipe.
It adds this new err pipe to the Stderr broadcaster.
This will block stderr if unconsumed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L78)  

```go
func (c *Config) StderrPipe() io.ReadCloser
```

##### Config.Stdin

Stdin returns the standard input in the configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L57)  

```go
func (c *Config) Stdin() io.ReadCloser
```

##### Config.StdinPipe

StdinPipe returns an input writer pipe as an io.WriteCloser.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L62)  

```go
func (c *Config) StdinPipe() io.WriteCloser
```

##### Config.Stdout

Stdout returns the standard output in the configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L47)  

```go
func (c *Config) Stdout() io.Writer
```

##### Config.StdoutPipe

StdoutPipe creates a new io.ReadCloser with an empty bytes pipe.
It adds this new out pipe to the Stdout broadcaster.
This will block stdout if unconsumed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L69)  

```go
func (c *Config) StdoutPipe() io.ReadCloser
```

##### Config.Wait

Wait for the stream to close
Wait supports timeouts via the context to unblock and forcefully
close the io streams

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/stream/streams.go#L172)  

```go
func (c *Config) Wait(ctx context.Context)
```

---


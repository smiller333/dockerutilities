# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:28:55 UTC

## Overview

Package logger defines interfaces that logger drivers implement to
log messages.

The other half of a logger driver is the implementation of the
factory, which holds the contextual instance information that
allows multiple loggers of the same type to perform different
actions, such as logging to different locations.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AddBuiltinLogOpts

AddBuiltinLogOpts updates the list of built-in log opts. This allows other packages to supplement additional log options
without having to register an actual log driver. This is used by things that are more proxy log drivers and should
not be exposed as a usable log driver to the API.
This should only be called on package initialization.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/log_cache_opts.go#L16)  

```go
func AddBuiltinLogOpts(opts map[string]bool)
```

---

### ListDrivers

ListDrivers gets the list of registered log driver names

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/factory.go#L39)  

```go
func ListDrivers() []string
```

---

### PutMessage

PutMessage puts the specified message back n the message pool.
The message fields are reset before putting into the pool.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L41)  

```go
func PutMessage(msg *Message)
```

---

### RegisterExternalValidator

RegisterExternalValidator adds the validator to the list of external validators.
External validators are used by packages outside this package that need to add their own validation logic.
This should only be called on package initialization.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/log_cache_opts.go#L8)  

```go
func RegisterExternalValidator(v LogOptValidator)
```

---

### RegisterLogDriver

RegisterLogDriver registers the given logging driver builder with given logging
driver name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/factory.go#L112)  
**Added in:** v1.7.0

```go
func RegisterLogDriver(name string, c Creator) error
```

---

### RegisterLogOptValidator

RegisterLogOptValidator registers the logging option validator with
the given logging driver name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/factory.go#L118)  
**Added in:** v1.8.0

```go
func RegisterLogOptValidator(name string, l LogOptValidator) error
```

---

### RegisterPluginGetter

RegisterPluginGetter sets the plugingetter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/plugin.go#L30)  

```go
func RegisterPluginGetter(g plugingetter.PluginGetter)
```

---

### ValidateLogOpts

ValidateLogOpts checks the options for the given log driver. The
options supported are specific to the LogDriver implementation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/factory.go#L134)  
**Added in:** v1.8.0

```go
func ValidateLogOpts(name string, cfg map[string]string) error
```

---

## Types

### Capability

Capability defines the list of capabilities that a driver can implement
These capabilities are not required to be a logging driver, however do
determine how a logging driver can be used

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L131)  

```go
type Capability struct {
	// Determines if a log driver can read back logs
	ReadLogs bool
}
```

---

### Copier

Copier can copy logs from specified sources to Logger and attach Timestamp.
Writes are concurrent, so you need implement some sync in your logger.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/copier.go#L27)  

```go
type Copier struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewCopier

NewCopier creates a new Copier

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/copier.go#L37)  

```go
func NewCopier(srcs map[string]io.Reader, dst Logger) *Copier
```

#### Methods

##### Copier.Close

Close closes the copier

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/copier.go#L186)  
**Added in:** v1.10.0

```go
func (c *Copier) Close()
```

##### Copier.Run

Run starts logs copying

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/copier.go#L46)  

```go
func (c *Copier) Run()
```

##### Copier.Wait

Wait waits until all copying is done

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/copier.go#L181)  

```go
func (c *Copier) Wait()
```

---

### Creator

Creator builds a logging driver instance with given context.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/factory.go#L15)  
**Added in:** v1.7.0

```go
type Creator func(Info) (Logger, error)
```

#### Functions

##### GetLogDriver

GetLogDriver provides the logging driver builder for a logging driver name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/factory.go#L123)  
**Added in:** v1.7.0

```go
func GetLogDriver(name string) (Creator, error)
```

---

### ErrReadLogsNotSupported

ErrReadLogsNotSupported is returned when the underlying log driver does not support reading

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L19)  
**Added in:** v1.8.0

```go
type ErrReadLogsNotSupported struct{}
```

#### Methods

##### ErrReadLogsNotSupported.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L21)  

```go
func (ErrReadLogsNotSupported) Error() string
```

##### ErrReadLogsNotSupported.NotImplemented

NotImplemented makes this error implement the `NotImplemented` interface from api/errdefs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L26)  

```go
func (ErrReadLogsNotSupported) NotImplemented()
```

---

### Info

Info provides enough information for a logging driver to do its function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L12)  

```go
type Info struct {
	Config              map[string]string
	ContainerID         string
	ContainerName       string
	ContainerEntrypoint string
	ContainerArgs       []string
	ContainerImageID    string
	ContainerImageName  string
	ContainerCreated    time.Time
	ContainerEnv        []string
	ContainerLabels     map[string]string
	LogPath             string
	DaemonName          string
}
```

#### Methods

##### Info.Command

Command returns the command that the container being logged was
started with. The Entrypoint is prepended to the container
arguments.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L112)  

```go
func (info *Info) Command() string
```

##### Info.ExtraAttributes

ExtraAttributes returns the user-defined extra attributes (labels,
environment variables) in key-value format. This can be used by log drivers
that support metadata to add more context to a log.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L30)  

```go
func (info *Info) ExtraAttributes(keyMod func(string) string) (map[string]string, error)
```

##### Info.FullID

FullID is an alias of ContainerID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L125)  

```go
func (info *Info) FullID() string
```

##### Info.Hostname

Hostname returns the hostname from the underlying OS.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L101)  

```go
func (info *Info) Hostname() (string, error)
```

##### Info.ID

ID Returns the Container ID shortened to 12 characters.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L120)  

```go
func (info *Info) ID() string
```

##### Info.ImageFullID

ImageFullID is an alias of ContainerImageID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L140)  

```go
func (info *Info) ImageFullID() string
```

##### Info.ImageID

ImageID returns the ContainerImageID shortened to 12 characters.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L135)  

```go
func (info *Info) ImageID() string
```

##### Info.ImageName

ImageName is an alias of ContainerImageName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L145)  

```go
func (info *Info) ImageName() string
```

##### Info.Name

Name returns the ContainerName without a preceding '/'.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loginfo.go#L130)  

```go
func (info *Info) Name() string
```

---

### LogOptValidator

LogOptValidator checks the options specific to the underlying
logging implementation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/factory.go#L19)  
**Added in:** v1.8.0

```go
type LogOptValidator func(cfg map[string]string) error
```

---

### LogReader

LogReader is the interface for reading log messages for loggers that support reading.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L90)  
**Added in:** v1.8.0

```go
type LogReader interface {
	// ReadLogs reads logs from underlying logging backend.
	ReadLogs(context.Context, ReadConfig) *LogWatcher
}
```

---

### LogWatcher

LogWatcher is used when consuming logs read from the LogReader interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L96)  
**Added in:** v1.8.0

```go
type LogWatcher struct {
	// For sending log messages to a reader.
	Msg chan *Message
	// For sending error messages that occur while reading logs.
	Err chan error
	// contains filtered or unexported fields
}
```

#### Functions

##### NewLogWatcher

NewLogWatcher returns a new LogWatcher.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L106)  
**Added in:** v1.8.0

```go
func NewLogWatcher() *LogWatcher
```

#### Methods

##### LogWatcher.ConsumerGone

ConsumerGone notifies that the logs consumer is gone.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L115)  

```go
func (w *LogWatcher) ConsumerGone()
```

##### LogWatcher.WatchConsumerGone

WatchConsumerGone returns a channel receiver that receives notification
when the log watcher consumer is gone.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L124)  

```go
func (w *LogWatcher) WatchConsumerGone() <-chan struct{}
```

---

### Logger

Logger is the interface for docker logging drivers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L68)  

```go
type Logger interface {
	Log(*Message) error
	Name() string
	Close() error
}
```

#### Functions

##### NewRingLogger

NewRingLogger creates a new Logger that is implemented as a RingBuffer wrapping
the passed in logger.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/ring.go#L55)  

```go
func NewRingLogger(driver Logger, logInfo Info, maxSize int64) Logger
```

---

### Message

Message is data structure that represents piece of output produced by some
container.  The Line member is a slice of an array whose contents can be
changed after a log driver's Log() method returns.

Message is subtyped from backend.LogMessage because there is a lot of
internal complexity around the Message type that should not be exposed
to any package not explicitly importing the logger type.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L53)  

```go
type Message backend.LogMessage
```

#### Functions

##### NewMessage

NewMessage returns a new message from the message sync.Pool

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L35)  

```go
func NewMessage() *Message
```

#### Methods

##### Message.AsLogMessage

AsLogMessage returns a pointer to the message as a pointer to
backend.LogMessage, which is an identical type with a different purpose

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L63)  

```go
func (m *Message) AsLogMessage() *backend.LogMessage
```

---

### ReadConfig

ReadConfig is the configuration passed into ReadLogs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L82)  
**Added in:** v1.8.0

```go
type ReadConfig struct {
	Since  time.Time
	Until  time.Time
	Tail   int
	Follow bool
}
```

---

### SizedLogger

SizedLogger is the interface for logging drivers that can control
the size of buffer used for their messages.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/logger.go#L76)  

```go
type SizedLogger interface {
	Logger
	BufSize() int
}
```

---


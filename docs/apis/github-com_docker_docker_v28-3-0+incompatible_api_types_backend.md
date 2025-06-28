# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/backend

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:43 UTC

## Overview

Package backend includes types to send information to server backends.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### BuildConfig

BuildConfig is the configuration used by a BuildManager to start a build

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/build.go#L34)  

```go
type BuildConfig struct {
	Source         io.ReadCloser
	ProgressWriter ProgressWriter
	Options        *build.ImageBuildOptions
}
```

---

### CommitConfig

CommitConfig is the configuration for creating an image as part of a build.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L160)  

```go
type CommitConfig struct {
	Author              string
	Comment             string
	Config              *container.Config // TODO(thaJeztah); change this to [dockerspec.DockerOCIImageConfig]
	ContainerConfig     *container.Config
	ContainerID         string
	ContainerMountLabel string
	ContainerOS         string
	ParentImageID       string
}
```

---

### ContainerAttachConfig

ContainerAttachConfig holds the streams to use when connecting to a container to view logs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L32)  

```go
type ContainerAttachConfig struct {
	GetStreams func(multiplexed bool, cancel func()) (io.ReadCloser, io.Writer, io.Writer, error)
	UseStdin   bool
	UseStdout  bool
	UseStderr  bool
	Logs       bool
	Stream     bool
	DetachKeys string
	// Used to signify that streams must be multiplexed by producer as endpoint can't manage multiple streams.
	// This is typically set by HTTP endpoint, while websocket can transport raw streams
	MuxStreams bool
}
```

---

### ContainerCreateConfig

ContainerCreateConfig is the parameter set to ContainerCreate()

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L15)  

```go
type ContainerCreateConfig struct {
	Name                        string
	Config                      *container.Config
	HostConfig                  *container.HostConfig
	NetworkingConfig            *network.NetworkingConfig
	Platform                    *ocispec.Platform
	DefaultReadOnlyNonRecursive bool
}
```

---

### ContainerInspectOptions

ContainerInspectOptions defines options for the backend.ContainerInspect
call.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L97)  

```go
type ContainerInspectOptions struct {
	// Size controls whether to propagate the container's size fields.
	Size bool
}
```

---

### ContainerRmConfig

ContainerRmConfig holds arguments for the container remove
operation. This struct is used to tell the backend what operations
to perform.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L27)  

```go
type ContainerRmConfig struct {
	ForceRemove, RemoveVolume, RemoveLink bool
}
```

---

### ContainerStatsConfig

ContainerStatsConfig holds information for configuring the runtime
behavior of a backend.ContainerStats() call.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L89)  

```go
type ContainerStatsConfig struct {
	Stream    bool
	OneShot   bool
	OutStream func() io.Writer
}
```

---

### CreateImageConfig

CreateImageConfig is the configuration for creating an image from a
container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L138)  

```go
type CreateImageConfig struct {
	Tag     reference.NamedTagged
	Pause   bool
	Author  string
	Comment string
	Config  *container.Config
	Changes []string
}
```

---

### ExecInspect

ExecInspect holds information about a running process started
with docker exec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L112)  

```go
type ExecInspect struct {
	ID            string
	Running       bool
	ExitCode      *int
	ProcessConfig *ExecProcessConfig
	OpenStdin     bool
	OpenStderr    bool
	OpenStdout    bool
	CanRemove     bool
	ContainerID   string
	DetachKeys    []byte
	Pid           int
}
```

---

### ExecProcessConfig

ExecProcessConfig holds information about the exec process
running on the host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L128)  

```go
type ExecProcessConfig struct {
	Tty        bool     `json:"tty"`
	Entrypoint string   `json:"entrypoint"`
	Arguments  []string `json:"arguments"`
	Privileged *bool    `json:"privileged,omitempty"`
	User       string   `json:"user,omitempty"`
}
```

---

### ExecStartConfig

ExecStartConfig holds the options to start container's exec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L103)  

```go
type ExecStartConfig struct {
	Stdin       io.Reader
	Stdout      io.Writer
	Stderr      io.Writer
	ConsoleSize *[2]uint `json:",omitempty"`
}
```

---

### GetImageAndLayerOptions

GetImageAndLayerOptions are the options supported by GetImageAndReleasableLayer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/build.go#L41)  

```go
type GetImageAndLayerOptions struct {
	PullOption PullOption
	AuthConfig map[string]registry.AuthConfig
	Output     io.Writer
	Platform   *ocispec.Platform
}
```

---

### GetImageOpts

GetImageOpts holds parameters to retrieve image information
from the backend.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L149)  

```go
type GetImageOpts struct {
	Platform *ocispec.Platform
}
```

---

### ImageInspectOpts

ImageInspectOpts holds parameters to inspect an image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L154)  

```go
type ImageInspectOpts struct {
	Manifests bool
	Platform  *ocispec.Platform
}
```

---

### LogAttr

LogAttr is used to hold the extra attributes available in the log message.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L72)  

```go
type LogAttr struct {
	Key   string
	Value string
}
```

---

### LogMessage

LogMessage is datastructure that represents piece of output produced by some
container.  The Line member is a slice of an array whose contents can be
changed after a log driver's Log() method returns.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L58)  

```go
type LogMessage struct {
	Line         []byte
	Source       string
	Timestamp    time.Time
	Attrs        []LogAttr
	PLogMetaData *PartialLogMetaData

	// Err is an error associated with a message. Completeness of a message
	// with Err is not expected, tho it may be partially complete (fields may
	// be missing, gibberish, or nil)
	Err error
}
```

---

### LogSelector

LogSelector is a list of services and tasks that should be returned as part
of a log stream. It is similar to swarmapi.LogSelector, with the difference
that the names don't have to be resolved to IDs; this is mostly to avoid
accidents later where a swarmapi LogSelector might have been incorrectly
used verbatim (and to avoid the handler having to import swarmapi types)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L82)  

```go
type LogSelector struct {
	Services []string
	Tasks    []string
}
```

---

### NetworkListConfig

NetworkListConfig stores the options available for listing networks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L187)  

```go
type NetworkListConfig struct {
	// TODO(@cpuguy83): naming is hard, this is pulled from what was being used in the router before moving here
	Detailed bool
	Verbose  bool
}
```

---

### PartialLogMetaData

PartialLogMetaData provides meta data for a partial log message. Messages
exceeding a predefined size are split into chunks with this metadata. The
expectation is for the logger endpoints to assemble the chunks using this
metadata.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L49)  

```go
type PartialLogMetaData struct {
	Last    bool   // true if this message is last of a partial
	ID      string // identifies group of messages comprising a single record
	Ordinal int    // ordering of message in partial group
}
```

---

### PluginDisableConfig

PluginDisableConfig holds arguments for plugin disable.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L182)  

```go
type PluginDisableConfig struct {
	ForceDisable bool
}
```

---

### PluginEnableConfig

PluginEnableConfig holds arguments for plugin enable

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L177)  

```go
type PluginEnableConfig struct {
	Timeout int
}
```

---

### PluginRmConfig

PluginRmConfig holds arguments for plugin remove.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/backend.go#L172)  

```go
type PluginRmConfig struct {
	ForceRemove bool
}
```

---

### ProgressWriter

ProgressWriter is a data object to transport progress streams to the client

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/build.go#L25)  
**Added in:** v1.12.0

```go
type ProgressWriter struct {
	Output             io.Writer
	StdoutFormatter    io.Writer
	StderrFormatter    io.Writer
	AuxFormatter       *streamformatter.AuxFormatter
	ProgressReaderFunc func(io.ReadCloser) io.ReadCloser
}
```

---

### PullOption

PullOption defines different modes for accessing images

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/backend/build.go#L13)  

```go
type PullOption int
```

---


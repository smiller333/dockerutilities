# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libcontainerd/supervisor

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:08:24 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Daemon

Daemon represents a running containerd daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/supervisor/remote_daemon.go#L56)  

```go
type Daemon interface {
	WaitTimeout(time.Duration) error
	Address() string
}
```

#### Functions

##### Start

Start starts a containerd daemon and monitors it

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/supervisor/remote_daemon.go#L65)  

```go
func Start(ctx context.Context, rootDir, stateDir string, opts ...DaemonOpt) (Daemon, error)
```

---

### DaemonOpt

DaemonOpt allows to configure parameters of container daemons

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/supervisor/remote_daemon.go#L62)  

```go
type DaemonOpt func(c *remote) error
```

#### Functions

##### WithCRIDisabled

WithCRIDisabled disables the CRI plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/supervisor/remote_daemon_options.go#L34)  

```go
func WithCRIDisabled() DaemonOpt
```

##### WithDetectLocalBinary

WithDetectLocalBinary checks if a containerd binary is present in the same
directory as the dockerd binary, and overrides the path of the containerd
binary to start if found. If no binary is found, no changes are made.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/supervisor/remote_daemon_options.go#L44)  

```go
func WithDetectLocalBinary() DaemonOpt
```

##### WithLogFormat

WithLogFormat defines the containerd log format.
This only makes sense if WithStartDaemon() was set to true.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/supervisor/remote_daemon_options.go#L26)  

```go
func WithLogFormat(format log.OutputFormat) DaemonOpt
```

##### WithLogLevel

WithLogLevel defines which log level to start containerd with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/supervisor/remote_daemon_options.go#L12)  

```go
func WithLogLevel(lvl string) DaemonOpt
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/plugin/executor/containerd

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:52 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Executor

Executor is the containerd client implementation of a plugin executor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/executor/containerd/containerd.go#L45)  

```go
type Executor struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates a new containerd plugin executor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/executor/containerd/containerd.go#L27)  

```go
func New(ctx context.Context, rootDir string, cli *containerd.Client, ns string, exitHandler ExitHandler, shim string, shimOpts interface{}) (*Executor, error)
```

#### Methods

##### Executor.Create

Create creates a new container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/executor/containerd/containerd.go#L77)  

```go
func (e *Executor) Create(id string, spec specs.Spec, stdout, stderr io.WriteCloser) error
```

##### Executor.IsRunning

IsRunning returns if the container with the given id is running

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/executor/containerd/containerd.go#L140)  

```go
func (e *Executor) IsRunning(id string) (bool, error)
```

##### Executor.ProcessEvent

ProcessEvent handles events from containerd
All events are ignored except the exit event, which is sent of to the stored handler

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/executor/containerd/containerd.go#L164)  

```go
func (e *Executor) ProcessEvent(id string, et libcontainerdtypes.EventType, ei libcontainerdtypes.EventInfo) error
```

##### Executor.Restore

Restore restores a container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/executor/containerd/containerd.go#L101)  

```go
func (e *Executor) Restore(id string, stdout, stderr io.WriteCloser) (bool, error)
```

##### Executor.Signal

Signal sends the specified signal to the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/executor/containerd/containerd.go#L152)  

```go
func (e *Executor) Signal(id string, signal syscall.Signal) error
```

---

### ExitHandler

ExitHandler represents an object that is called when the exit event is received from containerd

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/executor/containerd/containerd.go#L22)  

```go
type ExitHandler interface {
	HandleExitEvent(id string) error
}
```

---


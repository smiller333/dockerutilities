# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libcontainerd/types

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:08:27 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Backend

Backend defines callbacks that the client of the library needs to implement.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types.go#L40)  

```go
type Backend interface {
	ProcessEvent(containerID string, event EventType, ei EventInfo) error
}
```

---

### Checkpoints

Checkpoints contains the details of a checkpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types_linux.go#L33)  

```go
type Checkpoints struct{}
```

---

### Client

Client provides access to containerd features.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types.go#L57)  

```go
type Client interface {
	Version(ctx context.Context) (containerd.Version, error)
	// LoadContainer loads the metadata for a container from containerd.
	LoadContainer(ctx context.Context, containerID string) (Container, error)
	// NewContainer creates a new containerd container.
	NewContainer(ctx context.Context, containerID string, spec *specs.Spec, shim string, runtimeOptions interface{}, opts ...containerd.NewContainerOpts) (Container, error)
}
```

---

### Container

Container provides access to a containerd container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types.go#L66)  

```go
type Container interface {
	NewTask(ctx context.Context, checkpointDir string, withStdin bool, attachStdio StdioCallback) (Task, error)
	Task(ctx context.Context) (Task, error)
	// AttachTask returns the current task for the container and reattaches
	// to the IO for the running task. If no task exists for the container
	// a NotFound error is returned.
	//
	// Clients must make sure that only one reader is attached to the task.
	AttachTask(ctx context.Context, attachStdio StdioCallback) (Task, error)
	// Delete removes the container and associated resources
	Delete(context.Context) error
}
```

---

### EventInfo

EventInfo contains the event info

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types.go#L30)  

```go
type EventInfo struct {
	ContainerID string
	ProcessID   string
	Pid         uint32
	ExitCode    uint32
	ExitedAt    time.Time
	Error       error
}
```

---

### EventType

EventType represents a possible event from libcontainerd

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types.go#L14)  

```go
type EventType string
```

---

### Process

Process of a container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types.go#L45)  

```go
type Process interface {
	// Pid is the system specific process id
	Pid() uint32
	// Kill sends the provided signal to the process
	Kill(ctx context.Context, signal syscall.Signal) error
	// Resize changes the width and height of the process's terminal
	Resize(ctx context.Context, width, height uint32) error
	// Delete removes the process and any resources allocated returning the exit status
	Delete(context.Context) (*containerd.ExitStatus, error)
}
```

---

### Resources

Resources defines updatable container resource values. TODO: it must match containerd upcoming API

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types_linux.go#L30)  

```go
type Resources = specs.LinuxResources
```

---

### Stats

Stats holds metrics properties as returned by containerd

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types_linux.go#L13)  

```go
type Stats struct {
	Read time.Time
	// Metrics is expected to be either one of:
	// * github.com/containerd/cgroups/v3/cgroup1/stats.Metrics
	// * github.com/containerd/cgroups/v3/cgroup2/stats.Metrics
	Metrics interface{}
}
```

#### Functions

##### InterfaceToStats

InterfaceToStats returns a stats object from the platform-specific interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types_linux.go#L22)  

```go
func InterfaceToStats(read time.Time, v interface{}) *Stats
```

---

### StdioCallback

StdioCallback is called to connect a container or process stdio.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types.go#L103)  

```go
type StdioCallback func(io *cio.DirectIO) (cio.IO, error)
```

---

### Summary

Summary is not used on linux

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types_linux.go#L10)  

```go
type Summary struct{}
```

---

### Task

Task provides access to a running containerd container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/types/types.go#L80)  

```go
type Task interface {
	Process
	// Start begins execution of the task
	Start(context.Context) error
	// Pause suspends the execution of the task
	Pause(context.Context) error
	// Resume the execution of the task
	Resume(context.Context) error
	Stats(ctx context.Context) (*Stats, error)
	// Pids returns a list of system specific process ids inside the task
	Pids(context.Context) ([]containerd.ProcessInfo, error)
	Summary(ctx context.Context) ([]Summary, error)
	// ForceDelete forcefully kills the task's processes and deletes the task
	ForceDelete(context.Context) error
	// Status returns the executing status of the task
	Status(ctx context.Context) (containerd.Status, error)
	// Exec creates and starts a new process inside the task
	Exec(ctx context.Context, processID string, spec *specs.Process, withStdin bool, attachStdio StdioCallback) (Process, error)
	UpdateResources(ctx context.Context, resources *Resources) error
	CreateCheckpoint(ctx context.Context, checkpointDir string, exit bool) error
}
```

---


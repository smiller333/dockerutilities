# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/cluster/executor/container

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:55 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/executor/container/errors.go#L7)

```go
var (
	// ErrImageRequired returned if a task is missing the image definition.
	ErrImageRequired = errors.New("dockerexec: image required")

	// ErrContainerDestroyed returned when a container is prematurely destroyed
	// during a wait call.
	ErrContainerDestroyed = errors.New("dockerexec: container destroyed")

	// ErrContainerUnhealthy returned if controller detects the health check failure
	ErrContainerUnhealthy = errors.New("dockerexec: unhealthy container")
)
```

## Functions

### NewExecutor

NewExecutor returns an executor from the docker client.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/executor/container/executor.go#L48)  

```go
func NewExecutor(b executorpkg.Backend, p plugin.Backend, i executorpkg.ImageBackend, v executorpkg.VolumeBackend) exec.Executor
```

---

## Types

This section is empty.


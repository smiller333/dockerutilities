# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/restartmanager

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:11 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/restartmanager/restartmanager.go#L20)

```go
var ErrRestartCanceled = errors.New("restart canceled")
```

## Functions

This section is empty.

## Types

### RestartManager

RestartManager defines object that controls container restarting rules.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/restartmanager/restartmanager.go#L23)  

```go
type RestartManager struct {
	sync.Mutex
	sync.Once
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New returns a new RestartManager based on a policy.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/restartmanager/restartmanager.go#L35)  

```go
func New(policy container.RestartPolicy, restartCount int) *RestartManager
```

#### Methods

##### RestartManager.Cancel

Cancel tells the RestartManager to no longer restart the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/restartmanager/restartmanager.go#L126)  

```go
func (rm *RestartManager) Cancel()
```

##### RestartManager.SetPolicy

SetPolicy sets the restart-policy for the RestartManager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/restartmanager/restartmanager.go#L40)  

```go
func (rm *RestartManager) SetPolicy(policy container.RestartPolicy)
```

##### RestartManager.ShouldRestart

ShouldRestart returns whether the container should be restarted.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/restartmanager/restartmanager.go#L47)  

```go
func (rm *RestartManager) ShouldRestart(exitCode uint32, hasBeenManuallyStopped bool, executionDuration time.Duration) (bool, chan error, error)
```

---


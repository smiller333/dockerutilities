# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration-cli/daemon

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:56 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### WaitInspectWithArgs

WaitInspectWithArgs waits for the specified expression to be equals to the specified expected string in the given time.
Deprecated: use cli.WaitCmd instead

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L103)  

```go
func WaitInspectWithArgs(dockerBinary, name, expr, expected string, timeout time.Duration, arg ...string) error
```

---

## Types

### Daemon

Daemon represents a Docker daemon for the testing framework.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L17)  

```go
type Daemon struct {
	*daemon.Daemon
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New returns a Daemon instance to be used for testing.
This will create a directory such as d123456789 in the folder specified by $DOCKER_INTEGRATION_DAEMON_DEST or $DEST.
The daemon will not automatically start.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L25)  

```go
func New(t testing.TB, dockerBinary string, dockerdBinary string, ops ...daemon.Option) *Daemon
```

#### Methods

##### Daemon.CheckActiveContainerCount

CheckActiveContainerCount returns the number of active containers
FIXME(vdemeester) should re-use ActivateContainers in some way

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L83)  

```go
func (d *Daemon) CheckActiveContainerCount(ctx context.Context) func(t *testing.T) (interface{}, string)
```

##### Daemon.CheckControlAvailable

CheckControlAvailable returns the current swarm control available

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L164)  

```go
func (d *Daemon) CheckControlAvailable(ctx context.Context) func(t *testing.T) (interface{}, string)
```

##### Daemon.CheckLeader

CheckLeader returns whether there is a leader on the swarm or not

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L173)  

```go
func (d *Daemon) CheckLeader(ctx context.Context) func(t *testing.T) (interface{}, string)
```

##### Daemon.CheckLocalNodeState

CheckLocalNodeState returns the current swarm node state

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L156)  

```go
func (d *Daemon) CheckLocalNodeState(ctx context.Context) func(t *testing.T) (interface{}, string)
```

##### Daemon.CheckNodeReadyCount

CheckNodeReadyCount returns the number of ready node on the swarm

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L142)  

```go
func (d *Daemon) CheckNodeReadyCount(ctx context.Context) func(t *testing.T) (interface{}, string)
```

##### Daemon.CheckPluginImage

CheckPluginImage returns the runtime state of the plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L79)  

```go
func (d *Daemon) CheckPluginImage(ctx context.Context, plugin string) func(c *testing.T) (interface{}, string)
```

##### Daemon.CheckPluginRunning

CheckPluginRunning returns the runtime state of the plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L66)  

```go
func (d *Daemon) CheckPluginRunning(ctx context.Context, plugin string) func(c *testing.T) (interface{}, string)
```

##### Daemon.CheckRunningTaskImages

CheckRunningTaskImages returns the times each image is running as a task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L121)  

```go
func (d *Daemon) CheckRunningTaskImages(ctx context.Context) func(t *testing.T) (interface{}, string)
```

##### Daemon.CheckRunningTaskNetworks

CheckRunningTaskNetworks returns the number of times each network is referenced from a task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L100)  

```go
func (d *Daemon) CheckRunningTaskNetworks(ctx context.Context) func(t *testing.T) (interface{}, string)
```

##### Daemon.CheckServiceRunningTasks

CheckServiceRunningTasks returns the number of running tasks for the specified service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L50)  

```go
func (d *Daemon) CheckServiceRunningTasks(ctx context.Context, service string) func(*testing.T) (interface{}, string)
```

##### Daemon.CheckServiceTasks

CheckServiceTasks returns the number of tasks for the specified service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L92)  

```go
func (d *Daemon) CheckServiceTasks(ctx context.Context, service string) func(*testing.T) (interface{}, string)
```

##### Daemon.CheckServiceTasksInState

CheckServiceTasksInState returns the number of tasks with a matching state,
and optional message substring.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L17)  

```go
func (d *Daemon) CheckServiceTasksInState(ctx context.Context, service string, state swarm.TaskState, message string) func(*testing.T) (interface{}, string)
```

##### Daemon.CheckServiceTasksInStateWithError

CheckServiceTasksInStateWithError returns the number of tasks with a matching state,
and optional message substring.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L34)  

```go
func (d *Daemon) CheckServiceTasksInStateWithError(ctx context.Context, service string, state swarm.TaskState, errorMessage string) func(*testing.T) (interface{}, string)
```

##### Daemon.CheckServiceUpdateState

CheckServiceUpdateState returns the current update state for the specified service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L55)  

```go
func (d *Daemon) CheckServiceUpdateState(ctx context.Context, service string) func(*testing.T) (interface{}, string)
```

##### Daemon.Cmd

Cmd executes a docker CLI command against this daemon.
Example: d.Cmd("version") will run docker -H unix://path/to/unix.sock version

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L37)  

```go
func (d *Daemon) Cmd(args ...string) (string, error)
```

##### Daemon.CmdRetryOutOfSequence

CmdRetryOutOfSequence tries the specified command against the current daemon
up to 10 times, retrying if it encounters an "update out of sequence" error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon_swarm.go#L196)  

```go
func (d *Daemon) CmdRetryOutOfSequence(args ...string) (string, error)
```

##### Daemon.Command

Command creates a docker CLI command against this daemon, to be executed later.
Example: d.Command("version") creates a command to run "docker -H unix://path/to/unix.sock version"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L44)  

```go
func (d *Daemon) Command(args ...string) icmd.Cmd
```

##### Daemon.GetIDByName

GetIDByName returns the ID of an object (container, volume, …) given its name

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L59)  

```go
func (d *Daemon) GetIDByName(name string) (string, error)
```

##### Daemon.InspectField

InspectField returns the field filter by 'filter'

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L64)  

```go
func (d *Daemon) InspectField(name, filter string) (string, error)
```

##### Daemon.PrependHostArg

PrependHostArg prepend the specified arguments by the daemon host flags

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L49)  

```go
func (d *Daemon) PrependHostArg(args []string) []string
```

##### Daemon.WaitRun

WaitRun waits for a container to be running for 10s

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/daemon/daemon.go#L96)  

```go
func (d *Daemon) WaitRun(contID string) error
```

---


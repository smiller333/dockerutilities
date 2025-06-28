# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration/internal/swarm

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:06:25 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CreateService

CreateService creates a service on the passed in swarm daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L60)  

```go
func CreateService(ctx context.Context, t *testing.T, d *daemon.Daemon, opts ...ServiceSpecOpt) string
```

---

### CreateServiceSpec

CreateServiceSpec creates a default service-spec, and applies the provided options

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L73)  

```go
func CreateServiceSpec(t *testing.T, opts ...ServiceSpecOpt) swarmtypes.ServiceSpec
```

---

### ExecTask

ExecTask runs the passed in exec config on the given task

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L215)  

```go
func ExecTask(ctx context.Context, t *testing.T, d *daemon.Daemon, task swarmtypes.Task, options container.ExecOptions) types.HijackedResponse
```

---

### GetRunningTasks

GetRunningTasks gets the list of running tasks for a service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L200)  

```go
func GetRunningTasks(ctx context.Context, t *testing.T, c client.ServiceAPIClient, serviceID string) []swarmtypes.Task
```

---

### JobComplete

JobComplete is a poll function for determining that a ReplicatedJob is
completed additionally, while polling, it verifies that the job never
exceeds MaxConcurrent running tasks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/states.go#L91)  

```go
func JobComplete(ctx context.Context, client client.ServiceAPIClient, service swarmtypes.Service) func(log poll.LogT) poll.Result
```

---

### NetworkPoll

NetworkPoll tweaks the pollSettings for `network`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L32)  

```go
func NetworkPoll(config *poll.Settings)
```

---

### NewSwarm

NewSwarm creates a swarm daemon for testing

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L43)  

```go
func NewSwarm(ctx context.Context, t *testing.T, testEnv *environment.Execution, ops ...daemon.Option) *daemon.Daemon
```

---

### NoTasks

NoTasks verifies that all tasks are gone

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/states.go#L36)  

```go
func NoTasks(ctx context.Context, client client.ServiceAPIClient) func(log poll.LogT) poll.Result
```

---

### NoTasksForService

NoTasksForService verifies that there are no more tasks for the given service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/states.go#L14)  

```go
func NoTasksForService(ctx context.Context, client client.ServiceAPIClient, serviceID string) func(log poll.LogT) poll.Result
```

---

### RunningTasksCount

RunningTasksCount verifies there are `instances` tasks running for `serviceID`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/states.go#L51)  

```go
func RunningTasksCount(ctx context.Context, client client.ServiceAPIClient, serviceID string, instances uint64) func(log poll.LogT) poll.Result
```

---

### ServicePoll

ServicePoll tweaks the pollSettings for `service`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L22)  

```go
func ServicePoll(config *poll.Settings)
```

---

### ServiceWithImage

ServiceWithImage sets the image to use for the service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L102)  

```go
func ServiceWithImage(image string) func(*swarmtypes.ServiceSpec)
```

---

### ServiceWithInit

ServiceWithInit sets whether the service should use init or not

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L94)  

```go
func ServiceWithInit(b *bool) func(*swarmtypes.ServiceSpec)
```

---

### ServiceWithMode

ServiceWithMode sets the mode of the service to the provided mode.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L87)  

```go
func ServiceWithMode(mode swarmtypes.ServiceMode) func(*swarmtypes.ServiceSpec)
```

---

## Types

### ServiceSpecOpt

ServiceSpecOpt is used with `CreateService` to pass in service spec modifiers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L57)  

```go
type ServiceSpecOpt func(*swarmtypes.ServiceSpec)
```

#### Functions

##### ServiceWithCapabilities

ServiceWithCapabilities sets the Capabilities option of the service's ContainerSpec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L183)  

```go
func ServiceWithCapabilities(add []string, drop []string) ServiceSpecOpt
```

##### ServiceWithCommand

ServiceWithCommand sets the command to use for the service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L110)  

```go
func ServiceWithCommand(cmd []string) ServiceSpecOpt
```

##### ServiceWithConfig

ServiceWithConfig adds the config reference to the service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L118)  

```go
func ServiceWithConfig(configRef *swarmtypes.ConfigReference) ServiceSpecOpt
```

##### ServiceWithEndpoint

ServiceWithEndpoint sets the Endpoint of the service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L168)  

```go
func ServiceWithEndpoint(endpoint *swarmtypes.EndpointSpec) ServiceSpecOpt
```

##### ServiceWithMaxReplicas

ServiceWithMaxReplicas sets the max replicas for the service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L145)  

```go
func ServiceWithMaxReplicas(n uint64) ServiceSpecOpt
```

##### ServiceWithName

ServiceWithName sets the name of the service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L153)  

```go
func ServiceWithName(name string) ServiceSpecOpt
```

##### ServiceWithNetwork

ServiceWithNetwork sets the network of the service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L160)  

```go
func ServiceWithNetwork(network string) ServiceSpecOpt
```

##### ServiceWithPidsLimit

ServiceWithPidsLimit sets the PidsLimit option of the service's Resources.Limits.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L192)  

```go
func ServiceWithPidsLimit(limit int64) ServiceSpecOpt
```

##### ServiceWithReplicas

ServiceWithReplicas sets the replicas for the service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L134)  

```go
func ServiceWithReplicas(n uint64) ServiceSpecOpt
```

##### ServiceWithSecret

ServiceWithSecret adds the secret reference to the service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L126)  

```go
func ServiceWithSecret(secretRef *swarmtypes.SecretReference) ServiceSpecOpt
```

##### ServiceWithSysctls

ServiceWithSysctls sets the Sysctls option of the service's ContainerSpec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/swarm/service.go#L175)  

```go
func ServiceWithSysctls(sysctls map[string]string) ServiceSpecOpt
```

---


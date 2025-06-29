# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/swarm

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:31 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new build router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/swarm/cluster.go#L12)  

```go
func NewRouter(b Backend) router.Router
```

---

## Types

### Backend

Backend abstracts a swarm manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/swarm/backend.go#L12)  

```go
type Backend interface {
	Init(req swarm.InitRequest) (string, error)
	Join(req swarm.JoinRequest) error
	Leave(ctx context.Context, force bool) error
	Inspect() (swarm.Swarm, error)
	Update(uint64, swarm.Spec, swarm.UpdateFlags) error
	GetUnlockKey() (string, error)
	UnlockSwarm(req swarm.UnlockRequest) error
	GetServices(swarm.ServiceListOptions) ([]swarm.Service, error)
	GetService(idOrName string, insertDefaults bool) (swarm.Service, error)
	CreateService(swarm.ServiceSpec, string, bool) (*swarm.ServiceCreateResponse, error)
	UpdateService(string, uint64, swarm.ServiceSpec, swarm.ServiceUpdateOptions, bool) (*swarm.ServiceUpdateResponse, error)
	RemoveService(string) error
	ServiceLogs(context.Context, *backend.LogSelector, *container.LogsOptions) (<-chan *backend.LogMessage, error)
	GetNodes(swarm.NodeListOptions) ([]swarm.Node, error)
	GetNode(string) (swarm.Node, error)
	UpdateNode(string, uint64, swarm.NodeSpec) error
	RemoveNode(string, bool) error
	GetTasks(swarm.TaskListOptions) ([]swarm.Task, error)
	GetTask(string) (swarm.Task, error)
	GetSecrets(opts swarm.SecretListOptions) ([]swarm.Secret, error)
	CreateSecret(s swarm.SecretSpec) (string, error)
	RemoveSecret(idOrName string) error
	GetSecret(id string) (swarm.Secret, error)
	UpdateSecret(idOrName string, version uint64, spec swarm.SecretSpec) error
	GetConfigs(opts swarm.ConfigListOptions) ([]swarm.Config, error)
	CreateConfig(s swarm.ConfigSpec) (string, error)
	RemoveConfig(id string) error
	GetConfig(id string) (swarm.Config, error)
	UpdateConfig(idOrName string, version uint64, spec swarm.ConfigSpec) error
}
```

---


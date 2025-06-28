# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/cluster/executor

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:50 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Backend

Backend defines the executor component for a swarm agent.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/executor/backend.go#L32)  

```go
type Backend interface {
	CreateManagedNetwork(clustertypes.NetworkCreateRequest) error
	DeleteManagedNetwork(networkID string) error
	FindNetwork(idName string) (*libnetwork.Network, error)
	SetupIngress(clustertypes.NetworkCreateRequest, string) (<-chan struct{}, error)
	ReleaseIngress() (<-chan struct{}, error)
	CreateManagedContainer(ctx context.Context, config backend.ContainerCreateConfig) (container.CreateResponse, error)
	ContainerStart(ctx context.Context, name string, checkpoint string, checkpointDir string) error
	ContainerStop(ctx context.Context, name string, config container.StopOptions) error
	ContainerLogs(ctx context.Context, name string, config *container.LogsOptions) (msgs <-chan *backend.LogMessage, tty bool, err error)
	ConnectContainerToNetwork(ctx context.Context, containerName, networkName string, endpointConfig *network.EndpointSettings) error
	ActivateContainerServiceBinding(containerName string) error
	DeactivateContainerServiceBinding(containerName string) error
	UpdateContainerServiceConfig(containerName string, serviceConfig *clustertypes.ServiceConfig) error
	ContainerInspect(ctx context.Context, name string, options backend.ContainerInspectOptions) (*container.InspectResponse, error)
	ContainerWait(ctx context.Context, name string, condition container.WaitCondition) (<-chan container.StateStatus, error)
	ContainerRm(name string, config *backend.ContainerRmConfig) error
	ContainerKill(name string, sig string) error
	SetContainerDependencyStore(name string, store exec.DependencyGetter) error
	SetContainerSecretReferences(name string, refs []*swarm.SecretReference) error
	SetContainerConfigReferences(name string, refs []*swarm.ConfigReference) error
	SystemInfo(context.Context) (*system.Info, error)
	Containers(ctx context.Context, config *container.ListOptions) ([]*container.Summary, error)
	SetNetworkBootstrapKeys([]*networktypes.EncryptionKey) error
	DaemonJoinsCluster(provider cluster.Provider)
	DaemonLeavesCluster()
	IsSwarmCompatible() error
	SubscribeToEvents(since, until time.Time, filter filters.Args) ([]events.Message, chan interface{})
	UnsubscribeFromEvents(listener chan interface{})
	UpdateAttachment(string, string, string, *network.NetworkingConfig) error
	WaitForDetachment(context.Context, string, string, string, string) error
	PluginManager() *plugin.Manager
	PluginGetter() *plugin.Store
	GetAttachmentStore() *networkSettings.AttachmentStore
	HasExperimental() bool
}
```

---

### ImageBackend

ImageBackend is used by an executor to perform image operations

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/executor/backend.go#L75)  

```go
type ImageBackend interface {
	PullImage(ctx context.Context, ref reference.Named, platform *ocispec.Platform, metaHeaders map[string][]string, authConfig *registry.AuthConfig, outStream io.Writer) error
	GetRepositories(context.Context, reference.Named, *registry.AuthConfig) ([]distribution.Repository, error)
	GetImage(ctx context.Context, refOrID string, options backend.GetImageOpts) (*image.Image, error)
}
```

---

### VolumeBackend

VolumeBackend is used by an executor to perform volume operations

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/executor/backend.go#L70)  

```go
type VolumeBackend interface {
	Create(ctx context.Context, name, driverName string, opts ...volumeopts.CreateOption) (*volume.Volume, error)
}
```

---


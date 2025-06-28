# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/cluster

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:42 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Cluster

Cluster provides capabilities to participate in a cluster as a worker or a
manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L112)  

```go
type Cluster struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates a new Cluster instance using provided config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L138)  

```go
func New(config Config) (*Cluster, error)
```

#### Methods

##### Cluster.AttachNetwork

AttachNetwork generates an attachment request towards the manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/networks.go#L181)  
**Added in:** v1.13.0

```go
func (c *Cluster) AttachNetwork(target string, containerID string, addresses []string) (*network.NetworkingConfig, error)
```

##### Cluster.Cleanup

Cleanup stops active swarm node. This is run before daemon shutdown.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L369)  

```go
func (c *Cluster) Cleanup()
```

##### Cluster.CreateConfig

CreateConfig creates a new config in a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/configs.go#L65)  

```go
func (c *Cluster) CreateConfig(s types.ConfigSpec) (string, error)
```

##### Cluster.CreateNetwork

CreateNetwork creates a new cluster managed network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/networks.go#L270)  

```go
func (c *Cluster) CreateNetwork(s network.CreateRequest) (string, error)
```

##### Cluster.CreateSecret

CreateSecret creates a new secret in a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/secrets.go#L66)  
**Added in:** v1.13.0

```go
func (c *Cluster) CreateSecret(s types.SecretSpec) (string, error)
```

##### Cluster.CreateService

CreateService creates a new service in a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/services.go#L184)  

```go
func (c *Cluster) CreateService(s swarm.ServiceSpec, encodedAuth string, queryRegistry bool) (*swarm.ServiceCreateResponse, error)
```

##### Cluster.CreateVolume

CreateVolume creates a new cluster volume in the swarm cluster.

Returns the volume ID if creation is successful, or an error if not.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/volumes.go#L61)  

```go
func (c *Cluster) CreateVolume(v volumetypes.CreateOptions) (*volumetypes.Volume, error)
```

##### Cluster.DetachNetwork

DetachNetwork unblocks the waiters waiting on WaitForDetachment so
that a request to detach can be generated towards the manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/networks.go#L253)  
**Added in:** v1.13.0

```go
func (c *Cluster) DetachNetwork(target string, containerID string) error
```

##### Cluster.GetAdvertiseAddress

GetAdvertiseAddress returns the remotely reachable address of this node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L286)  

```go
func (c *Cluster) GetAdvertiseAddress() string
```

##### Cluster.GetConfig

GetConfig returns a config from a managed swarm cluster

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/configs.go#L13)  

```go
func (c *Cluster) GetConfig(input string) (types.Config, error)
```

##### Cluster.GetConfigs

GetConfigs returns all configs of a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/configs.go#L30)  

```go
func (c *Cluster) GetConfigs(options types.ConfigListOptions) ([]types.Config, error)
```

##### Cluster.GetDataPathAddress

GetDataPathAddress returns the address to be used for the data path traffic, if specified.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L297)  

```go
func (c *Cluster) GetDataPathAddress() string
```

##### Cluster.GetListenAddress

GetListenAddress returns the listen address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L276)  
**Added in:** v1.12.2

```go
func (c *Cluster) GetListenAddress() string
```

##### Cluster.GetLocalAddress

GetLocalAddress returns the local address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L269)  

```go
func (c *Cluster) GetLocalAddress() string
```

##### Cluster.GetNetwork

GetNetwork returns a cluster network by an ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/networks.go#L88)  

```go
func (c *Cluster) GetNetwork(input string) (network.Inspect, error)
```

##### Cluster.GetNetworks

GetNetworks returns all current cluster managed networks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/networks.go#L19)  

```go
func (c *Cluster) GetNetworks(filter filters.Args) ([]network.Inspect, error)
```

##### Cluster.GetNetworksByName

GetNetworksByName returns cluster managed networks by name.
It is ok to have multiple networks here. #18864

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/networks.go#L106)  
**Added in:** v1.13.0

```go
func (c *Cluster) GetNetworksByName(name string) ([]network.Inspect, error)
```

##### Cluster.GetNode

GetNode returns a node based on an ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/nodes.go#L50)  

```go
func (c *Cluster) GetNode(input string) (types.Node, error)
```

##### Cluster.GetNodes

GetNodes returns a list of all nodes known to a cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/nodes.go#L14)  

```go
func (c *Cluster) GetNodes(options types.NodeListOptions) ([]types.Node, error)
```

##### Cluster.GetRemoteAddressList

GetRemoteAddressList returns the advertise address for each of the remote managers if
available.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L308)  

```go
func (c *Cluster) GetRemoteAddressList() []string
```

##### Cluster.GetSecret

GetSecret returns a secret from a managed swarm cluster

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/secrets.go#L13)  
**Added in:** v1.13.0

```go
func (c *Cluster) GetSecret(input string) (types.Secret, error)
```

##### Cluster.GetSecrets

GetSecrets returns all secrets of a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/secrets.go#L30)  
**Added in:** v1.13.0

```go
func (c *Cluster) GetSecrets(options types.SecretListOptions) ([]types.Secret, error)
```

##### Cluster.GetService

GetService returns a service based on an ID or name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/services.go#L164)  

```go
func (c *Cluster) GetService(input string, insertDefaults bool) (swarm.Service, error)
```

##### Cluster.GetServices

GetServices returns all services of a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/services.go#L31)  

```go
func (c *Cluster) GetServices(options swarm.ServiceListOptions) ([]swarm.Service, error)
```

##### Cluster.GetTask

GetTask returns a task by an ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/tasks.go#L77)  

```go
func (c *Cluster) GetTask(input string) (types.Task, error)
```

##### Cluster.GetTasks

GetTasks returns a list of tasks matching the filter options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/tasks.go#L14)  

```go
func (c *Cluster) GetTasks(options types.TaskListOptions) ([]types.Task, error)
```

##### Cluster.GetUnlockKey

GetUnlockKey returns the unlock key for the swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/swarm.go#L285)  
**Added in:** v1.13.0

```go
func (c *Cluster) GetUnlockKey() (string, error)
```

##### Cluster.GetVolume

GetVolume returns a volume from the swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/volumes.go#L16)  

```go
func (c *Cluster) GetVolume(nameOrID string) (volumetypes.Volume, error)
```

##### Cluster.GetVolumes

GetVolumes returns all of the volumes matching the given options from a swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/volumes.go#L33)  

```go
func (c *Cluster) GetVolumes(options volumetypes.ListOptions) ([]*volumetypes.Volume, error)
```

##### Cluster.GetWatchStream

GetWatchStream returns the channel to pass changes from store watch API

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L315)  

```go
func (c *Cluster) GetWatchStream() chan *swarmapi.WatchMessage
```

##### Cluster.Info

Info returns information about the current cluster state.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/swarm.go#L432)  

```go
func (c *Cluster) Info(ctx context.Context) types.Info
```

##### Cluster.Init

Init initializes new cluster from user provided request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/swarm.go#L27)  

```go
func (c *Cluster) Init(req types.InitRequest) (string, error)
```

##### Cluster.Inspect

Inspect retrieves the configuration properties of a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/swarm.go#L219)  

```go
func (c *Cluster) Inspect() (types.Swarm, error)
```

##### Cluster.IsAgent

IsAgent returns true if Cluster is participating as a worker/agent.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L262)  

```go
func (c *Cluster) IsAgent() bool
```

##### Cluster.IsManager

IsManager returns true if Cluster is participating as a manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L255)  

```go
func (c *Cluster) IsManager() bool
```

##### Cluster.Join

Join makes current Cluster part of an existing swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/swarm.go#L148)  

```go
func (c *Cluster) Join(req types.JoinRequest) error
```

##### Cluster.Leave

Leave shuts down Cluster and removes current state.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/swarm.go#L360)  

```go
func (c *Cluster) Leave(ctx context.Context, force bool) error
```

##### Cluster.ListenClusterEvents

ListenClusterEvents returns a channel that receives messages on cluster
participation changes.
todo: make cancelable and accessible to multiple callers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L341)  

```go
func (c *Cluster) ListenClusterEvents() <-chan lncluster.ConfigEventType
```

##### Cluster.RemoveConfig

RemoveConfig removes a config from a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/configs.go#L84)  

```go
func (c *Cluster) RemoveConfig(input string) error
```

##### Cluster.RemoveNetwork

RemoveNetwork removes a cluster network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/networks.go#L293)  

```go
func (c *Cluster) RemoveNetwork(input string) error
```

##### Cluster.RemoveNode

RemoveNode removes a node from a cluster

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/nodes.go#L99)  

```go
func (c *Cluster) RemoveNode(input string, force bool) error
```

##### Cluster.RemoveSecret

RemoveSecret removes a secret from a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/secrets.go#L85)  
**Added in:** v1.13.0

```go
func (c *Cluster) RemoveSecret(input string) error
```

##### Cluster.RemoveService

RemoveService removes a service from a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/services.go#L415)  

```go
func (c *Cluster) RemoveService(input string) error
```

##### Cluster.RemoveVolume

RemoveVolume removes a volume from the swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/volumes.go#L90)  

```go
func (c *Cluster) RemoveVolume(nameOrID string, force bool) error
```

##### Cluster.SendClusterEvent

SendClusterEvent allows to send cluster events on the configEvent channel
TODO This method should not be exposed.
Currently it is used to notify the network controller that the keys are
available

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L455)  

```go
func (c *Cluster) SendClusterEvent(event lncluster.ConfigEventType)
```

##### Cluster.ServiceLogs

ServiceLogs collects service logs and writes them back to `config.OutStream`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/services.go#L428)  
**Added in:** v1.13.0

```go
func (c *Cluster) ServiceLogs(ctx context.Context, selector *backend.LogSelector, config *container.LogsOptions) (<-chan *backend.LogMessage, error)
```

##### Cluster.Start

Start the Cluster instance
TODO The split between New and Start can be join again when the SendClusterEvent
method is no longer required

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L168)  

```go
func (c *Cluster) Start() error
```

##### Cluster.Status

Status returns a textual representation of the node's swarm status and role (manager/worker)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/swarm.go#L497)  

```go
func (c *Cluster) Status() string
```

##### Cluster.UnlockSwarm

UnlockSwarm provides a key to decrypt data that is encrypted at rest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/swarm.go#L307)  
**Added in:** v1.13.0

```go
func (c *Cluster) UnlockSwarm(req types.UnlockRequest) error
```

##### Cluster.Update

Update updates configuration of a managed swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/swarm.go#L243)  

```go
func (c *Cluster) Update(version uint64, spec types.Spec, flags types.UpdateFlags) error
```

##### Cluster.UpdateAttachment

UpdateAttachment signals the attachment config to the attachment
waiter who is trying to start or attach the container to the
network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/networks.go#L121)  
**Added in:** v1.13.0

```go
func (c *Cluster) UpdateAttachment(target, containerID string, config *network.NetworkingConfig) error
```

##### Cluster.UpdateConfig

UpdateConfig updates a config in a managed swarm cluster.
Note: this is not exposed to the CLI but is available from the API only

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/configs.go#L102)  

```go
func (c *Cluster) UpdateConfig(input string, version uint64, spec types.ConfigSpec) error
```

##### Cluster.UpdateNode

UpdateNode updates existing nodes properties.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/nodes.go#L68)  

```go
func (c *Cluster) UpdateNode(input string, version uint64, spec types.NodeSpec) error
```

##### Cluster.UpdateSecret

UpdateSecret updates a secret in a managed swarm cluster.
Note: this is not exposed to the CLI but is available from the API only

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/secrets.go#L103)  
**Added in:** v1.13.0

```go
func (c *Cluster) UpdateSecret(input string, version uint64, spec types.SecretSpec) error
```

##### Cluster.UpdateService

UpdateService updates existing service to match new properties.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/services.go#L285)  

```go
func (c *Cluster) UpdateService(serviceIDOrName string, version uint64, spec swarm.ServiceSpec, flags swarm.ServiceUpdateOptions, queryRegistry bool) (*swarm.ServiceUpdateResponse, error)
```

##### Cluster.UpdateVolume

UpdateVolume updates a volume in the swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/volumes.go#L109)  

```go
func (c *Cluster) UpdateVolume(nameOrID string, version uint64, volume volumetypes.UpdateOptions) error
```

##### Cluster.WaitForDetachment

WaitForDetachment waits for the container to stop or detach from
the network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/networks.go#L143)  
**Added in:** v1.13.0

```go
func (c *Cluster) WaitForDetachment(ctx context.Context, networkName, networkID, taskID, containerID string) error
```

---

### Config

Config provides values for Cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L86)  

```go
type Config struct {
	Root                   string
	Name                   string
	Backend                executorpkg.Backend
	ImageBackend           executorpkg.ImageBackend
	PluginBackend          plugin.Backend
	VolumeBackend          executorpkg.VolumeBackend
	NetworkSubnetsProvider NetworkSubnetsProvider

	// DefaultAdvertiseAddr is the default host/IP or network interface to use
	// if no AdvertiseAddr value is specified.
	DefaultAdvertiseAddr string

	// path to store runtime state, such as the swarm control socket
	RuntimeRoot string

	// RaftHeartbeatTick is the number of ticks for heartbeat of quorum members
	RaftHeartbeatTick uint32

	// RaftElectionTick is the number of ticks to elapse before followers propose a new round of leader election
	// This value should be 10x that of RaftHeartbeatTick
	RaftElectionTick uint32
}
```

---

### NetworkSubnetsProvider

NetworkSubnetsProvider exposes functions for retrieving the subnets
of networks managed by Docker, so they can be filtered.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/cluster.go#L81)  

```go
type NetworkSubnetsProvider interface {
	Subnets() ([]net.IPNet, []net.IPNet)
}
```

---


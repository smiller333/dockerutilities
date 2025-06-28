# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/swarm

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:18 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime.go#L9)

```go
const (
	// RuntimeContainer is the container based runtime
	RuntimeContainer RuntimeType = "container"
	// RuntimePlugin is the plugin based runtime
	RuntimePlugin RuntimeType = "plugin"
	// RuntimeNetworkAttachment is the network attachment runtime
	RuntimeNetworkAttachment RuntimeType = "attachment"

	// RuntimeURLContainer is the proto url for the container type
	RuntimeURLContainer RuntimeURL = "types.docker.com/RuntimeContainer"
	// RuntimeURLPlugin is the proto url for the plugin type
	RuntimeURLPlugin RuntimeURL = "types.docker.com/RuntimePlugin"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L116)

```go
const (
	// UpdateFailureActionPause PAUSE
	UpdateFailureActionPause = "pause"
	// UpdateFailureActionContinue CONTINUE
	UpdateFailureActionContinue = "continue"
	// UpdateFailureActionRollback ROLLBACK
	UpdateFailureActionRollback = "rollback"

	// UpdateOrderStopFirst STOP_FIRST
	UpdateOrderStopFirst = "stop-first"
	// UpdateOrderStartFirst START_FIRST
	UpdateOrderStartFirst = "start-first"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L224)

```go
const (
	RegistryAuthFromSpec         = "spec"
	RegistryAuthFromPreviousSpec = "previous-spec"
)
```

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Annotations

Annotations represents how to describe an object.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/common.go#L26)  

```go
type Annotations struct {
	Name   string            `json:",omitempty"`
	Labels map[string]string `json:"Labels"`
}
```

---

### AppArmorMode

AppArmorMode is type used for the enumeration of possible AppArmor modes in
AppArmorOpts

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/container.go#L57)  

```go
type AppArmorMode string
```

---

### AppArmorOpts

AppArmorOpts defines the options for configuring AppArmor on a swarm-managed
container.  Currently, custom AppArmor profiles are not supported.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/container.go#L66)  

```go
type AppArmorOpts struct {
	Mode AppArmorMode `json:",omitempty"`
}
```

---

### CAConfig

CAConfig represents CA configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L110)  

```go
type CAConfig struct {
	// NodeCertExpiry is the duration certificates should be issued for
	NodeCertExpiry time.Duration `json:",omitempty"`

	// ExternalCAs is a list of CAs to which a manager node will make
	// certificate signing requests for node certificates.
	ExternalCAs []*ExternalCA `json:",omitempty"`

	// SigningCACert and SigningCAKey specify the desired signing root CA and
	// root CA key for the swarm.  When inspecting the cluster, the key will
	// be redacted.
	SigningCACert string `json:",omitempty"`
	SigningCAKey  string `json:",omitempty"`

	// If this value changes, and there is no specified signing cert and key,
	// then the swarm is forced to generate a new root certificate and key.
	ForceRotate uint64 `json:",omitempty"`
}
```

---

### ClusterInfo

ClusterInfo represents info about the cluster for outputting in "info"
it contains the same information as "Swarm", but without the JoinTokens

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L9)  

```go
type ClusterInfo struct {
	ID string
	Meta
	Spec                   Spec
	TLSInfo                TLSInfo
	RootRotationInProgress bool
	DefaultAddrPool        []string
	SubnetSize             uint32
	DataPathPort           uint32
}
```

---

### Config

Config represents a config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/config.go#L10)  

```go
type Config struct {
	ID string
	Meta
	Spec ConfigSpec
}
```

---

### ConfigCreateResponse

ConfigCreateResponse contains the information returned to a client
on the creation of a new config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/config.go#L54)  

```go
type ConfigCreateResponse struct {
	// ID is the id of the created config.
	ID string
}
```

---

### ConfigListOptions

ConfigListOptions holds parameters to list configs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/config.go#L60)  

```go
type ConfigListOptions struct {
	Filters filters.Args
}
```

---

### ConfigReference

ConfigReference is a reference to a config in swarm

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/config.go#L45)  

```go
type ConfigReference struct {
	File       *ConfigReferenceFileTarget    `json:",omitempty"`
	Runtime    *ConfigReferenceRuntimeTarget `json:",omitempty"`
	ConfigID   string
	ConfigName string
}
```

---

### ConfigReferenceFileTarget

ConfigReferenceFileTarget is a file target in a config reference

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/config.go#L33)  

```go
type ConfigReferenceFileTarget struct {
	Name string
	UID  string
	GID  string
	Mode os.FileMode
}
```

---

### ConfigReferenceRuntimeTarget

ConfigReferenceRuntimeTarget is a target for a config specifying that it
isn't mounted into the container but instead has some other purpose.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/config.go#L42)  

```go
type ConfigReferenceRuntimeTarget struct{}
```

---

### ConfigSpec

ConfigSpec represents a config specification from a config in swarm

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/config.go#L17)  

```go
type ConfigSpec struct {
	Annotations

	// Data is the data to store as a config.
	//
	// The maximum allowed size is 1000KB, as defined in [MaxConfigSize].
	//
	// [MaxConfigSize]: https://pkg.go.dev/github.com/moby/swarmkit/v2@v2.0.0-20250103191802-8c1959736554/manager/controlapi#MaxConfigSize
	Data []byte `json:",omitempty"`

	// Templating controls whether and how to evaluate the config payload as
	// a template. If it is not set, no templating is used.
	Templating *Driver `json:",omitempty"`
}
```

---

### ContainerSpec

ContainerSpec represents the spec of a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/container.go#L87)  

```go
type ContainerSpec struct {
	Image           string                  `json:",omitempty"`
	Labels          map[string]string       `json:",omitempty"`
	Command         []string                `json:",omitempty"`
	Args            []string                `json:",omitempty"`
	Hostname        string                  `json:",omitempty"`
	Env             []string                `json:",omitempty"`
	Dir             string                  `json:",omitempty"`
	User            string                  `json:",omitempty"`
	Groups          []string                `json:",omitempty"`
	Privileges      *Privileges             `json:",omitempty"`
	Init            *bool                   `json:",omitempty"`
	StopSignal      string                  `json:",omitempty"`
	TTY             bool                    `json:",omitempty"`
	OpenStdin       bool                    `json:",omitempty"`
	ReadOnly        bool                    `json:",omitempty"`
	Mounts          []mount.Mount           `json:",omitempty"`
	StopGracePeriod *time.Duration          `json:",omitempty"`
	Healthcheck     *container.HealthConfig `json:",omitempty"`
	// The format of extra hosts on swarmkit is specified in:
	// http://man7.org/linux/man-pages/man5/hosts.5.html
	//    IP_address canonical_hostname [aliases...]
	Hosts          []string            `json:",omitempty"`
	DNSConfig      *DNSConfig          `json:",omitempty"`
	Secrets        []*SecretReference  `json:",omitempty"`
	Configs        []*ConfigReference  `json:",omitempty"`
	Isolation      container.Isolation `json:",omitempty"`
	Sysctls        map[string]string   `json:",omitempty"`
	CapabilityAdd  []string            `json:",omitempty"`
	CapabilityDrop []string            `json:",omitempty"`
	Ulimits        []*container.Ulimit `json:",omitempty"`
	OomScoreAdj    int64               `json:",omitempty"`
}
```

---

### ContainerStatus

ContainerStatus represents the status of a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L202)  

```go
type ContainerStatus struct {
	ContainerID string
	PID         int
	ExitCode    int
}
```

---

### CredentialSpec

CredentialSpec for managed service account (Windows only)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/container.go#L71)  

```go
type CredentialSpec struct {
	Config   string
	File     string
	Registry string
}
```

---

### DNSConfig

DNSConfig specifies DNS related configurations in resolver configuration file (resolv.conf)
Detailed documentation is available in:
http://man7.org/linux/man-pages/man5/resolv.conf.5.html
`nameserver`, `search`, `options` have been supported.
TODO: `domain` is not supported yet.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/container.go#L15)  

```go
type DNSConfig struct {
	// Nameservers specifies the IP addresses of the name servers
	Nameservers []string `json:",omitempty"`
	// Search specifies the search list for host-name lookup
	Search []string `json:",omitempty"`
	// Options allows certain internal resolver variables to be modified
	Options []string `json:",omitempty"`
}
```

---

### DiscreteGenericResource

DiscreteGenericResource represents a "user defined" resource which is defined
as an integer
"Kind" is used to describe the Kind of a resource (e.g: "GPU", "FPGA", "SSD", ...)
Value is used to count the resource (SSD=5, HDD=3, ...)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L135)  

```go
type DiscreteGenericResource struct {
	Kind  string `json:",omitempty"`
	Value int64  `json:",omitempty"`
}
```

---

### DispatcherConfig

DispatcherConfig represents dispatcher configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L103)  

```go
type DispatcherConfig struct {
	// HeartbeatPeriod defines how often agent should send heartbeats to
	// dispatcher.
	HeartbeatPeriod time.Duration `json:",omitempty"`
}
```

---

### Driver

Driver represents a driver (network, logging, secrets backend).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/common.go#L32)  

```go
type Driver struct {
	Name    string            `json:",omitempty"`
	Options map[string]string `json:",omitempty"`
}
```

---

### EncryptionConfig

EncryptionConfig controls at-rest encryption of data and keys.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L65)  

```go
type EncryptionConfig struct {
	// AutoLockManagers specifies whether or not managers TLS keys and raft data
	// should be encrypted at rest in such a way that they must be unlocked
	// before the manager node starts up again.
	AutoLockManagers bool
}
```

---

### Endpoint

Endpoint represents an endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L8)  

```go
type Endpoint struct {
	Spec       EndpointSpec        `json:",omitempty"`
	Ports      []PortConfig        `json:",omitempty"`
	VirtualIPs []EndpointVirtualIP `json:",omitempty"`
}
```

---

### EndpointSpec

EndpointSpec represents the spec of an endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L15)  

```go
type EndpointSpec struct {
	Mode  ResolutionMode `json:",omitempty"`
	Ports []PortConfig   `json:",omitempty"`
}
```

---

### EndpointVirtualIP

EndpointVirtualIP represents the virtual ip of a port.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L70)  

```go
type EndpointVirtualIP struct {
	NetworkID string `json:",omitempty"`
	Addr      string `json:",omitempty"`
}
```

---

### EngineDescription

EngineDescription represents the description of an engine.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L68)  

```go
type EngineDescription struct {
	EngineVersion string              `json:",omitempty"`
	Labels        map[string]string   `json:",omitempty"`
	Plugins       []PluginDescription `json:",omitempty"`
}
```

---

### ExternalCA

ExternalCA defines external CA to be used by the cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L136)  

```go
type ExternalCA struct {
	// Protocol is the protocol used by this external CA.
	Protocol ExternalCAProtocol

	// URL is the URL where the external CA can be reached.
	URL string

	// Options is a set of additional key/value pairs whose interpretation
	// depends on the specified CA type.
	Options map[string]string `json:",omitempty"`

	// CACert specifies which root CA is used by this external CA.  This certificate must
	// be in PEM format.
	CACert string
}
```

---

### ExternalCAProtocol

ExternalCAProtocol represents type of external CA.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L130)  

```go
type ExternalCAProtocol string
```

---

### GenericResource

GenericResource represents a "user defined" resource which can
be either an integer (e.g: SSD=3) or a string (e.g: SSD=sda1)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L117)  

```go
type GenericResource struct {
	NamedResourceSpec    *NamedGenericResource    `json:",omitempty"`
	DiscreteResourceSpec *DiscreteGenericResource `json:",omitempty"`
}
```

---

### GlobalJob

GlobalJob is the type of a Service which executes a Task on every Node
matching the Service's placement constraints. These tasks run to completion
and then exit.

This type is deliberately empty.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L114)  

```go
type GlobalJob struct{}
```

---

### GlobalService

GlobalService is a kind of ServiceMode.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L88)  

```go
type GlobalService struct{}
```

---

### IPAMConfig

IPAMConfig represents ipam configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L117)  

```go
type IPAMConfig struct {
	Subnet  string `json:",omitempty"`
	Range   string `json:",omitempty"`
	Gateway string `json:",omitempty"`
}
```

---

### IPAMOptions

IPAMOptions represents ipam options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L111)  

```go
type IPAMOptions struct {
	Driver  Driver       `json:",omitempty"`
	Configs []IPAMConfig `json:",omitempty"`
}
```

---

### Info

Info represents generic information about swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L199)  

```go
type Info struct {
	NodeID   string
	NodeAddr string

	LocalNodeState   LocalNodeState
	ControlAvailable bool
	Error            string

	RemoteManagers []Peer
	Nodes          int `json:",omitempty"`
	Managers       int `json:",omitempty"`

	Cluster *ClusterInfo `json:",omitempty"`

	Warnings []string `json:",omitempty"`
}
```

---

### InitRequest

InitRequest is the request used to init a swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L153)  

```go
type InitRequest struct {
	ListenAddr       string
	AdvertiseAddr    string
	DataPathAddr     string
	DataPathPort     uint32
	ForceNewCluster  bool
	Spec             Spec
	AutoLockManagers bool
	Availability     NodeAvailability
	DefaultAddrPool  []string
	SubnetSize       uint32
}
```

---

### JobStatus

JobStatus is the status of a job-type service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L191)  

```go
type JobStatus struct {
	// JobIteration is a value increased each time a Job is executed,
	// successfully or otherwise. "Executed", in this case, means the job as a
	// whole has been started, not that an individual Task has been launched. A
	// job is "Executed" when its ServiceSpec is updated. JobIteration can be
	// used to disambiguate Tasks belonging to different executions of a job.
	//
	// Though JobIteration will increase with each subsequent execution, it may
	// not necessarily increase by 1, and so JobIteration should not be used to
	// keep track of the number of times a job has been executed.
	JobIteration Version

	// LastExecution is the time that the job was last executed, as observed by
	// Swarm manager.
	LastExecution time.Time `json:",omitempty"`
}
```

---

### JoinRequest

JoinRequest is the request used to join a swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L167)  

```go
type JoinRequest struct {
	ListenAddr    string
	AdvertiseAddr string
	DataPathAddr  string
	RemoteAddrs   []string
	JoinToken     string // accept by secret
	Availability  NodeAvailability
}
```

---

### JoinTokens

JoinTokens contains the tokens workers and managers need to join the swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L27)  

```go
type JoinTokens struct {
	// Worker is the join token workers may use to join the swarm.
	Worker string
	// Manager is the join token managers may use to join the swarm.
	Manager string
}
```

---

### Limit

Limit describes limits on resources which can be requested by a task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L109)  

```go
type Limit struct {
	NanoCPUs    int64 `json:",omitempty"`
	MemoryBytes int64 `json:",omitempty"`
	Pids        int64 `json:",omitempty"`
}
```

---

### LocalNodeState

LocalNodeState represents the state of the local node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L183)  

```go
type LocalNodeState string
```

---

### ManagerStatus

ManagerStatus represents the status of a manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L115)  

```go
type ManagerStatus struct {
	Leader       bool         `json:",omitempty"`
	Reachability Reachability `json:",omitempty"`
	Addr         string       `json:",omitempty"`
}
```

---

### Meta

Meta is a base object inherited by most of the other once.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/common.go#L19)  

```go
type Meta struct {
	Version   Version   `json:",omitempty"`
	CreatedAt time.Time `json:",omitempty"`
	UpdatedAt time.Time `json:",omitempty"`
}
```

---

### NamedGenericResource

NamedGenericResource represents a "user defined" resource which is defined
as a string.
"Kind" is used to describe the Kind of a resource (e.g: "GPU", "FPGA", "SSD", ...)
Value is used to identify the resource (GPU="UUID-1", FPGA="/dev/sdb5", ...)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L126)  

```go
type NamedGenericResource struct {
	Kind  string `json:",omitempty"`
	Value string `json:",omitempty"`
}
```

---

### Network

Network represents a network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L76)  

```go
type Network struct {
	ID string
	Meta
	Spec        NetworkSpec  `json:",omitempty"`
	DriverState Driver       `json:",omitempty"`
	IPAMOptions *IPAMOptions `json:",omitempty"`
}
```

---

### NetworkAttachment

NetworkAttachment represents a network attachment.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L105)  

```go
type NetworkAttachment struct {
	Network   Network  `json:",omitempty"`
	Addresses []string `json:",omitempty"`
}
```

---

### NetworkAttachmentConfig

NetworkAttachmentConfig represents the configuration of a network attachment.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L98)  

```go
type NetworkAttachmentConfig struct {
	Target     string            `json:",omitempty"`
	Aliases    []string          `json:",omitempty"`
	DriverOpts map[string]string `json:",omitempty"`
}
```

---

### NetworkAttachmentSpec

NetworkAttachmentSpec represents the runtime spec type for network
attachment tasks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime.go#L25)  

```go
type NetworkAttachmentSpec struct {
	ContainerID string
}
```

---

### NetworkSpec

NetworkSpec represents the spec of a network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L85)  

```go
type NetworkSpec struct {
	Annotations
	DriverConfiguration *Driver                  `json:",omitempty"`
	IPv6Enabled         bool                     `json:",omitempty"`
	Internal            bool                     `json:",omitempty"`
	Attachable          bool                     `json:",omitempty"`
	Ingress             bool                     `json:",omitempty"`
	IPAMOptions         *IPAMOptions             `json:",omitempty"`
	ConfigFrom          *network.ConfigReference `json:",omitempty"`
	Scope               string                   `json:",omitempty"`
}
```

---

### Node

Node represents a node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L6)  

```go
type Node struct {
	ID string
	Meta
	// Spec defines the desired state of the node as specified by the user.
	// The system will honor this and will *never* modify it.
	Spec NodeSpec `json:",omitempty"`
	// Description encapsulates the properties of the Node as reported by the
	// agent.
	Description NodeDescription `json:",omitempty"`
	// Status provides the current status of the node, as seen by the manager.
	Status NodeStatus `json:",omitempty"`
	// ManagerStatus provides the current status of the node's manager
	// component, if the node is a manager.
	ManagerStatus *ManagerStatus `json:",omitempty"`
}
```

---

### NodeAvailability

NodeAvailability represents the availability of a node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L40)  

```go
type NodeAvailability string
```

---

### NodeCSIInfo

NodeCSIInfo represents information about a CSI plugin available on the node

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L75)  

```go
type NodeCSIInfo struct {
	// PluginName is the name of the CSI plugin.
	PluginName string `json:",omitempty"`
	// NodeID is the ID of the node as reported by the CSI plugin. This is
	// different from the swarm node ID.
	NodeID string `json:",omitempty"`
	// MaxVolumesPerNode is the maximum number of volumes that may be published
	// to this node
	MaxVolumesPerNode int64 `json:",omitempty"`
	// AccessibleTopology indicates the location of this node in the CSI
	// plugin's topology
	AccessibleTopology *Topology `json:",omitempty"`
}
```

---

### NodeDescription

NodeDescription represents the description of a node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L52)  

```go
type NodeDescription struct {
	Hostname  string            `json:",omitempty"`
	Platform  Platform          `json:",omitempty"`
	Resources Resources         `json:",omitempty"`
	Engine    EngineDescription `json:",omitempty"`
	TLSInfo   TLSInfo           `json:",omitempty"`
	CSIInfo   []NodeCSIInfo     `json:",omitempty"`
}
```

---

### NodeListOptions

NodeListOptions holds parameters to list nodes with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L144)  

```go
type NodeListOptions struct {
	Filters filters.Args
}
```

---

### NodeRemoveOptions

NodeRemoveOptions holds parameters to remove nodes with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L149)  

```go
type NodeRemoveOptions struct {
	Force bool
}
```

---

### NodeRole

NodeRole represents the role of a node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L30)  

```go
type NodeRole string
```

---

### NodeSpec

NodeSpec represents the spec of a node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L23)  

```go
type NodeSpec struct {
	Annotations
	Role         NodeRole         `json:",omitempty"`
	Availability NodeAvailability `json:",omitempty"`
}
```

---

### NodeState

NodeState represents the state of a node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L122)  

```go
type NodeState string
```

---

### NodeStatus

NodeStatus represents the status of a node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L96)  

```go
type NodeStatus struct {
	State   NodeState `json:",omitempty"`
	Message string    `json:",omitempty"`
	Addr    string    `json:",omitempty"`
}
```

---

### OrchestrationConfig

OrchestrationConfig represents orchestration configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L47)  

```go
type OrchestrationConfig struct {
	// TaskHistoryRetentionLimit is the number of historic tasks to keep per instance or
	// node. If negative, never remove completed or failed tasks.
	TaskHistoryRetentionLimit *int64 `json:",omitempty"`
}
```

---

### Peer

Peer represents a peer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L227)  

```go
type Peer struct {
	NodeID string
	Addr   string
}
```

---

### Placement

Placement represents orchestration parameters.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L147)  

```go
type Placement struct {
	Constraints []string              `json:",omitempty"`
	Preferences []PlacementPreference `json:",omitempty"`
	MaxReplicas uint64                `json:",omitempty"`

	// Platforms stores all the platforms that the image can run on.
	// This field is used in the platform filter for scheduling. If empty,
	// then the platform filter is off, meaning there are no scheduling restrictions.
	Platforms []Platform `json:",omitempty"`
}
```

---

### PlacementPreference

PlacementPreference provides a way to make the scheduler aware of factors
such as topology.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L160)  

```go
type PlacementPreference struct {
	Spread *SpreadOver
}
```

---

### Platform

Platform represents the platform (Arch/OS).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L62)  

```go
type Platform struct {
	Architecture string `json:",omitempty"`
	OS           string `json:",omitempty"`
}
```

---

### PluginDescription

PluginDescription represents the description of an engine plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L90)  

```go
type PluginDescription struct {
	Type string `json:",omitempty"`
	Name string `json:",omitempty"`
}
```

---

### PortConfig

PortConfig represents the config of a port.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L31)  

```go
type PortConfig struct {
	Name     string             `json:",omitempty"`
	Protocol PortConfigProtocol `json:",omitempty"`
	// TargetPort is the port inside the container
	TargetPort uint32 `json:",omitempty"`
	// PublishedPort is the port on the swarm hosts
	PublishedPort uint32 `json:",omitempty"`
	// PublishMode is the mode in which port is published
	PublishMode PortConfigPublishMode `json:",omitempty"`
}
```

---

### PortConfigProtocol

PortConfigProtocol represents the protocol of a port.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L56)  

```go
type PortConfigProtocol string
```

---

### PortConfigPublishMode

PortConfigPublishMode represents the mode in which the port is to
be published.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L44)  

```go
type PortConfigPublishMode string
```

---

### PortStatus

PortStatus represents the port status of a task's host ports whose
service has published host ports

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L210)  

```go
type PortStatus struct {
	Ports []PortConfig `json:",omitempty"`
}
```

---

### Privileges

Privileges defines the security options for the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/container.go#L78)  

```go
type Privileges struct {
	CredentialSpec  *CredentialSpec
	SELinuxContext  *SELinuxContext
	Seccomp         *SeccompOpts  `json:",omitempty"`
	AppArmor        *AppArmorOpts `json:",omitempty"`
	NoNewPrivileges bool
}
```

---

### RaftConfig

RaftConfig represents raft configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L73)  

```go
type RaftConfig struct {
	// SnapshotInterval is the number of log entries between snapshots.
	SnapshotInterval uint64 `json:",omitempty"`

	// KeepOldSnapshots is the number of snapshots to keep beyond the
	// current snapshot.
	KeepOldSnapshots *uint64 `json:",omitempty"`

	// LogEntriesForSlowFollowers is the number of log entries to keep
	// around to sync up slow followers after a snapshot is created.
	LogEntriesForSlowFollowers uint64 `json:",omitempty"`

	// ElectionTick is the number of ticks that a follower will wait for a message
	// from the leader before becoming a candidate and starting an election.
	// ElectionTick must be greater than HeartbeatTick.
	//
	// A tick currently defaults to one second, so these translate directly to
	// seconds currently, but this is NOT guaranteed.
	ElectionTick int

	// HeartbeatTick is the number of ticks between heartbeats. Every
	// HeartbeatTick ticks, the leader will send a heartbeat to the
	// followers.
	//
	// A tick currently defaults to one second, so these translate directly to
	// seconds currently, but this is NOT guaranteed.
	HeartbeatTick int
}
```

---

### Reachability

Reachability represents the reachability of a node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L103)  

```go
type Reachability string
```

---

### ReplicatedJob

ReplicatedJob is the a type of Service which executes a defined Tasks
in parallel until the specified number of Tasks have succeeded.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L92)  

```go
type ReplicatedJob struct {
	// MaxConcurrent indicates the maximum number of Tasks that should be
	// executing simultaneously for this job at any given time. There may be
	// fewer Tasks that MaxConcurrent executing simultaneously; for example, if
	// there are fewer than MaxConcurrent tasks needed to reach
	// TotalCompletions.
	//
	// If this field is empty, it will default to a max concurrency of 1.
	MaxConcurrent *uint64 `json:",omitempty"`

	// TotalCompletions is the total number of Tasks desired to run to
	// completion.
	//
	// If this field is empty, the value of MaxConcurrent will be used.
	TotalCompletions *uint64 `json:",omitempty"`
}
```

---

### ReplicatedService

ReplicatedService is a kind of ServiceMode.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L83)  

```go
type ReplicatedService struct {
	Replicas *uint64 `json:",omitempty"`
}
```

---

### ResolutionMode

ResolutionMode represents a resolution mode.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/network.go#L21)  

```go
type ResolutionMode string
```

---

### ResourceRequirements

ResourceRequirements represents resources requirements.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L141)  

```go
type ResourceRequirements struct {
	Limits       *Limit     `json:",omitempty"`
	Reservations *Resources `json:",omitempty"`
}
```

---

### Resources

Resources represents resources (CPU/Memory) which can be advertised by a
node and requested to be reserved for a task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L102)  

```go
type Resources struct {
	NanoCPUs         int64             `json:",omitempty"`
	MemoryBytes      int64             `json:",omitempty"`
	GenericResources []GenericResource `json:",omitempty"`
}
```

---

### RestartPolicy

RestartPolicy represents the restart policy.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L172)  

```go
type RestartPolicy struct {
	Condition   RestartPolicyCondition `json:",omitempty"`
	Delay       *time.Duration         `json:",omitempty"`
	MaxAttempts *uint64                `json:",omitempty"`
	Window      *time.Duration         `json:",omitempty"`
}
```

---

### RestartPolicyCondition

RestartPolicyCondition represents when to restart.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L180)  

```go
type RestartPolicyCondition string
```

---

### RuntimeType

RuntimeType is the type of runtime used for the TaskSpec

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime.go#L4)  

```go
type RuntimeType string
```

---

### RuntimeURL

RuntimeURL is the proto type url

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime.go#L7)  

```go
type RuntimeURL string
```

---

### SELinuxContext

SELinuxContext contains the SELinux labels of the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/container.go#L25)  

```go
type SELinuxContext struct {
	Disable bool

	User  string
	Role  string
	Type  string
	Level string
}
```

---

### SeccompMode

SeccompMode is the type used for the enumeration of possible seccomp modes
in SeccompOpts

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/container.go#L36)  

```go
type SeccompMode string
```

---

### SeccompOpts

SeccompOpts defines the options for configuring seccomp on a swarm-managed
container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/container.go#L46)  

```go
type SeccompOpts struct {
	// Mode is the SeccompMode used for the container.
	Mode SeccompMode `json:",omitempty"`
	// Profile is the custom seccomp profile as a json object to be used with
	// the container. Mode should be set to SeccompModeCustom when using a
	// custom profile in this manner.
	Profile []byte `json:",omitempty"`
}
```

---

### Secret

Secret represents a secret.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/secret.go#L10)  

```go
type Secret struct {
	ID string
	Meta
	Spec SecretSpec
}
```

---

### SecretCreateResponse

SecretCreateResponse contains the information returned to a client
on the creation of a new secret.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/secret.go#L58)  

```go
type SecretCreateResponse struct {
	// ID is the id of the created secret.
	ID string
}
```

---

### SecretListOptions

SecretListOptions holds parameters to list secrets

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/secret.go#L64)  

```go
type SecretListOptions struct {
	Filters filters.Args
}
```

---

### SecretReference

SecretReference is a reference to a secret in swarm

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/secret.go#L50)  

```go
type SecretReference struct {
	File       *SecretReferenceFileTarget
	SecretID   string
	SecretName string
}
```

---

### SecretReferenceFileTarget

SecretReferenceFileTarget is a file target in a secret reference

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/secret.go#L42)  

```go
type SecretReferenceFileTarget struct {
	Name string
	UID  string
	GID  string
	Mode os.FileMode
}
```

---

### SecretSpec

SecretSpec represents a secret specification from a secret in swarm

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/secret.go#L17)  

```go
type SecretSpec struct {
	Annotations

	// Data is the data to store as a secret. It must be empty if a
	// [Driver] is used, in which case the data is loaded from an external
	// secret store. The maximum allowed size is 500KB, as defined in
	// [MaxSecretSize].
	//
	// This field is only used to create the secret, and is not returned
	// by other endpoints.
	//
	// [MaxSecretSize]: https://pkg.go.dev/github.com/moby/swarmkit/v2@v2.0.0-20250103191802-8c1959736554/api/validation#MaxSecretSize
	Data []byte `json:",omitempty"`

	// Driver is the name of the secrets driver used to fetch the secret's
	// value from an external secret store. If not set, the default built-in
	// store is used.
	Driver *Driver `json:",omitempty"`

	// Templating controls whether and how to evaluate the secret payload as
	// a template. If it is not set, no templating is used.
	Templating *Driver `json:",omitempty"`
}
```

---

### Service

Service represents a service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L10)  

```go
type Service struct {
	ID string
	Meta
	Spec         ServiceSpec   `json:",omitempty"`
	PreviousSpec *ServiceSpec  `json:",omitempty"`
	Endpoint     Endpoint      `json:",omitempty"`
	UpdateStatus *UpdateStatus `json:",omitempty"`

	// ServiceStatus is an optional, extra field indicating the number of
	// desired and running tasks. It is provided primarily as a shortcut to
	// calculating these values client-side, which otherwise would require
	// listing all tasks for a service, an operation that could be
	// computation and network expensive.
	ServiceStatus *ServiceStatus `json:",omitempty"`

	// JobStatus is the status of a Service which is in one of ReplicatedJob or
	// GlobalJob modes. It is absent on Replicated and Global services.
	JobStatus *JobStatus `json:",omitempty"`
}
```

---

### ServiceCreateOptions

ServiceCreateOptions contains the options to use when creating a service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L209)  

```go
type ServiceCreateOptions struct {
	// EncodedRegistryAuth is the encoded registry authorization credentials to
	// use when updating the service.
	//
	// This field follows the format of the X-Registry-Auth header.
	EncodedRegistryAuth string

	// QueryRegistry indicates whether the service update requires
	// contacting a registry. A registry may be contacted to retrieve
	// the image digest and manifest, which in turn can be used to update
	// platform or other information about the service.
	QueryRegistry bool
}
```

---

### ServiceCreateResponse

ServiceCreateResponse contains the information returned to a client on the
creation of a new service.

swagger:model ServiceCreateResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service_create_response.go#L10)  

```go
type ServiceCreateResponse struct {

	// The ID of the created service.
	ID string `json:"ID,omitempty"`

	// Optional warning message.
	//
	// FIXME(thaJeztah): this should have "omitempty" in the generated type.
	//
	Warnings []string `json:"Warnings"`
}
```

---

### ServiceInspectOptions

ServiceInspectOptions holds parameters related to the "service inspect"
operation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L270)  

```go
type ServiceInspectOptions struct {
	InsertDefaults bool
}
```

---

### ServiceListOptions

ServiceListOptions holds parameters to list services with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L260)  

```go
type ServiceListOptions struct {
	Filters filters.Args

	// Status indicates whether the server should include the service task
	// count of running and desired tasks.
	Status bool
}
```

---

### ServiceMode

ServiceMode represents the mode of a service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L49)  

```go
type ServiceMode struct {
	Replicated    *ReplicatedService `json:",omitempty"`
	Global        *GlobalService     `json:",omitempty"`
	ReplicatedJob *ReplicatedJob     `json:",omitempty"`
	GlobalJob     *GlobalJob         `json:",omitempty"`
}
```

---

### ServiceSpec

ServiceSpec represents the spec of a service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L31)  

```go
type ServiceSpec struct {
	Annotations

	// TaskTemplate defines how the service should construct new tasks when
	// orchestrating this service.
	TaskTemplate   TaskSpec      `json:",omitempty"`
	Mode           ServiceMode   `json:",omitempty"`
	UpdateConfig   *UpdateConfig `json:",omitempty"`
	RollbackConfig *UpdateConfig `json:",omitempty"`

	// Networks specifies which networks the service should attach to.
	//
	// Deprecated: This field is deprecated since v1.44. The Networks field in TaskSpec should be used instead.
	Networks     []NetworkAttachmentConfig `json:",omitempty"`
	EndpointSpec *EndpointSpec             `json:",omitempty"`
}
```

---

### ServiceStatus

ServiceStatus represents the number of running tasks in a service and the
number of tasks desired to be running.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L171)  

```go
type ServiceStatus struct {
	// RunningTasks is the number of tasks for the service actually in the
	// Running state
	RunningTasks uint64

	// DesiredTasks is the number of tasks desired to be running by the
	// service. For replicated services, this is the replica count. For global
	// services, this is computed by taking the number of tasks with desired
	// state of not-Shutdown.
	DesiredTasks uint64

	// CompletedTasks is the number of tasks in the state Completed, if this
	// service is in ReplicatedJob or GlobalJob mode. This field must be
	// cross-referenced with the service type, because the default value of 0
	// may mean that a service is not in a job mode, or it may mean that the
	// job has yet to complete any tasks.
	CompletedTasks uint64
}
```

---

### ServiceUpdateOptions

ServiceUpdateOptions contains the options to be used for updating services.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L230)  

```go
type ServiceUpdateOptions struct {
	// EncodedRegistryAuth is the encoded registry authorization credentials to
	// use when updating the service.
	//
	// This field follows the format of the X-Registry-Auth header.
	EncodedRegistryAuth string

	// RegistryAuthFrom specifies where to find the registry authorization
	// credentials if they are not given in EncodedRegistryAuth. Valid
	// values are "spec" and "previous-spec".
	RegistryAuthFrom string

	// Rollback indicates whether a server-side rollback should be
	// performed. When this is set, the provided spec will be ignored.
	// The valid values are "previous" and "none". An empty value is the
	// same as "none".
	Rollback string

	// QueryRegistry indicates whether the service update requires
	// contacting a registry. A registry may be contacted to retrieve
	// the image digest and manifest, which in turn can be used to update
	// platform or other information about the service.
	QueryRegistry bool
}
```

---

### ServiceUpdateResponse

ServiceUpdateResponse service update response
swagger:model ServiceUpdateResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service_update_response.go#L8)  

```go
type ServiceUpdateResponse struct {

	// Optional warning messages
	Warnings []string `json:"Warnings"`
}
```

---

### Spec

Spec represents the spec of a swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L35)  

```go
type Spec struct {
	Annotations

	Orchestration    OrchestrationConfig `json:",omitempty"`
	Raft             RaftConfig          `json:",omitempty"`
	Dispatcher       DispatcherConfig    `json:",omitempty"`
	CAConfig         CAConfig            `json:",omitempty"`
	TaskDefaults     TaskDefaults        `json:",omitempty"`
	EncryptionConfig EncryptionConfig    `json:",omitempty"`
}
```

---

### SpreadOver

SpreadOver is a scheduling preference that instructs the scheduler to spread
tasks evenly over groups of nodes identified by labels.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L166)  

```go
type SpreadOver struct {
	// label descriptor, such as engine.labels.az
	SpreadDescriptor string
}
```

---

### Status

Status provides information about the current swarm status and role,
obtained from the "Swarm" header in the API response.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L218)  

```go
type Status struct {
	// NodeState represents the state of the node.
	NodeState LocalNodeState

	// ControlAvailable indicates if the node is a swarm manager.
	ControlAvailable bool
}
```

---

### Swarm

Swarm represents a swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L21)  

```go
type Swarm struct {
	ClusterInfo
	JoinTokens JoinTokens
}
```

---

### TLSInfo

TLSInfo represents the TLS information about what CA certificate is trusted,
and who the issuer for a TLS certificate is

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/common.go#L39)  

```go
type TLSInfo struct {
	// TrustRoot is the trusted CA root certificate in PEM format
	TrustRoot string `json:",omitempty"`

	// CertIssuer is the raw subject bytes of the issuer
	CertIssuerSubject []byte `json:",omitempty"`

	// CertIssuerPublicKey is the raw public key bytes of the issuer
	CertIssuerPublicKey []byte `json:",omitempty"`
}
```

---

### Task

Task represents a task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L47)  

```go
type Task struct {
	ID string
	Meta
	Annotations

	Spec                TaskSpec            `json:",omitempty"`
	ServiceID           string              `json:",omitempty"`
	Slot                int                 `json:",omitempty"`
	NodeID              string              `json:",omitempty"`
	Status              TaskStatus          `json:",omitempty"`
	DesiredState        TaskState           `json:",omitempty"`
	NetworksAttachments []NetworkAttachment `json:",omitempty"`
	GenericResources    []GenericResource   `json:",omitempty"`

	// JobIteration is the JobIteration of the Service that this Task was
	// spawned from, if the Service is a ReplicatedJob or GlobalJob. This is
	// used to determine which Tasks belong to which run of the job. This field
	// is absent if the Service mode is Replicated or Global.
	JobIteration *Version `json:",omitempty"`

	// Volumes is the list of VolumeAttachments for this task. It specifies
	// which particular volumes are to be used by this particular task, and
	// fulfilling what mounts in the spec.
	Volumes []VolumeAttachment
}
```

---

### TaskDefaults

TaskDefaults parameterizes cluster-level task creation with default values.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L54)  

```go
type TaskDefaults struct {
	// LogDriver selects the log driver to use for tasks created in the
	// orchestrator if unspecified by a service.
	//
	// Updating this value will only have an affect on new tasks. Old tasks
	// will continue use their previously configured log driver until
	// recreated.
	LogDriver *Driver `json:",omitempty"`
}
```

---

### TaskListOptions

TaskListOptions holds parameters to list tasks with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L229)  

```go
type TaskListOptions struct {
	Filters filters.Args
}
```

---

### TaskSpec

TaskSpec represents the spec of a task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L74)  

```go
type TaskSpec struct {
	// ContainerSpec, NetworkAttachmentSpec, and PluginSpec are mutually exclusive.
	// PluginSpec is only used when the `Runtime` field is set to `plugin`
	// NetworkAttachmentSpec is used if the `Runtime` field is set to
	// `attachment`.
	ContainerSpec         *ContainerSpec         `json:",omitempty"`
	PluginSpec            *runtime.PluginSpec    `json:",omitempty"`
	NetworkAttachmentSpec *NetworkAttachmentSpec `json:",omitempty"`

	Resources     *ResourceRequirements     `json:",omitempty"`
	RestartPolicy *RestartPolicy            `json:",omitempty"`
	Placement     *Placement                `json:",omitempty"`
	Networks      []NetworkAttachmentConfig `json:",omitempty"`

	// LogDriver specifies the LogDriver to use for tasks created from this
	// spec. If not present, the one on cluster default on swarm.Spec will be
	// used, finally falling back to the engine default if not specified.
	LogDriver *Driver `json:",omitempty"`

	// ForceUpdate is a counter that triggers an update even if no relevant
	// parameters have been changed.
	ForceUpdate uint64

	Runtime RuntimeType `json:",omitempty"`
}
```

---

### TaskState

TaskState represents the state of a task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L11)  

```go
type TaskState string
```

---

### TaskStatus

TaskStatus represents the status of a task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L192)  

```go
type TaskStatus struct {
	Timestamp       time.Time        `json:",omitempty"`
	State           TaskState        `json:",omitempty"`
	Message         string           `json:",omitempty"`
	Err             string           `json:",omitempty"`
	ContainerStatus *ContainerStatus `json:",omitempty"`
	PortStatus      PortStatus       `json:",omitempty"`
}
```

---

### Topology

Topology defines the CSI topology of this node. This type is a duplicate of
github.com/docker/docker/api/types.Topology. Because the type definition
is so simple and to avoid complicated structure or circular imports, we just
duplicate it here. See that type for full documentation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/node.go#L139)  

```go
type Topology struct {
	Segments map[string]string `json:",omitempty"`
}
```

---

### UnlockKeyResponse

UnlockKeyResponse contains the response for Engine API:
GET /swarm/unlockkey

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L241)  

```go
type UnlockKeyResponse struct {
	// UnlockKey is the unlock key in ASCII-armored format.
	UnlockKey string
}
```

---

### UnlockRequest

UnlockRequest is the request used to unlock a swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L177)  

```go
type UnlockRequest struct {
	// UnlockKey is the unlock key in ASCII-armored format.
	UnlockKey string
}
```

---

### UpdateConfig

UpdateConfig represents the update configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L131)  

```go
type UpdateConfig struct {
	// Maximum number of tasks to be updated in one iteration.
	// 0 means unlimited parallelism.
	Parallelism uint64

	// Amount of time between updates.
	Delay time.Duration `json:",omitempty"`

	// FailureAction is the action to take when an update failures.
	FailureAction string `json:",omitempty"`

	// Monitor indicates how long to monitor a task for failure after it is
	// created. If the task fails by ending up in one of the states
	// REJECTED, COMPLETED, or FAILED, within Monitor from its creation,
	// this counts as a failure. If it fails after Monitor, it does not
	// count as a failure. If Monitor is unspecified, a default value will
	// be used.
	Monitor time.Duration `json:",omitempty"`

	// MaxFailureRatio is the fraction of tasks that may fail during
	// an update before the failure action is invoked. Any task created by
	// the current update which ends up in one of the states REJECTED,
	// COMPLETED or FAILED within Monitor from its creation counts as a
	// failure. The number of failures is divided by the number of tasks
	// being updated, and if this fraction is greater than
	// MaxFailureRatio, the failure action is invoked.
	//
	// If the failure action is CONTINUE, there is no effect.
	// If the failure action is PAUSE, no more tasks will be updated until
	// another update is started.
	MaxFailureRatio float32

	// Order indicates the order of operations when rolling out an updated
	// task. Either the old task is shut down before the new task is
	// started, or the new task is started before the old task is shut down.
	Order string
}
```

---

### UpdateFlags

UpdateFlags contains flags for SwarmUpdate.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/swarm.go#L233)  

```go
type UpdateFlags struct {
	RotateWorkerToken      bool
	RotateManagerToken     bool
	RotateManagerUnlockKey bool
}
```

---

### UpdateState

UpdateState is the state of a service update.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L57)  

```go
type UpdateState string
```

---

### UpdateStatus

UpdateStatus reports the status of a service update.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/service.go#L75)  

```go
type UpdateStatus struct {
	State       UpdateState `json:",omitempty"`
	StartedAt   *time.Time  `json:",omitempty"`
	CompletedAt *time.Time  `json:",omitempty"`
	Message     string      `json:",omitempty"`
}
```

---

### Version

Version represents the internal object version.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/common.go#L9)  

```go
type Version struct {
	Index uint64 `json:",omitempty"`
}
```

#### Methods

##### Version.String

String implements fmt.Stringer interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/common.go#L14)  

```go
func (v Version) String() string
```

---

### VolumeAttachment

VolumeAttachment contains the associating a Volume to a Task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/task.go#L215)  

```go
type VolumeAttachment struct {
	// ID is the Swarmkit ID of the Volume. This is not the CSI VolumeId.
	ID string `json:",omitempty"`

	// Source, together with Target, indicates the Mount, as specified in the
	// ContainerSpec, that this volume fulfills.
	Source string `json:",omitempty"`

	// Target, together with Source, indicates the Mount, as specified
	// in the ContainerSpec, that this volume fulfills.
	Target string `json:",omitempty"`
}
```

---


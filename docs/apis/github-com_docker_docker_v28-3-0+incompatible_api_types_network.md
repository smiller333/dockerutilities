# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/network

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:06 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L9)

```go
const (
	// NetworkDefault is a platform-independent alias to choose the platform-specific default network stack.
	NetworkDefault = "default"
	// NetworkHost is the name of the predefined network used when the NetworkMode host is selected (only available on Linux)
	NetworkHost = "host"
	// NetworkNone is the name of the predefined network used when the NetworkMode none is selected (available on both Linux and Windows)
	NetworkNone = "none"
	// NetworkBridge is the name of the default network on Linux
	NetworkBridge = "bridge"
	// NetworkNat is the name of the default network on Windows
	NetworkNat = "nat"
)
```

## Variables

This section is empty.

## Functions

### ValidateFilters

ValidateFilters validates the list of filter args with the available filters.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L160)  

```go
func ValidateFilters(filter filters.Args) error
```

---

### ValidateIPAM

ValidateIPAM checks whether the network's IPAM passed as argument is valid. It returns a joinError of the list of
errors found.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/ipam.go#L35)  

```go
func ValidateIPAM(ipam *IPAM, enableIPv6 bool) error
```

---

## Types

### Address

Address represents an IP address

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L101)  

```go
type Address struct {
	Addr      string
	PrefixLen int
}
```

---

### ConfigReference

ConfigReference specifies the source which provides a network's configuration

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L145)  

```go
type ConfigReference struct {
	Network string
}
```

---

### ConnectOptions

ConnectOptions represents the data to be used to connect a container to the
network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L61)  

```go
type ConnectOptions struct {
	Container      string
	EndpointConfig *EndpointSettings `json:",omitempty"`
}
```

---

### CreateOptions

CreateOptions holds options to create a network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L33)  

```go
type CreateOptions struct {
	Driver     string            // Driver is the driver-name used to create the network (e.g. `bridge`, `overlay`)
	Scope      string            // Scope describes the level at which the network exists (e.g. `swarm` for cluster-wide or `local` for machine level).
	EnableIPv4 *bool             `json:",omitempty"` // EnableIPv4 represents whether to enable IPv4.
	EnableIPv6 *bool             `json:",omitempty"` // EnableIPv6 represents whether to enable IPv6.
	IPAM       *IPAM             // IPAM is the network's IP Address Management.
	Internal   bool              // Internal represents if the network is used internal only.
	Attachable bool              // Attachable represents if the global scope is manually attachable by regular containers from workers in swarm mode.
	Ingress    bool              // Ingress indicates the network is providing the routing-mesh for the swarm cluster.
	ConfigOnly bool              // ConfigOnly creates a config-only network. Config-only networks are place-holder networks for network configurations to be used by other networks. ConfigOnly networks cannot be used directly to run containers or services.
	ConfigFrom *ConfigReference  // ConfigFrom specifies the source which will provide the configuration for this network. The specified network must be a config-only network; see [CreateOptions.ConfigOnly].
	Options    map[string]string // Options specifies the network-specific options to use for when creating the network.
	Labels     map[string]string // Labels holds metadata specific to the network being created.
}
```

---

### CreateRequest

CreateRequest is the request message sent to the server for network create call.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L23)  

```go
type CreateRequest struct {
	CreateOptions
	Name string // Name is the requested name of the network.

	// Deprecated: CheckDuplicate is deprecated since API v1.44, but it defaults to true when sent by the client
	// package to older daemons.
	CheckDuplicate *bool `json:",omitempty"`
}
```

---

### CreateResponse

CreateResponse NetworkCreateResponse

OK response to NetworkCreate operation
swagger:model CreateResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/create_response.go#L10)  

```go
type CreateResponse struct {

	// The ID of the created network.
	// Required: true
	ID string `json:"Id"`

	// Warnings encountered when creating the container
	// Required: true
	Warning string `json:"Warning"`
}
```

---

### DisconnectOptions

DisconnectOptions represents the data to be used to disconnect a container
from the network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L68)  

```go
type DisconnectOptions struct {
	Container string
	Force     bool
}
```

---

### EndpointIPAMConfig

EndpointIPAMConfig represents IPAM configurations for the endpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/endpoint.go#L68)  

```go
type EndpointIPAMConfig struct {
	IPv4Address  string   `json:",omitempty"`
	IPv6Address  string   `json:",omitempty"`
	LinkLocalIPs []string `json:",omitempty"`
}
```

#### Methods

##### EndpointIPAMConfig.Copy

Copy makes a copy of the endpoint ipam config

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/endpoint.go#L75)  

```go
func (cfg *EndpointIPAMConfig) Copy() *EndpointIPAMConfig
```

##### EndpointIPAMConfig.IsInRange

IsInRange checks whether static IP addresses are valid in a specific network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/endpoint.go#L92)  

```go
func (cfg *EndpointIPAMConfig) IsInRange(v4Subnets []NetworkSubnet, v6Subnets []NetworkSubnet) error
```

##### EndpointIPAMConfig.Validate

Validate checks whether cfg is valid.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/endpoint.go#L129)  

```go
func (cfg *EndpointIPAMConfig) Validate() error
```

---

### EndpointResource

EndpointResource contains network resources allocated and used for a
container in a network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L130)  

```go
type EndpointResource struct {
	Name        string
	EndpointID  string
	MacAddress  string
	IPv4Address string
	IPv6Address string
}
```

---

### EndpointSettings

EndpointSettings stores the network endpoint details

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/endpoint.go#L12)  

```go
type EndpointSettings struct {
	// Configurations
	IPAMConfig *EndpointIPAMConfig
	Links      []string
	Aliases    []string // Aliases holds the list of extra, user-specified DNS names for this endpoint.
	// MacAddress may be used to specify a MAC address when the container is created.
	// Once the container is running, it becomes operational data (it may contain a
	// generated address).
	MacAddress string
	DriverOpts map[string]string

	// GwPriority determines which endpoint will provide the default gateway
	// for the container. The endpoint with the highest priority will be used.
	// If multiple endpoints have the same priority, they are lexicographically
	// sorted based on their network name, and the one that sorts first is picked.
	GwPriority int
	// Operational data
	NetworkID           string
	EndpointID          string
	Gateway             string
	IPAddress           string
	IPPrefixLen         int
	IPv6Gateway         string
	GlobalIPv6Address   string
	GlobalIPv6PrefixLen int
	// DNSNames holds all the (non fully qualified) DNS names associated to this endpoint. First entry is used to
	// generate PTR records.
	DNSNames []string
}
```

#### Methods

##### EndpointSettings.Copy

Copy makes a deep copy of `EndpointSettings`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/endpoint.go#L43)  

```go
func (es *EndpointSettings) Copy() *EndpointSettings
```

---

### IPAM

IPAM represents IP Address Management

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/ipam.go#L12)  

```go
type IPAM struct {
	Driver  string
	Options map[string]string // Per network IPAM driver options
	Config  []IPAMConfig
}
```

---

### IPAMConfig

IPAMConfig represents IPAM configurations

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/ipam.go#L19)  

```go
type IPAMConfig struct {
	Subnet     string            `json:",omitempty"`
	IPRange    string            `json:",omitempty"`
	Gateway    string            `json:",omitempty"`
	AuxAddress map[string]string `json:"AuxiliaryAddresses,omitempty"`
}
```

---

### Inspect

Inspect is the body of the "get network" http response message.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L74)  

```go
type Inspect struct {
	Name       string                      // Name is the name of the network
	ID         string                      `json:"Id"` // ID uniquely identifies a network on a single machine
	Created    time.Time                   // Created is the time the network created
	Scope      string                      // Scope describes the level at which the network exists (e.g. `swarm` for cluster-wide or `local` for machine level)
	Driver     string                      // Driver is the Driver name used to create the network (e.g. `bridge`, `overlay`)
	EnableIPv4 bool                        // EnableIPv4 represents whether IPv4 is enabled
	EnableIPv6 bool                        // EnableIPv6 represents whether IPv6 is enabled
	IPAM       IPAM                        // IPAM is the network's IP Address Management
	Internal   bool                        // Internal represents if the network is used internal only
	Attachable bool                        // Attachable represents if the global scope is manually attachable by regular containers from workers in swarm mode.
	Ingress    bool                        // Ingress indicates the network is providing the routing-mesh for the swarm cluster.
	ConfigFrom ConfigReference             // ConfigFrom specifies the source which will provide the configuration for this network.
	ConfigOnly bool                        // ConfigOnly networks are place-holder networks for network configurations to be used by other networks. ConfigOnly networks cannot be used directly to run containers or services.
	Containers map[string]EndpointResource // Containers contains endpoints belonging to the network
	Options    map[string]string           // Options holds the network specific options to use for when creating the network
	Labels     map[string]string           // Labels holds metadata specific to the network being created
	Peers      []PeerInfo                  `json:",omitempty"` // List of peer nodes for an overlay network
	Services   map[string]ServiceInfo      `json:",omitempty"`
}
```

---

### InspectOptions

InspectOptions holds parameters to inspect network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L54)  

```go
type InspectOptions struct {
	Scope   string
	Verbose bool
}
```

---

### ListOptions

ListOptions holds parameters to filter the list of networks with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L49)  

```go
type ListOptions struct {
	Filters filters.Args
}
```

---

### NetworkSubnet

NetworkSubnet describes a user-defined subnet for a specific network. It's only used to validate if an
EndpointIPAMConfig is valid for a specific network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/endpoint.go#L84)  

```go
type NetworkSubnet interface {
	// Contains checks whether the NetworkSubnet contains [addr].
	Contains(addr net.IP) bool
	// IsStatic checks whether the subnet was statically allocated (ie. user-defined).
	IsStatic() bool
}
```

---

### NetworkingConfig

NetworkingConfig represents the container's networking configuration for each of its interfaces
Carries the networking configs specified in the `docker run` and `docker network connect` commands

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L140)  

```go
type NetworkingConfig struct {
	EndpointsConfig map[string]*EndpointSettings // Endpoint configs for each connecting network
}
```

---

### PeerInfo

PeerInfo represents one peer of an overlay network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L107)  

```go
type PeerInfo struct {
	Name string
	IP   string
}
```

---

### PruneReport

PruneReport contains the response for Engine API:
POST "/networks/prune"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L166)  

```go
type PruneReport struct {
	NetworksDeleted []string
}
```

---

### ServiceInfo

ServiceInfo represents service parameters with the list of service's tasks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L121)  

```go
type ServiceInfo struct {
	VIP          string
	Ports        []string
	LocalLBIndex int
	Tasks        []Task
}
```

---

### Summary

Summary is used as response when listing networks. It currently is an alias
for Inspect, but may diverge in the future, as not all information may
be included when listing networks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L98)  

```go
type Summary = Inspect
```

---

### Task

Task carries the information about one backend task

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/network/network.go#L113)  

```go
type Task struct {
	Name       string
	EndpointID string
	EndpointIP string
	Info       map[string]string
}
```

---


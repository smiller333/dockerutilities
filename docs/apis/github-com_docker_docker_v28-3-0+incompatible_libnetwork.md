# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:36 UTC

## Overview

Package libnetwork provides the basic functionality and extension points to
create network namespaces and allocate interfaces for containers to use.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L1164)

```go
var (
	ErrInvalidLengthAgent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAgent = fmt.Errorf("proto: unexpected end of group")
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L36)

```go
var PortConfig_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
	2: "SCTP",
}
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L42)

```go
var PortConfig_Protocol_value = map[string]int32{
	"TCP":  0,
	"UDP":  1,
	"SCTP": 2,
}
```

## Functions

### NetworkDeleteOptionRemoveLB

NetworkDeleteOptionRemoveLB informs a Network.Delete() operation that should
remove the load balancer endpoint for this network.  Note that the Delete()
method will automatically remove a load balancing endpoint for most networks
when the network is otherwise empty.  However, this does not occur for some
networks.  In particular, networks marked as ingress (which are supposed to
be more permanent than other overlay networks) won't automatically remove
the LB endpoint on Delete().  This method allows for explicit removal of
such networks provided there are no other endpoints present in the network.
If the network still has non-LB endpoints present, Delete() will not
remove the LB endpoint and will return an error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L942)  

```go
func NetworkDeleteOptionRemoveLB(p *networkDeleteParams)
```

---

## Types

### ActiveContainerError

ActiveContainerError is returned when an endpoint is deleted which has active
containers attached to it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L45)  

```go
type ActiveContainerError struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### ActiveContainerError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L50)  

```go
func (ace *ActiveContainerError) Error() string
```

##### ActiveContainerError.Forbidden

Forbidden denotes the type of this error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L55)  

```go
func (ace *ActiveContainerError) Forbidden()
```

---

### ActiveEndpointsError

ActiveEndpointsError is returned when a network is deleted which has active
endpoints in it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L30)  

```go
type ActiveEndpointsError struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### ActiveEndpointsError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L36)  

```go
func (aee *ActiveEndpointsError) Error() string
```

##### ActiveEndpointsError.Forbidden

Forbidden denotes the type of this error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L41)  

```go
func (aee *ActiveEndpointsError) Forbidden()
```

---

### ByNetworkType

ByNetworkType sorts a Endpoint slice based on the network-type
they're attached to. It implements sort.Interface and can be used
with sort.Stable or sort.Sort. It is used by Sandbox.ResolveName
when resolving names in swarm mode. In swarm mode, services with exposed
ports are connected to user overlay network, ingress network, and local
("docker_gwbridge") networks. Name resolution should prioritize returning
the VIP/IPs on user overlay network over ingress and local networks.

ByNetworkType re-orders the endpoints based on the network-type they
are attached to:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L42)  

```go
type ByNetworkType []*Endpoint
```

#### Methods

##### ByNetworkType.Len

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L44)  

```go
func (ep ByNetworkType) Len() int
```

##### ByNetworkType.Less

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L46)  

```go
func (ep ByNetworkType) Less(i, j int) bool
```

##### ByNetworkType.Swap

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L45)  

```go
func (ep ByNetworkType) Swap(i, j int)
```

---

### ByTime

ByTime implements sort.Interface for []*types.EncryptionKey based on
the LamportTime field.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.go#L32)  

```go
type ByTime []*types.EncryptionKey
```

#### Methods

##### ByTime.Len

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.go#L34)  

```go
func (b ByTime) Len() int
```

##### ByTime.Less

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.go#L36)  

```go
func (b ByTime) Less(i, j int) bool
```

##### ByTime.Swap

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.go#L35)  

```go
func (b ByTime) Swap(i, j int)
```

---

### Controller

Controller manages networks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L85)  

```go
type Controller struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates a new instance of network controller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L144)  

```go
func New(ctx context.Context, cfgOptions ...config.Option) (_ *Controller, retErr error)
```

#### Methods

##### Controller.AgentInitWait

AgentInitWait waits for agent initialization to be completed in the controller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L324)  

```go
func (c *Controller) AgentInitWait()
```

##### Controller.AgentStopWait

AgentStopWait waits for the Agent stop to be completed in the controller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L335)  

```go
func (c *Controller) AgentStopWait()
```

##### Controller.BuiltinDrivers

BuiltinDrivers returns the list of builtin network drivers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L405)  

```go
func (c *Controller) BuiltinDrivers() []string
```

##### Controller.BuiltinIPAMDrivers

BuiltinIPAMDrivers returns the list of builtin ipam drivers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L417)  

```go
func (c *Controller) BuiltinIPAMDrivers() []string
```

##### Controller.Config

Config returns the bootup configuration for the controller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L463)  

```go
func (c *Controller) Config() config.Config
```

##### Controller.FirewallBackend

FirewallBackend returns the name of the firewall backend for "docker info".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller_linux.go#L19)  

```go
func (c *Controller) FirewallBackend() *system.FirewallInfo
```

##### Controller.GetPluginGetter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L494)  

```go
func (c *Controller) GetPluginGetter() plugingetter.PluginGetter
```

##### Controller.GetSandbox

GetSandbox returns the Sandbox which has the passed id.

It returns an [ErrInvalidID] when passing an invalid ID, or an
types.NotFoundError if no Sandbox was found for the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L1012)  

```go
func (c *Controller) GetSandbox(containerID string) (*Sandbox, error)
```

##### Controller.ID

ID returns the controller's unique identity.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L400)  

```go
func (c *Controller) ID() string
```

##### Controller.IsDiagnosticEnabled

IsDiagnosticEnabled returns true if the diagnostic server is running.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L1142)  

```go
func (c *Controller) IsDiagnosticEnabled() bool
```

##### Controller.NetworkByID

NetworkByID returns the Network which has the passed id.
If not found, the error ErrNoSuchNetwork is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L898)  

```go
func (c *Controller) NetworkByID(id string) (*Network, error)
```

##### Controller.NetworkByName

NetworkByName returns the Network which has the passed name.
If not found, the error ErrNoSuchNetwork is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L875)  

```go
func (c *Controller) NetworkByName(name string) (*Network, error)
```

##### Controller.Networks

Networks returns the list of Network(s) managed by this controller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L851)  

```go
func (c *Controller) Networks(ctx context.Context) []*Network
```

##### Controller.NewNetwork

NewNetwork creates a new network of the specified network type. The options
are network specific and modeled in a generic way.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L510)  

```go
func (c *Controller) NewNetwork(ctx context.Context, networkType, name string, id string, options ...NetworkOption) (_ *Network, retErr error)
```

##### Controller.NewSandbox

NewSandbox creates a new sandbox for containerID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L906)  

```go
func (c *Controller) NewSandbox(ctx context.Context, containerID string, options ...SandboxOption) (_ *Sandbox, retErr error)
```

##### Controller.RegisterDriver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L498)  

```go
func (c *Controller) RegisterDriver(networkType string, driver driverapi.Driver, capability driverapi.Capability) error
```

##### Controller.SandboxByID

SandboxByID returns the Sandbox which has the passed id.
If not found, a types.NotFoundError is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L1036)  

```go
func (c *Controller) SandboxByID(id string) (*Sandbox, error)
```

##### Controller.SandboxDestroy

SandboxDestroy destroys a sandbox given a container ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L1050)  

```go
func (c *Controller) SandboxDestroy(ctx context.Context, id string) error
```

##### Controller.SetClusterProvider

SetClusterProvider sets the cluster provider.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L221)  

```go
func (c *Controller) SetClusterProvider(provider cluster.Provider)
```

##### Controller.SetKeys

SetKeys configures the encryption key for gossip and overlay data path.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L243)  

```go
func (c *Controller) SetKeys(keys []*types.EncryptionKey) error
```

##### Controller.StartDiagnostic

StartDiagnostic starts the network diagnostic server listening on port.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L1132)  

```go
func (c *Controller) StartDiagnostic(port int)
```

##### Controller.Stop

Stop stops the network controller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L1126)  

```go
func (c *Controller) Stop()
```

##### Controller.StopDiagnostic

StopDiagnostic stops the network diagnostic server.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L1137)  

```go
func (c *Controller) StopDiagnostic()
```

##### Controller.WalkNetworks

WalkNetworks uses the provided function to walk the Network(s) managed by this controller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L865)  

```go
func (c *Controller) WalkNetworks(walker NetworkWalker)
```

---

### DNSBackend

DNSBackend represents a backend DNS resolver used for DNS name
resolution. All the queries to the resolver are forwarded to the
backend resolver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L31)  

```go
type DNSBackend interface {
	// ResolveName resolves a service name to an IPv4 or IPv6 address by searching
	// the networks the sandbox is connected to. The second return value will be
	// true if the name exists in docker domain, even if there are no addresses of
	// the required type. Such queries shouldn't be forwarded to external nameservers.
	ResolveName(ctx context.Context, name string, ipType int) ([]net.IP, bool)
	// ResolveIP returns the service name for the passed in IP. IP is in reverse dotted
	// notation; the format used for DNS PTR records
	ResolveIP(ctx context.Context, name string) string
	// ResolveService returns all the backend details about the containers or hosts
	// backing a service. Its purpose is to satisfy an SRV query
	ResolveService(ctx context.Context, name string) ([]*net.SRV, []net.IP)
	// ExecFunc allows a function to be executed in the context of the backend
	// on behalf of the resolver.
	ExecFunc(f func()) error
	// NdotsSet queries the backends ndots dns option settings
	NdotsSet() bool
	// HandleQueryResp passes the name & IP from a response to the backend. backend
	// can use it to maintain any required state about the resolution
	HandleQueryResp(name string, ip net.IP)
}
```

---

### Endpoint

Endpoint represents a logical connection between a network and a sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L75)  

```go
type Endpoint struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Endpoint.AddStaticRoute

AddStaticRoute adds a route to the sandbox.
It may be used in addition to or instead of a default gateway (as above).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L303)  

```go
func (ep *Endpoint) AddStaticRoute(destination *net.IPNet, routeType int, nextHop net.IP) error
```

##### Endpoint.AddTableEntry

AddTableEntry adds a table entry to the gossip layer
passing the table name, key and an opaque value.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L322)  

```go
func (ep *Endpoint) AddTableEntry(tableName, key string, value []byte) error
```

##### Endpoint.CopyTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L275)  

```go
func (ep *Endpoint) CopyTo(o datastore.KVObject) error
```

##### Endpoint.Delete

Delete deletes and detaches this endpoint from the network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1007)  

```go
func (ep *Endpoint) Delete(ctx context.Context, force bool) error
```

##### Endpoint.DisableGatewayService

DisableGatewayService tells libnetwork not to provide Default GW for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L451)  

```go
func (ep *Endpoint) DisableGatewayService()
```

##### Endpoint.DriverInfo

DriverInfo returns a collection of driver operational data related to this endpoint retrieved from the driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info_unix.go#L8)  

```go
func (ep *Endpoint) DriverInfo() (map[string]interface{}, error)
```

##### Endpoint.Exists

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L445)  

```go
func (ep *Endpoint) Exists() bool
```

##### Endpoint.Gateway

Gateway returns the IPv4 gateway assigned by the driver.
This will only return a valid value if a container has joined the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L366)  

```go
func (ep *Endpoint) Gateway() net.IP
```

##### Endpoint.GatewayIPv6

GatewayIPv6 returns the IPv6 gateway assigned by the driver.
This will only return a valid value if a container has joined the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L379)  

```go
func (ep *Endpoint) GatewayIPv6() net.IP
```

##### Endpoint.ID

ID returns the system-generated id for this endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L327)  

```go
func (ep *Endpoint) ID() string
```

##### Endpoint.Iface

Iface returns information about the interface which was assigned to
the endpoint by the driver. This can be used after the
endpoint has been created.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L212)  

```go
func (ep *Endpoint) Iface() *EndpointInterface
```

##### Endpoint.Index

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L432)  

```go
func (ep *Endpoint) Index() uint64
```

##### Endpoint.Info

Info hydrates the endpoint and returns certain operational data belonging
to this endpoint.

TODO(thaJeztah): make sure that Endpoint is always fully hydrated, and remove the EndpointInfo interface, and use Endpoint directly.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L185)  

```go
func (ep *Endpoint) Info() EndpointInfo
```

##### Endpoint.InterfaceName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L295)  

```go
func (ep *Endpoint) InterfaceName() driverapi.InterfaceNameInfo
```

##### Endpoint.Join

Join joins the sandbox to the endpoint and populates into the sandbox
the network resources allocated for the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L483)  

```go
func (ep *Endpoint) Join(ctx context.Context, sb *Sandbox, options ...EndpointOption) error
```

##### Endpoint.Key

Key returns the endpoint's key.

Key structure: endpoint/network-id/endpoint-id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L392)  

```go
func (ep *Endpoint) Key() []string
```

##### Endpoint.KeyPrefix

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L400)  

```go
func (ep *Endpoint) KeyPrefix() []string
```

##### Endpoint.Leave

Leave detaches the network resources populated in the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L819)  

```go
func (ep *Endpoint) Leave(ctx context.Context, sb *Sandbox) error
```

##### Endpoint.Less

Less defines an ordering over endpoints, with better candidates for the default
gateway sorted first.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L630)  

```go
func (ep *Endpoint) Less(epj *Endpoint) bool
```

##### Endpoint.LoadBalancer

LoadBalancer returns whether the endpoint is the load balancer endpoint for the network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L345)  

```go
func (ep *Endpoint) LoadBalancer() bool
```

##### Endpoint.MarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L105)  

```go
func (ep *Endpoint) MarshalJSON() ([]byte, error)
```

##### Endpoint.Name

Name returns the name of this endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L335)  

```go
func (ep *Endpoint) Name() string
```

##### Endpoint.Network

Network returns the name of the network to which this endpoint is attached.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L343)  

```go
func (ep *Endpoint) Network() string
```

##### Endpoint.New

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L271)  

```go
func (ep *Endpoint) New() datastore.KVObject
```

##### Endpoint.Sandbox

Sandbox returns the attached sandbox if there, nil otherwise.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L336)  

```go
func (ep *Endpoint) Sandbox() *Sandbox
```

##### Endpoint.SetGateway

SetGateway sets the default IPv4 gateway when a container joins the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L391)  

```go
func (ep *Endpoint) SetGateway(gw net.IP) error
```

##### Endpoint.SetGatewayIPv6

SetGatewayIPv6 sets the default IPv6 gateway when a container joins the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L400)  

```go
func (ep *Endpoint) SetGatewayIPv6(gw6 net.IP) error
```

##### Endpoint.SetIndex

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L438)  

```go
func (ep *Endpoint) SetIndex(index uint64)
```

##### Endpoint.SetValue

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L428)  

```go
func (ep *Endpoint) SetValue(value []byte) error
```

##### Endpoint.Skip

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L451)  

```go
func (ep *Endpoint) Skip() bool
```

##### Endpoint.StaticRoutes

StaticRoutes returns the list of static routes configured by the network
driver when the container joins a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L353)  

```go
func (ep *Endpoint) StaticRoutes() []*types.StaticRoute
```

##### Endpoint.UnmarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L132)  

```go
func (ep *Endpoint) UnmarshalJSON(b []byte) (err error)
```

##### Endpoint.UpdateDNSNames

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L775)  

```go
func (ep *Endpoint) UpdateDNSNames(dnsNames []string) error
```

##### Endpoint.Value

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L408)  

```go
func (ep *Endpoint) Value() []byte
```

---

### EndpointInfo

EndpointInfo provides an interface to retrieve network resources bound to the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L13)  

```go
type EndpointInfo interface {
	// Iface returns information about the interface which was assigned to
	// the endpoint by the driver. This can be used after the
	// endpoint has been created.
	Iface() *EndpointInterface

	// Gateway returns the IPv4 gateway assigned by the driver.
	// This will only return a valid value if a container has joined the endpoint.
	Gateway() net.IP

	// GatewayIPv6 returns the IPv6 gateway assigned by the driver.
	// This will only return a valid value if a container has joined the endpoint.
	GatewayIPv6() net.IP

	// StaticRoutes returns the list of static routes configured by the network
	// driver when the container joins a network
	StaticRoutes() []*types.StaticRoute

	// Sandbox returns the attached sandbox if there, nil otherwise.
	Sandbox() *Sandbox

	// LoadBalancer returns whether the endpoint is the load balancer endpoint for the network.
	LoadBalancer() bool
}
```

---

### EndpointInterface

EndpointInterface holds interface addresses bound to the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L39)  

```go
type EndpointInterface struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### EndpointInterface.Address

Address returns the IPv4 address assigned to the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L255)  

```go
func (epi *EndpointInterface) Address() *net.IPNet
```

##### EndpointInterface.AddressIPv6

AddressIPv6 returns the IPv6 address assigned to the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L260)  

```go
func (epi *EndpointInterface) AddressIPv6() *net.IPNet
```

##### EndpointInterface.CopyTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L145)  

```go
func (epi *EndpointInterface) CopyTo(dstEpi *EndpointInterface) error
```

##### EndpointInterface.LinkLocalAddresses

LinkLocalAddresses returns the list of link-local (IPv4/IPv6) addresses assigned to the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L265)  

```go
func (epi *EndpointInterface) LinkLocalAddresses() []*net.IPNet
```

##### EndpointInterface.MacAddress

MacAddress returns the MAC address assigned to the endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L250)  

```go
func (epi *EndpointInterface) MacAddress() net.HardwareAddr
```

##### EndpointInterface.MarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L54)  

```go
func (epi *EndpointInterface) MarshalJSON() ([]byte, error)
```

##### EndpointInterface.NetnsPath

NetnsPath returns the path of the network namespace, if there is one. Else "".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L284)  

```go
func (epi *EndpointInterface) NetnsPath() string
```

##### EndpointInterface.SetCreatedInContainer

SetCreatedInContainer can be called by the driver to indicate that it's
created the network interface in the container's network namespace (so,
it doesn't need to be moved there).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L291)  

```go
func (epi *EndpointInterface) SetCreatedInContainer(cic bool)
```

##### EndpointInterface.SetIPAddress

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L231)  

```go
func (epi *EndpointInterface) SetIPAddress(address *net.IPNet) error
```

##### EndpointInterface.SetMacAddress

SetMacAddress allows the driver to set the mac address to the endpoint interface
during the call to CreateEndpoint, if the mac address is not already set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L220)  

```go
func (epi *EndpointInterface) SetMacAddress(mac net.HardwareAddr) error
```

##### EndpointInterface.SetNames

SetNames method assigns the srcName, dstName, and dstPrefix for the
interface. If both dstName and dstPrefix are set, dstName takes precedence.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L276)  

```go
func (epi *EndpointInterface) SetNames(srcName, dstPrefix, dstName string) error
```

##### EndpointInterface.SrcName

SrcName returns the name of the interface w/in the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L270)  

```go
func (epi *EndpointInterface) SrcName() string
```

##### EndpointInterface.UnmarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint_info.go#L86)  

```go
func (epi *EndpointInterface) UnmarshalJSON(b []byte) error
```

---

### EndpointOption

EndpointOption is an option setter function type used to pass various options to Network
and Endpoint interfaces methods. The various setter functions of type EndpointOption are
provided by libnetwork, they look like <Create|Join|Leave>Option[...](...)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L72)  

```go
type EndpointOption func(ep *Endpoint)
```

#### Functions

##### CreateOptionAlias

CreateOptionAlias function returns an option setter for setting endpoint alias

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1218)  

```go
func CreateOptionAlias(name string, alias string) EndpointOption
```

##### CreateOptionDNS

CreateOptionDNS function returns an option setter for dns entry option to
be passed to container Create method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1187)  

```go
func CreateOptionDNS(dns []string) EndpointOption
```

##### CreateOptionDNSNames

CreateOptionDNSNames specifies the list of (non fully qualified) DNS names associated to an endpoint. These will be
used to populate the embedded DNS server. Order matters: first name will be used to generate PTR records.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1195)  

```go
func CreateOptionDNSNames(names []string) EndpointOption
```

##### CreateOptionDisableIPv6

CreateOptionDisableIPv6 prevents allocation of an IPv6 address/gateway, even
if the container is connected to an IPv6-enabled network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1211)  

```go
func CreateOptionDisableIPv6() EndpointOption
```

##### CreateOptionDisableResolution

CreateOptionDisableResolution function returns an option setter to indicate
this endpoint doesn't want embedded DNS server functionality

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1203)  

```go
func CreateOptionDisableResolution() EndpointOption
```

##### CreateOptionExposedPorts

CreateOptionExposedPorts function returns an option setter for the container exposed
ports option to be passed to Network.CreateEndpoint method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1163)  

```go
func CreateOptionExposedPorts(exposedPorts []types.TransportPort) EndpointOption
```

##### CreateOptionIpam

CreateOptionIpam function returns an option setter for the ipam configuration for this endpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1144)  

```go
func CreateOptionIpam(ipV4, ipV6 net.IP, llIPs []net.IP, ipamOptions map[string]string) EndpointOption
```

##### CreateOptionLoadBalancer

CreateOptionLoadBalancer function returns an option setter for denoting the endpoint is a load balancer for a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1239)  

```go
func CreateOptionLoadBalancer() EndpointOption
```

##### CreateOptionPortMapping

CreateOptionPortMapping function returns an option setter for the mapping
ports option to be passed to Network.CreateEndpoint method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1176)  

```go
func CreateOptionPortMapping(portBindings []types.PortBinding) EndpointOption
```

##### CreateOptionService

CreateOptionService function returns an option setter for setting service binding configuration

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1228)  

```go
func CreateOptionService(name, id string, vip net.IP, ingressPorts []*PortConfig, aliases []string) EndpointOption
```

##### EndpointOptionGeneric

EndpointOptionGeneric function returns an option setter for a Generic option defined
in a Dictionary of Key-Value pair

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1130)  

```go
func EndpointOptionGeneric(generic map[string]any) EndpointOption
```

##### JoinOptionPriority

JoinOptionPriority function returns an option setter for priority option to
be passed to the endpoint.Join() method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1247)  

```go
func JoinOptionPriority(prio int) EndpointOption
```

##### WithNetnsPath

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/endpoint.go#L1262)  

```go
func WithNetnsPath(path string) EndpointOption
```

---

### EndpointRecord

EndpointRecord specifies all the endpoint specific information that
needs to gossiped to nodes participating in the network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L58)  

```go
type EndpointRecord struct {
	// Name of the container
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Service name of the service to which this endpoint belongs.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Service ID of the service to which this endpoint belongs.
	ServiceID string `protobuf:"bytes,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// Virtual IP of the service to which this endpoint belongs.
	VirtualIP string `protobuf:"bytes,4,opt,name=virtual_ip,json=virtualIp,proto3" json:"virtual_ip,omitempty"`
	// IP assigned to this endpoint.
	EndpointIP string `protobuf:"bytes,5,opt,name=endpoint_ip,json=endpointIp,proto3" json:"endpoint_ip,omitempty"`
	// IngressPorts exposed by the service to which this endpoint belongs.
	IngressPorts []*PortConfig `protobuf:"bytes,6,rep,name=ingress_ports,json=ingressPorts,proto3" json:"ingress_ports,omitempty"`
	// A list of aliases which are alternate names for the service
	Aliases []string `protobuf:"bytes,7,rep,name=aliases,proto3" json:"aliases,omitempty"`
	// List of aliases task specific aliases
	TaskAliases []string `protobuf:"bytes,8,rep,name=task_aliases,json=taskAliases,proto3" json:"task_aliases,omitempty"`
	// Whether this endpoint's service has been disabled
	ServiceDisabled bool `protobuf:"varint,9,opt,name=service_disabled,json=serviceDisabled,proto3" json:"service_disabled,omitempty"`
}
```

#### Methods

##### EndpointRecord.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L81)  

```go
func (*EndpointRecord) Descriptor() ([]byte, []int)
```

##### EndpointRecord.GetAliases

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L153)  

```go
func (m *EndpointRecord) GetAliases() []string
```

##### EndpointRecord.GetEndpointIP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L139)  

```go
func (m *EndpointRecord) GetEndpointIP() string
```

##### EndpointRecord.GetIngressPorts

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L146)  

```go
func (m *EndpointRecord) GetIngressPorts() []*PortConfig
```

##### EndpointRecord.GetName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L111)  

```go
func (m *EndpointRecord) GetName() string
```

##### EndpointRecord.GetServiceDisabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L167)  

```go
func (m *EndpointRecord) GetServiceDisabled() bool
```

##### EndpointRecord.GetServiceID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L125)  

```go
func (m *EndpointRecord) GetServiceID() string
```

##### EndpointRecord.GetServiceName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L118)  

```go
func (m *EndpointRecord) GetServiceName() string
```

##### EndpointRecord.GetTaskAliases

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L160)  

```go
func (m *EndpointRecord) GetTaskAliases() []string
```

##### EndpointRecord.GetVirtualIP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L132)  

```go
func (m *EndpointRecord) GetVirtualIP() string
```

##### EndpointRecord.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L299)  

```go
func (this *EndpointRecord) GoString() string
```

##### EndpointRecord.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L340)  

```go
func (m *EndpointRecord) Marshal() (dAtA []byte, err error)
```

##### EndpointRecord.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L350)  

```go
func (m *EndpointRecord) MarshalTo(dAtA []byte) (int, error)
```

##### EndpointRecord.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L355)  

```go
func (m *EndpointRecord) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### EndpointRecord.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L80)  

```go
func (*EndpointRecord) ProtoMessage()
```

##### EndpointRecord.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L79)  

```go
func (m *EndpointRecord) Reset()
```

##### EndpointRecord.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L496)  

```go
func (m *EndpointRecord) Size() (n int)
```

##### EndpointRecord.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L574)  

```go
func (this *EndpointRecord) String() string
```

##### EndpointRecord.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L618)  

```go
func (m *EndpointRecord) Unmarshal(dAtA []byte) error
```

##### EndpointRecord.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L105)  

```go
func (m *EndpointRecord) XXX_DiscardUnknown()
```

##### EndpointRecord.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L87)  

```go
func (m *EndpointRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### EndpointRecord.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L99)  

```go
func (m *EndpointRecord) XXX_Merge(src proto.Message)
```

##### EndpointRecord.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L102)  

```go
func (m *EndpointRecord) XXX_Size() int
```

##### EndpointRecord.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L84)  

```go
func (m *EndpointRecord) XXX_Unmarshal(b []byte) error
```

---

### EndpointWalker

EndpointWalker is a client provided function which will be used to walk the Endpoints.
When the function returns true, the walk will stop.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L41)  

```go
type EndpointWalker func(ep *Endpoint) bool
```

---

### ErrNoSuchNetwork

ErrNoSuchNetwork is returned when a network query finds no result

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L9)  

```go
type ErrNoSuchNetwork string
```

#### Methods

##### ErrNoSuchNetwork.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L11)  

```go
func (nsn ErrNoSuchNetwork) Error() string
```

##### ErrNoSuchNetwork.NotFound

NotFound denotes the type of this error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L16)  

```go
func (nsn ErrNoSuchNetwork) NotFound()
```

---

### IpamConf

IpamConf contains all the ipam related configurations for a network

TODO(aker): use proper net/* structs instead of string literals.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L87)  

```go
type IpamConf struct {
	// PreferredPool is the master address pool for containers and network interfaces.
	PreferredPool string
	// SubPool is a subset of the master pool. If specified,
	// this becomes the container pool for automatic address allocations.
	SubPool string
	// Gateway is the preferred Network Gateway address (optional).
	Gateway string
	// AuxAddresses contains auxiliary addresses for network driver. Must be within the master pool.
	// libnetwork will reserve them if they fall into the container pool.
	AuxAddresses map[string]string
}
```

#### Methods

##### IpamConf.Contains

Contains checks whether the ipam master address pool contains [addr].

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L109)  

```go
func (c *IpamConf) Contains(addr net.IP) bool
```

##### IpamConf.CopyTo

CopyTo deep copies to the destination IpamConfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L321)  

```go
func (c *IpamConf) CopyTo(dstC *IpamConf) error
```

##### IpamConf.IsStatic

IsStatic checks whether the subnet was statically allocated (ie. user-defined).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L123)  

```go
func (c *IpamConf) IsStatic() bool
```

##### IpamConf.Validate

Validate checks whether the configuration is valid

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L101)  

```go
func (c *IpamConf) Validate() error
```

---

### IpamInfo

IpamInfo contains all the ipam related operational info for a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L128)  

```go
type IpamInfo struct {
	PoolID string
	Meta   map[string]string
	driverapi.IPAMData
}
```

#### Methods

##### IpamInfo.CopyTo

CopyTo deep copies to the destination IpamInfo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L335)  

```go
func (i *IpamInfo) CopyTo(dstI *IpamInfo) error
```

##### IpamInfo.MarshalJSON

MarshalJSON encodes IpamInfo into json message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L135)  

```go
func (i *IpamInfo) MarshalJSON() ([]byte, error)
```

##### IpamInfo.UnmarshalJSON

UnmarshalJSON decodes json message into PoolData

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L152)  

```go
func (i *IpamInfo) UnmarshalJSON(data []byte) error
```

---

### ManagerRedirectError

ManagerRedirectError is returned when the request should be redirected to Manager

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L58)  

```go
type ManagerRedirectError string
```

#### Methods

##### ManagerRedirectError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L60)  

```go
func (mr ManagerRedirectError) Error() string
```

##### ManagerRedirectError.Maskable

Maskable denotes the type of this error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L65)  

```go
func (mr ManagerRedirectError) Maskable()
```

---

### Network

Network represents a logical connectivity zone that containers may
join using the Link method. A network is managed by a specific driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L177)  

```go
type Network struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Network.Attachable

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1832)  

```go
func (n *Network) Attachable() bool
```

##### Network.ConfigFrom

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1867)  

```go
func (n *Network) ConfigFrom() string
```

##### Network.ConfigOnly

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1874)  

```go
func (n *Network) ConfigOnly() bool
```

##### Network.CopyTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L461)  

```go
func (n *Network) CopyTo(o datastore.KVObject) error
```

##### Network.CreateEndpoint

CreateEndpoint creates a new endpoint to this network symbolically identified by the
specified unique name. The options parameter carries driver specific options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1179)  

```go
func (n *Network) CreateEndpoint(ctx context.Context, name string, options ...EndpointOption) (*Endpoint, error)
```

##### Network.Created

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L238)  

```go
func (n *Network) Created() time.Time
```

##### Network.Delete

Delete the network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1000)  

```go
func (n *Network) Delete(options ...NetworkDeleteOption) error
```

##### Network.DriverOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1758)  

```go
func (n *Network) DriverOptions() map[string]string
```

##### Network.Dynamic

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1846)  

```go
func (n *Network) Dynamic() bool
```

##### Network.EndpointByName

EndpointByName returns the Endpoint which has the passed name. If not found,
an errdefs.ErrNotFound is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1311)  

```go
func (n *Network) EndpointByName(name string) (*Endpoint, error)
```

##### Network.Endpoints

Endpoints returns the list of Endpoint(s) in this network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1292)  

```go
func (n *Network) Endpoints() []*Endpoint
```

##### Network.Exists

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L297)  

```go
func (n *Network) Exists() bool
```

##### Network.HandleQueryResp

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1979)  

```go
func (n *Network) HandleQueryResp(name string, ip net.IP)
```

##### Network.ID

ID returns a system generated id for this network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L231)  

```go
func (n *Network) ID() string
```

##### Network.IPv4Enabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1853)  

```go
func (n *Network) IPv4Enabled() bool
```

##### Network.IPv6Enabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1860)  

```go
func (n *Network) IPv6Enabled() bool
```

##### Network.Index

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L284)  

```go
func (n *Network) Index() uint64
```

##### Network.Ingress

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1839)  

```go
func (n *Network) Ingress() bool
```

##### Network.Internal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1825)  

```go
func (n *Network) Internal() bool
```

##### Network.IpamConfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1775)  

```go
func (n *Network) IpamConfig() (ipamType string, ipamOptions map[string]string, ipamV4Config []*IpamConf, ipamV6Config []*IpamConf)
```

##### Network.IpamInfo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1800)  

```go
func (n *Network) IpamInfo() (ipamV4Info []*IpamInfo, ipamV6Info []*IpamInfo)
```

##### Network.IsPruneable

IsPruneable returns true if n can be considered for removal as part of a
"docker network prune" (or system prune). The caller must still check that the
network should be removed. For example, it may have active endpoints.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network_unix.go#L78)  

```go
func (n *Network) IsPruneable() bool
```

##### Network.Key

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L260)  

```go
func (n *Network) Key() []string
```

##### Network.KeyPrefix

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L266)  

```go
func (n *Network) KeyPrefix() []string
```

##### Network.Labels

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1881)  

```go
func (n *Network) Labels() map[string]string
```

##### Network.MarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L574)  

```go
func (n *Network) MarshalJSON() ([]byte, error)
```

##### Network.Name

Name returns a user chosen name for this network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L223)  

```go
func (n *Network) Name() string
```

##### Network.NdotsSet

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L2084)  

```go
func (n *Network) NdotsSet() bool
```

##### Network.New

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L309)  

```go
func (n *Network) New() datastore.KVObject
```

##### Network.Peers

Peers returns a slice of PeerInfo structures which has the information about the peer
nodes participating in the same overlay network. This is currently the per-network
gossip cluster. For non-dynamic overlay networks and bridge networks it returns an
empty slice

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1745)  

```go
func (n *Network) Peers() []networkdb.PeerInfo
```

##### Network.ResolveIP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1999)  

```go
func (n *Network) ResolveIP(_ context.Context, ip string) string
```

##### Network.ResolveName

ResolveName looks up addresses of ipType for name req.
Returns (addresses, true) if req is found, but len(addresses) may be 0 if
there are no addresses of ipType. If the name is not found, the bool return
will be false.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1936)  

```go
func (n *Network) ResolveName(ctx context.Context, req string, ipType int) ([]net.IP, bool)
```

##### Network.ResolveService

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L2029)  

```go
func (n *Network) ResolveService(ctx context.Context, name string) ([]*net.SRV, []net.IP)
```

##### Network.Resolvers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L253)  

```go
func (n *Network) Resolvers() []*Resolver
```

##### Network.Scope

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1769)  

```go
func (n *Network) Scope() string
```

##### Network.Services

Services returns a map of services keyed by the service name with the details
of all the tasks that belong to the service. Applicable only in swarm mode.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.go#L463)  

```go
func (n *Network) Services() map[string]ServiceInfo
```

##### Network.SetIndex

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L290)  

```go
func (n *Network) SetIndex(index uint64)
```

##### Network.SetValue

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L280)  

```go
func (n *Network) SetValue(value []byte) error
```

##### Network.Skip

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L303)  

```go
func (n *Network) Skip() bool
```

##### Network.TableEventRegister

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1893)  

```go
func (n *Network) TableEventRegister(tableName string, objType driverapi.ObjectType) error
```

##### Network.Type

Type returns the type of network, which corresponds to its managing driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L246)  

```go
func (n *Network) Type() string
```

##### Network.UnmarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L633)  

```go
func (n *Network) UnmarshalJSON(b []byte) (err error)
```

##### Network.UpdateIpamConfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1908)  

```go
func (n *Network) UpdateIpamConfig(ipV4Data []driverapi.IPAMData)
```

##### Network.Value

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L270)  

```go
func (n *Network) Value() []byte
```

##### Network.WalkEndpoints

WalkEndpoints uses the provided function to walk the Endpoints.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L1301)  

```go
func (n *Network) WalkEndpoints(walker EndpointWalker)
```

---

### NetworkDeleteOption

NetworkDeleteOption is a type for optional parameters to pass to the
Network.Delete() function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L930)  

```go
type NetworkDeleteOption func(p *networkDeleteParams)
```

---

### NetworkNameError

NetworkNameError is returned when a network with the same name already exists.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L19)  

```go
type NetworkNameError string
```

#### Methods

##### NetworkNameError.Conflict

Conflict denotes the type of this error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L26)  

```go
func (nnr NetworkNameError) Conflict()
```

##### NetworkNameError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/error.go#L21)  

```go
func (nnr NetworkNameError) Error() string
```

---

### NetworkOption

NetworkOption is an option setter function type used to pass various options to
NewNetwork method. The various setter functions of type NetworkOption are
provided by libnetwork, they look like NetworkOptionXXXX(...)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L760)  

```go
type NetworkOption func(n *Network)
```

#### Functions

##### NetworkOptionAttachable

NetworkOptionAttachable returns an option setter to set attachable for a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L834)  

```go
func NetworkOptionAttachable(attachable bool) NetworkOption
```

##### NetworkOptionConfigFrom

NetworkOptionConfigFrom tells controller to pick the
network configuration from a configuration only network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L910)  

```go
func NetworkOptionConfigFrom(name string) NetworkOption
```

##### NetworkOptionConfigOnly

NetworkOptionConfigOnly tells controller this network is
a configuration only network. It serves as a configuration
for other networks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L902)  

```go
func NetworkOptionConfigOnly() NetworkOption
```

##### NetworkOptionDriverOpts

NetworkOptionDriverOpts function returns an option setter for any driver parameter described by a map

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L872)  

```go
func NetworkOptionDriverOpts(opts map[string]string) NetworkOption
```

##### NetworkOptionDynamic

NetworkOptionDynamic function returns an option setter for dynamic option for a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L893)  

```go
func NetworkOptionDynamic() NetworkOption
```

##### NetworkOptionEnableIPv4

NetworkOptionEnableIPv4 returns an option setter to explicitly configure IPv4

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L800)  

```go
func NetworkOptionEnableIPv4(enableIPv4 bool) NetworkOption
```

##### NetworkOptionEnableIPv6

NetworkOptionEnableIPv6 returns an option setter to explicitly configure IPv6

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L811)  

```go
func NetworkOptionEnableIPv6(enableIPv6 bool) NetworkOption
```

##### NetworkOptionGeneric

NetworkOptionGeneric function returns an option setter for a Generic option defined
in a Dictionary of Key-Value pair

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L764)  

```go
func NetworkOptionGeneric(generic map[string]any) NetworkOption
```

##### NetworkOptionIngress

NetworkOptionIngress returns an option setter to indicate if a network is
an ingress network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L786)  

```go
func NetworkOptionIngress(ingress bool) NetworkOption
```

##### NetworkOptionInternalNetwork

NetworkOptionInternalNetwork returns an option setter to config the network
to be internal which disables default gateway service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L823)  

```go
func NetworkOptionInternalNetwork() NetworkOption
```

##### NetworkOptionIpam

NetworkOptionIpam function returns an option setter for the ipam configuration for this network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L849)  

```go
func NetworkOptionIpam(ipamDriver string, addrSpace string, ipV4 []*IpamConf, ipV6 []*IpamConf, opts map[string]string) NetworkOption
```

##### NetworkOptionLBEndpoint

NetworkOptionLBEndpoint function returns an option setter for the configuration of the load balancer endpoint for this network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L865)  

```go
func NetworkOptionLBEndpoint(ip net.IP) NetworkOption
```

##### NetworkOptionLabels

NetworkOptionLabels function returns an option setter for labels specific to a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L886)  

```go
func NetworkOptionLabels(labels map[string]string) NetworkOption
```

##### NetworkOptionPersist

NetworkOptionPersist returns an option setter to set persistence policy for a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L793)  

```go
func NetworkOptionPersist(persist bool) NetworkOption
```

##### NetworkOptionScope

NetworkOptionScope returns an option setter to overwrite the network's scope.
By default the network's scope is set to the network driver's datascope.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/network.go#L842)  

```go
func NetworkOptionScope(scope string) NetworkOption
```

---

### NetworkWalker

NetworkWalker is a client provided function which will be used to walk the Networks.
When the function returns true, the walk will stop.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/controller.go#L82)  

```go
type NetworkWalker func(nw *Network) bool
```

---

### PortConfig

PortConfig specifies an exposed port which can be
addressed using the given name. This can be later queried
using a service discovery api or a DNS SRV query. The node
port specifies a port that can be used to address this
service external to the cluster by sending a connection
request to this port to any node on the cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L180)  

```go
type PortConfig struct {
	// Name for the port. If provided the port information can
	// be queried using the name as in a DNS SRV query.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Protocol for the port which is exposed.
	Protocol PortConfig_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=libnetwork.PortConfig_Protocol" json:"protocol,omitempty"`
	// The port which the application is exposing and is bound to.
	TargetPort uint32 `protobuf:"varint,3,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	// PublishedPort specifies the port on which the service is
	// exposed on all nodes on the cluster. If not specified an
	// arbitrary port in the node port range is allocated by the
	// system. If specified it should be within the node port
	// range and it should be available.
	PublishedPort uint32 `protobuf:"varint,4,opt,name=published_port,json=publishedPort,proto3" json:"published_port,omitempty"`
}
```

#### Methods

##### PortConfig.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L198)  

```go
func (*PortConfig) Descriptor() ([]byte, []int)
```

##### PortConfig.GetName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L228)  

```go
func (m *PortConfig) GetName() string
```

##### PortConfig.GetProtocol

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L235)  

```go
func (m *PortConfig) GetProtocol() PortConfig_Protocol
```

##### PortConfig.GetPublishedPort

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L249)  

```go
func (m *PortConfig) GetPublishedPort() uint32
```

##### PortConfig.GetTargetPort

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L242)  

```go
func (m *PortConfig) GetTargetPort() uint32
```

##### PortConfig.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L319)  

```go
func (this *PortConfig) GoString() string
```

##### PortConfig.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L440)  

```go
func (m *PortConfig) Marshal() (dAtA []byte, err error)
```

##### PortConfig.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L450)  

```go
func (m *PortConfig) MarshalTo(dAtA []byte) (int, error)
```

##### PortConfig.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L455)  

```go
func (m *PortConfig) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### PortConfig.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L197)  

```go
func (*PortConfig) ProtoMessage()
```

##### PortConfig.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L196)  

```go
func (m *PortConfig) Reset()
```

##### PortConfig.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L546)  

```go
func (m *PortConfig) Size() (n int)
```

##### PortConfig.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L597)  

```go
func (this *PortConfig) String() string
```

##### PortConfig.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L946)  

```go
func (m *PortConfig) Unmarshal(dAtA []byte) error
```

##### PortConfig.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L222)  

```go
func (m *PortConfig) XXX_DiscardUnknown()
```

##### PortConfig.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L204)  

```go
func (m *PortConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### PortConfig.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L216)  

```go
func (m *PortConfig) XXX_Merge(src proto.Message)
```

##### PortConfig.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L219)  

```go
func (m *PortConfig) XXX_Size() int
```

##### PortConfig.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L201)  

```go
func (m *PortConfig) XXX_Unmarshal(b []byte) error
```

---

### PortConfig_Protocol

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L28)  

```go
type PortConfig_Protocol int32
```

#### Methods

##### PortConfig_Protocol.EnumDescriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L52)  

```go
func (PortConfig_Protocol) EnumDescriptor() ([]byte, []int)
```

##### PortConfig_Protocol.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.pb.go#L48)  

```go
func (x PortConfig_Protocol) String() string
```

---

### Resolver

Resolver is the embedded DNS server in Docker. It operates by listening on
the container's loopback interface for DNS queries.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L79)  

```go
type Resolver struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewResolver

NewResolver creates a new instance of the Resolver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L98)  

```go
func NewResolver(address string, proxyDNS bool, backend DNSBackend) *Resolver
```

#### Methods

##### Resolver.NameServer

NameServer returns the IP of the DNS resolver for the containers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L246)  

```go
func (r *Resolver) NameServer() netip.Addr
```

##### Resolver.ResolverOptions

ResolverOptions returns resolv.conf options that should be set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L251)  

```go
func (r *Resolver) ResolverOptions() []string
```

##### Resolver.SetExtServers

SetExtServers configures the external nameservers the resolver should use
when forwarding queries, unless SetExtServersForSrc has configured servers
for the DNS client making the request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L226)  

```go
func (r *Resolver) SetExtServers(extDNS []extDNSEntry)
```

##### Resolver.SetExtServersForSrc

SetExtServersForSrc configures the external nameservers the resolver should
use when forwarding queries from srcAddr. If set, these servers will be used
in preference to servers set by SetExtServers. Supplying a nil or empty extDNS
deletes nameservers for srcAddr.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L240)  

```go
func (r *Resolver) SetExtServersForSrc(srcAddr netip.Addr, extDNS []extDNSEntry) error
```

##### Resolver.SetForwardingPolicy

SetForwardingPolicy re-configures the embedded DNS resolver to either enable or disable forwarding DNS queries to
external servers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L232)  

```go
func (r *Resolver) SetForwardingPolicy(policy bool)
```

##### Resolver.SetupFunc

SetupFunc returns the setup function that should be run in the container's
network namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L148)  

```go
func (r *Resolver) SetupFunc(port uint16) func()
```

##### Resolver.Start

Start starts the name server for the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L174)  

```go
func (r *Resolver) Start() error
```

##### Resolver.Stop

Stop stops the name server for the container. A stopped resolver can be
reused after running the SetupFunc again.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolver.go#L207)  

```go
func (r *Resolver) Stop()
```

---

### Sandbox

Sandbox provides the control over the network container entity.
It is a one to one mapping with the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L40)  

```go
type Sandbox struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Sandbox.AddHostsEntry

AddHostsEntry adds an entry to /etc/hosts.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_dns_unix.go#L32)  

```go
func (sb *Sandbox) AddHostsEntry(ctx context.Context, name, ip string) error
```

##### Sandbox.ContainerID

ContainerID returns the container id associated to this sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L106)  

```go
func (sb *Sandbox) ContainerID() string
```

##### Sandbox.Delete

Delete destroys this container after detaching it from all connected endpoints.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L130)  

```go
func (sb *Sandbox) Delete(ctx context.Context) error
```

##### Sandbox.DisableService

DisableService removes a managed container's endpoints from the load balancer
and service discovery.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L544)  

```go
func (sb *Sandbox) DisableService() (err error)
```

##### Sandbox.EnableService

EnableService makes a managed container's service available by adding the
endpoint to the service load balancer and service discovery.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L521)  

```go
func (sb *Sandbox) EnableService() (err error)
```

##### Sandbox.Endpoints

Endpoints returns all the endpoints connected to the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L303)  

```go
func (sb *Sandbox) Endpoints() []*Endpoint
```

##### Sandbox.ExecFunc

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_linux.go#L135)  

```go
func (sb *Sandbox) ExecFunc(f func()) error
```

##### Sandbox.GetEndpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L343)  

```go
func (sb *Sandbox) GetEndpoint(id string) *Endpoint
```

##### Sandbox.HandleQueryResp

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L356)  

```go
func (sb *Sandbox) HandleQueryResp(name string, ip net.IP)
```

##### Sandbox.ID

ID returns the ID of the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L101)  

```go
func (sb *Sandbox) ID() string
```

##### Sandbox.IPv6Enabled

IPv6Enabled determines whether a container supports IPv6.
IPv6 support can always be determined for host networking. For other network
types it can only be determined once there's a container namespace to probe,
return ok=false in that case.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_linux.go#L226)  

```go
func (sb *Sandbox) IPv6Enabled() (enabled, ok bool)
```

##### Sandbox.Key

Key returns the sandbox's key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L111)  

```go
func (sb *Sandbox) Key() string
```

##### Sandbox.Labels

Labels returns the sandbox's labels.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L119)  

```go
func (sb *Sandbox) Labels() map[string]any
```

##### Sandbox.MarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L282)  

```go
func (sb *Sandbox) MarshalJSON() ([]byte, error)
```

##### Sandbox.NdotsSet

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L677)  

```go
func (sb *Sandbox) NdotsSet() bool
```

##### Sandbox.NetnsPath

NetnsPath returns the network namespace's path and true, if a network has been
created - else the empty string and false.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_linux.go#L212)  

```go
func (sb *Sandbox) NetnsPath() (path string, ok bool)
```

##### Sandbox.Refresh

Refresh leaves all the endpoints, resets and re-applies the options,
re-joins all the endpoints without destroying the osl sandbox

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L243)  

```go
func (sb *Sandbox) Refresh(ctx context.Context, options ...SandboxOption) error
```

##### Sandbox.Rename

Rename changes the name of all attached Endpoints.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L215)  

```go
func (sb *Sandbox) Rename(name string) error
```

##### Sandbox.ResolveIP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L363)  

```go
func (sb *Sandbox) ResolveIP(ctx context.Context, ip string) string
```

##### Sandbox.ResolveName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L401)  

```go
func (sb *Sandbox) ResolveName(ctx context.Context, name string, ipType int) ([]net.IP, bool)
```

##### Sandbox.ResolveService

ResolveService returns all the backend details about the containers or hosts
backing a service. Its purpose is to satisfy an SRV query.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L380)  

```go
func (sb *Sandbox) ResolveService(ctx context.Context, name string) ([]*net.SRV, []net.IP)
```

##### Sandbox.SetKey

SetKey updates the Sandbox Key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_linux.go#L146)  

```go
func (sb *Sandbox) SetKey(ctx context.Context, basePath string) error
```

##### Sandbox.Statistics

Statistics retrieves the interfaces' statistics for the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_linux.go#L57)  

```go
func (sb *Sandbox) Statistics() (map[string]*types.InterfaceStatistics, error)
```

##### Sandbox.UnmarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L290)  

```go
func (sb *Sandbox) UnmarshalJSON(b []byte) (err error)
```

##### Sandbox.UpdateHostsEntry

UpdateHostsEntry updates the IP address in a /etc/hosts entry where the
name matches the regular expression regexp.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_dns_unix.go#L39)  

```go
func (sb *Sandbox) UpdateHostsEntry(regexp, ip string) error
```

##### Sandbox.UpdateLabels

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L273)  

```go
func (sb *Sandbox) UpdateLabels(labels map[string]any)
```

---

### SandboxOption

SandboxOption is an option setter function type used to pass various options to
NewNetContainer method. The various setter functions of type SandboxOption are
provided by libnetwork, they look like ContainerOptionXXXX(...)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox.go#L28)  

```go
type SandboxOption func(sb *Sandbox)
```

#### Functions

##### OptionDNS

OptionDNS function returns an option setter for dns entry option to
be passed to container Create method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L67)  

```go
func OptionDNS(dns []string) SandboxOption
```

##### OptionDNSOptions

OptionDNSOptions function returns an option setter for dns options entry option to
be passed to container Create method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L83)  

```go
func OptionDNSOptions(options []string) SandboxOption
```

##### OptionDNSSearch

OptionDNSSearch function returns an option setter for dns search entry option to
be passed to container Create method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L75)  

```go
func OptionDNSSearch(search []string) SandboxOption
```

##### OptionDomainname

OptionDomainname function returns an option setter for domainname option to
be passed to NewSandbox method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L19)  

```go
func OptionDomainname(name string) SandboxOption
```

##### OptionExposedPorts

OptionExposedPorts function returns an option setter for the container exposed
ports option to be passed to container Create method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L107)  

```go
func OptionExposedPorts(exposedPorts []types.TransportPort) SandboxOption
```

##### OptionExtraHost

OptionExtraHost function returns an option setter for extra /etc/hosts options
which is a name and IP as strings.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L43)  

```go
func OptionExtraHost(name string, IP string) SandboxOption
```

##### OptionHostname

OptionHostname function returns an option setter for hostname option to
be passed to NewSandbox method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L11)  

```go
func OptionHostname(name string) SandboxOption
```

##### OptionHostsPath

OptionHostsPath function returns an option setter for hostspath option to
be passed to NewSandbox method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L27)  

```go
func OptionHostsPath(path string) SandboxOption
```

##### OptionIngress

OptionIngress function returns an option setter for marking a
sandbox as the controller's ingress sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L137)  

```go
func OptionIngress() SandboxOption
```

##### OptionLoadBalancer

OptionLoadBalancer function returns an option setter for marking a
sandbox as a load balancer sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L146)  

```go
func OptionLoadBalancer(nid string) SandboxOption
```

##### OptionOriginHostsPath

OptionOriginHostsPath function returns an option setter for origin hosts file path
to be passed to NewSandbox method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L35)  

```go
func OptionOriginHostsPath(path string) SandboxOption
```

##### OptionOriginResolvConfPath

OptionOriginResolvConfPath function returns an option setter to set the path to the
origin resolv.conf file to be passed to net container methods.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L59)  

```go
func OptionOriginResolvConfPath(path string) SandboxOption
```

##### OptionPortMapping

OptionPortMapping function returns an option setter for the mapping
ports option to be passed to container Create method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L123)  

```go
func OptionPortMapping(portBindings []types.PortBinding) SandboxOption
```

##### OptionResolvConfPath

OptionResolvConfPath function returns an option setter for resolvconfpath option to
be passed to net container methods.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L51)  

```go
func OptionResolvConfPath(path string) SandboxOption
```

##### OptionUseDefaultSandbox

OptionUseDefaultSandbox function returns an option setter for using default sandbox
(host namespace) to be passed to container Create method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L91)  

```go
func OptionUseDefaultSandbox() SandboxOption
```

##### OptionUseExternalKey

OptionUseExternalKey function returns an option setter for using provided namespace
instead of creating one.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/sandbox_options.go#L99)  

```go
func OptionUseExternalKey() SandboxOption
```

---

### ServiceInfo

ServiceInfo has service specific details along with the list of backend tasks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.go#L448)  

```go
type ServiceInfo struct {
	VIP          string
	LocalLBIndex int
	Tasks        []Task
	Ports        []string
}
```

---

### Task

Task has the backend container details

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/agent.go#L440)  

```go
type Task struct {
	Name       string
	EndpointID string
	EndpointIP string
	Info       map[string]string
}
```

---


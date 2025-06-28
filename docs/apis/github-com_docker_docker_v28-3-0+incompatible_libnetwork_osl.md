# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/osl

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:56 UTC

## Overview

Package osl describes structures and interfaces which abstract os entities


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L34)

```go
const (
	// AdvertiseAddrNMsgsMin defines the minimum number of ARP/NA messages sent when an
	// interface is configured.
	// Zero can be used to disable unsolicited ARP/NA.
	AdvertiseAddrNMsgsMin = 0
	// AdvertiseAddrNMsgsMax defines the maximum number of ARP/NA messages sent when an
	// interface is configured. It's three, to match RFC-5227 Section 1.1
	//	// ("PROBE_NUM=3") and RFC-4861 MAX_NEIGHBOR_ADVERTISEMENT.
	AdvertiseAddrNMsgsMax = 3

	// AdvertiseAddrIntervalMin defines the minimum interval between ARP/NA messages
	// sent when an interface is configured. The min defined here is nonstandard,
	// RFC-5227 PROBE_MIN and the default for RetransTimer in RFC-4861 are one
	// second. But, faster resends may be useful in a bridge network (where packets
	// are not transmitted on a real network).
	AdvertiseAddrIntervalMin = 100 * time.Millisecond
	// AdvertiseAddrIntervalMax defines the maximum interval between ARP/NA messages
	// sent when an interface is configured. The max of 2s matches RFC-5227
	// PROBE_MAX.
	AdvertiseAddrIntervalMax = 2 * time.Second
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/sandbox.go#L7)

```go
const (
	// SandboxTypeIngress indicates that the sandbox is for the ingress
	SandboxTypeIngress = iota
	// SandboxTypeLoadBalancer indicates that the sandbox is a load balancer
	SandboxTypeLoadBalancer = iota
)
```

## Variables

This section is empty.

## Functions

### GenerateKey

GenerateKey generates a sandbox key based on the passed
container id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L63)  

```go
func GenerateKey(containerID string) string
```

---

### SetBasePath

SetBasePath sets the base url prefix for the ns path

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L46)  

```go
func SetBasePath(path string)
```

---

## Types

### Iface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/sandbox.go#L14)  

```go
type Iface struct {
	SrcName, DstPrefix, DstName string
}
```

---

### IfaceOption

IfaceOption is a function option type to set interface options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/sandbox.go#L19)  

```go
type IfaceOption func(i *Interface) error
```

#### Functions

##### WithAdvertiseAddrInterval

WithAdvertiseAddrInterval sets the interval between unsolicited ARP/NA messages
sent to advertise a network interface's addresses.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L112)  

```go
func WithAdvertiseAddrInterval(interval time.Duration) IfaceOption
```

##### WithAdvertiseAddrNMsgs

WithAdvertiseAddrNMsgs sets the number of unsolicited ARP/NA messages that will
be sent to advertise a network interface's addresses.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L99)  

```go
func WithAdvertiseAddrNMsgs(nMsgs int) IfaceOption
```

##### WithCreatedInContainer

WithCreatedInContainer can be used to say the network driver created the
interface in the container's network namespace (and, therefore, it doesn't
need to be moved into that namespace.)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L126)  

```go
func WithCreatedInContainer(cic bool) IfaceOption
```

##### WithIPv4Address

WithIPv4Address sets the IPv4 address of the interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L58)  

```go
func WithIPv4Address(addr *net.IPNet) IfaceOption
```

##### WithIPv6Address

WithIPv6Address sets the IPv6 address of the interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L66)  

```go
func WithIPv6Address(addr *net.IPNet) IfaceOption
```

##### WithIsBridge

WithIsBridge sets whether the interface is a bridge.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L32)  

```go
func WithIsBridge(isBridge bool) IfaceOption
```

##### WithLinkLocalAddresses

WithLinkLocalAddresses set the link-local IP addresses of the interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L74)  

```go
func WithLinkLocalAddresses(list []*net.IPNet) IfaceOption
```

##### WithMACAddress

WithMACAddress sets the interface MAC-address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L50)  

```go
func WithMACAddress(mac net.HardwareAddr) IfaceOption
```

##### WithMaster

WithMaster sets the master interface (if any) for this interface. The
master interface name should refer to the srcName of a previously added
interface of type bridge.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L42)  

```go
func WithMaster(name string) IfaceOption
```

##### WithRoutes

WithRoutes sets the interface routes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L82)  

```go
func WithRoutes(routes []*net.IPNet) IfaceOption
```

##### WithSysctls

WithSysctls sets the interface sysctls.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L90)  

```go
func WithSysctls(sysctls []string) IfaceOption
```

---

### Interface

Interface represents the settings and identity of a network device.
It is used as a return type for Network.Link, and it is common practice
for the caller to use this information when moving interface SrcName from
host namespace to DstName in a different net namespace with the appropriate
network settings.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L97)  

```go
type Interface struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Interface.Address

Address returns the IPv4 address for the interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L146)  

```go
func (i *Interface) Address() *net.IPNet
```

##### Interface.AddressIPv6

AddressIPv6 returns the IPv6 address for the interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L151)  

```go
func (i *Interface) AddressIPv6() *net.IPNet
```

##### Interface.Bridge

Bridge returns true if the interface is a bridge.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L137)  

```go
func (i *Interface) Bridge() bool
```

##### Interface.DstMaster

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L132)  

```go
func (i *Interface) DstMaster() string
```

##### Interface.DstName

DstName returns the final interface name in the target network namespace.
It's generated based on the prefix passed to Namespace.AddInterface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L128)  

```go
func (i *Interface) DstName() string
```

##### Interface.LinkLocalAddresses

LinkLocalAddresses returns the link-local IP addresses assigned to the
interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L157)  

```go
func (i *Interface) LinkLocalAddresses() []*net.IPNet
```

##### Interface.MacAddress

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L141)  

```go
func (i *Interface) MacAddress() net.HardwareAddr
```

##### Interface.Remove

Remove an interface from the sandbox by renaming to original name
and moving it out of the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L173)  

```go
func (i *Interface) Remove() error
```

##### Interface.Routes

Routes returns IP routes for the interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L162)  

```go
func (i *Interface) Routes() []*net.IPNet
```

##### Interface.SrcName

SrcName returns the name of the interface in the origin network namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L122)  

```go
func (i *Interface) SrcName() string
```

##### Interface.Statistics

Statistics returns the sandbox's side veth interface statistics.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L179)  

```go
func (i *Interface) Statistics() (*types.InterfaceStatistics, error)
```

---

### Namespace

Namespace represents a network sandbox. It represents a Linux network
namespace, and moves an interface into it when called on method AddInterface
or sets the gateway etc. It holds a list of Interfaces, routes etc., and more
can be added dynamically.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L227)  

```go
type Namespace struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### GetSandboxForExternalKey

GetSandboxForExternalKey returns sandbox object for the supplied path

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L151)  

```go
func GetSandboxForExternalKey(basePath string, key string) (*Namespace, error)
```

##### NewSandbox

NewSandbox provides a new Namespace instance created in an os specific way
provided a key which uniquely identifies the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L106)  

```go
func NewSandbox(key string, osCreate, isRestore bool) (*Namespace, error)
```

#### Methods

##### Namespace.AddAliasIP

AddAliasIP adds the passed IP address to the named interface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L278)  

```go
func (n *Namespace) AddAliasIP(ifName string, ip *net.IPNet) error
```

##### Namespace.AddInterface

AddInterface creates an Interface that represents an existing network
interface (except for bridge interfaces, which are created here).

The network interface will be reconfigured according the options passed, and
it'll be renamed from srcName into either dstName if it's not empty, or to
an auto-generated dest name that combines the provided dstPrefix and a
numeric suffix.

It's safe to call concurrently.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L235)  

```go
func (n *Namespace) AddInterface(ctx context.Context, srcName, dstPrefix, dstName string, options ...IfaceOption) error
```

##### Namespace.AddNeighbor

AddNeighbor adds a neighbor entry into the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/neigh_linux.go#L113)  

```go
func (n *Namespace) AddNeighbor(dstIP net.IP, dstMac net.HardwareAddr, force bool, options ...NeighOption) error
```

##### Namespace.AddStaticRoute

AddStaticRoute adds a static route to the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L179)  

```go
func (n *Namespace) AddStaticRoute(r *types.StaticRoute) error
```

##### Namespace.ApplyOSTweaks

ApplyOSTweaks applies operating system specific knobs on the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L526)  

```go
func (n *Namespace) ApplyOSTweaks(types []SandboxType)
```

##### Namespace.DeleteNeighbor

DeleteNeighbor deletes neighbor entry from the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/neigh_linux.go#L48)  

```go
func (n *Namespace) DeleteNeighbor(dstIP net.IP, dstMac net.HardwareAddr) error
```

##### Namespace.Destroy

Destroy destroys the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L385)  

```go
func (n *Namespace) Destroy() error
```

##### Namespace.DisableARPForVIP

DisableARPForVIP disables ARP replies and requests for VIP addresses
on a particular interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L297)  

```go
func (n *Namespace) DisableARPForVIP(srcName string) (retErr error)
```

##### Namespace.Gateway

Gateway returns the IPv4 gateway for the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L16)  

```go
func (n *Namespace) Gateway() net.IP
```

##### Namespace.GatewayIPv6

GatewayIPv6 returns the IPv6 gateway for the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L24)  

```go
func (n *Namespace) GatewayIPv6() net.IP
```

##### Namespace.GetLoopbackIfaceName

GetLoopbackIfaceName returns the name of the loopback interface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L273)  

```go
func (n *Namespace) GetLoopbackIfaceName() string
```

##### Namespace.IPv6LoEnabled

IPv6LoEnabled returns true if the loopback interface had an IPv6 address when
last checked. It's always checked on the first call, and by RefreshIPv6LoEnabled.
('::1' is assigned by the kernel if IPv6 is enabled.)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L496)  

```go
func (n *Namespace) IPv6LoEnabled() bool
```

##### Namespace.Interfaces

Interfaces returns the collection of Interface previously added with the AddInterface
method. Note that this doesn't include network interfaces added in any
other way (such as the default loopback interface which is automatically
created on creation of a sandbox).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L247)  

```go
func (n *Namespace) Interfaces() []*Interface
```

##### Namespace.InvokeFunc

InvokeFunc invoke a function in the network namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L328)  

```go
func (n *Namespace) InvokeFunc(f func()) error
```

##### Namespace.Key

Key returns the path where the network namespace is mounted.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L380)  

```go
func (n *Namespace) Key() string
```

##### Namespace.RefreshIPv6LoEnabled

RefreshIPv6LoEnabled refreshes the cached result returned by IPv6LoEnabled.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L506)  

```go
func (n *Namespace) RefreshIPv6LoEnabled()
```

##### Namespace.RemoveAliasIP

RemoveAliasIP removes the passed IP address from the named interface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L287)  

```go
func (n *Namespace) RemoveAliasIP(ifName string, ip *net.IPNet) error
```

##### Namespace.RemoveInterface

RemoveInterface removes an interface from the namespace by renaming to
original name and moving it out of the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/interface_linux.go#L805)  

```go
func (n *Namespace) RemoveInterface(i *Interface) error
```

##### Namespace.RemoveStaticRoute

RemoveStaticRoute removes a static route from the sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L191)  

```go
func (n *Namespace) RemoveStaticRoute(r *types.StaticRoute) error
```

##### Namespace.RestoreGateway

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L472)  

```go
func (n *Namespace) RestoreGateway(ipv4 bool, gw net.IP, srcName string)
```

##### Namespace.RestoreInterfaces

RestoreInterfaces restores the network namespace's interfaces.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L401)  

```go
func (n *Namespace) RestoreInterfaces(interfaces map[Iface][]IfaceOption) error
```

##### Namespace.RestoreRoutes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/namespace_linux.go#L466)  

```go
func (n *Namespace) RestoreRoutes(routes []*types.StaticRoute)
```

##### Namespace.SetDefaultRouteIPv4

SetDefaultRouteIPv4 sets up a connected route to 0.0.0.0 via the Interface
with srcName, if that Interface has a route to 0.0.0.0. Otherwise, it
returns an error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L214)  

```go
func (n *Namespace) SetDefaultRouteIPv4(srcName string) error
```

##### Namespace.SetDefaultRouteIPv6

SetDefaultRouteIPv6 sets up a connected route to [::] via the Interface
with srcName, if that Interface has a route to [::]. Otherwise, it
returns an error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L230)  

```go
func (n *Namespace) SetDefaultRouteIPv6(srcName string) error
```

##### Namespace.SetGateway

SetGateway sets the default IPv4 gateway for the sandbox. It is a no-op
if the given gateway is empty.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L49)  

```go
func (n *Namespace) SetGateway(gw net.IP) error
```

##### Namespace.SetGatewayIPv6

SetGatewayIPv6 sets the default IPv6 gateway for the sandbox. It is a no-op
if the given gateway is empty.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L145)  

```go
func (n *Namespace) SetGatewayIPv6(gwv6 net.IP) error
```

##### Namespace.StaticRoutes

StaticRoutes returns additional static routes for the sandbox. Note that
directly connected routes are stored on the particular interface they
refer to.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L34)  

```go
func (n *Namespace) StaticRoutes() []*types.StaticRoute
```

##### Namespace.UnsetDefaultRouteIPv4

UnsetDefaultRouteIPv4 unsets the previously set default IPv4 default route
in the sandbox. It is a no-op if no gateway was set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L271)  

```go
func (n *Namespace) UnsetDefaultRouteIPv4() error
```

##### Namespace.UnsetDefaultRouteIPv6

UnsetDefaultRouteIPv6 unsets the previously set default IPv6 default route
in the sandbox. It is a no-op if no gateway was set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L290)  

```go
func (n *Namespace) UnsetDefaultRouteIPv6() error
```

##### Namespace.UnsetGateway

UnsetGateway the previously set default IPv4 gateway in the sandbox.
It is a no-op if no gateway was set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L65)  

```go
func (n *Namespace) UnsetGateway() error
```

##### Namespace.UnsetGatewayIPv6

UnsetGatewayIPv6 unsets the previously set default IPv6 gateway in the sandbox.
It is a no-op if no gateway was set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/route_linux.go#L162)  

```go
func (n *Namespace) UnsetGatewayIPv6() error
```

---

### NeighOption

NeighOption is a function option type to set neighbor options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/sandbox.go#L22)  

```go
type NeighOption func(nh *neigh)
```

#### Functions

##### WithFamily

WithFamily sets the address-family for the neighbor entry. e.g. syscall.AF_BRIDGE.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L25)  

```go
func WithFamily(family int) NeighOption
```

##### WithLinkName

WithLinkName sets the srcName of the link to use in the neighbor entry.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/options_linux.go#L18)  

```go
func WithLinkName(name string) NeighOption
```

---

### NeighborSearchError

NeighborSearchError indicates that the neighbor is already present

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/neigh_linux.go#L16)  

```go
type NeighborSearchError struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### NeighborSearchError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/neigh_linux.go#L22)  

```go
func (n NeighborSearchError) Error() string
```

---

### SandboxType

SandboxType specify the time of the sandbox, this can be used to apply special configs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/sandbox.go#L5)  

```go
type SandboxType int
```

---


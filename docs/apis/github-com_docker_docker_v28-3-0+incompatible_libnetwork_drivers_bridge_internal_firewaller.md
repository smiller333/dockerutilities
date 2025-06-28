# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/bridge/internal/firewaller

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:33:14 UTC

## Overview

Package firewaller defines an interface that can be used to manipulate
firewall configuration for a bridge network.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Config

Config contains top-level settings for the firewaller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/firewaller.go#L22)  

```go
type Config struct {
	// IPv4 true means IPv4 firewalling is required.
	IPv4 bool
	// IPv6 true means IPv4 firewalling is required.
	IPv6 bool
	// Hairpin means the userland proxy will not be running.
	Hairpin bool
	// AllowDirectRouting means packets addressed directly to a container's IP address will be
	// accepted, regardless of which network interface they are from.
	AllowDirectRouting bool
	// WSL2Mirrored is true if running under WSL2 with mirrored networking enabled.
	WSL2Mirrored bool
}
```

---

### Firewaller

Firewaller implements firewall rules for bridge networks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/firewaller.go#L71)  

```go
type Firewaller interface {
	// NewNetwork returns an object that can be used to add published ports and legacy
	// links for a bridge network.
	NewNetwork(ctx context.Context, nc NetworkConfig) (Network, error)
	// FilterForwardDrop sets the default policy of the FORWARD chain in the filter
	// table to DROP.
	FilterForwardDrop(ctx context.Context, ipv IPVersion) error
}
```

---

### IPVersion

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/firewaller.go#L14)  

```go
type IPVersion uint8
```

---

### Network

Network can be used to manipulate firewall rules for a bridge network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/firewaller.go#L81)  

```go
type Network interface {
	// ReapplyNetworkLevelRules re-creates the initial set of network-level rules
	// created by [Firewaller.NewNetwork]. It can be called after, for example, a
	// firewalld reload has deleted the rules. Rules for port mappings and legacy
	// links are not re-created.
	ReapplyNetworkLevelRules(ctx context.Context) error
	// DelNetworkLevelRules deletes any configuration set up by [Firewaller.NewNetwork].
	// It does not delete per-port or per-link rules. The caller is responsible for tracking
	// those and deleting them when the network is removed.
	DelNetworkLevelRules(ctx context.Context) error

	// AddEndpoint is used to notify the firewaller about a new container on the
	// network, with its IP addresses.
	AddEndpoint(ctx context.Context, epIPv4, epIPv6 netip.Addr) error
	// DelEndpoint undoes configuration applied by AddEndpoint.
	DelEndpoint(ctx context.Context, epIPv4, epIPv6 netip.Addr) error

	// AddPorts adds the configuration needed for published ports.
	AddPorts(ctx context.Context, pbs []types.PortBinding) error
	// DelPorts deletes the configuration needed for published ports.
	DelPorts(ctx context.Context, pbs []types.PortBinding) error

	// AddLink adds the configuration needed for a legacy link.
	AddLink(ctx context.Context, parentIP, childIP netip.Addr, ports []types.TransportPort) error
	// DelLink deletes the configuration needed for a legacy link.
	DelLink(ctx context.Context, parentIP, childIP netip.Addr, ports []types.TransportPort)
}
```

---

### NetworkConfig

NetworkConfig contains settings for a single bridge network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/firewaller.go#L37)  

```go
type NetworkConfig struct {
	// IfName is the name of the bridge device.
	IfName string
	// Internal is true if the network should have no access to networks outside the Docker host.
	Internal bool
	// ICC is false if containers on the bridge should not be able to communicate (unless it's the
	// default bridge, and legacy-links are set up).
	ICC bool
	// Masquerade is true if the network should use masquerading/SNAT.
	Masquerade bool
	// TrustedHostInterfaces are interfaces that must be treated as part of the network (like the
	// bridge itself). In particular, these are not external interfaces for the purpose of
	// blocking direct-routing to a container's IP address.
	TrustedHostInterfaces []string
	// Config4 contains IPv4-specific configuration for the network.
	Config4 NetworkConfigFam
	// Config6 contains IPv6-specific configuration for the network.
	Config6 NetworkConfigFam
}
```

---

### NetworkConfigFam

NetworkConfigFam contains network configuration for a single address family.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/firewaller.go#L58)  

```go
type NetworkConfigFam struct {
	// HostIP is the address to use for SNAT. If unset, masquerading will be used instead.
	HostIP netip.Addr
	// Prefix is the bridge network's subnet.
	Prefix netip.Prefix
	// Routed is true if containers should be directly addressable, no NAT from the host.
	Routed bool
	// Unprotected is true if no rules to filter unpublished ports or direct access from
	// any remote host are required.
	Unprotected bool
}
```

---

### StubFirewaller

StubFirewaller implements a Firewaller for unit tests. It just tracks what it's been asked for.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L16)  

```go
type StubFirewaller struct {
	Config
	Networks map[string]*StubFirewallerNetwork
	FFD      map[IPVersion]bool // filter forward drop
}
```

#### Functions

##### NewStubFirewaller

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L22)  

```go
func NewStubFirewaller(config Config) *StubFirewaller
```

#### Methods

##### StubFirewaller.FilterForwardDrop

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L45)  

```go
func (fw *StubFirewaller) FilterForwardDrop(_ context.Context, ipv IPVersion) error
```

##### StubFirewaller.NewNetwork

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L32)  

```go
func (fw *StubFirewaller) NewNetwork(_ context.Context, nc NetworkConfig) (Network, error)
```

---

### StubFirewallerNetwork

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L61)  

```go
type StubFirewallerNetwork struct {
	NetworkConfig
	Deleted   bool
	Endpoints map[stubEndpoint]struct{}
	Ports     []types.PortBinding
	Links     []stubFirewallerLink
	// contains filtered or unexported fields
}
```

#### Methods

##### StubFirewallerNetwork.AddEndpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L97)  

```go
func (nw *StubFirewallerNetwork) AddEndpoint(_ context.Context, epIPv4, epIPv6 netip.Addr) error
```

##### StubFirewallerNetwork.AddLink

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L134)  

```go
func (nw *StubFirewallerNetwork) AddLink(_ context.Context, parentIP, childIP netip.Addr, ports []types.TransportPort) error
```

##### StubFirewallerNetwork.AddPorts

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L115)  

```go
func (nw *StubFirewallerNetwork) AddPorts(_ context.Context, pbs []types.PortBinding) error
```

##### StubFirewallerNetwork.DelEndpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L106)  

```go
func (nw *StubFirewallerNetwork) DelEndpoint(_ context.Context, epIPv4, epIPv6 netip.Addr) error
```

##### StubFirewallerNetwork.DelLink

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L152)  

```go
func (nw *StubFirewallerNetwork) DelLink(_ context.Context, parentIP, childIP netip.Addr, ports []types.TransportPort)
```

##### StubFirewallerNetwork.DelNetworkLevelRules

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L75)  

```go
func (nw *StubFirewallerNetwork) DelNetworkLevelRules(_ context.Context) error
```

##### StubFirewallerNetwork.DelPorts

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L125)  

```go
func (nw *StubFirewallerNetwork) DelPorts(_ context.Context, pbs []types.PortBinding) error
```

##### StubFirewallerNetwork.LinkExists

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L164)  

```go
func (nw *StubFirewallerNetwork) LinkExists(parentIP, childIP netip.Addr, ports []types.TransportPort) bool
```

##### StubFirewallerNetwork.PortExists

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L158)  

```go
func (nw *StubFirewallerNetwork) PortExists(pb types.PortBinding) bool
```

##### StubFirewallerNetwork.ReapplyNetworkLevelRules

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/firewaller/stub.go#L71)  

```go
func (nw *StubFirewallerNetwork) ReapplyNetworkLevelRules(_ context.Context) error
```

---


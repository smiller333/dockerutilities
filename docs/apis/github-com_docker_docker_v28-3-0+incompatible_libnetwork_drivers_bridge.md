# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/bridge

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:33:09 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/bridge_linux.go#L47)

```go
const (
	// DefaultGatewayV4AuxKey represents the default-gateway configured by the user
	DefaultGatewayV4AuxKey = "DefaultGatewayIPv4"
	// DefaultGatewayV6AuxKey represents the ipv6 default-gateway configured by the user
	DefaultGatewayV6AuxKey = "DefaultGatewayIPv6"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/labels.go#L3)

```go
const (
	// BridgeName label for bridge driver
	BridgeName = "com.docker.network.bridge.name"

	// EnableIPMasquerade label for bridge driver
	EnableIPMasquerade = "com.docker.network.bridge.enable_ip_masquerade"

	// IPv4GatewayMode label for bridge driver
	IPv4GatewayMode = "com.docker.network.bridge.gateway_mode_ipv4"
	// IPv6GatewayMode label for bridge driver
	IPv6GatewayMode = "com.docker.network.bridge.gateway_mode_ipv6"

	// EnableICC label
	EnableICC = "com.docker.network.bridge.enable_icc"

	// InhibitIPv4 label
	InhibitIPv4 = "com.docker.network.bridge.inhibit_ipv4"

	// DefaultBindingIP label
	DefaultBindingIP = "com.docker.network.bridge.host_binding_ipv4"

	// DefaultBridge label
	DefaultBridge = "com.docker.network.bridge.default_bridge"

	// TrustedHostInterfaces can be used to supply a list of host interfaces that are
	// allowed direct access to published ports on a container's address.
	TrustedHostInterfaces = "com.docker.network.bridge.trusted_host_interfaces"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/interface_linux.go#L17)

```go
const (
	// DefaultBridgeName is the default name for the bridge interface managed
	// by the driver when unspecified by the caller.
	DefaultBridgeName = "docker0"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/bridge_linux.go#L59)

```go
const DockerForwardChain = iptabler.DockerForwardChain
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/bridge_linux.go#L39)

```go
const (
	NetworkType = "bridge"
)
```

## Variables

This section is empty.

## Functions

### LegacyContainerLinkOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/bridge_linux.go#L1694)  

```go
func LegacyContainerLinkOptions(parentEndpoints, childEndpoints []string) map[string]interface{}
```

---

### Register

Register registers a new instance of bridge driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/bridge_linux.go#L190)  

```go
func Register(r driverapi.Registerer, store *datastore.Store, config map[string]interface{}) error
```

---

### ValidateFixedCIDRV6

ValidateFixedCIDRV6 checks that val is an IPv6 address and prefix length that
does not overlap with the link local subnet prefix 'fe80::/64'.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/bridge_linux.go#L231)  

```go
func ValidateFixedCIDRV6(val string) error
```

---

## Types

This section is empty.


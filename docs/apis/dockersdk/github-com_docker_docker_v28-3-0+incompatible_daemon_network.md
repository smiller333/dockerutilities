# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/network

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:40 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network/network_mode.go#L7)

```go
const DefaultNetwork = defaultNetwork
```

## Variables

This section is empty.

## Functions

### FilterNetworks

FilterNetworks filters network list according to user specified filter
and returns user chosen networks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network/filter.go#L12)  
**Added in:** v1.12.0

```go
func FilterNetworks(nws []network.Inspect, filter filters.Args) ([]network.Inspect, error)
```

---

### IsPredefined

IsPredefined indicates if a network is predefined by the daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network/network_mode.go#L10)  

```go
func IsPredefined(network string) bool
```

---

## Types

### AttachmentStore

AttachmentStore stores the load balancer IP address for a network id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network/settings.go#L42)  

```go
type AttachmentStore struct {
	sync.Mutex
	// contains filtered or unexported fields
}
```

#### Methods

##### AttachmentStore.ClearAttachments

ClearAttachments clears all the mappings of network to load balancer IP Address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network/settings.go#L67)  

```go
func (store *AttachmentStore) ClearAttachments()
```

##### AttachmentStore.GetIPForNetwork

GetIPForNetwork return the load balancer IP address for the given network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network/settings.go#L78)  

```go
func (store *AttachmentStore) GetIPForNetwork(networkID string) (net.IP, bool)
```

##### AttachmentStore.ResetAttachments

ResetAttachments clears any existing load balancer IP to network mapping and
sets the mapping to the given attachments.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network/settings.go#L51)  

```go
func (store *AttachmentStore) ResetAttachments(attachments map[string]string) error
```

---

### EndpointSettings

EndpointSettings is a package local wrapper for
networktypes.EndpointSettings which stores Endpoint state that
needs to be persisted to disk but not exposed in the api.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network/settings.go#L33)  
**Added in:** v1.9.0

```go
type EndpointSettings struct {
	*networktypes.EndpointSettings
	IPAMOperational bool
	// DesiredMacAddress is the configured value, it's copied from MacAddress (the
	// API param field) when the container is created.
	DesiredMacAddress string
}
```

---

### Settings

Settings stores configuration details about the daemon network config
TODO Windows. Many of these fields can be factored out.,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network/settings.go#L15)  

```go
type Settings struct {
	Bridge                 string
	SandboxID              string
	SandboxKey             string
	HairpinMode            bool
	LinkLocalIPv6Address   string
	LinkLocalIPv6PrefixLen int
	Networks               map[string]*EndpointSettings
	Service                *clustertypes.ServiceConfig
	Ports                  nat.PortMap
	SecondaryIPAddresses   []networktypes.Address
	SecondaryIPv6Addresses []networktypes.Address
	HasSwarmEndpoint       bool
}
```

---


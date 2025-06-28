# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/cluster/provider

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:55 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### NetworkCreateRequest

NetworkCreateRequest is a request when creating a network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/provider/network.go#L6)  

```go
type NetworkCreateRequest struct {
	ID string
	network.CreateRequest
}
```

---

### NetworkCreateResponse

NetworkCreateResponse is a response when creating a network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/provider/network.go#L12)  

```go
type NetworkCreateResponse struct {
	ID string `json:"Id"`
}
```

---

### PortConfig

PortConfig represents a port configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/provider/network.go#L23)  

```go
type PortConfig struct {
	Name          string
	Protocol      int32
	TargetPort    uint32
	PublishedPort uint32
}
```

---

### ServiceConfig

ServiceConfig represents a service configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/provider/network.go#L31)  

```go
type ServiceConfig struct {
	ID               string
	Name             string
	Aliases          map[string][]string
	VirtualAddresses map[string]*VirtualAddress
	ExposedPorts     []*PortConfig
}
```

---

### VirtualAddress

VirtualAddress represents a virtual address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/provider/network.go#L17)  

```go
type VirtualAddress struct {
	IPv4 string
	IPv6 string
}
```

---


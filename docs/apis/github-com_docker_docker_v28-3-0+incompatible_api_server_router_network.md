# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/network

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:25 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new network router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/network/network.go#L15)  

```go
func NewRouter(b Backend, c ClusterBackend) router.Router
```

---

## Types

### Backend

Backend is all the methods that need to be implemented
to provide network specific functionality.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/network/backend.go#L13)  
**Added in:** v1.10.0

```go
type Backend interface {
	GetNetworks(filters.Args, backend.NetworkListConfig) ([]network.Inspect, error)
	CreateNetwork(ctx context.Context, nc network.CreateRequest) (*network.CreateResponse, error)
	ConnectContainerToNetwork(ctx context.Context, containerName, networkName string, endpointConfig *network.EndpointSettings) error
	DisconnectContainerFromNetwork(containerName string, networkName string, force bool) error
	DeleteNetwork(networkID string) error
	NetworksPrune(ctx context.Context, pruneFilters filters.Args) (*network.PruneReport, error)
}
```

---

### ClusterBackend

ClusterBackend is all the methods that need to be implemented
to provide cluster network specific functionality.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/network/backend.go#L24)  

```go
type ClusterBackend interface {
	GetNetworks(filters.Args) ([]network.Inspect, error)
	GetNetwork(name string) (network.Inspect, error)
	GetNetworksByName(name string) ([]network.Inspect, error)
	CreateNetwork(nc network.CreateRequest) (string, error)
	RemoveNetwork(name string) error
}
```

---


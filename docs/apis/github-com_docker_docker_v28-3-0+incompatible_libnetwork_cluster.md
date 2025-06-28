# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/cluster

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:41 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cluster/provider.go#L9)

```go
const (
	// EventSocketChange control socket changed
	EventSocketChange = iota
	// EventNodeReady cluster node in ready state
	EventNodeReady
	// EventNodeLeave node is leaving the cluster
	EventNodeLeave
	// EventNetworkKeysAvailable network keys correctly configured in the networking layer
	EventNetworkKeysAvailable
)
```

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### ConfigEventType

ConfigEventType type of the event produced by the cluster

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cluster/provider.go#L21)  

```go
type ConfigEventType uint8
```

---

### Provider

Provider provides clustering config details

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cluster/provider.go#L24)  

```go
type Provider interface {
	IsManager() bool
	IsAgent() bool
	GetLocalAddress() string
	GetListenAddress() string
	GetAdvertiseAddress() string
	GetDataPathAddress() string
	GetRemoteAddressList() []string
	ListenClusterEvents() <-chan ConfigEventType
	AttachNetwork(string, string, []string) (*network.NetworkingConfig, error)
	DetachNetwork(string, string) error
	UpdateAttachment(string, string, *network.NetworkingConfig) error
	WaitForDetachment(context.Context, string, string, string, string) error
}
```

---


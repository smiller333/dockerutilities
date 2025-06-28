# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/discoverapi

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:33:04 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/discoverapi/discoverapi.go#L16)

```go
const (
	// NodeDiscovery represents Node join/leave events provided by discovery
	NodeDiscovery = iota + 1
	// EncryptionKeysConfig represents the initial key(s) for performing datapath encryption
	EncryptionKeysConfig
	// EncryptionKeysUpdate represents an update to the datapath encryption key(s)
	EncryptionKeysUpdate
)
```

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### DatastoreConfigData

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/discoverapi/discoverapi.go#L35)  

```go
type DatastoreConfigData struct {
	Scope    string
	Provider string
	Address  string
	Config   interface{}
}
```

---

### Discover

Discover is an interface to be implemented by the component interested in receiving discover events
like new node joining the cluster or datastore updates

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/discoverapi/discoverapi.go#L5)  

```go
type Discover interface {
	// DiscoverNew is a notification for a new discovery event, Example:a new node joining a cluster
	DiscoverNew(dType DiscoveryType, data interface{}) error

	// DiscoverDelete is a notification for a discovery delete event, Example:a node leaving a cluster
	DiscoverDelete(dType DiscoveryType, data interface{}) error
}
```

---

### DiscoveryType

DiscoveryType represents the type of discovery element the DiscoverNew function is invoked on

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/discoverapi/discoverapi.go#L14)  

```go
type DiscoveryType int
```

---

### DriverEncryptionConfig

DriverEncryptionConfig contains the initial datapath encryption key(s)
Key in first position is the primary key, the one to be used in tx.
Original key and tag types are []byte and uint64

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/discoverapi/discoverapi.go#L45)  

```go
type DriverEncryptionConfig struct {
	Keys [][]byte
	Tags []uint64
}
```

---

### DriverEncryptionUpdate

DriverEncryptionUpdate carries an update to the encryption key(s) as:
a new key and/or set a primary key and/or a removal of an existing key.
Original key and tag types are []byte and uint64

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/discoverapi/discoverapi.go#L53)  

```go
type DriverEncryptionUpdate struct {
	Key        []byte
	Tag        uint64
	Primary    []byte
	PrimaryTag uint64
	Prune      []byte
	PruneTag   uint64
}
```

---

### NodeDiscoveryData

NodeDiscoveryData represents the structure backing the node discovery data json string

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/discoverapi/discoverapi.go#L26)  

```go
type NodeDiscoveryData struct {
	Address     string
	BindAddress string
	Self        bool
}
```

---


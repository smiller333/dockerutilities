# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/macvlan

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:33:28 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/macvlan/macvlan.go#L16)

```go
const (
	NetworkType = "macvlan" // driver type name

)
```

## Variables

This section is empty.

## Functions

### Register

Register initializes and registers the libnetwork macvlan driver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/macvlan/macvlan.go#L60)  

```go
func Register(r driverapi.Registerer, store *datastore.Store, _ map[string]interface{}) error
```

---

## Types

This section is empty.


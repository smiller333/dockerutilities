# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/ipvlan

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:09:17 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/ipvlan/ipvlan.go#L16)

```go
const (
	NetworkType = "ipvlan" // driver type name

)
```

## Variables

This section is empty.

## Functions

### Register

Register initializes and registers the libnetwork ipvlan driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/ipvlan/ipvlan.go#L66)  

```go
func Register(r driverapi.Registerer, store *datastore.Store, config map[string]interface{}) error
```

---

## Types

This section is empty.


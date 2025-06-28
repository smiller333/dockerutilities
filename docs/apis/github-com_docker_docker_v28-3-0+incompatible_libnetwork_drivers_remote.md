# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/remote

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:09:37 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Register

Register makes sure a remote driver is registered with r when a network
driver plugin is activated.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/driver.go#L39)  

```go
func Register(r driverapi.Registerer, pg plugingetter.PluginGetter) error
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/ipams

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:21 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Register

Register registers all the builtin drivers (ie. default, windowsipam, null
and remote). If 'pg' is nil, the remote driver won't be registered.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/drivers.go#L15)  

```go
func Register(r ipamapi.Registerer, pg plugingetter.PluginGetter, lAddrPools, gAddrPools []*ipamutils.NetworkToSplit) error
```

---

## Types

This section is empty.


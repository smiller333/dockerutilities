# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/ipams/remote

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:28 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Register

Register registers a remote ipam when its plugin is activated.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/remote.go#L35)  

```go
func Register(cb ipamapi.Registerer, pg plugingetter.PluginGetter) error
```

---

## Types

### PluginResponse

PluginResponse is the interface for the plugin request responses

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/remote.go#L24)  

```go
type PluginResponse interface {
	IsSuccess() bool
	GetError() string
}
```

---


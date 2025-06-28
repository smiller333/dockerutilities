# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/plugingetter

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:00 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugingetter/getter.go#L10)

```go
const (
	// Lookup doesn't update RefCount
	Lookup = 0
	// Acquire increments RefCount
	Acquire = 1
	// Release decrements RefCount
	Release = -1
)
```

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### CompatPlugin

CompatPlugin is an abstraction to handle both v2(new) and v1(legacy) plugins.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugingetter/getter.go#L20)  

```go
type CompatPlugin interface {
	Name() string
	ScopedPath(string) string
	IsV1() bool
	PluginWithV1Client
}
```

---

### CountedPlugin

CountedPlugin is a plugin which is reference counted.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugingetter/getter.go#L40)  

```go
type CountedPlugin interface {
	Acquire()
	Release()
	CompatPlugin
}
```

---

### PluginAddr

PluginAddr is a plugin that exposes the socket address for creating custom clients rather than the built-in `*plugins.Client`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugingetter/getter.go#L33)  

```go
type PluginAddr interface {
	Addr() net.Addr
	Timeout() time.Duration
	Protocol() string
}
```

---

### PluginGetter

PluginGetter is the interface implemented by Store

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugingetter/getter.go#L47)  

```go
type PluginGetter interface {
	Get(name, capability string, mode int) (CompatPlugin, error)
	GetAllByCap(capability string) ([]CompatPlugin, error)
	GetAllManagedPluginsByCap(capability string) []CompatPlugin
	Handle(capability string, callback func(string, *plugins.Client))
}
```

---

### PluginWithV1Client

PluginWithV1Client is a plugin that directly utilizes the v1/http plugin client

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugingetter/getter.go#L28)  

```go
type PluginWithV1Client interface {
	Client() *plugins.Client
}
```

---


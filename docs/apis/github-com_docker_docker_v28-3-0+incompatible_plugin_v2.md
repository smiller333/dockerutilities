# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/plugin/v2

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:54 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### ErrInadequateCapability

ErrInadequateCapability indicates that the plugin did not have the requested capability.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L40)  

```go
type ErrInadequateCapability struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### ErrInadequateCapability.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L44)  

```go
func (e ErrInadequateCapability) Error() string
```

---

### Plugin

Plugin represents an individual plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L19)  

```go
type Plugin struct {
	PluginObj types.Plugin `json:"plugin"` // todo: embed struct

	Rootfs string // TODO: make private

	Config   digest.Digest
	Blobsums []digest.Digest
	Manifest digest.Digest

	SwarmServiceID string
	// contains filtered or unexported fields
}
```

#### Methods

##### Plugin.Acquire

Acquire increments the plugin's reference count
This should be followed up by `Release()` when the plugin is no longer in use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L260)  

```go
func (p *Plugin) Acquire()
```

##### Plugin.AddRefCount

AddRefCount adds to reference count.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L251)  

```go
func (p *Plugin) AddRefCount(count int)
```

##### Plugin.Addr

Addr returns the net.Addr to use to connect to the plugin socket

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L296)  

```go
func (p *Plugin) Addr() net.Addr
```

##### Plugin.Client

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L60)  

```go
func (p *Plugin) Client() *plugins.Client
```

##### Plugin.FilterByCap

FilterByCap query the plugin for a given capability.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L88)  

```go
func (p *Plugin) FilterByCap(capability string) (*Plugin, error)
```

##### Plugin.GetID

GetID returns the plugin's ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L219)  

```go
func (p *Plugin) GetID() string
```

##### Plugin.GetRefCount

GetRefCount returns the reference count.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L243)  

```go
func (p *Plugin) GetRefCount() int
```

##### Plugin.GetSocket

GetSocket returns the plugin socket.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L227)  

```go
func (p *Plugin) GetSocket() string
```

##### Plugin.GetTypes

GetTypes returns the interface types of a plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L235)  

```go
func (p *Plugin) GetTypes() []types.PluginInterfaceType
```

##### Plugin.InitEmptySettings

InitEmptySettings initializes empty settings for a plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L99)  

```go
func (p *Plugin) InitEmptySettings()
```

##### Plugin.InitSpec

InitSpec creates an OCI spec from the plugin's config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin_linux.go#L22)  

```go
func (p *Plugin) InitSpec(execRoot string) (*specs.Spec, error)
```

##### Plugin.IsEnabled

IsEnabled returns the active state of the plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L211)  

```go
func (p *Plugin) IsEnabled() bool
```

##### Plugin.IsV1

IsV1 returns true for V1 plugins and false otherwise.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L78)  

```go
func (p *Plugin) IsV1() bool
```

##### Plugin.Name

Name returns the plugin name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L83)  

```go
func (p *Plugin) Name() string
```

##### Plugin.Protocol

Protocol is the protocol that should be used for interacting with the plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L311)  

```go
func (p *Plugin) Protocol() string
```

##### Plugin.Release

Release decrements the plugin's reference count
This should only be called when the plugin is no longer in use, e.g. with
via `Acquire()` or getter.Get("name", "type", plugingetter.Acquire)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L267)  

```go
func (p *Plugin) Release()
```

##### Plugin.ScopedPath

ScopedPath returns the path scoped to the plugin rootfs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L49)  

```go
func (p *Plugin) ScopedPath(s string) string
```

##### Plugin.Set

Set is used to pass arguments to the plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L115)  

```go
func (p *Plugin) Set(args []string) error
```

##### Plugin.SetAddr

SetAddr sets the plugin address which can be used for dialing the plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L304)  

```go
func (p *Plugin) SetAddr(addr net.Addr)
```

##### Plugin.SetPClient

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L70)  

```go
func (p *Plugin) SetPClient(client *plugins.Client)
```

##### Plugin.SetSpecOptModifier

SetSpecOptModifier sets the function to use to modify the generated
runtime spec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L273)  

```go
func (p *Plugin) SetSpecOptModifier(f func(*specs.Spec))
```

##### Plugin.SetTimeout

SetTimeout sets the timeout to use for dialing.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L289)  

```go
func (p *Plugin) SetTimeout(t time.Duration)
```

##### Plugin.Timeout

Timeout gets the currently configured connection timeout.
This should be used when dialing the plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/v2/plugin.go#L281)  

```go
func (p *Plugin) Timeout() time.Duration
```

---


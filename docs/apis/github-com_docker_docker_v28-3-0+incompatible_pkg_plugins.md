# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/plugins

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:02 UTC

## Overview

Package plugins provides structures and helper functions to manage Docker
plugins.

Docker discovers plugins by looking for them in the plugin directory whenever
a user or container tries to use one by name. UNIX domain socket files must
be located under /run/docker/plugins, whereas spec files can be located
either under /etc/docker/plugins or /usr/lib/docker/plugins. This is handled
by the Registry interface, which lets you list all plugins or get a plugin by
its name if it exists.

The plugins need to implement an HTTP server and bind this to the UNIX socket
or the address specified in the spec files.
A handshake is send at /Plugin.Activate, and plugins are expected to return
a Manifest with a list of Docker subsystems which this plugin implements.

In order to use a plugins, you can use the `Get` with the name of the
plugin and the subsystem it implements.


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L37)

```go
const ProtocolSchemeHTTPV1 = "moby.plugins.http/v1"
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L30)

```go
const VersionMimetype = transport.VersionMimetype
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/discovery.go#L19)

```go
var ErrNotFound = errors.New("plugin not found")
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L40)

```go
var ErrNotImplements = errors.New("Plugin does not implement the requested driver")
```

## Functions

### Handle

Handle adds the specified function to the extpointHandlers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L276)  

```go
func Handle(iface string, fn func(string, *Client))
```

---

### IsNotFound

IsNotFound indicates if the passed in error is from an http.StatusNotFound from the plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/errors.go#L20)  
**Added in:** v1.10.0

```go
func IsNotFound(err error) bool
```

---

### SpecsPaths

SpecsPaths returns paths in which to look for plugins, in order of priority.

On Windows:

On Unix in non-rootless mode:

On Unix in rootless-mode:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/discovery.go#L138)  

```go
func SpecsPaths() []string
```

---

### WithRequestTimeout

WithRequestTimeout sets a timeout duration for plugin requests

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L114)  

```go
func WithRequestTimeout(t time.Duration) func(*RequestOpts)
```

---

## Types

### Client

Client represents a plugin client.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L100)  

```go
type Client struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewClient

NewClient creates a new plugin client (http).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L65)  

```go
func NewClient(addr string, tlsConfig *tlsconfig.Options) (*Client, error)
```

##### NewClientWithTimeout

NewClientWithTimeout creates a new plugin client (http).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L74)  
**Added in:** v1.13.0

```go
func NewClientWithTimeout(addr string, tlsConfig *tlsconfig.Options, timeout time.Duration) (*Client, error)
```

#### Methods

##### Client.Call

Call calls the specified method with the specified arguments for the plugin.
It will retry for 30 seconds if a failure occurs when calling.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L122)  

```go
func (c *Client) Call(serviceMethod string, args, ret interface{}) error
```

##### Client.CallWithOptions

CallWithOptions is just like call except it takes options

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L127)  

```go
func (c *Client) CallWithOptions(serviceMethod string, args interface{}, ret interface{}, opts ...func(*RequestOpts)) error
```

##### Client.SendFile

SendFile calls the specified method, and passes through the IO stream

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L158)  
**Added in:** v1.9.0

```go
func (c *Client) SendFile(serviceMethod string, data io.Reader, ret interface{}) error
```

##### Client.Stream

Stream calls the specified method with the specified arguments for the plugin and returns the response body

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L149)  
**Added in:** v1.9.0

```go
func (c *Client) Stream(serviceMethod string, args interface{}) (io.ReadCloser, error)
```

---

### LocalRegistry

LocalRegistry defines a registry that is local (using unix socket).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/discovery.go#L24)  

```go
type LocalRegistry struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewLocalRegistry

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/discovery.go#L29)  

```go
func NewLocalRegistry() LocalRegistry
```

#### Methods

##### LocalRegistry.GetAll

GetAll returns all the plugins for the specified implementation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L300)  

```go
func (l *LocalRegistry) GetAll(imp string) ([]*Plugin, error)
```

##### LocalRegistry.Plugin

Plugin returns the plugin registered with the given name (or returns an error).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/discovery.go#L98)  

```go
func (l *LocalRegistry) Plugin(name string) (*Plugin, error)
```

##### LocalRegistry.Scan

Scan scans all the plugin paths and returns all the names it found

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/discovery.go#L37)  

```go
func (l *LocalRegistry) Scan() ([]string, error)
```

---

### Manifest

Manifest lists what a plugin implements.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L58)  

```go
type Manifest struct {
	// List of subsystem the plugin implements.
	Implements []string
}
```

---

### Plugin

Plugin is the definition of a docker plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L64)  

```go
type Plugin struct {

	// Address of the plugin
	Addr string
	// TLS configuration of the plugin
	TLSConfig *tlsconfig.Options

	// Manifest of the plugin (see above)
	Manifest *Manifest `json:"-"`
	// contains filtered or unexported fields
}
```

#### Functions

##### Get

Get returns the plugin given the specified name and requested implementation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L260)  

```go
func Get(name, imp string) (*Plugin, error)
```

##### NewLocalPlugin

NewLocalPlugin creates a new local plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L111)  
**Added in:** v1.12.0

```go
func NewLocalPlugin(name, addr string) *Plugin
```

#### Methods

##### Plugin.Client

Client returns a ready-to-use plugin client that can be used to communicate with the plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L90)  

```go
func (p *Plugin) Client() *Client
```

##### Plugin.IsV1

IsV1 returns true for V1 plugins and false otherwise.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L100)  
**Added in:** v1.13.0

```go
func (p *Plugin) IsV1() bool
```

##### Plugin.Name

Name returns the name of the plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L85)  

```go
func (p *Plugin) Name() string
```

##### Plugin.Protocol

Protocol returns the protocol name/version used for plugins in this package.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L95)  

```go
func (p *Plugin) Protocol() string
```

##### Plugin.ScopedPath

ScopedPath returns the path scoped to the plugin's rootfs.
For v1 plugins, this always returns the path unchanged as v1 plugins run directly on the host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/plugins.go#L106)  

```go
func (p *Plugin) ScopedPath(s string) string
```

---

### RequestOpts

RequestOpts is the set of options that can be passed into a request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/client.go#L106)  

```go
type RequestOpts struct {
	Timeout time.Duration
	// contains filtered or unexported fields
}
```

---


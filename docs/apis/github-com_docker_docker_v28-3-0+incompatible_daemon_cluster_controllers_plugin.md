# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/cluster/controllers/plugin

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:45 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Backend

Backend is the interface for interacting with the plugin manager
Controller actions are passed to the configured backend to do the real work.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L45)  

```go
type Backend interface {
	Disable(name string, config *backend.PluginDisableConfig) error
	Enable(name string, config *backend.PluginEnableConfig) error
	Remove(name string, config *backend.PluginRmConfig) error
	Pull(ctx context.Context, ref reference.Named, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer, opts ...plugin.CreateOpt) error
	Upgrade(ctx context.Context, ref reference.Named, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer) error
	Get(name string) (*v2.Plugin, error)
	SubscribeEvents(buffer int, events ...plugin.Event) (eventCh <-chan interface{}, cancel func())
}
```

---

### Controller

Controller is the controller for the plugin backend.
Plugins are managed as a singleton object with a desired state (different from containers).
With the plugin controller instead of having a strict create->start->stop->remove
task lifecycle like containers, we manage the desired state of the plugin and let
the plugin manager do what it already does and monitor the plugin.
We'll also end up with many tasks all pointing to the same plugin ID.

TODO(@cpuguy83): registry auth is intentionally not supported until we work out
the right way to pass registry credentials via secrets.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L31)  

```go
type Controller struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewController

NewController returns a new cluster plugin controller

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L56)  

```go
func NewController(backend Backend, t *api.Task) (*Controller, error)
```

#### Methods

##### Controller.Close

Close is the close phase from swarmkit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L246)  

```go
func (p *Controller) Close() error
```

##### Controller.Prepare

Prepare is the prepare phase from swarmkit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L90)  

```go
func (p *Controller) Prepare(ctx context.Context) (retErr error)
```

##### Controller.Remove

Remove is the remove phase from swarmkit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L219)  

```go
func (p *Controller) Remove(ctx context.Context) error
```

##### Controller.Shutdown

Shutdown is the shutdown phase from swarmkit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L207)  

```go
func (p *Controller) Shutdown(ctx context.Context) error
```

##### Controller.Start

Start is the start phase from swarmkit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L139)  

```go
func (p *Controller) Start(ctx context.Context) error
```

##### Controller.Terminate

Terminate is the terminate phase from swarmkit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L213)  

```go
func (p *Controller) Terminate(ctx context.Context) error
```

##### Controller.Update

Update is the update phase from swarmkit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L84)  

```go
func (p *Controller) Update(ctx context.Context, t *api.Task) error
```

##### Controller.Wait

Wait causes the task to wait until returned

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/controllers/plugin/controller.go#L160)  

```go
func (p *Controller) Wait(ctx context.Context) error
```

---


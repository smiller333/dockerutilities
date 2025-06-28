# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/plugin

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:28 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### CreateOpt

CreateOpt is used to configure specific plugin details when created

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/defs.go#L36)  

```go
type CreateOpt func(p *v2.Plugin)
```

#### Functions

##### WithEnv

WithEnv is a CreateOpt that passes the user-provided environment variables
to the plugin container, de-duplicating variables with the same names case
sensitively and only appends valid key=value pairs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/defs.go#L49)  

```go
func WithEnv(env []string) CreateOpt
```

##### WithSwarmService

WithSwarmService is a CreateOpt that flags the passed in a plugin as a plugin
managed by swarm

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/defs.go#L40)  

```go
func WithSwarmService(id string) CreateOpt
```

---

### EndpointResolver

EndpointResolver provides looking up registry endpoints for pulling.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager.go#L49)  

```go
type EndpointResolver interface {
	LookupPullEndpoints(hostname string) (endpoints []registry.APIEndpoint, err error)
}
```

---

### Event

Event is emitted for actions performed on the plugin manager

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/events.go#L11)  

```go
type Event interface {
	// contains filtered or unexported methods
}
```

---

### EventCreate

EventCreate is an event which is emitted when a plugin is created
This is either by pull or create from context.

Use the `Interfaces` field to match only plugins that implement a specific
interface.
These are matched against using "or" logic.
If no interfaces are listed, all are matched.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/events.go#L22)  

```go
type EventCreate struct {
	Interfaces map[string]bool
	Plugin     types.Plugin
}
```

---

### EventDisable

EventDisable is an event that is emitted when a plugin is disabled
It matches on the passed in plugin's ID only.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/events.go#L62)  

```go
type EventDisable struct {
	Plugin types.Plugin
}
```

---

### EventEnable

EventEnable is an event that is emitted when a plugin is disabled
It matches on the passed in plugin's ID only.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/events.go#L76)  

```go
type EventEnable struct {
	Plugin types.Plugin
}
```

---

### EventRemove

EventRemove is an event which is emitted when a plugin is removed
It matches on the passed in plugin's ID only.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/events.go#L48)  

```go
type EventRemove struct {
	Plugin types.Plugin
}
```

---

### Executor

Executor is the interface that the plugin manager uses to interact with for starting/stopping plugins

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager.go#L41)  

```go
type Executor interface {
	Create(id string, spec specs.Spec, stdout, stderr io.WriteCloser) error
	IsRunning(id string) (bool, error)
	Restore(id string, stdout, stderr io.WriteCloser) (alive bool, err error)
	Signal(id string, signal syscall.Signal) error
}
```

---

### ExecutorCreator

ExecutorCreator is used in the manager config to pass in an `Executor`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager.go#L75)  

```go
type ExecutorCreator func(*Manager) (Executor, error)
```

---

### Manager

Manager controls the plugin subsystem.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager.go#L78)  
**Added in:** v1.13.0

```go
type Manager struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewManager

NewManager returns a new plugin manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager.go#L96)  
**Added in:** v1.13.0

```go
func NewManager(config ManagerConfig) (*Manager, error)
```

#### Methods

##### Manager.CreateFromContext

CreateFromContext creates a plugin from the given pluginDir which contains
both the rootfs and the config.json and a repoName with optional tag.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L623)  
**Added in:** v1.13.0

```go
func (pm *Manager) CreateFromContext(ctx context.Context, tarCtx io.ReadCloser, options *types.PluginCreateOptions) (retErr error)
```

##### Manager.Disable

Disable deactivates a plugin. This means resources (volumes, networks) cant use them.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L51)  
**Added in:** v1.13.0

```go
func (pm *Manager) Disable(refOrID string, config *backend.PluginDisableConfig) error
```

##### Manager.Enable

Enable activates a plugin, which implies that they are ready to be used by containers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L79)  
**Added in:** v1.13.0

```go
func (pm *Manager) Enable(refOrID string, config *backend.PluginEnableConfig) error
```

##### Manager.GC

GC cleans up unreferenced blobs. This is recommended to run in a goroutine

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager.go#L291)  
**Added in:** v1.13.0

```go
func (pm *Manager) GC()
```

##### Manager.Get

Get looks up the requested plugin in the store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager.go#L262)  

```go
func (pm *Manager) Get(idOrName string) (*v2.Plugin, error)
```

##### Manager.HandleExitEvent

HandleExitEvent is called when the executor receives the exit event
In the future we may change this, but for now all we care about is the exit event.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager.go#L131)  

```go
func (pm *Manager) HandleExitEvent(id string) error
```

##### Manager.Inspect

Inspect examines a plugin config

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L95)  
**Added in:** v1.13.0

```go
func (pm *Manager) Inspect(refOrID string) (*types.Plugin, error)
```

##### Manager.List

List displays the list of plugins and associated metadata.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L315)  
**Added in:** v1.13.0

```go
func (pm *Manager) List(pluginFilters filters.Args) ([]types.Plugin, error)
```

##### Manager.Privileges

Privileges pulls a plugin config and computes the privileges required to install it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L164)  
**Added in:** v1.13.0

```go
func (pm *Manager) Privileges(ctx context.Context, ref reference.Named, metaHeader http.Header, authConfig *registry.AuthConfig) (types.PluginPrivileges, error)
```

##### Manager.Pull

Pull pulls a plugin, check if the correct privileges are provided and install the plugin.

TODO: replace reference package usage with simpler url.Parse semantics

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L260)  
**Added in:** v1.13.0

```go
func (pm *Manager) Pull(ctx context.Context, ref reference.Named, name string, metaHeader http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer, opts ...CreateOpt) error
```

##### Manager.Push

Push pushes a plugin to the registry.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L359)  
**Added in:** v1.13.0

```go
func (pm *Manager) Push(ctx context.Context, name string, metaHeader http.Header, authConfig *registry.AuthConfig, outStream io.Writer) error
```

##### Manager.Remove

Remove deletes plugin's root directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L563)  
**Added in:** v1.13.0

```go
func (pm *Manager) Remove(name string, config *backend.PluginRmConfig) error
```

##### Manager.Set

Set sets plugin args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L610)  
**Added in:** v1.13.0

```go
func (pm *Manager) Set(name string, args []string) error
```

##### Manager.Shutdown

Shutdown stops all plugins and called during daemon shutdown.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager_linux.go#L194)  
**Added in:** v1.13.0

```go
func (pm *Manager) Shutdown()
```

##### Manager.SubscribeEvents

SubscribeEvents provides an event channel to listen for structured events from
the plugin manager actions, CRUD operations.
The caller must call the returned `cancel()` function once done with the channel
or this will leak resources.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/events.go#L92)  

```go
func (pm *Manager) SubscribeEvents(buffer int, watchEvents ...Event) (eventCh <-chan interface{}, cancel func())
```

##### Manager.Upgrade

Upgrade upgrades a plugin

TODO: replace reference package usage with simpler url.Parse semantics

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/backend_linux.go#L212)  
**Added in:** v1.13.1

```go
func (pm *Manager) Upgrade(ctx context.Context, ref reference.Named, name string, metaHeader http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer) error
```

---

### ManagerConfig

ManagerConfig defines configuration needed to start new manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/manager.go#L63)  
**Added in:** v1.13.0

```go
type ManagerConfig struct {
	Store              *Store // remove
	RegistryService    EndpointResolver
	LiveRestoreEnabled bool // TODO: remove
	LogPluginEvent     eventLogger
	Root               string
	ExecRoot           string
	CreateExecutor     ExecutorCreator
	AuthzMiddleware    *authorization.Middleware
}
```

---

### SpecOpt

SpecOpt is used for subsystems that need to modify the runtime spec of a plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/defs.go#L33)  

```go
type SpecOpt func(*specs.Spec)
```

#### Functions

##### WithSpecMounts

WithSpecMounts is a SpecOpt which appends the provided mounts to the runtime spec

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/defs.go#L70)  

```go
func WithSpecMounts(mounts []specs.Mount) SpecOpt
```

---

### Store

Store manages the plugin inventory in memory and on-disk

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/defs.go#L13)  
**Added in:** v1.13.0

```go
type Store struct {
	sync.RWMutex
	// contains filtered or unexported fields
}
```

#### Functions

##### NewStore

NewStore creates a Store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/defs.go#L24)  
**Added in:** v1.13.0

```go
func NewStore() *Store
```

#### Methods

##### Store.Add

Add adds a plugin to memory and plugindb.
An error will be returned if there is a collision.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L115)  
**Added in:** v1.13.0

```go
func (ps *Store) Add(p *v2.Plugin) error
```

##### Store.CallHandler

CallHandler calls the registered callback. It is invoked during plugin enable.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L244)  
**Added in:** v1.13.0

```go
func (ps *Store) CallHandler(p *v2.Plugin)
```

##### Store.Get

Get returns an enabled plugin matching the given name and capability.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L137)  
**Added in:** v1.13.0

```go
func (ps *Store) Get(name, capability string, mode int) (plugingetter.CompatPlugin, error)
```

##### Store.GetAll

GetAll retrieves all plugins.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L57)  
**Added in:** v1.13.0

```go
func (ps *Store) GetAll() map[string]*v2.Plugin
```

##### Store.GetAllByCap

GetAllByCap returns a list of enabled plugins matching the given capability.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L181)  
**Added in:** v1.13.0

```go
func (ps *Store) GetAllByCap(capability string) ([]plugingetter.CompatPlugin, error)
```

##### Store.GetAllManagedPluginsByCap

GetAllManagedPluginsByCap returns a list of managed plugins matching the given capability.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L176)  
**Added in:** v1.13.0

```go
func (ps *Store) GetAllManagedPluginsByCap(capability string) []plugingetter.CompatPlugin
```

##### Store.GetV2Plugin

GetV2Plugin retrieves a plugin by name, id or partial ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L29)  
**Added in:** v1.13.0

```go
func (ps *Store) GetV2Plugin(refOrID string) (*v2.Plugin, error)
```

##### Store.Handle

Handle sets a callback for a given capability. It is only used by network
and ipam drivers during plugin registration. The callback registers the
driver with the subsystem (network, ipam).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L215)  
**Added in:** v1.13.0

```go
func (ps *Store) Handle(capability string, callback func(string, *plugins.Client))
```

##### Store.RegisterRuntimeOpt

RegisterRuntimeOpt stores a list of SpecOpts for the provided capability.
These options are applied to the runtime spec before a plugin is started for the specified capability.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L236)  

```go
func (ps *Store) RegisterRuntimeOpt(cap string, opts ...SpecOpt)
```

##### Store.Remove

Remove removes a plugin from memory and plugindb.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L130)  
**Added in:** v1.13.0

```go
func (ps *Store) Remove(p *v2.Plugin)
```

##### Store.SetAll

SetAll initialized plugins during daemon restore.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L64)  
**Added in:** v1.13.0

```go
func (ps *Store) SetAll(plugins map[string]*v2.Plugin)
```

##### Store.SetState

SetState sets the active state of the plugin and updates plugindb.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/plugin/store.go#L90)  
**Added in:** v1.13.0

```go
func (ps *Store) SetState(p *v2.Plugin, state bool)
```

---


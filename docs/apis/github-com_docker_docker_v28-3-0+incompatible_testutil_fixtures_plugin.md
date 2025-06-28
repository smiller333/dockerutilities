# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil/fixtures/plugin

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:33 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Create

Create creates a new plugin with the specified name

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fixtures/plugin/plugin.go#L56)  

```go
func Create(ctx context.Context, c CreateClient, name string, opts ...CreateOpt) error
```

---

### CreateInRegistry

CreateInRegistry makes a plugin (locally) and pushes it to a registry.
This does not use a dockerd instance to create or push the plugin.
If you just want to create a plugin in some daemon, use `Create`.

This can be useful when testing plugins on swarm where you don't really want
the plugin to exist on any of the daemons (immediately) and there needs to be
some way to distribute the plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fixtures/plugin/plugin.go#L82)  

```go
func CreateInRegistry(ctx context.Context, repo string, auth *registry.AuthConfig, opts ...CreateOpt) error
```

---

## Types

### Config

Config wraps types.PluginConfig to provide some extra state for options
extra customizations on the plugin details, such as using a custom binary to
create the plugin with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fixtures/plugin/plugin.go#L28)  

```go
type Config struct {
	*types.PluginConfig

	RegistryConfig registrypkg.ServiceOptions
	// contains filtered or unexported fields
}
```

---

### CreateClient

CreateClient is the interface used for `BuildPlugin` to interact with the
daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fixtures/plugin/plugin.go#L51)  

```go
type CreateClient interface {
	PluginCreate(context.Context, io.Reader, types.PluginCreateOptions) error
}
```

---

### CreateOpt

CreateOpt is passed used to change the default plugin config before
creating it

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fixtures/plugin/plugin.go#L23)  

```go
type CreateOpt func(*Config)
```

#### Functions

##### WithBinary

WithBinary is a CreateOpt to set an custom binary to create the plugin with.
This binary must be statically compiled.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fixtures/plugin/plugin.go#L43)  

```go
func WithBinary(bin string) CreateOpt
```

##### WithInsecureRegistry

WithInsecureRegistry specifies that the given registry can skip host-key checking as well as fall back to plain http

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fixtures/plugin/plugin.go#L35)  

```go
func WithInsecureRegistry(url string) CreateOpt
```

---


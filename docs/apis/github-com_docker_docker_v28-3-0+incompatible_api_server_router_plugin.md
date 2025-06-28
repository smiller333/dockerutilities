# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/plugin

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:26 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new plugin router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/plugin/plugin.go#L12)  

```go
func NewRouter(b Backend) router.Router
```

---

## Types

### Backend

Backend for Plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/plugin/backend.go#L17)  

```go
type Backend interface {
	Disable(name string, config *backend.PluginDisableConfig) error
	Enable(name string, config *backend.PluginEnableConfig) error
	List(filters.Args) ([]types.Plugin, error)
	Inspect(name string) (*types.Plugin, error)
	Remove(name string, config *backend.PluginRmConfig) error
	Set(name string, args []string) error
	Privileges(ctx context.Context, ref reference.Named, metaHeaders http.Header, authConfig *registry.AuthConfig) (types.PluginPrivileges, error)
	Pull(ctx context.Context, ref reference.Named, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer, opts ...plugin.CreateOpt) error
	Push(ctx context.Context, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, outStream io.Writer) error
	Upgrade(ctx context.Context, ref reference.Named, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer) error
	CreateFromContext(ctx context.Context, tarCtx io.ReadCloser, options *types.PluginCreateOptions) error
}
```

---


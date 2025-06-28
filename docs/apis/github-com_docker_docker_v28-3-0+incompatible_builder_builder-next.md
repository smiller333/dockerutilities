# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/builder-next

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:37 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Builder

Builder can build using BuildKit backend

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/builder.go#L105)  

```go
type Builder struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates a new builder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/builder.go#L116)  

```go
func New(ctx context.Context, opt Opt) (*Builder, error)
```

#### Methods

##### Builder.Build

Build executes a build request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/builder.go#L245)  

```go
func (b *Builder) Build(ctx context.Context, opt backend.BuildConfig) (*builder.Result, error)
```

##### Builder.Cancel

Cancel cancels a build using ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/builder.go#L143)  

```go
func (b *Builder) Cancel(ctx context.Context, id string) error
```

##### Builder.Close

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/builder.go#L133)  

```go
func (b *Builder) Close() error
```

##### Builder.DiskUsage

DiskUsage returns a report about space used by build cache

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/builder.go#L153)  

```go
func (b *Builder) DiskUsage(ctx context.Context) ([]*build.CacheRecord, error)
```

##### Builder.Prune

Prune clears all reclaimable build cache.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/builder.go#L190)  

```go
func (b *Builder) Prune(ctx context.Context, opts build.CachePruneOptions) (int64, []string, error)
```

##### Builder.RegisterGRPC

RegisterGRPC registers controller to the grpc server.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/builder.go#L138)  

```go
func (b *Builder) RegisterGRPC(s *grpc.Server)
```

---

### Opt

Opt is option struct required for creating the builder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/builder.go#L82)  

```go
type Opt struct {
	SessionManager      *session.Manager
	Root                string
	EngineID            string
	Dist                images.DistributionServices
	ImageTagger         mobyexporter.ImageTagger
	NetworkController   *libnetwork.Controller
	DefaultCgroupParent string
	RegistryHosts       docker.RegistryHosts
	BuilderConfig       config.BuilderConfig
	Rootless            bool
	IdentityMapping     user.IdentityMapping
	DNSConfig           config.DNSConfig
	ApparmorProfile     string
	UseSnapshotter      bool
	Snapshotter         string
	ContainerdAddress   string
	ContainerdNamespace string
	Callbacks           exporter.BuildkitCallbacks
	CDICache            *cdi.Cache
}
```

---


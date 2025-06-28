# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/volume

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:36 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new volume router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/volume/volume.go#L13)  

```go
func NewRouter(b Backend, cb ClusterBackend) router.Router
```

---

## Types

### Backend

Backend is the methods that need to be implemented to provide
volume specific functionality

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/volume/backend.go#L13)  

```go
type Backend interface {
	List(ctx context.Context, filter filters.Args) ([]*volume.Volume, []string, error)
	Get(ctx context.Context, name string, opts ...opts.GetOption) (*volume.Volume, error)
	Create(ctx context.Context, name, driverName string, opts ...opts.CreateOption) (*volume.Volume, error)
	Remove(ctx context.Context, name string, opts ...opts.RemoveOption) error
	Prune(ctx context.Context, pruneFilters filters.Args) (*volume.PruneReport, error)
}
```

---

### ClusterBackend

ClusterBackend is the backend used for Swarm Cluster Volumes. Regular
volumes go through the volume service, but to avoid across-dependency
between the cluster package and the volume package, we simply provide two
backends here.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/volume/backend.go#L25)  

```go
type ClusterBackend interface {
	GetVolume(nameOrID string) (volume.Volume, error)
	GetVolumes(options volume.ListOptions) ([]*volume.Volume, error)
	CreateVolume(volume volume.CreateOptions) (*volume.Volume, error)
	RemoveVolume(nameOrID string, force bool) error
	UpdateVolume(nameOrID string, version uint64, volume volume.UpdateOptions) error
	IsManager() bool
}
```

---


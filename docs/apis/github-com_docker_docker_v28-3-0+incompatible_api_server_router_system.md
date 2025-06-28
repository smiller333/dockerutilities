# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/system

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:34 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new system router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/system/system.go#L27)  

```go
func NewRouter(b Backend, c ClusterBackend, builder BuildBackend, features func() map[string]bool) router.Router
```

---

## Types

### Backend

Backend is the methods that need to be implemented to provide
system specific functionality.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/system/backend.go#L30)  

```go
type Backend interface {
	SystemInfo(context.Context) (*system.Info, error)
	SystemVersion(context.Context) (types.Version, error)
	SystemDiskUsage(ctx context.Context, opts DiskUsageOptions) (*system.DiskUsage, error)
	SubscribeToEvents(since, until time.Time, ef filters.Args) ([]events.Message, chan interface{})
	UnsubscribeFromEvents(chan interface{})
	AuthenticateToRegistry(ctx context.Context, authConfig *registry.AuthConfig) (string, string, error)
}
```

---

### BuildBackend

BuildBackend provides build specific system information.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/system/backend.go#L46)  

```go
type BuildBackend interface {
	DiskUsage(context.Context) ([]*build.CacheRecord, error)
}
```

---

### ClusterBackend

ClusterBackend is all the methods that need to be implemented
to provide cluster system specific functionality.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/system/backend.go#L41)  

```go
type ClusterBackend interface {
	Info(context.Context) swarm.Info
}
```

---

### DiskUsageOptions

DiskUsageOptions holds parameters for system disk usage query.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/system/backend.go#L17)  

```go
type DiskUsageOptions struct {
	// Containers controls whether container disk usage should be computed.
	Containers bool

	// Images controls whether image disk usage should be computed.
	Images bool

	// Volumes controls whether volume disk usage should be computed.
	Volumes bool
}
```

---

### StatusProvider

StatusProvider provides methods to get the swarm status of the current node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/system/backend.go#L51)  

```go
type StatusProvider interface {
	Status() string
}
```

---


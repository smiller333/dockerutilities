# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/backend/build

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:24:59 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Backend

Backend provides build functionality to the API router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/backend/build/backend.go#L33)  

```go
type Backend struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewBackend

NewBackend creates a new build backend from components

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/backend/build/backend.go#L41)  

```go
func NewBackend(components ImageComponent, builder Builder, buildkit *buildkit.Builder, es *daemonevents.Events) (*Backend, error)
```

#### Methods

##### Backend.Build

Build builds an image from a Source

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/backend/build/backend.go#L53)  

```go
func (b *Backend) Build(ctx context.Context, config backend.BuildConfig) (string, error)
```

##### Backend.Cancel

Cancel cancels the build by ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/backend/build/backend.go#L114)  

```go
func (b *Backend) Cancel(ctx context.Context, id string) error
```

##### Backend.PruneCache

PruneCache removes all cached build sources

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/backend/build/backend.go#L100)  

```go
func (b *Backend) PruneCache(ctx context.Context, opts build.CachePruneOptions) (*build.CachePruneReport, error)
```

##### Backend.RegisterGRPC

RegisterGRPC registers buildkit controller to the grpc server.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/backend/build/backend.go#L46)  

```go
func (b *Backend) RegisterGRPC(s *grpc.Server)
```

---

### Builder

Builder defines interface for running a build

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/backend/build/backend.go#L28)  

```go
type Builder interface {
	Build(context.Context, backend.BuildConfig) (*builder.Result, error)
}
```

---

### ImageComponent

ImageComponent provides an interface for working with images

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/backend/build/backend.go#L22)  

```go
type ImageComponent interface {
	SquashImage(from string, to string) (string, error)
	TagImage(context.Context, image.ID, reference.Named) error
}
```

---


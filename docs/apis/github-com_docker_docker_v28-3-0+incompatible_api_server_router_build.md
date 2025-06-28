# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/build

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:10 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### BuilderVersion

BuilderVersion derives the default docker builder version from the config.

The default on Linux is version "2" (BuildKit), but the daemon can be
configured to recommend version "1" (classic Builder). Windows does not
yet support BuildKit for native Windows images, and uses "1" (classic builder)
as a default.

This value is only a recommendation as advertised by the daemon, and it is
up to the client to choose which builder to use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/build/build.go#L49)  

```go
func BuilderVersion(features map[string]bool) build.BuilderVersion
```

---

### NewRouter

NewRouter initializes a new build router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/build/build.go#L18)  

```go
func NewRouter(b Backend, d experimentalProvider) router.Router
```

---

## Types

### Backend

Backend abstracts an image builder whose only purpose is to build an image referenced by an imageID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/build/backend.go#L11)  

```go
type Backend interface {
	// Build a Docker image returning the id of the image
	// TODO: make this return a reference instead of string
	Build(context.Context, backend.BuildConfig) (string, error)

	// PruneCache prunes the build cache.
	PruneCache(context.Context, build.CachePruneOptions) (*build.CachePruneReport, error)
	Cancel(context.Context, string) error
}
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/distribution

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:18 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new distribution router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/distribution/distribution.go#L12)  

```go
func NewRouter(backend Backend) router.Router
```

---

## Types

### Backend

Backend is all the methods that need to be implemented
to provide image specific functionality.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/distribution/backend.go#L13)  

```go
type Backend interface {
	GetRepositories(context.Context, reference.Named, *registry.AuthConfig) ([]distribution.Repository, error)
}
```

---


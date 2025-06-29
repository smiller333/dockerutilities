# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/container

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:12 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new container router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/container/container.go#L17)  

```go
func NewRouter(b Backend, decoder httputils.ContainerDecoder, cgroup2 bool) router.Router
```

---

## Types

### Backend

Backend is all the methods that need to be implemented to provide container specific functionality.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/container/backend.go#L71)  

```go
type Backend interface {
	// contains filtered or unexported methods
}
```

---


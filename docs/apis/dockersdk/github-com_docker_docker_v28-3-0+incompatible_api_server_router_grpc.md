# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/grpc

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:19 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new grpc http router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/grpc/grpc.go#L31)  

```go
func NewRouter(backends ...Backend) router.Router
```

---

## Types

### Backend

Backend abstracts a registerable GRPC service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/grpc/backend.go#L6)  

```go
type Backend interface {
	RegisterGRPC(*grpc.Server)
}
```

---


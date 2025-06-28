# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/checkpoint

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:09 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new checkpoint router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/checkpoint/checkpoint.go#L16)  

```go
func NewRouter(b Backend, decoder httputils.ContainerDecoder) router.Router
```

---

## Types

### Backend

Backend for Checkpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/checkpoint/backend.go#L6)  

```go
type Backend interface {
	CheckpointCreate(container string, config checkpoint.CreateOptions) error
	CheckpointDelete(container string, config checkpoint.DeleteOptions) error
	CheckpointList(container string, config checkpoint.ListOptions) ([]checkpoint.Summary, error)
}
```

---


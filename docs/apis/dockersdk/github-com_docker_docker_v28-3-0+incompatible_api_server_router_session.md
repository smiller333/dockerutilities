# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/session

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:29 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new session router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/session/session.go#L12)  

```go
func NewRouter(b Backend) router.Router
```

---

## Types

### Backend

Backend abstracts an session receiver from an http request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/session/backend.go#L9)  

```go
type Backend interface {
	HandleHTTPRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}
```

---


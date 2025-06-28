# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router/image

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:23 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewRouter

NewRouter initializes a new image router

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/image/image.go#L15)  

```go
func NewRouter(backend Backend, searcher Searcher) router.Router
```

---

## Types

### Backend

Backend is all the methods that need to be implemented
to provide image specific functionality.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/image/backend.go#L18)  

```go
type Backend interface {
	// contains filtered or unexported methods
}
```

---

### Searcher

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/image/backend.go#L45)  

```go
type Searcher interface {
	Search(ctx context.Context, searchFilters filters.Args, term string, limit int, authConfig *registry.AuthConfig, headers map[string][]string) ([]registry.SearchResult, error)
}
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/middleware

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:02 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### DebugRequestMiddleware

DebugRequestMiddleware dumps the request to logger

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/middleware/debug.go#L19)  

```go
func DebugRequestMiddleware(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error) func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
```

---

## Types

### ExperimentalMiddleware

ExperimentalMiddleware is a the middleware in charge of adding the
'Docker-Experimental' header to every outgoing request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/middleware/experimental.go#L10)  
**Added in:** v1.13.0

```go
type ExperimentalMiddleware struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewExperimentalMiddleware

NewExperimentalMiddleware creates a new ExperimentalMiddleware

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/middleware/experimental.go#L15)  
**Added in:** v1.13.0

```go
func NewExperimentalMiddleware(experimentalEnabled bool) ExperimentalMiddleware
```

#### Methods

##### ExperimentalMiddleware.WrapHandler

WrapHandler returns a new handler function wrapping the previous one in the request chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/middleware/experimental.go#L23)  
**Added in:** v1.13.0

```go
func (e ExperimentalMiddleware) WrapHandler(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error) func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
```

---

### Middleware

Middleware is an interface to allow the use of ordinary functions as Docker API filters.
Any struct that has the appropriate signature can be registered as a middleware.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/middleware/middleware.go#L10)  

```go
type Middleware interface {
	WrapHandler(func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error) func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
}
```

---

### VersionMiddleware

VersionMiddleware is a middleware that
validates the client and server versions.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/middleware/version.go#L16)  
**Added in:** v1.12.0

```go
type VersionMiddleware struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewVersionMiddleware

NewVersionMiddleware creates a VersionMiddleware with the given versions.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/middleware/version.go#L36)  

```go
func NewVersionMiddleware(serverVersion, defaultAPIVersion, minAPIVersion string) (*VersionMiddleware, error)
```

#### Methods

##### VersionMiddleware.WrapHandler

WrapHandler returns a new handler function wrapping the previous one in the request chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/middleware/version.go#L67)  
**Added in:** v1.12.0

```go
func (v VersionMiddleware) WrapHandler(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error) func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
```

---


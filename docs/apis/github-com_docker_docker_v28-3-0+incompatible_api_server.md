# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:24:57 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Server

Server contains instance details for the server

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/server.go#L26)  
**Added in:** v1.4.0

```go
type Server struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Server.CreateMux

CreateMux returns a new mux with all the routers registered.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/server.go#L85)  
**Added in:** v1.9.0

```go
func (s *Server) CreateMux(ctx context.Context, routers ...router.Router) *mux.Router
```

##### Server.UseMiddleware

UseMiddleware appends a new middleware to the request chain.
This needs to be called before the API routes are configured.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/server.go#L32)  
**Added in:** v1.12.0

```go
func (s *Server) UseMiddleware(m middleware.Middleware)
```

---


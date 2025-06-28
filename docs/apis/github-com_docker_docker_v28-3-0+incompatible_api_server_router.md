# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/router

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:08 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### ExperimentalRoute

ExperimentalRoute defines an experimental API route that can be enabled or disabled.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/experimental.go#L11)  
**Added in:** v1.13.0

```go
type ExperimentalRoute interface {
	Route

	Enable()
	Disable()
}
```

---

### Route

Route defines an individual API route in the docker server.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/router.go#L12)  

```go
type Route interface {
	// Handler returns the raw function to create the http handler.
	Handler() httputils.APIFunc
	// Method returns the http method that the route responds to.
	Method() string
	// Path returns the subpath where the route responds to.
	Path() string
}
```

#### Functions

##### Experimental

Experimental will mark a route as experimental.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/experimental.go#L63)  
**Added in:** v1.13.0

```go
func Experimental(r Route) Route
```

##### NewDeleteRoute

NewDeleteRoute initializes a new route with the http method DELETE.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/local.go#L61)  
**Added in:** v1.11.0

```go
func NewDeleteRoute(path string, handler httputils.APIFunc, opts ...RouteWrapper) Route
```

##### NewGetRoute

NewGetRoute initializes a new route with the http method GET.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/local.go#L46)  
**Added in:** v1.11.0

```go
func NewGetRoute(path string, handler httputils.APIFunc, opts ...RouteWrapper) Route
```

##### NewHeadRoute

NewHeadRoute initializes a new route with the http method HEAD.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/local.go#L71)  
**Added in:** v1.11.0

```go
func NewHeadRoute(path string, handler httputils.APIFunc, opts ...RouteWrapper) Route
```

##### NewOptionsRoute

NewOptionsRoute initializes a new route with the http method OPTIONS.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/local.go#L66)  
**Added in:** v1.11.0

```go
func NewOptionsRoute(path string, handler httputils.APIFunc, opts ...RouteWrapper) Route
```

##### NewPostRoute

NewPostRoute initializes a new route with the http method POST.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/local.go#L51)  
**Added in:** v1.11.0

```go
func NewPostRoute(path string, handler httputils.APIFunc, opts ...RouteWrapper) Route
```

##### NewPutRoute

NewPutRoute initializes a new route with the http method PUT.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/local.go#L56)  
**Added in:** v1.11.0

```go
func NewPutRoute(path string, handler httputils.APIFunc, opts ...RouteWrapper) Route
```

##### NewRoute

NewRoute initializes a new local route for the router.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/local.go#L37)  
**Added in:** v1.11.0

```go
func NewRoute(method, path string, handler httputils.APIFunc, opts ...RouteWrapper) Route
```

---

### RouteWrapper

RouteWrapper wraps a route with extra functionality.
It is passed in when creating a new route.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/local.go#L11)  

```go
type RouteWrapper func(r Route) Route
```

---

### Router

Router defines an interface to specify a group of routes to add to the docker server.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/router/router.go#L6)  

```go
type Router interface {
	// Routes returns the list of routes to add to the docker server.
	Routes() []Route
}
```

---


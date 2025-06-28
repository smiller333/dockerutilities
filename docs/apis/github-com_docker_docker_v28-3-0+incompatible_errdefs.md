# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/errdefs

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:58 UTC

## Overview

Package errdefs defines a set of error interfaces that packages should use for communicating classes of errors.
Errors that cross the package boundary should implement one (and only one) of these interfaces.

Packages should not reference these interfaces directly, only implement them.
To check if a particular error implements one of these interfaces, there are helper
functions provided (e.g. `Is<SomeError>`) which can be used rather than asserting the interfaces directly.
If you must assert on these interfaces, be sure to check the causal chain (`err.Unwrap()`).


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L63)

```go
var IsCancelled = cerrdefs.IsCanceled
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L23)

```go
var IsConflict = cerrdefs.IsConflict
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L73)

```go
var IsDataLoss = cerrdefs.IsDataLoss
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L68)

```go
var IsDeadline = cerrdefs.IsDeadlineExceeded
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L38)

```go
var IsForbidden = cerrdefs.IsPermissionDenied
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L18)

```go
var IsInvalidParameter = cerrdefs.IsInvalidArgument
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L13)

```go
var IsNotFound = cerrdefs.IsNotFound
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L53)

```go
var IsNotImplemented = cerrdefs.IsNotImplemented
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L48)

```go
var IsNotModified = cerrdefs.IsNotModified
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L43)

```go
var IsSystem = cerrdefs.IsInternal
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L28)

```go
var IsUnauthorized = cerrdefs.IsUnauthorized
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L33)

```go
var IsUnavailable = cerrdefs.IsUnavailable
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L58)

```go
var IsUnknown = cerrdefs.IsUnknown
```

## Functions

### Cancelled

Cancelled creates an ErrCancelled error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrCancelled,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L244)  

```go
func Cancelled(err error) error
```

---

### Conflict

Conflict creates an ErrConflict error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrConflict,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L68)  

```go
func Conflict(err error) error
```

---

### DataLoss

DataLoss creates an ErrDataLoss error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrDataLoss,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L288)  

```go
func DataLoss(err error) error
```

---

### Deadline

Deadline creates an ErrDeadline error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrDeadline,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L266)  

```go
func Deadline(err error) error
```

---

### Forbidden

Forbidden creates an ErrForbidden error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrForbidden,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L134)  

```go
func Forbidden(err error) error
```

---

### FromContext

FromContext returns the error class from the passed in context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L296)  

```go
func FromContext(ctx context.Context) error
```

---

### FromStatusCode ⚠️ **DEPRECATED**

FromStatusCode creates an errdef error, based on the provided HTTP status-code

Deprecated: Use cerrdefs.ToNative instead

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/http_helpers.go#L10)  

```go
func FromStatusCode(err error, statusCode int) error
```

---

### InvalidParameter

InvalidParameter creates an ErrInvalidParameter error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrInvalidParameter,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L46)  

```go
func InvalidParameter(err error) error
```

---

### IsContext

IsContext returns if the passed in error is due to context cancellation or deadline exceeded.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/is.go#L76)  

```go
func IsContext(err error) bool
```

---

### NotFound

NotFound creates an ErrNotFound error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrNotFound,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L24)  

```go
func NotFound(err error) error
```

---

### NotImplemented

NotImplemented creates an ErrNotImplemented error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrNotImplemented,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L200)  

```go
func NotImplemented(err error) error
```

---

### NotModified

NotModified creates an ErrNotModified error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
NotModified,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L178)  

```go
func NotModified(err error) error
```

---

### System

System creates an ErrSystem error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrSystem,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L156)  

```go
func System(err error) error
```

---

### Unauthorized

Unauthorized creates an ErrUnauthorized error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrUnauthorized,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L90)  

```go
func Unauthorized(err error) error
```

---

### Unavailable

Unavailable creates an ErrUnavailable error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrUnavailable,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L112)  

```go
func Unavailable(err error) error
```

---

### Unknown

Unknown creates an ErrUnknown error from the given error.
It returns the error as-is if it is either nil (no error) or already implements
ErrUnknown,

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/helpers.go#L222)  

```go
func Unknown(err error) error
```

---

## Types

### ErrCancelled

ErrCancelled signals that the action was cancelled.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L57)  

```go
type ErrCancelled interface {
	Cancelled()
}
```

---

### ErrConflict

ErrConflict signals that some internal state conflicts with the requested action and can't be performed.
A change in state should be able to clear this error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L15)  

```go
type ErrConflict interface {
	Conflict()
}
```

---

### ErrDataLoss

ErrDataLoss indicates that data was lost or there is data corruption.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L67)  

```go
type ErrDataLoss interface {
	DataLoss()
}
```

---

### ErrDeadline

ErrDeadline signals that the deadline was reached before the action completed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L62)  

```go
type ErrDeadline interface {
	DeadlineExceeded()
}
```

---

### ErrForbidden

ErrForbidden signals that the requested action cannot be performed under any circumstances.
When a ErrForbidden is returned, the caller should never retry the action.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L31)  

```go
type ErrForbidden interface {
	Forbidden()
}
```

---

### ErrInvalidParameter

ErrInvalidParameter signals that the user input is invalid

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L9)  

```go
type ErrInvalidParameter interface {
	InvalidParameter()
}
```

---

### ErrNotFound

ErrNotFound signals that the requested object doesn't exist

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L4)  

```go
type ErrNotFound interface {
	NotFound()
}
```

---

### ErrNotImplemented

ErrNotImplemented signals that the requested action/feature is not implemented on the system as configured.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L47)  

```go
type ErrNotImplemented interface {
	NotImplemented()
}
```

---

### ErrNotModified

ErrNotModified signals that an action can't be performed because it's already in the desired state

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L42)  

```go
type ErrNotModified interface {
	NotModified()
}
```

---

### ErrSystem

ErrSystem signals that some internal error occurred.
An example of this would be a failed mount request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L37)  

```go
type ErrSystem interface {
	System()
}
```

---

### ErrUnauthorized

ErrUnauthorized is used to signify that the user is not authorized to perform a specific action

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L20)  

```go
type ErrUnauthorized interface {
	Unauthorized()
}
```

---

### ErrUnavailable

ErrUnavailable signals that the requested action/subsystem is not available.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L25)  

```go
type ErrUnavailable interface {
	Unavailable()
}
```

---

### ErrUnknown

ErrUnknown signals that the kind of error that occurred is not known.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/errdefs/defs.go#L52)  

```go
type ErrUnknown interface {
	Unknown()
}
```

---


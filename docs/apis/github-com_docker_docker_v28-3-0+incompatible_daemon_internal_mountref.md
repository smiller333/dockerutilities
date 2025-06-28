# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/internal/mountref

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:28:48 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Checker

Checker checks whether the provided path is mounted.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/mountref/counter.go#L18)  

```go
type Checker func(path string) bool
```

---

### Counter

Counter is a generic counter for use by graphdriver Get/Put calls

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/mountref/counter.go#L11)  

```go
type Counter struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewCounter

NewCounter returns a new Counter. It accepts a Checker to
determine whether a path is mounted.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/mountref/counter.go#L22)  

```go
func NewCounter(c Checker) *Counter
```

#### Methods

##### Counter.Decrement

Decrement decreases the ref count for the given id and returns the current count

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/mountref/counter.go#L37)  

```go
func (c *Counter) Decrement(path string) int
```

##### Counter.Increment

Increment increases the ref count for the given id and returns the current count

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/mountref/counter.go#L30)  

```go
func (c *Counter) Increment(path string) int
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/cleanups

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:31:03 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Composite

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/cleanups/composite.go#L9)  

```go
type Composite struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Composite.Add

Add adds a cleanup to be called.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/cleanups/composite.go#L14)  

```go
func (c *Composite) Add(f func(context.Context) error)
```

##### Composite.Call

Call calls all cleanups in reverse order and returns an error combining all
non-nil errors.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/cleanups/composite.go#L20)  

```go
func (c *Composite) Call(ctx context.Context) error
```

##### Composite.Release

Release removes all cleanups, turning Call into a no-op.
Caller still can call the cleanups by calling the returned function
which is equivalent to calling the Call before Release was called.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/cleanups/composite.go#L29)  

```go
func (c *Composite) Release() func(context.Context) error
```

---


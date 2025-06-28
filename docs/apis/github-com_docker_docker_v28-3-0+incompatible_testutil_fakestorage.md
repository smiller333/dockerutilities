# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil/fakestorage

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:28 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### SetTestEnvironment

SetTestEnvironment sets a static test environment
TODO: decouple this package from environment

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakestorage/storage.go#L38)  

```go
func SetTestEnvironment(env *environment.Execution)
```

---

## Types

### Fake

Fake is a static file server. It might be running locally or remotely
on test host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakestorage/storage.go#L30)  

```go
type Fake interface {
	Close() error
	URL() string
	CtxDir() string
}
```

#### Functions

##### New

New returns a static file server that is used as build context.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakestorage/storage.go#L43)  

```go
func New(t testing.TB, dir string, modifiers ...func(*fakecontext.Fake) error) Fake
```

---


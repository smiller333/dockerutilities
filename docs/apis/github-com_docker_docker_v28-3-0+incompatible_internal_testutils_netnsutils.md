# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/testutils/netnsutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:31:59 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AssertSocketSameNetNS

AssertSocketSameNetNS makes a best-effort attempt to assert that conn is in
the same network namespace as the current goroutine's thread.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/netnsutils/sanity_linux.go#L15)  

```go
func AssertSocketSameNetNS(t testing.TB, conn syscall.Conn)
```

---

### SetupTestOSContext

SetupTestOSContext joins the current goroutine to a new network namespace,
and returns its associated teardown function.

Example usage:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/netnsutils/context_unix.go#L32)  

```go
func SetupTestOSContext(t *testing.T) func()
```

---

## Types

### OSContext

OSContext is a handle to a test OS context.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/netnsutils/context_unix.go#L19)  

```go
type OSContext struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### SetupTestOSContextEx

SetupTestOSContextEx joins the current goroutine to a new network namespace.

Compared to SetupTestOSContext, this function allows goroutines to be
spawned which are associated with the same OS context via the returned
OSContext value.

Example usage:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/netnsutils/context_unix.go#L47)  

```go
func SetupTestOSContextEx(t *testing.T) *OSContext
```

#### Methods

##### OSContext.Cleanup

Cleanup tears down the OS context. It must be called from the same goroutine
as the SetupTestOSContextEx call which returned c.

Explicit cleanup is required as (*testing.T).Cleanup() makes no guarantees
about which goroutine the cleanup functions are invoked on.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/netnsutils/context_unix.go#L95)  

```go
func (c *OSContext) Cleanup(t *testing.T)
```

##### OSContext.Go

Go starts running fn in a new goroutine inside the test OS context.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/netnsutils/context_unix.go#L183)  

```go
func (c *OSContext) Go(t *testing.T, fn func())
```

##### OSContext.Set

Set sets the OS context of the calling goroutine to c and returns a teardown
function to restore the calling goroutine's OS context and release resources.
The teardown function accepts an optional Logger argument.

This is a lower-level interface which is less ergonomic than c.Go() but more
composable with other goroutine-spawning utilities such as sync.WaitGroup
or golang.org/x/sync/errgroup.Group.

Example usage:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/netnsutils/context_unix.go#L146)  

```go
func (c *OSContext) Set() (func(testutils.Logger), error)
```

---


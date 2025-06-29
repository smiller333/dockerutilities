# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:54 UTC

## Overview

Package testutil contains common testing tasks like running dockerd.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helpers.go#L29)

```go
var DevZero io.Reader = devZero{}
```

## Functions

### CheckNotParallel

CheckNotParallel checks if t.Parallel() was not called on the current test.
There's no public method to check this, so we use reflection to check the
internal field set by t.Parallel()
https://github.com/golang/go/blob/8e658eee9c7a67a8a79a8308695920ac9917566c/src/testing/testing.go#L1449

Since this is not a public API, it might change at any time.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helpers.go#L169)  

```go
func CheckNotParallel(t testing.TB)
```

---

### CleanupContext

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helpers.go#L159)  

```go
func CleanupContext(t *testing.T)
```

---

### ConfigureTracing

ConfigureTracing sets up an OTLP tracing exporter for use in tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helpers.go#L43)  

```go
func ConfigureTracing() func(context.Context)
```

---

### GenerateRandomAlphaOnlyString

GenerateRandomAlphaOnlyString generates an alphabetical random string with length n.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/stringutils.go#L6)  

```go
func GenerateRandomAlphaOnlyString(n int) string
```

---

### GetContext

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helpers.go#L151)  

```go
func GetContext(t TestingT) context.Context
```

---

### RunCommand

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helpers.go#L103)  

```go
func RunCommand(ctx context.Context, cmd string, args ...string) *icmd.Result
```

---

### SetContext

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helpers.go#L155)  

```go
func SetContext(t TestingT, ctx context.Context)
```

---

### StartSpan

StartSpan starts a span for the given test.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helpers.go#L91)  

```go
func StartSpan(ctx context.Context, t TestingT) context.Context
```

---

### TempDir

TempDir returns a temporary directory for use in tests.
t.TempDir() can't be used as the temporary directory returned by
that function cannot be accessed by the fake-root user for rootless
Docker. It creates a nested hierarchy of directories where the
outermost has permission 0700.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/temp_files.go#L14)  

```go
func TempDir(t *testing.T) string
```

---

## Types

### HelperT

HelperT is a subset of testing.T that implements the Helper function

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helper.go#L4)  

```go
type HelperT interface {
	Helper()
}
```

---

### TestingT

TestingT is an interface wrapper around *testing.T and *testing.B.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/helpers.go#L83)  

```go
type TestingT interface {
	Name() string
	Cleanup(func())
	Log(...any)
	Failed() bool
}
```

---


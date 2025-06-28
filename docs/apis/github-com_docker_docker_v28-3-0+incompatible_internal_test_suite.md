# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/test/suite

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:47 UTC

## Overview

Package suite is a simplified version of testify's suite package which has unnecessary dependencies.
Please remove this package whenever possible.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/test/suite/suite.go#L17)

```go
var TimeoutFlag = flag.Duration("timeout", 0, "DO NOT USE")
```

## Functions

### Run

Run takes a testing suite and runs all of the tests attached to it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/test/suite/suite.go#L22)  

```go
func Run(ctx context.Context, t *testing.T, suite interface{})
```

---

## Types

### SetupAllSuite

SetupAllSuite has a SetupSuite method, which will run before the
tests in the suite are run.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/test/suite/interfaces.go#L10)  

```go
type SetupAllSuite interface {
	SetUpSuite(context.Context, *testing.T)
}
```

---

### SetupTestSuite

SetupTestSuite has a SetupTest method, which will run before each
test in the suite.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/test/suite/interfaces.go#L16)  

```go
type SetupTestSuite interface {
	SetUpTest(context.Context, *testing.T)
}
```

---

### TearDownAllSuite

TearDownAllSuite has a TearDownSuite method, which will run after
all the tests in the suite have been run.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/test/suite/interfaces.go#L22)  

```go
type TearDownAllSuite interface {
	TearDownSuite(context.Context, *testing.T)
}
```

---

### TearDownTestSuite

TearDownTestSuite has a TearDownTest method, which will run after
each test in the suite.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/test/suite/interfaces.go#L28)  

```go
type TearDownTestSuite interface {
	TearDownTest(context.Context, *testing.T)
}
```

---

### TimeoutTestSuite

TimeoutTestSuite has a OnTimeout method, which will run after
a single test times out after a period specified by -timeout flag.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/test/suite/interfaces.go#L34)  

```go
type TimeoutTestSuite interface {
	OnTimeout()
}
```

---


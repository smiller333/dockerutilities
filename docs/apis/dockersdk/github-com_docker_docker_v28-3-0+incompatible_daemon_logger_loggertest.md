# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/loggertest

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:23 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Reader

Reader tests that a logger.LogReader implementation behaves as it should.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggertest/logreader.go#L35)  

```go
type Reader struct {
	// Factory returns a function which constructs loggers for the container
	// specified in info. Each call to the returned function must yield a
	// distinct logger instance which can read back logs written by earlier
	// instances.
	Factory func(*testing.T, logger.Info) func(*testing.T) logger.Logger
}
```

#### Methods

##### Reader.TestConcurrent

TestConcurrent tests the Logger and its LogReader implementation for
race conditions when logging from multiple goroutines concurrently.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggertest/logreader.go#L413)  

```go
func (tr Reader) TestConcurrent(t *testing.T)
```

##### Reader.TestFollow

TestFollow tests the LogReader's follow implementation.

The LogReader is expected to be able to follow an arbitrary number of
messages at a high rate with no dropped messages.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggertest/logreader.go#L196)  

```go
func (tr Reader) TestFollow(t *testing.T)
```

##### Reader.TestTail

TestTail tests the behavior of the LogReader's tail implementation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggertest/logreader.go#L54)  

```go
func (tr Reader) TestTail(t *testing.T)
```

---


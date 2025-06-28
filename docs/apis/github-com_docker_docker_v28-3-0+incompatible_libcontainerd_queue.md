# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libcontainerd/queue

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:24 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Queue

Queue is the structure used for holding functions in a queue.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/queue/queue.go#L6)  

```go
type Queue struct {
	sync.Mutex
	// contains filtered or unexported fields
}
```

#### Methods

##### Queue.Append

Append adds an item to a queue.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/queue/queue.go#L12)  

```go
func (q *Queue) Append(id string, f func())
```

---


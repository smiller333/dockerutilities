# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration/internal/system

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:06:27 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CheckGoroutineCount

CheckGoroutineCount returns a poll.Check that polls the daemon info API until the expected number of goroutines is hit.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/system/goroutines.go#L54)  

```go
func CheckGoroutineCount(ctx context.Context, apiClient client.SystemAPIClient, expected int) poll.Check
```

---

### StableGoroutineCount

StableGoroutineCount is a poll.Check that polls the daemon info API until the goroutine count is the same for 3 iterations.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/system/goroutines.go#L24)  

```go
func StableGoroutineCount(ctx context.Context, apiClient client.SystemAPIClient, count *int) poll.Check
```

---

### WaitForStableGoroutineCount

WaitForStableGoroutineCount polls the daemon Info API and returns the reported goroutine count
after multiple calls return the same number.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/system/goroutines.go#L13)  

```go
func WaitForStableGoroutineCount(ctx context.Context, t poll.TestingT, apiClient client.SystemAPIClient, opts ...poll.SettingOp) int
```

---

## Types

This section is empty.


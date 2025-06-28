# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/sliceutil

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:44 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Dedup

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/sliceutil/sliceutil.go#L6)  

```go
func Dedup[T comparable](slice []T) []T
```

---

### Map

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/sliceutil/sliceutil.go#L18)  

```go
func Map[S ~[]In, In, Out any](s S, fn func(In) Out) []Out
```

---

### Mapper

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/sliceutil/sliceutil.go#L26)  

```go
func Mapper[In, Out any](fn func(In) Out) func([]In) []Out
```

---

## Types

This section is empty.


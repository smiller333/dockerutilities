# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/names

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:37 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/names/names.go#L6)

```go
const RestrictedNameChars = `[a-zA-Z0-9][a-zA-Z0-9_.-]`
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/names/names.go#L9)

```go
var RestrictedNamePattern = lazyregexp.New(`^` + RestrictedNameChars + `+$`)
```

## Functions

This section is empty.

## Types

This section is empty.


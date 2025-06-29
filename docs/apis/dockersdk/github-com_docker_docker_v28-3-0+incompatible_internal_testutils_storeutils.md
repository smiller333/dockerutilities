# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/testutils/storeutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:08:01 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewTempStore

NewTempStore creates a new temporary libnetwork store for testing purposes.
The store is created in a temporary directory that is cleaned up when the
test finishes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/storeutils/store.go#L13)  

```go
func NewTempStore(t *testing.T) *datastore.Store
```

---

## Types

This section is empty.


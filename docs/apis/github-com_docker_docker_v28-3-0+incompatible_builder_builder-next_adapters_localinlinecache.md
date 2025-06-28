# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/builder-next/adapters/localinlinecache

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:42 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ResolveCacheImporterFunc

ResolveCacheImporterFunc returns a resolver function for local inline cache

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/localinlinecache/inlinecache.go#L26)  

```go
func ResolveCacheImporterFunc(sm *session.Manager, resolverFunc docker.RegistryHosts, cs content.Store, rs reference.Store, is imagestore.Store) remotecache.ResolveCacheImporterFunc
```

---

## Types

This section is empty.


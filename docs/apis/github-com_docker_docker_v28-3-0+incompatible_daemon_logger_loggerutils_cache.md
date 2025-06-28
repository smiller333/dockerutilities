# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/loggerutils/cache

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:31 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/cache/local_cache.go#L15)

```go
const (
	// DriverName is the name of the driver used for local log caching
	DriverName = local.Name
)
```

## Variables

This section is empty.

## Functions

### MergeDefaultLogConfig

MergeDefaultLogConfig reads the default log opts and makes sure that any caching related keys that exist there are
added to dst.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/cache/validate.go#L31)  

```go
func MergeDefaultLogConfig(dst, defaults map[string]string)
```

---

### ShouldUseCache

ShouldUseCache reads the log opts to determine if caching should be enabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/cache/local_cache.go#L103)  

```go
func ShouldUseCache(cfg map[string]string) bool
```

---

### WithLocalCache

WithLocalCache wraps the passed in logger with a logger caches all writes locally
in addition to writing to the passed in logger.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/loggerutils/cache/local_cache.go#L31)  

```go
func WithLocalCache(l logger.Logger, info logger.Info) (logger.Logger, error)
```

---

## Types

This section is empty.


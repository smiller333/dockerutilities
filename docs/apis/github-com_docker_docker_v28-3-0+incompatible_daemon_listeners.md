# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/listeners

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:49 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Init

Init creates new listeners for the server.
TODO: Clean up the fact that socketGroup and tlsConfig aren't always used.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/listeners/listeners_linux.go#L19)  

```go
func Init(proto, addr, socketGroup string, tlsConfig *tls.Config) ([]net.Listener, error)
```

---

## Types

This section is empty.


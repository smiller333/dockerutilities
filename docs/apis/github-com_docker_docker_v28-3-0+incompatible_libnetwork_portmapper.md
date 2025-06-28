# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/portmapper

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:03 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### StartProxy

StartProxy starts the proxy process at proxyPath.
If listenSock is not nil, it must be a bound socket that can be passed to
the proxy process for it to listen on.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/portmapper/proxy_linux.go#L20)  

```go
func StartProxy(pb types.PortBinding,
	proxyPath string,
	listenSock *os.File,
) (stop func() error, retErr error)
```

---

## Types

This section is empty.


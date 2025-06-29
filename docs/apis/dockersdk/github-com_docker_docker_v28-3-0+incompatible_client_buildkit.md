# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/client/buildkit

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:13 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ClientOpts

ClientOpts returns a list of buildkit client options which allows the
caller to create a buildkit client which will connect to the buildkit
API provided by the daemon. These options can be passed to bkclient.New.

Example:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/buildkit/buildkit.go#L18)  

```go
func ClientOpts(c client.HijackDialer) []bkclient.ClientOpt
```

---

## Types

This section is empty.


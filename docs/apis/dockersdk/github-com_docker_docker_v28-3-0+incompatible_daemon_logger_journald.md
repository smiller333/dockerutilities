# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/journald

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:05 UTC

## Overview

Package journald provides the log driver for forwarding server logs
to endpoints that receive the systemd format.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New creates a journald logger using the configuration passed in on
the context.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/journald/journald.go#L101)  

```go
func New(info logger.Info) (logger.Logger, error)
```

---

## Types

This section is empty.


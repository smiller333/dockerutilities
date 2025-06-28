# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/fluentd

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:02 UTC

## Overview

Package fluentd provides the log driver for forwarding server logs
to fluentd endpoints.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New creates a fluentd logger using the configuration passed in on
the context. The supported context configuration variable is
fluentd-address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/fluentd/fluentd.go#L80)  

```go
func New(info logger.Info) (logger.Logger, error)
```

---

### ValidateLogOpt

ValidateLogOpt looks for fluentd specific log option fluentd-address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/fluentd/fluentd.go#L147)  

```go
func ValidateLogOpt(cfg map[string]string) error
```

---

## Types

This section is empty.


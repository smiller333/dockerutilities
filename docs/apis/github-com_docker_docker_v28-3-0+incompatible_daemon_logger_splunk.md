# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/splunk

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:30 UTC

## Overview

Package splunk provides the log driver for forwarding server logs to
Splunk HTTP Event Collector endpoint.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New creates splunk logger driver using configuration passed in context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/splunk/splunk.go#L153)  

```go
func New(info logger.Info) (logger.Logger, error)
```

---

### ValidateLogOpt

ValidateLogOpt looks for all supported by splunk driver options

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/splunk/splunk.go#L553)  

```go
func ValidateLogOpt(cfg map[string]string) error
```

---

## Types

This section is empty.


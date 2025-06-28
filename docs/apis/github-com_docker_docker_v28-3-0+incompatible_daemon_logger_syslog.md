# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/syslog

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:35 UTC

## Overview

Package syslog provides the logdriver for forwarding server logs to syslog endpoints.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New creates a syslog logger using the configuration passed in on
the context. Supported context configuration variables are
syslog-address, syslog-facility, syslog-format.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/syslog/syslog.go#L88)  

```go
func New(info logger.Info) (logger.Logger, error)
```

---

### ValidateLogOpt

ValidateLogOpt looks for syslog specific log options
syslog-address, syslog-facility.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/syslog/syslog.go#L188)  
**Added in:** v1.8.0

```go
func ValidateLogOpt(cfg map[string]string) error
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/gcplogs

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:04 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New creates a new logger that logs to Google Cloud Logging using the application
default credentials.

See https://developers.google.com/identity/protocols/application-default-credentials

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/gcplogs/gcplogging.go#L108)  

```go
func New(info logger.Info) (logger.Logger, error)
```

---

### ValidateLogOpts

ValidateLogOpts validates the opts passed to the gcplogs driver. Currently, the gcplogs
driver doesn't take any arguments.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/gcplogs/gcplogging.go#L205)  

```go
func ValidateLogOpts(cfg map[string]string) error
```

---

## Types

This section is empty.


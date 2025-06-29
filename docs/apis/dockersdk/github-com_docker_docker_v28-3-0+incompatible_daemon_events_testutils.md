# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/events/testutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:04 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Scan

Scan turns an event string like the default ones formatted in the cli output
and turns it into an event message.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events/testutils/testutils.go#L42)  

```go
func Scan(text string) (*events.Message, error)
```

---

### ScanMap

ScanMap turns an event string like the default ones formatted in the cli output
and turns it into map.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events/testutils/testutils.go#L26)  

```go
func ScanMap(text string) map[string]string
```

---

## Types

This section is empty.


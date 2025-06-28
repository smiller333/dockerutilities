# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/time

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:27 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GetTimestamp

GetTimestamp tries to parse given string as golang duration,
then RFC3339 time and finally as a Unix timestamp. If
any of these were successful, it returns a Unix timestamp
as string otherwise returns the given value back.
In case of duration input, the returned timestamp is computed
as the given reference time minus the amount of the duration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/time/timestamp.go#L26)  

```go
func GetTimestamp(value string, reference time.Time) (string, error)
```

---

### ParseTimestamps

ParseTimestamps returns seconds and nanoseconds from a timestamp that has
the format ("%d.%09d", time.Unix(), int64(time.Nanosecond())).
If the incoming nanosecond portion is longer than 9 digits it is truncated.
The expectation is that the seconds and nanoseconds will be used to create a
time variable.  For example:

returns seconds as defaultSeconds if value == ""

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/time/timestamp.go#L108)  

```go
func ParseTimestamps(value string, defaultSeconds int64) (seconds int64, nanoseconds int64, _ error)
```

---

## Types

This section is empty.


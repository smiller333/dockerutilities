# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/parsers

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:42 UTC

## Overview

Package parsers provides helper functions to parse and validate different type
of string. It can be hosts, unix addresses, tcp addresses, filters, kernel
operating system versions.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ParseKeyValueOpt ⚠️ **DEPRECATED**

ParseKeyValueOpt parses and validates the specified string as a key/value
pair (key=value).

Deprecated: use strings.Cut instead. This utility was only used internally, and will be removed in the next release.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/parsers.go#L16)  

```go
func ParseKeyValueOpt(opt string) (key string, value string, _ error)
```

---

### ParseUintList ⚠️ **DEPRECATED**

ParseUintList parses and validates the specified string as the value
found in some cgroup file (e.g. `cpuset.cpus`, `cpuset.mems`), which could be
one of the formats below. Note that duplicates are actually allowed in the
input string. It returns a `map[int]bool` with available elements from `val`
set to `true`.
Supported formats:

Deprecated: ParseUintList was only used internally and will be removed in the next release.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/parsers.go#L61)  
**Added in:** v1.9.0

```go
func ParseUintList(val string) (map[int]bool, error)
```

---

### ParseUintListMaximum ⚠️ **DEPRECATED**

ParseUintListMaximum parses and validates the specified string as the value
found in some cgroup file (e.g. `cpuset.cpus`, `cpuset.mems`), which could be
one of the formats below. Note that duplicates are actually allowed in the
input string. It returns a `map[int]bool` with available elements from `val`
set to `true`. Values larger than `maximum` cause an error if max is non zero,
in order to stop the map becoming excessively large.
Supported formats:

Deprecated: ParseUintListMaximum was only used internally and will be removed in the next release.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/parsers.go#L41)  

```go
func ParseUintListMaximum(val string, maximum int) (map[int]bool, error)
```

---

## Types

This section is empty.


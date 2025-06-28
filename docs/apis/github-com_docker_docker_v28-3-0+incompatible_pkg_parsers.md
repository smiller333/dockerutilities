# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/parsers

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:48 UTC

## Overview

Package parsers provides helper functions to parse and validate different type
of string. It can be hosts, unix addresses, tcp addresses, filters, kernel
operating system versions.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ParseKeyValueOpt

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/parsers.go#L16)  

```go
func ParseKeyValueOpt(opt string) (key string, value string, _ error)
```

---

### ParseUintList

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/parsers.go#L61)  
**Added in:** v1.9.0

```go
func ParseUintList(val string) (map[int]bool, error)
```

---

### ParseUintListMaximum

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/parsers/parsers.go#L41)  

```go
func ParseUintListMaximum(val string, maximum int) (map[int]bool, error)
```

---

## Types

This section is empty.


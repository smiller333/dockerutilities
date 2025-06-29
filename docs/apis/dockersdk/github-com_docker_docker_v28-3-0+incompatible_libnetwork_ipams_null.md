# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/ipams/null

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:10:20 UTC

## Overview

Package null implements the null ipam driver. Null ipam driver satisfies ipamapi contract,
but does not effectively reserve/allocate any address pool or address


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/null/null.go#L13)

```go
const (
	// DriverName is the name of the built-in null ipam driver
	DriverName = "null"
)
```

## Variables

This section is empty.

## Functions

### Register

Register registers the null ipam driver with r.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/null/null.go#L80)  

```go
func Register(r ipamapi.Registerer) error
```

---

## Types

This section is empty.


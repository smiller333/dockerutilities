# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/internal/capabilities

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:37 UTC

## Overview

Package capabilities allows to generically handle capabilities.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Set

Set represents a set of capabilities.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/capabilities/caps.go#L5)  

```go
type Set map[string]struct{}
```

#### Methods

##### Set.Match

Match tries to match set with caps, which is an OR list of AND lists of capabilities.
The matched AND list of capabilities is returned; or nil if none are matched.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/capabilities/caps.go#L9)  

```go
func (set Set) Match(caps [][]string) []string
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/stringid

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:35 UTC

## Overview

Package stringid provides helper functions for dealing with string identifiers


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GenerateRandomID

GenerateRandomID returns a unique, 64-character ID consisting of a-z, 0-9.
It guarantees that the ID, when truncated (TruncateID) does not consist
of numbers only, so that the truncated ID can be used as hostname for
containers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/stringid/stringid.go#L33)  

```go
func GenerateRandomID() string
```

---

### TruncateID

TruncateID returns a shorthand version of a string identifier for convenience.
A collision with other shorthands is very unlikely, but possible.
In case of a collision a lookup with TruncIndex.Get() will fail, and the caller
will need to use a longer prefix, or the full-length Id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/stringid/stringid.go#L19)  

```go
func TruncateID(id string) string
```

---

## Types

This section is empty.


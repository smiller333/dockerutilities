# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/namesgenerator

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:46 UTC

## Overview

Package namesgenerator generates random names.

This package is officially "frozen" - no new additions will be accepted.

For a long time, this package provided a lot of joy within the project, but
at some point the conflicts of opinion became greater than the added joy.

At some future time, this may be replaced with something that sparks less
controversy, but for now it will remain as-is.

See also https://github.com/moby/moby/pull/43210#issuecomment-1029934277


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GetRandomName

GetRandomName generates a random name from the list of adjectives and surnames in this package
formatted as "adjective_surname". For example 'focused_turing'. If retry is non-zero, a random
integer between 0 and 10 will be added to the end of the name, e.g `focused_turing3`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/namesgenerator/names-generator.go#L849)  
**Added in:** v0.12.0

```go
func GetRandomName(retry int) string
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/useragent

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:25 UTC

## Overview

Package useragent provides helper functions to pack
version information into a single User-Agent header.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AppendVersions

AppendVersions converts versions to a string and appends the string to the string base.

Each VersionInfo will be converted to a string in the format of
"product/version", where the "product" is get from the name field, while
version is get from the version field. Several pieces of version information
will be concatenated and separated by space.

Example:
AppendVersions("base", VersionInfo{"foo", "1.0"}, VersionInfo{"bar", "2.0"})
results in "base foo/1.0 bar/2.0".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/useragent/useragent.go#L38)  

```go
func AppendVersions(base string, versions ...VersionInfo) string
```

---

## Types

### VersionInfo

VersionInfo is used to model UserAgent versions.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/useragent/useragent.go#L10)  

```go
type VersionInfo struct {
	Name    string
	Version string
}
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/builder-next/exporter/overrides

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:49 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### SanitizeRepoAndTags

SanitizeRepoAndTags parses the raw names to a slice of repoAndTag.
It removes duplicates and validates each repoName and tag to not contain a digest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/exporter/overrides/overrides.go#L11)  

```go
func SanitizeRepoAndTags(names []string) (repoAndTags []string, _ error)
```

---

## Types

This section is empty.


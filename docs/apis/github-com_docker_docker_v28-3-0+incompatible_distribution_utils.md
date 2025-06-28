# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/distribution/utils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:56 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### WriteDistributionProgress

WriteDistributionProgress is a helper for writing progress from chan to JSON
stream with an optional cancel function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/utils/progress.go#L16)  

```go
func WriteDistributionProgress(cancelFunc func(), outStream io.Writer, progressChan <-chan progress.Progress)
```

---

## Types

This section is empty.


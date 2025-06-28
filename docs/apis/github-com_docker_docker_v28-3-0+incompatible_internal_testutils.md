# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/testutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:31:55 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### UncompressedTarDigest

UncompressedTarDigest returns the canonical digest of the uncompressed tar stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/archive.go#L11)  

```go
func UncompressedTarDigest(compressedTar io.Reader) (digest.Digest, error)
```

---

## Types

### Logger

Logger is used to log non-fatal messages during tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/logger.go#L6)  

```go
type Logger interface {
	Logf(format string, args ...any)
}
```

---


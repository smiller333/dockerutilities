# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration/internal/build

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:18 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Do

Do builds an image from the given context and returns the image ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/build/build.go#L18)  

```go
func Do(ctx context.Context, t *testing.T, client client.APIClient, buildCtx *fakecontext.Fake) string
```

---

### GetImageIDFromBody

GetImageIDFromBody reads the image ID from the build response body.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/build/build.go#L32)  

```go
func GetImageIDFromBody(t *testing.T, body io.Reader) string
```

---

## Types

This section is empty.


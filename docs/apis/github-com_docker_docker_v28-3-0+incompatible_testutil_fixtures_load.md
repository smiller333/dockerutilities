# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil/fixtures/load

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:16:08 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### FrozenImagesLinux

FrozenImagesLinux loads the frozen image set for the integration suite
If the images are not available locally it will download them
TODO: This loads whatever is in the frozen image dir, regardless of what
images were passed in. If the images need to be downloaded, then it will respect
the passed in images

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fixtures/load/frozen.go#L31)  

```go
func FrozenImagesLinux(ctx context.Context, client client.APIClient, images ...string) error
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/builder-next/exporter

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:47 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/exporter/exporter.go#L3)

```go
const Moby = "moby"
```

## Variables

This section is empty.

## Functions

### NewWrapper

NewWrapper returns an exporter wrapper that applies moby specific attributes
and hooks the export process.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/exporter/wrapper.go#L34)  

```go
func NewWrapper(exp exporter.Exporter, callbacks BuildkitCallbacks) (exporter.Exporter, error)
```

---

## Types

### BuildkitCallbacks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/exporter/wrapper.go#L16)  

```go
type BuildkitCallbacks struct {
	// Exported is a Called when an image is exported by buildkit.
	Exported func(ctx context.Context, id string, desc ocispec.Descriptor)

	// Named is a callback that is called when an image is created in the
	// containerd image store by buildkit.
	Named func(ctx context.Context, ref reference.NamedTagged, desc ocispec.Descriptor)
}
```

---


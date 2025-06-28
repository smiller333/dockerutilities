# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/builder-next/exporter/mobyexporter

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:47 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New creates a new moby imagestore exporter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/exporter/mobyexporter/export.go#L48)  

```go
func New(opt Opt) (exporter.Exporter, error)
```

---

## Types

### Differ

Differ can make a moby layer from a snapshot

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/exporter/mobyexporter/export.go#L25)  

```go
type Differ interface {
	EnsureLayer(ctx context.Context, key string) ([]layer.DiffID, error)
}
```

---

### ImageTagger

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/exporter/mobyexporter/export.go#L29)  

```go
type ImageTagger interface {
	TagImage(ctx context.Context, imageID image.ID, newTag distref.Named) error
}
```

---

### Opt

Opt defines a struct for creating new exporter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/exporter/mobyexporter/export.go#L34)  

```go
type Opt struct {
	ImageStore            image.Store
	Differ                Differ
	ImageTagger           ImageTagger
	ContentStore          content.Store
	LeaseManager          leases.Manager
	ImageExportedCallback func(ctx context.Context, id string, desc ocispec.Descriptor)
}
```

---


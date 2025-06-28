# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/builder-next/imagerefchecker

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:51 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New creates new image reference checker that can be used to see if a reference
is being used by any of the images in the image store

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/imagerefchecker/checker.go#L25)  

```go
func New(opt Opt) cache.ExternalRefCheckerFunc
```

---

## Types

### LayerGetter

LayerGetter abstracts away the snapshotter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/imagerefchecker/checker.go#L13)  

```go
type LayerGetter interface {
	GetLayer(string) (layer.Layer, error)
}
```

---

### Opt

Opt represents the options needed to create a refchecker

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/imagerefchecker/checker.go#L18)  

```go
type Opt struct {
	LayerGetter LayerGetter
	ImageStore  image.Store
}
```

---


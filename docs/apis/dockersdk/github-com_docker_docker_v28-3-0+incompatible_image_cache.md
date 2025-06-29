# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/image/cache

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:06:03 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/cache/cache.go#L29)  

```go
func New(ctx context.Context, store ImageCacheStore, cacheFrom []string) (builder.ImageCache, error)
```

---

## Types

### ImageCache

ImageCache is cache based on history objects. Requires initial set of images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/cache/cache.go#L66)  

```go
type ImageCache struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### ImageCache.GetCache

GetCache returns the image id found in the cache

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/cache/cache.go#L78)  

```go
func (ic *ImageCache) GetCache(parentID string, cfg *containertypes.Config, platform ocispec.Platform) (string, error)
```

##### ImageCache.Populate

Populate adds an image to the cache (to be queried later)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/cache/cache.go#L73)  

```go
func (ic *ImageCache) Populate(image *image.Image)
```

---

### ImageCacheStore

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/cache/cache.go#L19)  

```go
type ImageCacheStore interface {
	Get(image.ID) (*image.Image, error)
	GetByRef(ctx context.Context, refOrId string) (*image.Image, error)
	SetParent(target, parent image.ID) error
	GetParent(target image.ID) (image.ID, error)
	Create(parent *image.Image, image image.Image, extraLayer layer.DiffID) (image.ID, error)
	IsBuiltLocally(id image.ID) (bool, error)
	Children(id image.ID) []image.ID
}
```

---

### LocalImageCache

LocalImageCache is cache based on parent chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/cache/cache.go#L56)  

```go
type LocalImageCache struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### LocalImageCache.GetCache

GetCache returns the image id found in the cache

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/cache/cache.go#L61)  

```go
func (lic *LocalImageCache) GetCache(imgID string, config *containertypes.Config, platform ocispec.Platform) (string, error)
```

---


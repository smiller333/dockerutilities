# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/testutils/specialimage

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:04 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ConfigTarget

ConfigTarget creates an image index with an image config being used as an
image target instead of a manifest or index.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/configtarget.go#L11)  

```go
func ConfigTarget(dir string) (*ocispec.Index, error)
```

---

### Dangling

Dangling creates an image with no layers and no tag.
It also has an extra org.mobyproject.test.specialimage=1 label set.
Layout: OCI.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/dangling.go#L19)  

```go
func Dangling(dir string) (*ocispec.Index, error)
```

---

### EmptyFS

EmptyFS builds an image with an empty rootfs.
Layout: Legacy Docker Archive
See https://github.com/docker/docker/pull/5262
and also https://github.com/docker/docker/issues/4242

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/emptyfs.go#L15)  

```go
func EmptyFS(dir string) (*ocispec.Index, error)
```

---

### EmptyIndex

EmptyIndex creates an image index with no manifests.
This is equivalent to `tianon/scratch:index`.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/emptyindex.go#L11)  

```go
func EmptyIndex(dir string) (*ocispec.Index, error)
```

---

### LegacyManifest

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/random.go#L90)  

```go
func LegacyManifest(dir string, imageRef string, mfstDesc ocispec.Descriptor) error
```

---

### Load

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/load.go#L21)  

```go
func Load(ctx context.Context, t *testing.T, apiClient client.APIClient, imageFunc SpecialImageFunc) string
```

---

### MultiLayer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/multilayer.go#L24)  

```go
func MultiLayer(dir string) (*ocispec.Index, error)
```

---

### MultiLayerCustom

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/multilayer.go#L32)  

```go
func MultiLayerCustom(dir string, imageRef string, layers []SingleFileLayer) (*ocispec.Index, error)
```

---

### MultiPlatform

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/multiplatform.go#L10)  

```go
func MultiPlatform(dir string, imageRef string, imagePlatforms []ocispec.Platform) (*ocispec.Index, []ocispec.Descriptor, error)
```

---

### PartialMultiPlatform

PartialMultiPlatform creates an index with all platforms in storedPlatforms
and missingPlatforms. However, only the blobs of the storedPlatforms are
created and stored, while the missingPlatforms are only referenced in the
index.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/partial.go#L20)  

```go
func PartialMultiPlatform(dir string, imageRef string, opts PartialOpts) (*ocispec.Index, []ocispec.Descriptor, error)
```

---

### RandomSinglePlatform

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/random.go#L15)  

```go
func RandomSinglePlatform(dir string, platform ocispec.Platform, source rand.Source) (*ocispec.Index, error)
```

---

### TextPlain

TextPlain creates an non-container image that only contains a text/plain blob.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/textplain.go#L11)  

```go
func TextPlain(dir string) (*ocispec.Index, error)
```

---

### TwoPlatform

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/twoplatform.go#L14)  

```go
func TwoPlatform(dir string) (*ocispec.Index, error)
```

---

## Types

### FileInLayer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/twoplatform.go#L38)  

```go
type FileInLayer struct {
	Path    string
	Content []byte
}
```

---

### PartialOpts

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/partial.go#L11)  

```go
type PartialOpts struct {
	Stored  []ocispec.Platform
	Missing []ocispec.Platform
}
```

---

### SingleFileLayer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/multilayer.go#L19)  

```go
type SingleFileLayer struct {
	Name    string
	Content []byte
}
```

---

### SpecialImageFunc

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/specialimage/load.go#L19)  

```go
type SpecialImageFunc func(string) (*ocispec.Index, error)
```

---


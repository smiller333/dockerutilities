# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/builder-next/adapters/containerimage

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:39 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Source

Source is the source implementation for accessing container images

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/containerimage/pull.go#L73)  

```go
type Source struct {
	SourceOpt
	// contains filtered or unexported fields
}
```

#### Functions

##### NewSource

NewSource creates a new image source

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/containerimage/pull.go#L79)  

```go
func NewSource(opt SourceOpt) (*Source, error)
```

#### Methods

##### Source.Identifier

Identifier constructs an Identifier from the given scheme, ref, and attrs,
all of which come from a SourceOp.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/containerimage/pull.go#L91)  

```go
func (is *Source) Identifier(scheme, ref string, attrs map[string]string, platform *pb.Platform) (source.Identifier, error)
```

##### Source.Resolve

Resolve returns access to pulling for an identifier

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/containerimage/pull.go#L249)  

```go
func (is *Source) Resolve(ctx context.Context, id source.Identifier, sm *session.Manager, vtx solver.Vertex) (source.SourceInstance, error)
```

##### Source.ResolveImageConfig

ResolveImageConfig returns image config for an image

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/containerimage/pull.go#L200)  

```go
func (is *Source) ResolveImageConfig(ctx context.Context, ref string, opt sourceresolver.Opt, sm *session.Manager, g session.Group) (digest.Digest, []byte, error)
```

##### Source.Schemes

Schemes returns a list of SourceOp identifier schemes that this source
should match.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/containerimage/pull.go#L85)  

```go
func (is *Source) Schemes() []string
```

---

### SourceOpt

SourceOpt is options for creating the image source

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/containerimage/pull.go#L59)  

```go
type SourceOpt struct {
	ContentStore    content.Store
	CacheAccessor   cache.Accessor
	ReferenceStore  reference.Store
	DownloadManager *xfer.LayerDownloadManager
	MetadataStore   metadata.V2MetadataService
	ImageStore      image.Store
	RegistryHosts   docker.RegistryHosts
	LayerStore      layer.Store
	LeaseManager    leases.Manager
	GarbageCollect  func(ctx context.Context) (gc.Stats, error)
}
```

---


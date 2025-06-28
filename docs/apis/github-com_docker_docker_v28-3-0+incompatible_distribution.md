# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/distribution

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:51 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### DeprecatedSchema1ImageError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/errors.go#L228)  

```go
func DeprecatedSchema1ImageError(ref reference.Named) error
```

---

### GetRepositories

GetRepositories returns a list of repositories configured for the given
reference. Multiple repositories can be returned if the reference is for
the default (Docker Hub) registry and a mirror is configured, but it omits
registries that were not reachable (pinging the /v2/ endpoint failed).

It returns an error if it was unable to reach any of the registries for
the given reference, or if the provided reference is invalid.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/repository.go#L19)  

```go
func GetRepositories(ctx context.Context, ref reference.Named, config *ImagePullConfig) ([]distribution.Repository, error)
```

---

### Pull

Pull initiates a pull operation. image is the repository name to pull, and
tag may be either empty, or indicate a specific tag to pull.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/pull.go#L19)  

```go
func Pull(ctx context.Context, ref reference.Named, config *ImagePullConfig, local ContentStore) error
```

---

### Push

Push initiates a push operation on ref. ref is the specific variant of the
image to push. If no tag is provided, all tags are pushed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/push.go#L21)  

```go
func Push(ctx context.Context, ref reference.Named, config *ImagePushConfig) error
```

---

### Tags

Tags returns available tags for the given image in the remote repository.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/pull.go#L33)  

```go
func Tags(ctx context.Context, ref reference.Named, config *Config) ([]string, error)
```

---

## Types

### AIModelNotSupportedError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/errors.go#L189)  

```go
type AIModelNotSupportedError struct{}
```

#### Methods

##### AIModelNotSupportedError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/errors.go#L191)  

```go
func (e AIModelNotSupportedError) Error() string
```

##### AIModelNotSupportedError.InvalidParameter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/errors.go#L195)  

```go
func (e AIModelNotSupportedError) InvalidParameter()
```

---

### Config

Config stores configuration for communicating
with a registry.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/config.go#L28)  
**Added in:** v1.13.0

```go
type Config struct {
	// MetaHeaders stores HTTP headers with metadata about the image
	MetaHeaders map[string][]string
	// AuthConfig holds authentication credentials for authenticating with
	// the registry.
	AuthConfig *registry.AuthConfig
	// ProgressOutput is the interface for showing the status of the pull
	// operation.
	ProgressOutput progress.Output
	// RegistryService is the registry service to use for TLS configuration
	// and endpoint lookup.
	RegistryService RegistryResolver
	// ImageEventLogger notifies events for a given image
	ImageEventLogger func(ctx context.Context, id, name string, action events.Action)
	// MetadataStore is the storage backend for distribution-specific
	// metadata.
	MetadataStore metadata.Store
	// ImageStore manages images.
	ImageStore ImageConfigStore
	// ReferenceStore manages tags. This value is optional, when excluded
	// content will not be tagged.
	ReferenceStore refstore.Store
}
```

---

### ContentStore

ContentStore is the interface used to persist registry blobs

Currently this is only used to persist manifests and manifest lists.
It is exported because `distribution.Pull` takes one as an argument.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/manifest.go#L44)  

```go
type ContentStore interface {
	content.Ingester
	content.Provider
	Info(ctx context.Context, dgst digest.Digest) (content.Info, error)
	Abort(ctx context.Context, ref string) error
	Update(ctx context.Context, info content.Info, fieldpaths ...string) (content.Info, error)
}
```

---

### ImageConfigStore

ImageConfigStore handles storing and getting image configurations
by digest. Allows getting an image configurations rootfs from the
configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/config.go#L89)  
**Added in:** v1.13.0

```go
type ImageConfigStore interface {
	Put(context.Context, []byte) (digest.Digest, error)
	Get(context.Context, digest.Digest) ([]byte, error)
}
```

#### Functions

##### NewImageConfigStoreFromStore

NewImageConfigStoreFromStore returns an ImageConfigStore backed
by an image.Store for container images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/config.go#L117)  
**Added in:** v1.13.0

```go
func NewImageConfigStoreFromStore(is image.Store) ImageConfigStore
```

---

### ImagePullConfig

ImagePullConfig stores pull configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/config.go#L53)  

```go
type ImagePullConfig struct {
	Config

	// DownloadManager manages concurrent pulls.
	DownloadManager *xfer.LayerDownloadManager
	// Schema2Types is an optional list of valid schema2 configuration types
	// allowed by the pull operation. If omitted, the default list of accepted
	// types is used.
	Schema2Types []string
	// Platform is the requested platform of the image being pulled
	Platform *ocispec.Platform
}
```

---

### ImagePushConfig

ImagePushConfig stores push configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/config.go#L67)  

```go
type ImagePushConfig struct {
	Config

	// ConfigMediaType is the configuration media type for
	// schema2 manifests.
	ConfigMediaType string
	// LayerStores manages layers.
	LayerStores PushLayerProvider
	// UploadManager dispatches uploads.
	UploadManager *xfer.LayerUploadManager
}
```

---

### PushLayer

PushLayer is a pushable layer with metadata about the layer
and access to the content of the layer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/config.go#L101)  
**Added in:** v1.13.0

```go
type PushLayer interface {
	ChainID() layer.ChainID
	DiffID() layer.DiffID
	Parent() PushLayer
	Open() (io.ReadCloser, error)
	Size() int64
	MediaType() string
	Release()
}
```

---

### PushLayerProvider

PushLayerProvider provides layers to be pushed by ChainID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/config.go#L95)  
**Added in:** v1.13.0

```go
type PushLayerProvider interface {
	Get(layer.ChainID) (PushLayer, error)
}
```

#### Functions

##### NewLayerProvidersFromStore

NewLayerProvidersFromStore returns layer providers backed by
an instance of LayerStore. Only getting layers as gzipped
tars is supported.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/config.go#L172)  

```go
func NewLayerProvidersFromStore(ls layer.Store) PushLayerProvider
```

---

### RegistryResolver

RegistryResolver is used for TLS configuration and endpoint lookup.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/config.go#L80)  

```go
type RegistryResolver interface {
	ResolveAuthConfig(map[string]registry.AuthConfig, reference.Named) registry.AuthConfig
	LookupPushEndpoints(hostname string) (endpoints []registrypkg.APIEndpoint, err error)
	LookupPullEndpoints(hostname string) (endpoints []registrypkg.APIEndpoint, err error)
}
```

---


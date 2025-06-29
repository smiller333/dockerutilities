# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/image

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:06:01 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/rootfs.go#L12)

```go
const TypeLayers = "layers"
```

## Variables

This section is empty.

## Functions

### CheckOS

CheckOS checks if the given OS matches the host's platform, and
returns an error otherwise.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image_os.go#L13)  

```go
func CheckOS(os string) error
```

---

## Types

### ChildConfig

ChildConfig is the configuration to apply to an Image to create a new
Child image. Other properties of the image are copied from the parent.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L202)  

```go
type ChildConfig struct {
	ContainerID     string
	Author          string
	Comment         string
	DiffID          layer.DiffID
	ContainerConfig *container.Config
	Config          *container.Config
}
```

---

### Details

Details provides additional image data

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L117)  

```go
type Details struct {
	// ManifestDescriptor is the descriptor of the platform-specific manifest
	// chosen by the [GetImage] call that returned this image.
	// The exact descriptor depends on the [GetImageOpts.Platform] field
	// passed to [GetImage] and the content availability.
	// This is only set by the containerd image service.
	ManifestDescriptor *ocispec.Descriptor
}
```

---

### DigestWalkFunc

DigestWalkFunc is function called by StoreBackend.Walk

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/fs.go#L17)  
**Added in:** v1.13.0

```go
type DigestWalkFunc func(id digest.Digest) error
```

---

### Exporter

Exporter provides interface for loading and saving images

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L282)  
**Added in:** v1.10.0

```go
type Exporter interface {
	Load(context.Context, io.ReadCloser, io.Writer, bool) error
	// TODO: Load(net.Context, io.ReadCloser, <- chan StatusMessage) error
	Save(context.Context, []string, io.Writer) error
}
```

---

### History

History stores build commands that were used to create an image

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L266)  
**Added in:** v1.10.0

```go
type History = ocispec.History
```

#### Functions

##### NewHistory

NewHistory creates a new history struct from arguments, and sets the created
time to the current time in UTC

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L270)  

```go
func NewHistory(author, comment, createdBy string, isEmptyLayer bool) History
```

---

### ID

ID is the content-addressable ID of an image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L20)  
**Added in:** v1.10.0

```go
type ID digest.Digest
```

#### Methods

##### ID.Digest

Digest converts ID into a digest

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L27)  
**Added in:** v1.13.0

```go
func (id ID) Digest() digest.Digest
```

##### ID.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L22)  
**Added in:** v1.10.0

```go
func (id ID) String() string
```

---

### Image

Image stores the image configuration

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L85)  

```go
type Image struct {
	V1Image

	// Parent is the ID of the parent image.
	//
	// Depending on how the image was created, this field may be empty and
	// is only set for images that were built/created locally. This field
	// is empty if the image was pulled from an image registry.
	Parent ID `json:"parent,omitempty"` //nolint:govet

	// RootFS contains information about the image's RootFS, including the
	// layer IDs.
	RootFS  *RootFS   `json:"rootfs,omitempty"`
	History []History `json:"history,omitempty"`

	// OsVersion is the version of the Operating System the image is built to
	// run on (especially for Windows).
	OSVersion  string   `json:"os.version,omitempty"`
	OSFeatures []string `json:"os.features,omitempty"`

	// Details holds additional details about image
	Details *Details `json:"-"`
	// contains filtered or unexported fields
}
```

#### Functions

##### Clone

Clone clones an image and changes ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L257)  

```go
func Clone(base *Image, id ID) *Image
```

##### NewChildImage

NewChildImage creates a new Image as a child of this image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L219)  

```go
func NewChildImage(img *Image, child ChildConfig, os string) *Image
```

##### NewFromJSON

NewFromJSON creates an Image configuration from json.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L289)  
**Added in:** v1.10.0

```go
func NewFromJSON(src []byte) (*Image, error)
```

##### NewImage

NewImage creates a new image with the given ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L212)  

```go
func NewImage(id ID) *Image
```

#### Methods

##### Image.BaseImgArch

BaseImgArch returns the image's architecture. If not populated, defaults to the host runtime arch.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L147)  

```go
func (img *Image) BaseImgArch() string
```

##### Image.BaseImgVariant

BaseImgVariant returns the image's variant, whether populated or not.
This avoids creating an inconsistency where the stored image variant
is "greater than" (i.e. v8 vs v6) the actual image variant.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L158)  

```go
func (img *Image) BaseImgVariant() string
```

##### Image.ID

ID returns the image's content-addressable ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L132)  

```go
func (img *Image) ID() ID
```

##### Image.ImageID

ImageID stringifies ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L137)  
**Added in:** v1.11.0

```go
func (img *Image) ImageID() string
```

##### Image.MarshalJSON

MarshalJSON serializes the image to JSON. It sorts the top-level keys so
that JSON that's been manipulated by a push/pull cycle with a legacy
registry won't end up with a different key order.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L185)  
**Added in:** v1.10.0

```go
func (img *Image) MarshalJSON() ([]byte, error)
```

##### Image.OperatingSystem

OperatingSystem returns the image's operating system. If not populated, defaults to the host runtime OS.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L163)  

```go
func (img *Image) OperatingSystem() string
```

##### Image.Platform

Platform generates an OCI platform from the image

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L172)  

```go
func (img *Image) Platform() ocispec.Platform
```

##### Image.RawJSON

RawJSON returns the immutable JSON associated with the image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L127)  
**Added in:** v1.10.0

```go
func (img *Image) RawJSON() []byte
```

##### Image.RunConfig

RunConfig returns the image's container config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L142)  
**Added in:** v1.11.0

```go
func (img *Image) RunConfig() *container.Config
```

---

### LayerGetReleaser

LayerGetReleaser is a minimal interface for getting and releasing images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/store.go#L37)  
**Added in:** v1.10.0

```go
type LayerGetReleaser interface {
	Get(layer.ChainID) (layer.Layer, error)
	Release(layer.Layer) ([]layer.Metadata, error)
}
```

---

### RootFS

RootFS describes images root filesystem
This is currently a placeholder that only supports layers. In the future
this can be made into an interface that supports different implementations.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/rootfs.go#L23)  
**Added in:** v1.10.0

```go
type RootFS struct {
	Type    string         `json:"type"`
	DiffIDs []layer.DiffID `json:"diff_ids,omitempty"`
}
```

#### Functions

##### NewRootFS

NewRootFS returns empty RootFS struct

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/rootfs.go#L29)  
**Added in:** v1.10.0

```go
func NewRootFS() *RootFS
```

#### Methods

##### RootFS.Append

Append appends a new diffID to rootfs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/rootfs.go#L34)  
**Added in:** v1.10.0

```go
func (r *RootFS) Append(id layer.DiffID)
```

##### RootFS.ChainID

ChainID returns the ChainID for the top layer in RootFS.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/rootfs.go#L48)  
**Added in:** v1.10.0

```go
func (r *RootFS) ChainID() layer.ChainID
```

##### RootFS.Clone

Clone returns a copy of the RootFS

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/rootfs.go#L39)  

```go
func (r *RootFS) Clone() *RootFS
```

---

### Store

Store is an interface for creating and accessing images

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/store.go#L19)  
**Added in:** v1.10.0

```go
type Store interface {
	Create(config []byte) (ID, error)
	Get(id ID) (*Image, error)
	Delete(id ID) ([]layer.Metadata, error)
	Search(partialID string) (ID, error)
	SetParent(id ID, parent ID) error
	GetParent(id ID) (ID, error)
	SetLastUpdated(id ID) error
	GetLastUpdated(id ID) (time.Time, error)
	SetBuiltLocally(id ID) error
	IsBuiltLocally(id ID) (bool, error)
	Children(id ID) []ID
	Map() map[ID]*Image
	Heads() map[ID]*Image
	Len() int
}
```

#### Functions

##### NewImageStore

NewImageStore returns new store object for given set of layer stores

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/store.go#L56)  
**Added in:** v1.10.0

```go
func NewImageStore(fs StoreBackend, lss LayerGetReleaser) (Store, error)
```

---

### StoreBackend

StoreBackend provides interface for image.Store persistence

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/fs.go#L20)  
**Added in:** v1.10.0

```go
type StoreBackend interface {
	Walk(f DigestWalkFunc) error
	Get(id digest.Digest) ([]byte, error)
	Set(data []byte) (digest.Digest, error)
	Delete(id digest.Digest) error
	SetMetadata(id digest.Digest, key string, data []byte) error
	GetMetadata(id digest.Digest, key string) ([]byte, error)
	DeleteMetadata(id digest.Digest, key string) error
}
```

#### Functions

##### NewFSStoreBackend

NewFSStoreBackend returns new filesystem based backend for image.Store

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/fs.go#L42)  
**Added in:** v1.10.0

```go
func NewFSStoreBackend(root string) (StoreBackend, error)
```

---

### V1Image

V1Image stores the V1 image configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/image.go#L32)  
**Added in:** v1.10.0

```go
type V1Image struct {
	// ID is a unique 64 character identifier of the image
	ID string `json:"id,omitempty"`

	// Parent is the ID of the parent image.
	//
	// Depending on how the image was created, this field may be empty and
	// is only set for images that were built/created locally. This field
	// is empty if the image was pulled from an image registry.
	Parent string `json:"parent,omitempty"`

	// Comment is an optional message that can be set when committing or
	// importing the image.
	Comment string `json:"comment,omitempty"`

	// Created is the timestamp at which the image was created
	Created *time.Time `json:"created"`

	// Container is the ID of the container that was used to create the image.
	//
	// Depending on how the image was created, this field may be empty.
	Container string `json:"container,omitempty"`

	// ContainerConfig is the configuration of the container that was committed
	// into the image.
	ContainerConfig container.Config `json:"container_config,omitempty"`

	// DockerVersion is the version of Docker that was used to build the image.
	//
	// Depending on how the image was created, this field may be empty.
	DockerVersion string `json:"docker_version,omitempty"`

	// Author is the name of the author that was specified when committing the
	// image, or as specified through MAINTAINER (deprecated) in the Dockerfile.
	Author string `json:"author,omitempty"`

	// Config is the configuration of the container received from the client.
	Config *container.Config `json:"config,omitempty"`

	// Architecture is the hardware CPU architecture that the image runs on.
	Architecture string `json:"architecture,omitempty"`

	// Variant is the CPU architecture variant (presently ARM-only).
	Variant string `json:"variant,omitempty"`

	// OS is the Operating System the image is built to run on.
	OS string `json:"os,omitempty"`

	// Size is the total size of the image including all layers it is composed of.
	Size int64 `json:",omitempty"`
}
```

---


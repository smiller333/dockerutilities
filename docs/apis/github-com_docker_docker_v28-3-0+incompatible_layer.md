# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/layer

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:16 UTC

## Overview

Package layer is package for managing read-only
and read-write mounts on the union file system
driver. Read-only mounts are referenced using a
content hash and are protected from mutation in
the exposed interface. The tar format is used
to create read-only layers and export both
read-only and writable layers. The exported
tar data for a read-only layer should match
the tar used to create the layer.


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/empty.go#L12)

```go
const DigestSHA256EmptyTar = DiffID("sha256:5f70bf18a086007016e948b04aed3b82103a36bea41755b6cddfaf10ace3c6ef")
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L23)

```go
var (
	// ErrLayerDoesNotExist is used when an operation is
	// attempted on a layer which does not exist.
	ErrLayerDoesNotExist = errors.New("layer does not exist")

	// ErrLayerNotRetained is used when a release is
	// attempted on a layer which is not retained.
	ErrLayerNotRetained = errors.New("layer not retained")

	// ErrMountDoesNotExist is used when an operation is
	// attempted on a mount layer which does not exist.
	ErrMountDoesNotExist = errors.New("mount does not exist")

	// ErrMountNameConflict is used when a mount is attempted
	// to be created but there is already a mount with the name
	// used for creation.
	ErrMountNameConflict = errors.New("mount already exists with name")

	// ErrMaxDepthExceeded is used when a layer is attempted
	// to be created which would result in a layer depth
	// greater than the 125 max.
	ErrMaxDepthExceeded = errors.New("max depth exceeded")
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/empty.go#L17)

```go
var EmptyLayer = &emptyLayer{}
```

## Functions

### IsEmpty

IsEmpty returns true if the layer is an EmptyLayer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/empty.go#L58)  

```go
func IsEmpty(diffID DiffID) bool
```

---

### LogReleaseMetadata

LogReleaseMetadata logs a metadata array, uses this to
ensure consistent logging for release metadata

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L221)  

```go
func LogReleaseMetadata(metadatas []Metadata)
```

---

### ReleaseAndLog

ReleaseAndLog releases the provided layer from the given layer
store, logging any error and release metadata

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L211)  

```go
func ReleaseAndLog(ls Store, l Layer)
```

---

## Types

### ChainID

ChainID is the content-addressable ID of a layer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L48)  

```go
type ChainID digest.Digest
```

#### Functions

##### CreateChainID

CreateChainID returns ID for a layerDigest slice

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L193)  

```go
func CreateChainID(dgsts []DiffID) ChainID
```

#### Methods

##### ChainID.String

String returns a string rendition of a layer ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L51)  

```go
func (id ChainID) String() string
```

---

### CreateRWLayerOpts

CreateRWLayerOpts contains optional arguments to be passed to CreateRWLayer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L164)  

```go
type CreateRWLayerOpts struct {
	MountLabel string
	InitFunc   MountInit
	StorageOpt map[string]string
}
```

---

### DescribableStore

DescribableStore represents a layer store capable of storing
descriptors for layers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L188)  
**Added in:** v1.12.0

```go
type DescribableStore interface {
	RegisterWithDescriptor(io.Reader, ChainID, distribution.Descriptor) (Layer, error)
}
```

---

### DiffID

DiffID is the hash of an individual layer tar.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L56)  

```go
type DiffID digest.Digest
```

#### Methods

##### DiffID.String

String returns a string rendition of a layer DiffID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L59)  

```go
func (diffID DiffID) String() string
```

---

### Layer

Layer represents a read-only layer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L72)  

```go
type Layer interface {
	TarStreamer

	// TarStreamFrom returns a tar archive stream for all the layer chain with
	// arbitrary depth.
	TarStreamFrom(ChainID) (io.ReadCloser, error)

	// ChainID returns the content hash of the entire layer chain. The hash
	// chain is made up of DiffID of top layer and all of its parents.
	ChainID() ChainID

	// DiffID returns the content hash of the layer
	// tar stream used to create this layer.
	DiffID() DiffID

	// Parent returns the next layer in the layer chain.
	Parent() Layer

	// Size returns the size of the entire layer chain. The size
	// is calculated from the total size of all files in the layers.
	Size() int64

	// DiffSize returns the size difference of the top layer
	// from parent layer.
	DiffSize() int64

	// Metadata returns the low level storage metadata associated
	// with layer.
	Metadata() (map[string]string, error)
}
```

---

### Metadata

Metadata holds information about a
read-only layer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L142)  

```go
type Metadata struct {
	// ChainID is the content hash of the layer
	ChainID ChainID

	// DiffID is the hash of the tar data used to
	// create the layer
	DiffID DiffID

	// Size is the size of the layer and all parents
	Size int64

	// DiffSize is the size of the top layer
	DiffSize int64
}
```

---

### MountInit

MountInit is a function to initialize a
writable mount. Changes made here will
not be included in the Tar stream of the
RWLayer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L161)  

```go
type MountInit func(root string) error
```

---

### RWLayer

RWLayer represents a layer which is
read and writable

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L105)  

```go
type RWLayer interface {
	TarStreamer

	// Name of mounted layer
	Name() string

	// Parent returns the layer which the writable
	// layer was created from.
	Parent() Layer

	// Mount mounts the RWLayer and returns the filesystem path
	// to the writable layer.
	Mount(mountLabel string) (string, error)

	// Unmount unmounts the RWLayer. This should be called
	// for every mount. If there are multiple mount calls
	// this operation will only decrement the internal mount counter.
	Unmount() error

	// Size represents the size of the writable layer
	// as calculated by the total size of the files
	// changed in the mutable layer.
	Size() (int64, error)

	// Changes returns the set of changes for the mutable layer
	// from the base layer.
	Changes() ([]archive.Change, error)

	// Metadata returns the low level metadata for the mutable layer
	Metadata() (map[string]string, error)

	// ApplyDiff applies the diff to the RW layer
	ApplyDiff(diff io.Reader) (int64, error)
}
```

---

### Store

Store represents a backend for managing both
read-only and read-write layers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L172)  

```go
type Store interface {
	Register(io.Reader, ChainID) (Layer, error)
	Get(ChainID) (Layer, error)
	Map() map[ChainID]Layer
	Release(Layer) ([]Metadata, error)
	CreateRWLayer(id string, parent ChainID, opts *CreateRWLayerOpts) (RWLayer, error)
	GetRWLayer(id string) (RWLayer, error)
	GetMountID(id string) (string, error)
	ReleaseRWLayer(RWLayer) ([]Metadata, error)
	Cleanup() error
	DriverStatus() [][2]string
	DriverName() string
}
```

#### Functions

##### NewStoreFromOptions

NewStoreFromOptions creates a new Store instance

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer_store.go#L53)  

```go
func NewStoreFromOptions(options StoreOptions) (Store, error)
```

---

### StoreOptions

StoreOptions are the options used to create a new Store instance

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer_store.go#L45)  

```go
type StoreOptions struct {
	Root               string
	GraphDriver        string
	GraphDriverOptions []string
	IDMapping          user.IdentityMapping
}
```

---

### TarStreamer

TarStreamer represents an object which may
have its contents exported as a tar stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/layer/layer.go#L65)  

```go
type TarStreamer interface {
	// TarStream returns a tar archive stream
	// for the contents of a layer.
	TarStream() (io.ReadCloser, error)
}
```

---


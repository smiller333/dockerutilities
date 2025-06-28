# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/graphdriver

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:28:10 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fsdiff.go#L17)

```go
var ApplyUncompressedLayer = chrootarchive.ApplyUncompressedLayer
```

## Functions

### IsDriverNotSupported

IsDriverNotSupported returns true if the error initializing
the graph driver is a non-supported error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/errors.go#L29)  

```go
func IsDriverNotSupported(err error) bool
```

---

### ParseStorageOptKeyValue

ParseStorageOptKeyValue parses and validates the specified string as a key/value
pair (key=value).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/utils.go#L10)  

```go
func ParseStorageOptKeyValue(opt string) (key string, value string, err error)
```

---

### Register

Register registers an InitFunc for the driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L119)  

```go
func Register(name string, initFunc InitFunc) error
```

---

## Types

### CreateOpts

CreateOpts contains optional arguments for Create() and CreateReadWrite()
methods.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L23)  
**Added in:** v1.13.0

```go
type CreateOpts struct {
	MountLabel string
	StorageOpt map[string]string
}
```

---

### DiffDriver

DiffDriver is the interface to use to implement graph diffs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L73)  
**Added in:** v1.13.0

```go
type DiffDriver interface {
	// Diff produces an archive of the changes between the specified
	// layer and its parent layer which may be "".
	Diff(id, parent string) (io.ReadCloser, error)
	// Changes produces a list of changes between the specified layer
	// and its parent layer. If parent is "", then all changes will be ADD changes.
	Changes(id, parent string) ([]archive.Change, error)
	// ApplyDiff extracts the changeset from the given diff into the
	// layer with the specified id and parent, returning the size of the
	// new layer in bytes.
	// The archive.Reader must be an uncompressed stream.
	ApplyDiff(id, parent string, diff io.Reader) (size int64, err error)
	// DiffSize calculates the changes between the specified id
	// and its parent and returns the size in bytes of the changes
	// relative to its base filesystem directory.
	DiffSize(id, parent string) (size int64, err error)
}
```

---

### DiffGetterDriver

DiffGetterDriver is the interface for layered file system drivers that
provide a specialized function for getting file contents for tar-split.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L99)  
**Added in:** v1.11.0

```go
type DiffGetterDriver interface {
	Driver
	// DiffGetter returns an interface to efficiently retrieve the contents
	// of files in a layer.
	DiffGetter(id string) (FileGetCloser, error)
}
```

---

### Driver

Driver is the interface for layered/snapshot file system drivers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L92)  

```go
type Driver interface {
	ProtoDriver
	DiffDriver
}
```

#### Functions

##### GetDriver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L131)  

```go
func GetDriver(name string, config Options) (Driver, error)
```

##### New

New creates the driver and initializes it at the specified root.

It is recommended to pass a name for the driver to use, but If no name
is provided, it attempts to detect the prior storage driver based on
existing state, or otherwise selects a storage driver based on a priority
list and the underlying filesystem.

It returns an error if the requested storage driver is not supported,
if scanning prior drivers is ambiguous (i.e., if state is found for
multiple drivers), or if no compatible driver is available for the
platform and underlying filesystem.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L169)  

```go
func New(driverName string, config Options) (Driver, error)
```

##### NewNaiveDiffDriver

NewNaiveDiffDriver returns a fully functional driver that wraps the
given ProtoDriver and adds the capability of the following methods which
it may or may not support on its own:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fsdiff.go#L41)  
**Added in:** v1.10.0

```go
func NewNaiveDiffDriver(driver ProtoDriver, idMap user.IdentityMapping) Driver
```

---

### ErrUnSupported

ErrUnSupported signals that the graph-driver is not supported on the current configuration

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/errors.go#L13)  

```go
type ErrUnSupported interface {
	NotSupported()
}
```

---

### FileGetCloser

FileGetCloser extends the storage.FileGetter interface with a Close method
for cleaning up.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L108)  
**Added in:** v1.11.0

```go
type FileGetCloser interface {
	storage.FileGetter
	// Close cleans up any resources associated with the FileGetCloser.
	Close() error
}
```

---

### InitFunc

InitFunc initializes the storage driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L29)  

```go
type InitFunc func(root string, options []string, idMap user.IdentityMapping) (Driver, error)
```

---

### NaiveDiffDriver

NaiveDiffDriver takes a ProtoDriver and adds the
capability of the Diffing methods on the local file system,
which it may or may not support on its own. See the comment
on the exported NewNaiveDiffDriver function below.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fsdiff.go#L23)  
**Added in:** v1.3.0

```go
type NaiveDiffDriver struct {
	ProtoDriver
	IDMap user.IdentityMapping
	// If true, allow ApplyDiff to succeed in spite of failures to set
	// extended attributes on the unpacked files due to the destination
	// filesystem not supporting them or a lack of permissions. The
	// resulting unpacked layer may be subtly broken.
	BestEffortXattrs bool
}
```

#### Methods

##### NaiveDiffDriver.ApplyDiff

ApplyDiff extracts the changeset from the given diff into the
layer with the specified id and parent, returning the size of the
new layer in bytes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fsdiff.go#L134)  
**Added in:** v1.10.0

```go
func (gdw *NaiveDiffDriver) ApplyDiff(id, parent string, diff io.Reader) (size int64, _ error)
```

##### NaiveDiffDriver.Changes

Changes produces a list of changes between the specified layer
and its parent layer. If parent is "", then all changes will be ADD changes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fsdiff.go#L109)  
**Added in:** v1.10.0

```go
func (gdw *NaiveDiffDriver) Changes(id, parent string) ([]archive.Change, error)
```

##### NaiveDiffDriver.Diff

Diff produces an archive of the changes between the specified
layer and its parent layer which may be "".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fsdiff.go#L50)  
**Added in:** v1.10.0

```go
func (gdw *NaiveDiffDriver) Diff(id, parent string) (arch io.ReadCloser, retErr error)
```

##### NaiveDiffDriver.DiffSize

DiffSize calculates the changes between the specified layer
and its parent and returns the size in bytes of the changes
relative to its base filesystem directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fsdiff.go#L159)  
**Added in:** v1.10.0

```go
func (gdw *NaiveDiffDriver) DiffSize(id, parent string) (size int64, _ error)
```

---

### NotSupportedError

NotSupportedError signals that the graph-driver is not supported on the current configuration

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/errors.go#L18)  

```go
type NotSupportedError string
```

#### Methods

##### NotSupportedError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/errors.go#L20)  

```go
func (e NotSupportedError) Error() string
```

##### NotSupportedError.NotSupported

NotSupported signals that a graph-driver is not supported.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/errors.go#L25)  

```go
func (e NotSupportedError) NotSupported()
```

---

### Options

Options is used to initialize a graphdriver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L151)  
**Added in:** v1.13.0

```go
type Options struct {
	Root                string
	DriverOptions       []string
	IDMap               user.IdentityMapping
	ExperimentalEnabled bool
}
```

---

### ProtoDriver

ProtoDriver defines the basic capabilities of a driver.
This interface exists solely to be a minimum set of methods
for client code which choose not to implement the entire Driver
interface and use the NaiveDiffDriver wrapper constructor.

Use of ProtoDriver directly by client code is not recommended.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/driver.go#L37)  
**Added in:** v1.3.0

```go
type ProtoDriver interface {
	// String returns a string representation of this driver.
	String() string
	// CreateReadWrite creates a new, empty filesystem layer that is ready
	// to be used as the storage for a container. Additional options can
	// be passed in opts. parent may be "" and opts may be nil.
	CreateReadWrite(id, parent string, opts *CreateOpts) error
	// Create creates a new, empty, filesystem layer with the
	// specified id and parent and options passed in opts. Parent
	// may be "" and opts may be nil.
	Create(id, parent string, opts *CreateOpts) error
	// Remove attempts to remove the filesystem layer with this id.
	Remove(id string) error
	// Get returns the mountpoint for the layered filesystem referred
	// to by this id. You can optionally specify a mountLabel or "".
	// Returns the absolute path to the mounted layered filesystem.
	Get(id, mountLabel string) (fs string, err error)
	// Put releases the system resources for the specified id,
	// e.g, unmounting layered filesystem.
	Put(id string) error
	// Exists returns whether a filesystem layer with the specified
	// ID exists on this driver.
	Exists(id string) bool
	// Status returns a set of key-value pairs which give low
	// level diagnostic status about this driver.
	Status() [][2]string
	// GetMetadata returns a set of key-value pairs which give driver-specific
	// low-level information about the image/container that the driver is managing.
	GetMetadata(id string) (map[string]string, error)
	// Cleanup performs necessary tasks to release resources
	// held by the driver, e.g., unmounting all layered filesystems
	// known to this driver.
	Cleanup() error
}
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/volume/service

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:52 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L63)

```go
const AnonymousLabel = "com.docker.volume.anonymous"
```

## Variables

This section is empty.

## Functions

### IsInUse

IsInUse returns a boolean indicating whether the error indicates that a
volume is in use

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/errors.go#L74)  

```go
func IsInUse(err error) bool
```

---

### IsNameConflict

IsNameConflict returns a boolean indicating whether the error indicates that a
volume name is already taken

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/errors.go#L85)  

```go
func IsNameConflict(err error) bool
```

---

### IsNotExist

IsNotExist returns a boolean indicating whether the error indicates that the volume does not exist

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/errors.go#L79)  

```go
func IsNotExist(err error) bool
```

---

## Types

### By

By is an interface which is used to implement filtering on volumes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/by.go#L9)  

```go
type By interface {
	// contains filtered or unexported methods
}
```

#### Functions

##### And

And creates a `By` combining all the passed in bys using AND logic.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/by.go#L28)  

```go
func And(bys ...By) By
```

##### ByDriver

ByDriver is `By` that filters based on the driver names that are passed in

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/by.go#L14)  

```go
func ByDriver(drivers ...string) By
```

##### FromList

FromList returns a By which sets the initial list of volumes to use

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/by.go#L60)  

```go
func FromList(ls *[]volume.Volume, by By) By
```

##### Or

Or creates a `By` combining all the passed in bys using OR logic.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/by.go#L41)  

```go
func Or(bys ...By) By
```

---

### ByReferenced

ByReferenced is a `By` that filters based on if the volume has references

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/by.go#L23)  

```go
type ByReferenced bool
```

---

### CustomFilter

CustomFilter is a `By` that is used by callers to provide custom filtering
logic.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/by.go#L55)  

```go
type CustomFilter filterFunc
```

---

### OpErr

OpErr is the error type returned by functions in the store package. It describes
the operation, volume name, and error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/errors.go#L34)  

```go
type OpErr struct {
	// Err is the error that occurred during the operation.
	Err error
	// Op is the operation which caused the error, such as "create", or "list".
	Op string
	// Name is the name of the resource being requested for this op, typically the volume name or the driver name.
	Name string
	// Refs is the list of references associated with the resource.
	Refs []string
}
```

#### Methods

##### OpErr.Cause

Cause returns the error the caused this error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/errors.go#L63)  

```go
func (e *OpErr) Cause() error
```

##### OpErr.Error

Error satisfies the built-in error interface type.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/errors.go#L46)  

```go
func (e *OpErr) Error() string
```

##### OpErr.Unwrap

Unwrap returns the error the caused this error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/errors.go#L68)  

```go
func (e *OpErr) Unwrap() error
```

---

### StoreOpt

StoreOpt sets options for a VolumeStore

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L81)  

```go
type StoreOpt func(store *VolumeStore) error
```

#### Functions

##### WithEventLogger

WithEventLogger configures the VolumeStore with the given VolumeEventLogger

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L131)  

```go
func WithEventLogger(logger VolumeEventLogger) StoreOpt
```

---

### VolumeEventLogger

VolumeEventLogger interface provides methods to log volume-related events

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L28)  

```go
type VolumeEventLogger interface {
	// LogVolumeEvent generates an event related to a volume.
	LogVolumeEvent(volumeID string, action events.Action, attributes map[string]string)
}
```

---

### VolumeStore

VolumeStore is responsible for storing and reference counting volumes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L212)  

```go
type VolumeStore struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewStore

NewStore creates a new volume store at the given path

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L84)  

```go
func NewStore(rootPath string, drivers *drivers.Store, opts ...StoreOpt) (*VolumeStore, error)
```

#### Methods

##### VolumeStore.CountReferences

CountReferences gives a count of all references for a given volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L868)  

```go
func (s *VolumeStore) CountReferences(v volume.Volume) int
```

##### VolumeStore.Create

Create creates a volume with the given name and driver
If the volume needs to be created with a reference to prevent race conditions
with volume cleanup, make sure to use the `CreateWithReference` option.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L480)  

```go
func (s *VolumeStore) Create(ctx context.Context, name, driverName string, createOpts ...opts.CreateOption) (volume.Volume, error)
```

##### VolumeStore.Find

Find lists volumes filtered by the past in filter.
If a driver returns a volume that has name which conflicts with another volume from a different driver,
the first volume is chosen and the conflicting volume is dropped.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L349)  

```go
func (s *VolumeStore) Find(ctx context.Context, by By) (vols []volume.Volume, warnings []string, err error)
```

##### VolumeStore.Get

Get looks if a volume with the given name exists and returns it if so

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L657)  

```go
func (s *VolumeStore) Get(ctx context.Context, name string, getOptions ...opts.GetOption) (volume.Volume, error)
```

##### VolumeStore.Release

Release releases the specified reference to the volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L843)  

```go
func (s *VolumeStore) Release(ctx context.Context, name string, ref string) error
```

##### VolumeStore.Remove

Remove removes the requested volume. A volume is not removed if it has any refs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L793)  

```go
func (s *VolumeStore) Remove(ctx context.Context, v volume.Volume, rmOpts ...opts.RemoveOption) error
```

##### VolumeStore.Shutdown

Shutdown releases all resources used by the volume store
It does not make any changes to volumes, drivers, etc.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/store.go#L889)  

```go
func (s *VolumeStore) Shutdown() error
```

---

### VolumesService

VolumesService manages access to volumes
This is used as the main access point for volumes to higher level services and the API.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L35)  

```go
type VolumesService struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewVolumeService

NewVolumeService creates a new volume service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L43)  

```go
func NewVolumeService(root string, pg plugingetter.PluginGetter, rootIDs idtools.Identity, logger VolumeEventLogger) (*VolumesService, error)
```

#### Methods

##### VolumesService.Create

Create creates a volume
If the caller is creating this volume to be consumed immediately, it is
expected that the caller specifies a reference ID.
This reference ID will protect this volume from removal.

A good example for a reference ID is a container's ID.
When whatever is going to reference this volume is removed the caller should dereference the volume by calling `Release`.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L72)  

```go
func (s *VolumesService) Create(ctx context.Context, name, driverName string, options ...opts.CreateOption) (*volumetypes.Volume, error)
```

##### VolumesService.Get

Get returns details about a volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L93)  

```go
func (s *VolumesService) Get(ctx context.Context, name string, getOpts ...opts.GetOption) (*volumetypes.Volume, error)
```

##### VolumesService.GetDriverList

GetDriverList gets the list of registered volume drivers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L57)  

```go
func (s *VolumesService) GetDriverList() []string
```

##### VolumesService.List

List gets the list of volumes which match the past in filters
If filters is nil or empty all volumes are returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L265)  

```go
func (s *VolumesService) List(ctx context.Context, filter filters.Args) (volumes []*volumetypes.Volume, warnings []string, _ error)
```

##### VolumesService.LiveRestoreVolume

LiveRestoreVolume passes through the LiveRestoreVolume call to the volume if it is implemented
otherwise it is a no-op.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L286)  

```go
func (s *VolumesService) LiveRestoreVolume(ctx context.Context, vol *volumetypes.Volume, ref string) error
```

##### VolumesService.LocalVolumesSize

LocalVolumesSize gets all local volumes and fetches their size on disk
Note that this intentionally skips volumes which have mount options. Typically
volumes with mount options are not really local even if they are using the
local driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L198)  

```go
func (s *VolumesService) LocalVolumesSize(ctx context.Context) ([]*volumetypes.Volume, error)
```

##### VolumesService.Mount

Mount mounts the volume
Callers should specify a unique reference for each Mount/Unmount pair.

Example:
```go
mountID := "randomString"
s.Mount(ctx, vol, mountID)
s.Unmount(ctx, vol, mountID)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L120)  

```go
func (s *VolumesService) Mount(ctx context.Context, vol *volumetypes.Volume, ref string) (string, error)
```

##### VolumesService.Prune

Prune removes (local) volumes which match the past in filter arguments.
Note that this intentionally skips volumes with mount options as there would
be no space reclaimed in this case.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L212)  

```go
func (s *VolumesService) Prune(ctx context.Context, filter filters.Args) (*volumetypes.PruneReport, error)
```

##### VolumesService.Release

Release releases a volume reference

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L149)  

```go
func (s *VolumesService) Release(ctx context.Context, name string, ref string) error
```

##### VolumesService.Remove

Remove removes a volume
An error is returned if the volume is still referenced.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L155)  

```go
func (s *VolumesService) Remove(ctx context.Context, name string, rmOpts ...opts.RemoveOption) error
```

##### VolumesService.Shutdown

Shutdown shuts down the image service and dependencies

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L280)  

```go
func (s *VolumesService) Shutdown() error
```

##### VolumesService.Unmount

Unmount unmounts the volume.
Note that depending on the implementation, the volume may still be mounted due to other resources using it.

The reference specified here should be the same reference specified during `Mount` and should be
unique for each mount/unmount pair.
See `Mount` documentation for an example.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/service.go#L137)  

```go
func (s *VolumesService) Unmount(ctx context.Context, vol *volumetypes.Volume, ref string) error
```

---


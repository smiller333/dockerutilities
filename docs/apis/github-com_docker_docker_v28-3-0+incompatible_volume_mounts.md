# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/volume/mounts

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:49 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/parser.go#L12)

```go
var ErrVolumeTargetIsRoot = errors.New("invalid specification: destination can't be '/'")
```

## Functions

This section is empty.

## Types

### MountPoint

MountPoint is the intersection point between a volume and a container. It
specifies which volume is to be used and where inside a container it should
be mounted.

Note that this type is embedded in `container.Container` object and persisted to disk.
Changes to this struct need to by synced with on disk state.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/mounts.go#L42)  

```go
type MountPoint struct {
	// Source is the source path of the mount.
	// E.g. `mount --bind /foo /bar`, `/foo` is the `Source`.
	Source string
	// Destination is the path relative to the container root (`/`) to the mount point
	// It is where the `Source` is mounted to
	Destination string
	// RW is set to true when the mountpoint should be mounted as read-write
	RW bool
	// Name is the name reference to the underlying data defined by `Source`
	// e.g., the volume name
	Name string
	// Driver is the volume driver used to create the volume (if it is a volume)
	Driver string
	// Type of mount to use, see `Type<foo>` definitions in github.com/docker/docker/api/types/mount
	Type mounttypes.Type `json:",omitempty"`
	// Volume is the volume providing data to this mountpoint.
	// This is nil unless `Type` is set to `TypeVolume`
	Volume volume.Volume `json:"-"`

	// Mode is the comma separated list of options supplied by the user when creating
	// the bind/volume mount.
	// Note Mode is not used on Windows
	Mode string `json:"Relabel,omitempty"` // Originally field was `Relabel`"

	// Propagation describes how the mounts are propagated from the host into the
	// mount point, and vice-versa.
	// See https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt
	// Note Propagation is not used on Windows
	Propagation mounttypes.Propagation `json:",omitempty"` // Mount propagation string

	// Specifies if data should be copied from the container before the first mount
	// Use a pointer here so we can tell if the user set this value explicitly
	// This allows us to error out when the user explicitly enabled copy but we can't copy due to the volume being populated
	CopyData bool `json:"-"`
	// ID is the opaque ID used to pass to the volume driver.
	// This should be set by calls to `Mount` and unset by calls to `Unmount`
	ID string `json:",omitempty"`

	// Spec is a copy of the API request that created this mount.
	Spec mounttypes.Mount

	// Some bind mounts should not be automatically created.
	// (Some are auto-created for backwards-compatibility)
	// This is checked on the API but setting this here prevents race conditions.
	// where a bind dir existed during validation was removed before reaching the setup code.
	SkipMountpointCreation bool

	Layer RWLayer `json:"-"`
	// contains filtered or unexported fields
}
```

#### Methods

##### MountPoint.Cleanup

Cleanup frees resources used by the mountpoint and cleans up all the paths
returned by Setup that hasn't been cleaned up by the caller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/mounts.go#L104)  

```go
func (m *MountPoint) Cleanup(ctx context.Context) error
```

##### MountPoint.LiveRestore

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/mounts.go#L269)  

```go
func (m *MountPoint) LiveRestore(ctx context.Context) error
```

##### MountPoint.Path

Path returns the path of a volume in a mount point.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/mounts.go#L296)  

```go
func (m *MountPoint) Path() string
```

##### MountPoint.Setup

Setup sets up a mount point by either mounting the volume if it is
configured, or creating the source directory if supplied.
The, optional, checkFun parameter allows doing additional checking
before creating the source directory on the host.

The returned path can be a temporary path, caller is responsible to
call the returned cleanup function as soon as the path is not needed.
Cleanup doesn't unmount the underlying volumes (if any), it only
frees up the resources that were needed to guarantee that the path
still points to the same target (to avoid TOCTOU attack).

Cleanup function doesn't need to be called when error is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/mounts.go#L158)  

```go
func (m *MountPoint) Setup(ctx context.Context, mountLabel string, rootIDs idtools.Identity, checkFun func(m *MountPoint) error) (mountPath string, cleanup func(context.Context) error, retErr error)
```

---

### Parser

Parser represents a platform specific parser for mount expressions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/parser.go#L29)  

```go
type Parser interface {
	ParseMountRaw(raw, volumeDriver string) (*MountPoint, error)
	ParseMountSpec(cfg mount.Mount) (*MountPoint, error)
	ParseVolumesFrom(spec string) (string, string, error)
	DefaultPropagationMode() mount.Propagation
	ConvertTmpfsOptions(opt *mount.TmpfsOptions, readOnly bool) (string, error)
	DefaultCopyMode() bool
	ValidateVolumeName(name string) error
	ReadWrite(mode string) bool
	IsBackwardCompatible(m *MountPoint) bool
	HasResource(m *MountPoint, absPath string) bool
	ValidateTmpfsMountDestination(dest string) error
	ValidateMountConfig(mt *mount.Mount) error
}
```

#### Functions

##### NewLCOWParser

NewLCOWParser creates a parser with Linux Containers on Windows semantics.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/lcow_parser.go#L14)  

```go
func NewLCOWParser() Parser
```

##### NewLinuxParser

NewLinuxParser creates a parser with Linux semantics.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/linux_parser.go#L15)  

```go
func NewLinuxParser() Parser
```

##### NewParser

NewParser creates a parser for the current host OS

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/parser.go#L45)  

```go
func NewParser() Parser
```

##### NewWindowsParser

NewWindowsParser creates a parser with Windows semantics.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/windows_parser.go#L15)  

```go
func NewWindowsParser() Parser
```

---

### RWLayer

RWLayer represents a writable layer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/mounts/mounts.go#L22)  

```go
type RWLayer interface {
	// Mount mounts the RWLayer and returns the filesystem path
	// to the writable layer.
	Mount(mountLabel string) (string, error)

	// Unmount unmounts the RWLayer. This should be called
	// for every mount. If there are multiple mount calls
	// this operation will only decrement the internal mount counter.
	Unmount() error

	// Metadata returns the low level metadata for the mutable layer
	Metadata() (map[string]string, error)
}
```

---


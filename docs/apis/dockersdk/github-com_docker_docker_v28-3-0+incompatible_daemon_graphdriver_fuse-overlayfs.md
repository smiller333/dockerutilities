# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/graphdriver/fuse-overlayfs

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:13 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Init

Init returns the naive diff driver for fuse-overlayfs.
If fuse-overlayfs is not supported on the host, the error
graphdriver.ErrNotSupported is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L77)  

```go
func Init(home string, options []string, idMap user.IdentityMapping) (graphdriver.Driver, error)
```

---

## Types

### Driver

Driver contains information about the home directory and the list of active
mounts that are created using this driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L60)  

```go
type Driver struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Driver.ApplyDiff

ApplyDiff applies the new layer into a root

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L453)  

```go
func (d *Driver) ApplyDiff(id string, parent string, diff io.Reader) (size int64, _ error)
```

##### Driver.Changes

Changes produces a list of changes between the specified layer and its
parent layer. If parent is "", then all changes will be ADD changes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L495)  

```go
func (d *Driver) Changes(id, parent string) ([]archive.Change, error)
```

##### Driver.Cleanup

Cleanup any state created by overlay which should be cleaned when daemon
is being shutdown. For now, we just have to unmount the bind mounted
we had created.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L150)  

```go
func (d *Driver) Cleanup() error
```

##### Driver.Create

Create is used to create the upper, lower, and merge directories required for overlay fs for a given id.
The parent filesystem is used to configure these directories for the overlay.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L165)  

```go
func (d *Driver) Create(id, parent string, opts *graphdriver.CreateOpts) (retErr error)
```

##### Driver.CreateReadWrite

CreateReadWrite creates a layer that is writable for use as a container
file system.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L156)  

```go
func (d *Driver) CreateReadWrite(id, parent string, opts *graphdriver.CreateOpts) error
```

##### Driver.Diff

Diff produces an archive of the changes between the specified
layer and its parent layer which may be "".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L489)  

```go
func (d *Driver) Diff(id, parent string) (io.ReadCloser, error)
```

##### Driver.DiffSize

DiffSize calculates the changes between the specified id
and its parent and returns the size in bytes of the changes
relative to its base filesystem directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L483)  

```go
func (d *Driver) DiffSize(id, parent string) (int64, error)
```

##### Driver.Exists

Exists checks to see if the id is already mounted.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L425)  

```go
func (d *Driver) Exists(id string) bool
```

##### Driver.Get

Get creates and mounts the required file system for the given id and returns the mount path.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L305)  

```go
func (d *Driver) Get(id, mountLabel string) (_ string, retErr error)
```

##### Driver.GetMetadata

GetMetadata returns metadata about the overlay driver such as the LowerDir,
UpperDir, WorkDir, and MergeDir used to store data.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L124)  

```go
func (d *Driver) GetMetadata(id string) (map[string]string, error)
```

##### Driver.Put

Put unmounts the mount path created for the give id.
It also removes the 'merged' directory to force the kernel to unmount the
overlay mount in other namespaces.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L389)  

```go
func (d *Driver) Put(id string) error
```

##### Driver.Remove

Remove cleans the directories that are created for this id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L282)  

```go
func (d *Driver) Remove(id string) error
```

##### Driver.Status

Status returns current driver information in a two dimensional string array.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L118)  

```go
func (d *Driver) Status() [][2]string
```

##### Driver.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/fuse-overlayfs/fuseoverlayfs.go#L113)  

```go
func (d *Driver) String() string
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/graphdriver/windows

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:28 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### InitFilter

InitFilter returns a new Windows storage filter driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L92)  

```go
func InitFilter(home string, options []string, _ user.IdentityMapping) (graphdriver.Driver, error)
```

---

## Types

### Driver

Driver represents a windows graph driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L80)  
**Added in:** v1.9.0

```go
type Driver struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Driver.ApplyDiff

ApplyDiff extracts the changeset from the given diff into the
layer with the specified id and parent, returning the size of the
new layer in bytes.
The layer should not be mounted when calling this function

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L573)  
**Added in:** v1.9.0

```go
func (d *Driver) ApplyDiff(id, parent string, diff io.Reader) (int64, error)
```

##### Driver.Changes

Changes produces a list of changes between the specified layer
and its parent layer. If parent is "", then all changes will be ADD changes.
The layer should not be mounted when calling this function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L518)  
**Added in:** v1.9.0

```go
func (d *Driver) Changes(id, _ string) ([]archive.Change, error)
```

##### Driver.Cleanup

Cleanup ensures the information the driver stores is properly removed.
We use this opportunity to cleanup any -removing folders which may be
still left if the daemon was killed while it was removing a layer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L453)  
**Added in:** v1.9.0

```go
func (d *Driver) Cleanup() error
```

##### Driver.Create

Create creates a new read-only layer with the given id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L177)  
**Added in:** v1.9.0

```go
func (d *Driver) Create(id, parent string, opts *graphdriver.CreateOpts) error
```

##### Driver.CreateReadWrite

CreateReadWrite creates a layer that is writable for use as a container
file system.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L169)  
**Added in:** v1.12.0

```go
func (d *Driver) CreateReadWrite(id, parent string, opts *graphdriver.CreateOpts) error
```

##### Driver.Diff

Diff produces an archive of the changes between the specified
layer and its parent layer which may be "".
The layer should be mounted when calling this function

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L482)  
**Added in:** v1.9.0

```go
func (d *Driver) Diff(id, _ string) (io.ReadCloser, error)
```

##### Driver.DiffGetter

DiffGetter returns a FileGetCloser that can read files from the directory that
contains files for the layer differences. Used for direct access for tar-split.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L922)  
**Added in:** v1.11.0

```go
func (d *Driver) DiffGetter(id string) (graphdriver.FileGetCloser, error)
```

##### Driver.DiffSize

DiffSize calculates the changes between the specified layer
and its parent and returns the size in bytes of the changes
relative to its base filesystem directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L607)  
**Added in:** v1.9.0

```go
func (d *Driver) DiffSize(id, parent string) (int64, error)
```

##### Driver.Exists

Exists returns true if the given id is registered with this driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L155)  
**Added in:** v1.9.0

```go
func (d *Driver) Exists(id string) bool
```

##### Driver.Get

Get returns the rootfs path for the id. This will mount the dir at its given path.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L366)  
**Added in:** v1.9.0

```go
func (d *Driver) Get(id, mountLabel string) (string, error)
```

##### Driver.GetLayerPath

GetLayerPath gets the layer path on host

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L361)  

```go
func (d *Driver) GetLayerPath(id string) (string, error)
```

##### Driver.GetMetadata

GetMetadata returns custom driver information.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L628)  
**Added in:** v1.9.0

```go
func (d *Driver) GetMetadata(id string) (map[string]string, error)
```

##### Driver.Put

Put adds a new layer to the driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L424)  
**Added in:** v1.9.0

```go
func (d *Driver) Put(id string) error
```

##### Driver.Remove

Remove unmounts and removes the dir information.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L265)  
**Added in:** v1.9.0

```go
func (d *Driver) Remove(id string) error
```

##### Driver.Status

Status returns the status of the driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L148)  
**Added in:** v1.9.0

```go
func (d *Driver) Status() [][2]string
```

##### Driver.String

String returns the string representation of a driver. This should match
the name the graph driver has been registered with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/windows/windows.go#L143)  
**Added in:** v1.9.0

```go
func (d *Driver) String() string
```

---


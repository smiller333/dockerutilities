# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/graphdriver/zfs

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:30 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Init

Init returns a new ZFS driver.
It takes base mount path and an array of options which are represented as key value pairs.
Each option is in the for key=value. 'zfs.fsname' is expected to be a valid key in the options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L49)  

```go
func Init(base string, opt []string, idMap user.IdentityMapping) (graphdriver.Driver, error)
```

---

## Types

### Driver

Driver holds information about the driver, such as zfs dataset, options and cache.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L176)  

```go
type Driver struct {
	sync.Mutex // protects filesystem cache against concurrent access
	// contains filtered or unexported fields
}
```

#### Methods

##### Driver.Cleanup

Cleanup is called on daemon shutdown, it is a no-op for ZFS.
TODO(@cpuguy83): Walk layer tree and check mounts?

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L192)  

```go
func (d *Driver) Cleanup() error
```

##### Driver.Create

Create prepares the dataset and filesystem for the ZFS driver for the given id under the parent.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L274)  

```go
func (d *Driver) Create(id, parent string, opts *graphdriver.CreateOpts) error
```

##### Driver.CreateReadWrite

CreateReadWrite creates a layer that is writable for use as a container
file system.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L269)  
**Added in:** v1.12.0

```go
func (d *Driver) CreateReadWrite(id, parent string, opts *graphdriver.CreateOpts) error
```

##### Driver.Exists

Exists checks to see if the cache entry exists for the given id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L447)  

```go
func (d *Driver) Exists(id string) bool
```

##### Driver.Get

Get returns the mountpoint for the given id after creating the target directories if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L380)  

```go
func (d *Driver) Get(id, mountLabel string) (_ string, retErr error)
```

##### Driver.GetMetadata

GetMetadata returns image/container metadata related to graph driver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L230)  
**Added in:** v1.8.0

```go
func (d *Driver) GetMetadata(id string) (map[string]string, error)
```

##### Driver.Put

Put removes the existing mountpoint for the given id if it exists.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L424)  

```go
func (d *Driver) Put(id string) error
```

##### Driver.Remove

Remove deletes the dataset, filesystem and the cache for the given id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L354)  

```go
func (d *Driver) Remove(id string) error
```

##### Driver.Status

Status returns information about the ZFS filesystem. It returns a two dimensional array of information
such as pool name, dataset name, disk usage, parent quota and compression used.
Currently it return 'Zpool', 'Zpool Health', 'Parent Dataset', 'Space Used By Parent',
'Space Available', 'Parent Quota' and 'Compression'.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L200)  

```go
func (d *Driver) Status() [][2]string
```

##### Driver.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L186)  

```go
func (d *Driver) String() string
```

---

### Logger

Logger returns a zfs logger implementation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L39)  

```go
type Logger struct{}
```

#### Methods

##### Logger.Log

Log wraps log message from ZFS driver with a prefix 'zfs'.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/zfs/zfs.go#L42)  

```go
func (*Logger) Log(cmd []string)
```

---


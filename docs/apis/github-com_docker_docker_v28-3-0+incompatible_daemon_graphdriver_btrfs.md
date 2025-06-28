# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/graphdriver/btrfs

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:28:12 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Init

Init returns a new BTRFS driver.
An error is returned if BTRFS is not supported.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L61)  

```go
func Init(home string, options []string, idMap user.IdentityMapping) (graphdriver.Driver, error)
```

---

## Types

### Driver

Driver contains information about the filesystem mounted.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L139)  

```go
type Driver struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Driver.Cleanup

Cleanup unmounts the home directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L166)  

```go
func (d *Driver) Cleanup() error
```

##### Driver.Create

Create the filesystem with given id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L482)  

```go
func (d *Driver) Create(id, parent string, opts *graphdriver.CreateOpts) error
```

##### Driver.CreateReadWrite

CreateReadWrite creates a layer that is writable for use as a container
file system.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L477)  
**Added in:** v1.12.0

```go
func (d *Driver) CreateReadWrite(id, parent string, opts *graphdriver.CreateOpts) error
```

##### Driver.Exists

Exists checks if the id exists in the filesystem.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L652)  

```go
func (d *Driver) Exists(id string) bool
```

##### Driver.Get

Get the requested filesystem id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L619)  

```go
func (d *Driver) Get(id, mountLabel string) (string, error)
```

##### Driver.GetMetadata

GetMetadata returns empty metadata for this driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L161)  
**Added in:** v1.8.0

```go
func (d *Driver) GetMetadata(id string) (map[string]string, error)
```

##### Driver.Put

Put is not implemented for BTRFS as there is no cleanup required for the id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L645)  

```go
func (d *Driver) Put(id string) error
```

##### Driver.Remove

Remove the filesystem with given id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L581)  

```go
func (d *Driver) Remove(id string) error
```

##### Driver.Status

Status returns the status of the driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L154)  

```go
func (d *Driver) Status() [][2]string
```

##### Driver.String

String prints the name of the driver (btrfs).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/btrfs/btrfs.go#L149)  

```go
func (d *Driver) String() string
```

---


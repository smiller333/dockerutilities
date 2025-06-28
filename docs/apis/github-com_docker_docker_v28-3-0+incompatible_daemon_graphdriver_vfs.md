# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/graphdriver/vfs

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:25 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L24)

```go
var CopyDir = dirCopy
```

## Functions

### Init

Init returns a new VFS driver.
This sets the home directory for the driver and returns NaiveDiffDriver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L32)  

```go
func Init(home string, options []string, idMap user.IdentityMapping) (graphdriver.Driver, error)
```

---

## Types

### Driver

Driver holds information about the driver, home directory of the driver.
Driver implements graphdriver.ProtoDriver. It uses only basic vfs operations.
In order to support layering, files are copied from the parent layer into the new layer. There is no copy-on-write support.
Driver must be wrapped in NaiveDiffDriver to be used as a graphdriver.Driver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L64)  

```go
type Driver struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### Driver.Cleanup

Cleanup is used to implement graphdriver.ProtoDriver. There is no cleanup required for this driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L93)  

```go
func (d *Driver) Cleanup() error
```

##### Driver.Create

Create prepares the filesystem for the VFS driver and copies the directory for the given id under the parent.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L151)  

```go
func (d *Driver) Create(id, parent string, opts *graphdriver.CreateOpts) error
```

##### Driver.CreateReadWrite

CreateReadWrite creates a layer that is writable for use as a container
file system.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L126)  
**Added in:** v1.12.0

```go
func (d *Driver) CreateReadWrite(id, parent string, opts *graphdriver.CreateOpts) error
```

##### Driver.Exists

Exists checks to see if the directory exists for the given id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L218)  

```go
func (d *Driver) Exists(id string) bool
```

##### Driver.Get

Get returns the directory for the given id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L200)  

```go
func (d *Driver) Get(id, mountLabel string) (string, error)
```

##### Driver.GetMetadata

GetMetadata is used for implementing the graphdriver.ProtoDriver interface. VFS does not currently have any meta data.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L88)  
**Added in:** v1.8.0

```go
func (d *Driver) GetMetadata(id string) (map[string]string, error)
```

##### Driver.Put

Put is a noop for vfs that return nil for the error, since this driver has no runtime resources to clean up.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L211)  

```go
func (d *Driver) Put(id string) error
```

##### Driver.Remove

Remove deletes the content from the directory for a given id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L195)  

```go
func (d *Driver) Remove(id string) error
```

##### Driver.Status

Status is used for implementing the graphdriver.ProtoDriver interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L76)  

```go
func (d *Driver) Status() [][2]string
```

##### Driver.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/vfs/driver.go#L71)  

```go
func (d *Driver) String() string
```

---


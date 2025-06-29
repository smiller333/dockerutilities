# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/volume/testutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:16:34 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### FakeRefs

FakeRefs checks ref count on a fake plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L218)  

```go
func FakeRefs(p plugingetter.CompatPlugin) int
```

---

### MakeFakePlugin

MakeFakePlugin creates a fake plugin from the passed in driver
Note: currently only "Create" is implemented because that's all that's needed
so far. If you need it to test something else, add it here, but probably you
shouldn't need to use this except for very specific cases with v2 plugin handling.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L146)  

```go
func MakeFakePlugin(d volume.Driver, l net.Listener) (plugingetter.CompatPlugin, error)
```

---

### NewFakeDriver

NewFakeDriver creates a new FakeDriver with the specified name

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L84)  
**Added in:** v1.10.0

```go
func NewFakeDriver(name string) volume.Driver
```

---

### NewFakePluginGetter

NewFakePluginGetter returns a plugin getter for fake plugins. It only
implements plugingetter.PluginGetter.Get, and panics when calling
any other method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L198)  

```go
func NewFakePluginGetter(pls ...plugingetter.CompatPlugin) plugingetter.PluginGetter
```

---

### NewFakeVolume

NewFakeVolume creates a new fake volume for testing

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L48)  

```go
func NewFakeVolume(name string, driverName string) volume.Volume
```

---

## Types

### FakeDriver

FakeDriver is a driver that generates fake volumes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L78)  

```go
type FakeDriver struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### FakeDriver.Create

Create initializes a fake volume.
It returns an error if the options include an "error" key with a message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L96)  

```go
func (d *FakeDriver) Create(name string, opts map[string]string) (volume.Volume, error)
```

##### FakeDriver.Get

Get gets the volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L124)  
**Added in:** v1.10.0

```go
func (d *FakeDriver) Get(name string) (volume.Volume, error)
```

##### FakeDriver.List

List lists the volumes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L115)  
**Added in:** v1.10.0

```go
func (d *FakeDriver) List() ([]volume.Volume, error)
```

##### FakeDriver.Name

Name is the name of the driver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L92)  

```go
func (d *FakeDriver) Name() string
```

##### FakeDriver.Remove

Remove deletes a volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L106)  

```go
func (d *FakeDriver) Remove(v volume.Volume) error
```

##### FakeDriver.Scope

Scope returns the local scope

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L132)  
**Added in:** v1.12.0

```go
func (*FakeDriver) Scope() string
```

---

### FakeVolume

FakeVolume is a fake volume with a random name

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L41)  

```go
type FakeVolume struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### FakeVolume.CreatedAt

CreatedAt provides the time the volume (directory) was created at

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L73)  

```go
func (f FakeVolume) CreatedAt() (time.Time, error)
```

##### FakeVolume.DriverName

DriverName is the name of the driver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L56)  

```go
func (f FakeVolume) DriverName() string
```

##### FakeVolume.Mount

Mount mounts the volume in the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L62)  

```go
func (FakeVolume) Mount(_ string) (string, error)
```

##### FakeVolume.Name

Name is the name of the volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L53)  

```go
func (f FakeVolume) Name() string
```

##### FakeVolume.Path

Path is the filesystem path to the volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L59)  

```go
func (FakeVolume) Path() string
```

##### FakeVolume.Status

Status provides low-level details about the volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L68)  
**Added in:** v1.12.0

```go
func (FakeVolume) Status() map[string]interface{}
```

##### FakeVolume.Unmount

Unmount unmounts the volume from the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L65)  

```go
func (FakeVolume) Unmount(_ string) error
```

---

### NoopVolume

NoopVolume is a volume that doesn't perform any operation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L17)  

```go
type NoopVolume struct{}
```

#### Methods

##### NoopVolume.CreatedAt

CreatedAt provides the time the volume (directory) was created at

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L38)  

```go
func (NoopVolume) CreatedAt() (time.Time, error)
```

##### NoopVolume.DriverName

DriverName is the name of the driver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L23)  

```go
func (NoopVolume) DriverName() string
```

##### NoopVolume.Mount

Mount mounts the volume in the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L29)  

```go
func (NoopVolume) Mount(_ string) (string, error)
```

##### NoopVolume.Name

Name is the name of the volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L20)  

```go
func (NoopVolume) Name() string
```

##### NoopVolume.Path

Path is the filesystem path to the volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L26)  

```go
func (NoopVolume) Path() string
```

##### NoopVolume.Status

Status provides low-level details about the volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L35)  
**Added in:** v1.12.0

```go
func (NoopVolume) Status() map[string]interface{}
```

##### NoopVolume.Unmount

Unmount unmounts the volume from the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/testutils/testutils.go#L32)  

```go
func (NoopVolume) Unmount(_ string) error
```

---


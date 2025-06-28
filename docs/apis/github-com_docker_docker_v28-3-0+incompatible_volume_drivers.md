# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/volume/drivers

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:16:22 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Store

Store is an in-memory store for volume drivers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/drivers/extpoint.go#L47)  

```go
type Store struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewStore

NewStore creates a new volume driver store

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/drivers/extpoint.go#L55)  

```go
func NewStore(pg plugingetter.PluginGetter) *Store
```

#### Methods

##### Store.CreateDriver

CreateDriver returns a volume driver by its name and increments RefCount.
If the driver is empty, it looks for the local driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/drivers/extpoint.go#L155)  

```go
func (s *Store) CreateDriver(name string) (volume.Driver, error)
```

##### Store.GetAllDrivers

GetAllDrivers lists all the registered drivers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/drivers/extpoint.go#L179)  

```go
func (s *Store) GetAllDrivers() ([]volume.Driver, error)
```

##### Store.GetDriver

GetDriver returns a volume driver by its name.
If the driver is empty, it looks for the local driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/drivers/extpoint.go#L149)  

```go
func (s *Store) GetDriver(name string) (volume.Driver, error)
```

##### Store.GetDriverList

GetDriverList returns list of volume drivers registered.
If no driver is registered, empty string list will be returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/drivers/extpoint.go#L167)  

```go
func (s *Store) GetDriverList() []string
```

##### Store.Register

Register associates the given driver to the given name, checking if
the name is already associated

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/drivers/extpoint.go#L127)  

```go
func (s *Store) Register(d volume.Driver, name string) bool
```

##### Store.ReleaseDriver

ReleaseDriver returns a volume driver by its name and decrements RefCount..
If the driver is empty, it looks for the local driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/drivers/extpoint.go#L161)  

```go
func (s *Store) ReleaseDriver(name string) (volume.Driver, error)
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/volume

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:42 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/volume.go#L14)

```go
const (
	LocalScope  = "local"
	GlobalScope = "global"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/volume.go#L10)

```go
const DefaultDriverName = "local"
```

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Capability

Capability defines a set of capabilities that a driver is able to handle.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/volume.go#L37)  
**Added in:** v1.12.0

```go
type Capability struct {
	// Scope is the scope of the driver, `global` or `local`
	// A `global` scope indicates that the driver manages volumes across the cluster
	// A `local` scope indicates that the driver only manages volumes resources local to the host
	// Scope is declared by the driver
	Scope string
}
```

---

### DetailedVolume

DetailedVolume wraps a Volume with user-defined labels, options, and cluster scope (e.g., `local` or `global`)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/volume.go#L74)  
**Added in:** v1.13.0

```go
type DetailedVolume interface {
	Labels() map[string]string
	Options() map[string]string
	Scope() string
	Volume
}
```

---

### Driver

Driver is for creating and removing volumes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/volume.go#L20)  

```go
type Driver interface {
	// Name returns the name of the volume driver.
	Name() string
	// Create makes a new volume with the given name.
	Create(name string, opts map[string]string) (Volume, error)
	// Remove deletes the volume.
	Remove(vol Volume) (err error)
	// List lists all the volumes the driver has
	List() ([]Volume, error)
	// Get retrieves the volume with the requested name
	Get(name string) (Volume, error)
	// Scope returns the scope of the driver (e.g. `global` or `local`).
	// Scope determines how the driver is handled at a cluster level
	Scope() string
}
```

---

### LiveRestorer

LiveRestorer is an optional interface that can be implemented by a volume driver
It is used to restore any resources that are necessary for a volume to be used by a live-restored container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/volume.go#L66)  

```go
type LiveRestorer interface {
	// LiveRestoreVolume allows a volume driver which implements this interface to restore any necessary resources (such as reference counting)
	// This is called only after the daemon is restarted with live-restored containers
	// It is called once per live-restored container.
	LiveRestoreVolume(_ context.Context, ref string) error
}
```

---

### Volume

Volume is a place to store data. It is backed by a specific driver, and can be mounted.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/volume.go#L46)  

```go
type Volume interface {
	// Name returns the name of the volume
	Name() string
	// DriverName returns the name of the driver which owns this volume.
	DriverName() string
	// Path returns the absolute path to the volume.
	Path() string
	// Mount mounts the volume and returns the absolute path to
	// where it can be consumed.
	Mount(id string) (string, error)
	// Unmount unmounts the volume when it is no longer in use.
	Unmount(id string) error
	// CreatedAt returns Volume Creation time
	CreatedAt() (time.Time, error)
	// Status returns low-level status information about a volume
	Status() map[string]interface{}
}
```

---


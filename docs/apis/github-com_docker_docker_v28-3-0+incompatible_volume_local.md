# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/volume/local

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:16:24 UTC

## Overview

Package local provides the default implementation for volumes. It
is used to mount data volume containers and directories local to
the host server.

Package local provides the default implementation for volumes. It
is used to mount data volume containers and directories local to
the host server.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/local/local.go#L35)

```go
var (
	// ErrNotFound is the typed error returned when the requested volume name can't be found
	ErrNotFound = errors.New("volume not found")
)
```

## Functions

This section is empty.

## Types

### Root

Root implements the Driver interface for the volume package and
manages the creation/removal of volumes. It uses only standard vfs
commands to create/remove dirs within its provided scope.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/local/local.go#L110)  

```go
type Root struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New instantiates a new Root instance with the provided scope. Scope
is the base path that the Root instance uses to store its
volumes. The base path is created here if it does not exist.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/local/local.go#L54)  

```go
func New(scope string, rootIdentity idtools.Identity) (*Root, error)
```

#### Methods

##### Root.Create

Create creates a new volume.Volume with the provided name, creating
the underlying directory tree required for this volume in the
process.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/local/local.go#L138)  

```go
func (r *Root) Create(name string, opts map[string]string) (volume.Volume, error)
```

##### Root.Get

Get looks up the volume for the given name and returns it if found

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/local/local.go#L242)  
**Added in:** v1.9.0

```go
func (r *Root) Get(name string) (volume.Volume, error)
```

##### Root.List

List lists all the volumes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/local/local.go#L120)  
**Added in:** v1.9.0

```go
func (r *Root) List() ([]volume.Volume, error)
```

##### Root.Name

Name returns the name of Root, defined in the volume package in the DefaultDriverName constant.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/local/local.go#L131)  

```go
func (r *Root) Name() string
```

##### Root.Remove

Remove removes the specified volume and all underlying data. If the
given volume does not belong to this driver and an error is
returned. The volume is reference counted, if all references are
not released then the volume is not removed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/local/local.go#L191)  

```go
func (r *Root) Remove(v volume.Volume) error
```

##### Root.Scope

Scope returns the local volume scope

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/local/local.go#L253)  
**Added in:** v1.12.0

```go
func (r *Root) Scope() string
```

---


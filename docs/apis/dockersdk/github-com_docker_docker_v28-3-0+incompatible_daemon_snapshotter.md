# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/snapshotter

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:42 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewMounter

NewMounter creates a new mounter for the provided snapshotter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/snapshotter/mount.go#L28)  

```go
func NewMounter(home string, snapshotter string, idMap user.IdentityMapping) *refCountMounter
```

---

## Types

### Mounter

Mounter handles mounting/unmounting things coming in from a snapshotter
with optional reference counting if needed by the filesystem

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/snapshotter/mount.go#L18)  

```go
type Mounter interface {
	// Mount mounts the rootfs for a container and returns the mount point
	Mount(mounts []mount.Mount, containerID string) (string, error)
	// Unmount unmounts the container rootfs
	Unmount(target string) error
	// Mounted returns a target mountpoint if it's already mounted
	Mounted(containerID string) (string, error)
}
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/mount

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:04 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L63)

```go
var Propagations = []Propagation{
	PropagationRPrivate,
	PropagationPrivate,
	PropagationRShared,
	PropagationShared,
	PropagationRSlave,
	PropagationSlave,
}
```

## Functions

This section is empty.

## Types

### BindOptions

BindOptions defines options specific to mounts of type "bind".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L87)  

```go
type BindOptions struct {
	Propagation      Propagation `json:",omitempty"`
	NonRecursive     bool        `json:",omitempty"`
	CreateMountpoint bool        `json:",omitempty"`
	// ReadOnlyNonRecursive makes the mount non-recursively read-only, but still leaves the mount recursive
	// (unless NonRecursive is set to true in conjunction).
	ReadOnlyNonRecursive bool `json:",omitempty"`
	// ReadOnlyForceRecursive raises an error if the mount cannot be made recursively read-only.
	ReadOnlyForceRecursive bool `json:",omitempty"`
}
```

---

### ClusterOptions

ClusterOptions specifies options for a Cluster volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L155)  

```go
type ClusterOptions struct {
}
```

---

### Consistency

Consistency represents the consistency requirements of a mount.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L73)  

```go
type Consistency string
```

---

### Driver

Driver represents a volume driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L111)  

```go
type Driver struct {
	Name    string            `json:",omitempty"`
	Options map[string]string `json:",omitempty"`
}
```

---

### ImageOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L106)  

```go
type ImageOptions struct {
	Subpath string `json:",omitempty"`
}
```

---

### Mount

Mount represents a mount (volume).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L27)  

```go
type Mount struct {
	Type Type `json:",omitempty"`
	// Source specifies the name of the mount. Depending on mount type, this
	// may be a volume name or a host path, or even ignored.
	// Source is not supported for tmpfs (must be an empty value)
	Source      string      `json:",omitempty"`
	Target      string      `json:",omitempty"`
	ReadOnly    bool        `json:",omitempty"` // attempts recursive read-only if possible
	Consistency Consistency `json:",omitempty"`

	BindOptions    *BindOptions    `json:",omitempty"`
	VolumeOptions  *VolumeOptions  `json:",omitempty"`
	ImageOptions   *ImageOptions   `json:",omitempty"`
	TmpfsOptions   *TmpfsOptions   `json:",omitempty"`
	ClusterOptions *ClusterOptions `json:",omitempty"`
}
```

---

### Propagation

Propagation represents the propagation of a mount.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L45)  

```go
type Propagation string
```

---

### TmpfsOptions

TmpfsOptions defines options specific to mounts of type "tmpfs".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L117)  

```go
type TmpfsOptions struct {
	// Size sets the size of the tmpfs, in bytes.
	//
	// This will be converted to an operating system specific value
	// depending on the host. For example, on linux, it will be converted to
	// use a 'k', 'm' or 'g' syntax. BSD, though not widely supported with
	// docker, uses a straight byte value.
	//
	// Percentages are not supported.
	SizeBytes int64 `json:",omitempty"`
	// Mode of the tmpfs upon creation
	Mode os.FileMode `json:",omitempty"`
	// Options to be passed to the tmpfs mount. An array of arrays. Flag
	// options should be provided as 1-length arrays. Other types should be
	// provided as 2-length arrays, where the first item is the key and the
	// second the value.
	Options [][]string `json:",omitempty"`
}
```

---

### Type

Type represents the type of a mount.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L8)  

```go
type Type string
```

---

### VolumeOptions

VolumeOptions represents the options for a mount of type volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/mount/mount.go#L99)  

```go
type VolumeOptions struct {
	NoCopy       bool              `json:",omitempty"`
	Labels       map[string]string `json:",omitempty"`
	Subpath      string            `json:",omitempty"`
	DriverConfig *Driver           `json:",omitempty"`
}
```

---


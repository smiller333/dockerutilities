# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/idtools

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:33 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GetRootUIDGID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L53)  

```go
func GetRootUIDGID(uidMap, gidMap []IDMap) (int, int, error)
```

---

### MkdirAllAndChown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L26)  

```go
func MkdirAllAndChown(path string, mode os.FileMode, owner Identity) error
```

---

### MkdirAllAndChownNew

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L45)  

```go
func MkdirAllAndChownNew(path string, mode os.FileMode, owner Identity) error
```

---

### MkdirAndChown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L36)  

```go
func MkdirAndChown(path string, mode os.FileMode, owner Identity) error
```

---

### ToUserIdentityMapping

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L154)  

```go
func ToUserIdentityMapping(u IdentityMapping) user.IdentityMapping
```

---

## Types

### IDMap

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L15)  

```go
type IDMap struct {
	ContainerID int `json:"container_id"`
	HostID      int `json:"host_id"`
	Size        int `json:"size"`
}
```

---

### Identity

Identity is either a UID and GID pair or a SID (but not both)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L104)  

```go
type Identity struct {
	UID int
	GID int
	SID string
}
```

#### Functions

##### CurrentIdentity

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L221)  

```go
func CurrentIdentity() Identity
```

#### Methods

##### Identity.Chown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L113)  

```go
func (id Identity) Chown(name string) error
```

---

### IdentityMapping

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L121)  

```go
type IdentityMapping struct {
	UIDMaps []IDMap `json:"UIDMaps"`
	GIDMaps []IDMap `json:"GIDMaps"`
}
```

---


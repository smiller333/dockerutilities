# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/idtools

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:28 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GetRootUIDGID ⚠️ **DEPRECATED**

GetRootUIDGID retrieves the remapped root uid/gid pair from the set of maps.
If the maps are empty, then the root uid/gid will default to "real" 0/0

Deprecated: use [(user.IdentityMapping).RootPair] instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L53)  

```go
func GetRootUIDGID(uidMap, gidMap []IDMap) (int, int, error)
```

---

### MkdirAllAndChown ⚠️ **DEPRECATED**

MkdirAllAndChown creates a directory (include any along the path) and then modifies
ownership to the requested uid/gid.  If the directory already exists, this
function will still change ownership and permissions.

Deprecated: use user.MkdirAllAndChown instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L26)  

```go
func MkdirAllAndChown(path string, mode os.FileMode, owner Identity) error
```

---

### MkdirAllAndChownNew ⚠️ **DEPRECATED**

MkdirAllAndChownNew creates a directory (include any along the path) and then modifies
ownership ONLY of newly created directories to the requested uid/gid. If the
directories along the path exist, no change of ownership or permissions will be performed

Deprecated: use user.MkdirAllAndChown with the user.WithOnlyNew option instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L45)  

```go
func MkdirAllAndChownNew(path string, mode os.FileMode, owner Identity) error
```

---

### MkdirAndChown ⚠️ **DEPRECATED**

MkdirAndChown creates a directory and then modifies ownership to the requested uid/gid.
If the directory already exists, this function still changes ownership and permissions.
Note that unlike os.Mkdir(), this function does not return IsExist error
in case path already exists.

Deprecated: use user.MkdirAndChown instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L36)  

```go
func MkdirAndChown(path string, mode os.FileMode, owner Identity) error
```

---

### ToUserIdentityMapping ⚠️ **DEPRECATED**

ToUserIdentityMapping converts an idtools.IdentityMapping to a user.IdentityMapping.

Deprecated: use user.IdentityMapping directly, this is transitioning to user package.

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

##### CurrentIdentity ⚠️ **DEPRECATED**

CurrentIdentity returns the identity of the current process

Deprecated: use os.Getuid and os.Getegid instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/idtools/idtools.go#L221)  

```go
func CurrentIdentity() Identity
```

#### Methods

##### Identity.Chown ⚠️ **DEPRECATED**

Chown changes the numeric uid and gid of the named file to id.UID and id.GID.

Deprecated: this method is deprecated and will be removed in the next release.

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


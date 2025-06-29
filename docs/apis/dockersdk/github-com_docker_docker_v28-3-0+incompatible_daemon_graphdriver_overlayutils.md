# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/graphdriver/overlayutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:21 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ErrDTypeNotSupported

ErrDTypeNotSupported denotes that the backing filesystem doesn't support d_type.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/overlayutils/overlayutils.go#L20)  

```go
func ErrDTypeNotSupported(driver, backingFs string) error
```

---

### GenerateID

GenerateID creates a new random string identifier with the given length

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/overlayutils/randomid.go#L19)  

```go
func GenerateID(l int, logger *log.Entry) string
```

---

### GetOverlayXattr

GetOverlayXattr combines the overlay module's xattr class with the named
xattr -- `user` when mounted inside a user namespace, and `trusted` when
mounted in the 'root' namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/overlayutils/overlayutils.go#L85)  

```go
func GetOverlayXattr(name string) string
```

---

### NeedsUserXAttr

NeedsUserXAttr returns whether overlayfs should be mounted with the "userxattr" mount option.

The "userxattr" option is needed for mounting overlayfs inside a user namespace with kernel >= 5.11.

The "userxattr" option is NOT needed for the initial user namespace (aka "the host").

Also, Ubuntu (since circa 2015) and Debian (since 10) with kernel < 5.11 can mount
the overlayfs in a user namespace without the "userxattr" option.

The corresponding kernel commit: https://github.com/torvalds/linux/commit/2d2f2d7322ff43e0fe92bf8cccdc0b09449bf2e1
> ovl: user xattr
>
> Optionally allow using "user.overlay." namespace instead of "trusted.overlay."
> ...
> Disable redirect_dir and metacopy options, because these would allow privilege escalation through direct manipulation of the
> "user.overlay.redirect" or "user.overlay.metacopy" xattrs.
> ...

The "userxattr" support is not exposed in "/sys/module/overlay/parameters".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/overlayutils/userxattr.go#L53)  

```go
func NeedsUserXAttr(d string) (bool, error)
```

---

### SupportsOverlay

SupportsOverlay checks if the system supports overlay filesystem
by performing an actual overlay mount.

checkMultipleLowers parameter enables check for multiple lowerdirs,
which is required for the overlay2 driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/overlayutils/overlayutils.go#L40)  

```go
func SupportsOverlay(d string, checkMultipleLowers bool) error
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/internal/fstype

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:28:45 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/fstype/fstype.go#L32)

```go
var FsNames = map[FsMagic]string{
	FsMagicUnsupported: "unsupported",

	FsMagicAufs:     "aufs",
	FsMagicBtrfs:    "btrfs",
	FsMagicCramfs:   "cramfs",
	FsMagicEcryptfs: "ecryptfs",
	FsMagicExtfs:    "extfs",
	FsMagicF2fs:     "f2fs",
	FsMagicFUSE:     "fuse",
	FsMagicGPFS:     "gpfs",
	FsMagicJffs2Fs:  "jffs2",
	FsMagicJfs:      "jfs",
	FsMagicNfsFs:    "nfs",
	FsMagicOverlay:  "overlayfs",
	FsMagicRAMFs:    "ramfs",
	FsMagicReiserFs: "reiserfs",
	FsMagicSmbFs:    "smb",
	FsMagicSquashFs: "squashfs",
	FsMagicTmpFs:    "tmpfs",
	FsMagicVxFS:     "vxfs",
	FsMagicXfs:      "xfs",
	FsMagicZfs:      "zfs",
}
```

## Functions

This section is empty.

## Types

### FsMagic

FsMagic unsigned id of the filesystem in use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/fstype/fstype.go#L4)  

```go
type FsMagic uint32
```

#### Functions

##### GetFSMagic

GetFSMagic returns the filesystem id given the path. It returns an error
when failing to detect the filesystem. it returns FsMagicUnsupported
if detection is not supported by the platform, but no error is returned
in this case.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/internal/fstype/fstype.go#L61)  

```go
func GetFSMagic(rootpath string) (FsMagic, error)
```

---


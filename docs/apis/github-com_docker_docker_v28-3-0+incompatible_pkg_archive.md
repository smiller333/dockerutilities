# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/archive

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:13 UTC

## Overview

Package archive provides helper functions for dealing with archive files.


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L79)

```go
const (
	Uncompressed = compression.None  // Deprecated: use [compression.None] instead.
	Bzip2        = compression.Bzip2 // Deprecated: use [compression.Bzip2] instead.
	Gzip         = compression.Gzip  // Deprecated: use [compression.Gzip] instead.
	Xz           = compression.Xz    // Deprecated: use [compression.Xz] instead.
	Zstd         = compression.Zstd  // Deprecated: use [compression.Zstd] instead.
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L87)

```go
const (
	AUFSWhiteoutFormat    = archive.AUFSWhiteoutFormat    // Deprecated: use [archive.AUFSWhiteoutFormat] instead.
	OverlayWhiteoutFormat = archive.OverlayWhiteoutFormat // Deprecated: use [archive.OverlayWhiteoutFormat] instead.
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/changes_deprecated.go#L15)

```go
const (
	ChangeModify = archive.ChangeModify // Deprecated: use [archive.ChangeModify] instead.
	ChangeAdd    = archive.ChangeAdd    // Deprecated: use [archive.ChangeAdd] instead.
	ChangeDelete = archive.ChangeDelete // Deprecated: use [archive.ChangeDelete] instead.
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/whiteouts_deprecated.go#L5)

```go
const (
	WhiteoutPrefix     = archive.WhiteoutPrefix     // Deprecated: use [archive.WhiteoutPrefix] instead.
	WhiteoutMetaPrefix = archive.WhiteoutMetaPrefix // Deprecated: use [archive.WhiteoutMetaPrefix] instead.
	WhiteoutLinkDir    = archive.WhiteoutLinkDir    // Deprecated: use [archive.WhiteoutLinkDir] instead.
	WhiteoutOpaqueDir  = archive.WhiteoutOpaqueDir  // Deprecated: use [archive.WhiteoutOpaqueDir] instead.
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L19)

```go
const ImpliedDirectoryMode = archive.ImpliedDirectoryMode
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L10)

```go
var (
	ErrNotDirectory      = archive.ErrNotDirectory      // Deprecated: use [archive.ErrNotDirectory] instead.
	ErrDirNotExists      = archive.ErrDirNotExists      // Deprecated: use [archive.ErrDirNotExists] instead.
	ErrCannotCopyDir     = archive.ErrCannotCopyDir     // Deprecated: use [archive.ErrCannotCopyDir] instead.
	ErrInvalidCopySource = archive.ErrInvalidCopySource // Deprecated: use [archive.ErrInvalidCopySource] instead.
)
```

## Functions

### ApplyLayer ⚠️ **DEPRECATED**

ApplyLayer parses a diff in the standard layer format from `layer`,
and applies it to the directory `dest`.

Deprecated: use archive.ApplyLayer instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/diff_deprecated.go#L20)  

```go
func ApplyLayer(dest string, layer io.Reader) (int64, error)
```

---

### ApplyUncompressedLayer ⚠️ **DEPRECATED**

ApplyUncompressedLayer parses a diff in the standard layer format from
`layer`, and applies it to the directory `dest`.

Deprecated: use archive.ApplyUncompressedLayer instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/diff_deprecated.go#L28)  
**Added in:** v1.8.0

```go
func ApplyUncompressedLayer(dest string, layer io.Reader, options *TarOptions) (int64, error)
```

---

### Changes ⚠️ **DEPRECATED**

Changes walks the path rw and determines changes for the files in the path,
with respect to the parent layers

Deprecated: use archive.Changes instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/changes_deprecated.go#L30)  

```go
func Changes(layers []string, rw string) ([]archive.Change, error)
```

---

### ChangesDirs ⚠️ **DEPRECATED**

ChangesDirs compares two directories and generates an array of Change objects describing the changes.

Deprecated: use archive.ChangesDirs instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/changes_deprecated.go#L42)  

```go
func ChangesDirs(newDir, oldDir string) ([]archive.Change, error)
```

---

### ChangesSize ⚠️ **DEPRECATED**

ChangesSize calculates the size in bytes of the provided changes, based on newDir.

Deprecated: use archive.ChangesSize instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/changes_deprecated.go#L49)  

```go
func ChangesSize(newDir string, changes []archive.Change) int64
```

---

### CheckSystemDriveAndRemoveDriveLetter ⚠️ **DEPRECATED**

CheckSystemDriveAndRemoveDriveLetter verifies that a path is the system drive.

Deprecated: use archive.CheckSystemDriveAndRemoveDriveLetter instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/path_deprecated.go#L8)  

```go
func CheckSystemDriveAndRemoveDriveLetter(path string) (string, error)
```

---

### CompressStream ⚠️ **DEPRECATED**

CompressStream compresses the dest with specified compression algorithm.

Deprecated: use compression.CompressStream instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L117)  

```go
func CompressStream(dest io.Writer, comp compression.Compression) (io.WriteCloser, error)
```

---

### CopyInfoDestinationPath ⚠️ **DEPRECATED**

CopyInfoDestinationPath stats the given path to create a CopyInfo
struct representing that resource for the destination of an archive copy
operation.

Deprecated: use archive.CopyInfoDestinationPath instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L83)  
**Added in:** v1.8.0

```go
func CopyInfoDestinationPath(path string) (info archive.CopyInfo, err error)
```

---

### CopyInfoSourcePath ⚠️ **DEPRECATED**

CopyInfoSourcePath stats the given path to create a CopyInfo struct.
struct representing that resource for the source of an archive copy
operation.

Deprecated: use archive.CopyInfoSourcePath instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L74)  
**Added in:** v1.8.0

```go
func CopyInfoSourcePath(path string, followLink bool) (archive.CopyInfo, error)
```

---

### CopyResource ⚠️ **DEPRECATED**

CopyResource performs an archive copy from the given source path to the
given destination path.

Deprecated: use archive.CopyResource instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L106)  
**Added in:** v1.8.0

```go
func CopyResource(srcPath, dstPath string, followLink bool) error
```

---

### CopyTo ⚠️ **DEPRECATED**

CopyTo handles extracting the given content whose
entries should be sourced from srcInfo to dstPath.

Deprecated: use archive.CopyTo instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L114)  
**Added in:** v1.8.0

```go
func CopyTo(content io.Reader, srcInfo archive.CopyInfo, dstPath string) error
```

---

### DecompressStream ⚠️ **DEPRECATED**

DecompressStream decompresses the archive and returns a ReaderCloser with the decompressed archive.

Deprecated: use compression.DecompressStream instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L110)  

```go
func DecompressStream(arch io.Reader) (io.ReadCloser, error)
```

---

### DetectCompression ⚠️ **DEPRECATED**

DetectCompression detects the compression algorithm of the source.

Deprecated: use compression.Detect instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L103)  

```go
func DetectCompression(source []byte) archive.Compression
```

---

### ExportChanges

ExportChanges produces an Archive from the provided changes, relative to dir.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/changes_deprecated.go#L54)  

```go
func ExportChanges(dir string, changes []archive.Change, idMap idtools.IdentityMapping) (io.ReadCloser, error)
```

---

### FileInfoHeader ⚠️ **DEPRECATED**

FileInfoHeader creates a populated Header from fi.

Deprecated: use archive.FileInfoHeader instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L143)  

```go
func FileInfoHeader(name string, fi os.FileInfo, link string) (*tar.Header, error)
```

---

### FileInfoHeaderNoLookups ⚠️ **DEPRECATED**

FileInfoHeaderNoLookups creates a partially-populated tar.Header from fi.

Deprecated: use tarheader.FileInfoHeaderNoLookups instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L136)  

```go
func FileInfoHeaderNoLookups(fi os.FileInfo, link string) (*tar.Header, error)
```

---

### Generate ⚠️ **DEPRECATED**

Generate generates a new archive from the content provided as input.

Deprecated: use archive.Generate instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/wrap_deprecated.go#L12)  

```go
func Generate(input ...string) (io.Reader, error)
```

---

### GetRebaseName ⚠️ **DEPRECATED**

GetRebaseName normalizes and compares path and resolvedPath.

Deprecated: use archive.GetRebaseName instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L128)  
**Added in:** v1.10.0

```go
func GetRebaseName(path, resolvedPath string) (string, string)
```

---

### IsArchivePath ⚠️ **DEPRECATED**

IsArchivePath checks if the (possibly compressed) file at the given path
starts with a tar file header.

Deprecated: use archive.IsArchivePath instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L96)  
**Added in:** v1.10.0

```go
func IsArchivePath(path string) bool
```

---

### IsEmpty ⚠️ **DEPRECATED**

IsEmpty checks if the tar archive is empty (doesn't contain any entries).

Deprecated: use archive.IsEmpty instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/diff_deprecated.go#L35)  

```go
func IsEmpty(rd io.Reader) (bool, error)
```

---

### NewTarballer ⚠️ **DEPRECATED**

NewTarballer constructs a new tarballer using TarWithOptions.

Deprecated: use archive.Tarballer instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L178)  

```go
func NewTarballer(srcPath string, options *TarOptions) (*archive.Tarballer, error)
```

---

### PrepareArchiveCopy ⚠️ **DEPRECATED**

PrepareArchiveCopy prepares the given srcContent archive.

Deprecated: use archive.PrepareArchiveCopy instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L90)  
**Added in:** v1.8.0

```go
func PrepareArchiveCopy(srcContent io.Reader, srcInfo, dstInfo archive.CopyInfo) (dstDir string, content io.ReadCloser, err error)
```

---

### PreserveTrailingDotOrSeparator ⚠️ **DEPRECATED**

PreserveTrailingDotOrSeparator returns the given cleaned path.

Deprecated: use archive.PreserveTrailingDotOrSeparator instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L20)  
**Added in:** v1.8.0

```go
func PreserveTrailingDotOrSeparator(cleanedPath string, originalPath string) string
```

---

### ReadSecurityXattrToTarHeader ⚠️ **DEPRECATED**

ReadSecurityXattrToTarHeader reads security.capability xattr from filesystem
to a tar header

Deprecated: use archive.ReadSecurityXattrToTarHeader instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L151)  

```go
func ReadSecurityXattrToTarHeader(path string, hdr *tar.Header) error
```

---

### RebaseArchiveEntries ⚠️ **DEPRECATED**

RebaseArchiveEntries rewrites the given srcContent archive replacing
an occurrence of oldBase with newBase at the beginning of entry names.

Deprecated: use archive.RebaseArchiveEntries instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L98)  
**Added in:** v1.10.0

```go
func RebaseArchiveEntries(srcContent io.Reader, oldBase, newBase string) io.ReadCloser
```

---

### ReplaceFileTarWrapper ⚠️ **DEPRECATED**

ReplaceFileTarWrapper converts inputTarStream to a new tar stream.

Deprecated: use archive.ReplaceFileTarWrapper instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L129)  

```go
func ReplaceFileTarWrapper(inputTarStream io.ReadCloser, mods map[string]archive.TarModifierFunc) io.ReadCloser
```

---

### ResolveHostSourcePath ⚠️ **DEPRECATED**

ResolveHostSourcePath decides real path need to be copied.

Deprecated: use archive.ResolveHostSourcePath instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L121)  
**Added in:** v1.10.0

```go
func ResolveHostSourcePath(path string, followLink bool) (resolvedPath, rebaseName string, _ error)
```

---

### SplitPathDirEntry ⚠️ **DEPRECATED**

SplitPathDirEntry splits the given path between its directory name and its
basename.

Deprecated: use archive.SplitPathDirEntry instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L28)  
**Added in:** v1.8.0

```go
func SplitPathDirEntry(path string) (dir, base string)
```

---

### Tar ⚠️ **DEPRECATED**

Tar creates an archive from the directory at `path`, and returns it as a
stream of bytes.

Deprecated: use archive.Tar instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L159)  

```go
func Tar(path string, compression archive.Compression) (io.ReadCloser, error)
```

---

### TarResource ⚠️ **DEPRECATED**

TarResource archives the resource described by the given CopyInfo to a Tar
archive.

Deprecated: use archive.TarResource instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L36)  
**Added in:** v1.8.0

```go
func TarResource(sourceInfo archive.CopyInfo) (content io.ReadCloser, err error)
```

---

### TarResourceRebase ⚠️ **DEPRECATED**

TarResourceRebase is like TarResource but renames the first path element of
items in the resulting tar archive to match the given rebaseName if not "".

Deprecated: use archive.TarResourceRebase instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L44)  
**Added in:** v1.8.0

```go
func TarResourceRebase(sourcePath, rebaseName string) (content io.ReadCloser, _ error)
```

---

### TarWithOptions ⚠️ **DEPRECATED**

TarWithOptions creates an archive with the given options.

Deprecated: use archive.TarWithOptions instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L166)  

```go
func TarWithOptions(srcPath string, options *TarOptions) (io.ReadCloser, error)
```

---

### ToArchiveOpt ⚠️ **DEPRECATED**

ToArchiveOpt converts an TarOptions to a archive.TarOptions.

Deprecated: use archive.TarOptions instead, this utility is for internal use to transition to the github.com/moby/go-archive module.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/utils.go#L11)  

```go
func ToArchiveOpt(options *TarOptions) *archive.TarOptions
```

---

### Unpack ⚠️ **DEPRECATED**

Unpack unpacks the decompressedArchive to dest with options.

Deprecated: use archive.Unpack instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L185)  
**Added in:** v1.3.3

```go
func Unpack(decompressedArchive io.Reader, dest string, options *TarOptions) error
```

---

### UnpackLayer ⚠️ **DEPRECATED**

UnpackLayer unpack `layer` to a `dest`.

Deprecated: use archive.UnpackLayer instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/diff_deprecated.go#L12)  
**Added in:** v1.3.3

```go
func UnpackLayer(dest string, layer io.Reader, options *TarOptions) (size int64, err error)
```

---

### Untar ⚠️ **DEPRECATED**

Untar reads a stream of bytes from `archive`, parses it as a tar archive,
and unpacks it into the directory at `dest`.

Deprecated: use archive.Untar instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L193)  

```go
func Untar(tarArchive io.Reader, dest string, options *TarOptions) error
```

---

### UntarUncompressed ⚠️ **DEPRECATED**

UntarUncompressed reads a stream of bytes from `tarArchive`, parses it as a tar archive,
and unpacks it into the directory at `dest`.
The archive must be an uncompressed stream.

Deprecated: use archive.UntarUncompressed instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L202)  
**Added in:** v1.8.0

```go
func UntarUncompressed(tarArchive io.Reader, dest string, options *TarOptions) error
```

---

## Types

### Archiver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L67)  
**Added in:** v1.3.2

```go
type Archiver struct {
	Untar     func(io.Reader, string, *TarOptions) error
	IDMapping idtools.IdentityMapping
}
```

---

### Change

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/changes_deprecated.go#L24)  

```go
type Change = archive.Change
```

---

### ChangeType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/changes_deprecated.go#L13)  

```go
type ChangeType = archive.ChangeType
```

---

### Compression

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L25)  

```go
type Compression = compression.Compression
```

---

### CopyInfo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/copy_deprecated.go#L67)  
**Added in:** v1.8.0

```go
type CopyInfo = archive.CopyInfo
```

---

### FileInfo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/changes_deprecated.go#L37)  

```go
type FileInfo = archive.FileInfo
```

---

### TarModifierFunc

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L124)  

```go
type TarModifierFunc = archive.TarModifierFunc
```

---

### TarOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L34)  

```go
type TarOptions struct {
	IncludeFiles     []string
	ExcludePatterns  []string
	Compression      compression.Compression
	NoLchown         bool
	IDMap            idtools.IdentityMapping
	ChownOpts        *idtools.Identity
	IncludeSourceDir bool
	// WhiteoutFormat is the expected on disk format for whiteout files.
	// This format will be converted to the standard format on pack
	// and from the standard format on unpack.
	WhiteoutFormat archive.WhiteoutFormat
	// When unpacking, specifies whether overwriting a directory with a
	// non-directory is allowed and vice versa.
	NoOverwriteDirNonDir bool
	// For each include when creating an archive, the included name will be
	// replaced with the matching name from this map.
	RebaseNames map[string]string
	InUserNS    bool
	// Allow unpacking to succeed in spite of failures to set extended
	// attributes on the unpacked files due to the destination filesystem
	// not supporting them or a lack of permissions. Extended attributes
	// were probably in the archive for a reason, so set this option at
	// your own peril.
	BestEffortXattrs bool
}
```

---

### Tarballer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L173)  

```go
type Tarballer = archive.Tarballer
```

---

### WhiteoutFormat

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/archive/archive_deprecated.go#L29)  
**Added in:** v1.12.0

```go
type WhiteoutFormat = archive.WhiteoutFormat
```

---


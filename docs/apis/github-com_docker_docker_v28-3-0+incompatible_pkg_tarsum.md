# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/tarsum

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:23 UTC

## Overview

Package tarsum provides algorithms to perform checksum calculation on
filesystem layers.

The transportation of filesystems, regarding Docker, is done with tar(1)
archives. There are a variety of tar serialization formats [2], and a key
concern here is ensuring a repeatable checksum given a set of inputs from a
generic tar archive. Types of transportation include distribution to and from a
registry endpoint, saving and loading through commands or Docker daemon APIs,
transferring the build context from client to Docker daemon, and committing the
filesystem of a container to become an image.

As tar archives are used for transit, but not preserved in many situations, the
focus of the algorithm is to ensure the integrity of the preserved filesystem,
while maintaining a deterministic accountability. This includes neither
constraining the ordering or manipulation of the files during the creation or
unpacking of the archive, nor include additional metadata state about the file
system attributes.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/versioning.go#L81)

```go
var (
	ErrNotVersion            = errors.New("string does not include a TarSum Version")
	ErrVersionNotImplemented = errors.New("TarSum Version is not yet implemented")
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/tarsum.go#L147)

```go
var DefaultTHash = NewTHash("sha256", sha256.New)
```

## Functions

### VersionLabelForChecksum

VersionLabelForChecksum returns the label for the given tarsum
checksum, i.e., everything before the first `+` character in
the string or an empty string if no label separator is found.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/versioning.go#L35)  
**Added in:** v1.5.0

```go
func VersionLabelForChecksum(checksum string) string
```

---

### WriteV1Header

WriteV1Header writes a tar header to a writer in V1 tarsum format.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/versioning.go#L26)  

```go
func WriteV1Header(h *tar.Header, w io.Writer)
```

---

## Types

### BuilderContext

BuilderContext is an interface extending TarSum by adding the Remove method.
In general there was concern about adding this method to TarSum itself
so instead it is being added just to "BuilderContext" which will then
only be used during the .dockerignore file processing
- see builder/evaluator.go

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/builder_context.go#L8)  
**Added in:** v1.5.0

```go
type BuilderContext interface {
	TarSum
	Remove(string)
}
```

---

### FileInfoSumInterface

FileInfoSumInterface provides an interface for accessing file checksum
information within a tar file. This info is accessed through interface
so the actual name and sum cannot be melded with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L12)  
**Added in:** v1.3.0

```go
type FileInfoSumInterface interface {
	// File name
	Name() string
	// Checksum of this particular file and its headers
	Sum() string
	// Position of file in the tar
	Pos() int64
}
```

---

### FileInfoSums

FileInfoSums provides a list of FileInfoSumInterfaces.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L40)  
**Added in:** v1.3.0

```go
type FileInfoSums []FileInfoSumInterface
```

#### Methods

##### FileInfoSums.GetAllFile

GetAllFile returns a FileInfoSums with all matching names.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L56)  
**Added in:** v1.3.0

```go
func (fis FileInfoSums) GetAllFile(name string) FileInfoSums
```

##### FileInfoSums.GetDuplicatePaths

GetDuplicatePaths returns a FileInfoSums with all duplicated paths.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L67)  
**Added in:** v1.3.0

```go
func (fis FileInfoSums) GetDuplicatePaths() (dups FileInfoSums)
```

##### FileInfoSums.GetFile

GetFile returns the first FileInfoSumInterface with a matching name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L43)  
**Added in:** v1.3.0

```go
func (fis FileInfoSums) GetFile(name string) FileInfoSumInterface
```

##### FileInfoSums.Len

Len returns the size of the FileInfoSums.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L81)  
**Added in:** v1.3.0

```go
func (fis FileInfoSums) Len() int
```

##### FileInfoSums.SortByNames

SortByNames sorts FileInfoSums content by name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L92)  
**Added in:** v1.3.0

```go
func (fis FileInfoSums) SortByNames()
```

##### FileInfoSums.SortByPos

SortByPos sorts FileInfoSums content by position.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L87)  
**Added in:** v1.3.0

```go
func (fis FileInfoSums) SortByPos()
```

##### FileInfoSums.SortBySums

SortBySums sorts FileInfoSums content by sums.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L97)  
**Added in:** v1.3.0

```go
func (fis FileInfoSums) SortBySums()
```

##### FileInfoSums.Swap

Swap swaps two FileInfoSum values if a FileInfoSums list.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/fileinfosums.go#L84)  
**Added in:** v1.3.0

```go
func (fis FileInfoSums) Swap(i, j int)
```

---

### THash

THash provides a hash.Hash type generator and its name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/tarsum.go#L125)  
**Added in:** v1.3.0

```go
type THash interface {
	Hash() hash.Hash
	Name() string
}
```

#### Functions

##### NewTHash

NewTHash is a convenience method for creating a THash.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/tarsum.go#L131)  
**Added in:** v1.3.0

```go
func NewTHash(name string, h func() hash.Hash) THash
```

---

### TarSum

TarSum is the generic interface for calculating fixed time
checksums of a tar archive.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/tarsum.go#L87)  

```go
type TarSum interface {
	io.Reader
	GetSums() FileInfoSums
	Sum([]byte) string
	Version() Version
	Hash() THash
}
```

#### Functions

##### NewTarSum

NewTarSum creates a new interface for calculating a fixed time checksum of a
tar archive.

This is used for calculating checksums of layers of an image, in some cases
including the byte payload of the image's json metadata as well, and for
calculating the checksums for buildcache.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/tarsum.go#L47)  
**Added in:** v1.3.0

```go
func NewTarSum(r io.Reader, dc bool, v Version) (TarSum, error)
```

##### NewTarSumForLabel

NewTarSumForLabel creates a new TarSum using the provided TarSum version+hash label.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/tarsum.go#L64)  
**Added in:** v1.5.0

```go
func NewTarSumForLabel(r io.Reader, disableCompression bool, label string) (TarSum, error)
```

##### NewTarSumHash

NewTarSumHash creates a new TarSum, providing a THash to use rather than
the DefaultTHash.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/tarsum.go#L53)  
**Added in:** v1.3.0

```go
func NewTarSumHash(r io.Reader, dc bool, v Version, tHash THash) (TarSum, error)
```

---

### Version

Version is used for versioning of the TarSum algorithm
based on the prefix of the hash used
i.e. "tarsum+sha256:e58fcf7418d4390dec8e8fb69d88c06ec07039d651fedd3aa72af9972e7d046b"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/versioning.go#L15)  
**Added in:** v1.3.0

```go
type Version int
```

#### Functions

##### GetVersionFromTarsum

GetVersionFromTarsum returns the Version from the provided string.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/versioning.go#L71)  
**Added in:** v1.3.0

```go
func GetVersionFromTarsum(tarsum string) (Version, error)
```

##### GetVersions

GetVersions gets a list of all known tarsum versions.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/versioning.go#L45)  
**Added in:** v1.3.0

```go
func GetVersions() []Version
```

#### Methods

##### Version.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/tarsum/versioning.go#L66)  
**Added in:** v1.3.0

```go
func (tsv Version) String() string
```

---


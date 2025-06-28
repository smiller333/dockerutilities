# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/containerfs

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:31:06 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### EnsureRemoveAll

EnsureRemoveAll wraps os.RemoveAll to check for specific errors that can
often be remedied.
Only use EnsureRemoveAll if you really want to make every effort to remove
a directory.

Because of the way os.Remove (and by extension os.RemoveAll) works, there
can be a race between reading directory entries and then actually attempting
to remove everything in the directory.
These types of errors do not need to be returned since it's ok for the dir to
be gone we can just retry the remove operation.

This should not return a os.ErrNotExist kind of error under any circumstances.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/containerfs/rm.go#L26)  

```go
func EnsureRemoveAll(dir string) error
```

---

## Types

This section is empty.


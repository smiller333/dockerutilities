# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/safepath

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:42 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### ErrEscapesBase

ErrEscapesBase is returned by Join when the resulting concatenation would
point outside of the specified base directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/errors.go#L27)  

```go
type ErrEscapesBase struct {
	Base, Subpath string
}
```

#### Methods

##### ErrEscapesBase.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/errors.go#L33)  

```go
func (e *ErrEscapesBase) Error() string
```

##### ErrEscapesBase.InvalidParameter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/errors.go#L31)  

```go
func (*ErrEscapesBase) InvalidParameter()
```

---

### ErrNotAccessible

ErrNotAccessible is returned by Join when the resulting path doesn't exist,
is not accessible, or any of the path components was replaced with a symlink
during the path traversal.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/errors.go#L6)  

```go
type ErrNotAccessible struct {
	Path  string
	Cause error
}
```

#### Methods

##### ErrNotAccessible.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/errors.go#L17)  

```go
func (e *ErrNotAccessible) Error() string
```

##### ErrNotAccessible.NotFound

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/errors.go#L11)  

```go
func (*ErrNotAccessible) NotFound()
```

##### ErrNotAccessible.Unwrap

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/errors.go#L13)  

```go
func (e *ErrNotAccessible) Unwrap() error
```

---

### SafePath

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/safepath.go#L11)  

```go
type SafePath struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### Join

Join makes sure that the concatenation of path and subpath doesn't
resolve to a path outside of path and returns a path to a temporary file that is
a bind mount to the exact same file/directory that was validated.

After use, it is the caller's responsibility to call Close on the returned
SafePath object, which will unmount the temporary file/directory
and remove it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/join_linux.go#L23)  

```go
func Join(ctx context.Context, path, subpath string) (*SafePath, error)
```

#### Methods

##### SafePath.Close

Close releases the resources used by the path.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/safepath.go#L21)  

```go
func (s *SafePath) Close(ctx context.Context) error
```

##### SafePath.IsValid

IsValid return true when path can still be used and wasn't cleaned up by Close.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/safepath.go#L43)  

```go
func (s *SafePath) IsValid() bool
```

##### SafePath.Path

Path returns a safe, temporary path that can be used to access the original path.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/safepath.go#L50)  

```go
func (s *SafePath) Path() string
```

##### SafePath.SourcePath

SourcePath returns the source path the safepath points to.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/safepath/safepath.go#L60)  

```go
func (s *SafePath) SourcePath() (string, string)
```

---


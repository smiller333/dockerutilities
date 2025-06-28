# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/internal/setmatrix

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:10:10 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### SetMatrix

SetMatrix is a map of Sets.
The zero value is an empty set matrix ready to use.

SetMatrix values are safe for concurrent use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/setmatrix/setmatrix.go#L16)  

```go
type SetMatrix[T comparable] struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### SetMatrix.Cardinality

Cardinality returns the number of elements in the set for a key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/setmatrix/setmatrix.go#L83)  

```go
func (s *SetMatrix[T]) Cardinality(key string) (cardinality int, ok bool)
```

##### SetMatrix.Contains

Contains is used to verify if an element is in a set for a specific key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/setmatrix/setmatrix.go#L34)  

```go
func (s *SetMatrix[T]) Contains(key string, value T) (containsElement, setExists bool)
```

##### SetMatrix.Get

Get returns the members of the set for a specific key as a slice.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/setmatrix/setmatrix.go#L23)  

```go
func (s *SetMatrix[T]) Get(key string) ([]T, bool)
```

##### SetMatrix.Insert

Insert inserts the value in the set of a key and returns whether the value is
inserted (was not already in the set) and the number of elements in the set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/setmatrix/setmatrix.go#L46)  

```go
func (s *SetMatrix[T]) Insert(key string, value T) (inserted bool, cardinality int)
```

##### SetMatrix.Keys

Keys returns all the keys in the map.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/setmatrix/setmatrix.go#L107)  

```go
func (s *SetMatrix[T]) Keys() []string
```

##### SetMatrix.Remove

Remove removes the value in the set for a specific key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/setmatrix/setmatrix.go#L62)  

```go
func (s *SetMatrix[T]) Remove(key string, value T) (removed bool, cardinality int)
```

##### SetMatrix.String

String returns the string version of the set.
The empty string is returned if there is no set for key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/setmatrix/setmatrix.go#L96)  

```go
func (s *SetMatrix[T]) String(key string) (v string, ok bool)
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/reference

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:04 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/reference/store.go#L19)

```go
var ErrDoesNotExist notFoundError = "reference does not exist"
```

## Functions

This section is empty.

## Types

### Association

An Association is a tuple associating a reference with an image ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/reference/store.go#L22)  

```go
type Association struct {
	Ref reference.Named
	ID  digest.Digest
}
```

---

### Store

Store provides the set of methods which can operate on a reference store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/reference/store.go#L28)  

```go
type Store interface {
	References(id digest.Digest) []reference.Named
	ReferencesByName(ref reference.Named) []Association
	AddTag(ref reference.Named, id digest.Digest, force bool) error
	AddDigest(ref reference.Canonical, id digest.Digest, force bool) error
	Delete(ref reference.Named) (bool, error)
	Get(ref reference.Named) (digest.Digest, error)
}
```

#### Functions

##### NewReferenceStore

NewReferenceStore creates a new reference store, tied to a file path where
the set of references are serialized in JSON format.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/reference/store.go#L71)  

```go
func NewReferenceStore(jsonPath string) (Store, error)
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/testutils/labelstore

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:31:57 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### InMemory

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/labelstore/memory_label_store.go#L9)  

```go
type InMemory struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### InMemory.Get

Get returns all the labels for the given digest

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/labelstore/memory_label_store.go#L15)  

```go
func (s *InMemory) Get(dgst digest.Digest) (map[string]string, error)
```

##### InMemory.Set

Set sets all the labels for a given digest

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/labelstore/memory_label_store.go#L23)  

```go
func (s *InMemory) Set(dgst digest.Digest, labels map[string]string) error
```

##### InMemory.Update

Update replaces the given labels for a digest,
a key with an empty value removes a label.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/labelstore/memory_label_store.go#L35)  

```go
func (s *InMemory) Update(dgst digest.Digest, update map[string]string) (map[string]string, error)
```

---


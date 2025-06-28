# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/distribution/metadata

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:53 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CheckV2MetadataHMAC

CheckV2MetadataHMAC returns true if the given "meta" is tagged with a hmac hashed by the given "key".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/v2_metadata_service.go#L42)  
**Added in:** v1.13.0

```go
func CheckV2MetadataHMAC(meta *V2Metadata, key []byte) bool
```

---

### ComputeV2MetadataHMAC

ComputeV2MetadataHMAC returns a hmac for the given "meta" hash by the given key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/v2_metadata_service.go#L60)  
**Added in:** v1.13.0

```go
func ComputeV2MetadataHMAC(key []byte, meta *V2Metadata) string
```

---

### ComputeV2MetadataHMACKey

ComputeV2MetadataHMACKey returns a key for the given "authConfig" that can be used to hash v2 metadata
entries.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/v2_metadata_service.go#L72)  
**Added in:** v1.13.0

```go
func ComputeV2MetadataHMACKey(authConfig *registry.AuthConfig) ([]byte, error)
```

---

## Types

### FSMetadataStore

FSMetadataStore uses the filesystem to associate metadata with layer and
image IDs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/metadata.go#L25)  

```go
type FSMetadataStore struct {
	sync.RWMutex
	// contains filtered or unexported fields
}
```

#### Functions

##### NewFSMetadataStore

NewFSMetadataStore creates a new filesystem-based metadata store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/metadata.go#L31)  

```go
func NewFSMetadataStore(basePath string) (*FSMetadataStore, error)
```

#### Methods

##### FSMetadataStore.Delete

Delete removes data indexed by namespace and key. The data file named after
the key, stored in the namespace's directory is deleted.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/metadata.go#L68)  

```go
func (store *FSMetadataStore) Delete(namespace, key string) error
```

##### FSMetadataStore.Get

Get retrieves data by namespace and key. The data is read from a file named
after the key, stored in the namespace's directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/metadata.go#L46)  

```go
func (store *FSMetadataStore) Get(namespace string, key string) ([]byte, error)
```

##### FSMetadataStore.Set

Set writes data indexed by namespace and key. The data is written to a file
named after the key, stored in the namespace's directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/metadata.go#L55)  

```go
func (store *FSMetadataStore) Set(namespace, key string, value []byte) error
```

---

### Store

Store implements a K/V store for mapping distribution-related IDs
to on-disk layer IDs and image IDs. The namespace identifies the type of
mapping (i.e. "v1ids" or "artifacts"). MetadataStore is goroutine-safe.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/metadata.go#L14)  

```go
type Store interface {
	// Get retrieves data by namespace and key.
	Get(namespace string, key string) ([]byte, error)
	// Set writes data indexed by namespace and key.
	Set(namespace, key string, value []byte) error
	// Delete removes data indexed by namespace and key.
	Delete(namespace, key string) error
}
```

---

### V2Metadata

V2Metadata contains the digest and source repository information for a layer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/v2_metadata_service.go#L33)  

```go
type V2Metadata struct {
	Digest           digest.Digest
	SourceRepository string
	// HMAC hashes above attributes with recent authconfig digest used as a key in order to determine matching
	// metadata entries accompanied by the same credentials without actually exposing them.
	HMAC string
}
```

---

### V2MetadataService

V2MetadataService maps layer IDs to a set of known metadata for
the layer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/v2_metadata_service.go#L17)  

```go
type V2MetadataService interface {
	GetMetadata(diffID layer.DiffID) ([]V2Metadata, error)
	GetDiffID(dgst digest.Digest) (layer.DiffID, error)
	Add(diffID layer.DiffID, metadata V2Metadata) error
	TagAndAdd(diffID layer.DiffID, hmacKey []byte, metadata V2Metadata) error
	Remove(metadata V2Metadata) error
}
```

#### Functions

##### NewV2MetadataService

NewV2MetadataService creates a new diff ID to v2 metadata mapping service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/metadata/v2_metadata_service.go#L105)  

```go
func NewV2MetadataService(store Store) V2MetadataService
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/datastore

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:33:00 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L53)

```go
const (
	// NetworkKeyPrefix is the prefix for network key in the kv store
	NetworkKeyPrefix = "network"
	// EndpointKeyPrefix is the prefix for endpoint key in the kv store
	EndpointKeyPrefix = "endpoint"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L65)

```go
const DefaultBucket = "libnetwork"
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L15)

```go
var (
	ErrKeyModified = store.ErrKeyModified
	ErrKeyNotFound = store.ErrKeyNotFound
)
```

## Functions

### Key

Key provides convenient method to create a Key

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L68)  

```go
func Key(key ...string) string
```

---

## Types

### KVObject

KVObject is Key/Value interface used by objects to be part of the Store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L27)  

```go
type KVObject interface {
	// Key method lets an object provide the Key to be used in KV Store
	Key() []string
	// KeyPrefix method lets an object return immediate parent key that can be used for tree walk
	KeyPrefix() []string
	// Value method lets an object marshal its content to be stored in the KV store
	Value() []byte
	// SetValue is used by the datastore to set the object's value when loaded from the data store.
	SetValue([]byte) error
	// Index method returns the latest DB Index as seen by the object
	Index() uint64
	// SetIndex method allows the datastore to store the latest DB Index into the object
	SetIndex(uint64)
	// Exists returns true if the object exists in the datastore, false if it hasn't been stored yet.
	// When SetIndex() is called, the object has been stored.
	Exists() bool
	// Skip provides a way for a KV Object to avoid persisting it in the KV Store
	Skip() bool
	// New returns a new object which is created based on the
	// source object
	New() KVObject
	// CopyTo deep copies the contents of the implementing object
	// to the passed destination object
	CopyTo(KVObject) error
}
```

---

### Store

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L20)  

```go
type Store struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates a new Store instance.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L80)  

```go
func New(dir, bucket string) (*Store, error)
```

#### Methods

##### Store.Close

Close closes the data store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L97)  

```go
func (ds *Store) Close()
```

##### Store.DeleteObject

DeleteObject deletes a kvObject from the on-disk DB and the in-memory cache.
Unlike DeleteObjectAtomic, it doesn't check the optimistic lock of the
passed kvObject.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L215)  

```go
func (ds *Store) DeleteObject(kvObject KVObject) error
```

##### Store.DeleteObjectAtomic

DeleteObjectAtomic performs atomic delete on a record.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L236)  

```go
func (ds *Store) DeleteObjectAtomic(kvObject KVObject) error
```

##### Store.GetObject

GetObject gets data from the store and unmarshals to the specified object.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L139)  

```go
func (ds *Store) GetObject(o KVObject) error
```

##### Store.List

List returns of a list of KVObjects belonging to the parent key. The caller
must pass a KVObject of the same type as the objects that need to be listed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L159)  

```go
func (ds *Store) List(kvObject KVObject) ([]KVObject, error)
```

##### Store.Map

Map returns a Map of KVObjects.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L197)  

```go
func (ds *Store) Map(key string, kvObject KVObject) (map[string]KVObject, error)
```

##### Store.PutObjectAtomic

PutObjectAtomic provides an atomic add and update operation for a Record.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/datastore/datastore.go#L102)  

```go
func (ds *Store) PutObjectAtomic(kvObject KVObject) error
```

---


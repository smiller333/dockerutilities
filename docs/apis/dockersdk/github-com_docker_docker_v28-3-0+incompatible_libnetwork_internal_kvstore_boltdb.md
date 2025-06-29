# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/internal/kvstore/boltdb

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:09:59 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New opens a new BoltDB connection to the specified path and bucket

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/kvstore/boltdb/boltdb.go#L33)  

```go
func New(path, bucket string) (store.Store, error)
```

---

## Types

### BoltDB

BoltDB type implements the Store interface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/kvstore/boltdb/boltdb.go#L22)  

```go
type BoltDB struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### BoltDB.AtomicDelete

AtomicDelete deletes a value at "key" if the key
has not been modified in the meantime, throws an
error if this is the case

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/kvstore/boltdb/boltdb.go#L151)  

```go
func (b *BoltDB) AtomicDelete(key string, previous *store.KVPair) error
```

##### BoltDB.AtomicPut

AtomicPut puts a value at "key" if the key has not been
modified since the last Put, throws an error if this is the case

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/kvstore/boltdb/boltdb.go#L194)  

```go
func (b *BoltDB) AtomicPut(key string, value []byte, previous *store.KVPair) (*store.KVPair, error)
```

##### BoltDB.Close

Close the db connection to the BoltDB

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/kvstore/boltdb/boltdb.go#L239)  

```go
func (b *BoltDB) Close()
```

##### BoltDB.Delete

Delete deletes a value at "key". Unlike AtomicDelete it doesn't check
whether the deleted key is at a specific version before deleting.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/kvstore/boltdb/boltdb.go#L179)  

```go
func (b *BoltDB) Delete(key string) error
```

##### BoltDB.Exists

Exists checks if the key exists inside the store

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/kvstore/boltdb/boltdb.go#L87)  

```go
func (b *BoltDB) Exists(key string) (bool, error)
```

##### BoltDB.List

List returns the range of keys starting with the passed in prefix

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/kvstore/boltdb/boltdb.go#L111)  

```go
func (b *BoltDB) List(keyPrefix string) ([]*store.KVPair, error)
```

##### BoltDB.Put

Put the key, value pair. index number metadata is prepended to the value

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/kvstore/boltdb/boltdb.go#L67)  

```go
func (b *BoltDB) Put(key string, value []byte) error
```

---


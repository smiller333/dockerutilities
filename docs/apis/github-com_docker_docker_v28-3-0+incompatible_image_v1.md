# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/image/v1

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:13 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CreateID

CreateID creates an ID from v1 image, layerID and parent ID.
Used for backwards compatibility with old clients.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/v1/imagev1.go#L44)  

```go
func CreateID(v1Image image.V1Image, layerID layer.ChainID, parent digest.Digest) (digest.Digest, error)
```

---

### HistoryFromConfig

HistoryFromConfig creates a History struct from v1 configuration JSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/v1/imagev1.go#L26)  

```go
func HistoryFromConfig(imageJSON []byte, emptyLayer bool) (image.History, error)
```

---

### MakeConfigFromV1Config

MakeConfigFromV1Config creates an image config from the legacy V1 config format.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/v1/imagev1.go#L72)  

```go
func MakeConfigFromV1Config(imageJSON []byte, rootfs *image.RootFS, history []image.History) ([]byte, error)
```

---

### ValidateID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/v1/imagev1.go#L124)  

```go
func ValidateID(id string) error
```

---

## Types

This section is empty.


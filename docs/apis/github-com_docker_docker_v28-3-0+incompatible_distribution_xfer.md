# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/distribution/xfer

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:54 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### IsDoNotRetryError

IsDoNotRetryError returns true if the error is caused by DoNotRetry error,
and the transfer should not be retried.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/transfer.go#L25)  

```go
func IsDoNotRetryError(err error) bool
```

---

## Types

### DigestRegisterer

DigestRegisterer can be implemented by a DownloadDescriptor, and provides a
Registered method which gets called after a downloaded layer is registered.
This allows the user of the download manager to know the DiffID of each
registered layer. This method is called if a cast to DigestRegisterer is
successful.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/download.go#L98)  

```go
type DigestRegisterer interface {
	Registered(diffID layer.DiffID)
}
```

---

### DoNotRetry

DoNotRetry is an error wrapper indicating that the error cannot be resolved
with a retry.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/transfer.go#L14)  

```go
type DoNotRetry struct {
	Err error
}
```

#### Methods

##### DoNotRetry.Error

Error returns the stringified representation of the encapsulated error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/transfer.go#L19)  

```go
func (e DoNotRetry) Error() string
```

---

### DownloadDescriptor

A DownloadDescriptor references a layer that may need to be downloaded.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/download.go#L76)  

```go
type DownloadDescriptor interface {
	// Key returns the key used to deduplicate downloads.
	Key() string
	// ID returns the ID for display purposes.
	ID() string
	// DiffID should return the DiffID for this layer, or an error
	// if it is unknown (for example, if it has not been downloaded
	// before).
	DiffID() (layer.DiffID, error)
	// Download is called to perform the download.
	Download(ctx context.Context, progressOutput progress.Output) (io.ReadCloser, int64, error)
	// Close is called when the download manager is finished with this
	// descriptor and will not call Download again or read from the reader
	// that Download returned.
	Close()
}
```

---

### DownloadOption

DownloadOption set options for the LayerDownloadManager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/download.go#L51)  

```go
type DownloadOption func(*LayerDownloadManager)
```

#### Functions

##### WithMaxDownloadAttempts

WithMaxDownloadAttempts configures the maximum number of download
attempts for a download manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/download.go#L55)  

```go
func WithMaxDownloadAttempts(max int) DownloadOption
```

---

### LayerDownloadManager

LayerDownloadManager figures out which layers need to be downloaded, then
registers and downloads those, taking into account dependencies between
layers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/download.go#L24)  

```go
type LayerDownloadManager struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewLayerDownloadManager

NewLayerDownloadManager returns a new LayerDownloadManager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/download.go#L37)  

```go
func NewLayerDownloadManager(layerStore layer.Store, concurrencyLimit int, options ...DownloadOption) *LayerDownloadManager
```

#### Methods

##### LayerDownloadManager.Download

Download is a blocking function which ensures the requested layers are
present in the layer store. It uses the string returned by the Key method to
deduplicate downloads. If a given layer is not already known to present in
the layer store, and the key is not used by an in-progress download, the
Download method is called to get the layer tar data. Layers are then
registered in the appropriate order.  The caller must call the returned
release function once it is done with the returned RootFS object.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/download.go#L113)  

```go
func (ldm *LayerDownloadManager) Download(ctx context.Context, initialRootFS image.RootFS, layers []DownloadDescriptor, progressOutput progress.Output) (image.RootFS, func(), error)
```

##### LayerDownloadManager.SetConcurrency

SetConcurrency sets the max concurrent downloads for each pull

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/download.go#L32)  
**Added in:** v1.12.0

```go
func (ldm *LayerDownloadManager) SetConcurrency(concurrency int)
```

---

### LayerUploadManager

LayerUploadManager provides task management and progress reporting for
uploads.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/upload.go#L18)  

```go
type LayerUploadManager struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewLayerUploadManager

NewLayerUploadManager returns a new LayerUploadManager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/upload.go#L29)  

```go
func NewLayerUploadManager(concurrencyLimit int, options ...func(*LayerUploadManager)) *LayerUploadManager
```

#### Methods

##### LayerUploadManager.SetConcurrency

SetConcurrency sets the max concurrent uploads for each push

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/upload.go#L24)  
**Added in:** v1.12.0

```go
func (lum *LayerUploadManager) SetConcurrency(concurrency int)
```

##### LayerUploadManager.Upload

Upload is a blocking function which ensures the listed layers are present on
the remote registry. It uses the string returned by the Key method to
deduplicate uploads.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/upload.go#L67)  

```go
func (lum *LayerUploadManager) Upload(ctx context.Context, layers []UploadDescriptor, progressOutput progress.Output) error
```

---

### UploadDescriptor

An UploadDescriptor references a layer that may need to be uploaded.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/distribution/xfer/upload.go#L48)  

```go
type UploadDescriptor interface {
	// Key returns the key used to deduplicate uploads.
	Key() string
	// ID returns the ID for display purposes.
	ID() string
	// DiffID should return the DiffID for this layer.
	DiffID() layer.DiffID
	// Upload is called to perform the Upload.
	Upload(ctx context.Context, progressOutput progress.Output) (distribution.Descriptor, error)
	// SetRemoteDescriptor provides the distribution.Descriptor that was
	// returned by Upload. This descriptor is not to be confused with
	// the UploadDescriptor interface, which is used for internally
	// identifying layers that are being uploaded.
	SetRemoteDescriptor(descriptor distribution.Descriptor)
}
```

---


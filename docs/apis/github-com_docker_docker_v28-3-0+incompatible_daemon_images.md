# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/images

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:28:36 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### OnlyPlatformWithFallback

OnlyPlatformWithFallback uses `platforms.Only` with a fallback to handle the case where the platform
being matched does not have a CPU variant.

The reason for this is that CPU variant is not even if the official image config spec as of this writing.
See: https://github.com/opencontainers/image-spec/pull/809
Since Docker tends to compare platforms from the image config, we need to handle this case.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image.go#L241)  

```go
func OnlyPlatformWithFallback(p ocispec.Platform) platforms.Matcher
```

---

## Types

### DistributionServices

DistributionServices provides daemon image storage services

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L86)  

```go
type DistributionServices struct {
	DownloadManager   *xfer.LayerDownloadManager
	V2MetadataService metadata.V2MetadataService
	LayerStore        layer.Store
	ImageStore        image.Store
	ReferenceStore    dockerreference.Store
}
```

---

### ErrImageDoesNotExist

ErrImageDoesNotExist is error returned when no image can be found for a reference.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image.go#L25)  

```go
type ErrImageDoesNotExist struct {
	Ref reference.Reference
}
```

#### Methods

##### ErrImageDoesNotExist.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image.go#L29)  

```go
func (e ErrImageDoesNotExist) Error() string
```

##### ErrImageDoesNotExist.NotFound

NotFound implements the NotFound interface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image.go#L38)  

```go
func (e ErrImageDoesNotExist) NotFound()
```

---

### ImageService

ImageService provides a backend for image management

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L69)  

```go
type ImageService struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewImageService

NewImageService returns a new ImageService from a configuration

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L51)  

```go
func NewImageService(config ImageServiceConfig) *ImageService
```

#### Methods

##### ImageService.Changes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_changes.go#L13)  

```go
func (i *ImageService) Changes(ctx context.Context, container *container.Container) ([]archive.Change, error)
```

##### ImageService.Children

Children returns the children image.IDs for a parent image.
called from list.go to filter containers
TODO: refactor to expose an ancestry for image.ID?

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L114)  

```go
func (i *ImageService) Children(_ context.Context, id image.ID) ([]image.ID, error)
```

##### ImageService.Cleanup

Cleanup resources before the process is shutdown.
called from daemon.go Daemon.Shutdown()

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L172)  

```go
func (i *ImageService) Cleanup() error
```

##### ImageService.CommitBuildStep

CommitBuildStep is used by the builder to create an image for each step in
the build.

This method is different from CreateImageFromContainer:

This is a temporary shim. Should be removed when builder stops using commit.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_commit.go#L122)  

```go
func (i *ImageService) CommitBuildStep(ctx context.Context, c backend.CommitConfig) (image.ID, error)
```

##### ImageService.CommitImage

CommitImage creates a new image from a commit config

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_commit.go#L17)  

```go
func (i *ImageService) CommitImage(ctx context.Context, c backend.CommitConfig) (image.ID, error)
```

##### ImageService.CountImages

CountImages returns the number of images stored by ImageService
called from info.go

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L107)  

```go
func (i *ImageService) CountImages(ctx context.Context) int
```

##### ImageService.CreateImage

CreateImage creates a new image by adding a config and ID to the image store.
This is similar to LoadImage() except that it receives JSON encoded bytes of
an image instead of a tar archive.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_builder.go#L240)  

```go
func (i *ImageService) CreateImage(ctx context.Context, config []byte, parent string, _ digest.Digest) (builder.Image, error)
```

##### ImageService.CreateLayer

CreateLayer creates a filesystem layer for a container.
called from create.go
TODO: accept an opt struct instead of container?

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L121)  

```go
func (i *ImageService) CreateLayer(container *container.Container, initFunc layer.MountInit) (container.RWLayer, error)
```

##### ImageService.CreateLayerFromImage

CreateLayerFromImage creates a file system from an arbitrary image
Used to mount an image inside another

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L142)  

```go
func (i *ImageService) CreateLayerFromImage(img *image.Image, layerName string, rwLayerOpts *layer.CreateRWLayerOpts) (container.RWLayer, error)
```

##### ImageService.DistributionServices

DistributionServices return services controlling daemon image storage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L95)  

```go
func (i *ImageService) DistributionServices() DistributionServices
```

##### ImageService.ExportImage

ExportImage exports a list of images to the given output stream. The
exported images are archived into a tar when written to the output
stream. All images with the given tag and all versions containing
the same tag are exported. names is the set of tags to export, and
outStream is the writer which the images are written to.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_exporter.go#L16)  

```go
func (i *ImageService) ExportImage(ctx context.Context, names []string, platform *ocispec.Platform, outStream io.Writer) error
```

##### ImageService.GetContainerLayerSize

GetContainerLayerSize returns the real size & virtual size of the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_unix.go#L20)  

```go
func (i *ImageService) GetContainerLayerSize(ctx context.Context, containerID string) (int64, int64, error)
```

##### ImageService.GetImage

GetImage returns an image corresponding to the image referred to by refOrID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image.go#L161)  

```go
func (i *ImageService) GetImage(ctx context.Context, refOrID string, options backend.GetImageOpts) (retImg *image.Image, retErr error)
```

##### ImageService.GetImageAndReleasableLayer

GetImageAndReleasableLayer returns an image and releaseable layer for a reference or ID.
Every call to GetImageAndReleasableLayer MUST call releasableLayer.Release() to prevent
leaking of layers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_builder.go#L195)  

```go
func (i *ImageService) GetImageAndReleasableLayer(ctx context.Context, refOrID string, opts backend.GetImageAndLayerOptions) (builder.Image, builder.ROLayer, error)
```

##### ImageService.GetLayerByID

GetLayerByID returns a layer by ID
called from daemon.go Daemon.restore().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L153)  

```go
func (i *ImageService) GetLayerByID(cid string) (container.RWLayer, error)
```

##### ImageService.GetLayerFolders

GetLayerFolders returns the layer folders from an image RootFS

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_unix.go#L14)  

```go
func (i *ImageService) GetLayerFolders(img *image.Image, rwLayer container.RWLayer, containerID string) ([]string, error)
```

##### ImageService.GetLayerMountID

GetLayerMountID returns the mount ID for a layer
called from daemon.go Daemon.Shutdown(), and Daemon.Cleanup() (cleanup is actually containerCleanup)
TODO: needs to be refactored to Unmount (see callers), or removed and replaced with GetLayerByID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L166)  

```go
func (i *ImageService) GetLayerMountID(cid string) (string, error)
```

##### ImageService.ImageDelete

ImageDelete deletes the image referenced by the given imageRef from this
daemon. The given imageRef can be an image ID, ID prefix, or a repository
reference (with an optional tag or digest, defaulting to the tag name
"latest"). There is differing behavior depending on whether the given
imageRef is a repository reference or not.

If the given imageRef is a repository reference then that repository
reference will be removed. However, if there exists any containers which
were created using the same image reference then the repository reference
cannot be removed unless either there are other repository references to the
same image or options.Force is true. Following removal of the repository reference,
the referenced image itself will attempt to be deleted as described below
but quietly, meaning any image delete conflicts will cause the image to not
be deleted and the conflict will not be reported.

There may be conflicts preventing deletion of an image and these conflicts
are divided into two categories grouped by their severity:

Hard Conflict:

Soft Conflict:

The image cannot be removed if there are any hard conflicts and can be
removed if there are soft conflicts only if options.Force is true.

If options.PruneChildren is true, ancestor images are attempted to be deleted quietly,
meaning any delete conflicts will cause the image to not be deleted and the
conflict will not be reported.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_delete.go#L66)  

```go
func (i *ImageService) ImageDelete(ctx context.Context, imageRef string, options imagetypes.RemoveOptions) ([]imagetypes.DeleteResponse, error)
```

##### ImageService.ImageDiskUsage

ImageDiskUsage returns the number of bytes used by content and layer stores
called from disk_usage.go

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L203)  

```go
func (i *ImageService) ImageDiskUsage(ctx context.Context) (int64, error)
```

##### ImageService.ImageHistory

ImageHistory returns a slice of ImageHistory structures for the specified image
name by walking the image lineage.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_history.go#L18)  

```go
func (i *ImageService) ImageHistory(ctx context.Context, name string, platform *ocispec.Platform) ([]*image.HistoryResponseItem, error)
```

##### ImageService.ImageInspect

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_inspect.go#L15)  

```go
func (i *ImageService) ImageInspect(ctx context.Context, refOrID string, opts backend.ImageInspectOpts) (*imagetypes.InspectResponse, error)
```

##### ImageService.Images

Images returns a filtered list of images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_list.go#L37)  

```go
func (i *ImageService) Images(ctx context.Context, opts imagetypes.ListOptions) ([]*imagetypes.Summary, error)
```

##### ImageService.ImagesPrune

ImagesPrune removes unused images

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_prune.go#L35)  

```go
func (i *ImageService) ImagesPrune(ctx context.Context, pruneFilters filters.Args) (*imagetypes.PruneReport, error)
```

##### ImageService.ImportImage

ImportImage imports an image, getting the archived layer data from layerReader.
Uncompressed layer archive is passed to the layerStore and handled by the
underlying graph driver.
Image is tagged with the given reference.
If the platform is nil, the default host platform is used.
Message is used as the image's history comment.
Image configuration is derived from the dockerfile instructions in changes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_import.go#L29)  

```go
func (i *ImageService) ImportImage(ctx context.Context, newRef reference.Named, platform *ocispec.Platform, msg string, layerReader io.Reader, changes []string) (image.ID, error)
```

##### ImageService.LayerStoreStatus

LayerStoreStatus returns the status for each layer store
called from info.go

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L159)  

```go
func (i *ImageService) LayerStoreStatus() [][2]string
```

##### ImageService.LoadImage

LoadImage uploads a set of images into the repository. This is the
complement of ExportImage.  The input stream is an uncompressed tar
ball containing images and metadata.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_exporter.go#L24)  

```go
func (i *ImageService) LoadImage(ctx context.Context, inTar io.ReadCloser, platform *ocispec.Platform, outStream io.Writer, quiet bool) error
```

##### ImageService.LogImageEvent

LogImageEvent generates an event related to an image with only the default attributes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_events.go#L11)  

```go
func (i *ImageService) LogImageEvent(ctx context.Context, imageID, refName string, action events.Action)
```

##### ImageService.MakeImageCache

MakeImageCache creates a stateful image cache.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/cache.go#L90)  

```go
func (i *ImageService) MakeImageCache(ctx context.Context, sourceRefs []string) (builder.ImageCache, error)
```

##### ImageService.PullImage

PullImage initiates a pull operation. image is the repository name to pull, and
tag may be either empty, or indicate a specific tag to pull.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_pull.go#L26)  

```go
func (i *ImageService) PullImage(ctx context.Context, ref reference.Named, platform *ocispec.Platform, metaHeaders map[string][]string, authConfig *registry.AuthConfig, outStream io.Writer) error
```

##### ImageService.PushImage

PushImage initiates a push operation on the repository named localName.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_push.go#L20)  

```go
func (i *ImageService) PushImage(ctx context.Context, ref reference.Named, platform *ocispec.Platform, metaHeaders map[string][]string, authConfig *registry.AuthConfig, outStream io.Writer) error
```

##### ImageService.ReleaseLayer

ReleaseLayer releases a layer allowing it to be removed
called from delete.go Daemon.cleanupContainer().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L186)  

```go
func (i *ImageService) ReleaseLayer(rwlayer container.RWLayer) error
```

##### ImageService.SquashImage

SquashImage creates a new image with the diff of the specified image and the specified parent.
This new image contains only the layers from it's parent + 1 extra layer which contains the diff of all the layers in between.
The existing image(s) is not destroyed.
If no parent is specified, a new image with the diff of all the specified image's layers merged into a new layer that has no parents.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_squash.go#L17)  

```go
func (i *ImageService) SquashImage(id, parent string) (string, error)
```

##### ImageService.StorageDriver

StorageDriver returns the name of the storage driver used by the ImageService.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L180)  

```go
func (i *ImageService) StorageDriver() string
```

##### ImageService.TagImage

TagImage adds the given reference to the image ID provided.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/image_tag.go#L12)  

```go
func (i *ImageService) TagImage(ctx context.Context, imageID image.ID, newTag reference.Named) error
```

##### ImageService.UpdateConfig

UpdateConfig values

called from reload.go

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L245)  

```go
func (i *ImageService) UpdateConfig(maxDownloads, maxUploads int)
```

---

### ImageServiceConfig

ImageServiceConfig is the configuration used to create a new ImageService

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/images/service.go#L34)  

```go
type ImageServiceConfig struct {
	ContainerStore            containerStore
	DistributionMetadataStore metadata.Store
	EventsService             *daemonevents.Events
	ImageStore                image.Store
	LayerStore                layer.Store
	MaxConcurrentDownloads    int
	MaxConcurrentUploads      int
	MaxDownloadAttempts       int
	ReferenceStore            dockerreference.Store
	RegistryService           distribution.RegistryResolver
	ContentStore              content.Store
	Leases                    leases.Manager
	ContentNamespace          string
}
```

---


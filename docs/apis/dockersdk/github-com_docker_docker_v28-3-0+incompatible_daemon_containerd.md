# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/containerd

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:59 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### ImageManifest

ImageManifest implements the containerd.Image interface, but all operations
act on the specific manifest instead of the index as opposed to the struct
returned by containerd.NewImageWithPlatform.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L99)  

```go
type ImageManifest struct {
	containerd.Image

	// Parent of the manifest (index/manifest list)
	RealTarget ocispec.Descriptor
	// contains filtered or unexported fields
}
```

#### Methods

##### ImageManifest.CheckContentAvailable

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L190)  

```go
func (im *ImageManifest) CheckContentAvailable(ctx context.Context) (bool, error)
```

##### ImageManifest.ImagePlatform

ImagePlatform returns the platform of the image manifest.
If the manifest list doesn't have a platform filled, it will be read from the config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L222)  

```go
func (im *ImageManifest) ImagePlatform(ctx context.Context) (ocispec.Platform, error)
```

##### ImageManifest.IsAttestation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L129)  

```go
func (im *ImageManifest) IsAttestation() bool
```

##### ImageManifest.IsPseudoImage

IsPseudoImage returns false when any of the below is true:
- The manifest has no layers
- None of its layers is a known image layer.
- The manifest has unknown/unknown platform.

Some manifests use the image media type for compatibility, even if they are not a real image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L146)  

```go
func (im *ImageManifest) IsPseudoImage(ctx context.Context) (bool, error)
```

##### ImageManifest.Manifest

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L176)  

```go
func (im *ImageManifest) Manifest(ctx context.Context) (ocispec.Manifest, error)
```

##### ImageManifest.Metadata

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L123)  

```go
func (im *ImageManifest) Metadata() c8dimages.Image
```

##### ImageManifest.PresentContentSize

PresentContentSize returns the size of the image's content that is present in the content store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L245)  

```go
func (im *ImageManifest) PresentContentSize(ctx context.Context) (int64, error)
```

##### ImageManifest.ReadConfig

ReadConfig gets the image config and unmarshals it into the provided struct.
The provided struct should be a pointer to the config struct or its subset.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L235)  

```go
func (im *ImageManifest) ReadConfig(ctx context.Context, outConfig interface{}) error
```

##### ImageManifest.SnapshotUsage

SnapshotUsage returns the disk usage of the image's snapshots.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L256)  

```go
func (im *ImageManifest) SnapshotUsage(ctx context.Context, snapshotter snapshots.Snapshotter) (snapshots.Usage, error)
```

---

### ImageService

ImageService implements daemon.ImageService

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L30)  

```go
type ImageService struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewService

NewService creates a new ImageService.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L60)  

```go
func NewService(config ImageServiceConfig) *ImageService
```

#### Methods

##### ImageService.Changes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_changes.go#L14)  

```go
func (i *ImageService) Changes(ctx context.Context, ctr *container.Container) ([]archive.Change, error)
```

##### ImageService.Children

Children returns a slice of image IDs that are children of the `id` image

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_children.go#L29)  

```go
func (i *ImageService) Children(ctx context.Context, id image.ID) ([]image.ID, error)
```

##### ImageService.Cleanup

Cleanup resources before the process is shutdown.
called from daemon.go Daemon.Shutdown()

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L122)  

```go
func (i *ImageService) Cleanup() error
```

##### ImageService.CommitBuildStep

CommitBuildStep is used by the builder to create an image for each step in
the build.

This method is different from CreateImageFromContainer:

This is a temporary shim. Should be removed when builder stops using commit.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_commit.go#L320)  

```go
func (i *ImageService) CommitBuildStep(ctx context.Context, c backend.CommitConfig) (image.ID, error)
```

##### ImageService.CommitImage

CommitImage creates a new image from a commit config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_commit.go#L36)  

```go
func (i *ImageService) CommitImage(ctx context.Context, cc backend.CommitConfig) (image.ID, error)
```

##### ImageService.CountImages

CountImages returns the number of images stored by ImageService
called from info.go

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L95)  

```go
func (i *ImageService) CountImages(ctx context.Context) int
```

##### ImageService.CreateImage

CreateImage creates a new image by adding a config and ID to the image store.
This is similar to LoadImage() except that it receives JSON encoded bytes of
an image instead of a tar archive.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_builder.go#L427)  

```go
func (i *ImageService) CreateImage(ctx context.Context, config []byte, parent string, layerDigest digest.Digest) (builder.Image, error)
```

##### ImageService.CreateLayer

CreateLayer creates a new layer for a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_snapshot.go#L25)  

```go
func (i *ImageService) CreateLayer(ctr *container.Container, initFunc layer.MountInit) (container.RWLayer, error)
```

##### ImageService.CreateLayerFromImage

CreateLayerFromImage creates a new layer from an image

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_snapshot.go#L39)  

```go
func (i *ImageService) CreateLayerFromImage(img *image.Image, layerName string, rwLayerOpts *layer.CreateRWLayerOpts) (container.RWLayer, error)
```

##### ImageService.DistributionServices

DistributionServices return services controlling daemon image storage.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L89)  

```go
func (i *ImageService) DistributionServices() dimages.DistributionServices
```

##### ImageService.ExportImage

ExportImage exports a list of images to the given output stream. The
exported images are archived into a tar when written to the output
stream. All images with the given tag and all versions containing
the same tag are exported. names is the set of tags to export, and
outStream is the writer which the images are written to.

TODO(thaJeztah): produce JSON stream progress response and image events; see https://github.com/moby/moby/issues/43910

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_exporter.go#L34)  

```go
func (i *ImageService) ExportImage(ctx context.Context, names []string, platform *ocispec.Platform, outStream io.Writer) error
```

##### ImageService.GetContainerLayerSize

GetContainerLayerSize returns the real size & virtual size of the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L187)  

```go
func (i *ImageService) GetContainerLayerSize(ctx context.Context, containerID string) (int64, int64, error)
```

##### ImageService.GetImage

GetImage returns an image corresponding to the image referred to by refOrID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image.go#L44)  

```go
func (i *ImageService) GetImage(ctx context.Context, refOrID string, options backend.GetImageOpts) (*image.Image, error)
```

##### ImageService.GetImageAndReleasableLayer

GetImageAndReleasableLayer returns an image and releaseable layer for a
reference or ID. Every call to GetImageAndReleasableLayer MUST call
releasableLayer.Release() to prevent leaking of layers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_builder.go#L67)  

```go
func (i *ImageService) GetImageAndReleasableLayer(ctx context.Context, refOrID string, opts backend.GetImageAndLayerOptions) (builder.Image, builder.ROLayer, error)
```

##### ImageService.GetLayerByID

GetLayerByID returns a layer by ID
called from daemon.go Daemon.restore().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_snapshot.go#L173)  

```go
func (i *ImageService) GetLayerByID(cid string) (container.RWLayer, error)
```

##### ImageService.GetLayerFolders

GetLayerFolders returns the layer folders from an image RootFS.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service_unix.go#L13)  

```go
func (i *ImageService) GetLayerFolders(img *image.Image, rwLayer container.RWLayer, containerID string) ([]string, error)
```

##### ImageService.GetLayerMountID

GetLayerMountID returns the mount ID for a layer
called from daemon.go Daemon.Shutdown(), and Daemon.Cleanup() (cleanup is actually containerCleanup)
TODO: needs to be refactored to Unmount (see callers), or removed and replaced with GetLayerByID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L116)  

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
reference is removed. However, if there exists any containers which
were created using the same image reference then the repository reference
cannot be removed unless either there are other repository references to the
same image or force is true. Following removal of the repository reference,
the referenced image itself is attempted to be deleted as described below
but quietly, meaning any image delete conflicts will cause the image to not
be deleted and the conflict will not be reported.

There may be conflicts preventing deletion of an image and these conflicts
are divided into two categories grouped by their severity:

Hard Conflict:

Soft Conflict:

The image cannot be removed if there are any hard conflicts and can be
removed if there are soft conflicts only if force is true.

If prune is true, ancestor images are attempted to be deleted quietly,
meaning any delete conflicts will cause the image to not be deleted and the
conflict will not be reported.

TODO(thaJeztah): image delete should send prometheus counters; see https://github.com/moby/moby/issues/45268

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_delete.go#L60)  

```go
func (i *ImageService) ImageDelete(ctx context.Context, imageRef string, options imagetypes.RemoveOptions) (response []imagetypes.DeleteResponse, retErr error)
```

##### ImageService.ImageDiskUsage

ImageDiskUsage returns the number of bytes used by content and layer stores
called from disk_usage.go

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L134)  

```go
func (i *ImageService) ImageDiskUsage(ctx context.Context) (int64, error)
```

##### ImageService.ImageHistory

ImageHistory returns a slice of HistoryResponseItem structures for the
specified image name by walking the image lineage.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_history.go#L21)  

```go
func (i *ImageService) ImageHistory(ctx context.Context, name string, platform *ocispec.Platform) ([]*imagetype.HistoryResponseItem, error)
```

##### ImageService.ImageInspect

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_inspect.go#L25)  

```go
func (i *ImageService) ImageInspect(ctx context.Context, refOrID string, opts backend.ImageInspectOpts) (*imagetypes.InspectResponse, error)
```

##### ImageService.Images

Images returns a filtered list of images.

TODO(thaJeztah): verify behavior of `RepoDigests` and `RepoTags` for images without (untagged) or multiple tags; see https://github.com/moby/moby/issues/43861
TODO(thaJeztah): verify "Size" vs "VirtualSize" in images; see https://github.com/moby/moby/issues/43862

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_list.go#L65)  

```go
func (i *ImageService) Images(ctx context.Context, opts imagetypes.ListOptions) ([]*imagetypes.Summary, error)
```

##### ImageService.ImagesPrune

ImagesPrune removes unused images

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_prune.go#L36)  

```go
func (i *ImageService) ImagesPrune(ctx context.Context, fltrs filters.Args) (*image.PruneReport, error)
```

##### ImageService.ImportImage

ImportImage imports an image, getting the archived layer data from layerReader.
Layer archive is imported as-is if the compression is gzip or zstd.
Uncompressed, xz and bzip2 archives are recompressed into gzip.
The image is tagged with the given reference.
If the platform is nil, the default host platform is used.
The message is used as the history comment.
Image configuration is derived from the dockerfile instructions in changes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_import.go#L40)  

```go
func (i *ImageService) ImportImage(ctx context.Context, ref reference.Named, platform *ocispec.Platform, msg string, layerReader io.Reader, changes []string) (image.ID, error)
```

##### ImageService.LayerStoreStatus

LayerStoreStatus returns the status for each layer store
called from info.go

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L106)  

```go
func (i *ImageService) LayerStoreStatus() [][2]string
```

##### ImageService.LoadImage

LoadImage uploads a set of images into the repository. This is the
complement of ExportImage.  The input stream is an uncompressed tar
ball containing images and metadata.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_exporter.go#L232)  

```go
func (i *ImageService) LoadImage(ctx context.Context, inTar io.ReadCloser, platform *ocispec.Platform, outStream io.Writer, quiet bool) error
```

##### ImageService.LogImageEvent

LogImageEvent generates an event related to an image with only the default attributes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_events.go#L12)  

```go
func (i *ImageService) LogImageEvent(ctx context.Context, imageID, refName string, action events.Action)
```

##### ImageService.MakeImageCache

MakeImageCache creates a stateful image cache.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/cache.go#L25)  

```go
func (i *ImageService) MakeImageCache(ctx context.Context, sourceRefs []string) (builder.ImageCache, error)
```

##### ImageService.NewImageManifest

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_manifest.go#L108)  

```go
func (i *ImageService) NewImageManifest(ctx context.Context, img c8dimages.Image, manifestDesc ocispec.Descriptor) (*ImageManifest, error)
```

##### ImageService.PullImage

PullImage initiates a pull operation. baseRef is the image to pull.
If reference is not tagged, all tags are pulled.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_pull.go#L33)  

```go
func (i *ImageService) PullImage(ctx context.Context, baseRef reference.Named, platform *ocispec.Platform, metaHeaders map[string][]string, authConfig *registrytypes.AuthConfig, outStream io.Writer) (retErr error)
```

##### ImageService.PushImage

PushImage initiates a push operation of the image pointed to by sourceRef.
If reference is untagged, all tags from the reference repository are pushed.
Image manifest (or index) is pushed as is, which will probably fail if you
don't have all content referenced by the index.
Cross-repo mounts will be attempted for non-existing blobs.

It will also add distribution source labels to the pushed content
pointing to the new target repository. This will allow subsequent pushes
to perform cross-repo mounts of the shared content when pushing to a different
repository on the same registry.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_push.go#L43)  

```go
func (i *ImageService) PushImage(ctx context.Context, sourceRef reference.Named, platform *ocispec.Platform, metaHeaders map[string][]string, authConfig *registry.AuthConfig, outStream io.Writer) (retErr error)
```

##### ImageService.ReleaseLayer

ReleaseLayer releases a layer allowing it to be removed
called from delete.go Daemon.cleanupContainer(), and Daemon.containerExport()

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_snapshot.go#L244)  

```go
func (i *ImageService) ReleaseLayer(rwlayer container.RWLayer) error
```

##### ImageService.ResolveImage

ResolveImage looks up an image by reference or identifier in the image store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image.go#L128)  

```go
func (i *ImageService) ResolveImage(ctx context.Context, refOrID string) (c8dimages.Image, error)
```

##### ImageService.SquashImage

SquashImage creates a new image with the diff of the specified image and
the specified parent. This new image contains only the layers from its
parent + 1 extra layer which contains the diff of all the layers in between.
The existing image(s) is not destroyed. If no parent is specified, a new
image with the diff of all the specified image's layers merged into a new
layer that has no parents.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_squash.go#L15)  

```go
func (i *ImageService) SquashImage(id, parent string) (string, error)
```

##### ImageService.StorageDriver

StorageDriver returns the name of the default storage-driver (snapshotter)
used by the ImageService.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L128)  

```go
func (i *ImageService) StorageDriver() string
```

##### ImageService.TagImage

TagImage creates an image named as newTag and targeting the given descriptor id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/image_tag.go#L18)  

```go
func (i *ImageService) TagImage(ctx context.Context, imageID image.ID, newTag reference.Named) error
```

##### ImageService.UpdateConfig

UpdateConfig values

called from reload.go

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L182)  

```go
func (i *ImageService) UpdateConfig(maxDownloads, maxUploads int)
```

---

### ImageServiceConfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/containerd/service.go#L48)  

```go
type ImageServiceConfig struct {
	Client          *containerd.Client
	Containers      container.Store
	Snapshotter     string
	RegistryHosts   docker.RegistryHosts
	Registry        distribution.RegistryResolver
	EventsService   *daemonevents.Events
	RefCountMounter snapshotter.Mounter
	IDMapping       user.IdentityMapping
}
```

---


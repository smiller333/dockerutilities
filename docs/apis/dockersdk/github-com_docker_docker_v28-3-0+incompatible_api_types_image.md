# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/image

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:01 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### AttestationProperties

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/manifest.go#L96)  

```go
type AttestationProperties struct {
	// For is the digest of the image manifest that this attestation is for.
	For digest.Digest `json:"For"`
}
```

---

### CreateOptions

CreateOptions holds information to create images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L26)  

```go
type CreateOptions struct {
	RegistryAuth string // RegistryAuth is the base64 encoded credentials for the registry.
	Platform     string // Platform is the target platform of the image if it needs to be pulled from the registry.
}
```

---

### DeleteResponse

DeleteResponse delete response
swagger:model DeleteResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/delete_response.go#L8)  

```go
type DeleteResponse struct {

	// The image ID of an image that was deleted
	Deleted string `json:"Deleted,omitempty"`

	// The image ID of an image that was untagged
	Untagged string `json:"Untagged,omitempty"`
}
```

---

### DiskUsage

DiskUsage contains disk usage for images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/disk_usage.go#L4)  

```go
type DiskUsage struct {
	TotalSize   int64
	Reclaimable int64
	Items       []*Summary
}
```

---

### HistoryOptions

HistoryOptions holds parameters to get image history.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L94)  

```go
type HistoryOptions struct {
	// Platform from the manifest list to use for history.
	Platform *ocispec.Platform
}
```

---

### HistoryResponseItem

HistoryResponseItem individual image layer information in response to ImageHistory operation
swagger:model HistoryResponseItem

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/image_history.go#L11)  

```go
type HistoryResponseItem struct {

	// comment
	// Required: true
	Comment string `json:"Comment"`

	// created
	// Required: true
	Created int64 `json:"Created"`

	// created by
	// Required: true
	CreatedBy string `json:"CreatedBy"`

	// Id
	// Required: true
	ID string `json:"Id"`

	// size
	// Required: true
	Size int64 `json:"Size"`

	// tags
	// Required: true
	Tags []string `json:"Tags"`
}
```

---

### ImageProperties

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/manifest.go#L70)  

```go
type ImageProperties struct {
	// Platform is the OCI platform object describing the platform of the image.
	//
	// Required: true
	Platform ocispec.Platform `json:"Platform"`

	Size struct {
		// Unpacked is the size (in bytes) of the locally unpacked
		// (uncompressed) image content that's directly usable by the containers
		// running this image.
		// It's independent of the distributable content - e.g.
		// the image might still have an unpacked data that's still used by
		// some container even when the distributable/compressed content is
		// already gone.
		//
		// Required: true
		Unpacked int64 `json:"Unpacked"`
	}

	// Containers is an array containing the IDs of the containers that are
	// using this image.
	//
	// Required: true
	Containers []string `json:"Containers"`
}
```

---

### ImportOptions

ImportOptions holds information to import images from the client host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L18)  

```go
type ImportOptions struct {
	Tag      string   // Tag is the name to tag this image with. This attribute is deprecated.
	Message  string   // Message is the message to tag the image with
	Changes  []string // Changes are the raw changes to apply to this image
	Platform string   // Platform is the target platform of the image
}
```

---

### ImportSource

ImportSource holds source information for ImageImport

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L12)  

```go
type ImportSource struct {
	Source     io.Reader // Source is the data to send to the server to create this image from. You must set SourceName to "-" to leverage this.
	SourceName string    // SourceName is the name of the image to pull. Set to "-" to leverage the Source attribute.
}
```

---

### InspectOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L109)  

```go
type InspectOptions struct {
	// Manifests returns the image manifests.
	Manifests bool

	// Platform selects the specific platform of a multi-platform image to inspect.
	//
	// This option is only available for API version 1.49 and up.
	Platform *ocispec.Platform
}
```

---

### InspectResponse

InspectResponse contains response of Engine API:
GET "/images/{name:.*}/json"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/image_inspect.go#L18)  

```go
type InspectResponse struct {
	// ID is the content-addressable ID of an image.
	//
	// This identifier is a content-addressable digest calculated from the
	// image's configuration (which includes the digests of layers used by
	// the image).
	//
	// Note that this digest differs from the `RepoDigests` below, which
	// holds digests of image manifests that reference the image.
	ID string `json:"Id"`

	// RepoTags is a list of image names/tags in the local image cache that
	// reference this image.
	//
	// Multiple image tags can refer to the same image, and this list may be
	// empty if no tags reference the image, in which case the image is
	// "untagged", in which case it can still be referenced by its ID.
	RepoTags []string

	// RepoDigests is a list of content-addressable digests of locally available
	// image manifests that the image is referenced from. Multiple manifests can
	// refer to the same image.
	//
	// These digests are usually only available if the image was either pulled
	// from a registry, or if the image was pushed to a registry, which is when
	// the manifest is generated and its digest calculated.
	RepoDigests []string

	// Parent is the ID of the parent image.
	//
	// Depending on how the image was created, this field may be empty and
	// is only set for images that were built/created locally. This field
	// is empty if the image was pulled from an image registry.
	Parent string

	// Comment is an optional message that can be set when committing or
	// importing the image.
	Comment string

	// Created is the date and time at which the image was created, formatted in
	// RFC 3339 nano-seconds (time.RFC3339Nano).
	//
	// This information is only available if present in the image,
	// and omitted otherwise.
	Created string `json:",omitempty"`

	// Container is the ID of the container that was used to create the image.
	//
	// Depending on how the image was created, this field may be empty.
	//
	// Deprecated: this field is omitted in API v1.45, but kept for backward compatibility.
	Container string `json:",omitempty"`

	// ContainerConfig is an optional field containing the configuration of the
	// container that was last committed when creating the image.
	//
	// Previous versions of Docker builder used this field to store build cache,
	// and it is not in active use anymore.
	//
	// Deprecated: this field is omitted in API v1.45, but kept for backward compatibility.
	ContainerConfig *container.Config `json:",omitempty"`

	// DockerVersion is the version of Docker that was used to build the image.
	//
	// Depending on how the image was created, this field may be empty.
	DockerVersion string

	// Author is the name of the author that was specified when committing the
	// image, or as specified through MAINTAINER (deprecated) in the Dockerfile.
	Author string
	Config *dockerspec.DockerOCIImageConfig

	// Architecture is the hardware CPU architecture that the image runs on.
	Architecture string

	// Variant is the CPU architecture variant (presently ARM-only).
	Variant string `json:",omitempty"`

	// OS is the Operating System the image is built to run on.
	Os string

	// OsVersion is the version of the Operating System the image is built to
	// run on (especially for Windows).
	OsVersion string `json:",omitempty"`

	// Size is the total size of the image including all layers it is composed of.
	Size int64

	// VirtualSize is the total size of the image including all layers it is
	// composed of.
	//
	// Deprecated: this field is omitted in API v1.44, but kept for backward compatibility. Use Size instead.
	VirtualSize int64 `json:"VirtualSize,omitempty"`

	// GraphDriver holds information about the storage driver used to store the
	// container's and image's filesystem.
	GraphDriver storage.DriverData

	// RootFS contains information about the image's RootFS, including the
	// layer IDs.
	RootFS RootFS

	// Metadata of the image in the local cache.
	//
	// This information is local to the daemon, and not part of the image itself.
	Metadata Metadata

	// Descriptor is the OCI descriptor of the image target.
	// It's only set if the daemon provides a multi-platform image store.
	//
	// WARNING: This is experimental and may change at any time without any backward
	// compatibility.
	Descriptor *ocispec.Descriptor `json:"Descriptor,omitempty"`

	// Manifests is a list of image manifests available in this image. It
	// provides a more detailed view of the platform-specific image manifests or
	// other image-attached data like build attestations.
	//
	// Only available if the daemon provides a multi-platform image store, the client
	// requests manifests AND does not request a specific platform.
	//
	// WARNING: This is experimental and may change at any time without any backward
	// compatibility.
	Manifests []ManifestSummary `json:"Manifests,omitempty"`
}
```

---

### ListOptions

ListOptions holds parameters to list images with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L66)  

```go
type ListOptions struct {
	// All controls whether all images in the graph are filtered, or just
	// the heads.
	All bool

	// Filters is a JSON-encoded set of filter arguments.
	Filters filters.Args

	// SharedSize indicates whether the shared size of images should be computed.
	SharedSize bool

	// ContainerCount indicates whether container count should be computed.
	//
	// Deprecated: This field has been unused and is no longer required and will be removed in a future version.
	ContainerCount bool

	// Manifests indicates whether the image manifests should be returned.
	Manifests bool
}
```

---

### LoadOptions

LoadOptions holds parameters to load images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L100)  

```go
type LoadOptions struct {
	// Quiet suppresses progress output
	Quiet bool

	// Platforms selects the platforms to load if the image is a
	// multi-platform image and has multiple variants.
	Platforms []ocispec.Platform
}
```

---

### LoadResponse

LoadResponse returns information to the client about a load process.

TODO(thaJeztah): remove this type, and just use an io.ReadCloser

This type was added in https://github.com/moby/moby/pull/18878, related
to https://github.com/moby/moby/issues/19177;

Make docker load to output json when the response content type is json
Swarm hijacks the response from docker load and returns JSON rather
than plain text like the Engine does. This makes the API library to return
information to figure that out.

However the "load" endpoint unconditionally returns JSON;
https://github.com/moby/moby/blob/7b9d2ef6e5518a3d3f3cc418459f8df786cfbbd1/api/server/router/image/image_routes.go#L248-L255

PR https://github.com/moby/moby/pull/21959 made the response-type depend
on whether "quiet" was set, but this logic got changed in a follow-up
https://github.com/moby/moby/pull/25557, which made the JSON response-type
unconditionally, but the output produced depend on whether"quiet" was set.

We should deprecated the "quiet" option, as it's really a client
responsibility.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/image.go#L43)  

```go
type LoadResponse struct {
	// Body must be closed to avoid a resource leak
	Body io.ReadCloser
	JSON bool
}
```

---

### ManifestKind

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/manifest.go#L8)  

```go
type ManifestKind string
```

---

### ManifestSummary

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/manifest.go#L16)  

```go
type ManifestSummary struct {
	// ID is the content-addressable ID of an image and is the same as the
	// digest of the image manifest.
	//
	// Required: true
	ID string `json:"ID"`

	// Descriptor is the OCI descriptor of the image.
	//
	// Required: true
	Descriptor ocispec.Descriptor `json:"Descriptor"`

	// Indicates whether all the child content (image config, layers) is
	// fully available locally
	//
	// Required: true
	Available bool `json:"Available"`

	// Size is the size information of the content related to this manifest.
	// Note: These sizes only take the locally available content into account.
	//
	// Required: true
	Size struct {
		// Content is the size (in bytes) of all the locally present
		// content in the content store (e.g. image config, layers)
		// referenced by this manifest and its children.
		// This only includes blobs in the content store.
		Content int64 `json:"Content"`

		// Total is the total size (in bytes) of all the locally present
		// data (both distributable and non-distributable) that's related to
		// this manifest and its children.
		// This equal to the sum of [Content] size AND all the sizes in the
		// [Size] struct present in the Kind-specific data struct.
		// For example, for an image kind (Kind == ManifestKindImage),
		// this would include the size of the image content and unpacked
		// image snapshots ([Size.Content] + [ImageData.Size.Unpacked]).
		Total int64 `json:"Total"`
	} `json:"Size"`

	// Kind is the kind of the image manifest.
	//
	// Required: true
	Kind ManifestKind `json:"Kind"`

	// Present only if Kind == ManifestKindImage.
	ImageData *ImageProperties `json:"ImageData,omitempty"`

	// Present only if Kind == ManifestKindAttestation.
	AttestationData *AttestationProperties `json:"AttestationData,omitempty"`
}
```

---

### Metadata

Metadata contains engine-local data about the image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/image.go#L9)  

```go
type Metadata struct {
	// LastTagTime is the date and time at which the image was last tagged.
	LastTagTime time.Time `json:",omitempty"`
}
```

---

### PruneReport

PruneReport contains the response for Engine API:
POST "/images/prune"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/image.go#L16)  

```go
type PruneReport struct {
	ImagesDeleted  []DeleteResponse
	SpaceReclaimed uint64
}
```

---

### PullOptions

PullOptions holds information to pull images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L32)  

```go
type PullOptions struct {
	All          bool
	RegistryAuth string // RegistryAuth is the base64 encoded credentials for the registry

	// PrivilegeFunc is a function that clients can supply to retry operations
	// after getting an authorization error. This function returns the registry
	// authentication header value in base64 encoded format, or an error if the
	// privilege request fails.
	//
	// For details, refer to [github.com/docker/docker/api/types/registry.RequestAuthConfig].
	PrivilegeFunc func(context.Context) (string, error)
	Platform      string
}
```

---

### PushOptions

PushOptions holds information to push images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L47)  

```go
type PushOptions struct {
	All          bool
	RegistryAuth string // RegistryAuth is the base64 encoded credentials for the registry

	// PrivilegeFunc is a function that clients can supply to retry operations
	// after getting an authorization error. This function returns the registry
	// authentication header value in base64 encoded format, or an error if the
	// privilege request fails.
	//
	// For details, refer to [github.com/docker/docker/api/types/registry.RequestAuthConfig].
	PrivilegeFunc func(context.Context) (string, error)

	// Platform is an optional field that selects a specific platform to push
	// when the image is a multi-platform image.
	// Using this will only push a single platform-specific manifest.
	Platform *ocispec.Platform `json:",omitempty"`
}
```

---

### RemoveOptions

RemoveOptions holds parameters to remove images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L87)  

```go
type RemoveOptions struct {
	Platforms     []ocispec.Platform
	Force         bool
	PruneChildren bool
}
```

---

### RootFS

RootFS returns Image's RootFS description including the layer IDs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/image_inspect.go#L11)  

```go
type RootFS struct {
	Type   string   `json:",omitempty"`
	Layers []string `json:",omitempty"`
}
```

---

### SaveOptions

SaveOptions holds parameters to save images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/opts.go#L120)  

```go
type SaveOptions struct {
	// Platforms selects the platforms to save if the image is a
	// multi-platform image and has multiple variants.
	Platforms []ocispec.Platform
}
```

---

### Summary

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/image/summary.go#L5)  

```go
type Summary struct {

	// Number of containers using this image. Includes both stopped and running
	// containers.
	//
	// This size is not calculated by default, and depends on which API endpoint
	// is used. `-1` indicates that the value has not been set / calculated.
	//
	// Required: true
	Containers int64 `json:"Containers"`

	// Date and time at which the image was created as a Unix timestamp
	// (number of seconds since EPOCH).
	//
	// Required: true
	Created int64 `json:"Created"`

	// ID is the content-addressable ID of an image.
	//
	// This identifier is a content-addressable digest calculated from the
	// image's configuration (which includes the digests of layers used by
	// the image).
	//
	// Note that this digest differs from the `RepoDigests` below, which
	// holds digests of image manifests that reference the image.
	//
	// Required: true
	ID string `json:"Id"`

	// User-defined key/value metadata.
	// Required: true
	Labels map[string]string `json:"Labels"`

	// ID of the parent image.
	//
	// Depending on how the image was created, this field may be empty and
	// is only set for images that were built/created locally. This field
	// is empty if the image was pulled from an image registry.
	//
	// Required: true
	ParentID string `json:"ParentId"`

	// Descriptor is the OCI descriptor of the image target.
	// It's only set if the daemon provides a multi-platform image store.
	//
	// WARNING: This is experimental and may change at any time without any backward
	// compatibility.
	Descriptor *ocispec.Descriptor `json:"Descriptor,omitempty"`

	// Manifests is a list of image manifests available in this image.  It
	// provides a more detailed view of the platform-specific image manifests or
	// other image-attached data like build attestations.
	//
	// WARNING: This is experimental and may change at any time without any backward
	// compatibility.
	Manifests []ManifestSummary `json:"Manifests,omitempty"`

	// List of content-addressable digests of locally available image manifests
	// that the image is referenced from. Multiple manifests can refer to the
	// same image.
	//
	// These digests are usually only available if the image was either pulled
	// from a registry, or if the image was pushed to a registry, which is when
	// the manifest is generated and its digest calculated.
	//
	// Required: true
	RepoDigests []string `json:"RepoDigests"`

	// List of image names/tags in the local image cache that reference this
	// image.
	//
	// Multiple image tags can refer to the same image, and this list may be
	// empty if no tags reference the image, in which case the image is
	// "untagged", in which case it can still be referenced by its ID.
	//
	// Required: true
	RepoTags []string `json:"RepoTags"`

	// Total size of image layers that are shared between this image and other
	// images.
	//
	// This size is not calculated by default. `-1` indicates that the value
	// has not been set / calculated.
	//
	// Required: true
	SharedSize int64 `json:"SharedSize"`

	// Total size of the image including all layers it is composed of.
	//
	// Required: true
	Size int64 `json:"Size"`

	// Total size of the image including all layers it is composed of.
	//
	// Deprecated: this field is omitted in API v1.44, but kept for backward compatibility. Use Size instead.
	VirtualSize int64 `json:"VirtualSize,omitempty"`
}
```

---


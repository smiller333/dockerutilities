# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/build

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:47 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### BuilderVersion

BuilderVersion sets the version of underlying builder to use

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/build/build.go#L11)  

```go
type BuilderVersion string
```

---

### CacheDiskUsage

CacheDiskUsage contains disk usage for the build cache.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/build/disk_usage.go#L4)  

```go
type CacheDiskUsage struct {
	TotalSize   int64
	Reclaimable int64
	Items       []*CacheRecord
}
```

---

### CachePruneOptions

CachePruneOptions hold parameters to prune the build cache.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/build/cache.go#L37)  

```go
type CachePruneOptions struct {
	All           bool
	ReservedSpace int64
	MaxUsedSpace  int64
	MinFreeSpace  int64
	Filters       filters.Args

	KeepStorage int64 // Deprecated: deprecated in API 1.48.
}
```

---

### CachePruneReport

CachePruneReport contains the response for Engine API:
POST "/build/prune"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/build/cache.go#L49)  

```go
type CachePruneReport struct {
	CachesDeleted  []string
	SpaceReclaimed uint64
}
```

---

### CacheRecord

CacheRecord contains information about a build cache record.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/build/cache.go#L10)  

```go
type CacheRecord struct {
	// ID is the unique ID of the build cache record.
	ID string
	// Parent is the ID of the parent build cache record.
	//
	// Deprecated: deprecated in API v1.42 and up, as it was deprecated in BuildKit; use Parents instead.
	Parent string `json:"Parent,omitempty"`
	// Parents is the list of parent build cache record IDs.
	Parents []string `json:" Parents,omitempty"`
	// Type is the cache record type.
	Type string
	// Description is a description of the build-step that produced the build cache.
	Description string
	// InUse indicates if the build cache is in use.
	InUse bool
	// Shared indicates if the build cache is shared.
	Shared bool
	// Size is the amount of disk space used by the build cache (in bytes).
	Size int64
	// CreatedAt is the date and time at which the build cache was created.
	CreatedAt time.Time
	// LastUsedAt is the date and time at which the build cache was last used.
	LastUsedAt *time.Time
	UsageCount int
}
```

---

### ImageBuildOptions

ImageBuildOptions holds the information
necessary to build images.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/build/build.go#L27)  

```go
type ImageBuildOptions struct {
	Tags           []string
	SuppressOutput bool
	RemoteContext  string
	NoCache        bool
	Remove         bool
	ForceRemove    bool
	PullParent     bool
	Isolation      container.Isolation
	CPUSetCPUs     string
	CPUSetMems     string
	CPUShares      int64
	CPUQuota       int64
	CPUPeriod      int64
	Memory         int64
	MemorySwap     int64
	CgroupParent   string
	NetworkMode    string
	ShmSize        int64
	Dockerfile     string
	Ulimits        []*container.Ulimit
	// BuildArgs needs to be a *string instead of just a string so that
	// we can tell the difference between "" (empty string) and no value
	// at all (nil). See the parsing of buildArgs in
	// api/server/router/build/build_routes.go for even more info.
	BuildArgs   map[string]*string
	AuthConfigs map[string]registry.AuthConfig
	Context     io.Reader
	Labels      map[string]string
	// squash the resulting image's layers to the parent
	// preserves the original image and creates a new one from the parent with all
	// the changes applied to a single layer
	Squash bool
	// CacheFrom specifies images that are used for matching cache. Images
	// specified here do not need to have a valid parent chain to match cache.
	CacheFrom   []string
	SecurityOpt []string
	ExtraHosts  []string // List of extra hosts
	Target      string
	SessionID   string
	Platform    string
	// Version specifies the version of the underlying builder to use
	Version BuilderVersion
	// BuildID is an optional identifier that can be passed together with the
	// build request. The same identifier can be used to gracefully cancel the
	// build with the cancel request.
	BuildID string
	// Outputs defines configurations for exporting build results. Only supported
	// in BuildKit mode
	Outputs []ImageBuildOutput
}
```

---

### ImageBuildOutput

ImageBuildOutput defines configuration for exporting a build result

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/build/build.go#L80)  

```go
type ImageBuildOutput struct {
	Type  string
	Attrs map[string]string
}
```

---

### ImageBuildResponse

ImageBuildResponse holds information
returned by a server after building
an image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/build/build.go#L88)  

```go
type ImageBuildResponse struct {
	Body   io.ReadCloser
	OSType string
}
```

---

### Result

Result contains the image id of a successful build.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/build/build.go#L21)  

```go
type Result struct {
	ID string
}
```

---


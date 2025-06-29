# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/builder-next/worker

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:54 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### DefaultGCPolicy

DefaultGCPolicy returns a default builder GC policy

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/gc.go#L24)  

```go
func DefaultGCPolicy(p string, reservedSpace, maxUsedSpace, minFreeSpace int64) []client.PruneInfo
```

---

## Types

### ContainerdWorker

ContainerdWorker is a local worker instance with dedicated snapshotter, cache, and so on.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/containerdworker.go#L17)  

```go
type ContainerdWorker struct {
	*base.Worker
	// contains filtered or unexported fields
}
```

#### Functions

##### NewContainerdWorker

NewContainerdWorker instantiates a local worker.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/containerdworker.go#L23)  

```go
func NewContainerdWorker(ctx context.Context, wo base.WorkerOpt, callbacks exporter.BuildkitCallbacks, rt nethttp.RoundTripper) (*ContainerdWorker, error)
```

#### Methods

##### ContainerdWorker.Exporter

Exporter returns exporter by name

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/containerdworker.go#L42)  

```go
func (w *ContainerdWorker) Exporter(name string, sm *session.Manager) (bkexporter.Exporter, error)
```

---

### LayerAccess

LayerAccess provides access to a moby layer from a snapshot

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L67)  

```go
type LayerAccess interface {
	GetDiffIDs(ctx context.Context, key string) ([]layer.DiffID, error)
	EnsureLayer(ctx context.Context, key string) ([]layer.DiffID, error)
}
```

---

### Opt

Opt defines a structure for creating a worker.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L73)  

```go
type Opt struct {
	ID                string
	Labels            map[string]string
	GCPolicy          []client.PruneInfo
	Executor          executor.Executor
	Snapshotter       snapshot.Snapshotter
	ContentStore      *containerdsnapshot.Store
	CacheManager      cache.Manager
	LeaseManager      *leaseutil.Manager
	GarbageCollect    func(context.Context) (gc.Stats, error)
	ImageSource       *imageadapter.Source
	DownloadManager   *xfer.LayerDownloadManager
	V2MetadataService distmetadata.V2MetadataService
	Transport         nethttp.RoundTripper
	Exporter          exporter.Exporter
	Layers            LayerAccess
	Platforms         []ocispec.Platform
	CDIManager        *cdidevices.Manager
}
```

---

### Worker

Worker is a local worker instance with dedicated snapshotter, cache, and so on.
TODO: s/Worker/OpWorker/g ?

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L95)  

```go
type Worker struct {
	Opt
	SourceManager *source.Manager
}
```

#### Functions

##### NewWorker

NewWorker instantiates a local worker

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L105)  

```go
func NewWorker(opt Opt) (*Worker, error)
```

#### Methods

##### Worker.BuildkitVersion

BuildkitVersion returns BuildKit version

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L199)  

```go
func (w *Worker) BuildkitVersion() client.BuildkitVersion
```

##### Worker.CDIManager

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L501)  

```go
func (w *Worker) CDIManager() *cdidevices.Manager
```

##### Worker.CacheManager

CacheManager returns cache.Manager for accessing local storage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L497)  

```go
func (w *Worker) CacheManager() cache.Manager
```

##### Worker.Close

Close closes the worker and releases all resources

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L216)  

```go
func (w *Worker) Close() error
```

##### Worker.ContentStore

ContentStore returns the wrapped content store

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L221)  

```go
func (w *Worker) ContentStore() *containerdsnapshot.Store
```

##### Worker.DiskUsage

DiskUsage returns disk usage report

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L317)  

```go
func (w *Worker) DiskUsage(ctx context.Context, opt client.DiskUsageInfo) ([]*client.UsageInfo, error)
```

##### Worker.Executor

Executor returns executor.Executor for running processes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L492)  

```go
func (w *Worker) Executor() executor.Executor
```

##### Worker.Exporter

Exporter returns exporter by name

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L327)  

```go
func (w *Worker) Exporter(name string, sm *session.Manager) (exporter.Exporter, error)
```

##### Worker.FromRemote

FromRemote converts a remote snapshot reference to a local one

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L431)  

```go
func (w *Worker) FromRemote(ctx context.Context, remote *solver.Remote) (cache.ImmutableRef, error)
```

##### Worker.GCPolicy

GCPolicy returns automatic GC Policy

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L194)  

```go
func (w *Worker) GCPolicy() []client.PruneInfo
```

##### Worker.GarbageCollect

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L207)  

```go
func (w *Worker) GarbageCollect(ctx context.Context) error
```

##### Worker.GetRemotes

GetRemotes returns the remote snapshot references given a local reference

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L345)  

```go
func (w *Worker) GetRemotes(ctx context.Context, ref cache.ImmutableRef, createIfNeeded bool, _ cacheconfig.RefConfig, all bool, s session.Group) ([]*solver.Remote, error)
```

##### Worker.ID

ID returns worker ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L149)  

```go
func (w *Worker) ID() string
```

##### Worker.Labels

Labels returns map of all worker labels

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L154)  

```go
func (w *Worker) Labels() map[string]string
```

##### Worker.LeaseManager

LeaseManager returns the wrapped lease manager

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L226)  

```go
func (w *Worker) LeaseManager() *leaseutil.Manager
```

##### Worker.LoadRef

LoadRef loads a reference by ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L231)  

```go
func (w *Worker) LoadRef(ctx context.Context, id string, hidden bool) (cache.ImmutableRef, error)
```

##### Worker.Platforms

Platforms returns one or more platforms supported by the image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L159)  

```go
func (w *Worker) Platforms(noCache bool) []ocispec.Platform
```

##### Worker.Prune

Prune deletes reclaimable build cache

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L322)  

```go
func (w *Worker) Prune(ctx context.Context, ch chan client.UsageInfo, info ...client.PruneInfo) error
```

##### Worker.PruneCacheMounts

PruneCacheMounts removes the current cache snapshots for specified IDs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L385)  

```go
func (w *Worker) PruneCacheMounts(ctx context.Context, ids map[string]bool) error
```

##### Worker.ResolveImageConfig

ResolveImageConfig returns image config for an image

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L312)  

```go
func (w *Worker) ResolveImageConfig(ctx context.Context, ref string, opt sourceresolver.Opt, sm *session.Manager, g session.Group) (digest.Digest, []byte, error)
```

##### Worker.ResolveOp

ResolveOp converts a LLB vertex into a LLB operation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L289)  

```go
func (w *Worker) ResolveOp(v solver.Vertex, s frontend.FrontendLLBBridge, sm *session.Manager) (solver.Op, error)
```

##### Worker.ResolveSourceMetadata

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/worker/worker.go#L245)  

```go
func (w *Worker) ResolveSourceMetadata(ctx context.Context, op *pb.SourceOp, opt sourceresolver.Opt, sm *session.Manager, g session.Group) (*sourceresolver.MetaResponse, error)
```

---


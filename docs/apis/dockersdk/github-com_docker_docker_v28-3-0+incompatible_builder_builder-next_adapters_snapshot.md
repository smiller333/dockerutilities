# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/builder-next/adapters/snapshot

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:42 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewSnapshotter

NewSnapshotter creates a new snapshotter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/snapshot/snapshot.go#L63)  

```go
func NewSnapshotter(opt Opt, prevLM leases.Manager, ns string) (snapshot.Snapshotter, *leaseutil.Manager, error)
```

---

## Types

### Opt

Opt defines options for creating the snapshotter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/builder-next/adapters/snapshot/snapshot.go#L35)  

```go
type Opt struct {
	GraphDriver     graphdriver.Driver
	LayerStore      layer.Store
	Root            string
	IdentityMapping user.IdentityMapping
}
```

---


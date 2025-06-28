# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libcontainerd

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:19 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewClient

NewClient creates a new libcontainerd client from a containerd client

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/libcontainerd_linux.go#L12)  

```go
func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend) (libcontainerdtypes.Client, error)
```

---

### ReplaceContainer

ReplaceContainer creates a new container, replacing any existing container
with the same id if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/replace.go#L17)  

```go
func ReplaceContainer(ctx context.Context, client types.Client, id string, spec *specs.Spec, shim string, runtimeOptions interface{}, opts ...containerd.NewContainerOpts) (types.Container, error)
```

---

## Types

This section is empty.


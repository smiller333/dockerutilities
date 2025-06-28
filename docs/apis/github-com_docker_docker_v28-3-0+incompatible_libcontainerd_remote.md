# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libcontainerd/remote

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:26 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/remote/client.go#L44)

```go
const DockerContainerBundlePath = "com.docker/engine.bundle.path"
```

## Variables

This section is empty.

## Functions

### NewClient

NewClient creates a new libcontainerd client from a containerd client

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/remote/client.go#L73)  

```go
func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend) (libcontainerdtypes.Client, error)
```

---

### WithBundle

WithBundle creates the bundle for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/remote/client_linux.go#L54)  

```go
func WithBundle(bundleDir string, ociSpec *specs.Spec) containerd.NewContainerOpts
```

---

## Types

This section is empty.


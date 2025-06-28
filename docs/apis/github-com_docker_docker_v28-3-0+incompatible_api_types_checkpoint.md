# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/checkpoint

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:51 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### CreateOptions

CreateOptions holds parameters to create a checkpoint from a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/checkpoint/options.go#L4)  

```go
type CreateOptions struct {
	CheckpointID  string
	CheckpointDir string
	Exit          bool
}
```

---

### DeleteOptions

DeleteOptions holds parameters to delete a checkpoint from a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/checkpoint/options.go#L16)  

```go
type DeleteOptions struct {
	CheckpointID  string
	CheckpointDir string
}
```

---

### ListOptions

ListOptions holds parameters to list checkpoints for a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/checkpoint/options.go#L11)  

```go
type ListOptions struct {
	CheckpointDir string
}
```

---

### Summary

Summary represents the details of a checkpoint when listing endpoints.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/checkpoint/list.go#L4)  

```go
type Summary struct {
	// Name is the name of the checkpoint.
	Name string
}
```

---


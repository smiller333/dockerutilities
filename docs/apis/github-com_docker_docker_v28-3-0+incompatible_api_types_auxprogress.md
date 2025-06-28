# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/auxprogress

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:41 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### ContentMissing

ContentMissing is a note that is sent when push fails because the content is missing.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/auxprogress/push.go#L21)  

```go
type ContentMissing struct {
	ContentMissing bool `json:"contentMissing"` // Always true

	// Desc is the descriptor of the root object that was attempted to be pushed.
	Desc ocispec.Descriptor `json:"desc"`
}
```

---

### ManifestPushedInsteadOfIndex

ManifestPushedInsteadOfIndex is a note that is sent when a manifest is pushed
instead of an index.  It is sent when the pushed image is an multi-platform
index, but the whole index couldn't be pushed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/auxprogress/push.go#L10)  

```go
type ManifestPushedInsteadOfIndex struct {
	ManifestPushedInsteadOfIndex bool `json:"manifestPushedInsteadOfIndex"` // Always true

	// OriginalIndex is the descriptor of the original image index.
	OriginalIndex ocispec.Descriptor `json:"originalIndex"`

	// SelectedManifest is the descriptor of the manifest that was pushed instead.
	SelectedManifest ocispec.Descriptor `json:"selectedManifest"`
}
```

---


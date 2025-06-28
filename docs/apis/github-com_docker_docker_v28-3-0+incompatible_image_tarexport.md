# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/image/tarexport

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:06:05 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### NewTarExporter

NewTarExporter returns new Exporter for tar packages

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/tarexport/tarexport.go#L46)  

```go
func NewTarExporter(is image.Store, lss layer.Store, rs refstore.Store, loggerImgEvent LogImageEvent, platform *ocispec.Platform) image.Exporter
```

---

## Types

### LogImageEvent

LogImageEvent defines interface for event generation related to image tar(load and save) operations

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/image/tarexport/tarexport.go#L40)  
**Added in:** v1.12.0

```go
type LogImageEvent interface {
	// LogImageEvent generates an event related to an image operation
	LogImageEvent(ctx context.Context, imageID, refName string, action events.Action)
}
```

---


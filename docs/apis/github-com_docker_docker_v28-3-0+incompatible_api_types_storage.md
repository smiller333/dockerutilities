# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/storage

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:14 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### DriverData

DriverData Information about the storage driver used to store the container's and
image's filesystem.

swagger:model DriverData

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/storage/driver_data.go#L10)  

```go
type DriverData struct {

	// Low-level storage metadata, provided as key/value pairs.
	//
	// This information is driver-specific, and depends on the storage-driver
	// in use, and should be used for informational purposes only.
	//
	// Required: true
	Data map[string]string `json:"Data"`

	// Name of the storage driver.
	// Required: true
	Name string `json:"Name"`
}
```

---


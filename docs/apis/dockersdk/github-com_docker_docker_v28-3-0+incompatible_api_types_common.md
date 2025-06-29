# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/common

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:52 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### IDResponse

IDResponse Response to an API call that returns just an Id
swagger:model IDResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/common/id_response.go#L8)  

```go
type IDResponse struct {

	// The id of the newly created object.
	// Required: true
	ID string `json:"Id"`
}
```

---


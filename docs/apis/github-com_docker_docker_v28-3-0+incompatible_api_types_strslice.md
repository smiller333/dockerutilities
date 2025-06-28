# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/strslice

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:17 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### StrSlice

StrSlice represents a string or an array of strings.
We need to override the json decoder to accept both options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/strslice/strslice.go#L7)  

```go
type StrSlice []string
```

#### Methods

##### StrSlice.UnmarshalJSON

UnmarshalJSON decodes the byte slice whether it's a string or an array of
strings. This method is needed to implement json.Unmarshaler.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/strslice/strslice.go#L11)  

```go
func (e *StrSlice) UnmarshalJSON(b []byte) error
```

---


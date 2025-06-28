# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libcontainerd/shimopts

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:28 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Generate

Generate converts opts into a runtime options value for the runtimeType which
can be passed into containerd.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libcontainerd/shimopts/convert.go#L13)  

```go
func Generate(runtimeType string, opts map[string]interface{}) (interface{}, error)
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/osl/kernel

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:58 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ApplyOSTweaks

ApplyOSTweaks applies the configuration values passed as arguments

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/kernel/knobs_linux.go#L30)  

```go
func ApplyOSTweaks(osConfig map[string]*OSValue)
```

---

## Types

### OSValue

OSValue represents a tuple, value defined, check function when to apply the value

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/osl/kernel/knobs.go#L6)  

```go
type OSValue struct {
	Value   string
	CheckFn conditionalCheck
}
```

---


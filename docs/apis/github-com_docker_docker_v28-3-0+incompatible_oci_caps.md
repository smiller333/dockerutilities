# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/oci/caps

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:14 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### DefaultCapabilities

DefaultCapabilities returns a Linux kernel default capabilities

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/caps/defaults.go#L4)  

```go
func DefaultCapabilities() []string
```

---

### GetAllCapabilities

GetAllCapabilities returns all capabilities that are available in the current
environment.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/caps/utils.go#L26)  

```go
func GetAllCapabilities() []string
```

---

### NormalizeLegacyCapabilities

NormalizeLegacyCapabilities normalizes, and validates CapAdd/CapDrop capabilities
by upper-casing them, and adding a CAP_ prefix (if not yet present).

This function also accepts the "ALL" magic-value, that's used by CapAdd/CapDrop.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/caps/utils.go#L56)  

```go
func NormalizeLegacyCapabilities(caps []string) ([]string, error)
```

---

### TweakCapabilities

TweakCapabilities tweaks capabilities by adding, dropping, or overriding
capabilities in the basics capabilities list.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/caps/utils.go#L83)  

```go
func TweakCapabilities(basics, adds, drops []string, privileged bool) ([]string, error)
```

---

## Types

This section is empty.


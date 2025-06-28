# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/profiles/apparmor

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:57 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### InstallDefault

InstallDefault generates a default profile in a temp directory determined by
os.TempDir(), then loads the profile into the kernel using 'apparmor_parser'.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/apparmor/apparmor.go#L59)  

```go
func InstallDefault(name string) error
```

---

### IsLoaded

IsLoaded checks if a profile with the given name has been loaded into the
kernel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/apparmor/apparmor.go#L95)  

```go
func IsLoaded(name string) (bool, error)
```

---

## Types

This section is empty.


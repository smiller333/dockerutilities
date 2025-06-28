# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/remotecontext/git

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:06 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Clone

Clone clones a repository into a newly created directory which
will be under "docker-build-git"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/git/gitutils.go#L36)  

```go
func Clone(remoteURL string, opts ...CloneOption) (string, error)
```

---

## Types

### CloneOption

CloneOption changes the behaviour of Clone().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/git/gitutils.go#L24)  

```go
type CloneOption func(*gitRepo)
```

#### Functions

##### WithIsolatedConfig

WithIsolatedConfig disables reading the user or system gitconfig files when
performing Git operations.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/git/gitutils.go#L28)  

```go
func WithIsolatedConfig(v bool) CloneOption
```

---


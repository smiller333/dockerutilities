# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil/fakegit

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:26 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### FakeGit

FakeGit is a fake git server

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakegit/fakegit.go#L35)  

```go
type FakeGit struct {
	RepoURL string
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New create a fake git server that can be used for git related tests

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakegit/fakegit.go#L48)  

```go
func New(c testing.TB, name string, files map[string]string, enforceLocalServer bool) *FakeGit
```

#### Methods

##### FakeGit.Close

Close closes the server, implements Closer interface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakegit/fakegit.go#L42)  

```go
func (g *FakeGit) Close()
```

---


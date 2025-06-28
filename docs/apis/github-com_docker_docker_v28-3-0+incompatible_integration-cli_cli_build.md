# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration-cli/cli/build

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:53 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### WithBuildContext

WithBuildContext sets up the build context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/build/build.go#L73)  

```go
func WithBuildContext(t testing.TB, contextOperators ...func(*fakecontext.Fake) error) func(*icmd.Cmd) func()
```

---

### WithBuildkit

WithBuildkit sets an DOCKER_BUILDKIT environment variable to make the build use buildkit (or not)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/build/build.go#L35)  

```go
func WithBuildkit(useBuildkit bool) func(*icmd.Cmd) func()
```

---

### WithContextPath

WithContextPath sets the build context path

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/build/build.go#L56)  

```go
func WithContextPath(path string) func(*icmd.Cmd) func()
```

---

### WithDockerfile

WithDockerfile creates / returns a CmdOperator to set the Dockerfile for a build operation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/build/build.go#L26)  

```go
func WithDockerfile(dockerfile string) func(*icmd.Cmd) func()
```

---

### WithExternalBuildContext

WithExternalBuildContext use the specified context as build context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/build/build.go#L64)  

```go
func WithExternalBuildContext(ctx *fakecontext.Fake) func(*icmd.Cmd) func()
```

---

### WithFile

WithFile adds the specified file (with content) in the build context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/build/build.go#L84)  

```go
func WithFile(name, content string) func(*fakecontext.Fake) error
```

---

### WithStdinContext

WithStdinContext sets the build context from the standard input with the specified reader

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/build/build.go#L14)  

```go
func WithStdinContext(closer io.ReadCloser) func(*icmd.Cmd) func()
```

---

### WithoutCache

WithoutCache makes the build ignore cache

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/build/build.go#L50)  

```go
func WithoutCache(cmd *icmd.Cmd) func()
```

---

## Types

This section is empty.


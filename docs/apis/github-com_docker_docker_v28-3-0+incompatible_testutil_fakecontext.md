# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil/fakecontext

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:23 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### WithBinaryFiles

WithBinaryFiles adds the specified files in the build context, content is binary

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L69)  

```go
func WithBinaryFiles(files map[string]*bytes.Buffer) func(*Fake) error
```

---

### WithDockerfile

WithDockerfile adds the specified content as Dockerfile in the build context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L52)  

```go
func WithDockerfile(content string) func(*Fake) error
```

---

### WithFile

WithFile adds the specified file (with content) in the build context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L45)  

```go
func WithFile(name, content string) func(*Fake) error
```

---

### WithFiles

WithFiles adds the specified files in the build context, content is a string

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L57)  

```go
func WithFiles(files map[string]string) func(*Fake) error
```

---

## Types

### Fake

Fake creates directories that can be used as a build context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L81)  

```go
type Fake struct {
	Dir string
}
```

#### Functions

##### New

New creates a fake build context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L14)  

```go
func New(t testing.TB, dir string, modifiers ...func(*Fake) error) *Fake
```

#### Methods

##### Fake.Add

Add a file at a path, creating directories where necessary

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L86)  

```go
func (f *Fake) Add(file, content string) error
```

##### Fake.AsTarReader

AsTarReader returns a ReadCloser with the contents of Dir as a tar archive.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L113)  

```go
func (f *Fake) AsTarReader(t testing.TB) io.ReadCloser
```

##### Fake.Close

Close deletes the context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L108)  

```go
func (f *Fake) Close() error
```

##### Fake.Delete

Delete a file at a path

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/fakecontext/context.go#L102)  

```go
func (f *Fake) Delete(file string) error
```

---


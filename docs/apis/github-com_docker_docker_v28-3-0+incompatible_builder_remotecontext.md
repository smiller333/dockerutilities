# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/remotecontext

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:04 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/detect.go#L27)

```go
const ClientSessionRemote = "client-session"
```

## Variables

This section is empty.

## Functions

### Detect

Detect returns a context and dockerfile from remote location or local
archive.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/detect.go#L31)  

```go
func Detect(config backend.BuildConfig) (remote builder.Source, dockerfile *parser.Result, _ error)
```

---

### FromArchive

FromArchive returns a build source from a tar stream.

It extracts the tar stream to a temporary folder that is deleted as soon as
the Context is closed.
As the extraction happens, a tarsum is calculated for every file, and the set of
all those sums then becomes the source of truth for all operations on this Context.

Closing tarStream has to be done by the caller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/archive.go#L52)  

```go
func FromArchive(tarStream io.Reader) (builder.Source, error)
```

---

### FullPath

FullPath is a helper for getting a full path for a path from a source

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/detect.go#L177)  

```go
func FullPath(remote builder.Source, path string) (string, error)
```

---

### GetWithStatusError

GetWithStatusError does an http.Get() and returns an error if the
status code is 4xx or 5xx.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/remote.go#L46)  

```go
func GetWithStatusError(address string) (*http.Response, error)
```

---

### MakeGitContext

MakeGitContext returns a Context from gitURL that is cloned in a temporary directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/git.go#L14)  

```go
func MakeGitContext(gitURL string) (builder.Source, error)
```

---

### NewFileHash

NewFileHash returns new hash that is used for the builder cache keys

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/filehash.go#L14)  

```go
func NewFileHash(path, name string, fi os.FileInfo) (hash.Hash, error)
```

---

### NewLazySource

NewLazySource creates a new LazyContext. LazyContext defines a hashed build
context based on a root directory. Individual files are hashed first time
they are asked. It is not safe to call methods of LazyContext concurrently.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/lazycontext.go#L16)  

```go
func NewLazySource(root string) (builder.Source, error)
```

---

### Rel

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/lazycontext.go#L89)  

```go
func Rel(basepath string, targpath string) (string, error)
```

---

### StatAt

StatAt is a helper for calling Stat on a path from a source

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/detect.go#L168)  

```go
func StatAt(remote builder.Source, path string) (os.FileInfo, error)
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/server/httputils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:00:59 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### BoolValue

BoolValue transforms a form value in different formats into a boolean type.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L18)  

```go
func BoolValue(r *http.Request, k string) bool
```

---

### BoolValueOrDefault

BoolValueOrDefault returns the default bool passed if the query param is
missing, otherwise it's just a proxy to boolValue above.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L29)  

```go
func BoolValueOrDefault(r *http.Request, k string, d bool) bool
```

---

### CheckForJSON

CheckForJSON makes sure that the request's Content-Type is application/json.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/httputils.go#L48)  

```go
func CheckForJSON(r *http.Request) error
```

---

### CloseStreams

CloseStreams ensures that a list for http streams are properly closed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/httputils.go#L35)  

```go
func CloseStreams(streams ...interface{})
```

---

### DecodePlatform

DecodePlatform decodes the OCI platform JSON string into a Platform struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L146)  

```go
func DecodePlatform(platformJSON string) (*ocispec.Platform, error)
```

---

### DecodePlatforms

DecodePlatforms decodes the OCI platform JSON string into a Platform struct.

Typically, the argument is a value of: r.Form["platform"]

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L169)  

```go
func DecodePlatforms(platformJSONs []string) ([]ocispec.Platform, error)
```

---

### HijackConnection

HijackConnection interrupts the http response writer to get the
underlying connection and operate with it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/httputils.go#L24)  

```go
func HijackConnection(w http.ResponseWriter) (io.ReadCloser, io.Writer, error)
```

---

### Int64ValueOrDefault

Int64ValueOrDefault parses a form value into an int64 type. If there is an
error, returns the error. If there is no value returns the default value.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L71)  

```go
func Int64ValueOrDefault(r *http.Request, field string, def int64) (int64, error)
```

---

### Int64ValueOrZero

Int64ValueOrZero parses a form value into an int64 type.
It returns 0 if the parsing fails.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L61)  

```go
func Int64ValueOrZero(r *http.Request, k string) int64
```

---

### ParseForm

ParseForm ensures the request form is parsed even with invalid content types.
If we don't do this, POST method without Content-type (even with empty body) will fail.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/httputils.go#L100)  

```go
func ParseForm(r *http.Request) error
```

---

### ReadJSON

ReadJSON validates the request to have the correct content-type, and decodes
the request's Body into out.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/httputils.go#L62)  

```go
func ReadJSON(r *http.Request, out interface{}) error
```

---

### RepoTagReference

RepoTagReference parses form values "repo" and "tag" and returns a valid
reference with repository and tag.
If repo is empty, then a nil reference is returned.
If no tag is given, then the default "latest" tag is set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L83)  

```go
func RepoTagReference(repo, tag string) (reference.NamedTagged, error)
```

---

### Uint32Value

Uint32Value parses a form value into an uint32 type. It returns an error
if the field is not set, empty, incorrectly formatted, or out of range.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L38)  

```go
func Uint32Value(r *http.Request, field string) (uint32, error)
```

---

### VersionFromContext

VersionFromContext returns an API version from the context using APIVersionKey.
It panics if the context value does not have version.Version type.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/httputils.go#L112)  

```go
func VersionFromContext(ctx context.Context) string
```

---

### WriteJSON

WriteJSON writes the value v to the http response stream as json with standard json encoding.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/httputils.go#L90)  

```go
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error
```

---

### WriteLogStream

WriteLogStream writes an encoded byte stream of log messages from the
messages channel, multiplexing them with a stdcopy.Writer if mux is true

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/write_log_stream.go#L20)  

```go
func WriteLogStream(_ context.Context, w http.ResponseWriter, msgs <-chan *backend.LogMessage, config *container.LogsOptions, mux bool)
```

---

## Types

### APIFunc

APIFunc is an adapter to allow the use of ordinary functions as Docker API endpoints.
Any function that has the appropriate signature can be registered as an API endpoint (e.g. getVersion).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/httputils.go#L20)  

```go
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
```

---

### APIVersionKey

APIVersionKey is the client's requested API version.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/httputils.go#L16)  

```go
type APIVersionKey struct{}
```

---

### ArchiveOptions

ArchiveOptions stores archive information for different operations.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L112)  

```go
type ArchiveOptions struct {
	Name string
	Path string
}
```

#### Functions

##### ArchiveFormValues

ArchiveFormValues parses form values and turns them into ArchiveOptions.
It fails if the archive name and path are not in the request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/form.go#L129)  

```go
func ArchiveFormValues(r *http.Request, vars map[string]string) (ArchiveOptions, error)
```

---

### ContainerDecoder

ContainerDecoder specifies how
to translate an io.Reader into
container configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/server/httputils/decoder.go#L13)  
**Added in:** v1.12.0

```go
type ContainerDecoder interface {
	DecodeConfig(src io.Reader) (*container.Config, *container.HostConfig, *network.NetworkingConfig, error)
}
```

---


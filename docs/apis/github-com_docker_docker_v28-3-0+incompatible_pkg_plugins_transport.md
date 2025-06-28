# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/plugins/transport

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:12 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/transport/mimetype.go#L6)

```go
const VersionMimetype = "application/vnd.docker.plugins.v1.2+json"
```

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### HTTPTransport

HTTPTransport holds an http.RoundTripper
and information about the scheme and address the transport
sends request to.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/transport/http.go#L12)  

```go
type HTTPTransport struct {
	http.RoundTripper
	// contains filtered or unexported fields
}
```

#### Functions

##### NewHTTPTransport

NewHTTPTransport creates a new HTTPTransport.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/transport/http.go#L19)  

```go
func NewHTTPTransport(r http.RoundTripper, scheme, addr string) *HTTPTransport
```

#### Methods

##### HTTPTransport.NewRequest

NewRequest creates a new http.Request and sets the URL
scheme and address with the transport's fields.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/transport/http.go#L29)  

```go
func (t HTTPTransport) NewRequest(path string, data io.Reader) (*http.Request, error)
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil/registry

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:37 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L17)

```go
const (
	// V2binary is the name of the registry v2 binary
	V2binary = "registry"
	// DefaultURL is the default url that will be used by the registry (if not specified otherwise)
	DefaultURL = "127.0.0.1:5000"
)
```

## Variables

This section is empty.

## Functions

### Htpasswd

Htpasswd sets the auth method with htpasswd

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/ops.go#L6)  

```go
func Htpasswd(c *Config)
```

---

### Token

Token sets the auth method to token, with the specified token url

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/ops.go#L11)  

```go
func Token(tokenURL string) func(*Config)
```

---

### URL

URL sets the registry url

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/ops.go#L19)  

```go
func URL(registryURL string) func(*Config)
```

---

### WithStderr

WithStderr sets the stdout of the registry command to the passed in writer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/ops.go#L33)  

```go
func WithStderr(w io.Writer) func(c *Config)
```

---

### WithStdout

WithStdout sets the stdout of the registry command to the passed in writer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/ops.go#L26)  

```go
func WithStdout(w io.Writer) func(c *Config)
```

---

## Types

### Config

Config contains the test registry configuration

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L36)  

```go
type Config struct {
	// contains filtered or unexported fields
}
```

---

### Mock

Mock represent a registry mock

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry_mock.go#L15)  

```go
type Mock struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewMock

NewMock creates a registry mock

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry_mock.go#L30)  

```go
func NewMock(t testing.TB) (*Mock, error)
```

#### Methods

##### Mock.Close

Close closes mock and releases resources

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry_mock.go#L66)  

```go
func (tr *Mock) Close()
```

##### Mock.RegisterHandler

RegisterHandler register the specified handler for the registry mock

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry_mock.go#L23)  

```go
func (tr *Mock) RegisterHandler(path string, h handlerFunc)
```

##### Mock.URL

URL returns the url of the registry

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry_mock.go#L61)  

```go
func (tr *Mock) URL() string
```

---

### V2

V2 represent a registry version 2

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L25)  

```go
type V2 struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewV2

NewV2 creates a v2 registry server

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L45)  

```go
func NewV2(t testing.TB, ops ...func(*Config)) *V2
```

#### Methods

##### V2.Close

Close kills the registry server

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L159)  

```go
func (r *V2) Close()
```

##### V2.Email

Email returns the configured email of the server

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L225)  

```go
func (r *V2) Email() string
```

##### V2.Password

Password returns the configured password of the server

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L220)  

```go
func (r *V2) Password() string
```

##### V2.Path

Path returns the path where the registry write data

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L230)  

```go
func (r *V2) Path() string
```

##### V2.Ping

Ping sends an http request to the current registry, and fail if it doesn't respond correctly

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L139)  

```go
func (r *V2) Ping() error
```

##### V2.ReadBlobContents

ReadBlobContents read the file corresponding to the specified digest

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L175)  

```go
func (r *V2) ReadBlobContents(t testing.TB, blobDigest digest.Digest) []byte
```

##### V2.TempMoveBlobData

TempMoveBlobData moves the existing data file aside, so that we can replace it with a
malicious blob of data for example.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L192)  

```go
func (r *V2) TempMoveBlobData(t testing.TB, blobDigest digest.Digest) (undo func())
```

##### V2.Username

Username returns the configured user name of the server

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L215)  

```go
func (r *V2) Username() string
```

##### V2.WaitReady

WaitReady waits for the registry to be ready to serve requests (or fail after a while)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L126)  

```go
func (r *V2) WaitReady(t testing.TB)
```

##### V2.WriteBlobContents

WriteBlobContents write the file corresponding to the specified digest with the given content

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/registry/registry.go#L184)  

```go
func (r *V2) WriteBlobContents(t testing.TB, blobDigest digest.Digest, data []byte)
```

---


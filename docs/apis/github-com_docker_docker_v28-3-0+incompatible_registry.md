# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/registry

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:06 UTC

## Overview

Package registry contains client primitives to interact with a remote Docker registry.


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/config.go#L36)

```go
const (
	// DefaultNamespace is the default namespace
	DefaultNamespace = "docker.io"
	// DefaultRegistryHost is the hostname for the default (Docker Hub) registry
	// used for pushing and pulling images. This hostname is hard-coded to handle
	// the conversion from image references without registry name (e.g. "ubuntu",
	// or "ubuntu:latest"), as well as references using the "docker.io" domain
	// name, which is used as canonical reference for images on Docker Hub, but
	// does not match the domain-name of Docker Hub's registry.
	DefaultRegistryHost = "registry-1.docker.io"
	// IndexHostname is the index hostname, used for authentication and image search.
	IndexHostname = "index.docker.io"
	// IndexServer is used for user auth and image search
	IndexServer = "https://" + IndexHostname + "/v1/"
	// IndexName is the name of the index
	IndexName = "docker.io"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/auth.go#L19)

```go
const AuthClientID = "docker"
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/config.go#L54)

```go
var (
	// DefaultV2Registry is the URI of the default (Docker Hub) registry.
	DefaultV2Registry = &url.URL{
		Scheme: "https",
		Host:   DefaultRegistryHost,
	}
)
```

## Functions

### CertsDir

CertsDir is the directory where certificates are stored.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/config.go#L104)  
**Added in:** v1.8.0

```go
func CertsDir() string
```

---

### ConvertToHostname

ConvertToHostname normalizes a registry URL which has http|https prepended
to just its hostname. It is used to match credentials, which may be either
stored as hostname or as hostname including scheme (in legacy configuration
files).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/auth.go#L134)  
**Added in:** v1.13.0

```go
func ConvertToHostname(url string) string
```

---

### GetAuthConfigKey

GetAuthConfigKey special-cases using the full index address of the official
index as the AuthConfig key, and uses the (host)name[:port] for private indexes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/config.go#L390)  
**Added in:** v1.10.0

```go
func GetAuthConfigKey(index *registry.IndexInfo) string
```

---

### Headers

Headers returns request modifiers with a User-Agent and metaHeaders

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/registry.go#L121)  

```go
func Headers(userAgent string, metaHeaders http.Header) []transport.RequestModifier
```

---

### HostCertsDir

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/registry.go#L22)  

```go
func HostCertsDir(hostname string) string
```

---

### NewStaticCredentialStore

NewStaticCredentialStore returns a credential store
which always returns the same credential values.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/auth.go#L43)  
**Added in:** v1.12.0

```go
func NewStaticCredentialStore(auth *registry.AuthConfig) auth.CredentialStore
```

---

### ParseSearchIndexInfo

ParseSearchIndexInfo will use repository name to get back an indexInfo.

TODO(thaJeztah) this function is only used by the CLI, and used to get
information of the registry (to provide credentials if needed). We should
move this function (or equivalent) to the CLI, as it's doing too much just
for that.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/search.go#L153)  
**Added in:** v1.10.0

```go
func ParseSearchIndexInfo(reposName string) (*registry.IndexInfo, error)
```

---

### PingV2Registry

PingV2Registry attempts to ping a v2 registry and on success return a
challenge manager for the supported authentication types.
If a response is received but cannot be interpreted, a PingResponseError will be returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/auth.go#L178)  
**Added in:** v1.11.0

```go
func PingV2Registry(endpoint *url.URL, transport http.RoundTripper) (challenge.Manager, error)
```

---

### ReadCertsDirectory

ReadCertsDirectory reads the directory for TLS certificates
including roots and certificate pairs and updates the
provided TLS configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/registry.go#L60)  
**Added in:** v1.8.0

```go
func ReadCertsDirectory(tlsConfig *tls.Config, directory string) error
```

---

### ResolveAuthConfig

ResolveAuthConfig matches an auth configuration to a server address or a URL

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/auth.go#L146)  
**Added in:** v1.7.0

```go
func ResolveAuthConfig(authConfigs map[string]registry.AuthConfig, index *registry.IndexInfo) registry.AuthConfig
```

---

### SetCertsDir

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/config.go#L99)  

```go
func SetCertsDir(path string)
```

---

### ValidateIndexName

ValidateIndexName validates an index name. It is used by the daemon to
validate the daemon configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/config.go#L326)  
**Added in:** v1.5.0

```go
func ValidateIndexName(val string) (string, error)
```

---

### ValidateMirror

ValidateMirror validates and normalizes an HTTP(S) registry mirror. It
returns an error if the given mirrorURL is invalid, or the normalized
format for the URL otherwise.

It is used by the daemon to validate the daemon configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/config.go#L300)  
**Added in:** v1.5.0

```go
func ValidateMirror(mirrorURL string) (string, error)
```

---

## Types

### APIEndpoint

APIEndpoint represents a remote API endpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L141)  
**Added in:** v1.8.0

```go
type APIEndpoint struct {
	Mirror                         bool
	URL                            *url.URL
	AllowNondistributableArtifacts bool // Deprecated: non-distributable artifacts are deprecated and enabled by default. This field will be removed in the next release.
	Official                       bool // Deprecated: this field was only used internally, and will be removed in the next release.
	TrimHostname                   bool // Deprecated: hostname is now trimmed unconditionally for remote names. This field will be removed in the next release.
	TLSConfig                      *tls.Config
}
```

---

### PingResponseError

PingResponseError is used when the response from a ping
was received but invalid.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/auth.go#L167)  
**Added in:** v1.11.0

```go
type PingResponseError struct {
	Err error
}
```

#### Methods

##### PingResponseError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/auth.go#L171)  
**Added in:** v1.11.0

```go
func (err PingResponseError) Error() string
```

---

### RepositoryInfo

RepositoryInfo describes a repository

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/types.go#L9)  
**Added in:** v1.5.0

```go
type RepositoryInfo struct {
	Name reference.Named
	// Index points to registry information
	Index *registry.IndexInfo
	// Official indicates whether the repository is considered official.
	// If the registry is official, and the normalized name does not
	// contain a '/' (e.g. "foo"), then it is considered an official repo.
	//
	// Deprecated: this field is no longer used and will be removed in the next release. The information captured in this field can be obtained from the [Name] field instead.
	Official bool
	// Class represents the class of the repository, such as "plugin"
	// or "image".
	//
	// Deprecated: this field is no longer used, and will be removed in the next release.
	Class string
}
```

#### Functions

##### ParseRepositoryInfo

ParseRepositoryInfo performs the breakdown of a repository name into a
RepositoryInfo, but lacks registry configuration.

It is used by the Docker cli to interact with registry-related endpoints.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/config.go#L420)  
**Added in:** v1.5.0

```go
func ParseRepositoryInfo(reposName reference.Named) (*RepositoryInfo, error)
```

---

### Service

Service is a registry service. It tracks configuration data such as a list
of mirrors.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L19)  
**Added in:** v0.11.0

```go
type Service struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewService

NewService returns a new instance of Service ready to be installed into
an engine.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L26)  
**Added in:** v0.11.0

```go
func NewService(options ServiceOptions) (*Service, error)
```

#### Methods

##### Service.Auth

Auth contacts the public registry with the provided credentials,
and returns OK if authentication was successful.
It can be used to verify the validity of a client's credentials.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L59)  
**Added in:** v0.11.0

```go
func (s *Service) Auth(ctx context.Context, authConfig *registry.AuthConfig, userAgent string) (statusMessage, token string, _ error)
```

##### Service.IsInsecureRegistry

IsInsecureRegistry returns true if the registry at given host is configured as
insecure registry.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L170)  

```go
func (s *Service) IsInsecureRegistry(host string) bool
```

##### Service.LookupPullEndpoints

LookupPullEndpoints creates a list of v2 endpoints to try to pull from, in order of preference.
It gives preference to mirrors over the actual registry, and HTTPS over plain HTTP.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L152)  
**Added in:** v1.8.0

```go
func (s *Service) LookupPullEndpoints(hostname string) ([]APIEndpoint, error)
```

##### Service.LookupPushEndpoints

LookupPushEndpoints creates a list of v2 endpoints to try to push to, in order of preference.
It gives preference to HTTPS over plain HTTP. Mirrors are not included.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L161)  
**Added in:** v1.8.0

```go
func (s *Service) LookupPushEndpoints(hostname string) ([]APIEndpoint, error)
```

##### Service.ReplaceConfig

ReplaceConfig prepares a transaction which will atomically replace the
registry service's configuration when the returned commit function is called.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L44)  

```go
func (s *Service) ReplaceConfig(options ServiceOptions) (commit func(), _ error)
```

##### Service.ResolveAuthConfig

ResolveAuthConfig looks up authentication for the given reference from the
given authConfigs.

IMPORTANT: This function is for internal use and should not be used by external projects.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L126)  

```go
func (s *Service) ResolveAuthConfig(authConfigs map[string]registry.AuthConfig, ref reference.Named) registry.AuthConfig
```

##### Service.ResolveRepository

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L115)  
**Added in:** v1.5.0

```go
func (s *Service) ResolveRepository(name reference.Named) (*RepositoryInfo, error)
```

##### Service.Search

Search queries the public registry for repositories matching the specified
search term and filters.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/search.go#L24)  
**Added in:** v0.11.0

```go
func (s *Service) Search(ctx context.Context, searchFilters filters.Args, term string, limit int, authConfig *registry.AuthConfig, headers map[string][]string) ([]registry.SearchResult, error)
```

##### Service.ServiceConfig

ServiceConfig returns a copy of the public registry service's configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/service.go#L36)  
**Added in:** v1.11.0

```go
func (s *Service) ServiceConfig() *registry.ServiceConfig
```

---

### ServiceOptions

ServiceOptions holds command line options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/registry/config.go#L21)  
**Added in:** v1.11.0

```go
type ServiceOptions struct {
	Mirrors            []string `json:"registry-mirrors,omitempty"`
	InsecureRegistries []string `json:"insecure-registries,omitempty"`
}
```

---


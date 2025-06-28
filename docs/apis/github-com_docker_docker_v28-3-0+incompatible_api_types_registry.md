# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/registry

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:11 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/authconfig.go#L14)

```go
const AuthHeader = "X-Registry-Auth"
```

## Variables

This section is empty.

## Functions

### EncodeAuthConfig

EncodeAuthConfig serializes the auth configuration as a base64url encoded
(RFC4648, section 5) JSON string for sending through the X-Registry-Auth header.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/authconfig.go#L53)  

```go
func EncodeAuthConfig(authConfig AuthConfig) (string, error)
```

---

## Types

### AuthConfig

AuthConfig contains authorization information for connecting to a Registry.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/authconfig.go#L29)  

```go
type AuthConfig struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Auth     string `json:"auth,omitempty"`

	// Email is an optional value associated with the username.
	// This field is deprecated and will be removed in a later
	// version of docker.
	Email string `json:"email,omitempty"`

	ServerAddress string `json:"serveraddress,omitempty"`

	// IdentityToken is used to authenticate the user and get
	// an access token for the registry.
	IdentityToken string `json:"identitytoken,omitempty"`

	// RegistryToken is a bearer token to be sent to a registry
	RegistryToken string `json:"registrytoken,omitempty"`
}
```

#### Functions

##### DecodeAuthConfig

DecodeAuthConfig decodes base64url encoded (RFC4648, section 5) JSON
authentication information as sent through the X-Registry-Auth header.

This function always returns an AuthConfig, even if an error occurs. It is up
to the caller to decide if authentication is required, and if the error can
be ignored.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/authconfig.go#L69)  

```go
func DecodeAuthConfig(authEncoded string) (*AuthConfig, error)
```

##### DecodeAuthConfigBody

DecodeAuthConfigBody decodes authentication information as sent as JSON in the
body of a request. This function is to provide backward compatibility with old
clients and API versions. Current clients and API versions expect authentication
to be provided through the X-Registry-Auth header.

Like DecodeAuthConfig, this function always returns an AuthConfig, even if an
error occurs. It is up to the caller to decide if authentication is required,
and if the error can be ignored.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/authconfig.go#L86)  

```go
func DecodeAuthConfigBody(rdr io.ReadCloser) (*AuthConfig, error)
```

---

### AuthenticateOKBody

AuthenticateOKBody authenticate o k body
swagger:model AuthenticateOKBody

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/authenticate.go#L12)  

```go
type AuthenticateOKBody struct {

	// An opaque token used to authenticate a user after a successful login
	// Required: true
	IdentityToken string `json:"IdentityToken"`

	// The status of the authentication
	// Required: true
	Status string `json:"Status"`
}
```

---

### DistributionInspect

DistributionInspect describes the result obtained from contacting the
registry to retrieve image metadata

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/registry.go#L115)  

```go
type DistributionInspect struct {
	// Descriptor contains information about the manifest, including
	// the content addressable digest
	Descriptor ocispec.Descriptor
	// Platforms contains the list of platforms supported by the image,
	// obtained by parsing the manifest
	Platforms []ocispec.Platform
}
```

---

### IndexInfo

IndexInfo contains information about a registry

RepositoryInfo Examples:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/registry.go#L100)  

```go
type IndexInfo struct {
	// Name is the name of the registry, such as "docker.io"
	Name string
	// Mirrors is a list of mirrors, expressed as URIs
	Mirrors []string
	// Secure is set to false if the registry is part of the list of
	// insecure registries. Insecure registries accept HTTP and/or accept
	// HTTPS with certificates from unknown CAs.
	Secure bool
	// Official indicates whether this is an official registry
	Official bool
}
```

---

### NetIPNet

NetIPNet is the net.IPNet type, which can be marshalled and
unmarshalled to JSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/registry.go#L45)  

```go
type NetIPNet net.IPNet
```

#### Methods

##### NetIPNet.MarshalJSON

MarshalJSON returns the JSON representation of the IPNet

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/registry.go#L53)  

```go
func (ipnet *NetIPNet) MarshalJSON() ([]byte, error)
```

##### NetIPNet.String

String returns the CIDR notation of ipnet

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/registry.go#L48)  

```go
func (ipnet *NetIPNet) String() string
```

##### NetIPNet.UnmarshalJSON

UnmarshalJSON sets the IPNet from a byte array of JSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/registry.go#L58)  

```go
func (ipnet *NetIPNet) UnmarshalJSON(b []byte) error
```

---

### RequestAuthConfig

RequestAuthConfig is a function interface that clients can supply
to retry operations after getting an authorization error.

The function must return the AuthHeader value (AuthConfig), encoded
in base64url format (RFC4648, section 5), which can be decoded by
DecodeAuthConfig.

It must return an error if the privilege request fails.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/authconfig.go#L26)  

```go
type RequestAuthConfig func(context.Context) (string, error)
```

---

### SearchOptions

SearchOptions holds parameters to search images with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/search.go#L10)  

```go
type SearchOptions struct {
	RegistryAuth string

	// PrivilegeFunc is a function that clients can supply to retry operations
	// after getting an authorization error. This function returns the registry
	// authentication header value in base64 encoded format, or an error if the
	// privilege request fails.
	//
	// For details, refer to [github.com/docker/docker/api/types/registry.RequestAuthConfig].
	PrivilegeFunc func(context.Context) (string, error)
	Filters       filters.Args
	Limit         int
}
```

---

### SearchResult

SearchResult describes a search result returned from a registry

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/search.go#L25)  

```go
type SearchResult struct {
	// StarCount indicates the number of stars this repository has
	StarCount int `json:"star_count"`
	// IsOfficial is true if the result is from an official repository.
	IsOfficial bool `json:"is_official"`
	// Name is the name of the repository
	Name string `json:"name"`
	// IsAutomated indicates whether the result is automated.
	//
	// Deprecated: the "is_automated" field is deprecated and will always be "false".
	IsAutomated bool `json:"is_automated"`
	// Description is a textual description of the repository
	Description string `json:"description"`
}
```

---

### SearchResults

SearchResults lists a collection search results returned from a registry

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/search.go#L41)  

```go
type SearchResults struct {
	// Query contains the query string that generated the search results
	Query string `json:"query"`
	// NumResults indicates the number of results the query returned
	NumResults int `json:"num_results"`
	// Results is a slice containing the actual results for the search
	Results []SearchResult `json:"results"`
}
```

---

### ServiceConfig

ServiceConfig stores daemon registry services configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/registry.go#L14)  

```go
type ServiceConfig struct {
	AllowNondistributableArtifactsCIDRs     []*NetIPNet `json:"AllowNondistributableArtifactsCIDRs,omitempty"`     // Deprecated: non-distributable artifacts are deprecated and enabled by default. This field will be removed in the next release.
	AllowNondistributableArtifactsHostnames []string    `json:"AllowNondistributableArtifactsHostnames,omitempty"` // Deprecated: non-distributable artifacts are deprecated and enabled by default. This field will be removed in the next release.

	InsecureRegistryCIDRs []*NetIPNet           `json:"InsecureRegistryCIDRs"`
	IndexConfigs          map[string]*IndexInfo `json:"IndexConfigs"`
	Mirrors               []string

	// ExtraFields is for internal use to include deprecated fields on older API versions.
	ExtraFields map[string]any `json:"-"`
}
```

#### Methods

##### ServiceConfig.MarshalJSON

MarshalJSON implements a custom marshaler to include legacy fields
in API responses.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/registry/registry.go#L28)  

```go
func (sc *ServiceConfig) MarshalJSON() ([]byte, error)
```

---


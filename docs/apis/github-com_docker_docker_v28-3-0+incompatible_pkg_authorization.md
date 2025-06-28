# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/authorization

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:24 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/api.go#L9)

```go
const (
	// AuthZApiRequest is the url for daemon request authorization
	AuthZApiRequest = "AuthZPlugin.AuthZReq"

	// AuthZApiResponse is the url for daemon response authorization
	AuthZApiResponse = "AuthZPlugin.AuthZRes"

	// AuthZApiImplements is the name of the interface all AuthZ plugins implement
	AuthZApiImplements = "authz"
)
```

## Variables

This section is empty.

## Functions

### GetPluginGetter

GetPluginGetter gets the plugingetter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/plugin.go#L45)  
**Added in:** v1.13.0

```go
func GetPluginGetter() plugingetter.PluginGetter
```

---

### SetPluginGetter

SetPluginGetter sets the plugingetter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/plugin.go#L40)  
**Added in:** v1.13.0

```go
func SetPluginGetter(pg plugingetter.PluginGetter)
```

---

## Types

### Ctx

Ctx stores a single request-response interaction context

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/authz.go#L48)  

```go
type Ctx struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewCtx

NewCtx creates new authZ context, it is used to store authorization information related to a specific docker
REST http session
A context provides two method:
Authenticate Request:
Call authZ plugins with current REST request and AuthN response
Request contains full HTTP packet sent to the docker daemon
https://docs.docker.com/reference/api/engine/

Authenticate Response:
Call authZ plugins with full info about current REST request, REST response and AuthN response
The response from this method may contains content that overrides the daemon response
This allows authZ plugins to filter privileged content

If multiple authZ plugins are specified, the block/allow decision is based on ANDing all plugin results
For response manipulation, the response from each plugin is piped between plugins. Plugin execution order
is determined according to daemon parameters

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/authz.go#L37)  

```go
func NewCtx(authZPlugins []Plugin, user, userAuthNMethod, requestMethod, requestURI string) *Ctx
```

#### Methods

##### Ctx.AuthZRequest

AuthZRequest authorized the request to the docker daemon using authZ plugins

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/authz.go#L72)  

```go
func (ctx *Ctx) AuthZRequest(w http.ResponseWriter, r *http.Request) error
```

##### Ctx.AuthZResponse

AuthZResponse authorized and manipulates the response from docker daemon using authZ plugins

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/authz.go#L120)  

```go
func (ctx *Ctx) AuthZResponse(rm ResponseModifier, r *http.Request) error
```

---

### Middleware

Middleware uses a list of plugins to
handle authorization in the API requests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/middleware.go#L14)  
**Added in:** v1.12.0

```go
type Middleware struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewMiddleware

NewMiddleware creates a new Middleware
with a slice of plugins names.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/middleware.go#L21)  
**Added in:** v1.12.0

```go
func NewMiddleware(names []string, pg plugingetter.PluginGetter) *Middleware
```

#### Methods

##### Middleware.RemovePlugin

RemovePlugin removes a single plugin from this authz middleware chain

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/middleware.go#L42)  

```go
func (m *Middleware) RemovePlugin(name string)
```

##### Middleware.SetPlugins

SetPlugins sets the plugin used for authorization

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/middleware.go#L35)  
**Added in:** v1.13.0

```go
func (m *Middleware) SetPlugins(names []string)
```

##### Middleware.WrapHandler

WrapHandler returns a new handler function wrapping the previous one in the request chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/middleware.go#L55)  
**Added in:** v1.12.0

```go
func (m *Middleware) WrapHandler(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error) func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
```

---

### PeerCertificate

PeerCertificate is a wrapper around x509.Certificate which provides a sane
encoding/decoding to/from PEM format and JSON.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/api.go#L22)  
**Added in:** v1.13.0

```go
type PeerCertificate x509.Certificate
```

#### Methods

##### PeerCertificate.MarshalJSON

MarshalJSON returns the JSON encoded pem bytes of a PeerCertificate.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/api.go#L25)  
**Added in:** v1.13.0

```go
func (pc *PeerCertificate) MarshalJSON() ([]byte, error)
```

##### PeerCertificate.UnmarshalJSON

UnmarshalJSON populates a new PeerCertificate struct from JSON data.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/api.go#L31)  
**Added in:** v1.13.0

```go
func (pc *PeerCertificate) UnmarshalJSON(b []byte) error
```

---

### Plugin

Plugin allows third party plugins to authorize requests and responses
in the context of docker API

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/plugin.go#L12)  

```go
type Plugin interface {
	// Name returns the registered plugin name
	Name() string

	// AuthZRequest authorizes the request from the client to the daemon
	AuthZRequest(*Request) (*Response, error)

	// AuthZResponse authorizes the response from the daemon to the client
	AuthZResponse(*Request) (*Response, error)
}
```

---

### Request

Request holds data required for authZ plugins

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/api.go#L46)  

```go
type Request struct {
	// User holds the user extracted by AuthN mechanism
	User string `json:"User,omitempty"`

	// UserAuthNMethod holds the mechanism used to extract user details (e.g., krb)
	UserAuthNMethod string `json:"UserAuthNMethod,omitempty"`

	// RequestMethod holds the HTTP method (GET/POST/PUT)
	RequestMethod string `json:"RequestMethod,omitempty"`

	// RequestUri holds the full HTTP uri (e.g., /v1.21/version)
	RequestURI string `json:"RequestUri,omitempty"`

	// RequestBody stores the raw request body sent to the docker daemon
	RequestBody []byte `json:"RequestBody,omitempty"`

	// RequestHeaders stores the raw request headers sent to the docker daemon
	RequestHeaders map[string]string `json:"RequestHeaders,omitempty"`

	// RequestPeerCertificates stores the request's TLS peer certificates in PEM format
	RequestPeerCertificates []*PeerCertificate `json:"RequestPeerCertificates,omitempty"`

	// ResponseStatusCode stores the status code returned from docker daemon
	ResponseStatusCode int `json:"ResponseStatusCode,omitempty"`

	// ResponseBody stores the raw response body sent from docker daemon
	ResponseBody []byte `json:"ResponseBody,omitempty"`

	// ResponseHeaders stores the response headers sent to the docker daemon
	ResponseHeaders map[string]string `json:"ResponseHeaders,omitempty"`
}
```

---

### Response

Response represents authZ plugin response

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/api.go#L79)  

```go
type Response struct {
	// Allow indicating whether the user is allowed or not
	Allow bool `json:"Allow"`

	// Msg stores the authorization message
	Msg string `json:"Msg,omitempty"`

	// Err stores a message in case there's an error
	Err string `json:"Err,omitempty"`
}
```

---

### ResponseModifier

ResponseModifier allows authorization plugins to read and modify the content of the http.response

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/response.go#L16)  

```go
type ResponseModifier interface {
	http.ResponseWriter
	http.Flusher

	// RawBody returns the current http content
	RawBody() []byte

	// RawHeaders returns the current content of the http headers
	RawHeaders() ([]byte, error)

	// StatusCode returns the current status code
	StatusCode() int

	// OverrideBody replaces the body of the HTTP reply
	OverrideBody(b []byte)

	// OverrideHeader replaces the headers of the HTTP reply
	OverrideHeader(b []byte) error

	// OverrideStatusCode replaces the status code of the HTTP reply
	OverrideStatusCode(statusCode int)

	// FlushAll flushes all data to the HTTP response
	FlushAll() error

	// Hijacked indicates the response has been hijacked by the Docker daemon
	Hijacked() bool
}
```

#### Functions

##### NewResponseModifier

NewResponseModifier creates a wrapper to an http.ResponseWriter to allow inspecting and modifying the content

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/authorization/response.go#L46)  

```go
func NewResponseModifier(rw http.ResponseWriter) ResponseModifier
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:39 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types.go#L11)

```go
const (
	// MediaTypeRawStream is vendor specific MIME-Type set for raw TTY streams
	MediaTypeRawStream = "application/vnd.docker.raw-stream"

	// MediaTypeMultiplexedStream is vendor specific MIME-Type set for stdin/stdout/stderr multiplexed streams
	MediaTypeMultiplexedStream = "application/vnd.docker.multiplexed-stream"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L66)

```go
const (
	NoHealthcheck = container.NoHealthcheck // Deprecated: use [container.NoHealthcheck].
	Starting      = container.Starting      // Deprecated: use [container.Starting].
	Healthy       = container.Healthy       // Deprecated: use [container.Healthy].
	Unhealthy     = container.Unhealthy     // Deprecated: use [container.Unhealthy].
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L166)

```go
const (
	RegistryAuthFromSpec         = swarm.RegistryAuthFromSpec         // Deprecated: use [swarm.RegistryAuthFromSpec].
	RegistryAuthFromPreviousSpec = swarm.RegistryAuthFromPreviousSpec // Deprecated: use [swarm.RegistryAuthFromPreviousSpec].
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L232)

```go
const (
	// BuilderV1 is the first generation builder in docker daemon
	//
	// Deprecated: use [build.BuilderV1].
	BuilderV1 = build.BuilderV1
	// BuilderBuildKit is builder based on moby/buildkit project
	//
	// Deprecated: use [build.BuilderBuildKit].
	BuilderBuildKit = build.BuilderBuildKit
)
```

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### BuildCache

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L191)  

```go
type BuildCache = build.CacheRecord
```

---

### BuildCachePruneOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L196)  

```go
type BuildCachePruneOptions = build.CachePruneOptions
```

---

### BuildCachePruneReport

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L202)  

```go
type BuildCachePruneReport = build.CachePruneReport
```

---

### BuildResult

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L207)  

```go
type BuildResult = build.Result
```

---

### BuilderVersion

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L230)  

```go
type BuilderVersion = build.BuilderVersion
```

---

### CloseWriter

CloseWriter is an interface that implements structs
that close input streams to prevent from writing.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L37)  
**Added in:** v1.13.0

```go
type CloseWriter interface {
	CloseWrite() error
}
```

---

### ComponentVersion

ComponentVersion describes the version information for a specific component.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types.go#L38)  

```go
type ComponentVersion struct {
	Name    string
	Version string
	Details map[string]string `json:",omitempty"`
}
```

---

### ConfigCreateResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L134)  

```go
type ConfigCreateResponse = swarm.ConfigCreateResponse
```

---

### ConfigListOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L139)  

```go
type ConfigListOptions = swarm.ConfigListOptions
```

---

### Container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L35)  
**Added in:** v1.7.0

```go
type Container = container.Summary
```

---

### ContainerJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L29)  
**Added in:** v1.7.0

```go
type ContainerJSON = container.InspectResponse
```

---

### ContainerJSONBase

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L23)  
**Added in:** v1.7.0

```go
type ContainerJSONBase = container.ContainerJSONBase
```

---

### ContainerState

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L40)  
**Added in:** v1.7.0

```go
type ContainerState = container.State
```

---

### DefaultNetworkSettings

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L57)  
**Added in:** v1.9.0

```go
type DefaultNetworkSettings = container.DefaultNetworkSettings
```

---

### DiskUsage

DiskUsage contains response of Engine API:
GET "/system/df"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types.go#L87)  
**Added in:** v1.13.0

```go
type DiskUsage struct {
	LayersSize  int64
	Images      []*image.Summary
	Containers  []*container.Summary
	Volumes     []*volume.Volume
	BuildCache  []*build.CacheRecord
	BuilderSize int64 `json:",omitempty"` // Deprecated: deprecated in API 1.38, and no longer used since API 1.40.
}
```

---

### DiskUsageObject

DiskUsageObject represents an object type used for disk usage query filtering.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types.go#L65)  

```go
type DiskUsageObject string
```

---

### DiskUsageOptions

DiskUsageOptions holds parameters for system disk usage query.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types.go#L79)  

```go
type DiskUsageOptions struct {
	// Types specifies what object types to include in the response. If empty,
	// all object types are returned.
	Types []DiskUsageObject
}
```

---

### ErrorResponse

ErrorResponse Represents an error.
swagger:model ErrorResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/error_response.go#L8)  
**Added in:** v1.13.0

```go
type ErrorResponse struct {

	// The error message.
	// Required: true
	Message string `json:"message"`
}
```

#### Methods

##### ErrorResponse.Error

Error returns the error message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/error_response_ext.go#L4)  

```go
func (e ErrorResponse) Error() string
```

---

### GraphDriverData

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L98)  
**Added in:** v1.8.0

```go
type GraphDriverData = storage.DriverData
```

---

### Health

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L76)  
**Added in:** v1.13.0

```go
type Health = container.Health
```

---

### HealthcheckResult

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L81)  
**Added in:** v1.13.0

```go
type HealthcheckResult = container.HealthcheckResult
```

---

### HijackedResponse

HijackedResponse holds connection information for a hijacked request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L15)  
**Added in:** v1.13.0

```go
type HijackedResponse struct {
	Conn   net.Conn
	Reader *bufio.Reader
	// contains filtered or unexported fields
}
```

#### Functions

##### NewHijackedResponse

NewHijackedResponse initializes a HijackedResponse type.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L10)  

```go
func NewHijackedResponse(conn net.Conn, mediaType string) HijackedResponse
```

#### Methods

##### HijackedResponse.Close

Close closes the hijacked connection and reader.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L22)  
**Added in:** v1.13.0

```go
func (h *HijackedResponse) Close()
```

##### HijackedResponse.CloseWrite

CloseWrite closes a readWriter for writing.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L42)  
**Added in:** v1.13.0

```go
func (h *HijackedResponse) CloseWrite() error
```

##### HijackedResponse.MediaType

MediaType let client know if HijackedResponse hold a raw or multiplexed stream.
returns false if HTTP Content-Type is not relevant, and container must be inspected

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L28)  

```go
func (h *HijackedResponse) MediaType() (string, bool)
```

---

### IDResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L17)  
**Added in:** v1.13.0

```go
type IDResponse = common.IDResponse
```

---

### ImageBuildOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L213)  
**Added in:** v1.13.0

```go
type ImageBuildOptions = build.ImageBuildOptions
```

---

### ImageBuildOutput

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L218)  

```go
type ImageBuildOutput = build.ImageBuildOutput
```

---

### ImageBuildResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L225)  
**Added in:** v1.13.0

```go
type ImageBuildResponse = build.ImageBuildResponse
```

---

### ImageInspect

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L109)  
**Added in:** v1.7.0

```go
type ImageInspect = image.InspectResponse
```

---

### MountPoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L87)  
**Added in:** v1.8.0

```go
type MountPoint = container.MountPoint
```

---

### NetworkSettings

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L45)  
**Added in:** v1.9.0

```go
type NetworkSettings = container.NetworkSettings
```

---

### NetworkSettingsBase

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L50)  
**Added in:** v1.9.0

```go
type NetworkSettingsBase = container.NetworkSettingsBase
```

---

### NodeListOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L144)  
**Added in:** v1.13.0

```go
type NodeListOptions = swarm.NodeListOptions
```

---

### NodeRemoveOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L149)  
**Added in:** v1.13.0

```go
type NodeRemoveOptions = swarm.NodeRemoveOptions
```

---

### Ping

Ping contains response of Engine API:
GET "/_ping"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types.go#L21)  
**Added in:** v1.13.0

```go
type Ping struct {
	APIVersion     string
	OSType         string
	Experimental   bool
	BuilderVersion build.BuilderVersion

	// SwarmStatus provides information about the current swarm status of the
	// engine, obtained from the "Swarm" header in the API response.
	//
	// It can be a nil struct if the API version does not provide this header
	// in the ping response, or if an error occurred, in which case the client
	// should use other ways to get the current swarm status, such as the /swarm
	// endpoint.
	SwarmStatus *swarm.Status
}
```

---

### Plugin

Plugin A plugin for the Engine API
swagger:model Plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin.go#L8)  
**Added in:** v1.13.0

```go
type Plugin struct {

	// config
	// Required: true
	Config PluginConfig `json:"Config"`

	// True if the plugin is running. False if the plugin is not running, only installed.
	// Required: true
	Enabled bool `json:"Enabled"`

	// Id
	ID string `json:"Id,omitempty"`

	// name
	// Required: true
	Name string `json:"Name"`

	// plugin remote reference used to push/pull the plugin
	PluginReference string `json:"PluginReference,omitempty"`

	// settings
	// Required: true
	Settings PluginSettings `json:"Settings"`
}
```

---

### PluginConfig

PluginConfig The config of a plugin.
swagger:model PluginConfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin.go#L35)  
**Added in:** v1.13.0

```go
type PluginConfig struct {

	// args
	// Required: true
	Args PluginConfigArgs `json:"Args"`

	// description
	// Required: true
	Description string `json:"Description"`

	// Docker Version used to create the plugin
	DockerVersion string `json:"DockerVersion,omitempty"`

	// documentation
	// Required: true
	Documentation string `json:"Documentation"`

	// entrypoint
	// Required: true
	Entrypoint []string `json:"Entrypoint"`

	// env
	// Required: true
	Env []PluginEnv `json:"Env"`

	// interface
	// Required: true
	Interface PluginConfigInterface `json:"Interface"`

	// ipc host
	// Required: true
	IpcHost bool `json:"IpcHost"`

	// linux
	// Required: true
	Linux PluginConfigLinux `json:"Linux"`

	// mounts
	// Required: true
	Mounts []PluginMount `json:"Mounts"`

	// network
	// Required: true
	Network PluginConfigNetwork `json:"Network"`

	// pid host
	// Required: true
	PidHost bool `json:"PidHost"`

	// propagated mount
	// Required: true
	PropagatedMount string `json:"PropagatedMount"`

	// user
	User PluginConfigUser `json:"User,omitempty"`

	// work dir
	// Required: true
	WorkDir string `json:"WorkDir"`

	// rootfs
	Rootfs *PluginConfigRootfs `json:"rootfs,omitempty"`
}
```

---

### PluginConfigArgs

PluginConfigArgs plugin config args
swagger:model PluginConfigArgs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin.go#L101)  
**Added in:** v1.13.0

```go
type PluginConfigArgs struct {

	// description
	// Required: true
	Description string `json:"Description"`

	// name
	// Required: true
	Name string `json:"Name"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`

	// value
	// Required: true
	Value []string `json:"Value"`
}
```

---

### PluginConfigInterface

PluginConfigInterface The interface between Docker and the plugin
swagger:model PluginConfigInterface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin.go#L122)  
**Added in:** v1.13.0

```go
type PluginConfigInterface struct {

	// Protocol to use for clients connecting to the plugin.
	ProtocolScheme string `json:"ProtocolScheme,omitempty"`

	// socket
	// Required: true
	Socket string `json:"Socket"`

	// types
	// Required: true
	Types []PluginInterfaceType `json:"Types"`
}
```

---

### PluginConfigLinux

PluginConfigLinux plugin config linux
swagger:model PluginConfigLinux

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin.go#L138)  
**Added in:** v1.13.0

```go
type PluginConfigLinux struct {

	// allow all devices
	// Required: true
	AllowAllDevices bool `json:"AllowAllDevices"`

	// capabilities
	// Required: true
	Capabilities []string `json:"Capabilities"`

	// devices
	// Required: true
	Devices []PluginDevice `json:"Devices"`
}
```

---

### PluginConfigNetwork

PluginConfigNetwork plugin config network
swagger:model PluginConfigNetwork

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin.go#L155)  
**Added in:** v1.13.0

```go
type PluginConfigNetwork struct {

	// type
	// Required: true
	Type string `json:"Type"`
}
```

---

### PluginConfigRootfs

PluginConfigRootfs plugin config rootfs
swagger:model PluginConfigRootfs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin.go#L164)  
**Added in:** v1.13.0

```go
type PluginConfigRootfs struct {

	// diff ids
	DiffIds []string `json:"diff_ids"`

	// type
	Type string `json:"type,omitempty"`
}
```

---

### PluginConfigUser

PluginConfigUser plugin config user
swagger:model PluginConfigUser

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin.go#L175)  
**Added in:** v1.13.0

```go
type PluginConfigUser struct {

	// g ID
	GID uint32 `json:"GID,omitempty"`

	// UID
	UID uint32 `json:"UID,omitempty"`
}
```

---

### PluginCreateOptions

PluginCreateOptions hold all options to plugin create.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L83)  
**Added in:** v1.13.0

```go
type PluginCreateOptions struct {
	RepoName string
}
```

---

### PluginDevice

PluginDevice plugin device
swagger:model PluginDevice

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_device.go#L8)  
**Added in:** v1.13.0

```go
type PluginDevice struct {

	// description
	// Required: true
	Description string `json:"Description"`

	// name
	// Required: true
	Name string `json:"Name"`

	// path
	// Required: true
	Path *string `json:"Path"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`
}
```

---

### PluginDisableOptions

PluginDisableOptions holds parameters to disable plugins.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L60)  
**Added in:** v1.13.0

```go
type PluginDisableOptions struct {
	Force bool
}
```

---

### PluginEnableOptions

PluginEnableOptions holds parameters to enable plugins.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L55)  
**Added in:** v1.13.0

```go
type PluginEnableOptions struct {
	Timeout int
}
```

---

### PluginEnv

PluginEnv plugin env
swagger:model PluginEnv

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_env.go#L8)  
**Added in:** v1.13.0

```go
type PluginEnv struct {

	// description
	// Required: true
	Description string `json:"Description"`

	// name
	// Required: true
	Name string `json:"Name"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`

	// value
	// Required: true
	Value *string `json:"Value"`
}
```

---

### PluginInstallOptions

PluginInstallOptions holds parameters to install a plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L65)  
**Added in:** v1.13.0

```go
type PluginInstallOptions struct {
	Disabled             bool
	AcceptAllPermissions bool
	RegistryAuth         string // RegistryAuth is the base64 encoded credentials for the registry
	RemoteRef            string // RemoteRef is the plugin name on the registry

	// PrivilegeFunc is a function that clients can supply to retry operations
	// after getting an authorization error. This function returns the registry
	// authentication header value in base64 encoded format, or an error if the
	// privilege request fails.
	//
	// For details, refer to [github.com/docker/docker/api/types/registry.RequestAuthConfig].
	PrivilegeFunc         func(context.Context) (string, error)
	AcceptPermissionsFunc func(context.Context, PluginPrivileges) (bool, error)
	Args                  []string
}
```

---

### PluginInterfaceType

PluginInterfaceType plugin interface type
swagger:model PluginInterfaceType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_interface_type.go#L8)  
**Added in:** v1.13.0

```go
type PluginInterfaceType struct {

	// capability
	// Required: true
	Capability string `json:"Capability"`

	// prefix
	// Required: true
	Prefix string `json:"Prefix"`

	// version
	// Required: true
	Version string `json:"Version"`
}
```

#### Methods

##### PluginInterfaceType.MarshalJSON

MarshalJSON implements json.Marshaler for PluginInterfaceType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_responses.go#L39)  
**Added in:** v1.13.0

```go
func (t *PluginInterfaceType) MarshalJSON() ([]byte, error)
```

##### PluginInterfaceType.String

String implements fmt.Stringer for PluginInterfaceType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_responses.go#L44)  
**Added in:** v1.13.0

```go
func (t PluginInterfaceType) String() string
```

##### PluginInterfaceType.UnmarshalJSON

UnmarshalJSON implements json.Unmarshaler for PluginInterfaceType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_responses.go#L13)  
**Added in:** v1.13.0

```go
func (t *PluginInterfaceType) UnmarshalJSON(p []byte) error
```

---

### PluginMount

PluginMount plugin mount
swagger:model PluginMount

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_mount.go#L8)  
**Added in:** v1.13.0

```go
type PluginMount struct {

	// description
	// Required: true
	Description string `json:"Description"`

	// destination
	// Required: true
	Destination string `json:"Destination"`

	// name
	// Required: true
	Name string `json:"Name"`

	// options
	// Required: true
	Options []string `json:"Options"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`

	// source
	// Required: true
	Source *string `json:"Source"`

	// type
	// Required: true
	Type string `json:"Type"`
}
```

---

### PluginPrivilege

PluginPrivilege describes a permission the user has to accept
upon installing a plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_responses.go#L50)  
**Added in:** v1.13.0

```go
type PluginPrivilege struct {
	Name        string
	Description string
	Value       []string
}
```

---

### PluginPrivileges

PluginPrivileges is a list of PluginPrivilege

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_responses.go#L57)  
**Added in:** v1.13.0

```go
type PluginPrivileges []PluginPrivilege
```

#### Methods

##### PluginPrivileges.Len

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_responses.go#L59)  

```go
func (s PluginPrivileges) Len() int
```

##### PluginPrivileges.Less

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_responses.go#L63)  

```go
func (s PluginPrivileges) Less(i, j int) bool
```

##### PluginPrivileges.Swap

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_responses.go#L67)  

```go
func (s PluginPrivileges) Swap(i, j int)
```

---

### PluginRemoveOptions

PluginRemoveOptions holds parameters to remove plugins.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/client.go#L50)  
**Added in:** v1.13.0

```go
type PluginRemoveOptions struct {
	Force bool
}
```

---

### PluginSettings

PluginSettings Settings that can be modified by users.
swagger:model PluginSettings

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin.go#L186)  
**Added in:** v1.13.0

```go
type PluginSettings struct {

	// args
	// Required: true
	Args []string `json:"Args"`

	// devices
	// Required: true
	Devices []PluginDevice `json:"Devices"`

	// env
	// Required: true
	Env []string `json:"Env"`

	// mounts
	// Required: true
	Mounts []PluginMount `json:"Mounts"`
}
```

---

### PluginsListResponse

PluginsListResponse contains the response for the Engine API

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugin_responses.go#L10)  
**Added in:** v1.13.0

```go
type PluginsListResponse []*Plugin
```

---

### Port

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L92)  
**Added in:** v1.7.0

```go
type Port = container.Port
```

---

### PushResult

PushResult contains the tag, manifest digest, and manifest size from the
push. It's used to signal this information to the trust code in the client
so it can sign the manifest if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types.go#L99)  

```go
type PushResult struct {
	Tag    string
	Digest string
	Size   int
}
```

---

### RequestPrivilegeFunc

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L117)  
**Added in:** v1.13.0

```go
type RequestPrivilegeFunc func(context.Context) (string, error)
```

---

### RootFS

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L103)  
**Added in:** v1.13.0

```go
type RootFS = image.RootFS
```

---

### SecretCreateResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L123)  
**Added in:** v1.13.0

```go
type SecretCreateResponse = swarm.SecretCreateResponse
```

---

### SecretListOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L128)  
**Added in:** v1.13.0

```go
type SecretListOptions = swarm.SecretListOptions
```

---

### ServiceCreateOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L159)  
**Added in:** v1.13.0

```go
type ServiceCreateOptions = swarm.ServiceCreateOptions
```

---

### ServiceInspectOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L180)  

```go
type ServiceInspectOptions = swarm.ServiceInspectOptions
```

---

### ServiceListOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L174)  
**Added in:** v1.13.0

```go
type ServiceListOptions = swarm.ServiceListOptions
```

---

### ServiceUpdateOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L164)  
**Added in:** v1.13.0

```go
type ServiceUpdateOptions = swarm.ServiceUpdateOptions
```

---

### SummaryNetworkSettings

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L63)  
**Added in:** v1.13.0

```go
type SummaryNetworkSettings = container.NetworkSettingsSummary
```

---

### SwarmUnlockKeyResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L186)  
**Added in:** v1.13.0

```go
type SwarmUnlockKeyResponse = swarm.UnlockKeyResponse
```

---

### TaskListOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types_deprecated.go#L154)  
**Added in:** v1.13.0

```go
type TaskListOptions = swarm.TaskListOptions
```

---

### Version

Version contains response of Engine API:
GET "/version"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/types.go#L46)  
**Added in:** v1.7.0

```go
type Version struct {
	Platform   struct{ Name string } `json:",omitempty"`
	Components []ComponentVersion    `json:",omitempty"`

	Version       string
	APIVersion    string `json:"ApiVersion"`
	MinAPIVersion string `json:"MinAPIVersion,omitempty"`
	GitCommit     string
	GoVersion     string
	Os            string
	Arch          string
	KernelVersion string `json:",omitempty"`
	Experimental  bool   `json:",omitempty"`
	BuildTime     string `json:",omitempty"`
}
```

---


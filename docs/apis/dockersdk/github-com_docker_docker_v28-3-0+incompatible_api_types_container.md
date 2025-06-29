# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/container

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:54 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/config.go#L15)

```go
const MinimumDuration = 1 * time.Millisecond
```

## Variables

This section is empty.

## Functions

### ValidateContainerState

ValidateContainerState checks if the provided string is a valid
container ContainerState.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/state.go#L29)  

```go
func ValidateContainerState(s ContainerState) error
```

---

### ValidateHealthStatus

ValidateHealthStatus checks if the provided string is a valid
container HealthStatus.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/health.go#L43)  

```go
func ValidateHealthStatus(s HealthStatus) error
```

---

### ValidateRestartPolicy

ValidateRestartPolicy validates the given RestartPolicy.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L321)  

```go
func ValidateRestartPolicy(policy RestartPolicy) error
```

---

## Types

### AttachOptions

AttachOptions holds parameters to attach to a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/options.go#L14)  

```go
type AttachOptions struct {
	Stream     bool
	Stdin      bool
	Stdout     bool
	Stderr     bool
	DetachKeys string
	Logs       bool
}
```

---

### BlkioStatEntry

BlkioStatEntry is one small entity to store a piece of Blkio stats
Not used on Windows.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L84)  

```go
type BlkioStatEntry struct {
	Major uint64 `json:"major"`
	Minor uint64 `json:"minor"`
	Op    string `json:"op"`
	Value uint64 `json:"value"`
}
```

---

### BlkioStats

BlkioStats stores All IO service stats for data read and write.
This is a Linux specific structure as the differences between expressing
block I/O on Windows and Linux are sufficiently significant to make
little sense attempting to morph into a combined structure.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L95)  

```go
type BlkioStats struct {
	// number of bytes transferred to and from the block device
	IoServiceBytesRecursive []BlkioStatEntry `json:"io_service_bytes_recursive"`
	IoServicedRecursive     []BlkioStatEntry `json:"io_serviced_recursive"`
	IoQueuedRecursive       []BlkioStatEntry `json:"io_queue_recursive"`
	IoServiceTimeRecursive  []BlkioStatEntry `json:"io_service_time_recursive"`
	IoWaitTimeRecursive     []BlkioStatEntry `json:"io_wait_time_recursive"`
	IoMergedRecursive       []BlkioStatEntry `json:"io_merged_recursive"`
	IoTimeRecursive         []BlkioStatEntry `json:"io_time_recursive"`
	SectorsRecursive        []BlkioStatEntry `json:"sectors_recursive"`
}
```

---

### CPUStats

CPUStats aggregates and wraps all CPU related info of container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L41)  

```go
type CPUStats struct {
	// CPU Usage. Linux and Windows.
	CPUUsage CPUUsage `json:"cpu_usage"`

	// System Usage. Linux only.
	SystemUsage uint64 `json:"system_cpu_usage,omitempty"`

	// Online CPUs. Linux only.
	OnlineCPUs uint32 `json:"online_cpus,omitempty"`

	// Throttling Data. Linux only.
	ThrottlingData ThrottlingData `json:"throttling_data,omitempty"`
}
```

---

### CPUUsage

CPUUsage stores All CPU stats aggregated since container inception.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L17)  

```go
type CPUUsage struct {
	// Total CPU time consumed.
	// Units: nanoseconds (Linux)
	// Units: 100's of nanoseconds (Windows)
	TotalUsage uint64 `json:"total_usage"`

	// Total CPU time consumed per core (Linux). Not used on Windows.
	// Units: nanoseconds.
	PercpuUsage []uint64 `json:"percpu_usage,omitempty"`

	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
	UsageInKernelmode uint64 `json:"usage_in_kernelmode"`

	// Time spent by tasks of the cgroup in user mode (Linux).
	// Time spent by all container processes in user mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers
	UsageInUsermode uint64 `json:"usage_in_usermode"`
}
```

---

### CgroupSpec

CgroupSpec represents the cgroup to use for the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L190)  

```go
type CgroupSpec string
```

#### Methods

##### CgroupSpec.Container

Container returns the ID or name of the container whose cgroup will be used.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L205)  

```go
func (c CgroupSpec) Container() (idOrName string)
```

##### CgroupSpec.IsContainer

IsContainer indicates whether the container is using another container cgroup

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L193)  

```go
func (c CgroupSpec) IsContainer() bool
```

##### CgroupSpec.Valid

Valid indicates whether the cgroup spec is valid.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L199)  

```go
func (c CgroupSpec) Valid() bool
```

---

### CgroupnsMode

CgroupnsMode represents the cgroup namespace mode of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L17)  

```go
type CgroupnsMode string
```

#### Methods

##### CgroupnsMode.IsEmpty

IsEmpty indicates whether the container cgroup namespace mode is unset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L37)  

```go
func (c CgroupnsMode) IsEmpty() bool
```

##### CgroupnsMode.IsHost

IsHost indicates whether the container shares the host's cgroup namespace

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L32)  

```go
func (c CgroupnsMode) IsHost() bool
```

##### CgroupnsMode.IsPrivate

IsPrivate indicates whether the container uses its own private cgroup namespace

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L27)  

```go
func (c CgroupnsMode) IsPrivate() bool
```

##### CgroupnsMode.Valid

Valid indicates whether the cgroup namespace mode is valid

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L42)  

```go
func (c CgroupnsMode) Valid() bool
```

---

### ChangeType

ChangeType Kind of change

Can be one of:

- `0`: Modified ("C")
- `1`: Added ("A")
- `2`: Deleted ("D")

swagger:model ChangeType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/change_type.go#L15)  

```go
type ChangeType uint8
```

#### Methods

##### ChangeType.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/change_types.go#L12)  

```go
func (ct ChangeType) String() string
```

---

### CommitOptions

CommitOptions holds parameters to commit changes into a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/options.go#L24)  

```go
type CommitOptions struct {
	Reference string
	Comment   string
	Author    string
	Changes   []string
	Pause     bool
	Config    *Config
}
```

---

### CommitResponse

CommitResponse response for the commit API call, containing the ID of the
image that was produced.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/commit.go#L7)  

```go
type CommitResponse = common.IDResponse
```

---

### Config

Config contains the configuration data about a container.
It should hold only portable information about the container.
Here, "portable" means "independent from the host we are running on".
Non-portable information *should* appear in HostConfig.
All fields added to this struct must be marked `omitempty` to keep getting
predictable hashes from the old `v1Compatibility` configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/config.go#L44)  

```go
type Config struct {
	Hostname        string              // Hostname
	Domainname      string              // Domainname
	User            string              // User that will run the command(s) inside the container, also support user:group
	AttachStdin     bool                // Attach the standard input, makes possible user interaction
	AttachStdout    bool                // Attach the standard output
	AttachStderr    bool                // Attach the standard error
	ExposedPorts    nat.PortSet         `json:",omitempty"` // List of exposed ports
	Tty             bool                // Attach standard streams to a tty, including stdin if it is not closed.
	OpenStdin       bool                // Open stdin
	StdinOnce       bool                // If true, close stdin after the 1 attached client disconnects.
	Env             []string            // List of environment variable to set in the container
	Cmd             strslice.StrSlice   // Command to run when starting the container
	Healthcheck     *HealthConfig       `json:",omitempty"` // Healthcheck describes how to check the container is healthy
	ArgsEscaped     bool                `json:",omitempty"` // True if command is already escaped (meaning treat as a command line) (Windows specific).
	Image           string              // Name of the image as it was passed by the operator (e.g. could be symbolic)
	Volumes         map[string]struct{} // List of volumes (mounts) used for the container
	WorkingDir      string              // Current directory (PWD) in the command will be launched
	Entrypoint      strslice.StrSlice   // Entrypoint to run when starting the container
	NetworkDisabled bool                `json:",omitempty"` // Is network disabled
	// Mac Address of the container.
	//
	// Deprecated: this field is deprecated since API v1.44. Use EndpointSettings.MacAddress instead.
	MacAddress  string            `json:",omitempty"`
	OnBuild     []string          // ONBUILD metadata that were defined on the image Dockerfile
	Labels      map[string]string // List of labels set to this container
	StopSignal  string            `json:",omitempty"` // Signal to stop a container
	StopTimeout *int              `json:",omitempty"` // Timeout (in seconds) to stop a container
	Shell       strslice.StrSlice `json:",omitempty"` // Shell for shell-form of RUN, CMD, ENTRYPOINT
}
```

---

### ContainerJSONBase

ContainerJSONBase contains response of Engine API GET "/containers/{name:.*}/json"
for API version 1.18 and older.

TODO(thaJeztah): combine ContainerJSONBase and InspectResponse into a single struct.
The split between ContainerJSONBase (ContainerJSONBase) and InspectResponse (InspectResponse)
was done in commit 6deaa58ba5f051039643cedceee97c8695e2af74 (https://github.com/moby/moby/pull/13675).
ContainerJSONBase contained all fields for API < 1.19, and InspectResponse
held fields that were added in API 1.19 and up. Given that the minimum
supported API version is now 1.24, we no longer use the separate type.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L154)  

```go
type ContainerJSONBase struct {
	ID              string `json:"Id"`
	Created         string
	Path            string
	Args            []string
	State           *State
	Image           string
	ResolvConfPath  string
	HostnamePath    string
	HostsPath       string
	LogPath         string
	Name            string
	RestartCount    int
	Driver          string
	Platform        string
	MountLabel      string
	ProcessLabel    string
	AppArmorProfile string
	ExecIDs         []string
	HostConfig      *HostConfig
	GraphDriver     storage.DriverData
	SizeRw          *int64 `json:",omitempty"`
	SizeRootFs      *int64 `json:",omitempty"`
}
```

---

### ContainerState

ContainerState is a string representation of the container's current state.

It currently is an alias for string, but may become a distinct type in the future.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/state.go#L11)  

```go
type ContainerState = string
```

---

### ContainerTopOKBody

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L21)  

```go
type ContainerTopOKBody = TopResponse
```

---

### ContainerUpdateOKBody

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L16)  

```go
type ContainerUpdateOKBody = UpdateResponse
```

---

### CopyToContainerOptions

CopyToContainerOptions holds information
about files to copy into a container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L43)  

```go
type CopyToContainerOptions struct {
	AllowOverwriteDirWithFile bool
	CopyUIDGID                bool
}
```

---

### CreateRequest

CreateRequest is the request message sent to the server for container
create calls. It is a config wrapper that holds the container Config
(portable) and the corresponding HostConfig (non-portable) and
network.NetworkingConfig.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/create_request.go#L9)  

```go
type CreateRequest struct {
	*Config
	HostConfig       *HostConfig               `json:"HostConfig,omitempty"`
	NetworkingConfig *network.NetworkingConfig `json:"NetworkingConfig,omitempty"`
}
```

---

### CreateResponse

CreateResponse ContainerCreateResponse

OK response to ContainerCreate operation
swagger:model CreateResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/create_response.go#L10)  

```go
type CreateResponse struct {

	// The ID of the created container
	// Required: true
	ID string `json:"Id"`

	// Warnings encountered when creating the container
	// Required: true
	Warnings []string `json:"Warnings"`
}
```

---

### DefaultNetworkSettings

DefaultNetworkSettings holds network information
during the 2 release deprecation period.
It will be removed in Docker 1.11.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/network_settings.go#L41)  

```go
type DefaultNetworkSettings struct {
	EndpointID          string // EndpointID uniquely represents a service endpoint in a Sandbox
	Gateway             string // Gateway holds the gateway address for the network
	GlobalIPv6Address   string // GlobalIPv6Address holds network's global IPv6 address
	GlobalIPv6PrefixLen int    // GlobalIPv6PrefixLen represents mask length of network's global IPv6 address
	IPAddress           string // IPAddress holds the IPv4 address for the network
	IPPrefixLen         int    // IPPrefixLen represents mask length of network's IPv4 address
	IPv6Gateway         string // IPv6Gateway holds gateway address specific for IPv6
	MacAddress          string // MacAddress holds the MAC address for the network
}
```

---

### DeviceMapping

DeviceMapping represents the device mapping between the host and the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L269)  

```go
type DeviceMapping struct {
	PathOnHost        string
	PathInContainer   string
	CgroupPermissions string
}
```

---

### DeviceRequest

DeviceRequest represents a request for devices from a device driver.
Used by GPU device drivers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L260)  

```go
type DeviceRequest struct {
	Driver       string            // Name of device driver
	Count        int               // Number of devices to request (-1 = All)
	DeviceIDs    []string          // List of device IDs as recognizable by the device driver
	Capabilities [][]string        // An OR list of AND lists of device capabilities (e.g. "gpu")
	Options      map[string]string // Options to pass onto the device driver
}
```

---

### DiskUsage

DiskUsage contains disk usage for containers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/disk_usage.go#L4)  

```go
type DiskUsage struct {
	TotalSize   int64
	Reclaimable int64
	Items       []*Summary
}
```

---

### ExecAttachOptions

ExecAttachOptions is a temp struct used by execAttach.

TODO(thaJeztah): make this a separate type; ContainerExecAttach does not use the Detach option, and cannot run detached.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/exec.go#L44)  

```go
type ExecAttachOptions = ExecStartOptions
```

---

### ExecCreateResponse

ExecCreateResponse is the response for a successful exec-create request.
It holds the ID of the exec that was created.

TODO(thaJeztah): make this a distinct type.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/exec.go#L9)  

```go
type ExecCreateResponse = common.IDResponse
```

---

### ExecInspect

ExecInspect holds information returned by exec inspect.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/exec.go#L47)  

```go
type ExecInspect struct {
	ExecID      string `json:"ID"`
	ContainerID string
	Running     bool
	ExitCode    int
	Pid         int
}
```

---

### ExecOptions

ExecOptions is a small subset of the Config struct that holds the configuration
for the exec feature of docker.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/exec.go#L13)  

```go
type ExecOptions struct {
	User         string   // User that will run the command
	Privileged   bool     // Is the container in privileged mode
	Tty          bool     // Attach standard streams to a tty.
	ConsoleSize  *[2]uint `json:",omitempty"` // Initial console size [height, width]
	AttachStdin  bool     // Attach the standard input, makes possible user interaction
	AttachStderr bool     // Attach the standard error
	AttachStdout bool     // Attach the standard output
	DetachKeys   string   // Escape keys for detach
	Env          []string // Environment variables
	WorkingDir   string   // Working directory
	Cmd          []string // Execution commands and args

	// Deprecated: the Detach field is not used, and will be removed in a future release.
	Detach bool
}
```

---

### ExecStartOptions

ExecStartOptions is a temp struct used by execStart
Config fields is part of ExecConfig in runconfig package

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/exec.go#L32)  

```go
type ExecStartOptions struct {
	// ExecStart will first check if it's detached
	Detach bool
	// Check if there's a tty
	Tty bool
	// Terminal size [height, width], unused if Tty == false
	ConsoleSize *[2]uint `json:",omitempty"`
}
```

---

### FilesystemChange

FilesystemChange Change in the container's filesystem.

swagger:model FilesystemChange

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/filesystem_change.go#L9)  

```go
type FilesystemChange struct {

	// kind
	// Required: true
	Kind ChangeType `json:"Kind"`

	// Path to file or directory that has changed.
	//
	// Required: true
	Path string `json:"Path"`
}
```

---

### Health

Health stores information about the container's healthcheck results

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/health.go#L23)  

```go
type Health struct {
	Status        HealthStatus         // Status is one of [Starting], [Healthy] or [Unhealthy].
	FailingStreak int                  // FailingStreak is the number of consecutive failures
	Log           []*HealthcheckResult // Log contains the last few results (oldest first)
}
```

---

### HealthConfig

HealthConfig holds configuration settings for the HEALTHCHECK feature.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/config.go#L36)  

```go
type HealthConfig = dockerspec.HealthcheckConfig
```

---

### HealthStatus

HealthStatus is a string representation of the container's health.

It currently is an alias for string, but may become a distinct type in future.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/health.go#L12)  

```go
type HealthStatus = string
```

---

### HealthcheckResult

HealthcheckResult stores information about a single run of a healthcheck probe

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/health.go#L30)  

```go
type HealthcheckResult struct {
	Start    time.Time // Start is the time this check started
	End      time.Time // End is the time this check ended
	ExitCode int       // ExitCode meanings: 0=healthy, 1=unhealthy, 2=reserved (considered unhealthy), else=error running probe
	Output   string    // Output from last check
}
```

---

### HostConfig

HostConfig the non-portable Config structure of a container.
Here, "non-portable" means "dependent of the host we are running on".
Portable information *should* appear in Config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L424)  

```go
type HostConfig struct {
	// Applicable to all platforms
	Binds           []string          // List of volume bindings for this container
	ContainerIDFile string            // File (path) where the containerId is written
	LogConfig       LogConfig         // Configuration of the logs for this container
	NetworkMode     NetworkMode       // Network mode to use for the container
	PortBindings    nat.PortMap       // Port mapping between the exposed port (container) and the host
	RestartPolicy   RestartPolicy     // Restart policy to be used for the container
	AutoRemove      bool              // Automatically remove container when it exits
	VolumeDriver    string            // Name of the volume driver used to mount volumes
	VolumesFrom     []string          // List of volumes to take from other container
	ConsoleSize     [2]uint           // Initial console size (height,width)
	Annotations     map[string]string `json:",omitempty"` // Arbitrary non-identifying metadata attached to container and provided to the runtime

	// Applicable to UNIX platforms
	CapAdd          strslice.StrSlice // List of kernel capabilities to add to the container
	CapDrop         strslice.StrSlice // List of kernel capabilities to remove from the container
	CgroupnsMode    CgroupnsMode      // Cgroup namespace mode to use for the container
	DNS             []string          `json:"Dns"`        // List of DNS server to lookup
	DNSOptions      []string          `json:"DnsOptions"` // List of DNSOption to look for
	DNSSearch       []string          `json:"DnsSearch"`  // List of DNSSearch to look for
	ExtraHosts      []string          // List of extra hosts
	GroupAdd        []string          // List of additional groups that the container process will run as
	IpcMode         IpcMode           // IPC namespace to use for the container
	Cgroup          CgroupSpec        // Cgroup to use for the container
	Links           []string          // List of links (in the name:alias form)
	OomScoreAdj     int               // Container preference for OOM-killing
	PidMode         PidMode           // PID namespace to use for the container
	Privileged      bool              // Is the container in privileged mode
	PublishAllPorts bool              // Should docker publish all exposed port for the container
	ReadonlyRootfs  bool              // Is the container root filesystem in read-only
	SecurityOpt     []string          // List of string values to customize labels for MLS systems, such as SELinux.
	StorageOpt      map[string]string `json:",omitempty"` // Storage driver options per container.
	Tmpfs           map[string]string `json:",omitempty"` // List of tmpfs (mounts) used for the container
	UTSMode         UTSMode           // UTS namespace to use for the container
	UsernsMode      UsernsMode        // The user namespace to use for the container
	ShmSize         int64             // Total shm memory usage
	Sysctls         map[string]string `json:",omitempty"` // List of Namespaced sysctls used for the container
	Runtime         string            `json:",omitempty"` // Runtime to use with this container

	// Applicable to Windows
	Isolation Isolation // Isolation technology of the container (e.g. default, hyperv)

	// Contains container's resources (cgroups, ulimits)
	Resources

	// Mounts specs used by the container
	Mounts []mount.Mount `json:",omitempty"`

	// MaskedPaths is the list of paths to be masked inside the container (this overrides the default set of paths)
	MaskedPaths []string

	// ReadonlyPaths is the list of paths to be set as read-only inside the container (this overrides the default set of paths)
	ReadonlyPaths []string

	// Run a custom init inside the container, if null, use the daemon's configured settings
	Init *bool `json:",omitempty"`
}
```

---

### InspectResponse

InspectResponse is the response for the GET "/containers/{name:.*}/json"
endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L181)  

```go
type InspectResponse struct {
	*ContainerJSONBase
	Mounts          []MountPoint
	Config          *Config
	NetworkSettings *NetworkSettings
	// ImageManifestDescriptor is the descriptor of a platform-specific manifest of the image used to create the container.
	ImageManifestDescriptor *ocispec.Descriptor `json:"ImageManifestDescriptor,omitempty"`
}
```

---

### IpcMode

IpcMode represents the container ipc stack.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L79)  

```go
type IpcMode string
```

#### Methods

##### IpcMode.Container

Container returns the name of the container ipc stack is going to be used.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L128)  

```go
func (n IpcMode) Container() (idOrName string)
```

##### IpcMode.IsContainer

IsContainer indicates whether the container uses another container's ipc namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L106)  

```go
func (n IpcMode) IsContainer() bool
```

##### IpcMode.IsEmpty

IsEmpty indicates whether container IpcMode is empty

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L117)  

```go
func (n IpcMode) IsEmpty() bool
```

##### IpcMode.IsHost

IsHost indicates whether the container shares the host's ipc namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L96)  

```go
func (n IpcMode) IsHost() bool
```

##### IpcMode.IsNone

IsNone indicates whether container IpcMode is set to "none".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L112)  

```go
func (n IpcMode) IsNone() bool
```

##### IpcMode.IsPrivate

IsPrivate indicates whether the container uses its own private ipc namespace which can not be shared.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L91)  

```go
func (n IpcMode) IsPrivate() bool
```

##### IpcMode.IsShareable

IsShareable indicates whether the container's ipc namespace can be shared with another container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L101)  

```go
func (n IpcMode) IsShareable() bool
```

##### IpcMode.Valid

Valid indicates whether the ipc mode is valid.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L122)  

```go
func (n IpcMode) Valid() bool
```

---

### Isolation

Isolation represents the isolation technology of a container. The supported
values are platform specific

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L48)  

```go
type Isolation string
```

#### Methods

##### Isolation.IsDefault

IsDefault indicates the default isolation technology of a container. On Linux this
is the native driver. On Windows, this is a Windows Server Container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L60)  

```go
func (i Isolation) IsDefault() bool
```

##### Isolation.IsHyperV

IsHyperV indicates the use of a Hyper-V partition for isolation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L67)  

```go
func (i Isolation) IsHyperV() bool
```

##### Isolation.IsProcess

IsProcess indicates the use of process isolation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L73)  

```go
func (i Isolation) IsProcess() bool
```

##### Isolation.IsValid

IsValid indicates if an isolation technology is valid

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig_unix.go#L8)  

```go
func (i Isolation) IsValid() bool
```

---

### ListOptions

ListOptions holds parameters to list containers with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/options.go#L47)  

```go
type ListOptions struct {
	Size    bool
	All     bool
	Latest  bool
	Since   string
	Before  string
	Limit   int
	Filters filters.Args
}
```

---

### LogConfig

LogConfig represents the logging configuration of the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L359)  

```go
type LogConfig struct {
	Type   string
	Config map[string]string
}
```

---

### LogMode

LogMode is a type to define the available modes for logging
These modes affect how logs are handled when log messages start piling up.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L349)  

```go
type LogMode string
```

---

### LogsOptions

LogsOptions holds parameters to filter logs with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/options.go#L58)  

```go
type LogsOptions struct {
	ShowStdout bool
	ShowStderr bool
	Since      string
	Until      string
	Timestamps bool
	Follow     bool
	Tail       string
	Details    bool
}
```

---

### MemoryStats

MemoryStats aggregates all memory stats since container inception on Linux.
Windows returns stats for commit and private working set only.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L57)  

```go
type MemoryStats struct {

	// current res_counter usage for memory
	Usage uint64 `json:"usage,omitempty"`
	// maximum usage ever recorded.
	MaxUsage uint64 `json:"max_usage,omitempty"`
	// TODO(vishh): Export these as stronger types.
	// all the stats exported via memory.stat.
	Stats map[string]uint64 `json:"stats,omitempty"`
	// number of times memory usage hits limits.
	Failcnt uint64 `json:"failcnt,omitempty"`
	Limit   uint64 `json:"limit,omitempty"`

	// committed bytes
	Commit uint64 `json:"commitbytes,omitempty"`
	// peak committed bytes
	CommitPeak uint64 `json:"commitpeakbytes,omitempty"`
	// private working set
	PrivateWorkingSet uint64 `json:"privateworkingset,omitempty"`
}
```

---

### MountPoint

MountPoint represents a mount point configuration inside the container.
This is used for reporting the mountpoints in use by a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L62)  

```go
type MountPoint struct {
	// Type is the type of mount, see `Type<foo>` definitions in
	// github.com/docker/docker/api/types/mount.Type
	Type mount.Type `json:",omitempty"`

	// Name is the name reference to the underlying data defined by `Source`
	// e.g., the volume name.
	Name string `json:",omitempty"`

	// Source is the source location of the mount.
	//
	// For volumes, this contains the storage location of the volume (within
	// `/var/lib/docker/volumes/`). For bind-mounts, and `npipe`, this contains
	// the source (host) part of the bind-mount. For `tmpfs` mount points, this
	// field is empty.
	Source string

	// Destination is the path relative to the container root (`/`) where the
	// Source is mounted inside the container.
	Destination string

	// Driver is the volume driver used to create the volume (if it is a volume).
	Driver string `json:",omitempty"`

	// Mode is a comma separated list of options supplied by the user when
	// creating the bind/volume mount.
	//
	// The default is platform-specific (`"z"` on Linux, empty on Windows).
	Mode string

	// RW indicates whether the mount is mounted writable (read-write).
	RW bool

	// Propagation describes how mounts are propagated from the host into the
	// mount point, and vice-versa. Refer to the Linux kernel documentation
	// for details:
	// https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt
	//
	// This field is not used on Windows.
	Propagation mount.Propagation
}
```

---

### NetworkMode

NetworkMode represents the container network stack.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L134)  

```go
type NetworkMode string
```

#### Methods

##### NetworkMode.ConnectedContainer

ConnectedContainer is the id of the container which network this container is connected to.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L158)  

```go
func (n NetworkMode) ConnectedContainer() (idOrName string)
```

##### NetworkMode.IsBridge

IsBridge indicates whether container uses the bridge network stack

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig_unix.go#L13)  

```go
func (n NetworkMode) IsBridge() bool
```

##### NetworkMode.IsContainer

IsContainer indicates whether container uses a container network stack.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L152)  

```go
func (n NetworkMode) IsContainer() bool
```

##### NetworkMode.IsDefault

IsDefault indicates whether container uses the default network stack.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L142)  

```go
func (n NetworkMode) IsDefault() bool
```

##### NetworkMode.IsHost

IsHost indicates whether container uses the host network stack.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig_unix.go#L18)  

```go
func (n NetworkMode) IsHost() bool
```

##### NetworkMode.IsNone

IsNone indicates whether container isn't using a network stack.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L137)  

```go
func (n NetworkMode) IsNone() bool
```

##### NetworkMode.IsPrivate

IsPrivate indicates whether container uses its private network stack.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L147)  

```go
func (n NetworkMode) IsPrivate() bool
```

##### NetworkMode.IsUserDefined

IsUserDefined indicates user-created network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig_unix.go#L23)  

```go
func (n NetworkMode) IsUserDefined() bool
```

##### NetworkMode.NetworkName

NetworkName returns the name of the network stack.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig_unix.go#L28)  

```go
func (n NetworkMode) NetworkName() string
```

##### NetworkMode.UserDefined

UserDefined indicates user-created network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L164)  

```go
func (n NetworkMode) UserDefined() string
```

---

### NetworkSettings

NetworkSettings exposes the network settings in the api

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/network_settings.go#L9)  

```go
type NetworkSettings struct {
	NetworkSettingsBase
	DefaultNetworkSettings
	Networks map[string]*network.EndpointSettings
}
```

---

### NetworkSettingsBase

NetworkSettingsBase holds networking state for a container when inspecting it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/network_settings.go#L16)  

```go
type NetworkSettingsBase struct {
	Bridge     string      // Bridge contains the name of the default bridge interface iff it was set through the daemon --bridge flag.
	SandboxID  string      // SandboxID uniquely represents a container's network stack
	SandboxKey string      // SandboxKey identifies the sandbox
	Ports      nat.PortMap // Ports is a collection of PortBinding indexed by Port

	// HairpinMode specifies if hairpin NAT should be enabled on the virtual interface
	//
	// Deprecated: This field is never set and will be removed in a future release.
	HairpinMode bool
	// LinkLocalIPv6Address is an IPv6 unicast address using the link-local prefix
	//
	// Deprecated: This field is never set and will be removed in a future release.
	LinkLocalIPv6Address string
	// LinkLocalIPv6PrefixLen is the prefix length of an IPv6 unicast address
	//
	// Deprecated: This field is never set and will be removed in a future release.
	LinkLocalIPv6PrefixLen int
	SecondaryIPAddresses   []network.Address // Deprecated: This field is never set and will be removed in a future release.
	SecondaryIPv6Addresses []network.Address // Deprecated: This field is never set and will be removed in a future release.
}
```

---

### NetworkSettingsSummary

NetworkSettingsSummary provides a summary of container's networks
in /containers/json

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/network_settings.go#L54)  

```go
type NetworkSettingsSummary struct {
	Networks map[string]*network.EndpointSettings
}
```

---

### NetworkStats

NetworkStats aggregates the network stats of one container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L116)  

```go
type NetworkStats struct {
	// Bytes received. Windows and Linux.
	RxBytes uint64 `json:"rx_bytes"`
	// Packets received. Windows and Linux.
	RxPackets uint64 `json:"rx_packets"`
	// Received errors. Not used on Windows. Note that we don't `omitempty` this
	// field as it is expected in the >=v1.21 API stats structure.
	RxErrors uint64 `json:"rx_errors"`
	// Incoming packets dropped. Windows and Linux.
	RxDropped uint64 `json:"rx_dropped"`
	// Bytes sent. Windows and Linux.
	TxBytes uint64 `json:"tx_bytes"`
	// Packets sent. Windows and Linux.
	TxPackets uint64 `json:"tx_packets"`
	// Sent errors. Not used on Windows. Note that we don't `omitempty` this
	// field as it is expected in the >=v1.21 API stats structure.
	TxErrors uint64 `json:"tx_errors"`
	// Outgoing packets dropped. Windows and Linux.
	TxDropped uint64 `json:"tx_dropped"`
	// Endpoint ID. Not used on Linux.
	EndpointID string `json:"endpoint_id,omitempty"`
	// Instance ID. Not used on Linux.
	InstanceID string `json:"instance_id,omitempty"`
}
```

---

### PathStat

PathStat is used to encode the header from
GET "/containers/{name:.*}/archive"
"Name" is the file or directory name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L33)  

```go
type PathStat struct {
	Name       string      `json:"name"`
	Size       int64       `json:"size"`
	Mode       os.FileMode `json:"mode"`
	Mtime      time.Time   `json:"mtime"`
	LinkTarget string      `json:"linkTarget"`
}
```

---

### PidMode

PidMode represents the pid namespace of the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L229)  

```go
type PidMode string
```

#### Methods

##### PidMode.Container

Container returns the name of the container whose pid namespace is going to be used.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L253)  

```go
func (n PidMode) Container() (idOrName string)
```

##### PidMode.IsContainer

IsContainer indicates whether the container uses a container's pid namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L242)  

```go
func (n PidMode) IsContainer() bool
```

##### PidMode.IsHost

IsHost indicates whether the container uses the host's pid namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L237)  

```go
func (n PidMode) IsHost() bool
```

##### PidMode.IsPrivate

IsPrivate indicates whether the container uses its own new pid namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L232)  

```go
func (n PidMode) IsPrivate() bool
```

##### PidMode.Valid

Valid indicates whether the pid namespace is valid.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L248)  

```go
func (n PidMode) Valid() bool
```

---

### PidsStats

PidsStats contains the stats of a container's pids

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L142)  

```go
type PidsStats struct {
	// Current is the number of pids in the cgroup
	Current uint64 `json:"current,omitempty"`
	// Limit is the hard limit on the number of pids in the cgroup.
	// A "Limit" of 0 means that there is no limit.
	Limit uint64 `json:"limit,omitempty"`
}
```

---

### Port

Port An open port on a container
swagger:model Port

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/port.go#L8)  

```go
type Port struct {

	// Host IP address that the container's port is mapped to
	IP string `json:"IP,omitempty"`

	// Port on the container
	// Required: true
	PrivatePort uint16 `json:"PrivatePort"`

	// Port exposed on the host
	PublicPort uint16 `json:"PublicPort,omitempty"`

	// type
	// Required: true
	Type string `json:"Type"`
}
```

---

### PruneReport

PruneReport contains the response for Engine API:
POST "/containers/prune"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L25)  

```go
type PruneReport struct {
	ContainersDeleted []string
	SpaceReclaimed    uint64
}
```

---

### RemoveOptions

RemoveOptions holds parameters to remove containers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/options.go#L34)  

```go
type RemoveOptions struct {
	RemoveVolumes bool
	RemoveLinks   bool
	Force         bool
}
```

---

### ResizeOptions

ResizeOptions holds parameters to resize a TTY.
It can be used to resize container TTYs and
exec process TTYs too.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/options.go#L8)  

```go
type ResizeOptions struct {
	Height uint
	Width  uint
}
```

---

### Resources

Resources contains container's resources (cgroups config, ulimits...)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L371)  

```go
type Resources struct {
	// Applicable to all platforms
	CPUShares int64 `json:"CpuShares"` // CPU shares (relative weight vs. other containers)
	Memory    int64 // Memory limit (in bytes)
	NanoCPUs  int64 `json:"NanoCpus"` // CPU quota in units of 10<sup>-9</sup> CPUs.

	// Applicable to UNIX platforms
	CgroupParent         string // Parent cgroup.
	BlkioWeight          uint16 // Block IO weight (relative weight vs. other containers)
	BlkioWeightDevice    []*blkiodev.WeightDevice
	BlkioDeviceReadBps   []*blkiodev.ThrottleDevice
	BlkioDeviceWriteBps  []*blkiodev.ThrottleDevice
	BlkioDeviceReadIOps  []*blkiodev.ThrottleDevice
	BlkioDeviceWriteIOps []*blkiodev.ThrottleDevice
	CPUPeriod            int64           `json:"CpuPeriod"`          // CPU CFS (Completely Fair Scheduler) period
	CPUQuota             int64           `json:"CpuQuota"`           // CPU CFS (Completely Fair Scheduler) quota
	CPURealtimePeriod    int64           `json:"CpuRealtimePeriod"`  // CPU real-time period
	CPURealtimeRuntime   int64           `json:"CpuRealtimeRuntime"` // CPU real-time runtime
	CpusetCpus           string          // CpusetCpus 0-2, 0,1
	CpusetMems           string          // CpusetMems 0-2, 0,1
	Devices              []DeviceMapping // List of devices to map inside the container
	DeviceCgroupRules    []string        // List of rule to be added to the device cgroup
	DeviceRequests       []DeviceRequest // List of device requests for device drivers

	// KernelMemory specifies the kernel memory limit (in bytes) for the container.
	// Deprecated: kernel 5.4 deprecated kmem.limit_in_bytes.
	KernelMemory      int64     `json:",omitempty"`
	KernelMemoryTCP   int64     `json:",omitempty"` // Hard limit for kernel TCP buffer memory (in bytes)
	MemoryReservation int64     // Memory soft limit (in bytes)
	MemorySwap        int64     // Total memory usage (memory + swap); set `-1` to enable unlimited swap
	MemorySwappiness  *int64    // Tuning container memory swappiness behaviour
	OomKillDisable    *bool     // Whether to disable OOM Killer or not
	PidsLimit         *int64    // Setting PIDs limit for a container; Set `0` or `-1` for unlimited, or `null` to not change.
	Ulimits           []*Ulimit // List of ulimits to be set in the container

	// Applicable to Windows
	CPUCount           int64  `json:"CpuCount"`   // CPU count
	CPUPercent         int64  `json:"CpuPercent"` // CPU percent
	IOMaximumIOps      uint64 // Maximum IOps for the container system drive
	IOMaximumBandwidth uint64 // Maximum IO in bytes per second for the container system drive
}
```

---

### RestartPolicy

RestartPolicy represents the restart policies of the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L276)  

```go
type RestartPolicy struct {
	Name              RestartPolicyMode
	MaximumRetryCount int
}
```

#### Methods

##### RestartPolicy.IsAlways

IsAlways indicates whether the container has the "always" restart policy.
This means the container will automatically restart regardless of the exit status.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L298)  

```go
func (rp *RestartPolicy) IsAlways() bool
```

##### RestartPolicy.IsNone

IsNone indicates whether the container has the "no" restart policy.
This means the container will not automatically restart when exiting.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L292)  

```go
func (rp *RestartPolicy) IsNone() bool
```

##### RestartPolicy.IsOnFailure

IsOnFailure indicates whether the container has the "on-failure" restart policy.
This means the container will automatically restart of exiting with a non-zero exit status.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L304)  

```go
func (rp *RestartPolicy) IsOnFailure() bool
```

##### RestartPolicy.IsSame

IsSame compares two RestartPolicy to see if they are the same

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L316)  

```go
func (rp *RestartPolicy) IsSame(tp *RestartPolicy) bool
```

##### RestartPolicy.IsUnlessStopped

IsUnlessStopped indicates whether the container has the
"unless-stopped" restart policy. This means the container will
automatically restart unless user has put it to stopped state.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L311)  

```go
func (rp *RestartPolicy) IsUnlessStopped() bool
```

---

### RestartPolicyMode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L281)  

```go
type RestartPolicyMode string
```

---

### StartOptions

StartOptions holds parameters to start containers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/options.go#L41)  

```go
type StartOptions struct {
	CheckpointID  string
	CheckpointDir string
}
```

---

### State

State stores container's running state
it's part of ContainerJSONBase and returned by "inspect" command

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L106)  

```go
type State struct {
	Status     ContainerState // String representation of the container state. Can be one of "created", "running", "paused", "restarting", "removing", "exited", or "dead"
	Running    bool
	Paused     bool
	Restarting bool
	OOMKilled  bool
	Dead       bool
	Pid        int
	ExitCode   int
	Error      string
	StartedAt  string
	FinishedAt string
	Health     *Health `json:",omitempty"`
}
```

---

### StateStatus

StateStatus is used to return container wait results.
Implements exec.ExitCode interface.
This type is needed as State include a sync.Mutex field which make
copying it unsafe.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/state.go#L42)  

```go
type StateStatus struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewStateStatus

NewStateStatus returns a new StateStatus with the given exit code and error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/state.go#L59)  

```go
func NewStateStatus(exitCode int, err error) StateStatus
```

#### Methods

##### StateStatus.Err

Err returns current error for the state. Returns nil if the container had
exited on its own.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/state.go#L54)  

```go
func (s StateStatus) Err() error
```

##### StateStatus.ExitCode

ExitCode returns current exitcode for the state.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/state.go#L48)  

```go
func (s StateStatus) ExitCode() int
```

---

### Stats

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L153)  

```go
type Stats = StatsResponse
```

---

### StatsResponse

StatsResponse aggregates all types of stats of one container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L156)  

```go
type StatsResponse struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`

	// Common stats
	Read    time.Time `json:"read"`
	PreRead time.Time `json:"preread"`

	// Linux specific stats, not populated on Windows.
	PidsStats  PidsStats  `json:"pids_stats,omitempty"`
	BlkioStats BlkioStats `json:"blkio_stats,omitempty"`

	// Windows specific stats, not populated on Linux.
	NumProcs     uint32       `json:"num_procs"`
	StorageStats StorageStats `json:"storage_stats,omitempty"`

	// Shared stats
	CPUStats    CPUStats                `json:"cpu_stats,omitempty"`
	PreCPUStats CPUStats                `json:"precpu_stats,omitempty"` // "Pre"="Previous"
	MemoryStats MemoryStats             `json:"memory_stats,omitempty"`
	Networks    map[string]NetworkStats `json:"networks,omitempty"`
}
```

---

### StatsResponseReader

StatsResponseReader wraps an io.ReadCloser to read (a stream of) stats
for a container, as produced by the GET "/stats" endpoint.

The OSType field is set to the server's platform to allow
platform-specific handling of the response.

TODO(thaJeztah): remove this wrapper, and make OSType part of StatsResponse.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L55)  

```go
type StatsResponseReader struct {
	Body   io.ReadCloser `json:"body"`
	OSType string        `json:"ostype"`
}
```

---

### StopOptions

StopOptions holds the options to stop or restart a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/config.go#L18)  

```go
type StopOptions struct {
	// Signal (optional) is the signal to send to the container to (gracefully)
	// stop it before forcibly terminating the container with SIGKILL after the
	// timeout expires. If not value is set, the default (SIGTERM) is used.
	Signal string `json:",omitempty"`

	// Timeout (optional) is the timeout (in seconds) to wait for the container
	// to stop gracefully before forcibly terminating it with SIGKILL.
	//
	// - Use nil to use the default timeout (10 seconds).
	// - Use '-1' to wait indefinitely.
	// - Use '0' to not wait for the container to exit gracefully, and
	//   immediately proceeds to forcibly terminating the container.
	// - Other positive values are used as timeout (in seconds).
	Timeout *int `json:",omitempty"`
}
```

---

### StorageStats

StorageStats is the disk I/O stats for read/write on Windows.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L108)  

```go
type StorageStats struct {
	ReadCountNormalized  uint64 `json:"read_count_normalized,omitempty"`
	ReadSizeBytes        uint64 `json:"read_size_bytes,omitempty"`
	WriteCountNormalized uint64 `json:"write_count_normalized,omitempty"`
	WriteSizeBytes       uint64 `json:"write_size_bytes,omitempty"`
}
```

---

### Summary

Summary contains response of Engine API:
GET "/containers/json"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/container.go#L123)  

```go
type Summary struct {
	ID                      string `json:"Id"`
	Names                   []string
	Image                   string
	ImageID                 string
	ImageManifestDescriptor *ocispec.Descriptor `json:"ImageManifestDescriptor,omitempty"`
	Command                 string
	Created                 int64
	Ports                   []Port
	SizeRw                  int64 `json:",omitempty"`
	SizeRootFs              int64 `json:",omitempty"`
	Labels                  map[string]string
	State                   ContainerState
	Status                  string
	HostConfig              struct {
		NetworkMode string            `json:",omitempty"`
		Annotations map[string]string `json:",omitempty"`
	}
	NetworkSettings *NetworkSettingsSummary
	Mounts          []MountPoint
}
```

---

### ThrottlingData

ThrottlingData stores CPU throttling stats of one running container.
Not used on Windows.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/stats.go#L7)  

```go
type ThrottlingData struct {
	// Number of periods with throttling active
	Periods uint64 `json:"periods"`
	// Number of periods when the container hits its throttling limit.
	ThrottledPeriods uint64 `json:"throttled_periods"`
	// Aggregate time the container was throttled for in nanoseconds.
	ThrottledTime uint64 `json:"throttled_time"`
}
```

---

### TopResponse

TopResponse ContainerTopResponse

Container "top" response.
swagger:model TopResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/top_response.go#L10)  

```go
type TopResponse struct {

	// Each process running in the container, where each process
	// is an array of values corresponding to the titles.
	Processes [][]string `json:"Processes"`

	// The ps column titles
	Titles []string `json:"Titles"`
}
```

---

### UTSMode

UTSMode represents the UTS namespace of the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L211)  

```go
type UTSMode string
```

#### Methods

##### UTSMode.IsHost

IsHost indicates whether the container uses the host's UTS namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L219)  

```go
func (n UTSMode) IsHost() bool
```

##### UTSMode.IsPrivate

IsPrivate indicates whether the container uses its private UTS namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L214)  

```go
func (n UTSMode) IsPrivate() bool
```

##### UTSMode.Valid

Valid indicates whether the UTS namespace is valid.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L224)  

```go
func (n UTSMode) Valid() bool
```

---

### Ulimit

Ulimit is an alias for units.Ulimit, which may be moving to a different
location or become a local type. This alias is to help transitioning.

Users are recommended to use this alias instead of using units.Ulimit directly.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L368)  

```go
type Ulimit = units.Ulimit
```

---

### UpdateConfig

UpdateConfig holds the mutable attributes of a Container.
Those attributes can be updated at runtime.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L415)  

```go
type UpdateConfig struct {
	// Contains container's resources (cgroups, ulimits)
	Resources
	RestartPolicy RestartPolicy
}
```

---

### UpdateResponse

UpdateResponse ContainerUpdateResponse

Response for a successful container-update.
swagger:model UpdateResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/update_response.go#L10)  

```go
type UpdateResponse struct {

	// Warnings encountered when updating the container.
	Warnings []string `json:"Warnings"`
}
```

---

### UsernsMode

UsernsMode represents userns mode in the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L172)  

```go
type UsernsMode string
```

#### Methods

##### UsernsMode.IsHost

IsHost indicates whether the container uses the host's userns.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L175)  

```go
func (n UsernsMode) IsHost() bool
```

##### UsernsMode.IsPrivate

IsPrivate indicates whether the container uses the a private userns.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L180)  

```go
func (n UsernsMode) IsPrivate() bool
```

##### UsernsMode.Valid

Valid indicates whether the userns is valid.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/hostconfig.go#L185)  

```go
func (n UsernsMode) Valid() bool
```

---

### WaitCondition

WaitCondition is a type used to specify a container state for which
to wait.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/waitcondition.go#L5)  

```go
type WaitCondition string
```

---

### WaitExitError

WaitExitError container waiting error, if any
swagger:model WaitExitError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/wait_exit_error.go#L8)  

```go
type WaitExitError struct {

	// Details of an error
	Message string `json:"Message,omitempty"`
}
```

---

### WaitResponse

WaitResponse ContainerWaitResponse

OK response to ContainerWait operation
swagger:model WaitResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/container/wait_response.go#L10)  

```go
type WaitResponse struct {

	// error
	Error *WaitExitError `json:"Error,omitempty"`

	// Exit code of the container
	// Required: true
	StatusCode int64 `json:"StatusCode"`
}
```

---


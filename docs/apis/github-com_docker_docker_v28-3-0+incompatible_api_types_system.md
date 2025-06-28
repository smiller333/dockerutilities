# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/system

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:25 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Commit

Commit holds the Git-commit (SHA1) that a binary was built from, as reported
in the version-string of external tools, such as containerd, or runC.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/info.go#L139)  

```go
type Commit struct {
	// ID is the actual commit ID or version of external tool.
	ID string

	// Expected is the commit ID of external tool expected by dockerd as set at build time.
	//
	// Deprecated: this field is no longer used in API v1.49, but kept for backward-compatibility with older API versions.
	Expected string `json:",omitempty"`
}
```

---

### ContainerdInfo

ContainerdInfo holds information about the containerd instance used by the daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/info.go#L88)  

```go
type ContainerdInfo struct {
	// Address is the path to the containerd socket.
	Address string `json:",omitempty"`
	// Namespaces is the containerd namespaces used by the daemon.
	Namespaces ContainerdNamespaces
}
```

---

### ContainerdNamespaces

ContainerdNamespaces reflects the containerd namespaces used by the daemon.

These namespaces can be configured in the daemon configuration, and are
considered to be used exclusively by the daemon,

As these namespaces are considered to be exclusively accessed
by the daemon, it is not recommended to change these values,
or to change them to a value that is used by other systems,
such as cri-containerd.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/info.go#L104)  

```go
type ContainerdNamespaces struct {
	// Containers holds the default containerd namespace used for
	// containers managed by the daemon.
	//
	// The default namespace for containers is "moby", but will be
	// suffixed with the `<uid>.<gid>` of the remapped `root` if
	// user-namespaces are enabled and the containerd image-store
	// is used.
	Containers string

	// Plugins holds the default containerd namespace used for
	// plugins managed by the daemon.
	//
	// The default namespace for plugins is "moby", but will be
	// suffixed with the `<uid>.<gid>` of the remapped `root` if
	// user-namespaces are enabled and the containerd image-store
	// is used.
	Plugins string
}
```

---

### DeviceInfo

DeviceInfo represents a discoverable device from a device driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/info.go#L164)  

```go
type DeviceInfo struct {
	// Source indicates the origin device driver.
	Source string `json:"Source"`
	// ID is the unique identifier for the device.
	// Example: CDI FQDN like "vendor.com/gpu=0", or other driver-specific device ID
	ID string `json:"ID"`
}
```

---

### DiskUsage

DiskUsage contains response of Engine API for API 1.49 and greater:
GET "/system/df"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/disk_usage.go#L12)  

```go
type DiskUsage struct {
	Images     *image.DiskUsage
	Containers *container.DiskUsage
	Volumes    *volume.DiskUsage
	BuildCache *build.CacheDiskUsage
}
```

---

### FirewallInfo

FirewallInfo describes the firewall backend.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/info.go#L156)  

```go
type FirewallInfo struct {
	// Driver is the name of the firewall backend driver.
	Driver string `json:"Driver"`
	// Info is a list of label/value pairs, containing information related to the firewall.
	Info [][2]string `json:"Info,omitempty"`
}
```

---

### Info

Info contains response of Engine API:
GET "/info"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/info.go#L11)  

```go
type Info struct {
	ID                 string
	Containers         int
	ContainersRunning  int
	ContainersPaused   int
	ContainersStopped  int
	Images             int
	Driver             string
	DriverStatus       [][2]string
	SystemStatus       [][2]string `json:",omitempty"` // SystemStatus is only propagated by the Swarm standalone API
	Plugins            PluginsInfo
	MemoryLimit        bool
	SwapLimit          bool
	KernelMemory       bool `json:",omitempty"` // Deprecated: kernel 5.4 deprecated kmem.limit_in_bytes
	KernelMemoryTCP    bool `json:",omitempty"` // KernelMemoryTCP is not supported on cgroups v2.
	CPUCfsPeriod       bool `json:"CpuCfsPeriod"`
	CPUCfsQuota        bool `json:"CpuCfsQuota"`
	CPUShares          bool
	CPUSet             bool
	PidsLimit          bool
	IPv4Forwarding     bool
	Debug              bool
	NFd                int
	OomKillDisable     bool
	NGoroutines        int
	SystemTime         string
	LoggingDriver      string
	CgroupDriver       string
	CgroupVersion      string `json:",omitempty"`
	NEventsListener    int
	KernelVersion      string
	OperatingSystem    string
	OSVersion          string
	OSType             string
	Architecture       string
	IndexServerAddress string
	RegistryConfig     *registry.ServiceConfig
	NCPU               int
	MemTotal           int64
	GenericResources   []swarm.GenericResource
	DockerRootDir      string
	HTTPProxy          string `json:"HttpProxy"`
	HTTPSProxy         string `json:"HttpsProxy"`
	NoProxy            string
	Name               string
	Labels             []string
	ExperimentalBuild  bool
	ServerVersion      string
	Runtimes           map[string]RuntimeWithStatus
	DefaultRuntime     string
	Swarm              swarm.Info
	// LiveRestoreEnabled determines whether containers should be kept
	// running when the daemon is shutdown or upon daemon start if
	// running containers are detected
	LiveRestoreEnabled  bool
	Isolation           container.Isolation
	InitBinary          string
	ContainerdCommit    Commit
	RuncCommit          Commit
	InitCommit          Commit
	SecurityOptions     []string
	ProductLicense      string               `json:",omitempty"`
	DefaultAddressPools []NetworkAddressPool `json:",omitempty"`
	FirewallBackend     *FirewallInfo        `json:"FirewallBackend,omitempty"`
	CDISpecDirs         []string
	DiscoveredDevices   []DeviceInfo `json:",omitempty"`

	Containerd *ContainerdInfo `json:",omitempty"`

	// Warnings contains a slice of warnings that occurred  while collecting
	// system information. These warnings are intended to be informational
	// messages for the user, and are not intended to be parsed / used for
	// other purposes, as they do not have a fixed format.
	Warnings []string
}
```

---

### KeyValue

KeyValue holds a key/value pair.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/security_opts.go#L46)  

```go
type KeyValue struct {
	Key, Value string
}
```

---

### NetworkAddressPool

NetworkAddressPool is a temp struct used by Info struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/info.go#L150)  

```go
type NetworkAddressPool struct {
	Base string
	Size int
}
```

---

### PluginsInfo

PluginsInfo is a temp struct holding Plugins name
registered with docker daemon. It is used by Info struct

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/info.go#L126)  

```go
type PluginsInfo struct {
	// List of Volume plugins registered
	Volume []string
	// List of Network plugins registered
	Network []string
	// List of Authorization plugins registered
	Authorization []string
	// List of Log plugins registered
	Log []string
}
```

---

### Runtime

Runtime describes an OCI runtime

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/runtime.go#L4)  

```go
type Runtime struct {
	Path string   `json:"path,omitempty"`
	Args []string `json:"runtimeArgs,omitempty"`

	Type    string                 `json:"runtimeType,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}
```

---

### RuntimeWithStatus

RuntimeWithStatus extends Runtime to hold [RuntimeStatus].

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/runtime.go#L17)  

```go
type RuntimeWithStatus struct {
	Runtime
	Status map[string]string `json:"status,omitempty"`
}
```

---

### SecurityOpt

SecurityOpt contains the name and options of a security option

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/security_opts.go#L10)  

```go
type SecurityOpt struct {
	Name    string
	Options []KeyValue
}
```

#### Functions

##### DecodeSecurityOptions

DecodeSecurityOptions decodes a security options string slice to a
type-safe SecurityOpt.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/system/security_opts.go#L17)  

```go
func DecodeSecurityOptions(opts []string) ([]SecurityOpt, error)
```

---


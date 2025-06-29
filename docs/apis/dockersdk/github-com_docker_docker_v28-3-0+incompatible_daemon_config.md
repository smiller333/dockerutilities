# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/config

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:57 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L28)

```go
const (
	// DefaultMaxConcurrentDownloads is the default value for
	// maximum number of downloads that
	// may take place at a time.
	DefaultMaxConcurrentDownloads = 3
	// DefaultMaxConcurrentUploads is the default value for
	// maximum number of uploads that
	// may take place at a time.
	DefaultMaxConcurrentUploads = 5
	// DefaultDownloadAttempts is the default value for
	// maximum number of attempts that
	// may take place at a time for each pull when the connection is lost.
	DefaultDownloadAttempts = 5
	// DefaultShmSize is the default value for container's shm size (64 MiB)
	DefaultShmSize int64 = 64 * 1024 * 1024
	// DefaultNetworkMtu is the default value for network MTU
	DefaultNetworkMtu = 1500
	// DisableNetworkBridge is the default value of the option to disable network bridge
	DisableNetworkBridge = "none"
	// DefaultShutdownTimeout is the default shutdown timeout (in seconds) for
	// the daemon for containers to stop when it is shutting down.
	DefaultShutdownTimeout = 15
	// DefaultInitBinary is the name of the default init binary
	DefaultInitBinary = "docker-init"
	// DefaultRuntimeBinary is the default runtime to be used by
	// containerd if none is specified
	DefaultRuntimeBinary = "runc"
	// DefaultContainersNamespace is the name of the default containerd namespace used for users containers.
	DefaultContainersNamespace = "moby"
	// DefaultPluginNamespace is the name of the default containerd namespace used for plugins.
	DefaultPluginNamespace = "plugins.moby"

	// SeccompProfileDefault is the built-in default seccomp profile.
	SeccompProfileDefault = "builtin"
	// SeccompProfileUnconfined is a special profile name for seccomp to use an
	// "unconfined" seccomp profile.
	SeccompProfileUnconfined = "unconfined"
	// LibnetDataPath is the path to libnetwork's data directory, relative to cfg.Root.
	// Windows tolerates the "/".
	LibnetDataPath = "network/files"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L22)

```go
const (
	// DefaultIpcMode is default for container's IpcMode, if not set otherwise
	DefaultIpcMode = container.IPCModePrivate

	// DefaultCgroupNamespaceMode is the default mode for containers cgroup namespace when using cgroups v2.
	DefaultCgroupNamespaceMode = container.CgroupnsModePrivate

	// DefaultCgroupV1NamespaceMode is the default mode for containers cgroup namespace when using cgroups v1.
	DefaultCgroupV1NamespaceMode = container.CgroupnsModeHost

	// StockRuntimeName is the reserved name/alias used to represent the
	// OCI runtime being shipped with the docker daemon package.
	StockRuntimeName = "runc"
)
```

## Variables

This section is empty.

## Functions

### GetConflictFreeLabels

GetConflictFreeLabels validates Labels for conflict
In swarm the duplicates for labels are removed
so we only take same values here, no conflict values
If the key-value is the same we will only take the last label

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L362)  

```go
func GetConflictFreeLabels(labels []string) ([]string, error)
```

---

### MaskCredentials

MaskCredentials masks credentials that are in an URL.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L793)  

```go
func MaskCredentials(rawURL string) string
```

---

### ParseGenericResources

ParseGenericResources parses and validates the specified string as a list of GenericResource

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/opts.go#L10)  

```go
func ParseGenericResources(value []string) ([]swarm.GenericResource, error)
```

---

### Reload

Reload reads the configuration in the host and reloads the daemon and server.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L383)  

```go
func Reload(configFile string, flags *pflag.FlagSet, reload func(*Config)) error
```

---

### Validate

Validate validates some specific configs.
such as config.DNS, config.Labels, config.DNSSearch,
as well as config.MaxConcurrentDownloads, config.MaxConcurrentUploads and config.MaxDownloadAttempts.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L682)  

```go
func Validate(config *Config) error
```

---

### ValidateMinAPIVersion

ValidateMinAPIVersion verifies if the given API version is within the
range supported by the daemon. It is used to validate a custom minimum
API version set through DOCKER_MIN_API_VERSION.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L663)  

```go
func ValidateMinAPIVersion(ver string) error
```

---

## Types

### BridgeConfig

BridgeConfig stores all the parameters for both the bridge driver and the default bridge network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L41)  

```go
type BridgeConfig struct {
	DefaultBridgeConfig

	EnableIPTables           bool   `json:"iptables,omitempty"`
	EnableIP6Tables          bool   `json:"ip6tables,omitempty"`
	EnableIPForward          bool   `json:"ip-forward,omitempty"`
	DisableFilterForwardDrop bool   `json:"ip-forward-no-drop,omitempty"`
	EnableIPMasq             bool   `json:"ip-masq,omitempty"`
	EnableUserlandProxy      bool   `json:"userland-proxy,omitempty"`
	UserlandProxyPath        string `json:"userland-proxy-path,omitempty"`
	AllowDirectRouting       bool   `json:"allow-direct-routing,omitempty"`
}
```

---

### BuilderConfig

BuilderConfig contains config for the builder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L141)  

```go
type BuilderConfig struct {
	GC           BuilderGCConfig       `json:",omitempty"`
	Entitlements BuilderEntitlements   `json:",omitempty"`
	History      *BuilderHistoryConfig `json:",omitempty"`
}
```

---

### BuilderEntitlements

BuilderEntitlements contains settings to enable/disable entitlements

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L135)  

```go
type BuilderEntitlements struct {
	NetworkHost      *bool `json:"network-host,omitempty"`
	SecurityInsecure *bool `json:"security-insecure,omitempty"`
}
```

---

### BuilderGCConfig

BuilderGCConfig contains GC config for a buildkit builder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L86)  

```go
type BuilderGCConfig struct {
	Enabled              *bool           `json:",omitempty"`
	Policy               []BuilderGCRule `json:",omitempty"`
	DefaultReservedSpace string          `json:",omitempty"`
	DefaultMaxUsedSpace  string          `json:",omitempty"`
	DefaultMinFreeSpace  string          `json:",omitempty"`
}
```

#### Methods

##### BuilderGCConfig.IsEnabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L94)  

```go
func (x *BuilderGCConfig) IsEnabled() bool
```

##### BuilderGCConfig.UnmarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L98)  

```go
func (x *BuilderGCConfig) UnmarshalJSON(data []byte) error
```

---

### BuilderGCFilter

BuilderGCFilter contains garbage-collection filter rules for a BuildKit builder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L48)  

```go
type BuilderGCFilter filters.Args
```

#### Methods

##### BuilderGCFilter.MarshalJSON

MarshalJSON returns a JSON byte representation of the BuilderGCFilter

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L51)  

```go
func (x *BuilderGCFilter) MarshalJSON() ([]byte, error)
```

##### BuilderGCFilter.UnmarshalJSON

UnmarshalJSON fills the BuilderGCFilter values structure from JSON input

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L66)  

```go
func (x *BuilderGCFilter) UnmarshalJSON(data []byte) error
```

---

### BuilderGCRule

BuilderGCRule represents a GC rule for buildkit cache

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L13)  

```go
type BuilderGCRule struct {
	All           bool            `json:",omitempty"`
	Filter        BuilderGCFilter `json:",omitempty"`
	ReservedSpace string          `json:",omitempty"`
	MaxUsedSpace  string          `json:",omitempty"`
	MinFreeSpace  string          `json:",omitempty"`
}
```

#### Methods

##### BuilderGCRule.UnmarshalJSON

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L21)  

```go
func (x *BuilderGCRule) UnmarshalJSON(data []byte) error
```

---

### BuilderHistoryConfig

BuilderHistoryConfig contains history config for a buildkit builder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/builder.go#L129)  

```go
type BuilderHistoryConfig struct {
	MaxAge     bkconfig.Duration `json:",omitempty"`
	MaxEntries int64             `json:",omitempty"`
}
```

---

### CommonConfig

CommonConfig defines the configuration of a docker daemon which is
common across platforms.
It includes json tags to deserialize configuration from a file
using the same names that the flags in the command line use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L175)  

```go
type CommonConfig struct {
	AuthorizationPlugins  []string `json:"authorization-plugins,omitempty"` // AuthorizationPlugins holds list of authorization plugins
	AutoRestart           bool     `json:"-"`
	DisableBridge         bool     `json:"-"`
	ExecOptions           []string `json:"exec-opts,omitempty"`
	GraphDriver           string   `json:"storage-driver,omitempty"`
	GraphOptions          []string `json:"storage-opts,omitempty"`
	Labels                []string `json:"labels,omitempty"`
	NetworkDiagnosticPort int      `json:"network-diagnostic-port,omitempty"`
	Pidfile               string   `json:"pidfile,omitempty"`
	RawLogs               bool     `json:"raw-logs,omitempty"`
	Root                  string   `json:"data-root,omitempty"`
	ExecRoot              string   `json:"exec-root,omitempty"`
	SocketGroup           string   `json:"group,omitempty"`
	CorsHeaders           string   `json:"api-cors-header,omitempty"` // Deprecated: CORS headers should not be set on the API. This feature will be removed in the next release. // TODO(thaJeztah): option is used to produce error when used; remove in next release

	// Proxies holds the proxies that are configured for the daemon.
	Proxies `json:"proxies"`

	// LiveRestoreEnabled determines whether we should keep containers
	// alive upon daemon shutdown/start
	LiveRestoreEnabled bool `json:"live-restore,omitempty"`

	// MaxConcurrentDownloads is the maximum number of downloads that
	// may take place at a time for each pull.
	MaxConcurrentDownloads int `json:"max-concurrent-downloads,omitempty"`

	// MaxConcurrentUploads is the maximum number of uploads that
	// may take place at a time for each push.
	MaxConcurrentUploads int `json:"max-concurrent-uploads,omitempty"`

	// MaxDownloadAttempts is the maximum number of attempts that
	// may take place at a time for each push.
	MaxDownloadAttempts int `json:"max-download-attempts,omitempty"`

	// ShutdownTimeout is the timeout value (in seconds) the daemon will wait for the container
	// to stop when daemon is being shutdown
	ShutdownTimeout int `json:"shutdown-timeout,omitempty"`

	Debug     bool             `json:"debug,omitempty"`
	Hosts     []string         `json:"hosts,omitempty"`
	LogLevel  string           `json:"log-level,omitempty"`
	LogFormat log.OutputFormat `json:"log-format,omitempty"`
	TLS       *bool            `json:"tls,omitempty"`
	TLSVerify *bool            `json:"tlsverify,omitempty"`

	// Embedded structs that allow config
	// deserialization without the full struct.
	TLSOptions

	// SwarmDefaultAdvertiseAddr is the default host/IP or network interface
	// to use if a wildcard address is specified in the ListenAddr value
	// given to the /swarm/init endpoint and no advertise address is
	// specified.
	SwarmDefaultAdvertiseAddr string `json:"swarm-default-advertise-addr"`

	// SwarmRaftHeartbeatTick is the number of ticks in time for swarm mode raft quorum heartbeat
	// Typical value is 1
	SwarmRaftHeartbeatTick uint32 `json:"swarm-raft-heartbeat-tick"`

	// SwarmRaftElectionTick is the number of ticks to elapse before followers in the quorum can propose
	// a new round of leader election.  Default, recommended value is at least 10X that of Heartbeat tick.
	// Higher values can make the quorum less sensitive to transient faults in the environment, but this also
	// means it takes longer for the managers to detect a down leader.
	SwarmRaftElectionTick uint32 `json:"swarm-raft-election-tick"`

	MetricsAddress string `json:"metrics-addr"`

	DNSConfig
	LogConfig
	BridgeConfig // BridgeConfig holds bridge network specific configuration.
	NetworkConfig
	registry.ServiceOptions

	// FIXME(vdemeester) This part is not that clear and is mainly dependent on cli flags
	// It should probably be handled outside this package.
	ValuesSet map[string]interface{} `json:"-"`

	Experimental bool `json:"experimental"` // Experimental indicates whether experimental features should be exposed or not

	// Exposed node Generic Resources
	// e.g: ["orange=red", "orange=green", "orange=blue", "apple=3"]
	NodeGenericResources []string `json:"node-generic-resources,omitempty"`

	// ContainerAddr is the address used to connect to containerd if we're
	// not starting it ourselves
	ContainerdAddr string `json:"containerd,omitempty"`

	// CriContainerd determines whether a supervised containerd instance
	// should be configured with the CRI plugin enabled. This allows using
	// Docker's containerd instance directly with a Kubernetes kubelet.
	CriContainerd bool `json:"cri-containerd,omitempty"`

	// Features contains a list of feature key value pairs indicating what features are enabled or disabled.
	// If a certain feature doesn't appear in this list then it's unset (i.e. neither true nor false).
	Features map[string]bool `json:"features,omitempty"`

	Builder BuilderConfig `json:"builder,omitempty"`

	ContainerdNamespace       string `json:"containerd-namespace,omitempty"`
	ContainerdPluginNamespace string `json:"containerd-plugin-namespace,omitempty"`

	DefaultRuntime string `json:"default-runtime,omitempty"`

	// CDISpecDirs is a list of directories in which CDI specifications can be found.
	CDISpecDirs []string `json:"cdi-spec-dirs,omitempty"`

	// The minimum API version provided by the daemon. Defaults to [defaultMinAPIVersion].
	//
	// The DOCKER_MIN_API_VERSION allows overriding the minimum API version within
	// constraints of the minimum and maximum (current) supported API versions.
	//
	// API versions older than [defaultMinAPIVersion] are deprecated and
	// to be removed in a future release. The "DOCKER_MIN_API_VERSION" env
	// var should only be used for exceptional cases, and the MinAPIVersion
	// field is therefore not included in the JSON representation.
	MinAPIVersion string `json:"-"`
}
```

---

### Config

Config defines the configuration of a docker daemon.
It includes json tags to deserialize configuration from a file
using the same names that the flags in the command line uses.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L73)  

```go
type Config struct {
	CommonConfig

	// Fields below here are platform specific.
	Runtimes             map[string]system.Runtime    `json:"runtimes,omitempty"`
	DefaultInitBinary    string                       `json:"default-init,omitempty"`
	CgroupParent         string                       `json:"cgroup-parent,omitempty"`
	EnableSelinuxSupport bool                         `json:"selinux-enabled,omitempty"`
	RemappedRoot         string                       `json:"userns-remap,omitempty"`
	Ulimits              map[string]*container.Ulimit `json:"default-ulimits,omitempty"`
	CPURealtimePeriod    int64                        `json:"cpu-rt-period,omitempty"`
	CPURealtimeRuntime   int64                        `json:"cpu-rt-runtime,omitempty"`
	Init                 bool                         `json:"init,omitempty"`
	InitPath             string                       `json:"init-path,omitempty"`
	SeccompProfile       string                       `json:"seccomp-profile,omitempty"`
	ShmSize              opts.MemBytes                `json:"default-shm-size,omitempty"`
	NoNewPrivileges      bool                         `json:"no-new-privileges,omitempty"`
	IpcMode              string                       `json:"default-ipc-mode,omitempty"`
	CgroupNamespaceMode  string                       `json:"default-cgroupns-mode,omitempty"`
	// ResolvConf is the path to the configuration of the host resolver
	ResolvConf string `json:"resolv-conf,omitempty"`
	Rootless   bool   `json:"rootless,omitempty"`
}
```

#### Functions

##### MergeDaemonConfigurations

MergeDaemonConfigurations reads a configuration file,
loads the file configuration in an isolated structure,
and merges the configuration provided from flags on top
if there are no conflicts.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L440)  

```go
func MergeDaemonConfigurations(flagsConfig *Config, flags *pflag.FlagSet, configFile string) (*Config, error)
```

##### New

New returns a new fully initialized Config struct with default values set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L312)  

```go
func New() (*Config, error)
```

##### Sanitize

Sanitize sanitizes the config for printing. It is currently limited to
masking usernames and passwords from Proxy URLs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L813)  

```go
func Sanitize(cfg Config) Config
```

#### Methods

##### Config.GetExecOpt

GetExecOpt looks up a user-configured exec-opt. It returns a boolean
if found, and an error if the configuration has invalid options set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L349)  

```go
func (conf *Config) GetExecOpt(name string) (val string, found bool, _ error)
```

##### Config.GetExecRoot

GetExecRoot returns the user configured Exec-root

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L98)  

```go
func (conf *Config) GetExecRoot() string
```

##### Config.GetInitPath

GetInitPath returns the configured docker-init path

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L103)  

```go
func (conf *Config) GetInitPath() string
```

##### Config.GetResolvConf

GetResolvConf returns the appropriate resolv.conf
Check setupResolvConf on how this is selected

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L120)  

```go
func (conf *Config) GetResolvConf() string
```

##### Config.IsRootless

IsRootless returns conf.Rootless on Linux but false on Windows

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L140)  

```go
func (conf *Config) IsRootless() bool
```

##### Config.IsSwarmCompatible

IsSwarmCompatible defines if swarm mode can be enabled in this config

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L125)  

```go
func (conf *Config) IsSwarmCompatible() error
```

##### Config.IsValueSet

IsValueSet returns true if a configuration value
was explicitly set in the configuration file.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L303)  

```go
func (conf *Config) IsValueSet(name string) bool
```

##### Config.LookupInitPath

LookupInitPath returns an absolute path to the "docker-init" binary by searching relevant "libexec" directories (per FHS 3.0 & 2.3) followed by PATH

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L114)  

```go
func (conf *Config) LookupInitPath() (string, error)
```

##### Config.ValidatePlatformConfig ⚠️ **DEPRECATED**

ValidatePlatformConfig checks if any platform-specific configuration settings are invalid.

Deprecated: this function was only used internally and is no longer used. Use Validate instead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L135)  

```go
func (conf *Config) ValidatePlatformConfig() error
```

---

### DNSConfig

DNSConfig defines the DNS configurations.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L163)  

```go
type DNSConfig struct {
	DNS            []net.IP     `json:"dns,omitempty"`
	DNSOptions     []string     `json:"dns-opts,omitempty"`
	DNSSearch      []string     `json:"dns-search,omitempty"`
	HostGatewayIP  net.IP       `json:"host-gateway-ip,omitempty"` // Deprecated: this single-IP is migrated to HostGatewayIPs
	HostGatewayIPs []netip.Addr `json:"host-gateway-ips,omitempty"`
}
```

---

### DefaultBridgeConfig

DefaultBridgeConfig stores all the parameters for the default bridge network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config_linux.go#L55)  

```go
type DefaultBridgeConfig struct {

	// Fields below here are platform specific.
	EnableIPv6                  bool   `json:"ipv6,omitempty"`
	FixedCIDRv6                 string `json:"fixed-cidr-v6,omitempty"`
	MTU                         int    `json:"mtu,omitempty"`
	DefaultIP                   net.IP `json:"ip,omitempty"`
	IP                          string `json:"bip,omitempty"`
	IP6                         string `json:"bip6,omitempty"`
	DefaultGatewayIPv4          net.IP `json:"default-gateway,omitempty"`
	DefaultGatewayIPv6          net.IP `json:"default-gateway-v6,omitempty"`
	InterContainerCommunication bool   `json:"icc,omitempty"`
	// contains filtered or unexported fields
}
```

---

### LogConfig

LogConfig represents the default log configuration.
It includes json tags to deserialize configuration from a file
using the same names that the flags in the command line use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L131)  

```go
type LogConfig struct {
	Type   string            `json:"log-driver,omitempty"`
	Config map[string]string `json:"log-opts,omitempty"`
}
```

---

### NetworkConfig

NetworkConfig stores the daemon-wide networking configurations

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L144)  

```go
type NetworkConfig struct {
	// Default address pools for docker networks
	DefaultAddressPools opts.PoolsOpt `json:"default-address-pools,omitempty"`
	// NetworkControlPlaneMTU allows to specify the control plane MTU, this will allow to optimize the network use in some components
	NetworkControlPlaneMTU int `json:"network-control-plane-mtu,omitempty"`
	// Default options for newly created networks
	DefaultNetworkOpts map[string]map[string]string `json:"default-network-opts,omitempty"`
}
```

---

### Proxies

Proxies holds the proxies that are configured for the daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L295)  

```go
type Proxies struct {
	HTTPProxy  string `json:"http-proxy,omitempty"`
	HTTPSProxy string `json:"https-proxy,omitempty"`
	NoProxy    string `json:"no-proxy,omitempty"`
}
```

---

### TLSOptions

TLSOptions defines TLS configuration for the daemon server.
It includes json tags to deserialize configuration from a file
using the same names that the flags in the command line use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/config/config.go#L156)  

```go
type TLSOptions struct {
	CAFile   string `json:"tlscacert,omitempty"`
	CertFile string `json:"tlscert,omitempty"`
	KeyFile  string `json:"tlskey,omitempty"`
}
```

---


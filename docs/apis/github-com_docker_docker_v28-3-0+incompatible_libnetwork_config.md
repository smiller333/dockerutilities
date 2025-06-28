# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/config

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:08:51 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Config

Config encapsulates configurations of various Libnetwork components

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L24)  

```go
type Config struct {
	DataDir string
	// ExecRoot is the base-path for libnetwork external key listeners
	// (created in "<ExecRoot>/libnetwork/<Controller-Short-ID>.sock"),
	// and is passed as "-exec-root: argument for "libnetwork-setkey".
	//
	// It is only used on Linux, but referenced in some "unix" files
	// (linux and freebsd).
	//
	// FIXME(thaJeztah): ExecRoot is only used for Controller.startExternalKeyListener(), but "libnetwork-setkey" is only implemented on Linux.
	ExecRoot       string
	DefaultNetwork string
	DefaultDriver  string
	Labels         []string

	ClusterProvider        cluster.Provider
	NetworkControlPlaneMTU int
	DefaultAddressPool     []*ipamutils.NetworkToSplit
	DatastoreBucket        string
	ActiveSandboxes        map[string]any
	PluginGetter           plugingetter.PluginGetter
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates a new Config and initializes it with the given Options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L48)  

```go
func New(opts ...Option) *Config
```

#### Methods

##### Config.DriverConfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L63)  

```go
func (c *Config) DriverConfig(name string) map[string]any
```

---

### Option

Option is an option setter function type used to pass various configurations
to the controller

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L69)  

```go
type Option func(c *Config)
```

#### Functions

##### OptionActiveSandboxes

OptionActiveSandboxes function returns an option setter for passing the sandboxes
which were active during previous daemon life

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L152)  

```go
func OptionActiveSandboxes(sandboxes map[string]any) Option
```

##### OptionDataDir

OptionDataDir function returns an option setter for data folder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L113)  

```go
func OptionDataDir(dataDir string) Option
```

##### OptionDefaultAddressPoolConfig

OptionDefaultAddressPoolConfig function returns an option setter for default address pool

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L88)  

```go
func OptionDefaultAddressPoolConfig(addressPool []*ipamutils.NetworkToSplit) Option
```

##### OptionDefaultDriver

OptionDefaultDriver function returns an option setter for default driver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L80)  

```go
func OptionDefaultDriver(dd string) Option
```

##### OptionDefaultNetwork

OptionDefaultNetwork function returns an option setter for a default network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L72)  

```go
func OptionDefaultNetwork(dn string) Option
```

##### OptionDriverConfig

OptionDriverConfig returns an option setter for driver configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L95)  

```go
func OptionDriverConfig(networkType string, config map[string]any) Option
```

##### OptionExecRoot

OptionExecRoot function returns an option setter for exec root folder.

On Linux, it sets both the controller's ExecRoot and osl.basePath, whereas
on FreeBSD, it only sets the controller's ExecRoot. It is a no-op on other
platforms.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L124)  

```go
func OptionExecRoot(execRoot string) Option
```

##### OptionLabels

OptionLabels function returns an option setter for labels

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L102)  

```go
func OptionLabels(labels []string) Option
```

##### OptionNetworkControlPlaneMTU

OptionNetworkControlPlaneMTU function returns an option setter for control plane MTU

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L136)  

```go
func OptionNetworkControlPlaneMTU(exp int) Option
```

##### OptionPluginGetter

OptionPluginGetter returns a plugingetter for remote drivers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/config/config.go#L129)  

```go
func OptionPluginGetter(pg plugingetter.PluginGetter) Option
```

---


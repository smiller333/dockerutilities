# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/sysinfo

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:16 UTC

## Overview

Package sysinfo stores information about which features a kernel supports.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Opt

Opt for New().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/sysinfo/sysinfo.go#L5)  

```go
type Opt func(info *SysInfo)
```

#### Functions

##### WithCgroup2GroupPath

WithCgroup2GroupPath specifies the cgroup v2 group path to inspect availability
of the controllers.

WithCgroup2GroupPath is expected to be used for rootless mode with systemd driver.

e.g. g = "/user.slice/user-1000.slice/user@1000.service"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/sysinfo/sysinfo_linux.go#L80)  

```go
func WithCgroup2GroupPath(g string) Opt
```

---

### SysInfo

SysInfo stores information about which features a kernel supports.
TODO Windows: Factor out platform specific capabilities.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/sysinfo/sysinfo.go#L9)  

```go
type SysInfo struct {
	// Whether the kernel supports AppArmor or not
	AppArmor bool
	// Whether the kernel supports Seccomp or not
	Seccomp bool

	// Whether the kernel supports cgroup namespaces or not
	CgroupNamespaces bool

	// Whether IPv4 forwarding is supported or not, if this was disabled, networking will not work
	IPv4ForwardingDisabled bool

	// Whether the cgroup has the mountpoint of "devices" or not
	CgroupDevicesEnabled bool

	// Whether the cgroup is in unified mode (v2).
	CgroupUnified bool

	// Warnings contains a slice of warnings that occurred  while collecting
	// system information. These warnings are intended to be informational
	// messages for the user, and can either be logged or returned to the
	// client; they are not intended to be parsed / used for other purposes,
	// and do not have a fixed format.
	Warnings []string
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New returns a new SysInfo, using the filesystem to detect which features
the kernel supports.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/sysinfo/sysinfo_linux.go#L90)  

```go
func New(options ...Opt) *SysInfo
```

#### Methods

##### SysInfo.IsCpusetCpusAvailable

IsCpusetCpusAvailable returns `true` if the provided string set is contained
in cgroup's cpuset.cpus set, `false` otherwise.
If error is not nil a parsing error occurred.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/sysinfo/sysinfo.go#L138)  
**Added in:** v1.9.0

```go
func (c SysInfo) IsCpusetCpusAvailable(requested string) (bool, error)
```

##### SysInfo.IsCpusetMemsAvailable

IsCpusetMemsAvailable returns `true` if the provided string set is contained
in cgroup's cpuset.mems set, `false` otherwise.
If error is not nil a parsing error occurred.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/sysinfo/sysinfo.go#L145)  
**Added in:** v1.9.0

```go
func (c SysInfo) IsCpusetMemsAvailable(requested string) (bool, error)
```

---


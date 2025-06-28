# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/oci

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:12 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AppendDevicePermissionsFromCgroupRules

AppendDevicePermissionsFromCgroupRules takes rules for the devices cgroup to append to the default set

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/oci.go#L35)  

```go
func AppendDevicePermissionsFromCgroupRules(devPermissions []specs.LinuxDeviceCgroup, rules []string) ([]specs.LinuxDeviceCgroup, error)
```

---

### DefaultLinuxSpec

DefaultLinuxSpec create a default spec for running Linux containers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/defaults.go#L56)  

```go
func DefaultLinuxSpec() specs.Spec
```

---

### DefaultPathEnv

DefaultPathEnv is unix style list of directories to search for
executables. Each directory is separated from the next by a colon
':' character .
For Windows containers, an empty string is returned as the default
path will be set by the container, and Docker has no context of what the
default path should be.

TODO(thaJeztah) align Windows default with BuildKit; see https://github.com/moby/buildkit/pull/1747
TODO(thaJeztah) use defaults from containerd (but align it with BuildKit; see https://github.com/moby/buildkit/pull/1747)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/defaults.go#L30)  

```go
func DefaultPathEnv(os string) string
```

---

### DefaultSpec

DefaultSpec returns the default spec used by docker for the current Platform

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/defaults.go#L38)  

```go
func DefaultSpec() specs.Spec
```

---

### DefaultWindowsSpec

DefaultWindowsSpec create a default spec for running Windows containers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/defaults.go#L46)  

```go
func DefaultWindowsSpec() specs.Spec
```

---

### DevicesFromPath

DevicesFromPath computes a list of devices and device permissions from paths (pathOnHost and pathInContainer) and cgroup permissions.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/devices_linux.go#L25)  
**Added in:** v1.13.0

```go
func DevicesFromPath(pathOnHost, pathInContainer, cgroupPermissions string) (devs []specs.LinuxDevice, devPermissions []specs.LinuxDeviceCgroup, _ error)
```

---

### NamespacePath

NamespacePath returns the configured Path of the first namespace in
s.Linux.Namespaces of type nsType.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/namespaces.go#L20)  

```go
func NamespacePath(s *specs.Spec, nsType specs.LinuxNamespaceType) (path string, ok bool)
```

---

### RemoveNamespace

RemoveNamespace removes the `nsType` namespace from OCI spec `s`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/namespaces.go#L6)  
**Added in:** v1.13.0

```go
func RemoveNamespace(s *specs.Spec, nsType specs.LinuxNamespaceType)
```

---

### SetCapabilities

SetCapabilities sets the provided capabilities on the spec
All capabilities are added if privileged is true.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/oci/oci.go#L22)  

```go
func SetCapabilities(s *specs.Spec, caplist []string) error
```

---

## Types

This section is empty.


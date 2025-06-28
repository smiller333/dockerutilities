# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/profiles/seccomp

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:59 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GetDefaultProfile

GetDefaultProfile returns the default seccomp profile.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp_linux.go#L15)  

```go
func GetDefaultProfile(rs *specs.Spec) (*specs.LinuxSeccomp, error)
```

---

### LoadProfile

LoadProfile takes a json string and decodes the seccomp profile.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp_linux.go#L20)  

```go
func LoadProfile(body string, rs *specs.Spec) (*specs.LinuxSeccomp, error)
```

---

## Types

### Architecture

Architecture is used to represent a specific architecture
and its sub-architectures

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp.go#L35)  

```go
type Architecture struct {
	Arch      specs.Arch   `json:"architecture"`
	SubArches []specs.Arch `json:"subArchitectures"`
}
```

---

### Filter

Filter is used to conditionally apply Seccomp rules

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp.go#L41)  

```go
type Filter struct {
	Caps   []string `json:"caps,omitempty"`
	Arches []string `json:"arches,omitempty"`

	// MinKernel describes the minimum kernel version the rule must be applied
	// on, in the format "<kernel version>.<major revision>" (e.g. "3.12").
	//
	// When matching the kernel version of the host, minor revisions, and distro-
	// specific suffixes are ignored, which means that "3.12.25-gentoo", "3.12-1-amd64",
	// "3.12", and "3.12-rc5" are considered equal (kernel 3, major revision 12).
	MinKernel *KernelVersion `json:"minKernel,omitempty"`
}
```

---

### KernelVersion

KernelVersion holds information about the kernel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp.go#L69)  

```go
type KernelVersion struct {
	Kernel uint64 // Version of the Kernel (i.e., the "4" in "4.1.2-generic")
	Major  uint64 // Major revision of the Kernel (i.e., the "1" in "4.1.2-generic")
}
```

#### Methods

##### KernelVersion.MarshalJSON

MarshalJSON implements json.Unmarshaler for KernelVersion

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp.go#L83)  

```go
func (k *KernelVersion) MarshalJSON() ([]byte, error)
```

##### KernelVersion.String

String implements fmt.Stringer for KernelVersion

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp.go#L75)  

```go
func (k *KernelVersion) String() string
```

##### KernelVersion.UnmarshalJSON

UnmarshalJSON implements json.Marshaler for KernelVersion

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp.go#L88)  

```go
func (k *KernelVersion) UnmarshalJSON(version []byte) error
```

---

### Seccomp

Seccomp represents the config for a seccomp profile for syscall restriction.
It is used to marshal/unmarshal the JSON profiles as accepted by docker, and
extends the runtime-spec's specs.LinuxSeccomp, overriding some fields to
provide the ability to define conditional rules based on the host's kernel
version, architecture, and the container's capabilities.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp.go#L17)  

```go
type Seccomp struct {
	specs.LinuxSeccomp

	// ArchMap contains a list of Architectures and Sub-architectures for the
	// profile. When generating the profile, this list is expanded to a
	// []specs.Arch, to propagate the Architectures field of the profile.
	ArchMap []Architecture `json:"archMap,omitempty"`

	// Syscalls contains lists of syscall rules. Rules can define conditions
	// for them to be included or excluded in the resulting profile (based on
	// kernel version, architecture, capabilities, etc.). These lists are
	// expanded to an specs.Syscall  When generating the profile, these lists
	// are expanded to a []specs.LinuxSyscall.
	Syscalls []*Syscall `json:"syscalls"`
}
```

#### Functions

##### DefaultProfile

DefaultProfile defines the allowed syscalls for the default seccomp profile.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/default_linux.go#L46)  

```go
func DefaultProfile() *Seccomp
```

---

### Syscall

Syscall is used to match a group of syscalls in Seccomp. It extends the
runtime-spec Syscall type, adding a "Name" field for backward compatibility
with older JSON representations, additional "Comment" metadata, and conditional
rules ("Includes", "Excludes") used to generate a runtime-spec Seccomp profile
based on the container (capabilities) and host's (arch, kernel) configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/profiles/seccomp/seccomp.go#L59)  

```go
type Syscall struct {
	specs.LinuxSyscall
	// Deprecated: kept for backward compatibility with old JSON profiles, use Names instead
	Name     string  `json:"name,omitempty"`
	Comment  string  `json:"comment,omitempty"`
	Includes *Filter `json:"includes,omitempty"`
	Excludes *Filter `json:"excludes,omitempty"`
}
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil/environment

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:21 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### EnsureFrozenImagesLinux

EnsureFrozenImagesLinux loads frozen test images into the daemon
if they aren't already loaded

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L216)  

```go
func EnsureFrozenImagesLinux(ctx context.Context, testEnv *Execution) error
```

---

### ProtectAll

ProtectAll protects the existing environment (containers, images, networks,
volumes, and, on Linux, plugins) from being cleaned up at the end of test
runs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L49)  

```go
func ProtectAll(ctx context.Context, t testing.TB, testEnv *Execution)
```

---

### ProtectContainers

ProtectContainers protects existing containers from being cleaned up at the
end of test runs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L77)  

```go
func ProtectContainers(ctx context.Context, t testing.TB, testEnv *Execution)
```

---

### ProtectDefaultBridge

ProtectDefaultBridge remembers default bridge settings so that, when a test
runs its own daemon and tramples settings of the bridge belonging to the
CI-started bridge, the bridge is restored to its old state before the next
test.

For example, a test may enable IPv6 with a link-local fixed-cidr-v6. That's
likely to break later tests, even if they also start their own daemon
(because, in the absence of any specific settings, the daemon learns default
bridge config from addresses on an existing bridge device).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect_linux.go#L35)  

```go
func ProtectDefaultBridge(_ context.Context, t testing.TB, testEnv *Execution)
```

---

### ProtectImages

ProtectImages protects existing images and on linux frozen images from being
cleaned up at the end of test runs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L108)  

```go
func ProtectImages(ctx context.Context, t testing.TB, testEnv *Execution)
```

---

### ProtectNetworks

ProtectNetworks protects existing networks from being cleaned up at the end
of test runs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L163)  

```go
func ProtectNetworks(ctx context.Context, t testing.TB, testEnv *Execution)
```

---

### ProtectPlugins

ProtectPlugins protects existing plugins from being cleaned up at the end of
test runs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L192)  

```go
func ProtectPlugins(ctx context.Context, t testing.TB, testEnv *Execution)
```

---

### ProtectVolumes

ProtectVolumes protects existing volumes from being cleaned up at the end of
test runs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L225)  

```go
func ProtectVolumes(ctx context.Context, t testing.TB, testEnv *Execution)
```

---

## Types

### Execution

Execution contains information about the current test execution and daemon
under test

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L23)  

```go
type Execution struct {
	DaemonInfo       system.Info
	DaemonVersion    types.Version
	PlatformDefaults PlatformDefaults
	// contains filtered or unexported fields
}
```

#### Functions

##### FromClient

FromClient creates a new Execution environment from the passed in client

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L49)  

```go
func FromClient(ctx context.Context, c *client.Client) (*Execution, error)
```

##### New

New creates a new Execution struct
This is configured using the env client (see client.FromEnv)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L40)  

```go
func New(ctx context.Context) (*Execution, error)
```

#### Methods

##### Execution.APIClient

APIClient returns an APIClient connected to the daemon under test

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L150)  

```go
func (e *Execution) APIClient() client.APIClient
```

##### Execution.Clean

Clean the environment, preserving protected objects (images, containers, ...)
and removing everything else. It's meant to run after any tests so that they don't
depend on each others.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/clean.go#L24)  

```go
func (e *Execution) Clean(ctx context.Context, t testing.TB)
```

##### Execution.DaemonAPIVersion

DaemonAPIVersion returns the negotiated daemon api version

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L131)  

```go
func (e *Execution) DaemonAPIVersion() string
```

##### Execution.FirewallBackendDriver

FirewallBackendDriver returns the value of FirewallBackend.Driver from
system Info if set, else the empty string.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L238)  

```go
func (e *Execution) FirewallBackendDriver() string
```

##### Execution.GitHubActions

GitHubActions is true if test is executed on a GitHub Runner.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L227)  

```go
func (e *Execution) GitHubActions() bool
```

##### Execution.HasExistingImage

HasExistingImage checks whether there is an image with the given reference.
Note that this is done by filtering and then checking whether there were any
results -- so ambiguous references might result in false-positives.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L201)  

```go
func (e *Execution) HasExistingImage(t testing.TB, reference string) bool
```

##### Execution.IsLocalDaemon

IsLocalDaemon is true if the daemon under test is on the same
host as the test process.

Deterministically working out the environment in which CI is running
to evaluate whether the daemon is local or remote is not possible through
a build tag.

For example Windows to Linux CI under Jenkins tests the 64-bit
Windows binary build with the daemon build tag, but calls a remote
Linux daemon.

We can't just say if Windows then assume the daemon is local as at
some point, we will be testing the Windows CLI against a Windows daemon.

Similarly, it will be perfectly valid to also run CLI tests from
a Linux CLI (built with the daemon tag) against a Windows daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L120)  

```go
func (e *Execution) IsLocalDaemon() bool
```

##### Execution.IsRemoteDaemon

IsRemoteDaemon is true if the daemon under test is on different host
as the test process.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L126)  

```go
func (e *Execution) IsRemoteDaemon() bool
```

##### Execution.IsRootless

IsRootless returns whether the rootless mode is enabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L166)  

```go
func (e *Execution) IsRootless() bool
```

##### Execution.IsUserNamespace

IsUserNamespace returns whether the user namespace remapping is enabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L155)  

```go
func (e *Execution) IsUserNamespace() bool
```

##### Execution.IsUserNamespaceInKernel

IsUserNamespaceInKernel returns whether the kernel supports user namespaces

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L171)  

```go
func (e *Execution) IsUserNamespaceInKernel() bool
```

##### Execution.NotAmd64

NotAmd64 returns true if the daemon's architecture is not amd64

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L232)  

```go
func (e *Execution) NotAmd64() bool
```

##### Execution.Print

Print the execution details to stdout
TODO: print everything

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L141)  

```go
func (e *Execution) Print()
```

##### Execution.ProtectContainer

ProtectContainer adds the specified container(s) to be protected in case of
clean

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L68)  

```go
func (e *Execution) ProtectContainer(t testing.TB, containers ...string)
```

##### Execution.ProtectDefaultBridge

ProtectDefaultBridge stores default bridge info, to be restored on clean.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect_linux.go#L64)  

```go
func (e *Execution) ProtectDefaultBridge(t testing.TB, info *defaultBridgeInfo)
```

##### Execution.ProtectImage

ProtectImage adds the specified image(s) to be protected in case of clean

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L99)  

```go
func (e *Execution) ProtectImage(t testing.TB, images ...string)
```

##### Execution.ProtectNetwork

ProtectNetwork adds the specified network(s) to be protected in case of
clean

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L154)  

```go
func (e *Execution) ProtectNetwork(t testing.TB, networks ...string)
```

##### Execution.ProtectPlugin

ProtectPlugin adds the specified plugin(s) to be protected in case of clean

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L183)  

```go
func (e *Execution) ProtectPlugin(t testing.TB, plugins ...string)
```

##### Execution.ProtectVolume

ProtectVolume adds the specified volume(s) to be protected in case of clean

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/protect.go#L216)  

```go
func (e *Execution) ProtectVolume(t testing.TB, volumes ...string)
```

##### Execution.RuntimeIsWindowsContainerd

RuntimeIsWindowsContainerd returns whether containerd runtime is used on Windows

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L161)  

```go
func (e *Execution) RuntimeIsWindowsContainerd() bool
```

##### Execution.UsingSnapshotter

UsingSnapshotter returns whether containerd snapshotters are used for the
tests by checking if the "TEST_INTEGRATION_USE_SNAPSHOTTER" is set to a
non-empty value.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L194)  

```go
func (e *Execution) UsingSnapshotter() bool
```

---

### PlatformDefaults

PlatformDefaults are defaults values for the platform of the daemon under test

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/environment/environment.go#L32)  

```go
type PlatformDefaults struct {
	BaseImage            string
	VolumesConfigPath    string
	ContainerStoragePath string
}
```

---


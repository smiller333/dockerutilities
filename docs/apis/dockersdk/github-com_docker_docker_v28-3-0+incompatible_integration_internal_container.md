# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration/internal/container

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:06:15 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Create

Create creates a container with the specified options, asserting that there was no error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L56)  

```go
func Create(ctx context.Context, t *testing.T, apiClient client.APIClient, ops ...func(*TestContainerConfig)) string
```

---

### CreateFromConfig

CreateFromConfig creates a container from the given TestContainerConfig.

Example use:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L71)  

```go
func CreateFromConfig(ctx context.Context, apiClient client.APIClient, config *TestContainerConfig) (container.CreateResponse, error)
```

---

### GetContainerNS

GetContainerNS gets the value of the specified namespace of a container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ns.go#L14)  

```go
func GetContainerNS(ctx context.Context, t *testing.T, apiClient client.APIClient, cID, nsName string) string
```

---

### Inspect

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L176)  

```go
func Inspect(ctx context.Context, t *testing.T, apiClient client.APIClient, containerRef string) container.InspectResponse
```

---

### IsInState

IsInState verifies the container is in one of the specified state, e.g., "running", "exited", etc.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/states.go#L36)  

```go
func IsInState(ctx context.Context, apiClient client.APIClient, containerID string, state ...container.ContainerState) func(log poll.LogT) poll.Result
```

---

### IsRemoved

IsRemoved verifies the container has been removed

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/states.go#L73)  

```go
func IsRemoved(ctx context.Context, apiClient client.APIClient, containerID string) func(log poll.LogT) poll.Result
```

---

### IsStopped

IsStopped verifies the container is in stopped state.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/states.go#L31)  

```go
func IsStopped(ctx context.Context, apiClient client.APIClient, containerID string) func(log poll.LogT) poll.Result
```

---

### IsSuccessful

IsSuccessful verifies state.Status == "exited" && state.ExitCode == 0

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/states.go#L56)  

```go
func IsSuccessful(ctx context.Context, apiClient client.APIClient, containerID string) func(log poll.LogT) poll.Result
```

---

### Remove

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L158)  

```go
func Remove(ctx context.Context, t *testing.T, apiClient client.APIClient, container string, options container.RemoveOptions)
```

---

### RemoveAll

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L165)  

```go
func RemoveAll(ctx context.Context, t *testing.T, apiClient client.APIClient)
```

---

### Run

Run creates and start a container with the specified options

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L76)  

```go
func Run(ctx context.Context, t *testing.T, apiClient client.APIClient, ops ...func(*TestContainerConfig)) string
```

---

### RunningStateFlagIs

RunningStateFlagIs polls for the container's Running state flag to be equal to running.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/states.go#L15)  

```go
func RunningStateFlagIs(ctx context.Context, apiClient client.APIClient, containerID string, running bool) func(log poll.LogT) poll.Result
```

---

### WithAdditionalGroups

WithAdditionalGroups sets the additional groups for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L241)  

```go
func WithAdditionalGroups(groups ...string) func(c *TestContainerConfig)
```

---

### WithAnnotations

WithAnnotations set the annotations for the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L304)  

```go
func WithAnnotations(annotations map[string]string) func(*TestContainerConfig)
```

---

### WithAutoRemove

WithAutoRemove sets the container to be removed on exit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L212)  

```go
func WithAutoRemove(c *TestContainerConfig)
```

---

### WithBind

WithBind sets the bind mount of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L123)  

```go
func WithBind(src, target string) func(*TestContainerConfig)
```

---

### WithBindRaw

WithBindRaw sets the bind mount of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L130)  

```go
func WithBindRaw(s string) func(*TestContainerConfig)
```

---

### WithCDIDevices

WithCDIDevices sets the CDI devices to use to start the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L318)  

```go
func WithCDIDevices(cdiDeviceNames ...string) func(*TestContainerConfig)
```

---

### WithCapability

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L328)  

```go
func WithCapability(capabilities ...string) func(*TestContainerConfig)
```

---

### WithCgroupnsMode

WithCgroupnsMode sets the cgroup namespace mode for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L258)  

```go
func WithCgroupnsMode(mode string) func(*TestContainerConfig)
```

---

### WithCmd

WithCmd sets the commands of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L44)  

```go
func WithCmd(cmds ...string) func(*TestContainerConfig)
```

---

### WithConsoleSize

WithConsoleSize sets the initial console size of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L297)  

```go
func WithConsoleSize(width, height uint) func(*TestContainerConfig)
```

---

### WithContainerWideMacAddress

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L359)  

```go
func WithContainerWideMacAddress(address string) func(c *TestContainerConfig)
```

---

### WithDNS

WithDNS sets external DNS servers for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L58)  

```go
func WithDNS(dns []string) func(*TestContainerConfig)
```

---

### WithDropCapability

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L334)  

```go
func WithDropCapability(capabilities ...string) func(*TestContainerConfig)
```

---

### WithEndpointSettings

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L193)  

```go
func WithEndpointSettings(nw string, config *network.EndpointSettings) func(*TestContainerConfig)
```

---

### WithExposedPorts

WithExposedPorts sets the exposed ports of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L72)  

```go
func WithExposedPorts(ports ...string) func(*TestContainerConfig)
```

---

### WithExtraHost

WithExtraHost sets the user defined IP:Host mappings in the container's
/etc/hosts file

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L269)  

```go
func WithExtraHost(extraHost string) func(*TestContainerConfig)
```

---

### WithHostname

WithHostname sets the hostname of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L23)  

```go
func WithHostname(name string) func(*TestContainerConfig)
```

---

### WithIPv4

WithIPv4 sets the specified ip for the specified network of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L162)  

```go
func WithIPv4(networkName, ip string) func(*TestContainerConfig)
```

---

### WithIPv6

WithIPv6 sets the specified ip6 for the specified network of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L178)  

```go
func WithIPv6(networkName, ip string) func(*TestContainerConfig)
```

---

### WithImage

WithImage sets the image of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L37)  

```go
func WithImage(image string) func(*TestContainerConfig)
```

---

### WithIsolation

WithIsolation specifies the isolation technology to apply to the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L290)  

```go
func WithIsolation(isolation container.Isolation) func(*TestContainerConfig)
```

---

### WithLinks

WithLinks sets the links of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L30)  

```go
func WithLinks(links ...string) func(*TestContainerConfig)
```

---

### WithLogDriver

WithLogDriver sets the log driver to use for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L205)  

```go
func WithLogDriver(driver string) func(*TestContainerConfig)
```

---

### WithMacAddress

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L149)  

```go
func WithMacAddress(networkName, mac string) func(config *TestContainerConfig)
```

---

### WithMount

WithMount adds an mount

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L106)  

```go
func WithMount(m mount.Mount) func(*TestContainerConfig)
```

---

### WithName

WithName sets the name of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L16)  

```go
func WithName(name string) func(*TestContainerConfig)
```

---

### WithNetworkMode

WithNetworkMode sets the network mode of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L51)  

```go
func WithNetworkMode(mode string) func(*TestContainerConfig)
```

---

### WithPIDMode

WithPIDMode sets the PID-mode for the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L347)  

```go
func WithPIDMode(mode container.PidMode) func(c *TestContainerConfig)
```

---

### WithPidsLimit

WithPidsLimit sets the container's "pids-limit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L217)  

```go
func WithPidsLimit(limit *int64) func(*TestContainerConfig)
```

---

### WithPlatform

WithPlatform specifies the desired platform the image should have.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L276)  

```go
func WithPlatform(p *ocispec.Platform) func(*TestContainerConfig)
```

---

### WithPortMap

WithPortMap sets/replaces port mappings.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L82)  

```go
func WithPortMap(pm nat.PortMap) func(*TestContainerConfig)
```

---

### WithPrivileged

WithPrivileged sets privileged mode for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L248)  

```go
func WithPrivileged(privileged bool) func(*TestContainerConfig)
```

---

### WithRestartPolicy

WithRestartPolicy sets container's restart policy

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L227)  

```go
func WithRestartPolicy(policy container.RestartPolicyMode) func(c *TestContainerConfig)
```

---

### WithRuntime

WithRuntime sets the runtime to use to start the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L311)  

```go
func WithRuntime(name string) func(*TestContainerConfig)
```

---

### WithSecurityOpt

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L340)  

```go
func WithSecurityOpt(opt string) func(*TestContainerConfig)
```

---

### WithStopSignal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L353)  

```go
func WithStopSignal(stopSignal string) func(c *TestContainerConfig)
```

---

### WithSysctls

WithSysctls sets sysctl options for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L65)  

```go
func WithSysctls(sysctls map[string]string) func(*TestContainerConfig)
```

---

### WithTmpfs

WithTmpfs sets a target path in the container to a tmpfs, with optional options
(separated with a colon).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L138)  

```go
func WithTmpfs(targetAndOpts string) func(config *TestContainerConfig)
```

---

### WithTty

WithTty sets the TTY mode of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L92)  

```go
func WithTty(tty bool) func(*TestContainerConfig)
```

---

### WithUser

WithUser sets the user

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L234)  

```go
func WithUser(user string) func(c *TestContainerConfig)
```

---

### WithVolume

WithVolume sets the volume of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L113)  

```go
func WithVolume(target string) func(*TestContainerConfig)
```

---

### WithWindowsDevice

WithWindowsDevice specifies a Windows Device, ala `--device` on the CLI

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L283)  

```go
func WithWindowsDevice(device string) func(*TestContainerConfig)
```

---

### WithWorkingDir

WithWorkingDir sets the working dir of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/ops.go#L99)  

```go
func WithWorkingDir(dir string) func(*TestContainerConfig)
```

---

## Types

### ContainerOutput

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L185)  

```go
type ContainerOutput struct {
	Stdout, Stderr string
}
```

#### Functions

##### Output

Output waits for the container to end running and returns its output.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L190)  

```go
func Output(ctx context.Context, client client.APIClient, id string) (ContainerOutput, error)
```

---

### ExecResult

ExecResult represents a result returned from Exec()

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/exec.go#L13)  

```go
type ExecResult struct {
	ExitCode int
	// contains filtered or unexported fields
}
```

#### Functions

##### Exec

Exec executes a command inside a container, returning the result
containing stdout, stderr, and exit code. Note:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/exec.go#L50)  

```go
func Exec(ctx context.Context, apiClient client.APIClient, id string, cmd []string, ops ...func(*container.ExecOptions)) (ExecResult, error)
```

##### ExecT

ExecT calls Exec() and aborts the test if an error occurs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/exec.go#L90)  

```go
func ExecT(ctx context.Context, t testing.TB, apiClient client.APIClient, id string, cmd []string, ops ...func(*container.ExecOptions)) ExecResult
```

#### Methods

##### ExecResult.AssertSuccess

AssertSuccess fails the test and stops execution if the command exited with a
nonzero status code.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/exec.go#L36)  

```go
func (res ExecResult) AssertSuccess(t testing.TB)
```

##### ExecResult.Combined

Combined returns combined stdout and stderr output of a command run by Exec()

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/exec.go#L30)  

```go
func (res ExecResult) Combined() string
```

##### ExecResult.Stderr

Stderr returns stderr output of a command run by Exec()

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/exec.go#L25)  

```go
func (res ExecResult) Stderr() string
```

##### ExecResult.Stdout

Stdout returns stdout output of a command run by Exec()

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/exec.go#L20)  

```go
func (res ExecResult) Stdout() string
```

---

### RunResult

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L86)  

```go
type RunResult struct {
	ContainerID string
	ExitCode    int
	Stdout      *bytes.Buffer
	Stderr      *bytes.Buffer
}
```

#### Functions

##### RunAttach

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L93)  

```go
func RunAttach(ctx context.Context, t *testing.T, apiClient client.APIClient, ops ...func(config *TestContainerConfig)) RunResult
```

---

### TestContainerConfig

TestContainerConfig holds container configuration struct that
are used in api calls.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L22)  

```go
type TestContainerConfig struct {
	Name             string
	Config           *container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
	Platform         *ocispec.Platform
}
```

#### Functions

##### NewTestConfig

NewTestConfig creates a new TestContainerConfig with the provided options.

If no options are passed, it creates a default config, which is a busybox
container running "top" (on Linux) or "sleep" (on Windows).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/container/container.go#L34)  

```go
func NewTestConfig(ops ...func(*TestContainerConfig)) *TestContainerConfig
```

---


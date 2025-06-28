# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:42 UTC

## Overview

Package daemon exposes the functions that occur on the host server
that the Docker daemon is running.

In implementing the various functions of the daemon, there is often
a method-specific struct for configuring the runtime behavior.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CreateDaemonRoot

CreateDaemonRoot creates the root for the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1535)  
**Added in:** v1.13.0

```go
func CreateDaemonRoot(config *config.Config) error
```

---

### DefaultApparmorProfile

DefaultApparmorProfile returns the name of the default apparmor profile

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/apparmor_default.go#L19)  

```go
func DefaultApparmorProfile() string
```

---

### LoadOrCreateID

LoadOrCreateID loads the engine's ID from the given root, or generates a new ID
if it doesn't exist. It returns the ID, and any error that occurred when
saving the file.

Note that this function expects the daemon's root directory to already have
been created with the right permissions and ownership (usually this would
be done by daemon.CreateDaemonRoot().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/id.go#L21)  

```go
func LoadOrCreateID(root string) (string, error)
```

---

### RegisterCDIDriver

RegisterCDIDriver registers the CDI device driver.
The driver injects CDI devices into an incoming OCI spec and is called for DeviceRequests associated with CDI devices.
If the list of CDI spec directories is empty, the driver is not registered.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cdi.go#L25)  

```go
func RegisterCDIDriver(cdiSpecDirs ...string) *cdi.Cache
```

---

### RemapContainerdNamespaces

RemapContainerdNamespaces returns the right containerd namespaces to use:
- if they are not already set in the config file
-  and the daemon is running with user namespace remapping enabled
Then it will return new namespace names, otherwise it will return the existing
namespaces

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1560)  

```go
func RemapContainerdNamespaces(config *config.Config) (ns string, pluginNs string, err error)
```

---

### Rootless

Rootless returns true if daemon is running in rootless mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/info_unix.go#L417)  

```go
func Rootless(cfg *config.Config) bool
```

---

### UsingSystemd

UsingSystemd returns true if cli option includes native.cgroupdriver=systemd

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon_unix.go#L615)  
**Added in:** v1.11.0

```go
func UsingSystemd(config *config.Config) bool
```

---

### WithApparmor

WithApparmor sets the apparmor profile

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_linux.go#L124)  

```go
func WithApparmor(c *container.Container) coci.SpecOpts
```

---

### WithCapabilities

WithCapabilities sets the container's capabilities

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_linux.go#L157)  

```go
func WithCapabilities(c *container.Container) coci.SpecOpts
```

---

### WithConsoleSize

WithConsoleSize sets the initial console size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_opts.go#L13)  

```go
func WithConsoleSize(c *container.Container) coci.SpecOpts
```

---

### WithDevices

WithDevices sets the container's devices

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_linux.go#L845)  

```go
func WithDevices(daemon *Daemon, c *container.Container) coci.SpecOpts
```

---

### WithNamespaces

WithNamespaces sets the container's namespaces

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_linux.go#L226)  

```go
func WithNamespaces(daemon *Daemon, c *container.Container) coci.SpecOpts
```

---

### WithOOMScore

WithOOMScore sets the oom score

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_linux.go#L98)  

```go
func WithOOMScore(score *int) coci.SpecOpts
```

---

### WithResources

WithResources applies the container resources

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_linux.go#L918)  

```go
func WithResources(c *container.Container) coci.SpecOpts
```

---

### WithSeccomp

WithSeccomp sets the seccomp profile

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/seccomp_linux.go#L19)  

```go
func WithSeccomp(daemon *Daemon, c *container.Container) coci.SpecOpts
```

---

### WithSelinux

WithSelinux sets the selinux labels

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_linux.go#L109)  

```go
func WithSelinux(c *container.Container) coci.SpecOpts
```

---

### WithSysctls

WithSysctls sets the container's sysctls

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_linux.go#L974)  

```go
func WithSysctls(c *container.Container) coci.SpecOpts
```

---

### WithUser

WithUser sets the container's user

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/oci_linux.go#L995)  

```go
func WithUser(c *container.Container) coci.SpecOpts
```

---

## Types

### Cluster

Cluster is the interface for github.com/docker/docker/daemon/cluster.(*Cluster).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster.go#L10)  
**Added in:** v1.13.0

```go
type Cluster interface {
	ClusterStatus
	NetworkManager
	SendClusterEvent(event lncluster.ConfigEventType)
}
```

---

### ClusterStatus

ClusterStatus interface provides information about the Swarm status of the Cluster

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster.go#L17)  

```go
type ClusterStatus interface {
	IsAgent() bool
	IsManager() bool
}
```

---

### Daemon

Daemon holds information about the Docker daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L99)  

```go
type Daemon struct {
	EventsService *events.Events

	PluginStore *plugin.Store // TODO: remove

	ReferenceStore refstore.Store

	CDICache *cdi.Cache
	// contains filtered or unexported fields
}
```

#### Functions

##### NewDaemon

NewDaemon sets up everything for the daemon to be able to service
requests from the webserver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L768)  

```go
func NewDaemon(ctx context.Context, config *config.Config, pluginStore *plugin.Store, authzMiddleware *authorization.Middleware) (_ *Daemon, retErr error)
```

#### Methods

##### Daemon.ActivateContainerServiceBinding

ActivateContainerServiceBinding puts this container into load balancer active rotation and DNS response

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/container_operations.go#L1067)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) ActivateContainerServiceBinding(containerName string) error
```

##### Daemon.AuthenticateToRegistry

AuthenticateToRegistry checks the validity of credentials in authConfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/auth.go#L11)  
**Added in:** v1.9.0

```go
func (daemon *Daemon) AuthenticateToRegistry(ctx context.Context, authConfig *registry.AuthConfig) (string, string, error)
```

##### Daemon.BuilderBackend

BuilderBackend returns the backend used by builder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1630)  

```go
func (daemon *Daemon) BuilderBackend() builder.Backend
```

##### Daemon.CheckpointCreate

CheckpointCreate checkpoints the process running in a container with CRIU

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/checkpoint.go#L55)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) CheckpointCreate(name string, config checkpoint.CreateOptions) error
```

##### Daemon.CheckpointDelete

CheckpointDelete deletes the specified checkpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/checkpoint.go#L89)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) CheckpointDelete(name string, config checkpoint.DeleteOptions) error
```

##### Daemon.CheckpointList

CheckpointList lists all checkpoints of the specified container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/checkpoint.go#L102)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) CheckpointList(name string, config checkpoint.ListOptions) ([]checkpoint.Summary, error)
```

##### Daemon.Cleanup

Cleanup releases any network resources allocated to the container along with any rules
around how containers are linked together.  It also unmounts the container's root filesystem.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/start.go#L270)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) Cleanup(ctx context.Context, container *container.Container)
```

##### Daemon.Config

Config returns daemon's config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L190)  

```go
func (daemon *Daemon) Config() config.Config
```

##### Daemon.ConnectContainerToNetwork

ConnectContainerToNetwork connects the given container to the given
network. If either cannot be found, an err is returned. If the
network cannot be set up, an err is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L484)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) ConnectContainerToNetwork(ctx context.Context, containerName, networkName string, endpointConfig *networktypes.EndpointSettings) error
```

##### Daemon.ConnectToNetwork

ConnectToNetwork connects a container to a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/container_operations.go#L994)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) ConnectToNetwork(ctx context.Context, ctr *container.Container, idOrName string, endpointConfig *networktypes.EndpointSettings) error
```

##### Daemon.ContainerArchivePath

ContainerArchivePath creates an archive of the filesystem resource at the
specified path in the container identified by the given name. Returns a
tar archive of the resource and whether it was a directory or a single file.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/archive.go#L37)  
**Added in:** v1.8.0

```go
func (daemon *Daemon) ContainerArchivePath(name string, path string) (content io.ReadCloser, stat *container.PathStat, _ error)
```

##### Daemon.ContainerAttach

ContainerAttach attaches to logs according to the config passed in. See ContainerAttachConfig.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/attach.go#L22)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerAttach(prefixOrName string, req *backend.ContainerAttachConfig) error
```

##### Daemon.ContainerAttachRaw

ContainerAttachRaw attaches the provided streams to the container's stdio

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/attach.go#L98)  
**Added in:** v1.11.0

```go
func (daemon *Daemon) ContainerAttachRaw(prefixOrName string, stdin io.ReadCloser, stdout, stderr io.Writer, doStream bool, attached chan struct{}) error
```

##### Daemon.ContainerChanges

ContainerChanges returns a list of container fs changes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/changes.go#L13)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerChanges(ctx context.Context, name string) ([]archive.Change, error)
```

##### Daemon.ContainerCreate

ContainerCreate creates a regular container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/create.go#L53)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerCreate(ctx context.Context, params backend.ContainerCreateConfig) (containertypes.CreateResponse, error)
```

##### Daemon.ContainerCreateIgnoreImagesArgsEscaped

ContainerCreateIgnoreImagesArgsEscaped creates a regular container. This is called from the builder RUN case
and ensures that we do not take the images ArgsEscaped

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/create.go#L61)  

```go
func (daemon *Daemon) ContainerCreateIgnoreImagesArgsEscaped(ctx context.Context, params backend.ContainerCreateConfig) (containertypes.CreateResponse, error)
```

##### Daemon.ContainerCreateWorkdir

ContainerCreateWorkdir creates the working directory. This solves the
issue arising from https://github.com/docker/docker/issues/27545,
which was initially fixed by https://github.com/docker/docker/pull/27884. But that fix
was too expensive in terms of performance on Windows. Instead,
https://github.com/docker/docker/pull/28514 introduces this new functionality
where the builder calls into the backend here to create the working directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/workdir.go#L9)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) ContainerCreateWorkdir(cID string) error
```

##### Daemon.ContainerExecCreate

ContainerExecCreate sets up an exec in a running container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/exec.go#L96)  
**Added in:** v1.3.0

```go
func (daemon *Daemon) ContainerExecCreate(name string, options *containertypes.ExecOptions) (string, error)
```

##### Daemon.ContainerExecInspect

ContainerExecInspect returns low-level information about the exec
command. An error is returned if the exec cannot be found.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/inspect.go#L199)  
**Added in:** v1.4.0

```go
func (daemon *Daemon) ContainerExecInspect(id string) (*backend.ExecInspect, error)
```

##### Daemon.ContainerExecResize

ContainerExecResize changes the size of the TTY of the process
running in the exec with the given name to the given height and
width.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/resize.go#L40)  
**Added in:** v1.3.0

```go
func (daemon *Daemon) ContainerExecResize(ctx context.Context, name string, height, width uint32) error
```

##### Daemon.ContainerExecStart

ContainerExecStart starts a previously set up exec instance. The
std streams are set up.
If ctx is cancelled, the process is terminated.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/exec.go#L162)  
**Added in:** v1.3.0

```go
func (daemon *Daemon) ContainerExecStart(ctx context.Context, name string, options backend.ExecStartConfig) (retErr error)
```

##### Daemon.ContainerExport

ContainerExport writes the contents of the container to the given
writer. An error is returned if the container cannot be found.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/export.go#L18)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerExport(ctx context.Context, name string, out io.Writer) error
```

##### Daemon.ContainerExtractToDir

ContainerExtractToDir extracts the given archive to the specified location
in the filesystem of the container identified by the given name. The given
path must be of a directory in the container. If it is not, the error will
be an errdefs.InvalidParameter. If noOverwriteDirNonDir is true then it will
be an error if unpacking the given content would cause an existing directory
to be replaced with a non-directory and vice versa.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/archive.go#L63)  
**Added in:** v1.8.0

```go
func (daemon *Daemon) ContainerExtractToDir(name, path string, copyUIDGID, noOverwriteDirNonDir bool, content io.Reader) error
```

##### Daemon.ContainerInspect

ContainerInspect returns low-level information about a
container. Returns an error if the container cannot be found, or if
there is an error getting the data.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/inspect.go#L25)  
**Added in:** v0.12.0

```go
func (daemon *Daemon) ContainerInspect(ctx context.Context, name string, options backend.ContainerInspectOptions) (*containertypes.InspectResponse, error)
```

##### Daemon.ContainerKill

ContainerKill sends signal to the container
If no signal is given, then Kill with SIGKILL and wait
for the container to exit.
If a signal is given, then just send it to the container and return.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/kill.go#L36)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerKill(name, stopSignal string) error
```

##### Daemon.ContainerLogs

ContainerLogs copies the container's log channel to the channel provided in
the config. If ContainerLogs returns an error, no messages have been copied.
and the channel will be closed without data.

if it returns nil, the config channel will be active and return log
messages until it runs out or the context is canceled.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logs.go#L27)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerLogs(ctx context.Context, containerName string, config *containertypes.LogsOptions) (messages <-chan *backend.LogMessage, isTTY bool, retErr error)
```

##### Daemon.ContainerPause

ContainerPause pauses a container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/pause.go#L13)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerPause(name string) error
```

##### Daemon.ContainerRename

ContainerRename changes the name of a container, using the oldName
to find the container. An error is returned if newName is already
reserved.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/rename.go#L20)  
**Added in:** v1.5.0

```go
func (daemon *Daemon) ContainerRename(oldName, newName string) (retErr error)
```

##### Daemon.ContainerResize

ContainerResize changes the size of the TTY of the process running
in the container with the given name to the given height and width.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/resize.go#L15)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerResize(ctx context.Context, name string, height, width uint32) error
```

##### Daemon.ContainerRestart

ContainerRestart stops and starts a container. It attempts to
gracefully stop the container within the given timeout, forcefully
stopping it if the timeout is exceeded. If given a negative
timeout, ContainerRestart will wait forever until a graceful
stop. Returns an error if the container cannot be found, or if
there is an underlying error at any stage of the restart.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/restart.go#L18)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerRestart(ctx context.Context, name string, options containertypes.StopOptions) error
```

##### Daemon.ContainerRm

ContainerRm removes the container id from the filesystem. An error
is returned if the container is not found, or if the remove
fails. If the remove succeeds, the container name is released, and
network links are removed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/delete.go#L28)  
**Added in:** v1.3.0

```go
func (daemon *Daemon) ContainerRm(name string, config *backend.ContainerRmConfig) error
```

##### Daemon.ContainerStart

ContainerStart starts a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/start.go#L49)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerStart(ctx context.Context, name string, checkpoint string, checkpointDir string) error
```

##### Daemon.ContainerStatPath

ContainerStatPath stats the filesystem resource at the specified path in the
container identified by the given name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/archive.go#L14)  
**Added in:** v1.8.0

```go
func (daemon *Daemon) ContainerStatPath(name string, path string) (*container.PathStat, error)
```

##### Daemon.ContainerStats

ContainerStats writes information about the container to the stream
given in the config object.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/stats.go#L19)  
**Added in:** v1.5.0

```go
func (daemon *Daemon) ContainerStats(ctx context.Context, prefixOrName string, config *backend.ContainerStatsConfig) error
```

##### Daemon.ContainerStop

ContainerStop looks for the given container and stops it.
In case the container fails to stop gracefully within a time duration
specified by the timeout argument, in seconds, it is forcefully
terminated (killed).

If the timeout is nil, the container's StopTimeout value is used, if set,
otherwise the engine default. A negative timeout value can be specified,
meaning no timeout, i.e. no forceful termination is performed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/stop.go#L24)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerStop(ctx context.Context, name string, options containertypes.StopOptions) error
```

##### Daemon.ContainerTop

ContainerTop lists the processes running inside of the given
container by calling ps with the given args, or with the flags
"-ef" if no args are given.  An error is returned if the container
is not found, or is not running, or if there are any problems
running ps, or parsing the output.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/top_unix.go#L140)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerTop(name string, psArgs string) (*container.TopResponse, error)
```

##### Daemon.ContainerUnpause

ContainerUnpause unpauses a container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/unpause.go#L13)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerUnpause(name string) error
```

##### Daemon.ContainerUpdate

ContainerUpdate updates configuration of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/update.go#L15)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) ContainerUpdate(name string, hostConfig *container.HostConfig) (container.UpdateResponse, error)
```

##### Daemon.ContainerWait

ContainerWait waits until the given container is in a certain state
indicated by the given condition. If the container is not found, a nil
channel and non-nil error is returned immediately. If the container is
found, a status result will be sent on the returned channel once the wait
condition is met or if an error occurs waiting for the container (such as a
context timeout or cancellation). On a successful wait, the exit code of the
container is returned in the status with a non-nil Err() value.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/wait.go#L16)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) ContainerWait(ctx context.Context, name string, condition containertypes.WaitCondition) (<-chan containertypes.StateStatus, error)
```

##### Daemon.Containers

Containers returns the list of containers to show given the user's filtering.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/list.go#L105)  
**Added in:** v1.2.0

```go
func (daemon *Daemon) Containers(ctx context.Context, config *containertypes.ListOptions) ([]*containertypes.Summary, error)
```

##### Daemon.ContainersPrune

ContainersPrune removes unused containers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/prune.go#L40)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (*container.PruneReport, error)
```

##### Daemon.CreateImageFromContainer

CreateImageFromContainer creates a new image from a container. The container
config will be updated by applying the change set to the custom config, then
applying that config over the existing container config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/commit.go#L126)  

```go
func (daemon *Daemon) CreateImageFromContainer(ctx context.Context, name string, c *backend.CreateImageConfig) (string, error)
```

##### Daemon.CreateManagedContainer

CreateManagedContainer creates a container that is managed by a Service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/create.go#L45)  
**Added in:** v1.12.0

```go
func (daemon *Daemon) CreateManagedContainer(ctx context.Context, params backend.ContainerCreateConfig) (containertypes.CreateResponse, error)
```

##### Daemon.CreateManagedNetwork

CreateManagedNetwork creates an agent network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L276)  
**Added in:** v1.12.0

```go
func (daemon *Daemon) CreateManagedNetwork(create clustertypes.NetworkCreateRequest) error
```

##### Daemon.CreateNetwork

CreateNetwork creates a network with the given name, driver and other optional parameters

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L282)  
**Added in:** v1.9.0

```go
func (daemon *Daemon) CreateNetwork(ctx context.Context, create networktypes.CreateRequest) (*networktypes.CreateResponse, error)
```

##### Daemon.DaemonJoinsCluster

DaemonJoinsCluster informs the daemon has joined the cluster and provides
the handler to query the cluster component

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L717)  
**Added in:** v1.13.1

```go
func (daemon *Daemon) DaemonJoinsCluster(clusterProvider cluster.Provider)
```

##### Daemon.DaemonLeavesCluster

DaemonLeavesCluster informs the daemon has left the cluster

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L722)  
**Added in:** v1.13.1

```go
func (daemon *Daemon) DaemonLeavesCluster()
```

##### Daemon.DeactivateContainerServiceBinding

DeactivateContainerServiceBinding removes this container from load balancer active rotation, and DNS response

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/container_operations.go#L1080)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) DeactivateContainerServiceBinding(containerName string) error
```

##### Daemon.DeleteManagedNetwork

DeleteManagedNetwork deletes an agent network.
The requirement of networkID is enforced.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L541)  
**Added in:** v1.12.0

```go
func (daemon *Daemon) DeleteManagedNetwork(networkID string) error
```

##### Daemon.DeleteNetwork

DeleteNetwork destroys a network unless it's one of docker's predefined networks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L550)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) DeleteNetwork(networkID string) error
```

##### Daemon.DisconnectContainerFromNetwork

DisconnectContainerFromNetwork disconnects the given container from
the given network. If either cannot be found, an err is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L494)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) DisconnectContainerFromNetwork(containerName string, networkName string, force bool) error
```

##### Daemon.DisconnectFromNetwork

DisconnectFromNetwork disconnects container from network n.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/container_operations.go#L1029)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) DisconnectFromNetwork(ctx context.Context, ctr *container.Container, networkName string, force bool) error
```

##### Daemon.DistributionServices

DistributionServices returns services controlling daemon storage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1214)  

```go
func (daemon *Daemon) DistributionServices() images.DistributionServices
```

##### Daemon.ExecExists

ExecExists looks up the exec instance and returns a bool if it exists or not.
It will also return the error produced by `getConfig`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/exec.go#L36)  
**Added in:** v1.9.0

```go
func (daemon *Daemon) ExecExists(name string) (bool, error)
```

##### Daemon.Features

Features returns the features map from configStore

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L200)  

```go
func (daemon *Daemon) Features() map[string]bool
```

##### Daemon.FindNetwork

FindNetwork returns a network based on:
1. Full ID
2. Full Name
3. Partial ID
as long as there is no ambiguity

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L59)  
**Added in:** v1.9.0

```go
func (daemon *Daemon) FindNetwork(term string) (*libnetwork.Network, error)
```

##### Daemon.ForceEndpointDelete

ForceEndpointDelete deletes an endpoint from a network forcefully

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/container_operations.go#L804)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) ForceEndpointDelete(name string, networkName string) error
```

##### Daemon.GetAttachmentStore

GetAttachmentStore returns current attachment store associated with the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1602)  

```go
func (daemon *Daemon) GetAttachmentStore() *network.AttachmentStore
```

##### Daemon.GetByName

GetByName returns a container given a name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/container.go#L180)  

```go
func (daemon *Daemon) GetByName(name string) (*container.Container, error)
```

##### Daemon.GetCluster

GetCluster returns the cluster

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1506)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) GetCluster() Cluster
```

##### Daemon.GetContainer

GetContainer looks for a container using the provided information, which could be
one of the following inputs from the caller:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/container.go#L38)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) GetContainer(prefixOrName string) (*container.Container, error)
```

##### Daemon.GetContainerStats

GetContainerStats collects all the stats published by a container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/stats.go#L102)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) GetContainerStats(container *container.Container) (*containertypes.StatsResponse, error)
```

##### Daemon.GetNetworkByID

GetNetworkByID function returns a network whose ID matches the given ID.
It fails with an error if no matching network is found.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L92)  
**Added in:** v1.11.0

```go
func (daemon *Daemon) GetNetworkByID(id string) (*libnetwork.Network, error)
```

##### Daemon.GetNetworkByName

GetNetworkByName function returns a network for a given network name.
If no network name is given, the default network is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L102)  
**Added in:** v1.11.0

```go
func (daemon *Daemon) GetNetworkByName(name string) (*libnetwork.Network, error)
```

##### Daemon.GetNetworkDriverList

GetNetworkDriverList returns the list of plugins drivers
registered for network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L507)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) GetNetworkDriverList(ctx context.Context) []string
```

##### Daemon.GetNetworks

GetNetworks returns a list of all networks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L591)  
**Added in:** v1.12.0

```go
func (daemon *Daemon) GetNetworks(filter filters.Args, config backend.NetworkListConfig) ([]networktypes.Inspect, error)
```

##### Daemon.GetNetworksByIDPrefix

GetNetworksByIDPrefix returns a list of networks whose ID partially matches zero or more networks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L114)  

```go
func (daemon *Daemon) GetNetworksByIDPrefix(partialID string) []*libnetwork.Network
```

##### Daemon.HasExperimental

HasExperimental returns whether the experimental features of the daemon are enabled or not

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L195)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) HasExperimental() bool
```

##### Daemon.ID

ID returns the daemon id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L158)  
**Added in:** v1.4.0

```go
func (daemon *Daemon) ID() string
```

##### Daemon.IdentityMapping

IdentityMapping returns uid/gid mapping or a SID (in the case of Windows) for the builder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1607)  

```go
func (daemon *Daemon) IdentityMapping() user.IdentityMapping
```

##### Daemon.ImageBackend

ImageBackend returns an image-backend for Swarm and the distribution router.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1617)  

```go
func (daemon *Daemon) ImageBackend() executorpkg.ImageBackend
```

##### Daemon.ImageExportedByBuildkit

ImageExportedByBuildkit is a callback that is called when an image is exported by buildkit.
This is used to log the image creation event for untagged images.
When no tag is given, buildkit doesn't call the image service so it has no
way of knowing the image was created.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/build.go#L15)  

```go
func (daemon *Daemon) ImageExportedByBuildkit(ctx context.Context, id string, desc ocispec.Descriptor)
```

##### Daemon.ImageNamedByBuildkit

ImageNamedByBuildkit is a callback that is called when an image is tagged by buildkit.
Note: It is only called if the buildkit didn't call the image service itself to perform the tagging.
Currently this only happens when the containerd image store is used.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/build.go#L22)  

```go
func (daemon *Daemon) ImageNamedByBuildkit(ctx context.Context, ref reference.NamedTagged, desc ocispec.Descriptor)
```

##### Daemon.ImageService

ImageService returns the Daemon's ImageService

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1612)  

```go
func (daemon *Daemon) ImageService() ImageService
```

##### Daemon.IsShuttingDown

IsShuttingDown tells whether the daemon is shutting down or not

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1444)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) IsShuttingDown() bool
```

##### Daemon.IsSwarmCompatible

IsSwarmCompatible verifies if the current daemon
configuration is compatible with the swarm mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L762)  
**Added in:** v1.12.0

```go
func (daemon *Daemon) IsSwarmCompatible() error
```

##### Daemon.Kill

Kill forcefully terminates a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/kill.go#L161)  

```go
func (daemon *Daemon) Kill(container *containerpkg.Container) error
```

##### Daemon.List

List returns an array of all containers registered in the daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/list.go#L56)  

```go
func (daemon *Daemon) List() []*container.Container
```

##### Daemon.LogContainerEvent

LogContainerEvent generates an event related to a container with only the default attributes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L20)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) LogContainerEvent(container *container.Container, action events.Action)
```

##### Daemon.LogContainerEventWithAttributes

LogContainerEventWithAttributes generates an event related to a container with specific given attributes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L25)  
**Added in:** v1.11.0

```go
func (daemon *Daemon) LogContainerEventWithAttributes(container *container.Container, action events.Action, attributes map[string]string)
```

##### Daemon.LogDaemonEventWithAttributes

LogDaemonEventWithAttributes generates an event related to the daemon itself with specific given attributes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L69)  
**Added in:** v1.12.0

```go
func (daemon *Daemon) LogDaemonEventWithAttributes(action events.Action, attributes map[string]string)
```

##### Daemon.LogNetworkEvent

LogNetworkEvent generates an event related to a network with only the default attributes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L54)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) LogNetworkEvent(nw *libnetwork.Network, action events.Action)
```

##### Daemon.LogNetworkEventWithAttributes

LogNetworkEventWithAttributes generates an event related to a network with specific given attributes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L59)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) LogNetworkEventWithAttributes(nw *libnetwork.Network, action events.Action, attributes map[string]string)
```

##### Daemon.LogPluginEvent

LogPluginEvent generates an event related to a plugin with only the default attributes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L38)  
**Added in:** v1.12.1

```go
func (daemon *Daemon) LogPluginEvent(pluginID, refName string, action events.Action)
```

##### Daemon.LogVolumeEvent

LogVolumeEvent generates an event related to a volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L46)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) LogVolumeEvent(volumeID string, action events.Action, attributes map[string]string)
```

##### Daemon.Mount

Mount sets container.BaseFS

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1345)  

```go
func (daemon *Daemon) Mount(container *container.Container) error
```

##### Daemon.NetworkController

NetworkController returns the network controller created by the daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L50)  

```go
func (daemon *Daemon) NetworkController() *libnetwork.Controller
```

##### Daemon.NetworksPrune

NetworksPrune removes unused networks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/prune.go#L190)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) NetworksPrune(ctx context.Context, pruneFilters filters.Args) (*network.PruneReport, error)
```

##### Daemon.PluginGetter

PluginGetter returns current pluginStore associated with the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1530)  
**Added in:** v1.13.1

```go
func (daemon *Daemon) PluginGetter() *plugin.Store
```

##### Daemon.PluginManager

PluginManager returns current pluginManager associated with the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1525)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) PluginManager() *plugin.Manager
```

##### Daemon.ProcessClusterNotifications

ProcessClusterNotifications gets changes from store and add them to event list

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L103)  

```go
func (daemon *Daemon) ProcessClusterNotifications(ctx context.Context, watchStream chan *swarmapi.WatchMessage)
```

##### Daemon.ProcessEvent

ProcessEvent is called by libcontainerd whenever an event occurs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/monitor.go#L167)  

```go
func (daemon *Daemon) ProcessEvent(id string, e libcontainerdtypes.EventType, ei libcontainerdtypes.EventInfo) error
```

##### Daemon.RawSysInfo

RawSysInfo returns *sysinfo.SysInfo .

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1638)  

```go
func (daemon *Daemon) RawSysInfo() *sysinfo.SysInfo
```

##### Daemon.Register

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/container.go#L101)  

```go
func (daemon *Daemon) Register(c *container.Container) error
```

##### Daemon.RegistryHosts

RegistryHosts returns the registry hosts configuration for the host component
of a distribution image reference.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/hosts.go#L29)  

```go
func (daemon *Daemon) RegistryHosts(host string) ([]docker.RegistryHost, error)
```

##### Daemon.RegistryService

RegistryService returns the Daemon's RegistryService

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1625)  
**Added in:** v1.7.0

```go
func (daemon *Daemon) RegistryService() *registry.Service
```

##### Daemon.ReleaseIngress

ReleaseIngress releases the ingress networking.
The function returns a channel which will signal the caller when the programming is completed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L190)  

```go
func (daemon *Daemon) ReleaseIngress() (<-chan struct{}, error)
```

##### Daemon.Reload

Reload modifies the live daemon configuration from conf.
conf is assumed to be a validated configuration.

These are the settings that Reload changes:
- Platform runtime
- Daemon debug log level
- Daemon max concurrent downloads
- Daemon max concurrent uploads
- Daemon max download attempts
- Daemon shutdown timeout (in seconds)
- Cluster discovery (reconfigure and restart)
- Daemon labels
- Insecure registries
- Registry mirrors
- Daemon live restore

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/reload.go#L74)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) Reload(conf *config.Config) error
```

##### Daemon.RestartSwarmContainers

RestartSwarmContainers restarts any autostart container which has a
swarm endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L660)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) RestartSwarmContainers()
```

##### Daemon.SetCluster

SetCluster sets the cluster

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1511)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) SetCluster(cluster Cluster)
```

##### Daemon.SetContainerConfigReferences

SetContainerConfigReferences sets the container config references needed

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/configs.go#L8)  

```go
func (daemon *Daemon) SetContainerConfigReferences(name string, refs []*swarmtypes.ConfigReference) error
```

##### Daemon.SetContainerDependencyStore

SetContainerDependencyStore sets the dependency store backend for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/dependency.go#L8)  

```go
func (daemon *Daemon) SetContainerDependencyStore(name string, store exec.DependencyGetter) error
```

##### Daemon.SetContainerSecretReferences

SetContainerSecretReferences sets the container secret references needed

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/secrets.go#L8)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) SetContainerSecretReferences(name string, refs []*swarmtypes.SecretReference) error
```

##### Daemon.SetNetworkBootstrapKeys

SetNetworkBootstrapKeys sets the bootstrap keys.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L243)  
**Added in:** v1.12.0

```go
func (daemon *Daemon) SetNetworkBootstrapKeys(keys []*lntypes.EncryptionKey) error
```

##### Daemon.SetupIngress

SetupIngress setups ingress networking.
The function returns a channel which will signal the caller when the programming is completed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L178)  
**Added in:** v1.12.0

```go
func (daemon *Daemon) SetupIngress(create clustertypes.NetworkCreateRequest, nodeIP string) (<-chan struct{}, error)
```

##### Daemon.Shutdown

Shutdown stops the daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1270)  
**Added in:** v1.7.0

```go
func (daemon *Daemon) Shutdown(ctx context.Context) error
```

##### Daemon.ShutdownTimeout

ShutdownTimeout returns the timeout (in seconds) before containers are forcibly
killed during shutdown. The default timeout can be configured both on the daemon
and per container, and the longest timeout will be used. A grace-period of
5 seconds is added to the configured timeout.

A negative (-1) timeout means "indefinitely", which means that containers
are not forcibly killed, and the daemon shuts down after all containers exit.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1243)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) ShutdownTimeout() int
```

##### Daemon.StoreHosts

StoreHosts stores the addresses the daemon is listening on

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L163)  

```go
func (daemon *Daemon) StoreHosts(hosts []string)
```

##### Daemon.Subnets

Subnets return the IPv4 and IPv6 subnets of networks that are manager by Docker.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1387)  

```go
func (daemon *Daemon) Subnets() ([]net.IPNet, []net.IPNet)
```

##### Daemon.SubscribeToEvents

SubscribeToEvents returns the currently record of events, a channel to stream new events from, and a function to cancel the stream of events.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L82)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) SubscribeToEvents(since, until time.Time, filter filters.Args) ([]events.Message, chan interface{})
```

##### Daemon.SystemDiskUsage

SystemDiskUsage returns information about the daemon data disk usage.
Callers must not mutate contents of the returned fields.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/disk_usage.go#L106)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) SystemDiskUsage(ctx context.Context, opts system.DiskUsageOptions) (*systemtypes.DiskUsage, error)
```

##### Daemon.SystemInfo

SystemInfo returns information about the host server the daemon is running on.

The only error this should return is due to context cancellation/deadline.
Anything else should be logged and ignored because this is looking up
multiple things and is often used for debugging.
The only case valid early return is when the caller doesn't want the result anymore (ie context cancelled).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/info.go#L47)  
**Added in:** v1.7.0

```go
func (daemon *Daemon) SystemInfo(ctx context.Context) (*system.Info, error)
```

##### Daemon.SystemVersion

SystemVersion returns version information about the daemon.

The only error this should return is due to context cancellation/deadline.
Anything else should be logged and ignored because this is looking up
multiple things and is often used for debugging.
The only case valid early return is when the caller doesn't want the result anymore (ie context cancelled).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/info.go#L108)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) SystemVersion(ctx context.Context) (types.Version, error)
```

##### Daemon.Unmount

Unmount unsets the container base filesystem

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L1373)  

```go
func (daemon *Daemon) Unmount(container *container.Container) error
```

##### Daemon.UnsubscribeFromEvents

UnsubscribeFromEvents stops the event subscription for a client by closing the
channel where the daemon sends events to.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/events.go#L88)  
**Added in:** v1.10.0

```go
func (daemon *Daemon) UnsubscribeFromEvents(listener chan interface{})
```

##### Daemon.UpdateAttachment

UpdateAttachment notifies the attacher about the attachment config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L253)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) UpdateAttachment(networkName, networkID, containerID string, config *networktypes.NetworkingConfig) error
```

##### Daemon.UpdateContainerServiceConfig

UpdateContainerServiceConfig updates a service configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L471)  
**Added in:** v1.12.0

```go
func (daemon *Daemon) UpdateContainerServiceConfig(containerName string, serviceConfig *clustertypes.ServiceConfig) error
```

##### Daemon.UsesSnapshotter

UsesSnapshotter returns true if feature flag to use containerd snapshotter is enabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/daemon.go#L205)  

```go
func (daemon *Daemon) UsesSnapshotter() bool
```

##### Daemon.VolumesService

VolumesService is used to perform volume operations

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/volumes.go#L322)  

```go
func (daemon *Daemon) VolumesService() *service.VolumesService
```

##### Daemon.WaitForDetachment

WaitForDetachment makes the cluster manager wait for detachment of
the container from the network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L267)  
**Added in:** v1.13.0

```go
func (daemon *Daemon) WaitForDetachment(ctx context.Context, networkName, networkID, taskID, containerID string) error
```

---

### ImageService

ImageService is a temporary interface to assist in the migration to the
containerd image-store. This interface should not be considered stable,
and may change over time.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/image_service.go#L26)  

```go
type ImageService interface {
	PullImage(ctx context.Context, ref reference.Named, platform *ocispec.Platform, metaHeaders map[string][]string, authConfig *registry.AuthConfig, outStream io.Writer) error
	PushImage(ctx context.Context, ref reference.Named, platform *ocispec.Platform, metaHeaders map[string][]string, authConfig *registry.AuthConfig, outStream io.Writer) error
	CreateImage(ctx context.Context, config []byte, parent string, contentStoreDigest digest.Digest) (builder.Image, error)
	ImageDelete(ctx context.Context, imageRef string, options imagetype.RemoveOptions) ([]imagetype.DeleteResponse, error)
	ExportImage(ctx context.Context, names []string, platform *ocispec.Platform, outStream io.Writer) error
	LoadImage(ctx context.Context, inTar io.ReadCloser, platform *ocispec.Platform, outStream io.Writer, quiet bool) error
	Images(ctx context.Context, opts imagetype.ListOptions) ([]*imagetype.Summary, error)
	LogImageEvent(ctx context.Context, imageID, refName string, action events.Action)
	CountImages(ctx context.Context) int
	ImagesPrune(ctx context.Context, pruneFilters filters.Args) (*imagetype.PruneReport, error)
	ImportImage(ctx context.Context, ref reference.Named, platform *ocispec.Platform, msg string, layerReader io.Reader, changes []string) (image.ID, error)
	TagImage(ctx context.Context, imageID image.ID, newTag reference.Named) error
	GetImage(ctx context.Context, refOrID string, options backend.GetImageOpts) (*image.Image, error)
	ImageHistory(ctx context.Context, name string, platform *ocispec.Platform) ([]*imagetype.HistoryResponseItem, error)
	CommitImage(ctx context.Context, c backend.CommitConfig) (image.ID, error)
	SquashImage(id, parent string) (string, error)
	ImageInspect(ctx context.Context, refOrID string, opts backend.ImageInspectOpts) (*imagetype.InspectResponse, error)
	ImageDiskUsage(ctx context.Context) (int64, error)

	GetImageAndReleasableLayer(ctx context.Context, refOrID string, opts backend.GetImageAndLayerOptions) (builder.Image, builder.ROLayer, error)
	CreateLayer(container *container.Container, initFunc layer.MountInit) (container.RWLayer, error)
	CreateLayerFromImage(img *image.Image, layerName string, rwLayerOpts *layer.CreateRWLayerOpts) (container.RWLayer, error)
	GetLayerByID(cid string) (container.RWLayer, error)
	LayerStoreStatus() [][2]string
	GetLayerMountID(cid string) (string, error)
	ReleaseLayer(rwlayer container.RWLayer) error
	GetContainerLayerSize(ctx context.Context, containerID string) (int64, int64, error)
	Changes(ctx context.Context, container *container.Container) ([]archive.Change, error)

	GetLayerFolders(img *image.Image, rwLayer container.RWLayer, containerID string) ([]string, error)

	MakeImageCache(ctx context.Context, cacheFrom []string) (builder.ImageCache, error)
	CommitBuildStep(ctx context.Context, c backend.CommitConfig) (image.ID, error)

	DistributionServices() images.DistributionServices
	Children(ctx context.Context, id image.ID) ([]image.ID, error)
	Cleanup() error
	StorageDriver() string
	UpdateConfig(maxDownloads, maxUploads int)
}
```

---

### NetworkManager

NetworkManager provides methods to manage networks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster.go#L23)  

```go
type NetworkManager interface {
	GetNetwork(input string) (network.Inspect, error)
	GetNetworks(filters.Args) ([]network.Inspect, error)
	RemoveNetwork(input string) error
}
```

---

### PredefinedNetworkError

PredefinedNetworkError is returned when user tries to create predefined network that already exists.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L40)  

```go
type PredefinedNetworkError string
```

#### Methods

##### PredefinedNetworkError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L42)  

```go
func (pnr PredefinedNetworkError) Error() string
```

##### PredefinedNetworkError.Forbidden

Forbidden denotes the type of this error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/network.go#L47)  

```go
func (pnr PredefinedNetworkError) Forbidden()
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/container

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:30 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L161)

```go
const (
	// Deprecated: use [container.WaitConditionNotRunning] instead.
	WaitConditionNotRunning = container.WaitConditionNotRunning
	// Deprecated: use [container.WaitConditionNextExit] instead.
	WaitConditionNextExit = container.WaitConditionNextExit
	// Deprecated: use [container.WaitConditionRemoved] instead.
	WaitConditionRemoved = container.WaitConditionRemoved
)
```

## Variables

This section is empty.

## Functions

### IsValidHealthString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L114)  
**Added in:** v1.13.0

```go
func IsValidHealthString(s string) bool
```

---

### IsValidStateString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L152)  

```go
func IsValidStateString(s container.ContainerState) bool
```

---

### ReplaceOrAppendEnvValues

ReplaceOrAppendEnvValues returns the defaults with the overrides either
replaced by env key or appended to the list

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/env.go#L9)  

```go
func ReplaceOrAppendEnvValues(defaults, overrides []string) []string
```

---

## Types

### Container

Container holds the structure defining a container object.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L68)  

```go
type Container struct {
	StreamConfig *stream.Config
	// We embed [State] here so that Container supports states directly,
	// but marshal it as a struct in JSON.
	//
	// State also provides a [sync.Mutex] which is used as lock for both
	// the Container and State.
	*State          `json:"State"`
	Root            string  `json:"-"` // Path to the "home" of the container, including metadata.
	BaseFS          string  `json:"-"` // Path to the graphdriver mountpoint
	RWLayer         RWLayer `json:"-"`
	ID              string
	Created         time.Time
	Managed         bool
	Path            string
	Args            []string
	Config          *containertypes.Config
	ImageID         image.ID `json:"Image"`
	ImageManifest   *ocispec.Descriptor
	NetworkSettings *network.Settings
	LogPath         string
	Name            string
	Driver          string

	// Deprecated: use [ImagePlatform.OS] instead.
	// TODO: Remove, see https://github.com/moby/moby/issues/48892
	OS string

	ImagePlatform ocispec.Platform

	RestartCount             int
	HasBeenStartedBefore     bool
	HasBeenManuallyStopped   bool // used for unless-stopped restart policy
	HasBeenManuallyRestarted bool `json:"-"` // used to distinguish restart caused by restart policy from the manual one
	MountPoints              map[string]*volumemounts.MountPoint
	HostConfig               *containertypes.HostConfig `json:"-"` // do not serialize the host config in the json, otherwise we'll make the container unportable
	ExecCommands             *ExecStore                 `json:"-"`
	DependencyStore          agentexec.DependencyGetter `json:"-"`
	SecretReferences         []*swarmtypes.SecretReference
	ConfigReferences         []*swarmtypes.ConfigReference
	// logDriver for closing
	LogDriver logger.Logger  `json:"-"`
	LogCopier *logger.Copier `json:"-"`

	// Fields here are specific to Unix platforms
	SecurityOptions
	HostnamePath   string
	HostsPath      string
	ShmPath        string
	ResolvConfPath string

	// Fields here are specific to Windows
	NetworkSharedContainerID string            `json:"-"`
	SharedEndpointList       []string          `json:"-"`
	LocalLogCacheMeta        localLogCacheMeta `json:",omitempty"`
	// contains filtered or unexported fields
}
```

#### Functions

##### NewBaseContainer

NewBaseContainer creates a new container with its
basic configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L143)  

```go
func NewBaseContainer(id, root string) *Container
```

#### Methods

##### Container.AddMountPointWithVolume

AddMountPointWithVolume adds a new mount point configured with a volume to the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L542)  

```go
func (container *Container) AddMountPointWithVolume(destination string, vol volume.Volume, rw bool)
```

##### Container.AttachContext

AttachContext returns the context for attach calls to track container liveness.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L662)  

```go
func (container *Container) AttachContext() context.Context
```

##### Container.BuildHostnameFile

BuildHostnameFile writes the container's hostname file.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L52)  

```go
func (container *Container) BuildHostnameFile() error
```

##### Container.CancelAttachContext

CancelAttachContext cancels attach context. All attach calls should detach
after this call.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L668)  
**Added in:** v1.11.0

```go
func (container *Container) CancelAttachContext()
```

##### Container.CheckpointDir

CheckpointDir returns the directory checkpoints are stored in

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L427)  
**Added in:** v1.13.0

```go
func (container *Container) CheckpointDir() string
```

##### Container.CheckpointTo

CheckpointTo makes the Container's current state visible to queries, and persists state.
Callers must hold a Container lock.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L225)  

```go
func (container *Container) CheckpointTo(ctx context.Context, store *ViewDB) error
```

##### Container.CloseStreams

CloseStreams closes the container's stdio streams

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L706)  
**Added in:** v1.12.5

```go
func (container *Container) CloseStreams() error
```

##### Container.CommitInMemory

CommitInMemory makes the Container's current state visible to queries,
but does not persist state.

Callers must hold a Container lock.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L301)  

```go
func (container *Container) CommitInMemory(store *ViewDB) error
```

##### Container.ConfigFilePath

ConfigFilePath returns the path to the on-disk location of a config.
On unix, configs are always considered secret

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L441)  

```go
func (container *Container) ConfigFilePath(configRef swarmtypes.ConfigReference) (string, error)
```

##### Container.ConfigPath

ConfigPath returns the path to the container's JSON config

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L422)  

```go
func (container *Container) ConfigPath() (string, error)
```

##### Container.CopyImagePathContent

CopyImagePathContent copies files in destination to the volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L126)  

```go
func (container *Container) CopyImagePathContent(volumePath, destination string) error
```

##### Container.CreateDaemonEnvironment

CreateDaemonEnvironment creates a new environment variable slice for this container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L769)  

```go
func (container *Container) CreateDaemonEnvironment(tty bool, linkedEnv []string) []string
```

##### Container.DetachAndUnmount

DetachAndUnmount uses a detached mount on all mount destinations, then
unmounts each volume normally.
This is used from daemon/archive for `docker cp`

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L332)  
**Added in:** v1.13.0

```go
func (container *Container) DetachAndUnmount(volumeEventLog func(name string, action events.Action, attributes map[string]string)) error
```

##### Container.ExitOnNext

ExitOnNext signals to the monitor that it should not restart the container
after we send the kill signal.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L412)  

```go
func (container *Container) ExitOnNext()
```

##### Container.FromDisk

FromDisk loads the container configuration stored in the host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L156)  

```go
func (container *Container) FromDisk() error
```

##### Container.FullHostname

FullHostname returns hostname and optional domain appended to it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L634)  
**Added in:** v1.11.0

```go
func (container *Container) FullHostname() string
```

##### Container.GetExecIDs

GetExecIDs returns the list of exec commands running on the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L530)  

```go
func (container *Container) GetExecIDs() []string
```

##### Container.GetMountLabel

GetMountLabel returns the mounting label for the container.
This label is empty if the container is privileged.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L525)  

```go
func (container *Container) GetMountLabel() string
```

##### Container.GetMountPoints

GetMountPoints gives a platform specific transformation to types.MountPoint. Callers must hold a Container lock.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L422)  

```go
func (container *Container) GetMountPoints() []containertypes.MountPoint
```

##### Container.GetProcessLabel

GetProcessLabel returns the process label for the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L514)  

```go
func (container *Container) GetProcessLabel() string
```

##### Container.GetResourcePath

GetResourcePath evaluates `path` in the scope of the container's BaseFS, with proper path
sanitization. Symlinks are all scoped to the BaseFS of the container, as
though the container's BaseFS was `/`.

The BaseFS of a container is the host-facing path which is bind-mounted as
`/` inside the container. This method is essentially used to access a
particular path inside the container as though you were a process in that
container.

# NOTE
The returned path is *only* safely scoped inside the container's BaseFS
if no component of the returned path changes (such as a component
symlinking to a different path) between using this method and using the
path. See symlink.FollowSymlinkInScope for more details.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L361)  

```go
func (container *Container) GetResourcePath(path string) (string, error)
```

##### Container.GetRootResourcePath

GetRootResourcePath evaluates `path` in the scope of the container's root, with proper path
sanitization. Symlinks are all scoped to the root of the container, as
though the container's root was `/`.

The root of a container is the host-facing configuration metadata directory.
Only use this method to safely access the container's `container.json` or
other metadata files. If in doubt, use container.GetResourcePath.

# NOTE
The returned path is *only* safely scoped inside the container's root
if no component of the returned path changes (such as a component
symlinking to a different path) between using this method and using the
path. See symlink.FollowSymlinkInScope for more details.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L403)  

```go
func (container *Container) GetRootResourcePath(path string) (string, error)
```

##### Container.GetRunningTask

GetRunningTask asserts that the container is running and returns the Task for
the container. An errdefs.Conflict error is returned if the container is not
in the Running state.

A system error is returned if container is in a bad state: Running is true
but has a nil Task.

The container lock must be held when calling this method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L832)  

```go
func (container *Container) GetRunningTask() (libcontainerdtypes.Task, error)
```

##### Container.HasMountFor

HasMountFor checks if path is a mountpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L139)  

```go
func (container *Container) HasMountFor(path string) bool
```

##### Container.HostConfigPath

HostConfigPath returns the path to the container's JSON hostconfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L417)  

```go
func (container *Container) HostConfigPath() (string, error)
```

##### Container.InitDNSHostConfig

InitDNSHostConfig ensures that the dns fields are never nil.
New containers don't ever have those fields nil,
but pre created containers can still have those nil values.
The non-recommended host configuration in the start api can
make these fields nil again, this corrects that issue until
we remove that behavior for good.
See https://github.com/docker/docker/pull/17779
for a more detailed explanation on why we don't want that.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L612)  

```go
func (container *Container) InitDNSHostConfig()
```

##### Container.InitializeStdio

InitializeStdio is called by libcontainerd to connect the stdio.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L711)  
**Added in:** v1.12.4

```go
func (container *Container) InitializeStdio(iop *cio.DirectIO) (cio.IO, error)
```

##### Container.IpcMounts

IpcMounts returns the list of IPC mounts

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L177)  

```go
func (container *Container) IpcMounts() []Mount
```

##### Container.IsDestinationMounted

IsDestinationMounted checks whether a path is mounted on the container or not.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L579)  

```go
func (container *Container) IsDestinationMounted(destination string) bool
```

##### Container.MountsResourcePath

MountsResourcePath returns the path where mounts are stored for the given mount

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L731)  

```go
func (container *Container) MountsResourcePath(mount string) (string, error)
```

##### Container.NetworkMounts

NetworkMounts returns the list of network mounts.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L62)  

```go
func (container *Container) NetworkMounts() []Mount
```

##### Container.Reset

Reset puts a container into a state where it can be restarted again.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/monitor.go#L15)  
**Added in:** v1.11.0

```go
func (container *Container) Reset(lock bool)
```

##### Container.ResetRestartManager

ResetRestartManager initializes new restartmanager based on container config

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L651)  
**Added in:** v1.13.0

```go
func (container *Container) ResetRestartManager(resetCount bool)
```

##### Container.RestartManager

RestartManager returns the current restartmanager instance connected to container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L643)  
**Added in:** v1.11.0

```go
func (container *Container) RestartManager() *restartmanager.RestartManager
```

##### Container.RestoreTask

RestoreTask restores the containerd container and task handles and reattaches
the IO for the running task. Container state is not synced with containerd's
state.

An errdefs.NotFound error is returned if the container does not exist in
containerd. However, a nil error is returned if the task does not exist in
containerd.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L809)  

```go
func (container *Container) RestoreTask(ctx context.Context, client libcontainerdtypes.Client) error
```

##### Container.SecretFilePath

SecretFilePath returns the path to the location of a secret on the host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L741)  

```go
func (container *Container) SecretFilePath(secretRef swarmtypes.SecretReference) (string, error)
```

##### Container.SecretMountPath

SecretMountPath returns the path of the secret mount for the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L736)  
**Added in:** v1.13.0

```go
func (container *Container) SecretMountPath() (string, error)
```

##### Container.SecretMounts

SecretMounts returns the mounts for the secret path.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L200)  

```go
func (container *Container) SecretMounts() ([]Mount, error)
```

##### Container.SetupWorkingDirectory

SetupWorkingDirectory sets up the container's working directory as set in container.Config.WorkingDir

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L324)  

```go
func (container *Container) SetupWorkingDirectory(uid int, gid int) error
```

##### Container.ShmResourcePath

ShmResourcePath returns path to shm

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L134)  

```go
func (container *Container) ShmResourcePath() (string, error)
```

##### Container.ShouldRestart

ShouldRestart decides whether the daemon should restart the container or not.
This is based on the container's restart policy.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L536)  

```go
func (container *Container) ShouldRestart() bool
```

##### Container.StartLogger

StartLogger starts a new logger driver for the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L432)  

```go
func (container *Container) StartLogger() (logger.Logger, error)
```

##### Container.StderrPipe

StderrPipe gets the stderr stream of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L701)  
**Added in:** v1.12.5

```go
func (container *Container) StderrPipe() io.ReadCloser
```

##### Container.StdinPipe

StdinPipe gets the stdin stream of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L691)  
**Added in:** v1.12.5

```go
func (container *Container) StdinPipe() io.WriteCloser
```

##### Container.StdoutPipe

StdoutPipe gets the stdout stream of the container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L696)  
**Added in:** v1.12.5

```go
func (container *Container) StdoutPipe() io.ReadCloser
```

##### Container.StopSignal

StopSignal returns the signal used to stop the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L584)  

```go
func (container *Container) StopSignal() syscall.Signal
```

##### Container.StopTimeout

StopTimeout returns the timeout (in seconds) used to stop the container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L597)  
**Added in:** v1.13.0

```go
func (container *Container) StopTimeout() int
```

##### Container.TmpfsMounts

TmpfsMounts returns the list of tmpfs mounts

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L395)  

```go
func (container *Container) TmpfsMounts() ([]Mount, error)
```

##### Container.TrySetNetworkMount

TrySetNetworkMount attempts to set the network mounts given a provided destination and
the path to use for it; return true if the given destination was a network mount file

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L34)  

```go
func (container *Container) TrySetNetworkMount(destination string, path string) bool
```

##### Container.UnmountIpcMount

UnmountIpcMount unmounts shm if it was mounted

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L156)  

```go
func (container *Container) UnmountIpcMount() error
```

##### Container.UnmountSecrets

UnmountSecrets unmounts the local tmpfs for secrets

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L232)  
**Added in:** v1.13.0

```go
func (container *Container) UnmountSecrets() error
```

##### Container.UnmountVolumes

UnmountVolumes unmounts all volumes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L556)  

```go
func (container *Container) UnmountVolumes(ctx context.Context, volumeEventLog func(name string, action events.Action, attributes map[string]string)) error
```

##### Container.UpdateContainer

UpdateContainer updates configuration of a container. Callers must hold a Lock on the Container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container_unix.go#L248)  

```go
func (container *Container) UpdateContainer(hostConfig *containertypes.HostConfig) error
```

##### Container.UpdateMonitor

UpdateMonitor updates monitor configure for running container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L629)  
**Added in:** v1.11.0

```go
func (container *Container) UpdateMonitor(restartPolicy containertypes.RestartPolicy)
```

##### Container.WriteHostConfig

WriteHostConfig saves the host configuration on disk for the container,
and returns a deep copy of the saved object. Callers must hold a Container lock.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L269)  

```go
func (container *Container) WriteHostConfig() (*containertypes.HostConfig, error)
```

---

### ExecConfig

ExecConfig holds the configurations for execs. The Daemon keeps
track of both running and finished execs so that they can be
examined both during and after completion.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L18)  

```go
type ExecConfig struct {
	sync.Mutex
	Started      chan struct{}
	StreamConfig *stream.Config
	ID           string
	Running      bool
	ExitCode     *int
	OpenStdin    bool
	OpenStderr   bool
	OpenStdout   bool
	CanRemove    bool
	Container    *Container
	DetachKeys   []byte
	Entrypoint   string
	Args         []string
	Tty          bool
	Privileged   bool
	User         string
	WorkingDir   string
	Env          []string
	Process      types.Process
	ConsoleSize  *[2]uint
}
```

#### Functions

##### NewExecConfig

NewExecConfig initializes the a new exec configuration

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L43)  

```go
func NewExecConfig(c *Container) *ExecConfig
```

#### Methods

##### ExecConfig.CloseStreams

CloseStreams closes the stdio streams for the exec

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L68)  

```go
func (c *ExecConfig) CloseStreams() error
```

##### ExecConfig.InitializeStdio

InitializeStdio is called by libcontainerd to connect the stdio.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L53)  

```go
func (c *ExecConfig) InitializeStdio(iop *cio.DirectIO) (cio.IO, error)
```

##### ExecConfig.SetExitCode

SetExitCode sets the exec config's exit code

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L73)  

```go
func (c *ExecConfig) SetExitCode(code int)
```

---

### ExecStore

ExecStore keeps track of the exec configurations.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L78)  

```go
type ExecStore struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewExecStore

NewExecStore initializes a new exec store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L84)  

```go
func NewExecStore() *ExecStore
```

#### Methods

##### ExecStore.Add

Add adds a new exec configuration to the store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L102)  

```go
func (e *ExecStore) Add(id string, Config *ExecConfig)
```

##### ExecStore.Commands

Commands returns the exec configurations in the store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L91)  

```go
func (e *ExecStore) Commands() map[string]*ExecConfig
```

##### ExecStore.Delete

Delete removes an exec configuration from the store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L117)  

```go
func (e *ExecStore) Delete(id string)
```

##### ExecStore.Get

Get returns an exec configuration by its id.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L109)  

```go
func (e *ExecStore) Get(id string) *ExecConfig
```

##### ExecStore.List

List returns the list of exec ids in the store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/exec.go#L124)  

```go
func (e *ExecStore) List() []string
```

---

### ExitStatus

ExitStatus provides exit reasons for a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L59)  
**Added in:** v1.11.0

```go
type ExitStatus struct {
	// The exit code with which the container exited.
	ExitCode int

	// Time at which the container died
	ExitedAt time.Time
}
```

---

### Health

Health holds the current container health-check state

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/health.go#L12)  
**Added in:** v1.12.0

```go
type Health struct {
	container.Health
	// contains filtered or unexported fields
}
```

#### Methods

##### Health.CloseMonitorChannel

CloseMonitorChannel closes any existing monitor channel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/health.go#L71)  
**Added in:** v1.12.0

```go
func (s *Health) CloseMonitorChannel()
```

##### Health.OpenMonitorChannel

OpenMonitorChannel creates and returns a new monitor channel. If there
already is one, it returns nil.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/health.go#L58)  
**Added in:** v1.12.0

```go
func (s *Health) OpenMonitorChannel() chan struct{}
```

##### Health.SetStatus

SetStatus writes the current status to the underlying health structure,
obeying the locking semantics.

Status may be set directly if another lock is used.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/health.go#L49)  

```go
func (s *Health) SetStatus(new container.HealthStatus)
```

##### Health.Status

Status returns the current health status.

Note that this takes a lock and the value may change after being read.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/health.go#L33)  

```go
func (s *Health) Status() container.HealthStatus
```

##### Health.String

String returns a human-readable description of the health-check state

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/health.go#L19)  
**Added in:** v1.12.0

```go
func (s *Health) String() string
```

---

### History

History is a convenience type for storing a list of containers,
sorted by creation date in descendant order.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/history.go#L7)  

```go
type History []*Container
```

#### Methods

##### History.Len

Len returns the number of containers in the history.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/history.go#L10)  

```go
func (history *History) Len() int
```

##### History.Less

Less compares two containers and returns true if the second one
was created before the first one.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/history.go#L16)  

```go
func (history *History) Less(i, j int) bool
```

##### History.Swap

Swap switches containers i and j positions in the history.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/history.go#L22)  

```go
func (history *History) Swap(i, j int)
```

---

### Mount

Mount contains information for a mount operation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/mounts_unix.go#L6)  
**Added in:** v1.11.0

```go
type Mount struct {
	Source                 string `json:"source"`
	Destination            string `json:"destination"`
	Writable               bool   `json:"writable"`
	Data                   string `json:"data"`
	Propagation            string `json:"mountpropagation"`
	NonRecursive           bool   `json:"nonrecursive"`
	ReadOnlyNonRecursive   bool   `json:"readonlynonrecursive"`
	ReadOnlyForceRecursive bool   `json:"readonlyforcerecursive"`
}
```

---

### RWLayer

RWLayer represents a writable layer for a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/rwlayer.go#L4)  

```go
type RWLayer interface {
	// Mount mounts the RWLayer and returns the filesystem path
	// to the writable layer.
	Mount(mountLabel string) (string, error)

	// Unmount unmounts the RWLayer. This should be called
	// for every mount. If there are multiple mount calls
	// this operation will only decrement the internal mount counter.
	Unmount() error

	// Metadata returns the low level metadata for the mutable layer
	Metadata() (map[string]string, error)
}
```

---

### SecurityOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/container.go#L127)  

```go
type SecurityOptions struct {
	// MountLabel contains the options for the "mount" command.
	MountLabel      string
	ProcessLabel    string
	AppArmorProfile string
	SeccompProfile  string
	NoNewPrivileges bool
	WritableCgroups *bool
}
```

---

### Snapshot

Snapshot is a read only view for Containers. It holds all information necessary to serve container queries in a
versioned ACID in-memory store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L33)  

```go
type Snapshot struct {
	container.Summary

	// additional info queries need to filter on
	// preserve nanosec resolution for queries
	CreatedAt    time.Time
	StartedAt    time.Time
	Name         string
	Pid          int
	ExitCode     int
	Running      bool
	Paused       bool
	Managed      bool
	ExposedPorts nat.PortSet
	PortBindings nat.PortSet
	Health       container.HealthStatus
	HostConfig   struct {
		Isolation string
	}
}
```

---

### State

State holds the current container state, and has methods to get and
set the state. State is embedded in the Container struct.

State contains an exported sync.Mutex which is used as a global lock
for both the State and the Container it's embedded in.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L20)  

```go
type State struct {
	// This Mutex is exported by design and is used as a global lock
	// for both the State and the Container it's embedded in.
	sync.Mutex
	// Note that [State.Running], [State.Restarting], and [State.Paused] are
	// not mutually exclusive.
	//
	// When pausing a container (on Linux), the freezer cgroup is used to suspend
	// all processes in the container. Freezing the process requires the process to
	// be running. As a result, paused containers can have both [State.Running]
	// and [State.Paused] set to true.
	//
	// In a similar fashion, [State.Running] and [State.Restarting] can both
	// be true in a situation where a container is in process of being restarted.
	// Refer to [State.StateString] for order of precedence.
	Running           bool
	Paused            bool
	Restarting        bool
	OOMKilled         bool
	RemovalInProgress bool `json:"-"` // No need for this to be persistent on disk.
	Dead              bool
	Pid               int
	ExitCodeValue     int    `json:"ExitCode"`
	ErrorMsg          string `json:"Error"` // contains last known error during container start, stop, or remove
	StartedAt         time.Time
	FinishedAt        time.Time
	Health            *Health
	Removed           bool `json:"-"`
	// contains filtered or unexported fields
}
```

#### Functions

##### NewState

NewState creates a default state object.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L70)  

```go
func NewState() *State
```

#### Methods

##### State.C8dContainer

C8dContainer returns a reference to the libcontainerd Container object for
the container and whether the reference is valid.

The container lock must be held when calling this method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L450)  

```go
func (s *State) C8dContainer() (_ libcontainerdtypes.Container, ok bool)
```

##### State.Err

Err returns an error if there is one.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L430)  

```go
func (s *State) Err() error
```

##### State.ExitCode

ExitCode returns current exitcode for the state. Take lock before if state
may be shared.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L257)  

```go
func (s *State) ExitCode() int
```

##### State.GetPID

GetPID holds the process id of a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L249)  

```go
func (s *State) GetPID() int
```

##### State.IsDead

IsDead returns whether the Dead flag is set. Used by Container to check whether a container is dead.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L407)  

```go
func (s *State) IsDead() bool
```

##### State.IsPaused

IsPaused returns whether the container is paused.

Note that [State.Running], [State.Restarting], and [State.Paused] are
not mutually exclusive.

When pausing a container (on Linux), the freezer cgroup is used to suspend
all processes in the container. Freezing the process requires the process to
be running. As a result, paused containers can have both [State.Running]
and [State.Paused] set to true.

In a similar fashion, [State.Running] and [State.Restarting] can both
be true in a situation where a container is in process of being restarted.
Refer to State.StateString for order of precedence.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L354)  

```go
func (s *State) IsPaused() bool
```

##### State.IsRemovalInProgress

IsRemovalInProgress returns whether the RemovalInProgress flag is set.
Used by Container to check whether a container is being removed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L400)  

```go
func (s *State) IsRemovalInProgress() bool
```

##### State.IsRestarting

IsRestarting returns whether the container is restarting.

Note that [State.Running], [State.Restarting], and [State.Paused] are
not mutually exclusive.

When pausing a container (on Linux), the freezer cgroup is used to suspend
all processes in the container. Freezing the process requires the process to
be running. As a result, paused containers can have both [State.Running]
and [State.Paused] set to true.

In a similar fashion, [State.Running] and [State.Restarting] can both
be true in a situation where a container is in process of being restarted.
Refer to State.StateString for order of precedence.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L373)  

```go
func (s *State) IsRestarting() bool
```

##### State.IsRunning

IsRunning returns whether the [State.Running] flag is set.

Note that [State.Running], [State.Restarting], and [State.Paused] are
not mutually exclusive.

When pausing a container (on Linux), the freezer cgroup is used to suspend
all processes in the container. Freezing the process requires the process to
be running. As a result, paused containers can have both [State.Running]
and [State.Paused] set to true.

In a similar fashion, [State.Running] and [State.Restarting] can both
be true in a situation where a container is in process of being restarted.
Refer to State.StateString for order of precedence.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L242)  

```go
func (s *State) IsRunning() bool
```

##### State.ResetRemovalInProgress

ResetRemovalInProgress makes the RemovalInProgress state to false.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L392)  

```go
func (s *State) ResetRemovalInProgress()
```

##### State.SetError

SetError sets the container's error state. This is useful when we want to
know the error that occurred when container transits to another state
when inspecting it

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L334)  

```go
func (s *State) SetError(err error)
```

##### State.SetExitCode

SetExitCode sets current exitcode for the state. Take lock before if state
may be shared.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L263)  
**Added in:** v1.12.0

```go
func (s *State) SetExitCode(ec int)
```

##### State.SetRemovalError

SetRemovalError is to be called in case a container remove failed.
It sets an error and notifies all waiters.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L420)  

```go
func (s *State) SetRemovalError(err error)
```

##### State.SetRemovalInProgress

SetRemovalInProgress sets the container state as being removed.
It returns true if the container was already in that state.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L381)  

```go
func (s *State) SetRemovalInProgress() bool
```

##### State.SetRemoved

SetRemoved assumes this container is already in the "dead" state and notifies all waiters.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L414)  

```go
func (s *State) SetRemoved()
```

##### State.SetRestarting

SetRestarting sets the container state to "restarting" without locking.
It also sets the container PID to 0.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L318)  

```go
func (s *State) SetRestarting(exitStatus *ExitStatus)
```

##### State.SetRunning

SetRunning sets the running state along with StartedAt time.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L268)  

```go
func (s *State) SetRunning(ctr libcontainerdtypes.Container, tsk libcontainerdtypes.Task, start time.Time)
```

##### State.SetRunningExternal

SetRunningExternal sets the running state without setting the `StartedAt` time (used for containers not started by Docker instead of SetRunning).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L273)  

```go
func (s *State) SetRunningExternal(ctr libcontainerdtypes.Container, tsk libcontainerdtypes.Task)
```

##### State.SetStopped

SetStopped sets the container state to "stopped" without locking.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L301)  

```go
func (s *State) SetStopped(exitStatus *ExitStatus)
```

##### State.StateString

StateString returns the container's current [ContainerState], based on the
[State.Running], [State.Paused], [State.Restarting], [State.RemovalInProgress],
[State.StartedAt] and [State.Dead] fields.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L121)  

```go
func (s *State) StateString() container.ContainerState
```

##### State.String

String returns a human-readable description of the state

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L75)  

```go
func (s *State) String() string
```

##### State.Task

Task returns a reference to the libcontainerd Task object for the container
and whether the reference is valid.

The container lock must be held when calling this method.

See also: (*Container).GetRunningTask().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L460)  

```go
func (s *State) Task() (_ libcontainerdtypes.Task, ok bool)
```

##### State.Wait

Wait waits until the container is in a certain state indicated by the given
condition. A context must be used for cancelling the request, controlling
timeouts, and avoiding goroutine leaks. Wait must be called without holding
the state lock. Returns a channel from which the caller will receive the
result. If the container exited on its own, the result's Err() method will
be nil and its ExitCode() method will return the container's exit code,
otherwise, the results Err() method will return an error indicating why the
wait operation failed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L178)  

```go
func (s *State) Wait(ctx context.Context, condition container.WaitCondition) <-chan container.StateStatus
```

---

### StateStatus

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L67)  
**Added in:** v1.13.0

```go
type StateStatus = container.StateStatus
```

---

### Store

Store defines an interface that
any container store must implement.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/store.go#L13)  

```go
type Store interface {
	// Add appends a new container to the store.
	Add(string, *Container)
	// Get returns a container from the store by the identifier it was stored with.
	Get(string) *Container
	// Delete removes a container from the store by the identifier it was stored with.
	Delete(string)
	// List returns a list of containers from the store.
	List() []*Container
	// Size returns the number of containers in the store.
	Size() int
	// First returns the first container found in the store by a given filter.
	First(StoreFilter) *Container
	// ApplyAll calls the reducer function with every container in the store.
	ApplyAll(StoreReducer)
}
```

#### Functions

##### NewMemoryStore

NewMemoryStore initializes a new memory store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/memory_store.go#L14)  

```go
func NewMemoryStore() Store
```

---

### StoreFilter

StoreFilter defines a function to filter
container in the store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/store.go#L5)  

```go
type StoreFilter func(*Container) bool
```

---

### StoreReducer

StoreReducer defines a function to
manipulate containers in the store

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/store.go#L9)  

```go
type StoreReducer func(*Container)
```

---

### View

View provides a consistent read-only view of the database.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L208)  

```go
type View struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### View.All

All returns a all items in this snapshot. Returned objects must never be modified.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L213)  

```go
func (v *View) All() ([]Snapshot, error)
```

##### View.Get

Get returns an item by id. Returned objects must never be modified.
It returns an errdefs.NotFound if the given id was not found.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L232)  

```go
func (v *View) Get(id string) (*Snapshot, error)
```

##### View.GetAllNames

GetAllNames returns all registered names.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L276)  

```go
func (v *View) GetAllNames() map[string][]string
```

##### View.GetID

GetID returns the container ID that the passed in name is reserved to.
It returns an errdefs.NotFound if the given id was not found.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L264)  

```go
func (v *View) GetID(name string) (string, error)
```

---

### ViewDB

ViewDB provides an in-memory transactional (ACID) container store.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L93)  

```go
type ViewDB struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewViewDB

NewViewDB provides the default implementation, with the default schema

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L98)  

```go
func NewViewDB() (*ViewDB, error)
```

#### Methods

##### ViewDB.Delete

Delete removes an item by ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L164)  

```go
func (db *ViewDB) Delete(c *Container) error
```

##### ViewDB.GetByPrefix

GetByPrefix returns a container with the given ID prefix. It returns an
error if an empty prefix was given or if multiple containers match the prefix.
It returns an errdefs.NotFound if the given s yielded no results.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L109)  

```go
func (db *ViewDB) GetByPrefix(s string) (string, error)
```

##### ViewDB.ReleaseName

ReleaseName releases the reserved name
Once released, a name can be reserved again

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L201)  

```go
func (db *ViewDB) ReleaseName(name string) error
```

##### ViewDB.ReserveName

ReserveName registers a container ID to a name. ReserveName is idempotent,
but returns an errdefs.Conflict when attempting to reserve a container ID
to a name that already is reserved.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L183)  

```go
func (db *ViewDB) ReserveName(name, containerID string) error
```

##### ViewDB.Save

Save atomically updates the in-memory store state for a Container.
Only read only (deep) copies of containers may be passed in.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L157)  

```go
func (db *ViewDB) Save(c *Container) error
```

##### ViewDB.Snapshot

Snapshot provides a consistent read-only view of the database.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/view.go#L138)  

```go
func (db *ViewDB) Snapshot() *View
```

---

### WaitCondition

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/container/state.go#L159)  

```go
type WaitCondition = container.WaitCondition
```

---


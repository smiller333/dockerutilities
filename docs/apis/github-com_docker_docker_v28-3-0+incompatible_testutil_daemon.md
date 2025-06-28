# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil/daemon

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:56 UTC

## Overview

Package daemon launches dockerd for testing purposes.


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L13)

```go
const (
	// DefaultSwarmPort is the default port use for swarm in the tests
	DefaultSwarmPort = 2477
)
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L61)

```go
var SockRoot = filepath.Join(os.TempDir(), "docker-integration")
```

## Functions

### ScanLogsMatchAll

ScanLogsMatchAll returns a function that can be used to scan the daemon logs until *all* the passed in strings are matched

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L375)  

```go
func ScanLogsMatchAll(contains ...string) func(string) bool
```

---

### ScanLogsMatchCount

ScanLogsMatchCount returns a function that can be used to scan the daemon logs until the passed in matcher function matches `count` times

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L364)  

```go
func ScanLogsMatchCount(f func(string) bool, count int) func(string) bool
```

---

### ScanLogsMatchString

ScanLogsMatchString returns a function that can be used to scan the daemon logs for the passed in string (`contains`).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L357)  

```go
func ScanLogsMatchString(contains string) func(string) bool
```

---

### SignalDaemonDump

SignalDaemonDump sends a signal to the daemon to write a dump file

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon_unix.go#L24)  

```go
func SignalDaemonDump(pid int)
```

---

## Types

### ConfigConstructor

ConfigConstructor defines a swarm config constructor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/config.go#L12)  

```go
type ConfigConstructor func(*swarm.Config)
```

---

### Daemon

Daemon represents a Docker daemon for the testing framework

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L70)  

```go
type Daemon struct {
	Root              string
	Folder            string
	Wait              chan error
	UseDefaultHost    bool
	UseDefaultTLSHost bool

	ResolvConfPathOverride string // Path to a replacement for "/etc/resolv.conf", or empty.

	SwarmPort       int // FIXME(vdemeester) should probably not be exported
	DefaultAddrPool []string
	SubnetSize      uint32
	DataPathPort    uint32
	OOMScoreAdjust  int
	// cached information
	CachedInfo system.Info
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New returns a Daemon instance to be used for testing.
This will create a directory such as d123456789 in the folder specified by
$DOCKER_INTEGRATION_DAEMON_DEST or $DEST.
The daemon will not automatically start.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L214)  

```go
func New(t testing.TB, ops ...Option) *Daemon
```

##### NewDaemon

NewDaemon returns a Daemon instance to be used for testing.
The daemon will not automatically start.
The daemon will modify and create files under workingDir.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L113)  

```go
func NewDaemon(workingDir string, ops ...Option) (*Daemon, error)
```

#### Methods

##### Daemon.ActiveContainers

ActiveContainers returns the list of ids of the currently running containers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/container.go#L12)  

```go
func (d *Daemon) ActiveContainers(ctx context.Context, t testing.TB) []string
```

##### Daemon.BinaryPath

BinaryPath returns the binary and its arguments.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L251)  

```go
func (d *Daemon) BinaryPath() (string, error)
```

##### Daemon.CgroupNamespace

CgroupNamespace returns the cgroup namespace the daemon is running in

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon_linux.go#L32)  

```go
func (d *Daemon) CgroupNamespace(t testing.TB) string
```

##### Daemon.Cleanup

Cleanup cleans the daemon files : exec root (network namespaces, ...), swarmkit files

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L321)  

```go
func (d *Daemon) Cleanup(t testing.TB)
```

##### Daemon.ContainersNamespace

ContainersNamespace returns the containerd namespace used for containers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L260)  

```go
func (d *Daemon) ContainersNamespace() string
```

##### Daemon.CreateConfig

CreateConfig creates a config given the specified spec

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/config.go#L15)  

```go
func (d *Daemon) CreateConfig(t testing.TB, configSpec swarm.ConfigSpec) string
```

##### Daemon.CreateSecret

CreateSecret creates a secret given the specified spec

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/secret.go#L15)  

```go
func (d *Daemon) CreateSecret(t testing.TB, secretSpec swarm.SecretSpec) string
```

##### Daemon.CreateService

CreateService creates a swarm service given the specified service constructor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/service.go#L35)  

```go
func (d *Daemon) CreateService(ctx context.Context, t testing.TB, f ...ServiceConstructor) string
```

##### Daemon.DeleteConfig

DeleteConfig removes the swarm config identified by the specified id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/config.go#L48)  

```go
func (d *Daemon) DeleteConfig(t testing.TB, id string)
```

##### Daemon.DeleteSecret

DeleteSecret removes the swarm secret identified by the specified id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/secret.go#L49)  

```go
func (d *Daemon) DeleteSecret(t testing.TB, id string)
```

##### Daemon.DumpStackAndQuit

DumpStackAndQuit sends SIGQUIT to the daemon, which triggers it to dump its
stack to its log file and exit
This is used primarily for gathering debug information on test timeout

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L680)  

```go
func (d *Daemon) DumpStackAndQuit()
```

##### Daemon.FindContainerIP

FindContainerIP returns the ip of the specified container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/container.go#L28)  

```go
func (d *Daemon) FindContainerIP(t testing.TB, id string) string
```

##### Daemon.FirewallBackendDriver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L991)  

```go
func (d *Daemon) FirewallBackendDriver(t testing.TB) string
```

##### Daemon.FirewallReloadedAt

FirewallReloadedAt fetches the daemon's Info and, if it contains a firewall
reload time, returns that time.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L1000)  

```go
func (d *Daemon) FirewallReloadedAt(t testing.TB) string
```

##### Daemon.GetConfig

GetConfig returns a swarm config identified by the specified id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/config.go#L37)  

```go
func (d *Daemon) GetConfig(t testing.TB, id string) *swarm.Config
```

##### Daemon.GetNode

GetNode returns a swarm node identified by the specified id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/node.go#L17)  

```go
func (d *Daemon) GetNode(ctx context.Context, t testing.TB, id string, errCheck ...func(error) bool) *swarm.Node
```

##### Daemon.GetSecret

GetSecret returns a swarm secret identified by the specified id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/secret.go#L38)  

```go
func (d *Daemon) GetSecret(t testing.TB, id string) *swarm.Secret
```

##### Daemon.GetService

GetService returns the swarm service corresponding to the specified id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/service.go#L41)  

```go
func (d *Daemon) GetService(ctx context.Context, t testing.TB, id string) *swarm.Service
```

##### Daemon.GetServiceTasks

GetServiceTasks returns the swarm tasks for the specified service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/service.go#L52)  

```go
func (d *Daemon) GetServiceTasks(ctx context.Context, t testing.TB, service string, additionalFilters ...filters.KeyValuePair) []swarm.Task
```

##### Daemon.GetSwarm

GetSwarm returns the current swarm object

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L150)  

```go
func (d *Daemon) GetSwarm(t testing.TB) swarm.Swarm
```

##### Daemon.GetTask

GetTask returns the swarm task identified by the specified id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/service.go#L110)  

```go
func (d *Daemon) GetTask(ctx context.Context, t testing.TB, id string) swarm.Task
```

##### Daemon.ID

ID returns the generated id of the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L270)  

```go
func (d *Daemon) ID() string
```

##### Daemon.Info

Info returns the info struct for this daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L982)  

```go
func (d *Daemon) Info(t testing.TB) system.Info
```

##### Daemon.Interrupt

Interrupt stops the daemon by sending it an Interrupt signal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L665)  

```go
func (d *Daemon) Interrupt() error
```

##### Daemon.JoinTokens

JoinTokens returns the current swarm join tokens

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L194)  

```go
func (d *Daemon) JoinTokens(t testing.TB) swarm.JoinTokens
```

##### Daemon.Kill

Kill will send a SIGKILL to the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L634)  

```go
func (d *Daemon) Kill() error
```

##### Daemon.ListConfigs

ListConfigs returns the list of the current swarm configs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/config.go#L26)  

```go
func (d *Daemon) ListConfigs(t testing.TB) []swarm.Config
```

##### Daemon.ListNodes

ListNodes returns the list of the current swarm nodes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/node.go#L71)  

```go
func (d *Daemon) ListNodes(ctx context.Context, t testing.TB) []swarm.Node
```

##### Daemon.ListSecrets

ListSecrets returns the list of the current swarm secrets

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/secret.go#L27)  

```go
func (d *Daemon) ListSecrets(t testing.TB) []swarm.Secret
```

##### Daemon.ListServices

ListServices returns the list of the current swarm services

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/service.go#L99)  

```go
func (d *Daemon) ListServices(ctx context.Context, t testing.TB) []swarm.Service
```

##### Daemon.LoadBusybox

LoadBusybox image into the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L867)  

```go
func (d *Daemon) LoadBusybox(ctx context.Context, t testing.TB)
```

##### Daemon.LoadImage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L871)  

```go
func (d *Daemon) LoadImage(ctx context.Context, t testing.TB, img string)
```

##### Daemon.LogFileName

LogFileName returns the path the daemon's log file

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L289)  

```go
func (d *Daemon) LogFileName() string
```

##### Daemon.NewClient

NewClient creates new client based on daemon's socket path

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L310)  

```go
func (d *Daemon) NewClient(extraOpts ...client.Opt) (*client.Client, error)
```

##### Daemon.NewClientT

NewClientT creates new client based on daemon's socket path

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L300)  

```go
func (d *Daemon) NewClientT(t testing.TB, extraOpts ...client.Opt) *client.Client
```

##### Daemon.NodeID

NodeID returns the swarm mode node ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L76)  

```go
func (d *Daemon) NodeID() string
```

##### Daemon.Pid

Pid returns the pid of the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L660)  

```go
func (d *Daemon) Pid() int
```

##### Daemon.PluginIsNotPresent

PluginIsNotPresent provides a poller to check if the specified plugin is not present

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/plugin.go#L34)  

```go
func (d *Daemon) PluginIsNotPresent(t testing.TB, name string) func(poll.LogT) poll.Result
```

##### Daemon.PluginIsNotRunning

PluginIsNotRunning provides a poller to check if the specified plugin is not running

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/plugin.go#L24)  

```go
func (d *Daemon) PluginIsNotRunning(t testing.TB, name string) func(poll.LogT) poll.Result
```

##### Daemon.PluginIsRunning

PluginIsRunning provides a poller to check if the specified plugin is running

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/plugin.go#L14)  

```go
func (d *Daemon) PluginIsRunning(t testing.TB, name string) func(poll.LogT) poll.Result
```

##### Daemon.PluginReferenceIs

PluginReferenceIs provides a poller to check if the specified plugin has the specified reference

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/plugin.go#L48)  

```go
func (d *Daemon) PluginReferenceIs(t testing.TB, name, expectedRef string) func(poll.LogT) poll.Result
```

##### Daemon.PollCheckLogs

PollCheckLogs is a poll.Check that checks the daemon logs using the passed in match function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L343)  

```go
func (d *Daemon) PollCheckLogs(ctx context.Context, match func(s string) bool) poll.Check
```

##### Daemon.ReadLogFile

ReadLogFile returns the content of the daemon log file

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L294)  

```go
func (d *Daemon) ReadLogFile() ([]byte, error)
```

##### Daemon.ReloadConfig

ReloadConfig asks the daemon to reload its configuration

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L807)  

```go
func (d *Daemon) ReloadConfig() error
```

##### Daemon.RemoveNode

RemoveNode removes the specified node

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/node.go#L36)  

```go
func (d *Daemon) RemoveNode(ctx context.Context, t testing.TB, id string, force bool)
```

##### Daemon.RemoveService

RemoveService removes the specified service

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/service.go#L89)  

```go
func (d *Daemon) RemoveService(ctx context.Context, t testing.TB, id string)
```

##### Daemon.Restart

Restart will restart the daemon by first stopping it and the starting it.
If an error occurs while starting the daemon, the test will fail.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L781)  

```go
func (d *Daemon) Restart(t testing.TB, args ...string)
```

##### Daemon.RestartNode

RestartNode restarts a daemon to be used as a swarm node

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L37)  

```go
func (d *Daemon) RestartNode(t testing.TB)
```

##### Daemon.RestartWithError

RestartWithError will restart the daemon by first stopping it and then starting it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L788)  

```go
func (d *Daemon) RestartWithError(arg ...string) error
```

##### Daemon.RootDir

RootDir returns the root directory of the daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L265)  

```go
func (d *Daemon) RootDir() string
```

##### Daemon.RotateTokens

RotateTokens update the swarm to rotate tokens

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L176)  

```go
func (d *Daemon) RotateTokens(t testing.TB)
```

##### Daemon.ScanLogs

ScanLogs scans the daemon logs and passes each line to the match function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L397)  

```go
func (d *Daemon) ScanLogs(ctx context.Context, match func(s string) bool) (bool, string, error)
```

##### Daemon.ScanLogsT

ScanLogsT uses `ScanLogs` to match the daemon logs using the passed in match function.
If there is an error or the match fails, the test will fail.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L389)  

```go
func (d *Daemon) ScanLogsT(ctx context.Context, t testing.TB, match func(s string) bool) (bool, string)
```

##### Daemon.SetEnvVar

SetEnvVar updates the set of extra env variables for the daemon, to take
effect on the next start/restart.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L857)  

```go
func (d *Daemon) SetEnvVar(name, val string)
```

##### Daemon.Signal

Signal sends the specified signal to the daemon if running

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L670)  

```go
func (d *Daemon) Signal(signal os.Signal) error
```

##### Daemon.Sock

Sock returns the socket path of the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L280)  

```go
func (d *Daemon) Sock() string
```

##### Daemon.Start

Start starts the daemon and return once it is ready to receive requests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L435)  

```go
func (d *Daemon) Start(t testing.TB, args ...string)
```

##### Daemon.StartAndSwarmInit

StartAndSwarmInit starts the daemon (with busybox) and init the swarm

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L45)  

```go
func (d *Daemon) StartAndSwarmInit(ctx context.Context, t testing.TB)
```

##### Daemon.StartAndSwarmJoin

StartAndSwarmJoin starts the daemon (with busybox) and join the specified swarm as worker or manager

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L51)  

```go
func (d *Daemon) StartAndSwarmJoin(ctx context.Context, t testing.TB, leader *Daemon, manager bool)
```

##### Daemon.StartNode

StartNode (re)starts the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L25)  

```go
func (d *Daemon) StartNode(t testing.TB)
```

##### Daemon.StartNodeWithBusybox

StartNodeWithBusybox starts daemon to be used as a swarm node, and loads the busybox image

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L31)  

```go
func (d *Daemon) StartNodeWithBusybox(ctx context.Context, t testing.TB)
```

##### Daemon.StartWithBusybox

StartWithBusybox will first start the daemon with Daemon.Start()
then save the busybox image from the main daemon and load it into this Daemon instance.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L627)  

```go
func (d *Daemon) StartWithBusybox(ctx context.Context, t testing.TB, arg ...string)
```

##### Daemon.StartWithError

StartWithError starts the daemon and return once it is ready to receive requests.
It returns an error in case it couldn't start.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L446)  

```go
func (d *Daemon) StartWithError(args ...string) error
```

##### Daemon.StartWithLogFile

StartWithLogFile will start the daemon and attach its streams to a given file.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L456)  

```go
func (d *Daemon) StartWithLogFile(out *os.File, providedArgs ...string) error
```

##### Daemon.Stop

Stop will send a SIGINT every second and wait for the daemon to stop.
If it times out, a SIGKILL is sent.
Stop will not delete the daemon directory. If a purged daemon is needed,
instantiate a new one with NewDaemon.
If an error occurs while starting the daemon, the test will fail.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L692)  

```go
func (d *Daemon) Stop(t testing.TB)
```

##### Daemon.StopWithError

StopWithError will send a SIGINT every second and wait for the daemon to stop.
If it timeouts, a SIGKILL is sent.
Stop will not delete the daemon directory. If a purged daemon is needed,
instantiate a new one with NewDaemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L708)  

```go
func (d *Daemon) StopWithError() (retErr error)
```

##### Daemon.StorageDriver

StorageDriver returns the configured storage driver of the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L275)  

```go
func (d *Daemon) StorageDriver() string
```

##### Daemon.SwarmInfo

SwarmInfo returns the swarm information of the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L125)  

```go
func (d *Daemon) SwarmInfo(ctx context.Context, t testing.TB) swarm.Info
```

##### Daemon.SwarmInit

SwarmInit initializes a new swarm cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L81)  

```go
func (d *Daemon) SwarmInit(ctx context.Context, t testing.TB, req swarm.InitRequest)
```

##### Daemon.SwarmJoin

SwarmJoin joins a daemon to an existing cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L101)  

```go
func (d *Daemon) SwarmJoin(ctx context.Context, t testing.TB, req swarm.JoinRequest)
```

##### Daemon.SwarmLeave

SwarmLeave forces daemon to leave current cluster.

The passed in testing.TB is only used to validate that the client was successfully created
Some tests rely on error checking the result of the actual unlock, so allow
the error to be returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L118)  

```go
func (d *Daemon) SwarmLeave(ctx context.Context, t testing.TB, force bool) error
```

##### Daemon.SwarmListenAddr

SwarmListenAddr returns the listen-addr used for the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L71)  

```go
func (d *Daemon) SwarmListenAddr() string
```

##### Daemon.SwarmUnlock

SwarmUnlock tries to unlock a locked swarm

The passed in testing.TB is only used to validate that the client was successfully created
Some tests rely on error checking the result of the actual unlock, so allow
the error to be returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L138)  

```go
func (d *Daemon) SwarmUnlock(t testing.TB, req swarm.UnlockRequest) error
```

##### Daemon.TailLogs

TailLogs tails N lines from the daemon logs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L419)  

```go
func (d *Daemon) TailLogs(n int) ([][]byte, error)
```

##### Daemon.TailLogsT

TailLogsT attempts to tail N lines from the daemon logs.
If there is an error the error is only logged, it does not cause an error with the test.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L331)  

```go
func (d *Daemon) TailLogsT(t LogT, n int)
```

##### Daemon.TamperWithContainerConfig

TamperWithContainerConfig modifies the on-disk config of a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L1015)  

```go
func (d *Daemon) TamperWithContainerConfig(t testing.TB, containerID string, tamper func(*container.Container))
```

##### Daemon.UpdateConfig

UpdateConfig updates the swarm config identified by the specified id
Currently, only label update is supported.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/config.go#L59)  

```go
func (d *Daemon) UpdateConfig(t testing.TB, id string, f ...ConfigConstructor)
```

##### Daemon.UpdateNode

UpdateNode updates a swarm node with the specified node constructor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/node.go#L49)  

```go
func (d *Daemon) UpdateNode(ctx context.Context, t testing.TB, id string, f ...NodeConstructor)
```

##### Daemon.UpdateSecret

UpdateSecret updates the swarm secret identified by the specified id
Currently, only label update is supported.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/secret.go#L60)  

```go
func (d *Daemon) UpdateSecret(t testing.TB, id string, f ...SecretConstructor)
```

##### Daemon.UpdateService

UpdateService updates a swarm service with the specified service constructor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/service.go#L75)  

```go
func (d *Daemon) UpdateService(ctx context.Context, t testing.TB, service *swarm.Service, f ...ServiceConstructor)
```

##### Daemon.UpdateSwarm

UpdateSwarm updates the current swarm object with the specified spec constructors

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L161)  

```go
func (d *Daemon) UpdateSwarm(t testing.TB, f ...SpecConstructor)
```

---

### LogT

LogT is the subset of the testing.TB interface used by the daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/daemon.go#L40)  

```go
type LogT interface {
	Logf(string, ...any)
}
```

---

### NodeConstructor

NodeConstructor defines a swarm node constructor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/node.go#L14)  

```go
type NodeConstructor func(*swarm.Node)
```

---

### Option

Option is used to configure a daemon.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L10)  

```go
type Option func(*Daemon)
```

#### Functions

##### WithContainerdSocket

WithContainerdSocket sets the --containerd option on the daemon.
Use an empty string to remove the option.

If unset the --containerd option will be used with a default value.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L16)  

```go
func WithContainerdSocket(socket string) Option
```

##### WithDefaultCgroupNamespaceMode

WithDefaultCgroupNamespaceMode sets the default cgroup namespace mode for the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L29)  

```go
func WithDefaultCgroupNamespaceMode(mode string) Option
```

##### WithDockerdBinary

WithDockerdBinary sets the dockerd binary to the specified one

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L57)  

```go
func WithDockerdBinary(dockerdBinary string) Option
```

##### WithEnvVars

WithEnvVars sets additional environment variables for the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L140)  

```go
func WithEnvVars(vars ...string) Option
```

##### WithEnvironment

WithEnvironment sets options from testutil/environment.Execution struct

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L106)  

```go
func WithEnvironment(e environment.Execution) Option
```

##### WithExperimental

WithExperimental sets the daemon in experimental mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L43)  

```go
func WithExperimental() Option
```

##### WithInit

WithInit sets the daemon init

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L50)  

```go
func WithInit() Option
```

##### WithOOMScoreAdjust

WithOOMScoreAdjust sets OOM score for the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L133)  

```go
func WithOOMScoreAdjust(score int) Option
```

##### WithResolvConf

WithResolvConf allows a test to provide content for a resolv.conf file to be used
as the basis for resolv.conf in the container, instead of the host's /etc/resolv.conf.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L148)  

```go
func WithResolvConf(content string) Option
```

##### WithRootlessUser

WithRootlessUser sets the daemon to be rootless

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L122)  

```go
func WithRootlessUser(username string) Option
```

##### WithStorageDriver

WithStorageDriver sets store driver option

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L115)  

```go
func WithStorageDriver(driver string) Option
```

##### WithSwarmDataPathPort

WithSwarmDataPathPort sets the  swarm datapath port to use for swarm mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L99)  

```go
func WithSwarmDataPathPort(datapathPort uint32) Option
```

##### WithSwarmDefaultAddrPool

WithSwarmDefaultAddrPool sets the swarm default address pool to use for swarm mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L85)  

```go
func WithSwarmDefaultAddrPool(defaultAddrPool []string) Option
```

##### WithSwarmDefaultAddrPoolSubnetSize

WithSwarmDefaultAddrPoolSubnetSize sets the subnet length mask of swarm default address pool to use for swarm mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L92)  

```go
func WithSwarmDefaultAddrPoolSubnetSize(subnetSize uint32) Option
```

##### WithSwarmIptables

WithSwarmIptables enabled/disables iptables for swarm nodes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L78)  

```go
func WithSwarmIptables(useIptables bool) Option
```

##### WithSwarmListenAddr

WithSwarmListenAddr sets the swarm listen addr to use for swarm mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L71)  

```go
func WithSwarmListenAddr(listenAddr string) Option
```

##### WithSwarmPort

WithSwarmPort sets the swarm port to use for swarm mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L64)  

```go
func WithSwarmPort(port int) Option
```

##### WithTestLogger

WithTestLogger causes the daemon to log certain actions to the provided test.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L36)  

```go
func WithTestLogger(t LogT) Option
```

##### WithUserNsRemap

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/ops.go#L22)  

```go
func WithUserNsRemap(remap string) Option
```

---

### SecretConstructor

SecretConstructor defines a swarm secret constructor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/secret.go#L12)  

```go
type SecretConstructor func(*swarm.Secret)
```

---

### ServiceConstructor

ServiceConstructor defines a swarm service constructor function

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/service.go#L14)  

```go
type ServiceConstructor func(*swarm.Service)
```

---

### SpecConstructor

SpecConstructor defines a swarm spec constructor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/daemon/swarm.go#L68)  

```go
type SpecConstructor func(*swarm.Spec)
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/client

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:14 UTC

## Overview

Package client is a Go client for the Docker Engine API.

For more information about the Engine API, see the documentation:
https://docs.docker.com/reference/api/engine/

You use the library by constructing a client object using NewClientWithOpts
and calling methods on it. The client can be configured from environment
variables by passing the FromEnv option, or configured manually by passing any
of the other available [Opts].

For example, to list running containers (the equivalent of "docker ps"):


## Examples

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/envvars.go#L3)

```go
const (
	// EnvOverrideHost is the name of the environment variable that can be used
	// to override the default host to connect to (DefaultDockerHost).
	//
	// This env-var is read by FromEnv and WithHostFromEnv and when set to a
	// non-empty value, takes precedence over the default host (which is platform
	// specific), or any host already set.
	EnvOverrideHost = "DOCKER_HOST"

	// EnvOverrideAPIVersion is the name of the environment variable that can
	// be used to override the API version to use. Value should be
	// formatted as MAJOR.MINOR, for example, "1.19".
	//
	// This env-var is read by FromEnv and WithVersionFromEnv and when set to a
	// non-empty value, takes precedence over API version negotiation.
	//
	// This environment variable should be used for debugging purposes only, as
	// it can set the client to use an incompatible (or invalid) API version.
	EnvOverrideAPIVersion = "DOCKER_API_VERSION"

	// EnvOverrideCertPath is the name of the environment variable that can be
	// used to specify the directory from which to load the TLS certificates
	// (ca.pem, cert.pem, key.pem) from. These certificates are used to configure
	// the Client for a TCP connection protected by TLS client authentication.
	//
	// TLS certificate verification is enabled by default if the Client is configured
	// to use a TLS connection. Refer to EnvTLSVerify below to learn how to
	// disable verification for testing purposes.
	//
	// WARNING: Access to the remote API is equivalent to root access to the
	// host where the daemon runs. Do not expose the API without protection,
	// and only if needed. Make sure you are familiar with the "daemon attack
	// surface" (https://docs.docker.com/go/attack-surface/).
	//
	// For local access to the API, it is recommended to connect with the daemon
	// using the default local socket connection (on Linux), or the named pipe
	// (on Windows).
	//
	// If you need to access the API of a remote daemon, consider using an SSH
	// (ssh://) connection, which is easier to set up, and requires no additional
	// configuration if the host is accessible using ssh.
	//
	// If you cannot use the alternatives above, and you must expose the API over
	// a TCP connection, refer to https://docs.docker.com/engine/security/protect-access/
	// to learn how to configure the daemon and client to use a TCP connection
	// with TLS client authentication. Make sure you know the differences between
	// a regular TLS connection and a TLS connection protected by TLS client
	// authentication, and verify that the API cannot be accessed by other clients.
	EnvOverrideCertPath = "DOCKER_CERT_PATH"

	// EnvTLSVerify is the name of the environment variable that can be used to
	// enable or disable TLS certificate verification. When set to a non-empty
	// value, TLS certificate verification is enabled, and the client is configured
	// to use a TLS connection, using certificates from the default directories
	// (within `~/.docker`); refer to EnvOverrideCertPath above for additional
	// details.
	//
	// WARNING: Access to the remote API is equivalent to root access to the
	// host where the daemon runs. Do not expose the API without protection,
	// and only if needed. Make sure you are familiar with the "daemon attack
	// surface" (https://docs.docker.com/go/attack-surface/).
	//
	// Before setting up your client and daemon to use a TCP connection with TLS
	// client authentication, consider using one of the alternatives mentioned
	// in EnvOverrideCertPath above.
	//
	// Disabling TLS certificate verification (for testing purposes)
	//
	// TLS certificate verification is enabled by default if the Client is configured
	// to use a TLS connection, and it is highly recommended to keep verification
	// enabled to prevent machine-in-the-middle attacks. Refer to the documentation
	// at https://docs.docker.com/engine/security/protect-access/ and pages linked
	// from that page to learn how to configure the daemon and client to use a
	// TCP connection with TLS client authentication enabled.
	//
	// Set the "DOCKER_TLS_VERIFY" environment to an empty string ("") to
	// disable TLS certificate verification. Disabling verification is insecure,
	// so should only be done for testing purposes. From the Go documentation
	// (https://pkg.go.dev/crypto/tls#Config):
	//
	// InsecureSkipVerify controls whether a client verifies the server's
	// certificate chain and host name. If InsecureSkipVerify is true, crypto/tls
	// accepts any certificate presented by the server and any host name in that
	// certificate. In this mode, TLS is susceptible to machine-in-the-middle
	// attacks unless custom verification is used. This should be used only for
	// testing or in combination with VerifyConnection or VerifyPeerCertificate.
	EnvTLSVerify = "DOCKER_TLS_VERIFY"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_unix.go#L7)

```go
const DefaultDockerHost = "unix:///var/run/docker.sock"
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L92)

```go
const DummyHost = "api.moby.localhost"
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L151)

```go
var ErrRedirect = errors.New("unexpected redirect in response")
```

## Functions

### CheckRedirect

CheckRedirect specifies the policy for dealing with redirect responses. It
can be set on http.Client.CheckRedirect to prevent HTTP redirects for
non-GET requests. It returns an ErrRedirect for non-GET request, otherwise
returns a http.ErrUseLastResponse, which is special-cased by http.Client
to use the last response.

Go 1.8 changed behavior for HTTP redirects (specifically 301, 307, and 308)
in the client. The client (and by extension API client) can be made to send
a request like "POST /containers//start" where what would normally be in the
name section of the URL is empty. This triggers an HTTP 301 from the daemon.

In go 1.8 this 301 is converted to a GET request, and ends up getting
a 404 from the daemon. This behavior change manifests in the client in that
before, the 301 was not followed and the client did not generate an error,
but now results in a message like "Error response from daemon: page not found".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L168)  

```go
func CheckRedirect(_ *http.Request, via []*http.Request) error
```

---

### ErrorConnectionFailed

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/errors.go#L36)  

```go
func ErrorConnectionFailed(host string) error
```

---

### FromEnv

FromEnv configures the client with values from environment variables. It
is the equivalent of using the WithTLSClientConfigFromEnv, WithHostFromEnv,
and WithVersionFromEnv options.

FromEnv uses the following environment variables:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L35)  

```go
func FromEnv(c *Client) error
```

---

### IsErrConnectionFailed

IsErrConnectionFailed returns true if the error is caused by connection failed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/errors.go#L29)  

```go
func IsErrConnectionFailed(err error) bool
```

---

### IsErrNotFound

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/errors.go#L56)  

```go
func IsErrNotFound(err error) bool
```

---

### ParseHostURL

ParseHostURL parses a url string, validates the string is a host url, and
returns the parsed URL

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L412)  

```go
func ParseHostURL(host string) (*url.URL, error)
```

---

## Types

### APIClient

APIClient is an interface that clients that talk with a docker server must implement.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L29)  

```go
type APIClient interface {
	CheckpointAPIClient // CheckpointAPIClient is still experimental.
	// contains filtered or unexported methods
}
```

---

### CheckpointAPIClient

CheckpointAPIClient defines API client methods for the checkpoints.

Experimental: checkpoint and restore is still an experimental feature,
and only available if the daemon is running with experimental features
enabled.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/checkpoint.go#L14)  

```go
type CheckpointAPIClient interface {
	CheckpointCreate(ctx context.Context, container string, options checkpoint.CreateOptions) error
	CheckpointDelete(ctx context.Context, container string, options checkpoint.DeleteOptions) error
	CheckpointList(ctx context.Context, container string, options checkpoint.ListOptions) ([]checkpoint.Summary, error)
}
```

---

### Client

Client is the API client that performs all operations
against a docker server.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L106)  

```go
type Client struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewClient

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_deprecated.go#L17)  

```go
func NewClient(host string, version string, client *http.Client, httpHeaders map[string]string) (*Client, error)
```

##### NewClientWithOpts

NewClientWithOpts initializes a new API client with a default HTTPClient, and
default API host and version. It also initializes the custom HTTP headers to
add to each request.

It takes an optional list of Opt functional arguments, which are applied in
the order they're provided, which allows modifying the defaults when creating
the client. For example, the following initializes a client that configures
itself with values from environment variables (FromEnv), and has automatic
API version negotiation enabled (WithAPIVersionNegotiation).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L189)  

```go
func NewClientWithOpts(ops ...Opt) (*Client, error)
```

##### NewEnvClient

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_deprecated.go#L25)  

```go
func NewEnvClient() (*Client, error)
```

#### Methods

##### Client.BuildCachePrune

BuildCachePrune requests the daemon to delete unused cache data

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/build_prune.go#L15)  

```go
func (cli *Client) BuildCachePrune(ctx context.Context, opts build.CachePruneOptions) (*build.CachePruneReport, error)
```

##### Client.BuildCancel

BuildCancel requests the daemon to cancel the ongoing build request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/build_cancel.go#L9)  

```go
func (cli *Client) BuildCancel(ctx context.Context, id string) error
```

##### Client.CheckpointCreate

CheckpointCreate creates a checkpoint from the given container with the given name

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/checkpoint_create.go#L10)  

```go
func (cli *Client) CheckpointCreate(ctx context.Context, containerID string, options checkpoint.CreateOptions) error
```

##### Client.CheckpointDelete

CheckpointDelete deletes the checkpoint with the given name from the given container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/checkpoint_delete.go#L11)  

```go
func (cli *Client) CheckpointDelete(ctx context.Context, containerID string, options checkpoint.DeleteOptions) error
```

##### Client.CheckpointList

CheckpointList returns the checkpoints of the given container in the docker host

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/checkpoint_list.go#L12)  

```go
func (cli *Client) CheckpointList(ctx context.Context, container string, options checkpoint.ListOptions) ([]checkpoint.Summary, error)
```

##### Client.ClientVersion

ClientVersion returns the API version used by this client.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L317)  

```go
func (cli *Client) ClientVersion() string
```

##### Client.Close

Close the transport used by the client

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L271)  

```go
func (cli *Client) Close() error
```

##### Client.ConfigCreate

ConfigCreate creates a new config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/config_create.go#L11)  

```go
func (cli *Client) ConfigCreate(ctx context.Context, config swarm.ConfigSpec) (swarm.ConfigCreateResponse, error)
```

##### Client.ConfigInspectWithRaw

ConfigInspectWithRaw returns the config information with raw data

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/config_inspect.go#L13)  

```go
func (cli *Client) ConfigInspectWithRaw(ctx context.Context, id string) (swarm.Config, []byte, error)
```

##### Client.ConfigList

ConfigList returns the list of configs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/config_list.go#L13)  

```go
func (cli *Client) ConfigList(ctx context.Context, options swarm.ConfigListOptions) ([]swarm.Config, error)
```

##### Client.ConfigRemove

ConfigRemove removes a config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/config_remove.go#L6)  

```go
func (cli *Client) ConfigRemove(ctx context.Context, id string) error
```

##### Client.ConfigUpdate

ConfigUpdate attempts to update a config

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/config_update.go#L11)  

```go
func (cli *Client) ConfigUpdate(ctx context.Context, id string, version swarm.Version, config swarm.ConfigSpec) error
```

##### Client.ContainerAttach

ContainerAttach attaches a connection to a container in the server.
It returns a types.HijackedConnection with the hijacked connection
and the a reader to get output. It's up to the called to close
the hijacked connection by calling types.HijackedResponse.Close.

The stream format on the response will be in one of two formats:

If the container is using a TTY, there is only a single stream (stdout), and
data is copied directly from the container output stream, no extra
multiplexing or headers.

If the container is *not* using a TTY, streams for stdout and stderr are
multiplexed.
The format of the multiplexed stream is as follows:

STREAM_TYPE can be 1 for stdout and 2 for stderr

SIZE1, SIZE2, SIZE3, and SIZE4 are four bytes of uint32 encoded as big endian.
This is the size of OUTPUT.

You can use github.com/docker/docker/pkg/stdcopy.StdCopy to demultiplex this
stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_attach.go#L36)  

```go
func (cli *Client) ContainerAttach(ctx context.Context, containerID string, options container.AttachOptions) (types.HijackedResponse, error)
```

##### Client.ContainerCommit

ContainerCommit applies changes to a container and creates a new tagged image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_commit.go#L14)  

```go
func (cli *Client) ContainerCommit(ctx context.Context, containerID string, options container.CommitOptions) (container.CommitResponse, error)
```

##### Client.ContainerCreate

ContainerCreate creates a new container based on the given configuration.
It can be associated with a name, but it's not mandatory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_create.go#L20)  

```go
func (cli *Client) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *ocispec.Platform, containerName string) (container.CreateResponse, error)
```

##### Client.ContainerDiff

ContainerDiff shows differences in a container filesystem since it was started.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_diff.go#L12)  

```go
func (cli *Client) ContainerDiff(ctx context.Context, containerID string) ([]container.FilesystemChange, error)
```

##### Client.ContainerExecAttach

ContainerExecAttach attaches a connection to an exec process in the server.
It returns a types.HijackedConnection with the hijacked connection
and the a reader to get output. It's up to the called to close
the hijacked connection by calling types.HijackedResponse.Close.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_exec.go#L61)  

```go
func (cli *Client) ContainerExecAttach(ctx context.Context, execID string, config container.ExecAttachOptions) (types.HijackedResponse, error)
```

##### Client.ContainerExecCreate

ContainerExecCreate creates a new exec configuration to run an exec process.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_exec.go#L14)  

```go
func (cli *Client) ContainerExecCreate(ctx context.Context, containerID string, options container.ExecOptions) (container.ExecCreateResponse, error)
```

##### Client.ContainerExecInspect

ContainerExecInspect returns information about a specific exec process on the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_exec.go#L71)  

```go
func (cli *Client) ContainerExecInspect(ctx context.Context, execID string) (container.ExecInspect, error)
```

##### Client.ContainerExecResize

ContainerExecResize changes the size of the tty for an exec process running inside a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_resize.go#L21)  

```go
func (cli *Client) ContainerExecResize(ctx context.Context, execID string, options container.ResizeOptions) error
```

##### Client.ContainerExecStart

ContainerExecStart starts an exec process already created in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_exec.go#L48)  

```go
func (cli *Client) ContainerExecStart(ctx context.Context, execID string, config container.ExecStartOptions) error
```

##### Client.ContainerExport

ContainerExport retrieves the raw contents of a container
and returns them as an io.ReadCloser. It's up to the caller
to close the stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_export.go#L12)  

```go
func (cli *Client) ContainerExport(ctx context.Context, containerID string) (io.ReadCloser, error)
```

##### Client.ContainerInspect

ContainerInspect returns the container information.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_inspect.go#L14)  

```go
func (cli *Client) ContainerInspect(ctx context.Context, containerID string) (container.InspectResponse, error)
```

##### Client.ContainerInspectWithRaw

ContainerInspectWithRaw returns the container information and its raw representation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_inspect.go#L32)  

```go
func (cli *Client) ContainerInspectWithRaw(ctx context.Context, containerID string, getSize bool) (container.InspectResponse, []byte, error)
```

##### Client.ContainerKill

ContainerKill terminates the container process but does not remove the container from the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_kill.go#L9)  

```go
func (cli *Client) ContainerKill(ctx context.Context, containerID, signal string) error
```

##### Client.ContainerList

ContainerList returns the list of containers in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_list.go#L14)  

```go
func (cli *Client) ContainerList(ctx context.Context, options container.ListOptions) ([]container.Summary, error)
```

##### Client.ContainerLogs

ContainerLogs returns the logs generated by a container in an io.ReadCloser.
It's up to the caller to close the stream.

The stream format on the response will be in one of two formats:

If the container is using a TTY, there is only a single stream (stdout), and
data is copied directly from the container output stream, no extra
multiplexing or headers.

If the container is *not* using a TTY, streams for stdout and stderr are
multiplexed.
The format of the multiplexed stream is as follows:

STREAM_TYPE can be 1 for stdout and 2 for stderr

SIZE1, SIZE2, SIZE3, and SIZE4 are four bytes of uint32 encoded as big endian.
This is the size of OUTPUT.

You can use github.com/docker/docker/pkg/stdcopy.StdCopy to demultiplex this
stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_logs.go#L36)  

```go
func (cli *Client) ContainerLogs(ctx context.Context, containerID string, options container.LogsOptions) (io.ReadCloser, error)
```

##### Client.ContainerPause

ContainerPause pauses the main process of a given container without terminating it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_pause.go#L6)  

```go
func (cli *Client) ContainerPause(ctx context.Context, containerID string) error
```

##### Client.ContainerRemove

ContainerRemove kills and removes a container from the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_remove.go#L11)  

```go
func (cli *Client) ContainerRemove(ctx context.Context, containerID string, options container.RemoveOptions) error
```

##### Client.ContainerRename

ContainerRename changes the name of a given container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_rename.go#L9)  

```go
func (cli *Client) ContainerRename(ctx context.Context, containerID, newContainerName string) error
```

##### Client.ContainerResize

ContainerResize changes the size of the tty for a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_resize.go#L12)  

```go
func (cli *Client) ContainerResize(ctx context.Context, containerID string, options container.ResizeOptions) error
```

##### Client.ContainerRestart

ContainerRestart stops and starts a container again.
It makes the daemon wait for the container to be up again for
a specific amount of time, given the timeout.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_restart.go#L15)  

```go
func (cli *Client) ContainerRestart(ctx context.Context, containerID string, options container.StopOptions) error
```

##### Client.ContainerStart

ContainerStart sends a request to the docker daemon to start a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_start.go#L11)  

```go
func (cli *Client) ContainerStart(ctx context.Context, containerID string, options container.StartOptions) error
```

##### Client.ContainerStatPath

ContainerStatPath returns stat information about a path inside the container filesystem.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_copy.go#L18)  

```go
func (cli *Client) ContainerStatPath(ctx context.Context, containerID, path string) (container.PathStat, error)
```

##### Client.ContainerStats

ContainerStats returns near realtime stats for a given container.
It's up to the caller to close the io.ReadCloser returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_stats.go#L12)  

```go
func (cli *Client) ContainerStats(ctx context.Context, containerID string, stream bool) (container.StatsResponseReader, error)
```

##### Client.ContainerStatsOneShot

ContainerStatsOneShot gets a single stat entry from a container.
It differs from `ContainerStats` in that the API should not wait to prime the stats

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_stats.go#L37)  

```go
func (cli *Client) ContainerStatsOneShot(ctx context.Context, containerID string) (container.StatsResponseReader, error)
```

##### Client.ContainerStop

ContainerStop stops a container. In case the container fails to stop
gracefully within a time frame specified by the timeout argument,
it is forcefully terminated (killed).

If the timeout is nil, the container's StopTimeout value is used, if set,
otherwise the engine default. A negative timeout value can be specified,
meaning no timeout, i.e. no forceful termination is performed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_stop.go#L19)  

```go
func (cli *Client) ContainerStop(ctx context.Context, containerID string, options container.StopOptions) error
```

##### Client.ContainerTop

ContainerTop shows process information from within a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_top.go#L13)  

```go
func (cli *Client) ContainerTop(ctx context.Context, containerID string, arguments []string) (container.TopResponse, error)
```

##### Client.ContainerUnpause

ContainerUnpause resumes the process execution within a container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_unpause.go#L6)  

```go
func (cli *Client) ContainerUnpause(ctx context.Context, containerID string) error
```

##### Client.ContainerUpdate

ContainerUpdate updates the resources of a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_update.go#L11)  

```go
func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.UpdateResponse, error)
```

##### Client.ContainerWait

ContainerWait waits until the specified container is in a certain state
indicated by the given condition, either "not-running" (default),
"next-exit", or "removed".

If this client's API version is before 1.30, condition is ignored and
ContainerWait will return immediately with the two channels, as the server
will wait as if the condition were "not-running".

If this client's API version is at least 1.30, ContainerWait blocks until
the request has been acknowledged by the server (with a response header),
then returns two channels on which the caller can wait for the exit status
of the container or an error if there was a problem either beginning the
wait request or in getting the response. This allows the caller to
synchronize ContainerWait with other calls, such as specifying a
"next-exit" condition before issuing a ContainerStart request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_wait.go#L32)  

```go
func (cli *Client) ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.WaitResponse, <-chan error)
```

##### Client.ContainersPrune

ContainersPrune requests the daemon to delete unused data

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_prune.go#L13)  

```go
func (cli *Client) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (container.PruneReport, error)
```

##### Client.CopyFromContainer

CopyFromContainer gets the content from the container and returns it as a Reader
for a TAR archive to manipulate it in the host. It's up to the caller to close the reader.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_copy.go#L65)  

```go
func (cli *Client) CopyFromContainer(ctx context.Context, containerID, srcPath string) (io.ReadCloser, container.PathStat, error)
```

##### Client.CopyToContainer

CopyToContainer copies content into the container filesystem.
Note that `content` must be a Reader for a TAR archive

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/container_copy.go#L37)  

```go
func (cli *Client) CopyToContainer(ctx context.Context, containerID, dstPath string, content io.Reader, options container.CopyToContainerOptions) error
```

##### Client.DaemonHost

DaemonHost returns the host address used by the client

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L400)  

```go
func (cli *Client) DaemonHost() string
```

##### Client.DialHijack

DialHijack returns a hijacked connection with negotiated protocol proto.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/hijack.go#L42)  

```go
func (cli *Client) DialHijack(ctx context.Context, url, proto string, meta map[string][]string) (net.Conn, error)
```

##### Client.Dialer

Dialer returns a dialer for a raw stream connection, with an HTTP/1.1 header,
that can be used for proxying the daemon connection. It is used by
"docker dial-stdio".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L453)  

```go
func (cli *Client) Dialer() func(context.Context) (net.Conn, error)
```

##### Client.DiskUsage

DiskUsage requests the current data usage from the daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/disk_usage.go#L13)  

```go
func (cli *Client) DiskUsage(ctx context.Context, options types.DiskUsageOptions) (types.DiskUsage, error)
```

##### Client.DistributionInspect

DistributionInspect returns the image digest with the full manifest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/distribution_inspect.go#L13)  

```go
func (cli *Client) DistributionInspect(ctx context.Context, imageRef, encodedRegistryAuth string) (registry.DistributionInspect, error)
```

##### Client.Events

Events returns a stream of events in the daemon. It's up to the caller to close the stream
by cancelling the context. Once the stream has been completely read an io.EOF error will
be sent over the error channel. If an error is sent all processing will be stopped. It's up
to the caller to reopen the stream in the event of an error by reinvoking this method.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/events.go#L18)  

```go
func (cli *Client) Events(ctx context.Context, options events.ListOptions) (<-chan events.Message, <-chan error)
```

##### Client.HTTPClient

HTTPClient returns a copy of the HTTP client bound to the server

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L405)  

```go
func (cli *Client) HTTPClient() *http.Client
```

##### Client.ImageBuild

ImageBuild sends a request to the daemon to build images.
The Body in the response implements an io.ReadCloser and it's up to the caller to
close it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_build.go#L21)  

```go
func (cli *Client) ImageBuild(ctx context.Context, buildContext io.Reader, options build.ImageBuildOptions) (build.ImageBuildResponse, error)
```

##### Client.ImageCreate

ImageCreate creates a new image based on the parent options.
It returns the JSON content in the response body.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_create.go#L17)  

```go
func (cli *Client) ImageCreate(ctx context.Context, parentReference string, options image.CreateOptions) (io.ReadCloser, error)
```

##### Client.ImageHistory

ImageHistory returns the changes in an image in history format.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_history.go#L25)  

```go
func (cli *Client) ImageHistory(ctx context.Context, imageID string, historyOpts ...ImageHistoryOption) ([]image.HistoryResponseItem, error)
```

##### Client.ImageImport

ImageImport creates a new image based on the source options.
It returns the JSON content in the response body.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_import.go#L15)  

```go
func (cli *Client) ImageImport(ctx context.Context, source image.ImportSource, ref string, options image.ImportOptions) (io.ReadCloser, error)
```

##### Client.ImageInspect

ImageInspect returns the image information.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_inspect.go#L15)  

```go
func (cli *Client) ImageInspect(ctx context.Context, imageID string, inspectOpts ...ImageInspectOption) (image.InspectResponse, error)
```

##### Client.ImageInspectWithRaw

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_inspect.go#L69)  

```go
func (cli *Client) ImageInspectWithRaw(ctx context.Context, imageID string) (image.InspectResponse, []byte, error)
```

##### Client.ImageList

ImageList returns a list of images in the docker host.

Experimental: Setting the [options.Manifest] will populate
image.Summary.Manifests with information about image manifests.
This is experimental and might change in the future without any backward
compatibility.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_list.go#L19)  

```go
func (cli *Client) ImageList(ctx context.Context, options image.ListOptions) ([]image.Summary, error)
```

##### Client.ImageLoad

ImageLoad loads an image in the docker host from the client host.
It's up to the caller to close the io.ReadCloser in the
ImageLoadResponse returned by this function.

Platform is an optional parameter that specifies the platform to load from
the provided multi-platform image. This is only has effect if the input image
is a multi-platform image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_load.go#L19)  

```go
func (cli *Client) ImageLoad(ctx context.Context, input io.Reader, loadOpts ...ImageLoadOption) (image.LoadResponse, error)
```

##### Client.ImagePull

ImagePull requests the docker host to pull an image from a remote registry.
It executes the privileged function if the operation is unauthorized
and it tries one more time.
It's up to the caller to handle the io.ReadCloser and close it properly.

FIXME(vdemeester): there is currently used in a few way in docker/docker
- if not in trusted content, ref is used to pass the whole reference, and tag is empty
- if in trusted content, ref is used to pass the reference name, and tag for the digest

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_pull.go#L22)  

```go
func (cli *Client) ImagePull(ctx context.Context, refStr string, options image.PullOptions) (io.ReadCloser, error)
```

##### Client.ImagePush

ImagePush requests the docker host to push an image to a remote registry.
It executes the privileged function if the operation is unauthorized
and it tries one more time.
It's up to the caller to handle the io.ReadCloser and close it properly.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_push.go#L22)  

```go
func (cli *Client) ImagePush(ctx context.Context, image string, options image.PushOptions) (io.ReadCloser, error)
```

##### Client.ImageRemove

ImageRemove removes an image from the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_remove.go#L12)  

```go
func (cli *Client) ImageRemove(ctx context.Context, imageID string, options image.RemoveOptions) ([]image.DeleteResponse, error)
```

##### Client.ImageSave

ImageSave retrieves one or more images from the docker host as an io.ReadCloser.

Platforms is an optional parameter that specifies the platforms to save from the image.
This is only has effect if the input image is a multi-platform image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_save.go#L13)  

```go
func (cli *Client) ImageSave(ctx context.Context, imageIDs []string, saveOpts ...ImageSaveOption) (io.ReadCloser, error)
```

##### Client.ImageSearch

ImageSearch makes the docker host search by a term in a remote registry.
The list of results is not sorted in any fashion.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_search.go#L17)  

```go
func (cli *Client) ImageSearch(ctx context.Context, term string, options registry.SearchOptions) ([]registry.SearchResult, error)
```

##### Client.ImageTag

ImageTag tags an image in the docker host

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_tag.go#L12)  

```go
func (cli *Client) ImageTag(ctx context.Context, source, target string) error
```

##### Client.ImagesPrune

ImagesPrune requests the daemon to delete unused data

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_prune.go#L13)  

```go
func (cli *Client) ImagesPrune(ctx context.Context, pruneFilters filters.Args) (image.PruneReport, error)
```

##### Client.Info

Info returns information about the docker server.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/info.go#L13)  

```go
func (cli *Client) Info(ctx context.Context) (system.Info, error)
```

##### Client.NegotiateAPIVersion

NegotiateAPIVersion queries the API and updates the version to match the API
version. NegotiateAPIVersion downgrades the client's API version to match the
APIVersion if the ping version is lower than the default version. If the API
version reported by the server is higher than the maximum version supported
by the client, it uses the client's maximum version.

If a manual override is in place, either through the "DOCKER_API_VERSION"
(EnvOverrideAPIVersion) environment variable, or if the client is initialized
with a fixed version (WithVersion), no negotiation is performed.

If the API server's ping response does not contain an API version, or if the
client did not get a successful ping response, it assumes it is connected with
an old daemon that does not support API version negotiation, in which case it
downgrades to the latest version of the API before version negotiation was
added (1.24).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L336)  

```go
func (cli *Client) NegotiateAPIVersion(ctx context.Context)
```

##### Client.NegotiateAPIVersionPing

NegotiateAPIVersionPing downgrades the client's API version to match the
APIVersion in the ping response. If the API version in pingResponse is higher
than the maximum version supported by the client, it uses the client's maximum
version.

If a manual override is in place, either through the "DOCKER_API_VERSION"
(EnvOverrideAPIVersion) environment variable, or if the client is initialized
with a fixed version (WithVersion), no negotiation is performed.

If the API server's ping response does not contain an API version, we assume
we are connected with an old daemon without API version negotiation support,
and downgrade to the latest version of the API before version negotiation was
added (1.24).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client.go#L364)  

```go
func (cli *Client) NegotiateAPIVersionPing(pingResponse types.Ping)
```

##### Client.NetworkConnect

NetworkConnect connects a container to an existent network in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/network_connect.go#L10)  

```go
func (cli *Client) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error
```

##### Client.NetworkCreate

NetworkCreate creates a new network in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/network_create.go#L12)  

```go
func (cli *Client) NetworkCreate(ctx context.Context, name string, options network.CreateOptions) (network.CreateResponse, error)
```

##### Client.NetworkDisconnect

NetworkDisconnect disconnects a container from an existent network in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/network_disconnect.go#L10)  

```go
func (cli *Client) NetworkDisconnect(ctx context.Context, networkID, containerID string, force bool) error
```

##### Client.NetworkInspect

NetworkInspect returns the information for a specific network configured in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/network_inspect.go#L14)  

```go
func (cli *Client) NetworkInspect(ctx context.Context, networkID string, options network.InspectOptions) (network.Inspect, error)
```

##### Client.NetworkInspectWithRaw

NetworkInspectWithRaw returns the information for a specific network configured in the docker host and its raw representation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/network_inspect.go#L20)  

```go
func (cli *Client) NetworkInspectWithRaw(ctx context.Context, networkID string, options network.InspectOptions) (network.Inspect, []byte, error)
```

##### Client.NetworkList

NetworkList returns the list of networks configured in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/network_list.go#L13)  

```go
func (cli *Client) NetworkList(ctx context.Context, options network.ListOptions) ([]network.Summary, error)
```

##### Client.NetworkRemove

NetworkRemove removes an existent network from the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/network_remove.go#L6)  

```go
func (cli *Client) NetworkRemove(ctx context.Context, networkID string) error
```

##### Client.NetworksPrune

NetworksPrune requests the daemon to delete unused networks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/network_prune.go#L13)  

```go
func (cli *Client) NetworksPrune(ctx context.Context, pruneFilters filters.Args) (network.PruneReport, error)
```

##### Client.NewVersionError

NewVersionError returns an error if the APIVersion required is less than the
current supported version.

It performs API-version negotiation if the Client is configured with this
option, otherwise it assumes the latest API version is used.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/errors.go#L76)  

```go
func (cli *Client) NewVersionError(ctx context.Context, APIrequired, feature string) error
```

##### Client.NodeInspectWithRaw

NodeInspectWithRaw returns the node information.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/node_inspect.go#L13)  

```go
func (cli *Client) NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error)
```

##### Client.NodeList

NodeList returns the list of nodes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/node_list.go#L13)  

```go
func (cli *Client) NodeList(ctx context.Context, options swarm.NodeListOptions) ([]swarm.Node, error)
```

##### Client.NodeRemove

NodeRemove removes a Node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/node_remove.go#L11)  

```go
func (cli *Client) NodeRemove(ctx context.Context, nodeID string, options swarm.NodeRemoveOptions) error
```

##### Client.NodeUpdate

NodeUpdate updates a Node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/node_update.go#L11)  

```go
func (cli *Client) NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error
```

##### Client.Ping

Ping pings the server and returns the value of the "Docker-Experimental",
"Builder-Version", "OS-Type" & "API-Version" headers. It attempts to use
a HEAD request on the endpoint, but falls back to GET if HEAD is not supported
by the daemon. It ignores internal server errors returned by the API, which
may be returned if the daemon is in an unhealthy state, but returns errors
for other non-success status codes, failing to connect to the API, or failing
to parse the API response.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/ping.go#L21)  

```go
func (cli *Client) Ping(ctx context.Context) (types.Ping, error)
```

##### Client.PluginCreate

PluginCreate creates a plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_create.go#L13)  

```go
func (cli *Client) PluginCreate(ctx context.Context, createContext io.Reader, createOptions types.PluginCreateOptions) error
```

##### Client.PluginDisable

PluginDisable disables a plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_disable.go#L11)  

```go
func (cli *Client) PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error
```

##### Client.PluginEnable

PluginEnable enables a plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_enable.go#L12)  

```go
func (cli *Client) PluginEnable(ctx context.Context, name string, options types.PluginEnableOptions) error
```

##### Client.PluginInspectWithRaw

PluginInspectWithRaw inspects an existing plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_inspect.go#L13)  

```go
func (cli *Client) PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error)
```

##### Client.PluginInstall

PluginInstall installs a plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_install.go#L18)  

```go
func (cli *Client) PluginInstall(ctx context.Context, name string, options types.PluginInstallOptions) (_ io.ReadCloser, retErr error)
```

##### Client.PluginList

PluginList returns the installed plugins

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_list.go#L13)  

```go
func (cli *Client) PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error)
```

##### Client.PluginPush

PluginPush pushes a plugin to a registry

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_push.go#L12)  

```go
func (cli *Client) PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error)
```

##### Client.PluginRemove

PluginRemove removes a plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_remove.go#L11)  

```go
func (cli *Client) PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error
```

##### Client.PluginSet

PluginSet modifies settings for an existing plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_set.go#L8)  

```go
func (cli *Client) PluginSet(ctx context.Context, name string, args []string) error
```

##### Client.PluginUpgrade

PluginUpgrade upgrades a plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/plugin_upgrade.go#L16)  
**Added in:** v1.13.1

```go
func (cli *Client) PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (io.ReadCloser, error)
```

##### Client.RegistryLogin

RegistryLogin authenticates the docker server with a given docker registry.
It returns unauthorizedError when the authentication fails.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/login.go#L13)  

```go
func (cli *Client) RegistryLogin(ctx context.Context, auth registry.AuthConfig) (registry.AuthenticateOKBody, error)
```

##### Client.SecretCreate

SecretCreate creates a new secret.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/secret_create.go#L11)  

```go
func (cli *Client) SecretCreate(ctx context.Context, secret swarm.SecretSpec) (swarm.SecretCreateResponse, error)
```

##### Client.SecretInspectWithRaw

SecretInspectWithRaw returns the secret information with raw data

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/secret_inspect.go#L13)  

```go
func (cli *Client) SecretInspectWithRaw(ctx context.Context, id string) (swarm.Secret, []byte, error)
```

##### Client.SecretList

SecretList returns the list of secrets.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/secret_list.go#L13)  

```go
func (cli *Client) SecretList(ctx context.Context, options swarm.SecretListOptions) ([]swarm.Secret, error)
```

##### Client.SecretRemove

SecretRemove removes a secret.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/secret_remove.go#L6)  

```go
func (cli *Client) SecretRemove(ctx context.Context, id string) error
```

##### Client.SecretUpdate

SecretUpdate attempts to update a secret.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/secret_update.go#L11)  
**Added in:** v1.13.1

```go
func (cli *Client) SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error
```

##### Client.ServerVersion

ServerVersion returns information of the docker client and server host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/version.go#L11)  

```go
func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error)
```

##### Client.ServiceCreate

ServiceCreate creates a new service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/service_create.go#L19)  

```go
func (cli *Client) ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options swarm.ServiceCreateOptions) (swarm.ServiceCreateResponse, error)
```

##### Client.ServiceInspectWithRaw

ServiceInspectWithRaw returns the service information and the raw data.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/service_inspect.go#L15)  

```go
func (cli *Client) ServiceInspectWithRaw(ctx context.Context, serviceID string, opts swarm.ServiceInspectOptions) (swarm.Service, []byte, error)
```

##### Client.ServiceList

ServiceList returns the list of services.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/service_list.go#L13)  

```go
func (cli *Client) ServiceList(ctx context.Context, options swarm.ServiceListOptions) ([]swarm.Service, error)
```

##### Client.ServiceLogs

ServiceLogs returns the logs generated by a service in an io.ReadCloser.
It's up to the caller to close the stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/service_logs.go#L16)  

```go
func (cli *Client) ServiceLogs(ctx context.Context, serviceID string, options container.LogsOptions) (io.ReadCloser, error)
```

##### Client.ServiceRemove

ServiceRemove kills and removes a service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/service_remove.go#L6)  

```go
func (cli *Client) ServiceRemove(ctx context.Context, serviceID string) error
```

##### Client.ServiceUpdate

ServiceUpdate updates a Service. The version number is required to avoid conflicting writes.
It should be the value as set *before* the update. You can find this value in the Meta field
of swarm.Service, which can be found using ServiceInspectWithRaw.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/service_update.go#L17)  

```go
func (cli *Client) ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options swarm.ServiceUpdateOptions) (swarm.ServiceUpdateResponse, error)
```

##### Client.SwarmGetUnlockKey

SwarmGetUnlockKey retrieves the swarm's unlock key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/swarm_get_unlock_key.go#L11)  

```go
func (cli *Client) SwarmGetUnlockKey(ctx context.Context) (swarm.UnlockKeyResponse, error)
```

##### Client.SwarmInit

SwarmInit initializes the swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/swarm_init.go#L11)  

```go
func (cli *Client) SwarmInit(ctx context.Context, req swarm.InitRequest) (string, error)
```

##### Client.SwarmInspect

SwarmInspect inspects the swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/swarm_inspect.go#L11)  

```go
func (cli *Client) SwarmInspect(ctx context.Context) (swarm.Swarm, error)
```

##### Client.SwarmJoin

SwarmJoin joins the swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/swarm_join.go#L10)  

```go
func (cli *Client) SwarmJoin(ctx context.Context, req swarm.JoinRequest) error
```

##### Client.SwarmLeave

SwarmLeave leaves the swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/swarm_leave.go#L9)  

```go
func (cli *Client) SwarmLeave(ctx context.Context, force bool) error
```

##### Client.SwarmUnlock

SwarmUnlock unlocks locked swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/swarm_unlock.go#L10)  

```go
func (cli *Client) SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error
```

##### Client.SwarmUpdate

SwarmUpdate updates the swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/swarm_update.go#L12)  

```go
func (cli *Client) SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error
```

##### Client.TaskInspectWithRaw

TaskInspectWithRaw returns the task information and its raw representation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/task_inspect.go#L13)  

```go
func (cli *Client) TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error)
```

##### Client.TaskList

TaskList returns the list of tasks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/task_list.go#L13)  

```go
func (cli *Client) TaskList(ctx context.Context, options swarm.TaskListOptions) ([]swarm.Task, error)
```

##### Client.TaskLogs

TaskLogs returns the logs generated by a task in an io.ReadCloser.
It's up to the caller to close the stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/task_logs.go#L15)  

```go
func (cli *Client) TaskLogs(ctx context.Context, taskID string, options container.LogsOptions) (io.ReadCloser, error)
```

##### Client.VolumeCreate

VolumeCreate creates a volume in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/volume_create.go#L11)  

```go
func (cli *Client) VolumeCreate(ctx context.Context, options volume.CreateOptions) (volume.Volume, error)
```

##### Client.VolumeInspect

VolumeInspect returns the information about a specific volume in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/volume_inspect.go#L13)  

```go
func (cli *Client) VolumeInspect(ctx context.Context, volumeID string) (volume.Volume, error)
```

##### Client.VolumeInspectWithRaw

VolumeInspectWithRaw returns the information about a specific volume in the docker host and its raw representation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/volume_inspect.go#L19)  

```go
func (cli *Client) VolumeInspectWithRaw(ctx context.Context, volumeID string) (volume.Volume, []byte, error)
```

##### Client.VolumeList

VolumeList returns the volumes configured in the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/volume_list.go#L13)  

```go
func (cli *Client) VolumeList(ctx context.Context, options volume.ListOptions) (volume.ListResponse, error)
```

##### Client.VolumeRemove

VolumeRemove removes a volume from the docker host.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/volume_remove.go#L11)  

```go
func (cli *Client) VolumeRemove(ctx context.Context, volumeID string, force bool) error
```

##### Client.VolumeUpdate

VolumeUpdate updates a volume. This only works for Cluster Volumes, and
only some fields can be updated.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/volume_update.go#L13)  

```go
func (cli *Client) VolumeUpdate(ctx context.Context, volumeID string, version swarm.Version, options volume.UpdateOptions) error
```

##### Client.VolumesPrune

VolumesPrune requests the daemon to delete unused data

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/volume_prune.go#L13)  

```go
func (cli *Client) VolumesPrune(ctx context.Context, pruneFilters filters.Args) (volume.PruneReport, error)
```

---

### CommonAPIClient

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L26)  

```go
type CommonAPIClient = stableAPIClient
```

---

### ConfigAPIClient

ConfigAPIClient defines API client methods for configs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L231)  

```go
type ConfigAPIClient interface {
	ConfigList(ctx context.Context, options swarm.ConfigListOptions) ([]swarm.Config, error)
	ConfigCreate(ctx context.Context, config swarm.ConfigSpec) (swarm.ConfigCreateResponse, error)
	ConfigRemove(ctx context.Context, id string) error
	ConfigInspectWithRaw(ctx context.Context, name string) (swarm.Config, []byte, error)
	ConfigUpdate(ctx context.Context, id string, version swarm.Version, config swarm.ConfigSpec) error
}
```

---

### ContainerAPIClient

ContainerAPIClient defines API client methods for the containers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L71)  

```go
type ContainerAPIClient interface {
	ContainerAttach(ctx context.Context, container string, options container.AttachOptions) (types.HijackedResponse, error)
	ContainerCommit(ctx context.Context, container string, options container.CommitOptions) (container.CommitResponse, error)
	ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *ocispec.Platform, containerName string) (container.CreateResponse, error)
	ContainerDiff(ctx context.Context, container string) ([]container.FilesystemChange, error)
	ContainerExecAttach(ctx context.Context, execID string, options container.ExecAttachOptions) (types.HijackedResponse, error)
	ContainerExecCreate(ctx context.Context, container string, options container.ExecOptions) (container.ExecCreateResponse, error)
	ContainerExecInspect(ctx context.Context, execID string) (container.ExecInspect, error)
	ContainerExecResize(ctx context.Context, execID string, options container.ResizeOptions) error
	ContainerExecStart(ctx context.Context, execID string, options container.ExecStartOptions) error
	ContainerExport(ctx context.Context, container string) (io.ReadCloser, error)
	ContainerInspect(ctx context.Context, container string) (container.InspectResponse, error)
	ContainerInspectWithRaw(ctx context.Context, container string, getSize bool) (container.InspectResponse, []byte, error)
	ContainerKill(ctx context.Context, container, signal string) error
	ContainerList(ctx context.Context, options container.ListOptions) ([]container.Summary, error)
	ContainerLogs(ctx context.Context, container string, options container.LogsOptions) (io.ReadCloser, error)
	ContainerPause(ctx context.Context, container string) error
	ContainerRemove(ctx context.Context, container string, options container.RemoveOptions) error
	ContainerRename(ctx context.Context, container, newContainerName string) error
	ContainerResize(ctx context.Context, container string, options container.ResizeOptions) error
	ContainerRestart(ctx context.Context, container string, options container.StopOptions) error
	ContainerStatPath(ctx context.Context, container, path string) (container.PathStat, error)
	ContainerStats(ctx context.Context, container string, stream bool) (container.StatsResponseReader, error)
	ContainerStatsOneShot(ctx context.Context, container string) (container.StatsResponseReader, error)
	ContainerStart(ctx context.Context, container string, options container.StartOptions) error
	ContainerStop(ctx context.Context, container string, options container.StopOptions) error
	ContainerTop(ctx context.Context, container string, arguments []string) (container.TopResponse, error)
	ContainerUnpause(ctx context.Context, container string) error
	ContainerUpdate(ctx context.Context, container string, updateConfig container.UpdateConfig) (container.UpdateResponse, error)
	ContainerWait(ctx context.Context, container string, condition container.WaitCondition) (<-chan container.WaitResponse, <-chan error)
	CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, container.PathStat, error)
	CopyToContainer(ctx context.Context, container, path string, content io.Reader, options container.CopyToContainerOptions) error
	ContainersPrune(ctx context.Context, pruneFilters filters.Args) (container.PruneReport, error)
}
```

---

### DistributionAPIClient

DistributionAPIClient defines API client methods for the registry

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L107)  

```go
type DistributionAPIClient interface {
	DistributionInspect(ctx context.Context, image, encodedRegistryAuth string) (registry.DistributionInspect, error)
}
```

---

### HijackDialer

HijackDialer defines methods for a hijack dialer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L66)  

```go
type HijackDialer interface {
	DialHijack(ctx context.Context, url, proto string, meta map[string][]string) (net.Conn, error)
}
```

---

### ImageAPIClient

ImageAPIClient defines API client methods for the images

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L112)  

```go
type ImageAPIClient interface {
	ImageBuild(ctx context.Context, context io.Reader, options build.ImageBuildOptions) (build.ImageBuildResponse, error)
	BuildCachePrune(ctx context.Context, opts build.CachePruneOptions) (*build.CachePruneReport, error)
	BuildCancel(ctx context.Context, id string) error
	ImageCreate(ctx context.Context, parentReference string, options image.CreateOptions) (io.ReadCloser, error)
	ImageImport(ctx context.Context, source image.ImportSource, ref string, options image.ImportOptions) (io.ReadCloser, error)

	ImageList(ctx context.Context, options image.ListOptions) ([]image.Summary, error)
	ImagePull(ctx context.Context, ref string, options image.PullOptions) (io.ReadCloser, error)
	ImagePush(ctx context.Context, ref string, options image.PushOptions) (io.ReadCloser, error)
	ImageRemove(ctx context.Context, image string, options image.RemoveOptions) ([]image.DeleteResponse, error)
	ImageSearch(ctx context.Context, term string, options registry.SearchOptions) ([]registry.SearchResult, error)
	ImageTag(ctx context.Context, image, ref string) error
	ImagesPrune(ctx context.Context, pruneFilter filters.Args) (image.PruneReport, error)

	ImageInspect(ctx context.Context, image string, _ ...ImageInspectOption) (image.InspectResponse, error)
	ImageHistory(ctx context.Context, image string, _ ...ImageHistoryOption) ([]image.HistoryResponseItem, error)
	ImageLoad(ctx context.Context, input io.Reader, _ ...ImageLoadOption) (image.LoadResponse, error)
	ImageSave(ctx context.Context, images []string, _ ...ImageSaveOption) (io.ReadCloser, error)

	ImageAPIClientDeprecated
}
```

---

### ImageAPIClientDeprecated

ImageAPIClientDeprecated defines deprecated methods of the ImageAPIClient.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L136)  

```go
type ImageAPIClientDeprecated interface {
	// ImageInspectWithRaw returns the image information and its raw representation.
	//
	// Deprecated: Use [Client.ImageInspect] instead. Raw response can be obtained using the [ImageInspectWithRawResponse] option.
	ImageInspectWithRaw(ctx context.Context, image string) (image.InspectResponse, []byte, error)
}
```

---

### ImageHistoryOption

ImageHistoryOption is a type representing functional options for the image history operation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_history_opts.go#L8)  

```go
type ImageHistoryOption interface {
	Apply(*imageHistoryOpts) error
}
```

#### Functions

##### ImageHistoryWithPlatform

ImageHistoryWithPlatform sets the platform for the image history operation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_history.go#L14)  

```go
func ImageHistoryWithPlatform(platform ocispec.Platform) ImageHistoryOption
```

---

### ImageInspectOption

ImageInspectOption is a type representing functional options for the image inspect operation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_inspect_opts.go#L11)  

```go
type ImageInspectOption interface {
	Apply(*imageInspectOpts) error
}
```

#### Functions

##### ImageInspectWithAPIOpts

ImageInspectWithAPIOpts sets the API options for the image inspect operation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_inspect_opts.go#L52)  

```go
func ImageInspectWithAPIOpts(opts image.InspectOptions) ImageInspectOption
```

##### ImageInspectWithManifests

ImageInspectWithManifests sets manifests API option for the image inspect operation.
This option is only available for API version 1.48 and up.
With this option set, the image inspect operation response will have the
image.InspectResponse.Manifests field populated if the server is multi-platform capable.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_inspect_opts.go#L33)  

```go
func ImageInspectWithManifests(manifests bool) ImageInspectOption
```

##### ImageInspectWithPlatform

ImageInspectWithPlatform sets platform API option for the image inspect operation.
This option is only available for API version 1.49 and up.
With this option set, the image inspect operation will return information for the
specified platform variant of the multi-platform image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_inspect_opts.go#L44)  

```go
func ImageInspectWithPlatform(platform *ocispec.Platform) ImageInspectOption
```

##### ImageInspectWithRawResponse

ImageInspectWithRawResponse instructs the client to additionally store the
raw inspect response in the provided buffer.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_inspect_opts.go#L22)  

```go
func ImageInspectWithRawResponse(raw *bytes.Buffer) ImageInspectOption
```

---

### ImageLoadOption

ImageLoadOption is a type representing functional options for the image load operation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_load_opts.go#L11)  

```go
type ImageLoadOption interface {
	Apply(*imageLoadOpts) error
}
```

#### Functions

##### ImageLoadWithPlatforms

ImageLoadWithPlatforms sets the platforms to be loaded from the image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_load_opts.go#L33)  

```go
func ImageLoadWithPlatforms(platforms ...ocispec.Platform) ImageLoadOption
```

##### ImageLoadWithQuiet

ImageLoadWithQuiet sets the quiet option for the image load operation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_load_opts.go#L25)  

```go
func ImageLoadWithQuiet(quiet bool) ImageLoadOption
```

---

### ImageSaveOption

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_save_opts.go#L10)  

```go
type ImageSaveOption interface {
	Apply(*imageSaveOpts) error
}
```

#### Functions

##### ImageSaveWithPlatforms

ImageSaveWithPlatforms sets the platforms to be saved from the image.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/image_save_opts.go#L21)  

```go
func ImageSaveWithPlatforms(platforms ...ocispec.Platform) ImageSaveOption
```

---

### NetworkAPIClient

NetworkAPIClient defines API client methods for the networks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L144)  

```go
type NetworkAPIClient interface {
	NetworkConnect(ctx context.Context, network, container string, config *network.EndpointSettings) error
	NetworkCreate(ctx context.Context, name string, options network.CreateOptions) (network.CreateResponse, error)
	NetworkDisconnect(ctx context.Context, network, container string, force bool) error
	NetworkInspect(ctx context.Context, network string, options network.InspectOptions) (network.Inspect, error)
	NetworkInspectWithRaw(ctx context.Context, network string, options network.InspectOptions) (network.Inspect, []byte, error)
	NetworkList(ctx context.Context, options network.ListOptions) ([]network.Summary, error)
	NetworkRemove(ctx context.Context, network string) error
	NetworksPrune(ctx context.Context, pruneFilter filters.Args) (network.PruneReport, error)
}
```

---

### NodeAPIClient

NodeAPIClient defines API client methods for the nodes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L156)  

```go
type NodeAPIClient interface {
	NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error)
	NodeList(ctx context.Context, options swarm.NodeListOptions) ([]swarm.Node, error)
	NodeRemove(ctx context.Context, nodeID string, options swarm.NodeRemoveOptions) error
	NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error
}
```

---

### Opt

Opt is a configuration option to initialize a Client.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L20)  

```go
type Opt func(*Client) error
```

#### Functions

##### WithAPIVersionNegotiation

WithAPIVersionNegotiation enables automatic API version negotiation for the client.
With this option enabled, the client automatically negotiates the API version
to use when making requests. API version negotiation is performed on the first
request; subsequent requests do not re-negotiate.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L221)  

```go
func WithAPIVersionNegotiation() Opt
```

##### WithDialContext

WithDialContext applies the dialer to the client transport. This can be
used to set the Timeout and KeepAlive settings of the client. It returns
an error if the client does not have a http.Transport configured.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L52)  

```go
func WithDialContext(dialContext func(ctx context.Context, network, addr string) (net.Conn, error)) Opt
```

##### WithHTTPClient

WithHTTPClient overrides the client's HTTP client with the specified one.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L93)  

```go
func WithHTTPClient(client *http.Client) Opt
```

##### WithHTTPHeaders

WithHTTPHeaders appends custom HTTP headers to the client's default headers.
It does not allow for built-in headers (such as "User-Agent", if set) to
be overridden. Also see WithUserAgent.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L123)  

```go
func WithHTTPHeaders(headers map[string]string) Opt
```

##### WithHost

WithHost overrides the client host with the specified one.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L63)  

```go
func WithHost(host string) Opt
```

##### WithHostFromEnv

WithHostFromEnv overrides the client host with the host specified in the
DOCKER_HOST (EnvOverrideHost) environment variable. If DOCKER_HOST is not set,
or set to an empty value, the host is not modified.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L83)  

```go
func WithHostFromEnv() Opt
```

##### WithScheme

WithScheme overrides the client scheme with the specified one.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L131)  

```go
func WithScheme(scheme string) Opt
```

##### WithTLSClientConfig

WithTLSClientConfig applies a TLS config to the client transport.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L139)  

```go
func WithTLSClientConfig(cacertPath, certPath, keyPath string) Opt
```

##### WithTLSClientConfigFromEnv

WithTLSClientConfigFromEnv configures the client's TLS settings with the
settings in the DOCKER_CERT_PATH (EnvOverrideCertPath) and DOCKER_TLS_VERIFY
(EnvTLSVerify) environment variables. If DOCKER_CERT_PATH is not set or empty,
TLS configuration is not modified.

WithTLSClientConfigFromEnv uses the following environment variables:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L170)  

```go
func WithTLSClientConfigFromEnv() Opt
```

##### WithTimeout

WithTimeout configures the time limit for requests made by the HTTP client.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L103)  

```go
func WithTimeout(timeout time.Duration) Opt
```

##### WithTraceOptions

WithTraceOptions sets tracing span options for the client.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L235)  

```go
func WithTraceOptions(opts ...otelhttp.Option) Opt
```

##### WithTraceProvider

WithTraceProvider sets the trace provider for the client.
If this is not set then the global trace provider will be used.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L230)  

```go
func WithTraceProvider(provider trace.TracerProvider) Opt
```

##### WithUserAgent

WithUserAgent configures the User-Agent header to use for HTTP requests.
It overrides any User-Agent set in headers. When set to an empty string,
the User-Agent header is removed, and no header is sent.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L113)  

```go
func WithUserAgent(ua string) Opt
```

##### WithVersion

WithVersion overrides the client version with the specified one. If an empty
version is provided, the value is ignored to allow version negotiation
(see WithAPIVersionNegotiation).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L197)  

```go
func WithVersion(version string) Opt
```

##### WithVersionFromEnv

WithVersionFromEnv overrides the client version with the version specified in
the DOCKER_API_VERSION (EnvOverrideAPIVersion) environment variable.
If DOCKER_API_VERSION is not set, or set to an empty value, the version
is not modified.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/options.go#L211)  

```go
func WithVersionFromEnv() Opt
```

---

### PluginAPIClient

PluginAPIClient defines API client methods for the plugins

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L164)  

```go
type PluginAPIClient interface {
	PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error)
	PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error
	PluginEnable(ctx context.Context, name string, options types.PluginEnableOptions) error
	PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error
	PluginInstall(ctx context.Context, name string, options types.PluginInstallOptions) (io.ReadCloser, error)
	PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (io.ReadCloser, error)
	PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error)
	PluginSet(ctx context.Context, name string, args []string) error
	PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error)
	PluginCreate(ctx context.Context, createContext io.Reader, options types.PluginCreateOptions) error
}
```

---

### SecretAPIClient

SecretAPIClient defines API client methods for secrets

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L222)  

```go
type SecretAPIClient interface {
	SecretList(ctx context.Context, options swarm.SecretListOptions) ([]swarm.Secret, error)
	SecretCreate(ctx context.Context, secret swarm.SecretSpec) (swarm.SecretCreateResponse, error)
	SecretRemove(ctx context.Context, id string) error
	SecretInspectWithRaw(ctx context.Context, name string) (swarm.Secret, []byte, error)
	SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error
}
```

---

### ServiceAPIClient

ServiceAPIClient defines API client methods for the services

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L178)  

```go
type ServiceAPIClient interface {
	ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options swarm.ServiceCreateOptions) (swarm.ServiceCreateResponse, error)
	ServiceInspectWithRaw(ctx context.Context, serviceID string, options swarm.ServiceInspectOptions) (swarm.Service, []byte, error)
	ServiceList(ctx context.Context, options swarm.ServiceListOptions) ([]swarm.Service, error)
	ServiceRemove(ctx context.Context, serviceID string) error
	ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options swarm.ServiceUpdateOptions) (swarm.ServiceUpdateResponse, error)
	ServiceLogs(ctx context.Context, serviceID string, options container.LogsOptions) (io.ReadCloser, error)
	TaskLogs(ctx context.Context, taskID string, options container.LogsOptions) (io.ReadCloser, error)
	TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error)
	TaskList(ctx context.Context, options swarm.TaskListOptions) ([]swarm.Task, error)
}
```

---

### SwarmAPIClient

SwarmAPIClient defines API client methods for the swarm

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L191)  

```go
type SwarmAPIClient interface {
	SwarmInit(ctx context.Context, req swarm.InitRequest) (string, error)
	SwarmJoin(ctx context.Context, req swarm.JoinRequest) error
	SwarmGetUnlockKey(ctx context.Context) (swarm.UnlockKeyResponse, error)
	SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error
	SwarmLeave(ctx context.Context, force bool) error
	SwarmInspect(ctx context.Context) (swarm.Swarm, error)
	SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error
}
```

---

### SwarmManagementAPIClient

SwarmManagementAPIClient defines all methods for managing Swarm-specific
objects.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L57)  

```go
type SwarmManagementAPIClient interface {
	SwarmAPIClient
	NodeAPIClient
	ServiceAPIClient
	SecretAPIClient
	ConfigAPIClient
}
```

---

### SystemAPIClient

SystemAPIClient defines API client methods for the system

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L202)  

```go
type SystemAPIClient interface {
	Events(ctx context.Context, options events.ListOptions) (<-chan events.Message, <-chan error)
	Info(ctx context.Context) (system.Info, error)
	RegistryLogin(ctx context.Context, auth registry.AuthConfig) (registry.AuthenticateOKBody, error)
	DiskUsage(ctx context.Context, options types.DiskUsageOptions) (types.DiskUsage, error)
	Ping(ctx context.Context) (types.Ping, error)
}
```

---

### VolumeAPIClient

VolumeAPIClient defines API client methods for the volumes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/client/client_interfaces.go#L211)  

```go
type VolumeAPIClient interface {
	VolumeCreate(ctx context.Context, options volume.CreateOptions) (volume.Volume, error)
	VolumeInspect(ctx context.Context, volumeID string) (volume.Volume, error)
	VolumeInspectWithRaw(ctx context.Context, volumeID string) (volume.Volume, []byte, error)
	VolumeList(ctx context.Context, options volume.ListOptions) (volume.ListResponse, error)
	VolumeRemove(ctx context.Context, volumeID string, force bool) error
	VolumesPrune(ctx context.Context, pruneFilter filters.Args) (volume.PruneReport, error)
	VolumeUpdate(ctx context.Context, volumeID string, version swarm.Version, options volume.UpdateOptions) error
}
```

---


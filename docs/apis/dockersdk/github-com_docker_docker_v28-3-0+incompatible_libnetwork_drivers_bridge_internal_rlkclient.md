# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/bridge/internal/rlkclient

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:09:13 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### PortDriverClient

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/rlkclient/rootlesskit_client_linux.go#L23)  

```go
type PortDriverClient struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewPortDriverClient

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/rlkclient/rootlesskit_client_linux.go#L30)  

```go
func NewPortDriverClient(ctx context.Context) (*PortDriverClient, error)
```

#### Methods

##### PortDriverClient.AddPort

AddPort makes a request to RootlessKit asking it to set up a port
mapping between a host IP address and a child host IP address.

AddPort may return ProtocolUnsupportedError.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/rlkclient/rootlesskit_client_linux.go#L110)  

```go
func (c *PortDriverClient) AddPort(
	ctx context.Context,
	proto string,
	hostIP netip.Addr,
	childIP netip.Addr,
	hostPort int,
) (func() error, error)
```

##### PortDriverClient.ChildHostIP

ChildHostIP returns the address that must be used in the child network
namespace in place of hostIP, a host IP address. In particular, port
mappings from host IP addresses, and DNAT rules, must use this child
address in place of the real host address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/rlkclient/rootlesskit_client_linux.go#L82)  

```go
func (c *PortDriverClient) ChildHostIP(hostIP netip.Addr) netip.Addr
```

---

### ProtocolUnsupportedError

ProtocolUnsupportedError is returned when apiProto is not supported by portDriverName.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/rlkclient/rootlesskit_client_linux.go#L96)  

```go
type ProtocolUnsupportedError struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### ProtocolUnsupportedError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/rlkclient/rootlesskit_client_linux.go#L101)  

```go
func (e *ProtocolUnsupportedError) Error() string
```

---


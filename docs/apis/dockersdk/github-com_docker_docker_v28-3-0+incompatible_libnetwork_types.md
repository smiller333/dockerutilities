# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/types

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:11:04 UTC

## Overview

Package types contains types that are common across libnetwork project


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L17)

```go
const (
	IP = iota // IPv4 and IPv6
	IPv4
	IPv6
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L174)

```go
const (
	// ICMP is for the ICMP ip protocol
	ICMP = 1
	// TCP is for the TCP ip protocol
	TCP = 6
	// UDP is for the UDP ip protocol
	UDP = 17
	// SCTP is for the SCTP ip protocol
	SCTP = 132
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L344)

```go
const (
	// NEXTHOP indicates a StaticRoute with an IP next hop.
	NEXTHOP = iota

	// CONNECTED indicates a StaticRoute with an interface for directly connected peers.
	CONNECTED
)
```

## Variables

This section is empty.

## Functions

### CompareIPNet

CompareIPNet returns equal if the two IP Networks are equal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L260)  

```go
func CompareIPNet(a, b *net.IPNet) bool
```

---

### ForbiddenErrorf

ForbiddenErrorf creates an instance of ForbiddenError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L436)  

```go
func ForbiddenErrorf(format string, params ...interface{}) error
```

---

### GetBroadcastIP

GetBroadcastIP returns the broadcast ip address for the passed network (ip and mask).
IP address representation is not modified. If address and mask are not compatible
an error is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L318)  

```go
func GetBroadcastIP(ip net.IP, mask net.IPMask) (net.IP, error)
```

---

### GetHostPartIP

GetHostPartIP returns the host portion of the ip address identified by the mask.
IP address representation is not modified. If address and mask are not compatible
an error is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L299)  

```go
func GetHostPartIP(ip net.IP, mask net.IPMask) (net.IP, error)
```

---

### GetIPCopy

GetIPCopy returns a copy of the passed IP address

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L230)  

```go
func GetIPCopy(from net.IP) net.IP
```

---

### GetIPNetCanonical

GetIPNetCanonical returns the canonical form for the passed network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L250)  

```go
func GetIPNetCanonical(nw *net.IPNet) *net.IPNet
```

---

### GetIPNetCopy

GetIPNetCopy returns a copy of the passed IP Network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L240)  

```go
func GetIPNetCopy(from *net.IPNet) *net.IPNet
```

---

### GetMacCopy

GetMacCopy returns a copy of the passed MAC address

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L220)  

```go
func GetMacCopy(from net.HardwareAddr) net.HardwareAddr
```

---

### InternalErrorf

InternalErrorf creates an instance of InternalError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L451)  

```go
func InternalErrorf(format string, params ...interface{}) error
```

---

### InternalMaskableErrorf

InternalMaskableErrorf creates an instance of InternalError and MaskableError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L456)  

```go
func InternalMaskableErrorf(format string, params ...interface{}) error
```

---

### InvalidParameterErrorf

InvalidParameterErrorf creates an instance of InvalidParameterError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L426)  

```go
func InvalidParameterErrorf(format string, params ...interface{}) error
```

---

### IsIPNetValid

IsIPNetValid returns true if the ipnet is a valid network/mask
combination. Otherwise returns false.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L272)  

```go
func IsIPNetValid(nw *net.IPNet) bool
```

---

### NotFoundErrorf

NotFoundErrorf creates an instance of NotFoundError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L431)  

```go
func NotFoundErrorf(format string, params ...interface{}) error
```

---

### NotImplementedErrorf

NotImplementedErrorf creates an instance of NotImplementedError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L446)  

```go
func NotImplementedErrorf(format string, params ...interface{}) error
```

---

### ParseCIDR

ParseCIDR returns the *net.IPNet represented by the passed CIDR notation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L335)  

```go
func ParseCIDR(cidr string) (*net.IPNet, error)
```

---

### UnavailableErrorf

UnavailableErrorf creates an instance of UnavailableError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L441)  

```go
func UnavailableErrorf(format string, params ...interface{}) error
```

---

## Types

### EncryptionKey

EncryptionKey is the libnetwork representation of the key distributed by the lead
manager.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L25)  

```go
type EncryptionKey struct {
	Subsystem   string
	Algorithm   int32
	Key         []byte
	LamportTime uint64
}
```

---

### ForbiddenError

ForbiddenError is an interface for errors which denote a valid request that cannot be honored

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L407)  

```go
type ForbiddenError = errdefs.ErrForbidden
```

---

### InterfaceStatistics

InterfaceStatistics represents the interface's statistics

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L374)  

```go
type InterfaceStatistics struct {
	RxBytes   uint64
	RxPackets uint64
	RxErrors  uint64
	RxDropped uint64
	TxBytes   uint64
	TxPackets uint64
	TxErrors  uint64
	TxDropped uint64
}
```

#### Methods

##### InterfaceStatistics.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L385)  

```go
func (is *InterfaceStatistics) String() string
```

---

### InternalError

InternalError is an interface for errors raised because of an internal error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L416)  

```go
type InternalError interface {
	// Internal makes implementer into InternalError type
	Internal()
}
```

---

### InvalidParameterError

InvalidParameterError is an interface for errors originated by a bad request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L401)  

```go
type InvalidParameterError = errdefs.ErrInvalidParameter
```

---

### MaskableError

MaskableError is an interface for errors which can be ignored by caller

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L395)  

```go
type MaskableError interface {
	// Maskable makes implementer into MaskableError type
	Maskable()
}
```

---

### NotFoundError

NotFoundError is an interface for errors raised because a needed resource is not available

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L404)  

```go
type NotFoundError = errdefs.ErrNotFound
```

---

### NotImplementedError

NotImplementedError is an interface for errors raised because of requested functionality is not yet implemented

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L413)  

```go
type NotImplementedError = errdefs.ErrNotImplemented
```

---

### PortBinding

PortBinding represents a port binding between the container and the host

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L71)  

```go
type PortBinding struct {
	Proto       Protocol
	IP          net.IP
	Port        uint16
	HostIP      net.IP
	HostPort    uint16
	HostPortEnd uint16
}
```

#### Methods

##### PortBinding.ContainerAddr

ContainerAddr returns the container side transport address

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L95)  

```go
func (p PortBinding) ContainerAddr() (net.Addr, error)
```

##### PortBinding.Equal

Equal returns true if o has the same values as p, else false.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L121)  

```go
func (p *PortBinding) Equal(o *PortBinding) bool
```

##### PortBinding.GetCopy

GetCopy returns a copy of this PortBinding structure instance

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L109)  

```go
func (p *PortBinding) GetCopy() PortBinding
```

##### PortBinding.HostAddr

HostAddr returns the host side transport address

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L81)  

```go
func (p PortBinding) HostAddr() (net.Addr, error)
```

##### PortBinding.String

String returns the PortBinding structure in the form "HostIP:HostPort:IP:Port/Proto",
omitting un-set fields apart from Port.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L132)  

```go
func (p PortBinding) String() string
```

---

### Protocol

Protocol represents an IP protocol number

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L186)  

```go
type Protocol uint8
```

#### Functions

##### ParseProtocol

ParseProtocol returns the respective Protocol type for the passed string

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L204)  

```go
func ParseProtocol(s string) Protocol
```

#### Methods

##### Protocol.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L188)  

```go
func (p Protocol) String() string
```

---

### QosPolicy

QosPolicy represents a quality of service policy on an endpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L33)  

```go
type QosPolicy struct {
	MaxEgressBandwidth uint64
}
```

---

### StaticRoute

StaticRoute is a statically-provisioned IP route.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L353)  

```go
type StaticRoute struct {
	Destination *net.IPNet

	RouteType int // NEXT_HOP or CONNECTED

	// NextHop will be resolved by the kernel (i.e. as a loose hop).
	NextHop net.IP
}
```

#### Methods

##### StaticRoute.GetCopy

GetCopy returns a copy of this StaticRoute structure

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L363)  

```go
func (r *StaticRoute) GetCopy() *StaticRoute
```

---

### TransportPort

TransportPort represents a local Layer 4 endpoint

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L38)  

```go
type TransportPort struct {
	Proto Protocol
	Port  uint16
}
```

#### Methods

##### TransportPort.Equal

Equal checks if this instance of TransportPort is equal to the passed one

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L44)  

```go
func (t *TransportPort) Equal(o *TransportPort) bool
```

##### TransportPort.GetCopy

GetCopy returns a copy of this TransportPort structure instance

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L61)  

```go
func (t *TransportPort) GetCopy() TransportPort
```

##### TransportPort.String

String returns the TransportPort structure in string form

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L66)  

```go
func (t *TransportPort) String() string
```

---

### UnavailableError

UnavailableError is an interface for errors returned when the required service is not available

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/types/types.go#L410)  

```go
type UnavailableError = errdefs.ErrUnavailable
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/windows/overlay

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:33:50 UTC

## Overview

Package overlay is a generated protocol buffer package.

It is generated from these files:

It has these top-level messages:


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L434)

```go
var (
	ErrInvalidLengthOverlay = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOverlay   = fmt.Errorf("proto: integer overflow")
)
```

## Functions

This section is empty.

## Types

### PeerRecord

PeerRecord defines the information corresponding to a peer
container in the overlay network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L40)  

```go
type PeerRecord struct {
	// Endpoint IP is the IP of the container attachment on the
	// given overlay network.
	EndpointIP string `protobuf:"bytes,1,opt,name=endpoint_ip,json=endpointIp,proto3" json:"endpoint_ip,omitempty"`
	// Endpoint MAC is the mac address of the container attachment
	// on the given overlay network.
	EndpointMAC string `protobuf:"bytes,2,opt,name=endpoint_mac,json=endpointMac,proto3" json:"endpoint_mac,omitempty"`
	// Tunnel Endpoint IP defines the host IP for the host in
	// which this container is running and can be reached by
	// building a tunnel to that host IP.
	TunnelEndpointIP string `protobuf:"bytes,3,opt,name=tunnel_endpoint_ip,json=tunnelEndpointIp,proto3" json:"tunnel_endpoint_ip,omitempty"`
}
```

#### Methods

##### PeerRecord.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L55)  

```go
func (*PeerRecord) Descriptor() ([]byte, []int)
```

##### PeerRecord.GetEndpointIP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L57)  

```go
func (m *PeerRecord) GetEndpointIP() string
```

##### PeerRecord.GetEndpointMAC

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L64)  

```go
func (m *PeerRecord) GetEndpointMAC() string
```

##### PeerRecord.GetTunnelEndpointIP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L71)  

```go
func (m *PeerRecord) GetTunnelEndpointIP() string
```

##### PeerRecord.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L81)  

```go
func (this *PeerRecord) GoString() string
```

##### PeerRecord.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L101)  

```go
func (m *PeerRecord) Marshal() (dAtA []byte, err error)
```

##### PeerRecord.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L111)  

```go
func (m *PeerRecord) MarshalTo(dAtA []byte) (int, error)
```

##### PeerRecord.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L54)  

```go
func (*PeerRecord) ProtoMessage()
```

##### PeerRecord.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L53)  

```go
func (m *PeerRecord) Reset()
```

##### PeerRecord.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L146)  

```go
func (m *PeerRecord) Size() (n int)
```

##### PeerRecord.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L177)  

```go
func (this *PeerRecord) String() string
```

##### PeerRecord.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/overlay/overlay.pb.go#L197)  

```go
func (m *PeerRecord) Unmarshal(dAtA []byte) error
```

---


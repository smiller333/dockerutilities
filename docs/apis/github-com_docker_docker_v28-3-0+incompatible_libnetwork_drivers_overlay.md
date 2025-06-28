# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/overlay

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:09:30 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.go#L19)

```go
const (
	NetworkType = "overlay"
)
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L468)

```go
var (
	ErrInvalidLengthOverlay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOverlay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOverlay = fmt.Errorf("proto: unexpected end of group")
)
```

## Functions

### Register

Register registers a new instance of the overlay driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.go#L45)  

```go
func Register(r driverapi.Registerer, config map[string]interface{}) error
```

---

## Types

### PeerRecord

PeerRecord defines the information corresponding to a peer
container in the overlay network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L30)  

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

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L45)  

```go
func (*PeerRecord) Descriptor() ([]byte, []int)
```

##### PeerRecord.GetEndpointIP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L75)  

```go
func (m *PeerRecord) GetEndpointIP() string
```

##### PeerRecord.GetEndpointMAC

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L82)  

```go
func (m *PeerRecord) GetEndpointMAC() string
```

##### PeerRecord.GetTunnelEndpointIP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L89)  

```go
func (m *PeerRecord) GetTunnelEndpointIP() string
```

##### PeerRecord.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L121)  

```go
func (this *PeerRecord) GoString() string
```

##### PeerRecord.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L141)  

```go
func (m *PeerRecord) Marshal() (dAtA []byte, err error)
```

##### PeerRecord.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L151)  

```go
func (m *PeerRecord) MarshalTo(dAtA []byte) (int, error)
```

##### PeerRecord.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L156)  

```go
func (m *PeerRecord) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### PeerRecord.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L44)  

```go
func (*PeerRecord) ProtoMessage()
```

##### PeerRecord.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L43)  

```go
func (m *PeerRecord) Reset()
```

##### PeerRecord.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L196)  

```go
func (m *PeerRecord) Size() (n int)
```

##### PeerRecord.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L223)  

```go
func (this *PeerRecord) String() string
```

##### PeerRecord.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L243)  

```go
func (m *PeerRecord) Unmarshal(dAtA []byte) error
```

##### PeerRecord.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L69)  

```go
func (m *PeerRecord) XXX_DiscardUnknown()
```

##### PeerRecord.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L51)  

```go
func (m *PeerRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### PeerRecord.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L63)  

```go
func (m *PeerRecord) XXX_Merge(src proto.Message)
```

##### PeerRecord.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L66)  

```go
func (m *PeerRecord) XXX_Size() int
```

##### PeerRecord.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlay.pb.go#L48)  

```go
func (m *PeerRecord) XXX_Unmarshal(b []byte) error
```

---


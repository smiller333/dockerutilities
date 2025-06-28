# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/networkdb

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:48 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/watch.go#L25)

```go
const NodeTable = "NodeTable"
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L3087)

```go
var (
	ErrInvalidLengthNetworkdb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkdb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNetworkdb = fmt.Errorf("proto: unexpected end of group")
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L58)

```go
var MessageType_name = map[int32]string{
	0: "INVALID",
	1: "NETWORK_EVENT",
	2: "TABLE_EVENT",
	3: "PUSH_PULL",
	4: "BULK_SYNC",
	5: "COMPOUND",
	6: "NODE_EVENT",
}
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L68)

```go
var MessageType_value = map[string]int32{
	"INVALID":       0,
	"NETWORK_EVENT": 1,
	"TABLE_EVENT":   2,
	"PUSH_PULL":     3,
	"BULK_SYNC":     4,
	"COMPOUND":      5,
	"NODE_EVENT":    6,
}
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L126)

```go
var NetworkEvent_Type_name = map[int32]string{
	0: "INVALID",
	1: "JOIN",
	2: "LEAVE",
}
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L132)

```go
var NetworkEvent_Type_value = map[string]int32{
	"INVALID": 0,
	"JOIN":    1,
	"LEAVE":   2,
}
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L96)

```go
var NodeEvent_Type_name = map[int32]string{
	0: "INVALID",
	1: "JOIN",
	2: "LEAVE",
}
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L102)

```go
var NodeEvent_Type_value = map[string]int32{
	"INVALID": 0,
	"JOIN":    1,
	"LEAVE":   2,
}
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L161)

```go
var TableEvent_Type_name = map[int32]string{
	0: "INVALID",
	1: "CREATE",
	2: "UPDATE",
	3: "DELETE",
}
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L168)

```go
var TableEvent_Type_value = map[string]int32{
	"INVALID": 0,
	"CREATE":  1,
	"UPDATE":  2,
	"DELETE":  3,
}
```

## Functions

This section is empty.

## Types

### BulkSyncMessage

BulkSync message payload definition.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L580)  

```go
type BulkSyncMessage struct {
	// Lamport time when this bulk sync was initiated.
	LTime github_com_hashicorp_serf_serf.LamportTime `protobuf:"varint,1,opt,name=l_time,json=lTime,proto3,customtype=github.com/hashicorp/serf/serf.LamportTime" json:"l_time"`
	// Indicates if this bulksync is a response to a bulk sync
	// request from a peer node.
	Unsolicited bool `protobuf:"varint,2,opt,name=unsolicited,proto3" json:"unsolicited,omitempty"`
	// Name of the node which is producing this bulk sync message.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// List of network names whose table entries are getting
	// bulksynced as part of the bulksync.
	Networks []string `protobuf:"bytes,4,rep,name=networks,proto3" json:"networks,omitempty"`
	// Bulksync payload
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}
```

#### Methods

##### BulkSyncMessage.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L597)  

```go
func (*BulkSyncMessage) Descriptor() ([]byte, []int)
```

##### BulkSyncMessage.GetNetworks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L641)  

```go
func (m *BulkSyncMessage) GetNetworks() []string
```

##### BulkSyncMessage.GetNodeName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L634)  

```go
func (m *BulkSyncMessage) GetNodeName() string
```

##### BulkSyncMessage.GetPayload

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L648)  

```go
func (m *BulkSyncMessage) GetPayload() []byte
```

##### BulkSyncMessage.GetUnsolicited

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L627)  

```go
func (m *BulkSyncMessage) GetUnsolicited() bool
```

##### BulkSyncMessage.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L908)  

```go
func (this *BulkSyncMessage) GoString() string
```

##### BulkSyncMessage.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1248)  

```go
func (m *BulkSyncMessage) Marshal() (dAtA []byte, err error)
```

##### BulkSyncMessage.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1258)  

```go
func (m *BulkSyncMessage) MarshalTo(dAtA []byte) (int, error)
```

##### BulkSyncMessage.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1263)  

```go
func (m *BulkSyncMessage) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### BulkSyncMessage.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L596)  

```go
func (*BulkSyncMessage) ProtoMessage()
```

##### BulkSyncMessage.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L595)  

```go
func (m *BulkSyncMessage) Reset()
```

##### BulkSyncMessage.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1528)  

```go
func (m *BulkSyncMessage) Size() (n int)
```

##### BulkSyncMessage.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1674)  

```go
func (this *BulkSyncMessage) String() string
```

##### BulkSyncMessage.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L2653)  

```go
func (m *BulkSyncMessage) Unmarshal(dAtA []byte) error
```

##### BulkSyncMessage.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L621)  

```go
func (m *BulkSyncMessage) XXX_DiscardUnknown()
```

##### BulkSyncMessage.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L603)  

```go
func (m *BulkSyncMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### BulkSyncMessage.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L615)  

```go
func (m *BulkSyncMessage) XXX_Merge(src proto.Message)
```

##### BulkSyncMessage.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L618)  

```go
func (m *BulkSyncMessage) XXX_Size() int
```

##### BulkSyncMessage.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L600)  

```go
func (m *BulkSyncMessage) XXX_Unmarshal(b []byte) error
```

---

### CompoundMessage

Compound message payload definition.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L656)  

```go
type CompoundMessage struct {
	// A list of simple messages.
	Messages []*CompoundMessage_SimpleMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}
```

#### Methods

##### CompoundMessage.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L663)  

```go
func (*CompoundMessage) Descriptor() ([]byte, []int)
```

##### CompoundMessage.GetMessages

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L693)  

```go
func (m *CompoundMessage) GetMessages() []*CompoundMessage_SimpleMessage
```

##### CompoundMessage.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L922)  

```go
func (this *CompoundMessage) GoString() string
```

##### CompoundMessage.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1309)  

```go
func (m *CompoundMessage) Marshal() (dAtA []byte, err error)
```

##### CompoundMessage.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1319)  

```go
func (m *CompoundMessage) MarshalTo(dAtA []byte) (int, error)
```

##### CompoundMessage.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1324)  

```go
func (m *CompoundMessage) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### CompoundMessage.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L662)  

```go
func (*CompoundMessage) ProtoMessage()
```

##### CompoundMessage.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L661)  

```go
func (m *CompoundMessage) Reset()
```

##### CompoundMessage.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1557)  

```go
func (m *CompoundMessage) Size() (n int)
```

##### CompoundMessage.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1688)  

```go
func (this *CompoundMessage) String() string
```

##### CompoundMessage.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L2840)  

```go
func (m *CompoundMessage) Unmarshal(dAtA []byte) error
```

##### CompoundMessage.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L687)  

```go
func (m *CompoundMessage) XXX_DiscardUnknown()
```

##### CompoundMessage.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L669)  

```go
func (m *CompoundMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### CompoundMessage.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L681)  

```go
func (m *CompoundMessage) XXX_Merge(src proto.Message)
```

##### CompoundMessage.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L684)  

```go
func (m *CompoundMessage) XXX_Size() int
```

##### CompoundMessage.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L666)  

```go
func (m *CompoundMessage) XXX_Unmarshal(b []byte) error
```

---

### CompoundMessage_SimpleMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L700)  

```go
type CompoundMessage_SimpleMessage struct {
	// Bytestring payload of a message constructed using
	// other message type definitions.
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
}
```

#### Methods

##### CompoundMessage_SimpleMessage.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L708)  

```go
func (*CompoundMessage_SimpleMessage) Descriptor() ([]byte, []int)
```

##### CompoundMessage_SimpleMessage.GetPayload

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L738)  

```go
func (m *CompoundMessage_SimpleMessage) GetPayload() []byte
```

##### CompoundMessage_SimpleMessage.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L934)  

```go
func (this *CompoundMessage_SimpleMessage) GoString() string
```

##### CompoundMessage_SimpleMessage.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1346)  

```go
func (m *CompoundMessage_SimpleMessage) Marshal() (dAtA []byte, err error)
```

##### CompoundMessage_SimpleMessage.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1356)  

```go
func (m *CompoundMessage_SimpleMessage) MarshalTo(dAtA []byte) (int, error)
```

##### CompoundMessage_SimpleMessage.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1361)  

```go
func (m *CompoundMessage_SimpleMessage) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### CompoundMessage_SimpleMessage.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L707)  

```go
func (*CompoundMessage_SimpleMessage) ProtoMessage()
```

##### CompoundMessage_SimpleMessage.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L706)  

```go
func (m *CompoundMessage_SimpleMessage) Reset()
```

##### CompoundMessage_SimpleMessage.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1572)  

```go
func (m *CompoundMessage_SimpleMessage) Size() (n int)
```

##### CompoundMessage_SimpleMessage.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1703)  

```go
func (this *CompoundMessage_SimpleMessage) String() string
```

##### CompoundMessage_SimpleMessage.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L2924)  

```go
func (m *CompoundMessage_SimpleMessage) Unmarshal(dAtA []byte) error
```

##### CompoundMessage_SimpleMessage.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L732)  

```go
func (m *CompoundMessage_SimpleMessage) XXX_DiscardUnknown()
```

##### CompoundMessage_SimpleMessage.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L714)  

```go
func (m *CompoundMessage_SimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### CompoundMessage_SimpleMessage.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L726)  

```go
func (m *CompoundMessage_SimpleMessage) XXX_Merge(src proto.Message)
```

##### CompoundMessage_SimpleMessage.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L729)  

```go
func (m *CompoundMessage_SimpleMessage) XXX_Size() int
```

##### CompoundMessage_SimpleMessage.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L711)  

```go
func (m *CompoundMessage_SimpleMessage) XXX_Unmarshal(b []byte) error
```

---

### Config

Config represents the configuration of the networkdb instance and
can be passed by the caller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L162)  

```go
type Config struct {
	// NodeID is the node unique identifier of the node when is part of the cluster
	NodeID string

	// Hostname is the node hostname.
	Hostname string

	// BindAddr is the IP on which networkdb listens. It can be
	// 0.0.0.0 to listen on all addresses on the host.
	BindAddr string

	// AdvertiseAddr is the node's IP address that we advertise for
	// cluster communication.
	AdvertiseAddr string

	// BindPort is the local node's port to which we bind to for
	// cluster communication.
	BindPort int

	// Keys to be added to the Keyring of the memberlist. Key at index
	// 0 is the primary key
	Keys [][]byte

	// PacketBufferSize is the maximum number of bytes that memberlist will
	// put in a packet (this will be for UDP packets by default with a NetTransport).
	// A safe value for this is typically 1400 bytes (which is the default). However,
	// depending on your network's MTU (Maximum Transmission Unit) you may
	// be able to increase this to get more content into each gossip packet.
	PacketBufferSize int

	// StatsPrintPeriod the period to use to print queue stats
	// Default is 5min
	StatsPrintPeriod time.Duration

	// HealthPrintPeriod the period to use to print the health score
	// Default is 1min
	HealthPrintPeriod time.Duration
	// contains filtered or unexported fields
}
```

#### Functions

##### DefaultConfig

DefaultConfig returns a NetworkDB config with default values

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L237)  

```go
func DefaultConfig() *Config
```

---

### CreateEvent

CreateEvent generates a table entry create event to the watchers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/watch.go#L33)  

```go
type CreateEvent event
```

---

### DeleteEvent

DeleteEvent generates a table entry delete event to the watchers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/watch.go#L39)  

```go
type DeleteEvent event
```

---

### GossipMessage

GossipMessage is a basic message header used by all messages types.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L184)  

```go
type GossipMessage struct {
	Type MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=networkdb.MessageType" json:"type,omitempty"`
	Data []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}
```

#### Methods

##### GossipMessage.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L191)  

```go
func (*GossipMessage) Descriptor() ([]byte, []int)
```

##### GossipMessage.GetData

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L228)  

```go
func (m *GossipMessage) GetData() []byte
```

##### GossipMessage.GetType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L221)  

```go
func (m *GossipMessage) GetType() MessageType
```

##### GossipMessage.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L828)  

```go
func (this *GossipMessage) GoString() string
```

##### GossipMessage.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L952)  

```go
func (m *GossipMessage) Marshal() (dAtA []byte, err error)
```

##### GossipMessage.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L962)  

```go
func (m *GossipMessage) MarshalTo(dAtA []byte) (int, error)
```

##### GossipMessage.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L967)  

```go
func (m *GossipMessage) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### GossipMessage.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L190)  

```go
func (*GossipMessage) ProtoMessage()
```

##### GossipMessage.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L189)  

```go
func (m *GossipMessage) Reset()
```

##### GossipMessage.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1387)  

```go
func (m *GossipMessage) Size() (n int)
```

##### GossipMessage.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1591)  

```go
func (this *GossipMessage) String() string
```

##### GossipMessage.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1721)  

```go
func (m *GossipMessage) Unmarshal(dAtA []byte) error
```

##### GossipMessage.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L215)  

```go
func (m *GossipMessage) XXX_DiscardUnknown()
```

##### GossipMessage.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L197)  

```go
func (m *GossipMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### GossipMessage.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L209)  

```go
func (m *GossipMessage) XXX_Merge(src proto.Message)
```

##### GossipMessage.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L212)  

```go
func (m *GossipMessage) XXX_Size() int
```

##### GossipMessage.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L194)  

```go
func (m *GossipMessage) XXX_Unmarshal(b []byte) error
```

---

### MessageType

MessageType enum defines all the core message types that networkdb
uses to communicate to peers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L31)  

```go
type MessageType int32
```

#### Methods

##### MessageType.EnumDescriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L82)  

```go
func (MessageType) EnumDescriptor() ([]byte, []int)
```

##### MessageType.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L78)  

```go
func (x MessageType) String() string
```

---

### Mux

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdbdiagnostic.go#L20)  

```go
type Mux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}
```

---

### NetworkDB

NetworkDB instance drives the networkdb cluster and acts the broker
for cluster-scoped and network-scoped gossip and watches.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L33)  

```go
type NetworkDB struct {
	sync.RWMutex
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates a new instance of NetworkDB using the Config passed by
the caller.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L254)  

```go
func New(c *Config) (*NetworkDB, error)
```

#### Methods

##### NetworkDB.Close

Close destroys this NetworkDB instance by leave the cluster,
stopping timers, canceling goroutines etc.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L295)  

```go
func (nDB *NetworkDB) Close()
```

##### NetworkDB.ClusterPeers

ClusterPeers returns all the gossip cluster peers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L305)  

```go
func (nDB *NetworkDB) ClusterPeers() []PeerInfo
```

##### NetworkDB.CreateEntry

CreateEntry creates a table entry in NetworkDB for given (network,
table, key) tuple and if the NetworkDB is part of the cluster
propagates this event to the cluster. It is an error to create an
entry for the same tuple for which there is already an existing
entry unless the current entry is deleting state.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L372)  

```go
func (nDB *NetworkDB) CreateEntry(tname, nid, key string, value []byte) error
```

##### NetworkDB.DeleteEntry

DeleteEntry deletes a table entry in NetworkDB for given (network,
table, key) tuple and if the NetworkDB is part of the cluster
propagates this event to the cluster.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L448)  

```go
func (nDB *NetworkDB) DeleteEntry(tname, nid, key string) error
```

##### NetworkDB.GetEntry

GetEntry retrieves the value of a table entry in a given (network,
table, key) tuple

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L340)  

```go
func (nDB *NetworkDB) GetEntry(tname, nid, key string) ([]byte, error)
```

##### NetworkDB.GetTableByNetwork

GetTableByNetwork walks the networkdb by the give table and network id and
returns a map of keys and values

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L431)  

```go
func (nDB *NetworkDB) GetTableByNetwork(tname, nid string) map[string]*TableElem
```

##### NetworkDB.Join

Join joins this NetworkDB instance with a list of peer NetworkDB
instances passed by the caller in the form of addr:port

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L285)  

```go
func (nDB *NetworkDB) Join(members []string) error
```

##### NetworkDB.JoinNetwork

JoinNetwork joins this node to a given network and propagates this
event across the cluster. This triggers this node joining the
sub-cluster of this network and participates in the network-scoped
gossip and bulk sync for this network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L611)  

```go
func (nDB *NetworkDB) JoinNetwork(nid string) error
```

##### NetworkDB.LeaveNetwork

LeaveNetwork leaves this node from a given network and propagates
this event across the cluster. This triggers this node leaving the
sub-cluster of this network and as a result will no longer
participate in the network-scoped gossip and bulk sync for this
network. Also remove all the table entries for this network from
networkdb

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L667)  

```go
func (nDB *NetworkDB) LeaveNetwork(nid string) error
```

##### NetworkDB.Peers

Peers returns the gossip peers for a given network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L319)  

```go
func (nDB *NetworkDB) Peers(nid string) []PeerInfo
```

##### NetworkDB.RegisterDiagnosticHandlers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdbdiagnostic.go#L24)  

```go
func (nDB *NetworkDB) RegisterDiagnosticHandlers(m Mux)
```

##### NetworkDB.RemoveKey

RemoveKey removes a key from the key ring. The key being removed
can't be the primary key

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/cluster.go#L88)  

```go
func (nDB *NetworkDB) RemoveKey(key []byte)
```

##### NetworkDB.SetKey

SetKey adds a new key to the key ring

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/cluster.go#L55)  

```go
func (nDB *NetworkDB) SetKey(key []byte)
```

##### NetworkDB.SetPrimaryKey

SetPrimaryKey sets the given key as the primary key. This should have
been added apriori through SetKey

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/cluster.go#L72)  

```go
func (nDB *NetworkDB) SetPrimaryKey(key []byte)
```

##### NetworkDB.UpdateEntry

UpdateEntry updates a table entry in NetworkDB for given (network,
table, key) tuple and if the NetworkDB is part of the cluster
propagates this event to the cluster. It is an error to update a
non-existent entry.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L400)  

```go
func (nDB *NetworkDB) UpdateEntry(tname, nid, key string, value []byte) error
```

##### NetworkDB.WalkTable

WalkTable walks a single table in NetworkDB and invokes the passed
function for each entry in the table passing the network, key,
value. The walk stops if the passed function returns a true.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L586)  

```go
func (nDB *NetworkDB) WalkTable(tname string, fn func(string, string, []byte, bool) bool) error
```

##### NetworkDB.Watch

Watch creates a watcher with filters for a particular table or
network or any combination of the tuple. If any of the
filter is an empty string it acts as a wildcard for that
field. Watch returns a channel of events, where the events will be
sent.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/watch.go#L46)  

```go
func (nDB *NetworkDB) Watch(tname, nid string) (*events.Channel, func())
```

---

### NetworkEntry

NetworkEntry for push pull of networks.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L359)  

```go
type NetworkEntry struct {
	// ID of the network
	NetworkID string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Latest lamport time of the network attachment when this
	// network event was recorded.
	LTime github_com_hashicorp_serf_serf.LamportTime `protobuf:"varint,2,opt,name=l_time,json=lTime,proto3,customtype=github.com/hashicorp/serf/serf.LamportTime" json:"l_time"`
	// Source node name where this network attachment happened.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Indicates if a leave from this network is in progress.
	Leaving bool `protobuf:"varint,4,opt,name=leaving,proto3" json:"leaving,omitempty"`
}
```

#### Methods

##### NetworkEntry.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L373)  

```go
func (*NetworkEntry) Descriptor() ([]byte, []int)
```

##### NetworkEntry.GetLeaving

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L417)  

```go
func (m *NetworkEntry) GetLeaving() bool
```

##### NetworkEntry.GetNetworkID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L403)  

```go
func (m *NetworkEntry) GetNetworkID() string
```

##### NetworkEntry.GetNodeName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L410)  

```go
func (m *NetworkEntry) GetNodeName() string
```

##### NetworkEntry.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L864)  

```go
func (this *NetworkEntry) GoString() string
```

##### NetworkEntry.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1074)  

```go
func (m *NetworkEntry) Marshal() (dAtA []byte, err error)
```

##### NetworkEntry.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1084)  

```go
func (m *NetworkEntry) MarshalTo(dAtA []byte) (int, error)
```

##### NetworkEntry.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1089)  

```go
func (m *NetworkEntry) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### NetworkEntry.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L372)  

```go
func (*NetworkEntry) ProtoMessage()
```

##### NetworkEntry.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L371)  

```go
func (m *NetworkEntry) Reset()
```

##### NetworkEntry.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1445)  

```go
func (m *NetworkEntry) Size() (n int)
```

##### NetworkEntry.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1627)  

```go
func (this *NetworkEntry) String() string
```

##### NetworkEntry.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L2096)  

```go
func (m *NetworkEntry) Unmarshal(dAtA []byte) error
```

##### NetworkEntry.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L397)  

```go
func (m *NetworkEntry) XXX_DiscardUnknown()
```

##### NetworkEntry.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L379)  

```go
func (m *NetworkEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### NetworkEntry.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L391)  

```go
func (m *NetworkEntry) XXX_Merge(src proto.Message)
```

##### NetworkEntry.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L394)  

```go
func (m *NetworkEntry) XXX_Size() int
```

##### NetworkEntry.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L376)  

```go
func (m *NetworkEntry) XXX_Unmarshal(b []byte) error
```

---

### NetworkEvent

NetworkEvent message payload definition.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L293)  

```go
type NetworkEvent struct {
	Type NetworkEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=networkdb.NetworkEvent_Type" json:"type,omitempty"`
	// Lamport time using a network lamport clock indicating the
	// time this event was generated on the node where it was
	// generated.
	LTime github_com_hashicorp_serf_serf.LamportTime `protobuf:"varint,2,opt,name=l_time,json=lTime,proto3,customtype=github.com/hashicorp/serf/serf.LamportTime" json:"l_time"`
	// Source node name.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// ID of the network for which the event is generated.
	NetworkID string `protobuf:"bytes,4,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}
```

#### Methods

##### NetworkEvent.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L307)  

```go
func (*NetworkEvent) Descriptor() ([]byte, []int)
```

##### NetworkEvent.GetNetworkID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L351)  

```go
func (m *NetworkEvent) GetNetworkID() string
```

##### NetworkEvent.GetNodeName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L344)  

```go
func (m *NetworkEvent) GetNodeName() string
```

##### NetworkEvent.GetType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L337)  

```go
func (m *NetworkEvent) GetType() NetworkEvent_Type
```

##### NetworkEvent.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L851)  

```go
func (this *NetworkEvent) GoString() string
```

##### NetworkEvent.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1027)  

```go
func (m *NetworkEvent) Marshal() (dAtA []byte, err error)
```

##### NetworkEvent.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1037)  

```go
func (m *NetworkEvent) MarshalTo(dAtA []byte) (int, error)
```

##### NetworkEvent.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1042)  

```go
func (m *NetworkEvent) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### NetworkEvent.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L306)  

```go
func (*NetworkEvent) ProtoMessage()
```

##### NetworkEvent.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L305)  

```go
func (m *NetworkEvent) Reset()
```

##### NetworkEvent.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1422)  

```go
func (m *NetworkEvent) Size() (n int)
```

##### NetworkEvent.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1614)  

```go
func (this *NetworkEvent) String() string
```

##### NetworkEvent.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1944)  

```go
func (m *NetworkEvent) Unmarshal(dAtA []byte) error
```

##### NetworkEvent.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L331)  

```go
func (m *NetworkEvent) XXX_DiscardUnknown()
```

##### NetworkEvent.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L313)  

```go
func (m *NetworkEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### NetworkEvent.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L325)  

```go
func (m *NetworkEvent) XXX_Merge(src proto.Message)
```

##### NetworkEvent.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L328)  

```go
func (m *NetworkEvent) XXX_Size() int
```

##### NetworkEvent.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L310)  

```go
func (m *NetworkEvent) XXX_Unmarshal(b []byte) error
```

---

### NetworkEvent_Type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L116)  

```go
type NetworkEvent_Type int32
```

#### Methods

##### NetworkEvent_Type.EnumDescriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L142)  

```go
func (NetworkEvent_Type) EnumDescriptor() ([]byte, []int)
```

##### NetworkEvent_Type.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L138)  

```go
func (x NetworkEvent_Type) String() string
```

---

### NetworkPushPull

NetworkPushpull message payload definition.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L425)  

```go
type NetworkPushPull struct {
	// Lamport time when this push pull was initiated.
	LTime    github_com_hashicorp_serf_serf.LamportTime `protobuf:"varint,1,opt,name=l_time,json=lTime,proto3,customtype=github.com/hashicorp/serf/serf.LamportTime" json:"l_time"`
	Networks []*NetworkEntry                            `protobuf:"bytes,2,rep,name=networks,proto3" json:"networks,omitempty"`
	// Name of the node sending this push pull payload.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
}
```

#### Methods

##### NetworkPushPull.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L435)  

```go
func (*NetworkPushPull) Descriptor() ([]byte, []int)
```

##### NetworkPushPull.GetNetworks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L465)  

```go
func (m *NetworkPushPull) GetNetworks() []*NetworkEntry
```

##### NetworkPushPull.GetNodeName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L472)  

```go
func (m *NetworkPushPull) GetNodeName() string
```

##### NetworkPushPull.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L877)  

```go
func (this *NetworkPushPull) GoString() string
```

##### NetworkPushPull.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1126)  

```go
func (m *NetworkPushPull) Marshal() (dAtA []byte, err error)
```

##### NetworkPushPull.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1136)  

```go
func (m *NetworkPushPull) MarshalTo(dAtA []byte) (int, error)
```

##### NetworkPushPull.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1141)  

```go
func (m *NetworkPushPull) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### NetworkPushPull.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L434)  

```go
func (*NetworkPushPull) ProtoMessage()
```

##### NetworkPushPull.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L433)  

```go
func (m *NetworkPushPull) Reset()
```

##### NetworkPushPull.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1468)  

```go
func (m *NetworkPushPull) Size() (n int)
```

##### NetworkPushPull.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1640)  

```go
func (this *NetworkPushPull) String() string
```

##### NetworkPushPull.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L2249)  

```go
func (m *NetworkPushPull) Unmarshal(dAtA []byte) error
```

##### NetworkPushPull.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L459)  

```go
func (m *NetworkPushPull) XXX_DiscardUnknown()
```

##### NetworkPushPull.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L441)  

```go
func (m *NetworkPushPull) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### NetworkPushPull.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L453)  

```go
func (m *NetworkPushPull) XXX_Merge(src proto.Message)
```

##### NetworkPushPull.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L456)  

```go
func (m *NetworkPushPull) XXX_Size() int
```

##### NetworkPushPull.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L438)  

```go
func (m *NetworkPushPull) XXX_Unmarshal(b []byte) error
```

---

### NodeAddr

NodeAddr represents the value carried for node event in NodeTable

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/watch.go#L28)  

```go
type NodeAddr struct {
	Addr net.IP
}
```

---

### NodeEvent

NodeEvent message payload definition.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L236)  

```go
type NodeEvent struct {
	Type NodeEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=networkdb.NodeEvent_Type" json:"type,omitempty"`
	// Lamport time using a network lamport clock indicating the
	// time this event was generated on the node where it was
	// generated.
	LTime github_com_hashicorp_serf_serf.LamportTime `protobuf:"varint,2,opt,name=l_time,json=lTime,proto3,customtype=github.com/hashicorp/serf/serf.LamportTime" json:"l_time"`
	// Source node name.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
}
```

#### Methods

##### NodeEvent.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L248)  

```go
func (*NodeEvent) Descriptor() ([]byte, []int)
```

##### NodeEvent.GetNodeName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L285)  

```go
func (m *NodeEvent) GetNodeName() string
```

##### NodeEvent.GetType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L278)  

```go
func (m *NodeEvent) GetType() NodeEvent_Type
```

##### NodeEvent.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L839)  

```go
func (this *NodeEvent) GoString() string
```

##### NodeEvent.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L987)  

```go
func (m *NodeEvent) Marshal() (dAtA []byte, err error)
```

##### NodeEvent.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L997)  

```go
func (m *NodeEvent) MarshalTo(dAtA []byte) (int, error)
```

##### NodeEvent.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1002)  

```go
func (m *NodeEvent) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### NodeEvent.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L247)  

```go
func (*NodeEvent) ProtoMessage()
```

##### NodeEvent.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L246)  

```go
func (m *NodeEvent) Reset()
```

##### NodeEvent.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1403)  

```go
func (m *NodeEvent) Size() (n int)
```

##### NodeEvent.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1602)  

```go
func (this *NodeEvent) String() string
```

##### NodeEvent.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1824)  

```go
func (m *NodeEvent) Unmarshal(dAtA []byte) error
```

##### NodeEvent.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L272)  

```go
func (m *NodeEvent) XXX_DiscardUnknown()
```

##### NodeEvent.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L254)  

```go
func (m *NodeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### NodeEvent.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L266)  

```go
func (m *NodeEvent) XXX_Merge(src proto.Message)
```

##### NodeEvent.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L269)  

```go
func (m *NodeEvent) XXX_Size() int
```

##### NodeEvent.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L251)  

```go
func (m *NodeEvent) XXX_Unmarshal(b []byte) error
```

---

### NodeEvent_Type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L86)  

```go
type NodeEvent_Type int32
```

#### Methods

##### NodeEvent_Type.EnumDescriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L112)  

```go
func (NodeEvent_Type) EnumDescriptor() ([]byte, []int)
```

##### NodeEvent_Type.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L108)  

```go
func (x NodeEvent_Type) String() string
```

---

### PeerClusterInfo

PeerClusterInfo represents the peer (gossip cluster) nodes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L118)  

```go
type PeerClusterInfo struct {
	PeerInfo
}
```

---

### PeerInfo

PeerInfo represents the peer (gossip cluster) nodes of a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L112)  

```go
type PeerInfo struct {
	Name string
	IP   string
}
```

---

### TableElem

TableElem elem

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.go#L424)  

```go
type TableElem struct {
	Value []byte
	// contains filtered or unexported fields
}
```

---

### TableEvent

TableEvent message payload definition.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L480)  

```go
type TableEvent struct {
	Type TableEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=networkdb.TableEvent_Type" json:"type,omitempty"`
	// Lamport time when this event was generated.
	LTime github_com_hashicorp_serf_serf.LamportTime `protobuf:"varint,2,opt,name=l_time,json=lTime,proto3,customtype=github.com/hashicorp/serf/serf.LamportTime" json:"l_time"`
	// Node name where this event originated.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// ID of the network to which this table entry belongs.
	NetworkID string `protobuf:"bytes,4,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Name of the table to which this table entry belongs.
	TableName string `protobuf:"bytes,5,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// Entry key.
	Key string `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	// Entry value.
	Value []byte `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	// Residual reap time for the entry before getting deleted in seconds
	ResidualReapTime int32 `protobuf:"varint,8,opt,name=residual_reap_time,json=residualReapTime,proto3" json:"residual_reap_time,omitempty"`
}
```

#### Methods

##### TableEvent.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L500)  

```go
func (*TableEvent) Descriptor() ([]byte, []int)
```

##### TableEvent.GetKey

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L558)  

```go
func (m *TableEvent) GetKey() string
```

##### TableEvent.GetNetworkID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L544)  

```go
func (m *TableEvent) GetNetworkID() string
```

##### TableEvent.GetNodeName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L537)  

```go
func (m *TableEvent) GetNodeName() string
```

##### TableEvent.GetResidualReapTime

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L572)  

```go
func (m *TableEvent) GetResidualReapTime() int32
```

##### TableEvent.GetTableName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L551)  

```go
func (m *TableEvent) GetTableName() string
```

##### TableEvent.GetType

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L530)  

```go
func (m *TableEvent) GetType() TableEvent_Type
```

##### TableEvent.GetValue

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L565)  

```go
func (m *TableEvent) GetValue() []byte
```

##### TableEvent.GoString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L891)  

```go
func (this *TableEvent) GoString() string
```

##### TableEvent.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1175)  

```go
func (m *TableEvent) Marshal() (dAtA []byte, err error)
```

##### TableEvent.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1185)  

```go
func (m *TableEvent) MarshalTo(dAtA []byte) (int, error)
```

##### TableEvent.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1190)  

```go
func (m *TableEvent) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### TableEvent.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L499)  

```go
func (*TableEvent) ProtoMessage()
```

##### TableEvent.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L498)  

```go
func (m *TableEvent) Reset()
```

##### TableEvent.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1490)  

```go
func (m *TableEvent) Size() (n int)
```

##### TableEvent.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L1657)  

```go
func (this *TableEvent) String() string
```

##### TableEvent.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L2384)  

```go
func (m *TableEvent) Unmarshal(dAtA []byte) error
```

##### TableEvent.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L524)  

```go
func (m *TableEvent) XXX_DiscardUnknown()
```

##### TableEvent.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L506)  

```go
func (m *TableEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### TableEvent.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L518)  

```go
func (m *TableEvent) XXX_Merge(src proto.Message)
```

##### TableEvent.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L521)  

```go
func (m *TableEvent) XXX_Size() int
```

##### TableEvent.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L503)  

```go
func (m *TableEvent) XXX_Unmarshal(b []byte) error
```

---

### TableEvent_Type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L146)  

```go
type TableEvent_Type int32
```

#### Methods

##### TableEvent_Type.EnumDescriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L179)  

```go
func (TableEvent_Type) EnumDescriptor() ([]byte, []int)
```

##### TableEvent_Type.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/networkdb.pb.go#L175)  

```go
func (x TableEvent_Type) String() string
```

---

### UpdateEvent

UpdateEvent generates a table entry update event to the watchers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/networkdb/watch.go#L36)  

```go
type UpdateEvent event
```

---


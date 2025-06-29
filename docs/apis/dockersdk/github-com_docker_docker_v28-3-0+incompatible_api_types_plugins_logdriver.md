# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/plugins/logdriver

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:09 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L751)

```go
var (
	ErrInvalidLengthEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEntry = fmt.Errorf("proto: unexpected end of group")
)
```

## Functions

This section is empty.

## Types

### LogEntry

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L25)  

```go
type LogEntry struct {
	Source             string                   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	TimeNano           int64                    `protobuf:"varint,2,opt,name=time_nano,json=timeNano,proto3" json:"time_nano,omitempty"`
	Line               []byte                   `protobuf:"bytes,3,opt,name=line,proto3" json:"line,omitempty"`
	Partial            bool                     `protobuf:"varint,4,opt,name=partial,proto3" json:"partial,omitempty"`
	PartialLogMetadata *PartialLogEntryMetadata `protobuf:"bytes,5,opt,name=partial_log_metadata,json=partialLogMetadata,proto3" json:"partial_log_metadata,omitempty"`
}
```

#### Methods

##### LogEntry.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L36)  

```go
func (*LogEntry) Descriptor() ([]byte, []int)
```

##### LogEntry.GetLine

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L80)  

```go
func (m *LogEntry) GetLine() []byte
```

##### LogEntry.GetPartial

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L87)  

```go
func (m *LogEntry) GetPartial() bool
```

##### LogEntry.GetPartialLogMetadata

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L94)  

```go
func (m *LogEntry) GetPartialLogMetadata() *PartialLogEntryMetadata
```

##### LogEntry.GetSource

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L66)  

```go
func (m *LogEntry) GetSource() string
```

##### LogEntry.GetTimeNano

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L73)  

```go
func (m *LogEntry) GetTimeNano() int64
```

##### LogEntry.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L188)  

```go
func (m *LogEntry) Marshal() (dAtA []byte, err error)
```

##### LogEntry.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L198)  

```go
func (m *LogEntry) MarshalTo(dAtA []byte) (int, error)
```

##### LogEntry.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L203)  

```go
func (m *LogEntry) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### LogEntry.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L35)  

```go
func (*LogEntry) ProtoMessage()
```

##### LogEntry.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L33)  

```go
func (m *LogEntry) Reset()
```

##### LogEntry.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L308)  

```go
func (m *LogEntry) Size() (n int)
```

##### LogEntry.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L34)  

```go
func (m *LogEntry) String() string
```

##### LogEntry.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L360)  

```go
func (m *LogEntry) Unmarshal(dAtA []byte) error
```

##### LogEntry.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L60)  

```go
func (m *LogEntry) XXX_DiscardUnknown()
```

##### LogEntry.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L42)  

```go
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### LogEntry.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L54)  

```go
func (m *LogEntry) XXX_Merge(src proto.Message)
```

##### LogEntry.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L57)  

```go
func (m *LogEntry) XXX_Size() int
```

##### LogEntry.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L39)  

```go
func (m *LogEntry) XXX_Unmarshal(b []byte) error
```

---

### LogEntryDecoder

LogEntryDecoder decodes log entries from a stream
It is expected that the wire format is as defined by LogEntryEncoder.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/io.go#L53)  

```go
type LogEntryDecoder interface {
	Decode(*LogEntry) error
}
```

#### Functions

##### NewLogEntryDecoder

NewLogEntryDecoder creates a new stream decoder for log entries

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/io.go#L58)  

```go
func NewLogEntryDecoder(r io.Reader) LogEntryDecoder
```

---

### LogEntryEncoder

LogEntryEncoder encodes a LogEntry to a protobuf stream
The stream should look like:

[uint32 binary encoded message size][protobuf message]

To decode an entry, read the first 4 bytes to get the size of the entry,
then read `size` bytes from the stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/io.go#L17)  

```go
type LogEntryEncoder interface {
	Encode(*LogEntry) error
}
```

#### Functions

##### NewLogEntryEncoder

NewLogEntryEncoder creates a protobuf stream encoder for log entries.
This is used to write out  log entries to a stream.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/io.go#L23)  

```go
func NewLogEntryEncoder(w io.Writer) LogEntryEncoder
```

---

### PartialLogEntryMetadata

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L101)  

```go
type PartialLogEntryMetadata struct {
	Last    bool   `protobuf:"varint,1,opt,name=last,proto3" json:"last,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Ordinal int32  `protobuf:"varint,3,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
}
```

#### Methods

##### PartialLogEntryMetadata.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L110)  

```go
func (*PartialLogEntryMetadata) Descriptor() ([]byte, []int)
```

##### PartialLogEntryMetadata.GetId

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L147)  

```go
func (m *PartialLogEntryMetadata) GetId() string
```

##### PartialLogEntryMetadata.GetLast

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L140)  

```go
func (m *PartialLogEntryMetadata) GetLast() bool
```

##### PartialLogEntryMetadata.GetOrdinal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L154)  

```go
func (m *PartialLogEntryMetadata) GetOrdinal() int32
```

##### PartialLogEntryMetadata.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L252)  

```go
func (m *PartialLogEntryMetadata) Marshal() (dAtA []byte, err error)
```

##### PartialLogEntryMetadata.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L262)  

```go
func (m *PartialLogEntryMetadata) MarshalTo(dAtA []byte) (int, error)
```

##### PartialLogEntryMetadata.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L267)  

```go
func (m *PartialLogEntryMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### PartialLogEntryMetadata.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L109)  

```go
func (*PartialLogEntryMetadata) ProtoMessage()
```

##### PartialLogEntryMetadata.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L107)  

```go
func (m *PartialLogEntryMetadata) Reset()
```

##### PartialLogEntryMetadata.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L335)  

```go
func (m *PartialLogEntryMetadata) Size() (n int)
```

##### PartialLogEntryMetadata.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L108)  

```go
func (m *PartialLogEntryMetadata) String() string
```

##### PartialLogEntryMetadata.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L551)  

```go
func (m *PartialLogEntryMetadata) Unmarshal(dAtA []byte) error
```

##### PartialLogEntryMetadata.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L134)  

```go
func (m *PartialLogEntryMetadata) XXX_DiscardUnknown()
```

##### PartialLogEntryMetadata.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L116)  

```go
func (m *PartialLogEntryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### PartialLogEntryMetadata.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L128)  

```go
func (m *PartialLogEntryMetadata) XXX_Merge(src proto.Message)
```

##### PartialLogEntryMetadata.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L131)  

```go
func (m *PartialLogEntryMetadata) XXX_Size() int
```

##### PartialLogEntryMetadata.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/plugins/logdriver/entry.pb.go#L113)  

```go
func (m *PartialLogEntryMetadata) XXX_Unmarshal(b []byte) error
```

---


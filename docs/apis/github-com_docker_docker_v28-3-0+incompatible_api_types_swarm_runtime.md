# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/swarm/runtime

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:20 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L804)

```go
var (
	ErrInvalidLengthPlugin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlugin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlugin = fmt.Errorf("proto: unexpected end of group")
)
```

## Functions

This section is empty.

## Types

### PluginPrivilege

PluginPrivilege describes a permission the user has to accept
upon installing a plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L105)  

```go
type PluginPrivilege struct {
	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Value       []string `protobuf:"bytes,3,rep,name=value,proto3" json:"value,omitempty"`
}
```

#### Methods

##### PluginPrivilege.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L114)  

```go
func (*PluginPrivilege) Descriptor() ([]byte, []int)
```

##### PluginPrivilege.GetDescription

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L151)  

```go
func (m *PluginPrivilege) GetDescription() string
```

##### PluginPrivilege.GetName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L144)  

```go
func (m *PluginPrivilege) GetName() string
```

##### PluginPrivilege.GetValue

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L158)  

```go
func (m *PluginPrivilege) GetValue() []string
```

##### PluginPrivilege.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L261)  

```go
func (m *PluginPrivilege) Marshal() (dAtA []byte, err error)
```

##### PluginPrivilege.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L271)  

```go
func (m *PluginPrivilege) MarshalTo(dAtA []byte) (int, error)
```

##### PluginPrivilege.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L276)  

```go
func (m *PluginPrivilege) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### PluginPrivilege.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L113)  

```go
func (*PluginPrivilege) ProtoMessage()
```

##### PluginPrivilege.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L111)  

```go
func (m *PluginPrivilege) Reset()
```

##### PluginPrivilege.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L350)  

```go
func (m *PluginPrivilege) Size() (n int)
```

##### PluginPrivilege.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L112)  

```go
func (m *PluginPrivilege) String() string
```

##### PluginPrivilege.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L579)  

```go
func (m *PluginPrivilege) Unmarshal(dAtA []byte) error
```

##### PluginPrivilege.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L138)  

```go
func (m *PluginPrivilege) XXX_DiscardUnknown()
```

##### PluginPrivilege.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L120)  

```go
func (m *PluginPrivilege) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### PluginPrivilege.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L132)  

```go
func (m *PluginPrivilege) XXX_Merge(src proto.Message)
```

##### PluginPrivilege.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L135)  

```go
func (m *PluginPrivilege) XXX_Size() int
```

##### PluginPrivilege.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L117)  

```go
func (m *PluginPrivilege) XXX_Unmarshal(b []byte) error
```

---

### PluginSpec

PluginSpec defines the base payload which clients can specify for creating
a service with the plugin runtime.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L27)  

```go
type PluginSpec struct {
	Name       string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remote     string             `protobuf:"bytes,2,opt,name=remote,proto3" json:"remote,omitempty"`
	Privileges []*PluginPrivilege `protobuf:"bytes,3,rep,name=privileges,proto3" json:"privileges,omitempty"`
	Disabled   bool               `protobuf:"varint,4,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Env        []string           `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty"`
}
```

#### Methods

##### PluginSpec.Descriptor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L38)  

```go
func (*PluginSpec) Descriptor() ([]byte, []int)
```

##### PluginSpec.GetDisabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L89)  

```go
func (m *PluginSpec) GetDisabled() bool
```

##### PluginSpec.GetEnv

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L96)  

```go
func (m *PluginSpec) GetEnv() []string
```

##### PluginSpec.GetName

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L68)  

```go
func (m *PluginSpec) GetName() string
```

##### PluginSpec.GetPrivileges

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L82)  

```go
func (m *PluginSpec) GetPrivileges() []*PluginPrivilege
```

##### PluginSpec.GetRemote

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L75)  

```go
func (m *PluginSpec) GetRemote() string
```

##### PluginSpec.Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L191)  

```go
func (m *PluginSpec) Marshal() (dAtA []byte, err error)
```

##### PluginSpec.MarshalTo

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L201)  

```go
func (m *PluginSpec) MarshalTo(dAtA []byte) (int, error)
```

##### PluginSpec.MarshalToSizedBuffer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L206)  

```go
func (m *PluginSpec) MarshalToSizedBuffer(dAtA []byte) (int, error)
```

##### PluginSpec.ProtoMessage

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L37)  

```go
func (*PluginSpec) ProtoMessage()
```

##### PluginSpec.Reset

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L35)  

```go
func (m *PluginSpec) Reset()
```

##### PluginSpec.Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L318)  

```go
func (m *PluginSpec) Size() (n int)
```

##### PluginSpec.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L36)  

```go
func (m *PluginSpec) String() string
```

##### PluginSpec.Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L379)  

```go
func (m *PluginSpec) Unmarshal(dAtA []byte) error
```

##### PluginSpec.XXX_DiscardUnknown

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L62)  

```go
func (m *PluginSpec) XXX_DiscardUnknown()
```

##### PluginSpec.XXX_Marshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L44)  

```go
func (m *PluginSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

##### PluginSpec.XXX_Merge

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L56)  

```go
func (m *PluginSpec) XXX_Merge(src proto.Message)
```

##### PluginSpec.XXX_Size

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L59)  

```go
func (m *PluginSpec) XXX_Size() int
```

##### PluginSpec.XXX_Unmarshal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/swarm/runtime/plugin.pb.go#L41)  

```go
func (m *PluginSpec) XXX_Unmarshal(b []byte) error
```

---


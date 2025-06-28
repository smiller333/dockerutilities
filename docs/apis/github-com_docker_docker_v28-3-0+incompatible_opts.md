# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/opts

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:16 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/hosts.go#L14)

```go
const (
	// DefaultHTTPPort Default HTTP Port used if only the protocol is provided to -H flag e.g. dockerd -H tcp://
	// These are the IANA registered port numbers for use with Docker
	// see http://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml?search=docker
	DefaultHTTPPort = 2375 // Default HTTP Port
	// DefaultTLSHTTPPort Default HTTP Port used when TLS enabled
	DefaultTLSHTTPPort = 2376 // Default TLS encrypted HTTP Port
	// DefaultUnixSocket Path for the unix socket.
	// Docker daemon by default always listens on the default unix socket
	DefaultUnixSocket = "/var/run/docker.sock"
	// DefaultTCPHost constant defines the default host string used by docker on Windows
	DefaultTCPHost = "tcp://" + DefaultHTTPHost + ":2375"
	// DefaultTLSHost constant defines the default host string used by docker for TLS sockets
	DefaultTLSHost = "tcp://" + DefaultHTTPHost + ":2376"
	// DefaultNamedPipe defines the default named pipe used by docker on Windows
	DefaultNamedPipe = `//./pipe/docker_engine`
	// HostGatewayName is the string value that can be passed
	// to the IPAddr section in --add-host that is replaced by
	// the value of HostGatewayIP daemon config value
	HostGatewayName = "host-gateway"
)
```

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/hosts_unix.go#L5)

```go
const (
	// DefaultHTTPHost Default HTTP Host used if only port is provided to -H flag e.g. dockerd -H tcp://:8080
	DefaultHTTPHost = "localhost"

	// DefaultHost constant defines the default host string used by docker on other hosts than Windows
	DefaultHost = "unix://" + DefaultUnixSocket
)
```

## Variables

This section is empty.

## Functions

### ParseHost

ParseHost and set defaults for a Daemon host string.
defaultToTLS is preferred over defaultToUnixXDG.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/hosts.go#L53)  
**Added in:** v1.9.0

```go
func ParseHost(defaultToTLS, defaultToUnixXDG bool, val string) (string, error)
```

---

### ParseLink

ParseLink parses and validates the specified string as a link format (name:alias)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L366)  

```go
func ParseLink(val string) (string, string, error)
```

---

### ParseTCPAddr

ParseTCPAddr parses and validates that the specified address is a valid TCP
address. It returns a formatted TCP address, either using the address parsed
from tryAddr, or the contents of defaultAddr if tryAddr is a blank string.
tryAddr is expected to have already been Trim()'d
defaultAddr must be in the full `tcp://host:port` form

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/hosts.go#L127)  
**Added in:** v1.12.0

```go
func ParseTCPAddr(tryAddr string, defaultAddr string) (string, error)
```

---

### ValidateDNSSearch

ValidateDNSSearch validates domain for resolvconf search configuration.
A zero length domain is represented by a dot (.).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L317)  
**Added in:** v1.8.0

```go
func ValidateDNSSearch(val string) (string, error)
```

---

### ValidateEnv

ValidateEnv validates an environment variable and returns it.
If no value is specified, it obtains its value from the current environment

As on ParseEnvFile and related to #16585, environment variable names
are not validate whatsoever, it's up to application inside docker
to validate them or not.

The only validation here is to check if name is empty, per #25099

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/env.go#L18)  

```go
func ValidateEnv(val string) (string, error)
```

---

### ValidateExtraHost

ValidateExtraHost validates that the specified string is a valid extrahost and returns it.
ExtraHost is in the form of name:ip where the ip has to be a valid ip (IPv4 or IPv6).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/hosts.go#L180)  
**Added in:** v1.3.0

```go
func ValidateExtraHost(val string) (string, error)
```

---

### ValidateHost

ValidateHost validates that the specified string is a valid host and returns it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/hosts.go#L37)  
**Added in:** v1.7.0

```go
func ValidateHost(val string) (string, error)
```

---

### ValidateIPAddress

ValidateIPAddress validates if the given value is a correctly formatted
IP address, and returns the value in normalized form. Leading and trailing
whitespace is allowed, but it does not allow IPv6 addresses surrounded by
square brackets ("[::1]").

Refer to net.ParseIP for accepted formats.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L307)  
**Added in:** v1.2.0

```go
func ValidateIPAddress(val string) (string, error)
```

---

### ValidateLabel

ValidateLabel validates that the specified string is a valid label,
it does not use the reserved namespaces com.docker.*, io.docker.*, org.dockerproject.*
and returns it.
Labels are in the form on key=value.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L339)  
**Added in:** v1.4.0

```go
func ValidateLabel(val string) (string, error)
```

---

### ValidateSingleGenericResource

ValidateSingleGenericResource validates that a single entry in the
generic resource list is valid.
i.e 'GPU=UID1' is valid however 'GPU:UID1' or 'UID1' isn't

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L358)  

```go
func ValidateSingleGenericResource(val string) (string, error)
```

---

## Types

### ListOpts

ListOpts holds a list of values and a validation function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L19)  

```go
type ListOpts struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewListOpts

NewListOpts creates a new ListOpts with the specified validator.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L25)  

```go
func NewListOpts(validator ValidatorFctType) ListOpts
```

##### NewListOptsRef

NewListOptsRef creates a new ListOpts with the specified values and validator.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L31)  
**Added in:** v1.8.0

```go
func NewListOptsRef(values *[]string, validator ValidatorFctType) *ListOpts
```

#### Methods

##### ListOpts.Delete

Delete removes the specified element from the slice.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L60)  

```go
func (opts *ListOpts) Delete(key string)
```

##### ListOpts.Get

Get checks the existence of the specified key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L95)  

```go
func (opts *ListOpts) Get(key string) bool
```

##### ListOpts.GetAll

GetAll returns the values of slice.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L80)  

```go
func (opts *ListOpts) GetAll() []string
```

##### ListOpts.GetAllOrEmpty

GetAllOrEmpty returns the values of the slice
or an empty slice when there are no values.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L86)  
**Added in:** v1.9.1

```go
func (opts *ListOpts) GetAllOrEmpty() []string
```

##### ListOpts.GetMap

GetMap returns the content of values in a map in order to avoid
duplicates.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L71)  

```go
func (opts *ListOpts) GetMap() map[string]struct{}
```

##### ListOpts.Len

Len returns the amount of element in the slice.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L105)  

```go
func (opts *ListOpts) Len() int
```

##### ListOpts.Set

Set validates if needed the input value and adds it to the
internal slice.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L47)  

```go
func (opts *ListOpts) Set(value string) error
```

##### ListOpts.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L38)  

```go
func (opts *ListOpts) String() string
```

##### ListOpts.Type

Type returns a string name for this Option type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L110)  
**Added in:** v1.12.0

```go
func (opts *ListOpts) Type() string
```

##### ListOpts.WithValidator

WithValidator returns the ListOpts with validator set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L115)  

```go
func (opts *ListOpts) WithValidator(validator ValidatorFctType) *ListOpts
```

---

### MapMapOpts

MapMapOpts holds a map of maps of values and a validation function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L171)  

```go
type MapMapOpts struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewMapMapOpts

NewMapMapOpts creates a new MapMapOpts with the specified map of values and a validator.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L216)  

```go
func NewMapMapOpts(values map[string]map[string]string, validator ValidatorFctType) *MapMapOpts
```

#### Methods

##### MapMapOpts.GetAll

GetAll returns the values of MapOpts as a map.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L202)  

```go
func (opts *MapMapOpts) GetAll() map[string]map[string]string
```

##### MapMapOpts.Set

Set validates if needed the input value and add it to the
internal map, by splitting on '='.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L178)  

```go
func (opts *MapMapOpts) Set(value string) error
```

##### MapMapOpts.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L206)  

```go
func (opts *MapMapOpts) String() string
```

##### MapMapOpts.Type

Type returns a string name for this Option type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L211)  

```go
func (opts *MapMapOpts) Type() string
```

---

### MapOpts

MapOpts holds a map of values and a validation function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L227)  
**Added in:** v1.7.0

```go
type MapOpts struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewMapOpts

NewMapOpts creates a new MapOpts with the specified map of values and a validator.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L262)  
**Added in:** v1.8.0

```go
func NewMapOpts(values map[string]string, validator ValidatorFctType) *MapOpts
```

#### Methods

##### MapOpts.GetAll

GetAll returns the values of MapOpts as a map.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L248)  
**Added in:** v1.9.0

```go
func (opts *MapOpts) GetAll() map[string]string
```

##### MapOpts.Set

Set validates if needed the input value and add it to the
internal map, by splitting on '='.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L234)  
**Added in:** v1.7.0

```go
func (opts *MapOpts) Set(value string) error
```

##### MapOpts.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L252)  
**Added in:** v1.7.0

```go
func (opts *MapOpts) String() string
```

##### MapOpts.Type

Type returns a string name for this Option type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L257)  
**Added in:** v1.12.0

```go
func (opts *MapOpts) Type() string
```

---

### MemBytes

MemBytes is a type for human readable memory bytes (like 128M, 2g, etc)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L388)  

```go
type MemBytes int64
```

#### Methods

##### MemBytes.Set

Set sets the value of the MemBytes by passing a string

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L402)  

```go
func (m *MemBytes) Set(value string) error
```

##### MemBytes.String

String returns the string format of the human readable memory bytes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L391)  

```go
func (m *MemBytes) String() string
```

##### MemBytes.Type

Type returns the type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L409)  

```go
func (m *MemBytes) Type() string
```

##### MemBytes.UnmarshalJSON

UnmarshalJSON is the customized unmarshaler for MemBytes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L419)  

```go
func (m *MemBytes) UnmarshalJSON(s []byte) error
```

##### MemBytes.Value

Value returns the value in int64

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L414)  

```go
func (m *MemBytes) Value() int64
```

---

### NamedListOpts

NamedListOpts is a ListOpts with a configuration name.
This struct is useful to keep reference to the assigned
field name in the internal configuration struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L129)  
**Added in:** v1.10.0

```go
type NamedListOpts struct {
	ListOpts
	// contains filtered or unexported fields
}
```

#### Functions

##### NewNamedListOptsRef

NewNamedListOptsRef creates a reference to a new NamedListOpts struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L137)  
**Added in:** v1.10.0

```go
func NewNamedListOptsRef(name string, values *[]string, validator ValidatorFctType) *NamedListOpts
```

#### Methods

##### NamedListOpts.Name

Name returns the name of the NamedListOpts in the configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L145)  
**Added in:** v1.10.0

```go
func (o *NamedListOpts) Name() string
```

---

### NamedMapMapOpts

NamedMapMapOpts is a MapMapOpts with a configuration name.
This struct is useful to keep reference to the assigned
field name in the internal configuration struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L152)  

```go
type NamedMapMapOpts struct {
	MapMapOpts
	// contains filtered or unexported fields
}
```

#### Functions

##### NewNamedMapMapOpts

NewNamedMapMapOpts creates a reference to a new NamedMapOpts struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L158)  

```go
func NewNamedMapMapOpts(name string, values map[string]map[string]string, validator ValidatorFctType) *NamedMapMapOpts
```

#### Methods

##### NamedMapMapOpts.Name

Name returns the name of the NamedListOpts in the configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L166)  

```go
func (o *NamedMapMapOpts) Name() string
```

---

### NamedMapOpts

NamedMapOpts is a MapOpts struct with a configuration name.
This struct is useful to keep reference to the assigned
field name in the internal configuration struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L275)  
**Added in:** v1.10.0

```go
type NamedMapOpts struct {
	MapOpts
	// contains filtered or unexported fields
}
```

#### Functions

##### NewNamedMapOpts

NewNamedMapOpts creates a reference to a new NamedMapOpts struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L283)  
**Added in:** v1.10.0

```go
func NewNamedMapOpts(name string, values map[string]string, validator ValidatorFctType) *NamedMapOpts
```

#### Methods

##### NamedMapOpts.Name

Name returns the name of the NamedMapOpts in the configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L291)  
**Added in:** v1.10.0

```go
func (o *NamedMapOpts) Name() string
```

---

### NamedOption

NamedOption is an interface that list and map options
with names implement.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L122)  
**Added in:** v1.10.0

```go
type NamedOption interface {
	Name() string
}
```

---

### NamedUlimitOpt

NamedUlimitOpt defines a named map of Ulimits

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/ulimit.go#L63)  

```go
type NamedUlimitOpt struct {
	UlimitOpt
	// contains filtered or unexported fields
}
```

#### Functions

##### NewNamedUlimitOpt

NewNamedUlimitOpt creates a new NamedUlimitOpt

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/ulimit.go#L71)  

```go
func NewNamedUlimitOpt(name string, ref *map[string]*container.Ulimit) *NamedUlimitOpt
```

#### Methods

##### NamedUlimitOpt.Name

Name returns the option name

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/ulimit.go#L82)  

```go
func (o *NamedUlimitOpt) Name() string
```

---

### PoolsOpt

PoolsOpt is a Value type for parsing the default address pools definitions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/address_pools.go#L15)  

```go
type PoolsOpt struct {
	Values []*ipamutils.NetworkToSplit
}
```

#### Methods

##### PoolsOpt.Name

Name returns the flag name of this option

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/address_pools.go#L85)  

```go
func (p *PoolsOpt) Name() string
```

##### PoolsOpt.Set

Set predefined pools

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/address_pools.go#L25)  

```go
func (p *PoolsOpt) Set(value string) error
```

##### PoolsOpt.String

String returns a string repr of this option

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/address_pools.go#L70)  

```go
func (p *PoolsOpt) String() string
```

##### PoolsOpt.Type

Type returns the type of this option

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/address_pools.go#L65)  

```go
func (p *PoolsOpt) Type() string
```

##### PoolsOpt.UnmarshalJSON

UnmarshalJSON fills values structure  info from JSON input

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/address_pools.go#L20)  

```go
func (p *PoolsOpt) UnmarshalJSON(raw []byte) error
```

##### PoolsOpt.Value

Value returns the mounts

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/address_pools.go#L80)  

```go
func (p *PoolsOpt) Value() []*ipamutils.NetworkToSplit
```

---

### RuntimeOpt

RuntimeOpt defines a map of Runtimes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/runtime.go#L11)  

```go
type RuntimeOpt struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewNamedRuntimeOpt

NewNamedRuntimeOpt creates a new RuntimeOpt

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/runtime.go#L18)  

```go
func NewNamedRuntimeOpt(name string, ref *map[string]system.Runtime, stockRuntime string) *RuntimeOpt
```

#### Methods

##### RuntimeOpt.GetMap

GetMap returns a map of Runtimes (name: path)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/runtime.go#L70)  

```go
func (o *RuntimeOpt) GetMap() map[string]system.Runtime
```

##### RuntimeOpt.Name

Name returns the name of the NamedListOpts in the configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/runtime.go#L26)  

```go
func (o *RuntimeOpt) Name() string
```

##### RuntimeOpt.Set

Set validates and updates the list of Runtimes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/runtime.go#L31)  

```go
func (o *RuntimeOpt) Set(val string) error
```

##### RuntimeOpt.String

String returns Runtime values as a string.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/runtime.go#L60)  

```go
func (o *RuntimeOpt) String() string
```

##### RuntimeOpt.Type

Type returns the type of the option

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/runtime.go#L79)  

```go
func (o *RuntimeOpt) Type() string
```

---

### UlimitOpt

UlimitOpt defines a map of Ulimits

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/ulimit.go#L11)  
**Added in:** v1.6.0

```go
type UlimitOpt struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewUlimitOpt

NewUlimitOpt creates a new UlimitOpt

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/ulimit.go#L16)  
**Added in:** v1.6.0

```go
func NewUlimitOpt(ref *map[string]*container.Ulimit) *UlimitOpt
```

#### Methods

##### UlimitOpt.GetList

GetList returns a slice of pointers to Ulimits.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/ulimit.go#L48)  
**Added in:** v1.6.0

```go
func (o *UlimitOpt) GetList() []*container.Ulimit
```

##### UlimitOpt.Set

Set validates a Ulimit and sets its name as a key in UlimitOpt

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/ulimit.go#L25)  
**Added in:** v1.6.0

```go
func (o *UlimitOpt) Set(val string) error
```

##### UlimitOpt.String

String returns Ulimit values as a string.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/ulimit.go#L38)  
**Added in:** v1.6.0

```go
func (o *UlimitOpt) String() string
```

##### UlimitOpt.Type

Type returns the option type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/ulimit.go#L58)  

```go
func (o *UlimitOpt) Type() string
```

---

### ValidatorFctListType

ValidatorFctListType defines a validator function that returns a validated list of string and/or an error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L299)  
**Added in:** v1.5.0

```go
type ValidatorFctListType func(val string) ([]string, error)
```

---

### ValidatorFctType

ValidatorFctType defines a validator function that returns a validated string and/or an error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/opts/opts.go#L296)  

```go
type ValidatorFctType func(val string) (string, error)
```

---


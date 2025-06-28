# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/opts

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:29 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ValidateHostGatewayIPs

ValidateHostGatewayIPs makes sure the addresses are valid, and there's at-most one IPv4 and one IPv6 address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/host_gateway_opts.go#L9)  

```go
func ValidateHostGatewayIPs(hostGatewayIPs []netip.Addr) error
```

---

## Types

### NamedIPListOpts

NamedIPListOpts appends to an underlying []netip.Addr.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/named_iplist_opts.go#L9)  

```go
type NamedIPListOpts struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewNamedIPListOptsRef

NewNamedIPListOptsRef constructs a NamedIPListOpts and returns its address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/named_iplist_opts.go#L15)  

```go
func NewNamedIPListOptsRef(name string, values *[]netip.Addr) *NamedIPListOpts
```

#### Methods

##### NamedIPListOpts.Name

Name returns the name of the NamedIPListOpts in the configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/named_iplist_opts.go#L46)  

```go
func (o *NamedIPListOpts) Name() string
```

##### NamedIPListOpts.Set

Set converts value to a netip.Addr and appends it to the underlying []netip.Addr.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/named_iplist_opts.go#L31)  

```go
func (o *NamedIPListOpts) Set(value string) error
```

##### NamedIPListOpts.String

String returns a string representation of the addresses in the underlying []netip.Addr.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/named_iplist_opts.go#L23)  

```go
func (o *NamedIPListOpts) String() string
```

##### NamedIPListOpts.Type

Type returns a string name for this Option type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/named_iplist_opts.go#L41)  

```go
func (o *NamedIPListOpts) Type() string
```

---

### NamedSetOpts

NamedSetOpts is a SetOpts struct with a configuration name.
This struct is useful to keep reference to the assigned
field name in the internal configuration struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/opts.go#L65)  

```go
type NamedSetOpts struct {
	SetOpts
	// contains filtered or unexported fields
}
```

#### Functions

##### NewNamedSetOpts

NewNamedSetOpts creates a reference to a new NamedSetOpts struct.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/opts.go#L73)  

```go
func NewNamedSetOpts(name string, values map[string]bool) *NamedSetOpts
```

#### Methods

##### NamedSetOpts.Name

Name returns the name of the NamedSetOpts in the configuration.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/opts.go#L81)  

```go
func (o *NamedSetOpts) Name() string
```

---

### SetOpts

SetOpts holds a map of values and a validation function.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/opts.go#L13)  

```go
type SetOpts struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewSetOpts

NewSetOpts creates a new SetOpts with the specified set of values as a map of string to bool.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/opts.go#L53)  

```go
func NewSetOpts(values map[string]bool) *SetOpts
```

#### Methods

##### SetOpts.GetAll

GetAll returns the values of SetOpts as a map.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/opts.go#L39)  

```go
func (opts *SetOpts) GetAll() map[string]bool
```

##### SetOpts.Set

Set validates if needed the input value and add it to the
internal map, by splitting on '='.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/opts.go#L19)  

```go
func (opts *SetOpts) Set(value string) error
```

##### SetOpts.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/opts.go#L43)  

```go
func (opts *SetOpts) String() string
```

##### SetOpts.Type

Type returns a string name for this Option type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/opts/opts.go#L48)  

```go
func (opts *SetOpts) Type() string
```

---


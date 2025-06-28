# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/ipams/defaultipam

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:23 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L18)

```go
const (
	// DriverName is the name of the built-in default IPAM driver.
	DriverName = "default"
)
```

## Variables

This section is empty.

## Functions

### Register

Register registers the default ipam driver with libnetwork. It takes
two optional address pools respectively containing the list of user-defined
address pools for 'local' and 'global' address spaces.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L29)  

```go
func Register(ic ipamapi.Registerer, lAddrPools, gAddrPools []*ipamutils.NetworkToSplit) error
```

---

## Types

### Allocator

Allocator provides per address space ipv4/ipv6 bookkeeping

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L45)  

```go
type Allocator struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewAllocator

NewAllocator returns an instance of libnetwork ipam

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L51)  

```go
func NewAllocator(lcAs, glAs []*ipamutils.NetworkToSplit) (*Allocator, error)
```

#### Methods

##### Allocator.GetDefaultAddressSpaces

GetDefaultAddressSpaces returns the local and global default address spaces

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L111)  

```go
func (a *Allocator) GetDefaultAddressSpaces() (string, string, error)
```

##### Allocator.IsBuiltIn

IsBuiltIn returns true for builtin drivers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L316)  

```go
func (a *Allocator) IsBuiltIn() bool
```

##### Allocator.ReleaseAddress

ReleaseAddress releases the address from the specified pool ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L264)  

```go
func (a *Allocator) ReleaseAddress(poolID string, address net.IP) error
```

##### Allocator.ReleasePool

ReleasePool releases the address pool identified by the passed id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L179)  

```go
func (a *Allocator) ReleasePool(poolID string) error
```

##### Allocator.RequestAddress

RequestAddress returns an address from the specified pool ID

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L234)  

```go
func (a *Allocator) RequestAddress(poolID string, prefAddress net.IP, opts map[string]string) (*net.IPNet, map[string]string, error)
```

##### Allocator.RequestPool

RequestPool returns an address pool along with its unique id.
addressSpace must be a valid address space name and must not be the empty string.
If requestedPool is the empty string then the default predefined pool for addressSpace will be used, otherwise pool must be a valid IP address and length in CIDR notation.
If requestedSubPool is not empty, it must be a valid IP address and length in CIDR notation which is a sub-range of requestedPool.
requestedSubPool must be empty if requestedPool is empty.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/allocator.go#L120)  

```go
func (a *Allocator) RequestPool(req ipamapi.PoolRequest) (ipamapi.AllocatedPool, error)
```

---

### PoolData

PoolData contains the configured pool data

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/structures.go#L19)  

```go
type PoolData struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### PoolData.String

String returns the string form of the PoolData object

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/structures.go#L72)  

```go
func (p *PoolData) String() string
```

---

### PoolID

PoolID is the pointer to the configured pools in each address space

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/structures.go#L13)  

```go
type PoolID struct {
	AddressSpace string
	SubnetKey
}
```

#### Functions

##### PoolIDFromString

PoolIDFromString creates a new PoolID and populates the SubnetKey object
reading it from the given string.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/structures.go#L38)  

```go
func PoolIDFromString(str string) (pID PoolID, err error)
```

#### Methods

##### PoolID.String

String returns the string form of the SubnetKey object

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/structures.go#L63)  

```go
func (s *PoolID) String() string
```

---

### SubnetKey

SubnetKey is the composite key to an address pool within an address space.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/structures.go#L28)  

```go
type SubnetKey struct {
	Subnet, ChildSubnet netip.Prefix
}
```

#### Methods

##### SubnetKey.Is6

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/defaultipam/structures.go#L32)  

```go
func (k SubnetKey) Is6() bool
```

---


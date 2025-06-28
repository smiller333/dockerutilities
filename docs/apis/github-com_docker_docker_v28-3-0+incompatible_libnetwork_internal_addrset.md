# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/internal/addrset

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:09:51 UTC

## Overview

Package addrset implements a set of IP addresses.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/addrset/addrset.go#L16)

```go
var (
	// ErrNotAvailable is returned when no more addresses are available to set
	ErrNotAvailable = errors.New("address not available")
	// ErrAllocated is returned when the specific address requested is already allocated
	ErrAllocated = errors.New("address already allocated")
)
```

## Functions

This section is empty.

## Types

### AddrSet

AddrSet is a set of IP addresses.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/addrset/addrset.go#L39)  

```go
type AddrSet struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New returns an AddrSet for the range of addresses in pool.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/addrset/addrset.go#L45)  

```go
func New(pool netip.Prefix) *AddrSet
```

#### Methods

##### AddrSet.Add

Add adds address addr to the set. If addr is already in the set, it returns a
wrapped ErrAllocated. If addr is not in the set's address range, it returns
an error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/addrset/addrset.go#L55)  

```go
func (as *AddrSet) Add(addr netip.Addr) error
```

##### AddrSet.AddAny

AddAny adds an arbitrary address to the set, and returns that address. Or, if
no addresses are available, it returns a wrapped ErrNotAvailable.

If the address set's pool contains fewer than 1<<maxBitsPerBitmap addresses,
AddAny will add any address from the entire set. If the pool is bigger than
that, AddAny will only consider the first 1<<maxBitsPerBitmap addresses. If
those are all allocated, it returns ErrNotAvailable.

When serial=true, the set is scanned starting from the address following
the address most recently set by AddrSet.AddAny (or AddrSet.AddAnyInRange
if the range is in the same 1<<maxBitsPerBitmap .

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/addrset/addrset.go#L81)  

```go
func (as *AddrSet) AddAny(serial bool) (netip.Addr, error)
```

##### AddrSet.AddAnyInRange

AddAnyInRange adds an arbitrary address from ipr to the set, and returns that
address. Or, if no addresses are available, it returns a wrapped ErrNotAvailable.
If ipr is not fully contained within the set's range, it returns an error.

When serial=true, the set is scanned starting from the address following
the address most recently set by AddrSet.AddAny or AddrSet.AddAnyInRange.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/addrset/addrset.go#L102)  

```go
func (as *AddrSet) AddAnyInRange(ipr netip.Prefix, serial bool) (netip.Addr, error)
```

##### AddrSet.Remove

Remove removes addr from the set or, if addr is not in the set's address range it
returns an error. If addr is not in the set, it returns nil (removing an address
that's not in the set is not an error).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/addrset/addrset.go#L127)  

```go
func (as *AddrSet) Remove(addr netip.Addr) error
```

##### AddrSet.String

String returns a description of the address set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/addrset/addrset.go#L146)  

```go
func (as *AddrSet) String() string
```

---


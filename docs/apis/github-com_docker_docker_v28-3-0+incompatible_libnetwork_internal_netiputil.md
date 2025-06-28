# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/internal/netiputil

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:09 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AddrPortFromNet

AddrPortFromNet converts a net.Addr into a netip.AddrPort.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/netiputil/netiputil.go#L56)  

```go
func AddrPortFromNet(addr net.Addr) netip.AddrPort
```

---

### HostID

HostID masks out the 'bits' most-significant bits of addr. The result is
undefined if bits > addr.BitLen().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/netiputil/netiputil.go#L43)  

```go
func HostID(addr netip.Addr, bits uint) uint64
```

---

### LastAddr

LastAddr returns the last address of prefix 'p'.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/netiputil/netiputil.go#L64)  

```go
func LastAddr(p netip.Prefix) netip.Addr
```

---

### PrefixAfter

PrefixAfter returns the prefix of size 'sz' right after 'prev'.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/netiputil/netiputil.go#L80)  

```go
func PrefixAfter(prev netip.Prefix, sz int) netip.Prefix
```

---

### PrefixCompare

PrefixCompare two prefixes and return a negative, 0, or a positive integer as
required by slices.SortFunc. When two prefixes with the same address is
provided, the shortest one will be sorted first.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/netiputil/netiputil.go#L71)  

```go
func PrefixCompare(a, b netip.Prefix) int
```

---

### SubnetRange

SubnetRange returns the amount to add to network.Addr() in order to yield the
first and last addresses in subnet, respectively.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/netiputil/netiputil.go#L49)  

```go
func SubnetRange(network, subnet netip.Prefix) (start, end uint64)
```

---

### ToIPNet

ToIPNet converts p into a *net.IPNet, returning nil if p is not valid.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/netiputil/netiputil.go#L11)  

```go
func ToIPNet(p netip.Prefix) *net.IPNet
```

---

### ToPrefix

ToPrefix converts n into a netip.Prefix. If n is not a valid IPv4 or IPV6
address, ToPrefix returns netip.Prefix{}, false.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/netiputil/netiputil.go#L23)  

```go
func ToPrefix(n *net.IPNet) (netip.Prefix, bool)
```

---

## Types

This section is empty.


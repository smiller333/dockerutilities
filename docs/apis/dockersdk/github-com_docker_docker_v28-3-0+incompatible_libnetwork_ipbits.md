# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/ipbits

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:10:32 UTC

## Overview

Package ipbits contains utilities for manipulating netip.Addr values as
numbers or bitfields.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Add

Add returns ip + (x << shift).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipbits/ipbits.go#L11)  

```go
func Add(ip netip.Addr, x uint64, shift uint) netip.Addr
```

---

### Field

Field returns the value of the bitfield [u, v] in ip as an integer,
where bit 0 is the most-significant bit of ip.

The result is undefined if u > v, if v-u > 64, or if u or v is larger than
ip.BitLen().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipbits/ipbits.go#L53)  

```go
func Field(ip netip.Addr, u, v uint) uint64
```

---

### SubnetsBetween

SubnetsBetween computes the number of subnets of size 'sz' available between 'a1'
and 'a2'. The result is capped at math.MaxUint64. It returns 0 when one of
'a1' or 'a2' is invalid, if both aren't of the same family, or when 'a2' is
less than 'a1'.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipbits/ipbits.go#L31)  

```go
func SubnetsBetween(a1 netip.Addr, a2 netip.Addr, sz int) uint64
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/ipamutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:10:30 UTC

## Overview

Package ipamutils provides utility functions for ipam management


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### NetworkToSplit

NetworkToSplit represent a network that has to be split in chunks with mask length Size.
Each subnet in the set is derived from the Base pool. Base is to be passed
in CIDR format.
Example: a Base "10.10.0.0/16 with Size 24 will define the set of 256
10.10.[0-255].0/24 address pools

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipamutils/utils.go#L32)  

```go
type NetworkToSplit struct {
	Base netip.Prefix `json:"base"`
	Size int          `json:"size"`
}
```

#### Functions

##### GetGlobalScopeDefaultNetworks

GetGlobalScopeDefaultNetworks returns a copy of the global-scope network list.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipamutils/utils.go#L48)  

```go
func GetGlobalScopeDefaultNetworks() []*NetworkToSplit
```

##### GetLocalScopeDefaultNetworks

GetLocalScopeDefaultNetworks returns a copy of the default local-scope network list.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipamutils/utils.go#L53)  

```go
func GetLocalScopeDefaultNetworks() []*NetworkToSplit
```

#### Methods

##### NetworkToSplit.FirstPrefix

FirstPrefix returns the first prefix available in NetworkToSplit.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipamutils/utils.go#L38)  

```go
func (n NetworkToSplit) FirstPrefix() netip.Prefix
```

##### NetworkToSplit.Overlaps

Overlaps is a util function checking whether 'p' overlaps with 'n'.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipamutils/utils.go#L43)  

```go
func (n NetworkToSplit) Overlaps(p netip.Prefix) bool
```

---


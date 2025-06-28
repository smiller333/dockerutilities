# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/netutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:10:39 UTC

## Overview

Package netutils provides network utility functions.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GenerateIfaceName

GenerateIfaceName returns an interface name using the passed in
prefix and the length of random bytes. The api ensures that the
there is no interface which exists with that name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/netutils/utils_linux.go#L103)  

```go
func GenerateIfaceName(nlh nlwrap.Handle, prefix string, len int) (string, error)
```

---

### GenerateMACFromIP

GenerateMACFromIP returns a locally administered MAC address where the 4 least
significant bytes are derived from the IPv4 address.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/netutils/utils.go#L19)  

```go
func GenerateMACFromIP(ip net.IP) net.HardwareAddr
```

---

### GenerateRandomMAC

GenerateRandomMAC returns a new 6-byte(48-bit) hardware address (MAC)
that is not multicast and has the local assignment bit set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/netutils/utils.go#L40)  

```go
func GenerateRandomMAC() net.HardwareAddr
```

---

### GenerateRandomName

GenerateRandomName returns a string of the specified length, created by joining the prefix to random hex characters.
The length must be strictly larger than len(prefix), or an error will be returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/netutils/utils.go#L50)  

```go
func GenerateRandomName(prefix string, length int) (string, error)
```

---

### InferReservedNetworks

InferReservedNetworks returns a list of network prefixes that seem to be
used by the system and that would likely break it if they were assigned to
some Docker networks. It uses two heuristics to build that list:

1. Nameservers configured in /etc/resolv.conf ;
2. On-link routes ;

That 2nd heuristic was originally not limited to on-links -- all non-default
routes were checked (see 1). This proved to be not ideal at best and
highly problematic at worst:

The 2nd heuristic was modified to be limited to on-link routes in PR #42598
(first released in v23.0, see 3).

If these heuristics don't detect an overlap, users should change their daemon
config to remove that overlapping prefix from `default-address-pools`. If a
prefix is found to overlap but users care enough about it being associated
to a Docker network they can still rely on static allocation.

For IPv6, the 2nd heuristic isn't applied as there's no such thing as
on-link routes for IPv6.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/netutils/utils_linux.go#L60)  

```go
func InferReservedNetworks(v6 bool) []netip.Prefix
```

---

### IsV6Listenable

IsV6Listenable returns true when `[::1]:0` is listenable.
IsV6Listenable returns false mostly when the kernel was booted with `ipv6.disable=1` option.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/netutils/utils.go#L108)  

```go
func IsV6Listenable() bool
```

---

### MustParseMAC

MustParseMAC returns a net.HardwareAddr or panic.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/netutils/utils.go#L125)  

```go
func MustParseMAC(s string) net.HardwareAddr
```

---

### ReverseIP

ReverseIP accepts a V4 or V6 IP string in the canonical form and returns a reversed IP in
the dotted decimal form . This is used to setup the IP to service name mapping in the optimal
way for the DNS PTR queries.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/netutils/utils.go#L68)  

```go
func ReverseIP(IP string) string
```

---

## Types

This section is empty.


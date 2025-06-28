# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/resolvconf

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:10:59 UTC

## Overview

Package resolvconf provides utility code to query and update DNS configuration in /etc/resolv.conf


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L15)

```go
const (
	IP = iota // IPv4 and IPv6
	IPv4
	IPv6
)
```

## Variables

This section is empty.

## Functions

### GetNameservers

GetNameservers returns nameservers (if any) listed in /etc/resolv.conf

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L67)  

```go
func GetNameservers(resolvConf []byte, kind int) []string
```

---

### GetNameserversAsPrefix

GetNameserversAsPrefix returns nameservers (if any) listed in
/etc/resolv.conf as CIDR blocks (e.g., "1.2.3.4/32")

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L88)  

```go
func GetNameserversAsPrefix(resolvConf []byte) []netip.Prefix
```

---

### GetOptions

GetOptions returns options (if any) listed in /etc/resolv.conf
If more than one options line is encountered, only the contents of the last
one is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L115)  

```go
func GetOptions(resolvConf []byte) []string
```

---

### GetSearchDomains

GetSearchDomains returns search domains (if any) listed in /etc/resolv.conf
If more than one search line is encountered, only the contents of the last
one is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L104)  

```go
func GetSearchDomains(resolvConf []byte) []string
```

---

### Path

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L27)  

```go
func Path() string
```

---

## Types

### File

File contains the resolv.conf content and its hash

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L22)  

```go
type File struct {
	Content []byte
	Hash    []byte
}
```

#### Functions

##### Build

Build generates and writes a configuration file to path containing a nameserver
entry for every element in nameservers, a "search" entry for every element in
dnsSearch, and an "options" entry for every element in dnsOptions. It returns
a File containing the generated content and its (sha256) hash.

Note that the resolv.conf file is written, but the hash file is not.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L129)  

```go
func Build(path string, nameservers, dnsSearch, dnsOptions []string) (*File, error)
```

##### FilterResolvDNS

FilterResolvDNS cleans up the config in resolvConf.  It has two main jobs:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L52)  

```go
func FilterResolvDNS(resolvConf []byte, ipv6Enabled bool) (*File, error)
```

##### Get

Get returns the contents of /etc/resolv.conf and its hash

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L32)  

```go
func Get() (*File, error)
```

##### GetSpecific

GetSpecific returns the contents of the user specified resolv.conf file and its hash

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/resolvconf/resolvconf.go#L37)  

```go
func GetSpecific(path string) (*File, error)
```

---


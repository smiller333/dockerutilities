# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/internal/resolvconf

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:14 UTC

## Overview

Package resolvconf is used to generate a container's /etc/resolv.conf file.

Constructor Load and Parse read a resolv.conf file from the filesystem or
a reader respectively, and return a ResolvConf object.

The ResolvConf object can then be updated with overrides for nameserver,
search domains, and DNS options.

ResolvConf can then be transformed to make it suitable for legacy networking,
a network with an internal nameserver, or used as-is for host networking.

This package includes methods to write the file for the container, along with
a hash that can be used to detect modifications made by the user to avoid
overwriting those updates.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Path

Path returns the path to the resolv.conf file that libnetwork should use.

When /etc/resolv.conf contains 127.0.0.53 as the only nameserver, then
it is assumed systemd-resolved manages DNS. Because inside the container 127.0.0.53
is not a valid DNS server, Path() returns /run/systemd/resolve/resolv.conf
which is the resolv.conf that systemd-resolved generates and manages.
Otherwise Path() returns /etc/resolv.conf.

Errors are silenced as they will inevitably resurface at future open/read calls.

More information at https://www.freedesktop.org/software/systemd/man/systemd-resolved.service.html#/etc/resolv.conf

TODO(robmry) - alternatePath is only needed for legacy networking ...

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf_path.go#L42)  

```go
func Path() string
```

---

### UserModified

UserModified can be used to determine whether the resolv.conf file has been
modified since it was generated. It returns false with no error if the file
matches the hash, true with no error if the file no longer matches the hash,
and false with an error if the result cannot be determined.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L394)  

```go
func UserModified(rcPath, rcHashPath string) (bool, error)
```

---

## Types

### ExtDNSEntry

ExtDNSEntry represents a nameserver address that was removed from the
container's resolv.conf when it was transformed by TransformForIntNS(). These
are addresses read from the host's file, or applied via an override ('--dns').

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L66)  

```go
type ExtDNSEntry struct {
	Addr         netip.Addr
	HostLoopback bool // The address is loopback, in the host's namespace.
}
```

#### Methods

##### ExtDNSEntry.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L71)  

```go
func (ed ExtDNSEntry) String() string
```

---

### ResolvConf

ResolvConf represents a resolv.conf file. It can be constructed by
reading a resolv.conf file, using method Parse().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L54)  

```go
type ResolvConf struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### Load

Load opens a file at path and parses it as a resolv.conf file.
On error, the returned ResolvConf will be zero-valued.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L96)  

```go
func Load(path string) (ResolvConf, error)
```

##### Parse

Parse parses a resolv.conf file from reader.
path is optional if reader is an *os.File.
On error, the returned ResolvConf will be zero-valued.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L108)  

```go
func Parse(reader io.Reader, path string) (ResolvConf, error)
```

#### Methods

##### ResolvConf.AddOption

AddOption adds a single DNS option.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L204)  

```go
func (rc *ResolvConf) AddOption(option string)
```

##### ResolvConf.Generate

Generate returns content suitable for writing to a resolv.conf file. If comments
is true, the file will include header information if supplied, and a trailing
comment that describes how the file was constructed and lists external resolvers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L292)  

```go
func (rc *ResolvConf) Generate(comments bool) ([]byte, error)
```

##### ResolvConf.NameServers

NameServers returns addresses used in nameserver directives.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L143)  

```go
func (rc *ResolvConf) NameServers() []netip.Addr
```

##### ResolvConf.Option

Option finds the last option named search, and returns (value, true) if
found, else ("", false). Options are treated as "name:value", where the
":value" may be omitted.

For example, for "ndots:1 edns0":

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L183)  

```go
func (rc *ResolvConf) Option(search string) (string, bool)
```

##### ResolvConf.Options

Options returns the current options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L171)  

```go
func (rc *ResolvConf) Options() []string
```

##### ResolvConf.OverrideNameServers

OverrideNameServers replaces the current set of nameservers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L148)  

```go
func (rc *ResolvConf) OverrideNameServers(nameServers []netip.Addr)
```

##### ResolvConf.OverrideOptions

OverrideOptions replaces the current DNS options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L194)  

```go
func (rc *ResolvConf) OverrideOptions(options []string)
```

##### ResolvConf.OverrideSearch

OverrideSearch replaces the current DNS search domains.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L159)  

```go
func (rc *ResolvConf) OverrideSearch(search []string)
```

##### ResolvConf.Search

Search returns the current DNS search domains.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L154)  

```go
func (rc *ResolvConf) Search() []string
```

##### ResolvConf.SetHeader

SetHeader sets the content to be included verbatim at the top of the
generated resolv.conf file. No formatting or checking is done on the
string. It must be valid resolv.conf syntax. (Comments must have '#'
or ';' in the first column of each line).

For example:

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L138)  

```go
func (rc *ResolvConf) SetHeader(c string)
```

##### ResolvConf.TransformForIntNS

TransformForIntNS makes sure the resolv.conf file will be suitable for
use in a network sandbox that has an internal DNS resolver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L246)  

```go
func (rc *ResolvConf) TransformForIntNS(
	internalNS netip.Addr,
	reqdOptions []string,
) ([]ExtDNSEntry, error)
```

##### ResolvConf.TransformForLegacyNw

TransformForLegacyNw makes sure the resolv.conf file will be suitable for
use in a legacy network (one that has no internal resolver).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L217)  

```go
func (rc *ResolvConf) TransformForLegacyNw(ipv6 bool)
```

##### ResolvConf.WriteFile

WriteFile generates content and writes it to path. If hashPath is non-zero, it
also writes a file containing a hash of the content, to enable UserModified()
to determine whether the file has been modified.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/resolvconf/resolvconf.go#L362)  

```go
func (rc *ResolvConf) WriteFile(path, hashPath string, perm os.FileMode) error
```

---


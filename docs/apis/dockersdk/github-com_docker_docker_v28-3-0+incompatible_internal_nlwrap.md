# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/nlwrap

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:27 UTC

## Overview

Package nlwrap wraps vishvandanda/netlink functions that may return EINTR.

A Handle instantiated using NewHandle or NewHandleAt can be used in place
of a netlink.Handle, it's a wrapper that replaces methods that need to be
wrapped. Functions that use the package handle need to be called as "nlwrap.X"
instead of "netlink.X".

When netlink.ErrDumpInterrupted is returned, the wrapped functions retry up to
maxAttempts times. This error means NLM_F_DUMP_INTR was flagged in a netlink
response, meaning something changed during the dump so results may be
incomplete or inconsistent.

To avoid retrying indefinitely, if netlink.ErrDumpInterrupted is still
returned after maxAttempts, the wrapped functions will discard the error, log
a stack trace to make the issue visible and aid in debugging, and return the
possibly inconsistent results. Returning possibly inconsistent results matches
the behaviour of vishvananda/netlink versions prior to 1.2.1, in which the
NLM_F_DUMP_INTR flag was ignored.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AddrList

AddrList calls netlink.AddrList, retrying if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L90)  

```go
func AddrList(link netlink.Link, family int) (addrs []netlink.Addr, err error)
```

---

### ConntrackTableList

ConntrackTableList calls netlink.ConntrackTableList, retrying if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L112)  

```go
func ConntrackTableList(
	table netlink.ConntrackTableType,
	family netlink.InetFamily,
) (flows []*netlink.ConntrackFlow, err error)
```

---

### LinkByName

LinkByName calls netlink.LinkByName, retrying if necessary. The netlink
function doesn't normally ask the kernel for a dump of links. But, on an old
kernel, it will do as a fallback and that dump may get inconsistent results.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L137)  

```go
func LinkByName(name string) (link netlink.Link, err error)
```

---

### LinkList

LinkList calls netlink.Handle.LinkList, retrying if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L155)  

```go
func LinkList() (links []netlink.Link, err error)
```

---

### LinkSubscribeWithOptions

LinkSubscribeWithOptions calls netlink.LinkSubscribeWithOptions, retrying if necessary.
Close the done channel when done (rather than just sending on it), so that goroutines
started by the netlink package are all stopped.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L166)  

```go
func LinkSubscribeWithOptions(ch chan<- netlink.LinkUpdate, done <-chan struct{}, options netlink.LinkSubscribeOptions) (err error)
```

---

## Types

### Handle

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L33)  

```go
type Handle struct {
	*netlink.Handle
}
```

#### Functions

##### NewHandle

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L37)  

```go
func NewHandle(nlFamilies ...int) (Handle, error)
```

##### NewHandleAt

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L45)  

```go
func NewHandleAt(ns netns.NsHandle, nlFamilies ...int) (Handle, error)
```

#### Methods

##### Handle.AddrList

AddrList calls nlh.Handle.AddrList, retrying if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L81)  

```go
func (nlh Handle) AddrList(link netlink.Link, family int) (addrs []netlink.Addr, err error)
```

##### Handle.Close

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L53)  

```go
func (nlh Handle) Close()
```

##### Handle.ConntrackDeleteFilters

ConntrackDeleteFilters calls nlh.Handle.ConntrackDeleteFilters, retrying if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L99)  

```go
func (nlh Handle) ConntrackDeleteFilters(
	table netlink.ConntrackTableType,
	family netlink.InetFamily,
	filters ...netlink.CustomConntrackFilter,
) (matched uint, err error)
```

##### Handle.LinkByName

LinkByName calls nlh.Handle.LinkByName, retrying if necessary. The netlink function
doesn't normally ask the kernel for a dump of links. But, on an old kernel, it
will do as a fallback and that dump may get inconsistent results.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L126)  

```go
func (nlh Handle) LinkByName(name string) (link netlink.Link, err error)
```

##### Handle.LinkList

LinkList calls nlh.Handle.LinkList, retrying if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L146)  

```go
func (nlh Handle) LinkList() (links []netlink.Link, err error)
```

##### Handle.RouteList

RouteList calls nlh.Handle.RouteList, retrying if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L175)  

```go
func (nlh Handle) RouteList(link netlink.Link, family int) (routes []netlink.Route, err error)
```

##### Handle.XfrmPolicyList

XfrmPolicyList calls nlh.Handle.XfrmPolicyList, retrying if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L184)  

```go
func (nlh Handle) XfrmPolicyList(family int) (policies []netlink.XfrmPolicy, err error)
```

##### Handle.XfrmStateList

XfrmStateList calls nlh.Handle.XfrmStateList, retrying if necessary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/nlwrap/nlwrap_linux.go#L193)  

```go
func (nlh Handle) XfrmStateList(family int) (states []netlink.XfrmState, err error)
```

---


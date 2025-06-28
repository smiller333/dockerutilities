# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/portallocator

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:00 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GetPortRange

GetPortRange returns the PortAllocator's default port range.

This function is for internal use in tests, and must not be used
for other purposes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/portallocator/portallocator.go#L62)  

```go
func GetPortRange() (start, end uint16)
```

---

## Types

### PortAllocator

PortAllocator manages the transport ports database

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/portallocator/portallocator.go#L38)  

```go
type PortAllocator struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### Get

Get returns the PortAllocator

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/portallocator/portallocator.go#L68)  

```go
func Get() *PortAllocator
```

#### Methods

##### PortAllocator.ReleaseAll

ReleaseAll releases all ports for all ips.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/portallocator/portallocator.go#L280)  

```go
func (p *PortAllocator) ReleaseAll()
```

##### PortAllocator.ReleasePort

ReleasePort releases port from global ports pool for specified ip and proto.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/portallocator/portallocator.go#L261)  

```go
func (p *PortAllocator) ReleasePort(ip net.IP, proto string, port int)
```

##### PortAllocator.RequestPort

RequestPort requests new port from global ports pool for specified ip and proto.
If port is 0 it returns first free port. Otherwise it checks port availability
in proto's pool and returns that port or error if port is already busy.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/portallocator/portallocator.go#L113)  

```go
func (p *PortAllocator) RequestPort(ip net.IP, proto string, port int) (int, error)
```

##### PortAllocator.RequestPortInRange

RequestPortInRange is equivalent to PortAllocator.RequestPortsInRange with
a single IP address. If ip is nil, a port is instead requested for the
default IP (0.0.0.0).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/portallocator/portallocator.go#L123)  

```go
func (p *PortAllocator) RequestPortInRange(ip net.IP, proto string, portStart, portEnd int) (int, error)
```

##### PortAllocator.RequestPortsInRange

RequestPortsInRange requests new ports from the global ports pool, for proto and each of ips.
If portStart and portEnd are 0 it returns the first free port in the default ephemeral range.
If portStart != portEnd it returns the first free port in the requested range.
Otherwise, (portStart == portEnd) it checks port availability in the requested proto's port-pool
and returns that port or error if port is already busy.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/portallocator/portallocator.go#L135)  

```go
func (p *PortAllocator) RequestPortsInRange(ips []net.IP, proto string, portStart, portEnd int) (int, error)
```

---


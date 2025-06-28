# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drvregistry

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:09:47 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### DriverWalkFunc

DriverWalkFunc defines the network driver table walker function signature.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/networks.go#L12)  

```go
type DriverWalkFunc func(name string, driver driverapi.Driver, capability driverapi.Capability) bool
```

---

### IPAMWalkFunc

IPAMWalkFunc defines the IPAM driver table walker function signature.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/ipams.go#L63)  

```go
type IPAMWalkFunc func(name string, driver ipamapi.Ipam, cap *ipamapi.Capability) bool
```

---

### IPAMs

IPAMs is a registry of IPAM drivers. The zero value is an empty IPAM driver
registry, ready to use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/ipams.go#L19)  

```go
type IPAMs struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### IPAMs.IPAM

IPAM returns the actual IPAM driver instance and its capability which registered with the passed name.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/ipams.go#L27)  

```go
func (ir *IPAMs) IPAM(name string) (ipamapi.Ipam, *ipamapi.Capability)
```

##### IPAMs.RegisterIpamDriver

RegisterIpamDriver registers the IPAM driver discovered with default capabilities.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/ipams.go#L58)  

```go
func (ir *IPAMs) RegisterIpamDriver(name string, driver ipamapi.Ipam) error
```

##### IPAMs.RegisterIpamDriverWithCapabilities

RegisterIpamDriverWithCapabilities registers the IPAM driver discovered with specified capabilities.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/ipams.go#L36)  

```go
func (ir *IPAMs) RegisterIpamDriverWithCapabilities(name string, driver ipamapi.Ipam, caps *ipamapi.Capability) error
```

##### IPAMs.WalkIPAMs

WalkIPAMs walks the IPAM drivers registered in the registry and invokes the passed walk function and each one of them.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/ipams.go#L66)  

```go
func (ir *IPAMs) WalkIPAMs(ifn IPAMWalkFunc)
```

---

### Networks

Networks is a registry of network drivers. The zero value is an empty network
driver registry, ready to use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/networks.go#L21)  

```go
type Networks struct {
	// Notify is called whenever a network driver is registered.
	Notify driverapi.Registerer
	// contains filtered or unexported fields
}
```

#### Methods

##### Networks.Driver

Driver returns the network driver instance registered under name, and its capability.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/networks.go#L53)  

```go
func (nr *Networks) Driver(name string) (driverapi.Driver, driverapi.Capability)
```

##### Networks.RegisterDriver

RegisterDriver registers the network driver with nr.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/networks.go#L62)  

```go
func (nr *Networks) RegisterDriver(ntype string, driver driverapi.Driver, capability driverapi.Capability) error
```

##### Networks.WalkDrivers

WalkDrivers walks the network drivers registered in the registry and invokes the passed walk function and each one of them.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drvregistry/networks.go#L32)  

```go
func (nr *Networks) WalkDrivers(dfn DriverWalkFunc)
```

---


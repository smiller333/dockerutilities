# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/cnmallocator

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:55 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/networkallocator.go#L24)

```go
const (
	// DefaultDriver defines the name of the driver to be used by
	// default if a network without any driver name specified is
	// created.
	DefaultDriver = "overlay"
)
```

## Variables

This section is empty.

## Functions

### IsBuiltInDriver

IsBuiltInDriver returns whether the passed driver is an internal network driver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/networkallocator.go#L967)  

```go
func IsBuiltInDriver(name string) bool
```

---

### RegisterManager

RegisterManager registers a new instance of the manager driver for networkType with r.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/manager.go#L16)  

```go
func RegisterManager(r driverapi.Registerer, networkType string) error
```

---

## Types

### Provider

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/provider.go#L17)  

```go
type Provider struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewProvider

NewProvider returns a new cnmallocator provider.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/provider.go#L24)  

```go
func NewProvider(pg plugingetter.PluginGetter) *Provider
```

#### Methods

##### Provider.NewAllocator

NewAllocator returns a new NetworkAllocator handle

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/networkallocator.go#L89)  

```go
func (p *Provider) NewAllocator(netConfig *networkallocator.Config) (networkallocator.NetworkAllocator, error)
```

##### Provider.PredefinedNetworks

PredefinedNetworks returns the list of predefined network structures

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/drivers_network_linux.go#L22)  

```go
func (*Provider) PredefinedNetworks() []networkallocator.PredefinedNetworkData
```

##### Provider.SetDefaultVXLANUDPPort

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/provider.go#L90)  

```go
func (p *Provider) SetDefaultVXLANUDPPort(port uint32) error
```

##### Provider.ValidateIPAMDriver

ValidateIPAMDriver implements networkallocator.NetworkProvider.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/provider.go#L29)  

```go
func (p *Provider) ValidateIPAMDriver(driver *api.Driver) error
```

##### Provider.ValidateIngressNetworkDriver

ValidateIngressNetworkDriver implements networkallocator.NetworkProvider.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/provider.go#L46)  

```go
func (p *Provider) ValidateIngressNetworkDriver(driver *api.Driver) error
```

##### Provider.ValidateNetworkDriver

ValidateNetworkDriver implements networkallocator.NetworkProvider.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cnmallocator/provider.go#L54)  

```go
func (p *Provider) ValidateNetworkDriver(driver *api.Driver) error
```

---


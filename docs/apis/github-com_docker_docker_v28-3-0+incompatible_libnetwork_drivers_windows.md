# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/windows

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:09:42 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/windows/labels.go#L3)

```go
const (
	// NetworkName label for bridge driver
	NetworkName = "com.docker.network.windowsshim.networkname"

	// HNSID of the discovered network
	HNSID = "com.docker.network.windowsshim.hnsid"

	// HNSOwned indicates that the network was learned from the host, not created by docker.
	HNSOwned = "com.docker.network.windowsshim.hnsowned"

	// RoutingDomain of the network
	RoutingDomain = "com.docker.network.windowsshim.routingdomain"

	// Interface of the network
	Interface = "com.docker.network.windowsshim.interface"

	// QosPolicies of the endpoint
	QosPolicies = "com.docker.endpoint.windowsshim.qospolicies"

	// VLAN of the network
	VLAN = "com.docker.network.windowsshim.vlanid"

	// VSID of the network
	VSID = "com.docker.network.windowsshim.vsid"

	// DNSSuffix of the network
	DNSSuffix = "com.docker.network.windowsshim.dnssuffix"

	// DNSServers of the network
	DNSServers = "com.docker.network.windowsshim.dnsservers"

	// MacPool of the network
	MacPool = "com.docker.network.windowsshim.macpool"

	// SourceMac of the network
	SourceMac = "com.docker.network.windowsshim.sourcemac"

	// DisableICC label
	DisableICC = "com.docker.network.windowsshim.disableicc"

	// DisableDNS label
	DisableDNS = "com.docker.network.windowsshim.disable_dns"

	// DisableGatewayDNS label
	DisableGatewayDNS = "com.docker.network.windowsshim.disable_gatewaydns"

	// EnableOutboundNat label
	EnableOutboundNat = "com.docker.network.windowsshim.enable_outboundnat"

	// OutboundNatExceptions label
	OutboundNatExceptions = "com.docker.network.windowsshim.outboundnat_exceptions"
)
```

## Variables

This section is empty.

## Functions

This section is empty.

## Types

This section is empty.


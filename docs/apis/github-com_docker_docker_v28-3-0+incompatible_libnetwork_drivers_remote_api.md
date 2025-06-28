# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/remote/api

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:09:40 UTC

## Overview

Package api represents all requests and responses suitable for conversation
with a remote driver.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### AllocateNetworkRequest

AllocateNetworkRequest requests allocation of new network by manager

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L36)  

```go
type AllocateNetworkRequest struct {
	// A network ID that remote plugins are expected to store for future
	// reference.
	NetworkID string

	// A free form map->object interface for communication of options.
	Options map[string]string

	// IPAMData contains the address pool information for this network
	IPv4Data, IPv6Data []driverapi.IPAMData
}
```

---

### AllocateNetworkResponse

AllocateNetworkResponse is the response to the AllocateNetworkRequest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L49)  

```go
type AllocateNetworkResponse struct {
	Response
	// A free form plugin specific string->string object to be sent in
	// CreateNetworkRequest call in the libnetwork agents
	Options map[string]string
}
```

---

### CreateEndpointRequest

CreateEndpointRequest is the request to create an endpoint within a network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L118)  

```go
type CreateEndpointRequest struct {
	// Provided at create time, this will be the network id referenced.
	NetworkID string
	// The ID of the endpoint for later reference.
	EndpointID string
	Interface  *EndpointInterface
	Options    map[string]interface{}
}
```

---

### CreateEndpointResponse

CreateEndpointResponse is the response to the CreateEndpoint action.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L135)  

```go
type CreateEndpointResponse struct {
	Response
	Interface *EndpointInterface
}
```

---

### CreateNetworkRequest

CreateNetworkRequest requests a new network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L89)  

```go
type CreateNetworkRequest struct {
	// A network ID that remote plugins are expected to store for future
	// reference.
	NetworkID string

	// A free form map->object interface for communication of options.
	Options map[string]interface{}

	// IPAMData contains the address pool information for this network
	IPv4Data, IPv6Data []driverapi.IPAMData
}
```

---

### CreateNetworkResponse

CreateNetworkResponse is the response to the CreateNetworkRequest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L102)  

```go
type CreateNetworkResponse struct {
	Response
}
```

---

### DeleteEndpointRequest

DeleteEndpointRequest describes the API for deleting an endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L148)  

```go
type DeleteEndpointRequest struct {
	NetworkID  string
	EndpointID string
}
```

---

### DeleteEndpointResponse

DeleteEndpointResponse is the response to the DeleteEndpoint action.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L154)  

```go
type DeleteEndpointResponse struct {
	Response
}
```

---

### DeleteNetworkRequest

DeleteNetworkRequest is the request to delete an existing network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L107)  

```go
type DeleteNetworkRequest struct {
	// The ID of the network to delete.
	NetworkID string
}
```

---

### DeleteNetworkResponse

DeleteNetworkResponse is the response to a request for deleting a network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L113)  

```go
type DeleteNetworkResponse struct {
	Response
}
```

---

### DiscoveryNotification

DiscoveryNotification represents a discovery notification

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L238)  

```go
type DiscoveryNotification struct {
	DiscoveryType discoverapi.DiscoveryType
	DiscoveryData interface{}
}
```

---

### DiscoveryResponse

DiscoveryResponse is used by libnetwork to log any plugin error processing the discovery notifications

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L244)  

```go
type DiscoveryResponse struct {
	Response
}
```

---

### EndpointInfoRequest

EndpointInfoRequest retrieves information about the endpoint from the network driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L159)  

```go
type EndpointInfoRequest struct {
	NetworkID  string
	EndpointID string
}
```

---

### EndpointInfoResponse

EndpointInfoResponse is the response to an EndpointInfoRequest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L165)  

```go
type EndpointInfoResponse struct {
	Response
	Value map[string]interface{}
}
```

---

### EndpointInterface

EndpointInterface represents an interface endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L128)  

```go
type EndpointInterface struct {
	Address     string
	AddressIPv6 string
	MacAddress  string
}
```

---

### FreeNetworkRequest

FreeNetworkRequest is the request to free allocated network in the manager

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L57)  

```go
type FreeNetworkRequest struct {
	// The ID of the network to be freed.
	NetworkID string
}
```

---

### FreeNetworkResponse

FreeNetworkResponse is the response to a request for freeing a network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L63)  

```go
type FreeNetworkResponse struct {
	Response
}
```

---

### GetCapabilityResponse

GetCapabilityResponse is the response of GetCapability request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L25)  

```go
type GetCapabilityResponse struct {
	Response
	Scope             string
	ConnectivityScope string

	// GwAllocChecker is used by the driver to report that it will accept a
	// [GwAllocCheckerRequest] at "GwAllocCheck".
	GwAllocChecker bool
}
```

---

### GwAllocCheckerRequest

GwAllocCheckerRequest is the body of a request sent to "GwAllocCheck", if the
driver reported capability "GwAllocChecker". This request is sent before the
CreateNetworkRequest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L70)  

```go
type GwAllocCheckerRequest struct {
	// Options has the same form as Options in [CreateNetworkRequest].
	Options map[string]interface{}
}
```

---

### GwAllocCheckerResponse

GwAllocCheckerResponse is the response to a GwAllocCheckerRequest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L76)  

```go
type GwAllocCheckerResponse struct {
	Response
	// SkipIPv4, if true, tells Docker that when it creates a network with the
	// Options in the [GwAllocCheckerRequest] it should not reserve an IPv4
	// gateway address.
	SkipIPv4 bool
	// SkipIPv6, if true, tells Docker that when it creates a network with the
	// Options in the [GwAllocCheckerRequest] it should not reserve an IPv6
	// gateway address.
	SkipIPv6 bool
}
```

---

### Interface

Interface is the representation of a linux interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L141)  

```go
type Interface struct {
	Address     *net.IPNet
	AddressIPv6 *net.IPNet
	MacAddress  net.HardwareAddr
}
```

---

### InterfaceName

InterfaceName is the struct representation of a pair of devices with source
and destination, for the purposes of putting an endpoint into a container.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L180)  

```go
type InterfaceName struct {
	SrcName   string
	DstName   string
	DstPrefix string
}
```

---

### JoinRequest

JoinRequest describes the API for joining an endpoint to a sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L171)  

```go
type JoinRequest struct {
	NetworkID  string
	EndpointID string
	SandboxKey string
	Options    map[string]interface{}
}
```

---

### JoinResponse

JoinResponse is the response to a JoinRequest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L194)  

```go
type JoinResponse struct {
	Response
	InterfaceName         *InterfaceName
	Gateway               string
	GatewayIPv6           string
	StaticRoutes          []StaticRoute
	DisableGatewayService bool
}
```

---

### LeaveRequest

LeaveRequest describes the API for detaching an endpoint from a sandbox.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L204)  

```go
type LeaveRequest struct {
	NetworkID  string
	EndpointID string
}
```

---

### LeaveResponse

LeaveResponse is the answer to LeaveRequest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L210)  

```go
type LeaveResponse struct {
	Response
}
```

---

### ProgramExternalConnectivityRequest

ProgramExternalConnectivityRequest describes the API for programming the external connectivity for the given endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L215)  

```go
type ProgramExternalConnectivityRequest struct {
	NetworkID  string
	EndpointID string
	Options    map[string]interface{}
}
```

---

### ProgramExternalConnectivityResponse

ProgramExternalConnectivityResponse is the answer to ProgramExternalConnectivityRequest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L222)  

```go
type ProgramExternalConnectivityResponse struct {
	Response
}
```

---

### Response

Response is the basic response structure used in all responses.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L15)  

```go
type Response struct {
	Err string
}
```

#### Methods

##### Response.GetError

GetError returns the error from the response, if any.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L20)  

```go
func (r *Response) GetError() string
```

---

### RevokeExternalConnectivityRequest

RevokeExternalConnectivityRequest describes the API for revoking the external connectivity for the given endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L227)  

```go
type RevokeExternalConnectivityRequest struct {
	NetworkID  string
	EndpointID string
}
```

---

### RevokeExternalConnectivityResponse

RevokeExternalConnectivityResponse is the answer to RevokeExternalConnectivityRequest.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L233)  

```go
type RevokeExternalConnectivityResponse struct {
	Response
}
```

---

### StaticRoute

StaticRoute is the plain JSON representation of a static route.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/remote/api/api.go#L187)  

```go
type StaticRoute struct {
	Destination string
	RouteType   int
	NextHop     string
}
```

---


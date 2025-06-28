# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/ipams/remote/api

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:31 UTC

## Overview

Package api defines the data structure to be used in the request/response
messages between libnetwork and the remote ipam plugin


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### GetAddressSpacesResponse

GetAddressSpacesResponse is the response to the “get default address spaces“ request message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L38)  

```go
type GetAddressSpacesResponse struct {
	Response
	LocalDefaultAddressSpace  string
	GlobalDefaultAddressSpace string
}
```

---

### GetCapabilityResponse

GetCapabilityResponse is the response of GetCapability request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L23)  

```go
type GetCapabilityResponse struct {
	Response
	RequiresMACAddress    bool
	RequiresRequestReplay bool
}
```

#### Methods

##### GetCapabilityResponse.ToCapability

ToCapability converts the capability response into the internal ipam driver capability structure

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L30)  

```go
func (capRes GetCapabilityResponse) ToCapability() *ipamapi.Capability
```

---

### ReleaseAddressRequest

ReleaseAddressRequest represents the expected data in a “release address“ request message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L86)  

```go
type ReleaseAddressRequest struct {
	PoolID  string
	Address string
}
```

---

### ReleaseAddressResponse

ReleaseAddressResponse represents the response message to a “release address“ request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L92)  

```go
type ReleaseAddressResponse struct {
	Response
}
```

---

### ReleasePoolRequest

ReleasePoolRequest represents the expected data in a “release address pool“ request message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L62)  

```go
type ReleasePoolRequest struct {
	PoolID string
}
```

---

### ReleasePoolResponse

ReleasePoolResponse represents the response message to a “release address pool“ request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L67)  

```go
type ReleasePoolResponse struct {
	Response
}
```

---

### RequestAddressRequest

RequestAddressRequest represents the expected data in a “request address“ request message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L72)  

```go
type RequestAddressRequest struct {
	PoolID  string
	Address string
	Options map[string]string
}
```

---

### RequestAddressResponse

RequestAddressResponse represents the expected data in the response message to a “request address“ request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L79)  

```go
type RequestAddressResponse struct {
	Response
	Address string // in CIDR format
	Data    map[string]string
}
```

---

### RequestPoolRequest

RequestPoolRequest represents the expected data in a “request address pool“ request message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L45)  

```go
type RequestPoolRequest struct {
	AddressSpace string
	Pool         string
	SubPool      string
	Options      map[string]string
	V6           bool
}
```

---

### RequestPoolResponse

RequestPoolResponse represents the response message to a “request address pool“ request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L54)  

```go
type RequestPoolResponse struct {
	Response
	PoolID string
	Pool   string // CIDR format
	Data   map[string]string
}
```

---

### Response

Response is the basic response structure used in all responses

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L8)  

```go
type Response struct {
	Error string
}
```

#### Methods

##### Response.GetError

GetError returns the error from the response, if any.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L18)  

```go
func (r *Response) GetError() string
```

##### Response.IsSuccess

IsSuccess returns whether the plugin response is successful

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ipams/remote/api/api.go#L13)  

```go
func (r *Response) IsSuccess() bool
```

---


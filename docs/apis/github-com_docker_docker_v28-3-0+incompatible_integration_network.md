# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration/network

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:37 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### CreateMasterDummy

CreateMasterDummy creates a dummy network interface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/network/helpers.go#L18)  

```go
func CreateMasterDummy(ctx context.Context, t *testing.T, master string)
```

---

### CreateVlanInterface

CreateVlanInterface creates a vlan network interface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/network/helpers.go#L25)  

```go
func CreateVlanInterface(ctx context.Context, t *testing.T, master, slave, id string)
```

---

### DeleteInterface

DeleteInterface deletes a network interface

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/network/helpers.go#L33)  

```go
func DeleteInterface(ctx context.Context, t *testing.T, ifName string)
```

---

### IsNetworkAvailable

IsNetworkAvailable provides a comparison to check if a docker network is available

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/network/helpers.go#L55)  

```go
func IsNetworkAvailable(ctx context.Context, c client.NetworkAPIClient, name string) is.Comparison
```

---

### IsNetworkNotAvailable

IsNetworkNotAvailable provides a comparison to check if a docker network is not available

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/network/helpers.go#L71)  

```go
func IsNetworkNotAvailable(ctx context.Context, c client.NetworkAPIClient, name string) is.Comparison
```

---

### LinkDoesntExist

LinkDoesntExist verifies that a link doesn't exist

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/network/helpers.go#L46)  

```go
func LinkDoesntExist(ctx context.Context, t *testing.T, master string)
```

---

### LinkExists

LinkExists verifies that a link exists

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/network/helpers.go#L40)  

```go
func LinkExists(ctx context.Context, t *testing.T, master string)
```

---

## Types

This section is empty.


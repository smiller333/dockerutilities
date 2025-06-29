# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/cmd/networkdb-test/dummyclient

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:08:46 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### RegisterDiagnosticHandlers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cmd/networkdb-test/dummyclient/dummyClient.go#L18)  

```go
func RegisterDiagnosticHandlers(mux Mux, nDB *networkdb.NetworkDB)
```

---

## Types

### Mux

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/cmd/networkdb-test/dummyclient/dummyClient.go#L14)  

```go
type Mux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}
```

---


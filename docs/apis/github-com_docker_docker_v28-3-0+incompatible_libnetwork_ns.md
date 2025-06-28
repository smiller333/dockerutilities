# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/ns

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:51 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ns/init_linux.go#L15)

```go
var (

	// NetlinkSocketsTimeout represents the default timeout duration for the sockets
	NetlinkSocketsTimeout = 3 * time.Second
)
```

## Functions

### Init

Init initializes a new network namespace

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ns/init_linux.go#L24)  

```go
func Init()
```

---

### NlHandle

NlHandle returns the netlink handler

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ns/init_linux.go#L52)  

```go
func NlHandle() nlwrap.Handle
```

---

### ParseHandlerInt

ParseHandlerInt transforms the namespace handler into an integer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/ns/init_linux.go#L41)  

```go
func ParseHandlerInt() int
```

---

## Types

This section is empty.


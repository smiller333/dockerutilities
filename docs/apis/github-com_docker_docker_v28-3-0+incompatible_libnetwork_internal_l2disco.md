# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/internal/l2disco

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:07 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### UnsolARP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/l2disco/unsol_arp_linux.go#L31)  

```go
type UnsolARP struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewUnsolARP

NewUnsolARP returns a pointer to an object that can send unsolicited ARPs on
the interface with ifIndex, for ip and mac.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/l2disco/unsol_arp_linux.go#L39)  

```go
func NewUnsolARP(_ context.Context, ip net.IP, mac net.HardwareAddr, ifIndex int) (*UnsolARP, error)
```

#### Methods

##### UnsolARP.Close

Close releases resources.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/l2disco/unsol_arp_linux.go#L71)  

```go
func (ua *UnsolARP) Close() error
```

##### UnsolARP.Send

Send sends an unsolicited ARP message.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/l2disco/unsol_arp_linux.go#L66)  

```go
func (ua *UnsolARP) Send() error
```

---

### UnsolNA

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/l2disco/unsol_na_linux.go#L28)  

```go
type UnsolNA struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewUnsolNA

NewUnsolNA returns a pointer to an object that can send unsolicited Neighbour
Advertisements for ip and mac.
https://datatracker.ietf.org/doc/html/rfc4861#section-4.4

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/l2disco/unsol_na_linux.go#L37)  

```go
func NewUnsolNA(ctx context.Context, ip net.IP, mac net.HardwareAddr, ifIndex int) (*UnsolNA, error)
```

#### Methods

##### UnsolNA.Close

Close releases resources.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/l2disco/unsol_na_linux.go#L87)  

```go
func (un *UnsolNA) Close() error
```

##### UnsolNA.Send

Send sends an unsolicited ARP message.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/l2disco/unsol_na_linux.go#L75)  

```go
func (un *UnsolNA) Send() error
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration/internal/network

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:23 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/dns.go#L11)

```go
const DNSRespAddr = "10.11.12.13"
```

## Variables

This section is empty.

## Functions

### CollectBcastARPs

CollectBcastARPs collects broadcast ARPs from interface ifname.
It returns a stop function, to stop collection and return a slice of collected packets (with
timestamps added when they were received in userspace).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/l2disco_linux.go#L31)  

```go
func CollectBcastARPs(t *testing.T, ifname string) (stop func() []TimestampedPkt)
```

---

### CollectICMP6

CollectICMP6 collects ICMP6 packets sent to the all nodes address.
It returns a stop function, to stop collection and return a slice of collected packets (with
timestamps added when they were received in userspace).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/l2disco_linux.go#L59)  

```go
func CollectICMP6(t *testing.T, ifname string) (stop func() []TimestampedPkt)
```

---

### Create

Create creates a network with the specified options

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/network.go#L24)  

```go
func Create(ctx context.Context, client client.APIClient, name string, ops ...func(*network.CreateOptions)) (string, error)
```

---

### CreateNoError

CreateNoError creates a network with the specified options and verifies there were no errors

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/network.go#L29)  

```go
func CreateNoError(ctx context.Context, t *testing.T, client client.APIClient, name string, ops ...func(*network.CreateOptions)) string
```

---

### GenResolvConf

GenResolvConf generates a resolv.conf that only contains a single
nameserver line, with address addr, and returns the file content.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/dns.go#L15)  

```go
func GenResolvConf(addr string) string
```

---

### IsRemoved

IsRemoved verifies the network is removed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/states.go#L12)  

```go
func IsRemoved(ctx context.Context, client client.NetworkAPIClient, networkID string) func(log poll.LogT) poll.Result
```

---

### RemoveNoError

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/network.go#L37)  

```go
func RemoveNoError(ctx context.Context, t *testing.T, apiClient client.APIClient, name string)
```

---

### StartDaftDNS

StartDaftDNS starts and returns a really, really daft DNS server that only
responds to type-A requests, and always with address dnsRespAddr.
The DNS server will be stopped automatically by a t.Cleanup().

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/dns.go#L22)  

```go
func StartDaftDNS(t *testing.T, addr string)
```

---

### UnpackUnsolARP

UnpackUnsolARP checks the packet is a valid Ethernet unsolicited/broadcast ARP
request packet. It returns sender hardware and protocol addresses,
and true if it is - else, an error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/l2disco_linux.go#L130)  

```go
func UnpackUnsolARP(pkt TimestampedPkt) (sh net.HardwareAddr, sp netip.Addr, err error)
```

---

### UnpackUnsolNA

UnpackUnsolNA returns the hardware (MAC) and protocol (IP) addresses from the
packet, if it is an unsolicited Neighbour Advertisement message with a
link address option. Otherwise, it returns an error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/l2disco_linux.go#L177)  

```go
func UnpackUnsolNA(pkt TimestampedPkt) (th net.HardwareAddr, tp netip.Addr, err error)
```

---

### WithAttachable

WithAttachable sets Attachable flag on the create network request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L68)  

```go
func WithAttachable() func(*network.CreateOptions)
```

---

### WithConfigFrom

WithConfigFrom sets the ConfigOnly flag in the create network request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L61)  

```go
func WithConfigFrom(name string) func(*network.CreateOptions)
```

---

### WithConfigOnly

WithConfigOnly sets the ConfigOnly flag in the create network request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L54)  

```go
func WithConfigOnly(co bool) func(*network.CreateOptions)
```

---

### WithDriver

WithDriver sets the driver of the network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L8)  

```go
func WithDriver(driver string) func(*network.CreateOptions)
```

---

### WithIPAM

WithIPAM adds an IPAM with the specified Subnet and Gateway to the network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L133)  

```go
func WithIPAM(subnet, gateway string) func(*network.CreateOptions)
```

---

### WithIPAMRange

WithIPAMRange adds an IPAM with the specified Subnet, IPRange and Gateway to the network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L138)  

```go
func WithIPAMRange(subnet, iprange, gateway string) func(*network.CreateOptions)
```

---

### WithIPv4

WithIPv4 enables/disables IPv4 on the network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L15)  

```go
func WithIPv4(enable bool) func(*network.CreateOptions)
```

---

### WithIPv4Disabled

WithIPv4Disabled makes sure IPv4 is disabled on the network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L31)  

```go
func WithIPv4Disabled() func(*network.CreateOptions)
```

---

### WithIPv6

WithIPv6 Enables IPv6 on the network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L23)  

```go
func WithIPv6() func(*network.CreateOptions)
```

---

### WithIPv6Disabled

WithIPv6Disabled makes sure IPv6 is disabled on the network.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L39)  

```go
func WithIPv6Disabled() func(*network.CreateOptions)
```

---

### WithIPvlan

WithIPvlan sets the network as ipvlan with the specified parent and mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L107)  

```go
func WithIPvlan(parent, mode string) func(*network.CreateOptions)
```

---

### WithInternal

WithInternal enables Internal flag on the create network request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L47)  

```go
func WithInternal() func(*network.CreateOptions)
```

---

### WithMacvlan

WithMacvlan sets the network as macvlan with the specified parent

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L82)  

```go
func WithMacvlan(parent string) func(*network.CreateOptions)
```

---

### WithMacvlanPassthru

WithMacvlanPassthru sets the network as macvlan with the specified parent in passthru mode

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L94)  

```go
func WithMacvlanPassthru(parent string) func(options *network.CreateOptions)
```

---

### WithOption

WithOption adds the specified key/value pair to network's options

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L123)  

```go
func WithOption(key, value string) func(*network.CreateOptions)
```

---

### WithScope

WithScope sets the network scope.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/ops.go#L75)  

```go
func WithScope(s string) func(*network.CreateOptions)
```

---

## Types

### TimestampedPkt

TimestampedPkt has a Data slice representing a packet, ReceivedAt is a timestamp
set after the packet was received in user-space.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/network/l2disco_linux.go#L23)  

```go
type TimestampedPkt struct {
	ReceivedAt time.Time
	Data       []byte
}
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/testutils/networking

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:01 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L17)

```go
const CurrentNetns = ""
```

## Variables

This section is empty.

## Functions

### FirewalldReload

FirewalldReload reloads firewalld and waits for the daemon to re-create its rules.
It's a no-op if firewalld is not running, and the test fails if the reload does
not complete.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/firewall.go#L110)  

```go
func FirewalldReload(t *testing.T, d *daemon.Daemon)
```

---

### FirewalldRunning

FirewalldRunning returns true if "firewall-cmd --state" reports "running".

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/firewall.go#L102)  

```go
func FirewalldRunning() bool
```

---

### SetFilterForwardPolicies

SetFilterForwardPolicies sets the default policy for the FORWARD chain in
the filter tables for both IPv4 and IPv6. The original policy is restored
using t.Cleanup().

There's only one filter-FORWARD policy, so this won't behave well if used by
tests running in parallel in a single network namespace that expect different
behaviour.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/firewall.go#L33)  

```go
func SetFilterForwardPolicies(t *testing.T, firewallBackend string, policy string)
```

---

## Types

### Host

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L94)  

```go
type Host struct {
	Name  string
	Iface string // Iface is the interface name in the host network namespace.
	// contains filtered or unexported fields
}
```

#### Methods

##### Host.Destroy

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L174)  

```go
func (h Host) Destroy(t *testing.T)
```

##### Host.Do

Do run the provided function in the host's network namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L146)  

```go
func (h Host) Do(t *testing.T, fn func())
```

##### Host.MustRun

MustRun executes the provided command in the host's network namespace
and returns its combined stdout/stderr, failing the test if the
command returns an error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L135)  

```go
func (h Host) MustRun(t *testing.T, cmd string, args ...string) string
```

##### Host.Run

Run executes the provided command in the host's network namespace,
returns its combined stdout/stderr, and error.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L123)  

```go
func (h Host) Run(t *testing.T, cmd string, args ...string) (string, error)
```

---

### L3Segment

L3Segment simulates a switched, dual-stack capable network that
interconnects multiple hosts running in their own network namespace.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L33)  

```go
type L3Segment struct {
	Hosts map[string]Host
	// contains filtered or unexported fields
}
```

#### Functions

##### NewL3Segment

NewL3Segment creates a new L3Segment. The bridge interface interconnecting
all the hosts is created in a new network namespace named nsName and it's
assigned one or more IP addresses. Those need to be unmasked netip.Prefix.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L41)  

```go
func NewL3Segment(t *testing.T, nsName string, addrs ...netip.Prefix) *L3Segment
```

#### Methods

##### L3Segment.AddHost

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L64)  

```go
func (l3 *L3Segment) AddHost(t *testing.T, hostname, nsName, ifname string, addrs ...netip.Prefix)
```

##### L3Segment.Destroy

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/testutils/networking/l3_segment_linux.go#L86)  

```go
func (l3 *L3Segment) Destroy(t *testing.T)
```

---


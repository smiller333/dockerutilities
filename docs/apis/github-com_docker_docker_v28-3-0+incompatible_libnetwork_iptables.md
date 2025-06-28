# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/iptables

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:10:35 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AddInterfaceFirewalld

AddInterfaceFirewalld adds the interface to the trusted zone. It is a
no-op if firewalld is not running.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/firewalld.go#L315)  

```go
func AddInterfaceFirewalld(intf string) error
```

---

### DelInterfaceFirewalld

DelInterfaceFirewalld removes the interface from the trusted zone It is a
no-op if firewalld is not running.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/firewalld.go#L341)  

```go
func DelInterfaceFirewalld(intf string) error
```

---

### DeleteConntrackEntries

DeleteConntrackEntries deletes all the conntrack connections on the host for the specified IP
Returns the number of flows deleted for IPv4, IPv6 else error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/conntrack.go#L28)  

```go
func DeleteConntrackEntries(nlh nlwrap.Handle, ipv4List []net.IP, ipv6List []net.IP) error
```

---

### DeleteConntrackEntriesByPort

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/conntrack.go#L60)  

```go
func DeleteConntrackEntriesByPort(nlh nlwrap.Handle, proto types.Protocol, ports []uint16) error
```

---

### FirewalldReloadedAt

FirewalldReloadedAt returns the time at which the daemon last completed a
firewalld reload, or a zero-valued time.Time if it has not been reloaded
since the daemon started.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/firewalld.go#L63)  

```go
func FirewalldReloadedAt() time.Time
```

---

### OnReloaded

OnReloaded add callback

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/firewalld.go#L179)  

```go
func OnReloaded(callback func())
```

---

### UsingFirewalld

UsingFirewalld returns true if iptables rules will be applied via firewalld's
passthrough interface. The error return is non-nil if the status cannot be
determined because the initialisation function has not been called.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/firewalld.go#L50)  

```go
func UsingFirewalld() (bool, error)
```

---

## Types

### Action

Action signifies the iptable action.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L23)  

```go
type Action string
```

---

### ChainError

ChainError is returned to represent errors during ip table operation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L87)  

```go
type ChainError struct {
	Chain  string
	Output []byte
}
```

#### Methods

##### ChainError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L92)  

```go
func (e ChainError) Error() string
```

---

### ChainInfo

ChainInfo defines the iptables chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L80)  

```go
type ChainInfo struct {
	Name      string
	Table     Table
	IPVersion IPVersion
}
```

#### Methods

##### ChainInfo.Link

Link adds reciprocal ACCEPT rule for two supplied IP addresses.
Traffic is allowed from ip1 to ip2 and vice-versa

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L207)  

```go
func (c *ChainInfo) Link(action Action, ip1, ip2 netip.Addr, port int, proto string, bridgeName string) error
```

##### ChainInfo.Output

Output adds linking rule to an OUTPUT chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L254)  

```go
func (c *ChainInfo) Output(action Action, args ...string) error
```

##### ChainInfo.Prerouting

Prerouting adds linking rule to nat/PREROUTING chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L239)  

```go
func (c *ChainInfo) Prerouting(action Action, args ...string) error
```

##### ChainInfo.Remove

Remove removes the chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L268)  

```go
func (c *ChainInfo) Remove() error
```

---

### Conn

Conn is a connection to firewalld dbus endpoint.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/firewalld.go#L28)  

```go
type Conn struct {
	// contains filtered or unexported fields
}
```

---

### IPTable

IPTable defines struct with IPVersion.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L75)  

```go
type IPTable struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### GetIptable

GetIptable returns an instance of IPTable with specified version (IPv4
or IPv6). It panics if an invalid IPVersion is provided.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L153)  

```go
func GetIptable(version IPVersion) *IPTable
```

#### Methods

##### IPTable.AddReturnRule

AddReturnRule adds a return rule for the chain in the filter table

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L408)  

```go
func (iptable IPTable) AddReturnRule(chain string) error
```

##### IPTable.DeleteJumpRule

DeleteJumpRule deletes a rule added by EnsureJumpRule. It's a no-op if the rule
doesn't exist.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L432)  

```go
func (iptable IPTable) DeleteJumpRule(fromChain, toChain string, rule ...string) error
```

##### IPTable.EnsureJumpRule

EnsureJumpRule ensures the jump rule is on top

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L419)  

```go
func (iptable IPTable) EnsureJumpRule(fromChain, toChain string, rule ...string) error
```

##### IPTable.ExistChain

ExistChain checks if a chain exists

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L394)  

```go
func (iptable IPTable) ExistChain(chain string, table Table) bool
```

##### IPTable.Exists

Exists checks if a rule exists

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L284)  

```go
func (iptable IPTable) Exists(table Table, chain string, rule ...string) bool
```

##### IPTable.ExistsNative

ExistsNative behaves as Exists with the difference it
will always invoke `iptables` binary.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L290)  

```go
func (iptable IPTable) ExistsNative(table Table, chain string, rule ...string) bool
```

##### IPTable.NewChain

NewChain adds a new chain to ip table.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L167)  

```go
func (iptable IPTable) NewChain(name string, table Table) (*ChainInfo, error)
```

##### IPTable.ProgramRule

ProgramRule adds the rule specified by args only if the
rule is not already present in the chain. Reciprocally,
it removes the rule only if present.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L231)  

```go
func (iptable IPTable) ProgramRule(table Table, chain string, action Action, args []string) error
```

##### IPTable.Raw

Raw calls 'iptables' system command, passing supplied arguments.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L338)  

```go
func (iptable IPTable) Raw(args ...string) ([]byte, error)
```

##### IPTable.RawCombinedOutput

RawCombinedOutput internally calls the Raw function and returns a non nil
error if Raw returned a non nil error or a non empty output

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L377)  

```go
func (iptable IPTable) RawCombinedOutput(args ...string) error
```

##### IPTable.RawCombinedOutputNative

RawCombinedOutputNative behave as RawCombinedOutput with the difference it
will always invoke `iptables` binary

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L386)  

```go
func (iptable IPTable) RawCombinedOutputNative(args ...string) error
```

##### IPTable.RemoveExistingChain

RemoveExistingChain removes existing chain from the table.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L190)  

```go
func (iptable IPTable) RemoveExistingChain(name string, table Table) error
```

##### IPTable.SetDefaultPolicy

SetDefaultPolicy sets the passed default policy for the table/chain

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L400)  

```go
func (iptable IPTable) SetDefaultPolicy(table Table, chain string, policy Policy) error
```

---

### IPVersion

IPVersion refers to IP version, v4 or v6

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L59)  

```go
type IPVersion string
```

---

### Policy

Policy is the default iptable policies

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L35)  

```go
type Policy string
```

---

### Rule

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L442)  

```go
type Rule struct {
	IPVer IPVersion
	Table Table
	Chain string
	Args  []string
}
```

#### Methods

##### Rule.Append

Append appends the rule to the end of the chain. If the rule already exists anywhere in the
chain, this is a no-op.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L472)  

```go
func (r Rule) Append() error
```

##### Rule.Delete

Delete deletes the rule from the kernel. If the rule does not exist, this is a no-op.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L489)  

```go
func (r Rule) Delete() error
```

##### Rule.Exists

Exists returns true if the rule exists in the kernel.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L450)  

```go
func (r Rule) Exists() bool
```

##### Rule.Insert

Insert inserts the rule at the head of the chain. If the rule already exists anywhere in the
chain, this is a no-op.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L481)  

```go
func (r Rule) Insert() error
```

##### Rule.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L496)  

```go
func (r Rule) String() string
```

##### Rule.WithChain

WithChain returns a version of the rule with its Chain field set to chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L463)  

```go
func (r Rule) WithChain(chain string) Rule
```

---

### Table

Table refers to Nat, Filter or Mangle.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/iptables/iptables.go#L45)  

```go
type Table string
```

---


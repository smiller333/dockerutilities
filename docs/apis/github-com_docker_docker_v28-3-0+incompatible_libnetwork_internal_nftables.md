# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/internal/nftables

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:34:11 UTC

## Overview

Package nftables provides methods to create an nftables table and manage its maps, sets,
chains, and rules.

To use it, the first step is to create a TableRef using NewTable. The table can
then be populated and managed using that ref.

Modifications to the table are only applied (sent to "nft") when TableRef.Apply is
called. This means a number of updates can be made, for example, adding all the
rules needed for a docker network - and those rules will then be applied atomically
in a single "nft" run.

TableRef.Apply can only be called after Enable, and only if Enable returns
true (meaning an "nft" executable was found). Enabled can be called to check
whether nftables has been enabled.

Be aware:


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L98)

```go
const (
	BaseChainPriorityRaw      = -300
	BaseChainPriorityMangle   = -150
	BaseChainPriorityDstNAT   = -100
	BaseChainPriorityFilter   = 0
	BaseChainPrioritySecurity = 50
	BaseChainPrioritySrcNAT   = 100
)
```

## Variables

This section is empty.

## Functions

### Enable

Enable checks whether the "nft" tool is available, and returns true if it is.
Subsequent calls to Enabled will return the same result.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L130)  

```go
func Enable() bool
```

---

### Enabled

Enabled returns true if the "nft" tool is available and Enable has been called.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L146)  

```go
func Enabled() bool
```

---

## Types

### BaseChainHook

BaseChainHook enumerates the base chain hook types.
See https://wiki.nftables.org/wiki-nftables/index.php/Configuring_chains#Base_chain_hooks

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L85)  

```go
type BaseChainHook string
```

---

### BaseChainType

BaseChainType enumerates the base chain types.
See https://wiki.nftables.org/wiki-nftables/index.php/Configuring_chains#Base_chain_types

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L75)  

```go
type BaseChainType string
```

---

### ChainRef

ChainRef is a handle for an nftables chain.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L392)  

```go
type ChainRef struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### ChainRef.AppendRule

AppendRule appends a rule to a RuleGroup in a ChainRef.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L499)  

```go
func (c ChainRef) AppendRule(group RuleGroup, rule string, args ...interface{}) error
```

##### ChainRef.DeleteRule

DeleteRule deletes a rule from a RuleGroup in a ChainRef. It is an error
to delete from a group that does not exist, or to delete a rule that does not
exist.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L521)  

```go
func (c ChainRef) DeleteRule(group RuleGroup, rule string, args ...interface{}) error
```

##### ChainRef.SetPolicy

SetPolicy sets the default policy for a base chain. It is an error to call this
for a non-base ChainRef.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L489)  

```go
func (c ChainRef) SetPolicy(policy string) error
```

---

### ChainUpdateFunc

ChainUpdateFunc is a function that can add rules to a chain, or remove rules from it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L457)  

```go
type ChainUpdateFunc func(RuleGroup, string, ...interface{}) error
```

---

### Family

Family enumerates address families.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L108)  

```go
type Family string
```

---

### RuleGroup

RuleGroup is used to allocate rules within a chain to a group. These groups are
purely an internal construct, nftables knows nothing about them. Within groups
rules retain the order in which they were added, and groups are ordered from
lowest to highest numbered group.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L375)  

```go
type RuleGroup int
```

---

### SetRef

SetRef is a handle for an nftables named set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L648)  

```go
type SetRef struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### SetRef.AddElement

AddElement adds an element to a set. It is the caller's responsibility to make sure
the element has the correct type. It is an error to add an element that is already
in the set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L690)  

```go
func (s SetRef) AddElement(element string) error
```

##### SetRef.DeleteElement

DeleteElement deletes an element from the set. It is an error to delete an
element that is not in the set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L707)  

```go
func (s SetRef) DeleteElement(element string) error
```

---

### TableRef

TableRef is a handle for an nftables table.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L169)  

```go
type TableRef struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewTable

NewTable creates a new nftables table and returns a TableRef

See https://wiki.nftables.org/wiki-nftables/index.php/Configuring_tables

The table will be created and flushed when TableRef.Apply is next called.
It's flushed in case it already exists in the host's nftables - when that
happens, rules in its chains will be deleted but not the chains themselves,
maps, sets, or elements of maps or sets. But, those un-flushed items can't do
anything disruptive unless referred to by rules, and they will be flushed if
they get re-created via the TableRef, when TableRef.Apply is next called
(so, before they can be used by a new rule).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L184)  

```go
func NewTable(family Family, name string) (TableRef, error)
```

#### Methods

##### TableRef.Apply

Apply makes incremental updates to nftables, corresponding to changes to the TableRef
since Apply was last called.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L299)  

```go
func (t TableRef) Apply(ctx context.Context) error
```

##### TableRef.BaseChain

BaseChain constructs a new nftables base chain and returns a ChainRef.

See https://wiki.nftables.org/wiki-nftables/index.php/Configuring_chains#Adding_base_chains

It is an error to create a base chain that already exists.
If the underlying chain already exists, it will be flushed by the
next TableRef.Apply before new rules are added.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L403)  

```go
func (t TableRef) BaseChain(name string, chainType BaseChainType, hook BaseChainHook, priority int) (ChainRef, error)
```

##### TableRef.Chain

Chain returns a ChainRef for an existing chain (which may be a base chain).
If there is no existing chain, a regular chain is added and its ChainRef is
returned.

See https://wiki.nftables.org/wiki-nftables/index.php/Configuring_chains#Adding_regular_chains

If a new ChainRef is created and the underlying chain already exists, it
will be flushed by the next TableRef.Apply before new rules are added.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L437)  

```go
func (t TableRef) Chain(name string) ChainRef
```

##### TableRef.ChainUpdateFunc

ChainUpdateFunc returns a ChainUpdateFunc to add rules to the named chain if
enable is true, or to remove rules from the chain if enable is false.
(Written as a convenience function to ease migration of iptables functions
originally written with an enable flag.)

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L463)  

```go
func (t TableRef) ChainUpdateFunc(name string, enable bool) ChainUpdateFunc
```

##### TableRef.DeleteChain

DeleteChain deletes a chain. It is an error to delete a chain that does not exist.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L472)  

```go
func (t TableRef) DeleteChain(name string) error
```

##### TableRef.Family

Family returns the address family of the nftables table described by TableRef.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L199)  

```go
func (t TableRef) Family() Family
```

##### TableRef.InterfaceVMap

InterfaceVMap creates a map from interface name to a verdict and returns a VMapRef,
or returns an existing VMapRef if it has already been created.

See https://wiki.nftables.org/wiki-nftables/index.php/Verdict_Maps_(vmaps)

If a VMapRef is created and the underlying map already exists, it will be flushed
by the next TableRef.Apply before new elements are added.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L573)  

```go
func (t TableRef) InterfaceVMap(name string) VMapRef
```

##### TableRef.PrefixSet

PrefixSet creates a new named nftables set for IPv4 or IPv6 addresses (depending
on the address family of the TableRef), and returns its SetRef. Or, if the
set has already been created, its SetRef is returned.

(TableRef does not support "inet", only "ip" or "ip6". So the element type can
always be determined. But, there's no "inet" element type, so this will need to
change if we need an "inet" table.)

See https://wiki.nftables.org/wiki-nftables/index.php/Sets#Named_sets

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L661)  

```go
func (t TableRef) PrefixSet(name string) SetRef
```

##### TableRef.Reload

Reload deletes the table, then re-creates it, atomically.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L336)  

```go
func (t TableRef) Reload(ctx context.Context) error
```

---

### VMapRef

VMapRef is a handle for an nftables verdict map.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L562)  

```go
type VMapRef struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### VMapRef.AddElement

AddElement adds an element to a verdict map. The caller must ensure the key has
the correct type. It is an error to add a key that already exists.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L597)  

```go
func (v VMapRef) AddElement(key string, verdict string) error
```

##### VMapRef.DeleteElement

DeleteElement deletes an element from a verdict map. It is an error to delete
an element that does not exist.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/internal/nftables/nftables_linux.go#L615)  

```go
func (v VMapRef) DeleteElement(key string) error
```

---


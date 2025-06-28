# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/etchosts

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:33:55 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Add

Add adds an arbitrary number of Records to an already existing /etc/hosts file

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/etchosts/etchosts.go#L108)  

```go
func Add(path string, recs []Record) error
```

---

### Build

Build function
path is path to host file string required
extraContent is an array of extra host records.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/etchosts/etchosts.go#L75)  

```go
func Build(path string, extraContent []Record) error
```

---

### BuildNoIPv6

BuildNoIPv6 is the same as Build, but will not include IPv6 entries.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/etchosts/etchosts.go#L80)  

```go
func BuildNoIPv6(path string, extraContent []Record) error
```

---

### Delete

Delete deletes Records from /etc/hosts.
The hostnames must be an exact match (if the user has modified the record,
it won't be deleted). The address, parsed as a netip.Addr must also match
the value in recs.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/etchosts/etchosts.go#L135)  

```go
func Delete(path string, recs []Record) error
```

---

### Drop

Drop drops the path string from the path cache

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/etchosts/etchosts.go#L65)  

```go
func Drop(path string)
```

---

### Update

Update all IP addresses where hostname matches.
path is path to host file
IP is new IP address
hostname is hostname to search for to replace IP

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/etchosts/etchosts.go#L186)  

```go
func Update(path, IP, hostname string) error
```

---

## Types

### Record

Record Structure for a single host record

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/etchosts/etchosts.go#L16)  

```go
type Record struct {
	Hosts string
	IP    netip.Addr
}
```

#### Methods

##### Record.WriteTo

WriteTo writes record to file and returns bytes written or error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/etchosts/etchosts.go#L22)  

```go
func (r Record) WriteTo(w io.Writer) (int64, error)
```

---


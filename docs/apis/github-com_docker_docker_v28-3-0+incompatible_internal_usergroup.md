# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/usergroup

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:08:08 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AddNamespaceRangesUser

AddNamespaceRangesUser takes a username and uses the standard system
utility to create a system user/group pair used to hold the
/etc/sub{uid,gid} ranges which will be used for user namespace
mapping ranges in containers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/usergroup/add_linux.go#L36)  

```go
func AddNamespaceRangesUser(name string) (int, int, error)
```

---

### LoadIdentityMapping

LoadIdentityMapping takes a requested username and
using the data from /etc/sub{uid,gid} ranges, creates the
proper uid and gid remapping ranges for that user/group pair

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/usergroup/lookup_unix.go#L142)  

```go
func LoadIdentityMapping(name string) (user.IdentityMapping, error)
```

---

### LookupGID

LookupGID uses traditional local system files lookup (from libcontainer/user) on a group ID,
followed by a call to `getent` for supporting host configured non-files passwd and group dbs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/usergroup/lookup_unix.go#L73)  

```go
func LookupGID(gid int) (user.Group, error)
```

---

### LookupGroup

LookupGroup uses traditional local system files lookup (from libcontainer/user) on a group name,
followed by a call to `getent` for supporting host configured non-files passwd and group dbs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/usergroup/lookup_unix.go#L61)  

```go
func LookupGroup(name string) (user.Group, error)
```

---

### LookupUID

LookupUID uses traditional local system files lookup (from libcontainer/user) on a uid,
followed by a call to `getent` for supporting host configured non-files passwd and group dbs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/usergroup/lookup_unix.go#L34)  

```go
func LookupUID(uid int) (user.User, error)
```

---

### LookupUser

LookupUser uses traditional local system files lookup (from libcontainer/user) on a username,
followed by a call to `getent` for supporting host configured non-files passwd and group dbs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/usergroup/lookup_unix.go#L18)  

```go
func LookupUser(name string) (user.User, error)
```

---

## Types

This section is empty.


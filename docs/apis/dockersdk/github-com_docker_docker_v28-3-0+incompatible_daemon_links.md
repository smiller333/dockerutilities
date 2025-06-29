# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/links

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:47 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### EnvVars

EnvVars generates environment variables for the linked container
for the Link with the given options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/links/links.go#L27)  

```go
func EnvVars(parentIP, childIP, name string, env []string, exposedPorts map[nat.Port]struct{}) []string
```

---

## Types

### Link

Link struct holds information about parent/child linked container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/links/links.go#L12)  

```go
type Link struct {
	// Parent container IP address
	ParentIP string
	// Child container IP address
	ChildIP string
	// Link name
	Name string
	// Child environments variables
	ChildEnvironment []string
	// Child exposed ports
	Ports []nat.Port
}
```

#### Functions

##### NewLink

NewLink initializes a new Link struct with the provided options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/links/links.go#L32)  

```go
func NewLink(parentIP, childIP, name string, env []string, exposedPorts map[nat.Port]struct{}) *Link
```

#### Methods

##### Link.ToEnv

ToEnv creates a string's slice containing child container information in
the form of environment variables which will be later exported on container
startup.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/links/links.go#L50)  

```go
func (l *Link) ToEnv() []string
```

---


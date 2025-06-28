# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/volume/service/opts

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:54 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### WithGetResolveStatus

WithGetResolveStatus indicates to `Get` to also fetch the volume status.
This can cause significant overhead in the volume lookup.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L80)  

```go
func WithGetResolveStatus(cfg *GetConfig)
```

---

## Types

### CreateConfig

CreateConfig is the set of config options that can be set when creating
a volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L8)  

```go
type CreateConfig struct {
	Options   map[string]string
	Labels    map[string]string
	Reference string
}
```

---

### CreateOption

CreateOption is used to pass options in when creating a volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L4)  

```go
type CreateOption func(*CreateConfig)
```

#### Functions

##### WithCreateLabel

WithCreateLabel creates a CreateOption which adds a label with the given key/value pair

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L15)  

```go
func WithCreateLabel(key, value string) CreateOption
```

##### WithCreateLabels

WithCreateLabels creates a CreateOption which sets the labels to the
passed in value

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L26)  

```go
func WithCreateLabels(labels map[string]string) CreateOption
```

##### WithCreateOptions

WithCreateOptions creates a CreateOption which sets the options passed
to the volume driver when creating a volume to the options passed in.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L34)  

```go
func WithCreateOptions(opts map[string]string) CreateOption
```

##### WithCreateReference

WithCreateReference creates a CreateOption which sets a reference to use
when creating a volume. This ensures that the volume is created with a reference
already attached to it to prevent race conditions with Create and volume cleanup.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L43)  

```go
func WithCreateReference(ref string) CreateOption
```

---

### GetConfig

GetConfig is used with `GetOption` to set options for the volumes service's
`Get` implementation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L51)  

```go
type GetConfig struct {
	Driver        string
	Reference     string
	ResolveStatus bool
}
```

---

### GetOption

GetOption is passed to the service `Get` add extra details on the get request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L58)  

```go
type GetOption func(*GetConfig)
```

#### Functions

##### WithGetDriver

WithGetDriver provides the driver to get the volume from
If no driver is provided to `Get`, first the available metadata is checked
to see which driver it belongs to, if that is not available all drivers are
probed to find the volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L64)  

```go
func WithGetDriver(name string) GetOption
```

##### WithGetReference

WithGetReference indicates to `Get` to increment the reference count for the
retrieved volume with the provided reference ID.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L72)  

```go
func WithGetReference(ref string) GetOption
```

---

### RemoveConfig

RemoveConfig is used by `RemoveOption` to store config options for remove

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L85)  

```go
type RemoveConfig struct {
	PurgeOnError bool
}
```

---

### RemoveOption

RemoveOption is used to pass options to the volumes service `Remove` implementation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L90)  

```go
type RemoveOption func(*RemoveConfig)
```

#### Functions

##### WithPurgeOnError

WithPurgeOnError is an option passed to `Remove` which will purge all cached
data about a volume even if there was an error while attempting to remove the
volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/volume/service/opts/opts.go#L95)  

```go
func WithPurgeOnError(b bool) RemoveOption
```

---


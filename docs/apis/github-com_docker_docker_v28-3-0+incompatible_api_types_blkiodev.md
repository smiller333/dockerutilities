# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/blkiodev

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:01:45 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### ThrottleDevice

ThrottleDevice is a structure that holds device:rate_per_second pair

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/blkiodev/blkio.go#L16)  

```go
type ThrottleDevice struct {
	Path string
	Rate uint64
}
```

#### Methods

##### ThrottleDevice.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/blkiodev/blkio.go#L21)  

```go
func (t *ThrottleDevice) String() string
```

---

### WeightDevice

WeightDevice is a structure that holds device:weight pair

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/blkiodev/blkio.go#L6)  

```go
type WeightDevice struct {
	Path   string
	Weight uint16
}
```

#### Methods

##### WeightDevice.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/blkiodev/blkio.go#L11)  

```go
func (w *WeightDevice) String() string
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/stats

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:49 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Collector

Collector manages and provides container resource stats

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/stats/collector.go#L13)  

```go
type Collector struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewCollector

NewCollector creates a stats collector that will poll the supervisor with the specified interval

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/stats/collector.go#L22)  

```go
func NewCollector(supervisor supervisor, interval time.Duration) *Collector
```

#### Methods

##### Collector.Collect

Collect registers the container with the collector and adds it to
the event loop for collection on the specified interval returning
a channel for the subscriber to receive on.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/stats/collector.go#L40)  

```go
func (s *Collector) Collect(c *container.Container) chan interface{}
```

##### Collector.Run

Run starts the collectors and will indefinitely collect stats from the supervisor

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/stats/collector.go#L79)  

```go
func (s *Collector) Run()
```

##### Collector.StopCollection

StopCollection closes the channels for all subscribers and removes
the container from metrics collection.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/stats/collector.go#L56)  

```go
func (s *Collector) StopCollection(c *container.Container)
```

##### Collector.Unsubscribe

Unsubscribe removes a specific subscriber from receiving updates for a container's stats.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/stats/collector.go#L66)  

```go
func (s *Collector) Unsubscribe(c *container.Container, ch chan interface{})
```

---


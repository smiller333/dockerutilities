# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/metrics

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:15 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/metrics.go#L10)

```go
var (

	// ContainerActions tracks the time taken to process container operations
	ContainerActions = metricsNS.NewLabeledTimer("container_actions", "The number of seconds it takes to process each container action", "action")
	// NetworkActions tracks the time taken to process network operations
	NetworkActions = metricsNS.NewLabeledTimer("network_actions", "The number of seconds it takes to process each network action", "action")
	// HostInfoFunctions tracks the time taken to gather host information
	HostInfoFunctions = metricsNS.NewLabeledTimer("host_info_functions", "The number of seconds it takes to call functions gathering info about the host", "function")
	ImageActions      = metricsNS.NewLabeledTimer("image_actions", "The number of seconds it takes to process each image action", "action")

	// EngineInfo provides information about the engine and its environment
	EngineInfo = metricsNS.NewLabeledGauge("engine", "The information related to the engine and the OS it is running on", gometrics.Unit("info"),
		"version",
		"commit",
		"architecture",
		"graphdriver",
		"kernel",
		"os",
		"os_type",
		"os_version",
		"daemon_id",
	)
	// EngineCPUs tracks the number of CPUs available to the engine
	EngineCPUs = metricsNS.NewGauge("engine_cpus", "The number of cpus that the host system of the engine has", gometrics.Unit("cpus"))
	// EngineMemory tracks the amount of memory available to the engine
	EngineMemory = metricsNS.NewGauge("engine_memory", "The number of bytes of memory that the host system of the engine has", gometrics.Bytes)

	// HealthChecksCounter tracks the total number of health checks
	HealthChecksCounter = metricsNS.NewCounter("health_checks", "The total number of health checks")
	// HealthChecksFailedCounter tracks the number of failed health checks
	HealthChecksFailedCounter = metricsNS.NewCounter("health_checks_failed", "The total number of failed health checks")
	// HealthCheckStartDuration tracks the time taken to prepare health checks
	HealthCheckStartDuration = metricsNS.NewTimer("health_check_start_duration", "The number of seconds it takes to prepare to run health checks")

	// StateCtr tracks container states
	StateCtr = newStateCounter(metricsNS, metricsNS.NewDesc("container_states", "The count of containers in various states", gometrics.Unit("containers"), "state"))

	// EventsCounter tracks the number of events logged
	EventsCounter = metricsNS.NewCounter("events", "The number of events logged")

	// EventSubscribers tracks the number of current subscribers to events
	EventSubscribers = metricsNS.NewGauge("events_subscribers", "The number of current subscribers to events", gometrics.Total)
)
```

## Functions

### CleanupPlugin

CleanupPlugin stops metrics collection for all plugins

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/plugin_unix.go#L82)  

```go
func CleanupPlugin(store plugingetter.PluginGetter)
```

---

### RegisterPlugin

RegisterPlugin starts the metrics server listener and registers the metrics plugin
callback with the plugin store

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/plugin_unix.go#L50)  

```go
func RegisterPlugin(store *plugin.Store, path string) error
```

---

### StartTimer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/metrics.go#L69)  

```go
func StartTimer(t gometrics.Timer) func()
```

---

## Types

### Plugin

Plugin represents a metrics collector plugin

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/plugin_unix.go#L26)  

```go
type Plugin interface {
	StartMetrics() error
	StopMetrics() error
}
```

---

### StateCounter

StateCounter tracks container states

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/metrics.go#L74)  

```go
type StateCounter struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### StateCounter.Collect

Collect implements prometheus.Collector

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/metrics.go#L129)  

```go
func (ctr *StateCounter) Collect(ch chan<- prometheus.Metric)
```

##### StateCounter.Delete

Delete removes a container's state

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/metrics.go#L116)  

```go
func (ctr *StateCounter) Delete(id string)
```

##### StateCounter.Describe

Describe implements prometheus.Collector

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/metrics.go#L124)  

```go
func (ctr *StateCounter) Describe(ch chan<- *prometheus.Desc)
```

##### StateCounter.Get

Get returns the count of containers in running, paused, and stopped states

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/metrics.go#L90)  

```go
func (ctr *StateCounter) Get() (running int, paused int, stopped int)
```

##### StateCounter.Set

Set updates the state for a container

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/metrics/metrics.go#L108)  

```go
func (ctr *StateCounter) Set(id, label string)
```

---


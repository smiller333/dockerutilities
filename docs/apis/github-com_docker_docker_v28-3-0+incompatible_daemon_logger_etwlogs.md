# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/etwlogs

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:00 UTC

## Overview

Package etwlogs provides a log driver for forwarding container logs
as ETW events.(ETW stands for Event Tracing for Windows)
A client can then create an ETW listener to listen for events that are sent
by the ETW provider that we register, using the provider's GUID "a3693192-9ed6-46d2-a981-f8226c8363bd".
Here is an example of how to do this using the logman utility:
1. logman start -ets DockerContainerLogs -p {a3693192-9ed6-46d2-a981-f8226c8363bd} 0 0 -o trace.etl
2. Run container(s) and generate log messages
3. logman stop -ets DockerContainerLogs
4. You can then convert the etl log file to XML using: tracerpt -y trace.etl

Each container log message generates an ETW event that also contains:
the container name and ID, the timestamp, and the stream type.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New creates a new etwLogs logger for the given container and registers the EWT provider.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/etwlogs/etwlogs_windows.go#L60)  

```go
func New(info logger.Info) (logger.Logger, error)
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/process

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:14:57 UTC

## Overview

Package process provides a set of basic functions to manage individual
processes.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Alive

Alive returns true if process with a given pid is running. It only considers
positive PIDs; 0 (all processes in the current process group), -1 (all processes
with a PID larger than 1), and negative (-n, all processes in process group
"n") values for pid are never considered to be alive.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/process/process_unix.go#L21)  

```go
func Alive(pid int) bool
```

---

### Kill

Kill force-stops a process. It only considers positive PIDs; 0 (all processes
in the current process group), -1 (all processes with a PID larger than 1),
and negative (-n, all processes in process group "n") values for pid are
ignored. Refer to KILL(2) for details.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/process/process_unix.go#L50)  

```go
func Kill(pid int) error
```

---

### Zombie

Zombie return true if process has a state with "Z". It only considers positive
PIDs; 0 (all processes in the current process group), -1 (all processes with
a PID larger than 1), and negative (-n, all processes in process group "n")
values for pid are ignored. Refer to PROC(5) for details.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/process/process_unix.go#L67)  

```go
func Zombie(pid int) (bool, error)
```

---

## Types

This section is empty.


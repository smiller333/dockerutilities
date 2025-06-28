# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/pidfile

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:55 UTC

## Overview

Package pidfile provides structure and helper functions to create and remove
PID file. A PID file is usually a file used to store the process ID of a
running process.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Read

Read reads the "PID file" at path, and returns the PID if it contains a
valid PID of a running process, or 0 otherwise. It returns an error when
failing to read the file, or if the file doesn't exist, but malformed content
is ignored. Consumers should therefore check if the returned PID is a non-zero
value before use.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pidfile/pidfile.go#L20)  

```go
func Read(path string) (pid int, _ error)
```

---

### Write

Write writes a "PID file" at the specified path. It returns an error if the
file exists and contains a valid PID of a running process, or when failing
to write the file.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/pidfile/pidfile.go#L38)  

```go
func Write(path string, pid int) error
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/modprobe

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:07:20 UTC

## Overview

Package modprobe attempts to load kernel modules. It may have more success
than simply running "modprobe", particularly for docker-in-docker.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### LoadModules

LoadModules attempts to load kernel modules, if necessary.

isLoaded must be a function that checks whether the modules are loaded. It may
be called multiple times. isLoaded must return an error to indicate that the
modules still need to be loaded, otherwise nil.

For each method of loading modules, LoadModules will attempt the load for each
of modNames, then it will call isLoaded to check the result - moving on to try
the next method if needed, and there is one.

The returned error is the result of the final call to isLoaded.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/modprobe/modprobe_linux.go#L27)  

```go
func LoadModules(ctx context.Context, isLoaded func() error, modNames ...string) error
```

---

## Types

This section is empty.


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/overlay/overlayutils

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:33:38 UTC

## Overview

Package overlayutils provides utility functions for overlay networks


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### AppendVNIList

AppendVNIList appends the VNI values encoded as a CSV string to slice.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlayutils/utils.go#L47)  

```go
func AppendVNIList(vnis []uint32, csv string) ([]uint32, error)
```

---

### ConfigVXLANUDPPort

ConfigVXLANUDPPort configures the VXLAN UDP port (data path port) number.
If no port is set, the default (4789) is returned. Valid port numbers are
between 1024 and 49151.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlayutils/utils.go#L21)  

```go
func ConfigVXLANUDPPort(vxlanPort uint32) error
```

---

### VXLANUDPPort

VXLANUDPPort returns Vxlan UDP port number

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/overlay/overlayutils/utils.go#L40)  

```go
func VXLANUDPPort() uint32
```

---

## Types

This section is empty.


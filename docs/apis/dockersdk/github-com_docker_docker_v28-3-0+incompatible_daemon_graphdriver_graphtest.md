# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/graphdriver/graphtest

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:16 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### DriverBenchDeepLayerDiff

DriverBenchDeepLayerDiff benchmarks calls to diff on top of a given number of layers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphbench_unix.go#L190)  
**Added in:** v1.12.0

```go
func DriverBenchDeepLayerDiff(b *testing.B, layerCount int, drivername string, driveroptions ...string)
```

---

### DriverBenchDeepLayerRead

DriverBenchDeepLayerRead benchmarks calls to read a file under a given number of layers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphbench_unix.go#L223)  
**Added in:** v1.12.0

```go
func DriverBenchDeepLayerRead(b *testing.B, layerCount int, drivername string, driveroptions ...string)
```

---

### DriverBenchDiffApplyN

DriverBenchDiffApplyN benchmarks calls to diff and apply together

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphbench_unix.go#L124)  
**Added in:** v1.12.0

```go
func DriverBenchDiffApplyN(b *testing.B, fileCount int, drivername string, driveroptions ...string)
```

---

### DriverBenchDiffBase

DriverBenchDiffBase benchmarks calls to diff on a root layer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphbench_unix.go#L60)  
**Added in:** v1.12.0

```go
func DriverBenchDiffBase(b *testing.B, drivername string, driveroptions ...string)
```

---

### DriverBenchDiffN

DriverBenchDiffN benchmarks calls to diff on two layers with
a provided number of files on the lower and upper layers.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphbench_unix.go#L89)  
**Added in:** v1.12.0

```go
func DriverBenchDiffN(b *testing.B, bottom, top int, drivername string, driveroptions ...string)
```

---

### DriverBenchExists

DriverBenchExists benchmarks calls to exist

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphbench_unix.go#L16)  
**Added in:** v1.12.0

```go
func DriverBenchExists(b *testing.B, drivername string, driveroptions ...string)
```

---

### DriverBenchGetEmpty

DriverBenchGetEmpty benchmarks calls to get on an empty layer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphbench_unix.go#L35)  
**Added in:** v1.12.0

```go
func DriverBenchGetEmpty(b *testing.B, drivername string, driveroptions ...string)
```

---

### DriverTestChanges

DriverTestChanges tests computed changes on a layer matches changes made

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L265)  
**Added in:** v1.12.0

```go
func DriverTestChanges(t testing.TB, drivername string, driverOptions ...string)
```

---

### DriverTestCreateBase

DriverTestCreateBase create a base driver and verify.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L113)  

```go
func DriverTestCreateBase(t testing.TB, drivername string, driverOptions ...string)
```

---

### DriverTestCreateEmpty

DriverTestCreateEmpty creates a new image and verifies it is empty and the right metadata

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L84)  

```go
func DriverTestCreateEmpty(t testing.TB, drivername string, driverOptions ...string)
```

---

### DriverTestCreateSnap

DriverTestCreateSnap Create a driver and snap and verify.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L125)  

```go
func DriverTestCreateSnap(t testing.TB, drivername string, driverOptions ...string)
```

---

### DriverTestDeepLayerRead

DriverTestDeepLayerRead reads a file from a lower layer under a given number of layers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L144)  
**Added in:** v1.12.0

```go
func DriverTestDeepLayerRead(t testing.TB, layerCount int, drivername string, driverOptions ...string)
```

---

### DriverTestDiffApply

DriverTestDiffApply tests diffing and applying produces the same layer

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L174)  
**Added in:** v1.12.0

```go
func DriverTestDiffApply(t testing.TB, fileCount int, drivername string, driverOptions ...string)
```

---

### DriverTestSetQuota

DriverTestSetQuota Create a driver and test setting quota.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L307)  
**Added in:** v1.12.0

```go
func DriverTestSetQuota(t *testing.T, drivername string, required bool)
```

---

### GetDriver

GetDriver create a new driver with given name or return an existing driver with the name updating the reference count.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L62)  

```go
func GetDriver(t testing.TB, name string, options ...string) graphdriver.Driver
```

---

### PutDriver

PutDriver removes the driver if it is no longer used and updates the reference count.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L72)  

```go
func PutDriver(t testing.TB)
```

---

## Types

### Driver

Driver conforms to graphdriver.Driver interface and
contains information such as root and reference count of the number of clients using it.
This helps in testing drivers added into the framework.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/graphdriver/graphtest/graphtest_unix.go#L27)  

```go
type Driver struct {
	graphdriver.Driver
	// contains filtered or unexported fields
}
```

---


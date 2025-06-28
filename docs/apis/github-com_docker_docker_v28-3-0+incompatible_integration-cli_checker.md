# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration-cli/checker

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:48 UTC

## Overview

Package checker provides helpers for gotest.tools/assert.
Please remove this package whenever possible.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Compare

Compare defines the interface to compare values

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L13)  

```go
type Compare func(x interface{}) assert.BoolOrComparison
```

#### Functions

##### Contains

Contains checks if the value contains the given value

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L37)  

```go
func Contains(y interface{}) Compare
```

##### DeepEquals

DeepEquals checks if two values are equal

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L59)  

```go
func DeepEquals(y interface{}) Compare
```

##### Equals

Equals checks if the value is equal to the given value

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L30)  

```go
func Equals(y interface{}) Compare
```

##### False

False checks if the value is false

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L16)  

```go
func False() Compare
```

##### GreaterThan

GreaterThan checks if the value is greater than the given value

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L80)  

```go
func GreaterThan(y int) Compare
```

##### HasLen

HasLen checks if the value has the expected number of elements

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L66)  

```go
func HasLen(y int) Compare
```

##### IsNil

IsNil checks if the value is nil

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L73)  

```go
func IsNil() Compare
```

##### Not

Not checks if two values are not

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L44)  

```go
func Not(c Compare) Compare
```

##### True

True checks if the value is true

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/checker/checker.go#L23)  

```go
func True() Compare
```

---


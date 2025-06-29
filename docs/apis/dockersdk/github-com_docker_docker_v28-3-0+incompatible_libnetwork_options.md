# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/options

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:10:46 UTC

## Overview

Package options provides a way to pass unstructured sets of options to a
component expecting a strongly-typed configuration structure.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### GenerateFromModel

GenerateFromModel takes the generic options, and tries to build a new
instance of the model's type by matching keys from the generic options to
fields in the model.

The return value is of the same type than the model (including a potential
pointer qualifier).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/options/options.go#L56)  

```go
func GenerateFromModel(options Generic, model any) (any, error)
```

---

## Types

### CannotSetFieldError

CannotSetFieldError is the error returned when the generic parameters hold a
value for a field that cannot be set in the destination structure.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/options/options.go#L26)  

```go
type CannotSetFieldError struct {
	Field string
	Type  string
}
```

#### Methods

##### CannotSetFieldError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/options/options.go#L31)  

```go
func (e CannotSetFieldError) Error() string
```

---

### Generic

Generic is a basic type to store arbitrary settings.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/options/options.go#L48)  

```go
type Generic map[string]any
```

---

### NoSuchFieldError

NoSuchFieldError is the error returned when the generic parameters hold a
value for a field absent from the destination structure.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/options/options.go#L15)  

```go
type NoSuchFieldError struct {
	Field string
	Type  string
}
```

#### Methods

##### NoSuchFieldError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/options/options.go#L20)  

```go
func (e NoSuchFieldError) Error() string
```

---

### TypeMismatchError

TypeMismatchError is the error returned when the type of the generic value
for a field mismatches the type of the destination structure.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/options/options.go#L37)  

```go
type TypeMismatchError struct {
	Field      string
	ExpectType string
	ActualType string
}
```

#### Methods

##### TypeMismatchError.Error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/options/options.go#L43)  

```go
func (e TypeMismatchError) Error() string
```

---


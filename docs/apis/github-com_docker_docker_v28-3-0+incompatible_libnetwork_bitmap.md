# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/bitmap

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:08:32 UTC

## Overview

Package bitmap provides a datatype for long vectors of bits.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L21)

```go
var (
	// ErrNoBitAvailable is returned when no more bits are available to set
	ErrNoBitAvailable = errors.New("no bit available")
	// ErrBitAllocated is returned when the specific bit requested is already set
	ErrBitAllocated = errors.New("requested bit is already allocated")
)
```

## Functions

This section is empty.

## Types

### Bitmap

Bitmap is a fixed-length bit vector. It is not safe for concurrent use.

The data is stored as a list of run-length encoded blocks. It operates
directly on the encoded representation, without decompressing.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L37)  

```go
type Bitmap struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### Copy

Copy returns a deep copy of b.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L63)  

```go
func Copy(b *Bitmap) *Bitmap
```

##### New

New returns a new Bitmap of ordinals in the interval [0, n).

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L51)  

```go
func New(n uint64) *Bitmap
```

#### Methods

##### Bitmap.Bits

Bits returns the length of the bit sequence

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L325)  

```go
func (h *Bitmap) Bits() uint64
```

##### Bitmap.IsSet

IsSet atomically checks if the ordinal bit is set. In case ordinal
is outside of the bit sequence limits, false is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L226)  

```go
func (h *Bitmap) IsSet(ordinal uint64) bool
```

##### Bitmap.MarshalBinary

MarshalBinary encodes h into a binary representation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L288)  

```go
func (h *Bitmap) MarshalBinary() ([]byte, error)
```

##### Bitmap.MarshalJSON

MarshalJSON encodes h into a JSON message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L340)  

```go
func (h *Bitmap) MarshalJSON() ([]byte, error)
```

##### Bitmap.Set

Set atomically sets the corresponding bit in the sequence

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L207)  

```go
func (h *Bitmap) Set(ordinal uint64) error
```

##### Bitmap.SetAny

SetAny sets the first unset bit in the sequence and returns the ordinal of
the set bit.

When serial=true, the bitmap is scanned starting from the ordinal following
the bit most recently set by Bitmap.SetAny or Bitmap.SetAnyInRange.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L199)  

```go
func (h *Bitmap) SetAny(serial bool) (uint64, error)
```

##### Bitmap.SetAnyInRange

SetAnyInRange sets the first unset bit in the range [start, end] and returns
the ordinal of the set bit.

When serial=true, the bitmap is scanned starting from the ordinal following
the bit most recently set by Bitmap.SetAny or Bitmap.SetAnyInRange.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L184)  

```go
func (h *Bitmap) SetAnyInRange(start, end uint64, serial bool) (uint64, error)
```

##### Bitmap.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L334)  

```go
func (h *Bitmap) String() string
```

##### Bitmap.UnmarshalBinary

UnmarshalBinary decodes a binary representation of a Bitmap value which was
generated using Bitmap.MarshalBinary.

The scan position for serial Bitmap.SetAny and Bitmap.SetAnyInRange
operations is neither unmarshaled nor reset.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L307)  

```go
func (h *Bitmap) UnmarshalBinary(ba []byte) error
```

##### Bitmap.UnmarshalJSON

UnmarshalJSON decodes JSON message into h

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L349)  

```go
func (h *Bitmap) UnmarshalJSON(data []byte) error
```

##### Bitmap.Unselected

Unselected returns the number of bits which are not selected

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L330)  

```go
func (h *Bitmap) Unselected() uint64
```

##### Bitmap.Unset

Unset atomically unsets the corresponding bit in the sequence

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/bitmap/sequence.go#L216)  

```go
func (h *Bitmap) Unset(ordinal uint64) error
```

---


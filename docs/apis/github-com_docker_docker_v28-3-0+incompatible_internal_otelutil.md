# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/otelutil

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:31:38 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/baggage.go#L12)

```go
const TriggerKey = "trigger"
```

## Variables

This section is empty.

## Functions

### MustNewBaggage

MustNewBaggage creates an OTel Baggage containing the provided members. It
panics if the baggage cannot be created.

DO NOT USE this function with dynamic values.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/baggage.go#L18)  

```go
func MustNewBaggage(members ...baggage.Member) baggage.Baggage
```

---

### MustNewMemberRaw

MustNewMemberRaw creates an OTel Baggage member with the provided key and
value. It panics if the key or value aren't valid UTF-8 strings.

DO NOT USE this function with dynamic key/value.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/baggage.go#L33)  

```go
func MustNewMemberRaw(key, value string) baggage.Member
```

---

### NewTracerProvider

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/provider.go#L16)  

```go
func NewTracerProvider(ctx context.Context, allowNoop bool) (trace.TracerProvider, func(context.Context) error)
```

---

### RecordStatus

RecordStatus records the status of a span based on the error provided.

If err is nil, the span status is unmodified. If err is not nil, the span
takes status Error, and the error message is recorded.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/status.go#L12)  

```go
func RecordStatus(span trace.Span, err error)
```

---

## Types

### EnvironCarrier

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/environ_carrier.go#L17)  

```go
type EnvironCarrier struct {
	TraceParent, TraceState string
}
```

#### Functions

##### PropagateFromEnvironment

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/environ_carrier.go#L66)  

```go
func PropagateFromEnvironment() *EnvironCarrier
```

#### Methods

##### EnvironCarrier.Environ

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/environ_carrier.go#L55)  

```go
func (c *EnvironCarrier) Environ() []string
```

##### EnvironCarrier.Get

Get returns the value associated with the passed key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/environ_carrier.go#L22)  

```go
func (c *EnvironCarrier) Get(key string) string
```

##### EnvironCarrier.Keys

Keys lists the keys stored in this carrier.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/environ_carrier.go#L44)  

```go
func (c *EnvironCarrier) Keys() []string
```

##### EnvironCarrier.Set

Set stores the key-value pair.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/otelutil/environ_carrier.go#L33)  

```go
func (c *EnvironCarrier) Set(key, value string)
```

---


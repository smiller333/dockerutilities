# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/events

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:25:58 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Action

Action is used for event-actions.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/events/events.go#L24)  

```go
type Action string
```

---

### Actor

Actor describes something that generates events,
like a container, or a network, or a volume.
It has a defined name and a set of attributes.
The container attributes are its labels, other actors
can generate these attributes from other properties.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/events/events.go#L108)  

```go
type Actor struct {
	ID         string
	Attributes map[string]string
}
```

---

### ListOptions

ListOptions holds parameters to filter events with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/events/events.go#L135)  

```go
type ListOptions struct {
	Since   string
	Until   string
	Filters filters.Args
}
```

---

### Message

Message represents the information an event contains

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/events/events.go#L114)  

```go
type Message struct {
	// Deprecated: use Action instead.
	// Information from JSONMessage.
	// With data only in container events.
	Status string `json:"status,omitempty"`
	// Deprecated: use Actor.ID instead.
	ID string `json:"id,omitempty"`
	// Deprecated: use Actor.Attributes["image"] instead.
	From string `json:"from,omitempty"`

	Type   Type
	Action Action
	Actor  Actor
	// Engine events are local scope. Cluster events are swarm scope.
	Scope string `json:"scope,omitempty"`

	Time     int64 `json:"time,omitempty"`
	TimeNano int64 `json:"timeNano,omitempty"`
}
```

---

### Type

Type is used for event-types.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/events/events.go#L6)  

```go
type Type string
```

---


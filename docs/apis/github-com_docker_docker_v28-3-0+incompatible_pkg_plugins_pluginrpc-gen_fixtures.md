# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/plugins/pluginrpc-gen/fixtures

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:36:07 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Bar

Bar is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L40)  

```go
type Bar interface {
	Boo(a string, b string) (s string, err error)
}
```

---

### Fooer

Fooer is an empty interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L16)  

```go
type Fooer interface{}
```

---

### Fooer10

Fooer10 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L71)  
**Added in:** v1.12.0

```go
type Fooer10 interface {
	Foo(a []otherfixture.Spaceship)
}
```

---

### Fooer11

Fooer11 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L76)  
**Added in:** v1.12.0

```go
type Fooer11 interface {
	Foo(a []*otherfixture.Spaceship)
}
```

---

### Fooer12

Fooer12 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L81)  
**Added in:** v1.12.0

```go
type Fooer12 interface {
	Foo(a aliasedio.Reader)
}
```

---

### Fooer2

Fooer2 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L19)  

```go
type Fooer2 interface {
	Foo()
}
```

---

### Fooer3

Fooer3 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L24)  

```go
type Fooer3 interface {
	Foo()
	Bar(a string)
	Baz(a string) (err error)
	Qux(a, b string) (val string, err error)
	Wobble() (w *wobble)
	Wiggle() (w wobble)
	WiggleWobble(a []*wobble, b []wobble, c map[string]*wobble, d map[*wobble]wobble, e map[string][]wobble, f []*otherfixture.Spaceship) (g map[*wobble]wobble, h [][]*wobble, i otherfixture.Spaceship, j *otherfixture.Spaceship, k map[*otherfixture.Spaceship]otherfixture.Spaceship, l []otherfixture.Spaceship)
}
```

---

### Fooer4

Fooer4 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L35)  

```go
type Fooer4 interface {
	Foo() error
}
```

---

### Fooer5

Fooer5 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L45)  

```go
type Fooer5 interface {
	Foo()
	Bar
}
```

---

### Fooer6

Fooer6 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L51)  
**Added in:** v1.12.0

```go
type Fooer6 interface {
	Foo(a otherfixture.Spaceship)
}
```

---

### Fooer7

Fooer7 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L56)  
**Added in:** v1.12.0

```go
type Fooer7 interface {
	Foo(a *otherfixture.Spaceship)
}
```

---

### Fooer8

Fooer8 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L61)  
**Added in:** v1.12.0

```go
type Fooer8 interface {
	Foo(a map[string]otherfixture.Spaceship)
}
```

---

### Fooer9

Fooer9 is an interface used for tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/plugins/pluginrpc-gen/fixtures/foo.go#L66)  
**Added in:** v1.12.0

```go
type Fooer9 interface {
	Foo(a map[string]*otherfixture.Spaceship)
}
```

---


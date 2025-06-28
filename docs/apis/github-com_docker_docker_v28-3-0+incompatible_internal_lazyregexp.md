# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/lazyregexp

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:31:17 UTC

## Overview

Package lazyregexp is a thin wrapper over regexp, allowing the use of global
regexp variables without forcing them to be compiled at init.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### Regexp

Regexp is a wrapper around regexp.Regexp, where the underlying regexp will be
compiled the first time it is needed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L22)  

```go
type Regexp struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates a new lazy regexp, delaying the compiling work until it is first
needed. If the code is being run as part of tests, the regexp compiling will
happen immediately.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L83)  

```go
func New(str string) *Regexp
```

#### Methods

##### Regexp.FindAllString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L62)  

```go
func (r *Regexp) FindAllString(s string, n int) []string
```

##### Regexp.FindAllStringSubmatch

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L42)  

```go
func (r *Regexp) FindAllStringSubmatch(s string, n int) [][]string
```

##### Regexp.FindString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L58)  

```go
func (r *Regexp) FindString(s string) string
```

##### Regexp.FindStringSubmatch

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L46)  

```go
func (r *Regexp) FindStringSubmatch(s string) []string
```

##### Regexp.FindStringSubmatchIndex

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L50)  

```go
func (r *Regexp) FindStringSubmatchIndex(s string) []int
```

##### Regexp.FindSubmatch

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L38)  

```go
func (r *Regexp) FindSubmatch(s []byte) [][]byte
```

##### Regexp.MatchString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L66)  

```go
func (r *Regexp) MatchString(s string) bool
```

##### Regexp.ReplaceAllString

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L54)  

```go
func (r *Regexp) ReplaceAllString(src, repl string) string
```

##### Regexp.ReplaceAllStringFunc

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L70)  

```go
func (r *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
```

##### Regexp.SubexpNames

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/lazyregexp/lazyregexp.go#L74)  

```go
func (r *Regexp) SubexpNames() []string
```

---


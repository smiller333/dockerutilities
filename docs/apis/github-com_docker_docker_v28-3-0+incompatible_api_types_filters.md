# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/filters

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:00 UTC

## Overview

Package filters provides tools for encoding a mapping of keys to a set of
multiple values.


## Examples

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ToJSON

ToJSON returns the Args as a JSON encoded string

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L58)  

```go
func ToJSON(a Args) (string, error)
```

---

### ToParamWithVersion

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L71)  

```go
func ToParamWithVersion(version string, a Args) (string, error)
```

---

## Types

### Args

Args stores a mapping of keys to a set of multiple values.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L16)  

```go
type Args struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### FromJSON

FromJSON decodes a JSON encoded string into Args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L85)  

```go
func FromJSON(p string) (Args, error)
```

##### NewArgs

NewArgs returns a new Args populated with the initial args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L32)  

```go
func NewArgs(initialArgs ...KeyValuePair) Args
```

#### Methods

##### Args.Add

Add a new value to the set of values

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L127)  

```go
func (args Args) Add(key, value string)
```

##### Args.Clone

Clone returns a copy of args.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L297)  

```go
func (args Args) Clone() (newArgs Args)
```

##### Args.Contains

Contains returns true if the key exists in the mapping

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L265)  

```go
func (args Args) Contains(field string) bool
```

##### Args.Del

Del removes a value from the set

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L136)  

```go
func (args Args) Del(key, value string)
```

##### Args.ExactMatch

ExactMatch returns true if the source matches exactly one of the values.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L221)  

```go
func (args Args) ExactMatch(key, source string) bool
```

##### Args.FuzzyMatch

FuzzyMatch returns true if the source matches exactly one value,  or the
source has one of the values as a prefix.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L250)  

```go
func (args Args) FuzzyMatch(key, source string) bool
```

##### Args.Get

Get returns the list of values associated with the key

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L114)  

```go
func (args Args) Get(key string) []string
```

##### Args.GetBoolOrDefault

GetBoolOrDefault returns a boolean value of the key if the key is present
and is interpretable as a boolean value. Otherwise the default value is returned.
Error is not nil only if the filter values are not valid boolean or are conflicting.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L201)  

```go
func (args Args) GetBoolOrDefault(key string, defaultValue bool) (bool, error)
```

##### Args.Keys

Keys returns all the keys in list of Args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L41)  

```go
func (args Args) Keys() []string
```

##### Args.Len

Len returns the number of keys in the mapping

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L146)  

```go
func (args Args) Len() int
```

##### Args.MarshalJSON

MarshalJSON returns a JSON byte representation of the Args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L50)  

```go
func (args Args) MarshalJSON() ([]byte, error)
```

##### Args.Match

Match returns true if any of the values at key match the source string

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L180)  

```go
func (args Args) Match(field, source string) bool
```

##### Args.MatchKVList

MatchKVList returns true if all the pairs in sources exist as key=value
pairs in the mapping at key, or if there are no values at key.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L152)  

```go
func (args Args) MatchKVList(key string, sources map[string]string) bool
```

##### Args.UniqueExactMatch

UniqueExactMatch returns true if there is only one value and the source
matches exactly the value.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L234)  

```go
func (args Args) UniqueExactMatch(key, source string) bool
```

##### Args.UnmarshalJSON

UnmarshalJSON populates the Args from JSON encode bytes

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L109)  

```go
func (args Args) UnmarshalJSON(raw []byte) error
```

##### Args.Validate

Validate compared the set of accepted keys against the keys in the mapping.
An error is returned if any mapping keys are not in the accepted set.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L272)  

```go
func (args Args) Validate(accepted map[string]bool) error
```

##### Args.WalkValues

WalkValues iterates over the list of values for a key in the mapping and calls
op() for each value. If op returns an error the iteration stops and the
error is returned.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L284)  

```go
func (args Args) WalkValues(field string, op func(value string) error) error
```

---

### KeyValuePair

KeyValuePair are used to initialize a new Args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L21)  

```go
type KeyValuePair struct {
	Key   string
	Value string
}
```

#### Functions

##### Arg

Arg creates a new KeyValuePair for initializing Args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/filters/parse.go#L27)  

```go
func Arg(key, value string) KeyValuePair
```

---


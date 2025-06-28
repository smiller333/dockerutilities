# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/dockerfile

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:02:59 UTC

## Overview

Package dockerfile is the evaluation step in the Dockerfile parse/evaluate pipeline.

It incorporates a dispatch table based on the parser.Node values (see the
parser package for more information) that are yielded from the parser itself.
Calling newBuilder with the BuildOpts struct can be used to customize the
experience for execution purposes only. Parsing is controlled in the parser
package, and this division of responsibility should be respected.

Please see the jump table targets for the actual invocations, most of which
will call out to the functions in internals.go to deal with their tasks.

ONBUILD is a special case, which is covered in the onbuild() func in
dispatchers.go.

The evaluator uses the concept of "steps", which are usually each processable
line in the Dockerfile. Each step is numbered and certain actions are taken
before and after each step, such as creating an image ID and removing temporary
containers and images. Note that ONBUILD creates a kinda-sorta "sub run" which
includes its own set of steps (usually only one of them).


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### BuildFromConfig

BuildFromConfig builds directly from `changes`, treating it as if it were the contents of a Dockerfile
It will:
- Call parse.Parse() to get an AST root for the concatenated Dockerfile entries.
- Do build by calling builder.dispatch() to call all entries' handling routines

BuildFromConfig is used by the /commit endpoint, with the changes
coming from the query parameter of the same name.

TODO: Remove?

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/builder.go#L323)  

```go
func BuildFromConfig(ctx context.Context, config *container.Config, changes []string, os string) (*container.Config, error)
```

---

## Types

### BuildArgs

BuildArgs manages arguments used by the builder

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L26)  

```go
type BuildArgs struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewBuildArgs

NewBuildArgs creates a new BuildArgs type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L38)  

```go
func NewBuildArgs(argsFromOptions map[string]*string) *BuildArgs
```

#### Methods

##### BuildArgs.AddArg

AddArg adds a new arg that can be used by directives

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L99)  

```go
func (b *BuildArgs) AddArg(key string, value *string)
```

##### BuildArgs.AddMetaArg

AddMetaArg adds a new meta arg that can be used by FROM directives

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L94)  

```go
func (b *BuildArgs) AddMetaArg(key string, value *string)
```

##### BuildArgs.Clone

Clone returns a copy of the BuildArgs type

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L48)  

```go
func (b *BuildArgs) Clone() *BuildArgs
```

##### BuildArgs.FilterAllowed

FilterAllowed returns all allowed args without the filtered args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L137)  

```go
func (b *BuildArgs) FilterAllowed(filter []string) []string
```

##### BuildArgs.GetAllAllowed

GetAllAllowed returns a mapping with all the allowed args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L114)  

```go
func (b *BuildArgs) GetAllAllowed() map[string]string
```

##### BuildArgs.GetAllMeta

GetAllMeta returns a mapping with all the meta args

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L119)  

```go
func (b *BuildArgs) GetAllMeta() map[string]string
```

##### BuildArgs.IsReferencedOrNotBuiltin

IsReferencedOrNotBuiltin checks if the key is a built-in arg, or if it has been
referenced by the Dockerfile. Returns true if the arg is not a builtin or
if the builtin has been referenced in the Dockerfile.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L107)  

```go
func (b *BuildArgs) IsReferencedOrNotBuiltin(key string) bool
```

##### BuildArgs.MergeReferencedArgs

MergeReferencedArgs merges referenced args from another BuildArgs
object into the current one

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L64)  

```go
func (b *BuildArgs) MergeReferencedArgs(other *BuildArgs)
```

##### BuildArgs.ResetAllowed

ResetAllowed clears the list of args that are allowed to be used by a
directive

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L89)  

```go
func (b *BuildArgs) ResetAllowed()
```

##### BuildArgs.WarnOnUnusedBuildArgs

WarnOnUnusedBuildArgs checks if there are any leftover build-args that were
passed but not consumed during build. Print a warning, if there are any.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/buildargs.go#L72)  

```go
func (b *BuildArgs) WarnOnUnusedBuildArgs(out io.Writer)
```

---

### BuildManager

BuildManager is shared across all Builder objects

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/builder.go#L49)  
**Added in:** v1.11.0

```go
type BuildManager struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### NewBuildManager

NewBuildManager creates a BuildManager

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/builder.go#L56)  
**Added in:** v1.11.0

```go
func NewBuildManager(b builder.Backend, identityMapping user.IdentityMapping) (*BuildManager, error)
```

#### Methods

##### BuildManager.Build

Build starts a new build from a BuildConfig

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/builder.go#L66)  
**Added in:** v1.11.0

```go
func (bm *BuildManager) Build(ctx context.Context, config backend.BuildConfig) (*builder.Result, error)
```

---

### Builder

Builder is a Dockerfile builder
It implements the builder.Backend interface.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/builder.go#L111)  

```go
type Builder struct {
	Stdout io.Writer
	Stderr io.Writer
	Aux    *streamformatter.AuxFormatter
	Output io.Writer
	// contains filtered or unexported fields
}
```

---

### ImageProber

ImageProber exposes an Image cache to the Builder. It supports resetting a
cache.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/dockerfile/imageprobe.go#L14)  

```go
type ImageProber interface {
	Reset(ctx context.Context) error
	Probe(parentID string, runConfig *container.Config, platform ocispec.Platform) (string, error)
}
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/stack

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:15:07 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Dump

Dump outputs the runtime stack to os.StdErr.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/stack/stackdump.go#L17)  

```go
func Dump()
```

---

### DumpToFile

DumpToFile appends the runtime stack into a file named "goroutine-stacks-<timestamp>.log"
in dir and returns the full path to that file. If no directory name is
provided, it outputs to os.Stderr.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/stack/stackdump.go#L24)  

```go
func DumpToFile(dir string) (string, error)
```

---

## Types

This section is empty.


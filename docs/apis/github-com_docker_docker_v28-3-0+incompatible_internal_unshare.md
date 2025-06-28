# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/unshare

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:11 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Go

Go calls the given functions in a new goroutine, locked to an OS thread,
which has had the parts of its execution state disassociated from the rest of
the current process using unshare(2). It blocks until the new goroutine has
started and setupfn has returned. fn is only called if setupfn returns nil. A
nil setupfn or fn is equivalent to passing a no-op function.

The disassociated execution state and any changes made to it are only visible
to the goroutine which the functions are called in. Any other goroutines,
including ones started from the function, will see the same execution state
as the rest of the process.

The acceptable flags are documented in the unshare(2) Linux man-page.
The corresponding CLONE_* constants are defined in package unix.

This function may terminate the thread which the new goroutine executed on
after fn returns, which could cause subprocesses started with the
syscall.SysProcAttr Pdeathsig field set to be signaled before process
termination. Any subprocess started before this function is called may be
affected, in addition to any subprocesses started inside setupfn or fn.
There are more details at https://go.dev/issue/27505.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unshare/unshare_linux.go#L100)  

```go
func Go(flags int, setupfn func() error, fn func()) error
```

---

## Types

This section is empty.


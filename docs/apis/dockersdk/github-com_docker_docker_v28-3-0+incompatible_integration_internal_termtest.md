# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration/internal/termtest

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:06:30 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### StripANSICommands

StripANSICommands attempts to strip ANSI console escape and control sequences
from s, returning a string containing only the final printed characters which
would be visible onscreen if the string was to be processed by a terminal
emulator. Basic cursor positioning and screen erase control sequences are
parsed and processed such that the output of simple CLI commands passed
through a Windows Pseudoterminal and then this function yields the same
string as if the output of those commands was redirected to a file.

The only correct way to represent the result of processing ANSI console
output would be a two-dimensional array of an emulated terminal's display
buffer. That would be awkward to test against, so this function instead
attempts to render to a one-dimensional string without extra padding. This is
an inherently lossy process, and any attempts to render a string containing
many cursor positioning commands are unlikely to yield satisfactory results.
Handlers for several ANSI control sequences are also unimplemented; attempts
to parse a string containing one will panic.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration/internal/termtest/stripansi.go#L25)  

```go
func StripANSICommands(s string, opts ...ansiterm.Option) (string, error)
```

---

## Types

This section is empty.


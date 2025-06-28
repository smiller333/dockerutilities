# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/journald/internal/fake

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:14 UTC

## Overview

Package fake implements a journal writer for testing which is decoupled from
the system's journald.

The systemd project does not have any facilities to support testing of
journal reader clients (although it has been requested:
https://github.com/systemd/systemd/issues/14120) so we have to get creative.
The systemd-journal-remote command reads serialized journal entries in the
Journal Export Format and writes them to journal files. This format is
well-documented and straightforward to generate.


## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/journald/internal/fake/sender.go#L43)

```go
var ErrCommandNotFound = errors.New("systemd-journal-remote command not found")
```

## Functions

### JournalRemoteCmdPath

JournalRemoteCmdPath searches for the systemd-journal-remote command in
well-known paths and the directories named in the $PATH environment variable.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/journald/internal/fake/sender.go#L47)  

```go
func JournalRemoteCmdPath() (string, error)
```

---

## Types

### Sender

Sender fakes github.com/coreos/go-systemd/v22/journal.Send, writing journal
entries to an arbitrary journal file without depending on a running journald
process.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/journald/internal/fake/sender.go#L59)  

```go
type Sender struct {
	CmdName    string
	OutputPath string

	// Clock for timestamping sent messages.
	Clock clock.Clock
	// Whether to assign the event's realtime timestamp to the time
	// specified by the SYSLOG_TIMESTAMP variable value. This is roughly
	// analogous to journald receiving the event and assigning it a
	// timestamp in zero time after the SYSLOG_TIMESTAMP value was set,
	// which is highly unrealistic in practice.
	AssignEventTimestampFromSyslogTimestamp bool
	// Boot ID for journal entries. Required by systemd-journal-remote as of
	// https://github.com/systemd/systemd/commit/1eede158519e4e5ed22738c90cb57a91dbecb7f2
	// (systemd 255).
	BootID uuid.UUID

	// When set, Send will act as a test helper and redirect
	// systemd-journal-remote command output to the test log.
	TB testing.TB
}
```

#### Functions

##### New

New constructs a new Sender which will write journal entries to outpath. The
file name must end in '.journal' and the directory must already exist. The
journal file will be created if it does not exist. An existing journal file
will be appended to.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/journald/internal/fake/sender.go#L85)  

```go
func New(outpath string) (*Sender, error)
```

##### NewT

NewT is like New but will skip the test if the systemd-journal-remote command
is not available.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/journald/internal/fake/sender.go#L101)  

```go
func NewT(t *testing.T, outpath string) *Sender
```

#### Methods

##### Sender.Send

Send is a drop-in replacement for
github.com/coreos/go-systemd/v22/journal.Send.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/journald/internal/fake/sender.go#L116)  

```go
func (s *Sender) Send(message string, priority journal.Priority, vars map[string]string) error
```

---


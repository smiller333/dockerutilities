# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/journald/internal/export

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:08 UTC

## Overview

Package export implements a serializer for the systemd Journal Export Format
as documented at https://systemd.io/JOURNAL_EXPORT_FORMATS/


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### WriteEndOfEntry

WriteEndOfEntry terminates the journal entry.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/journald/internal/export/export.go#L47)  

```go
func WriteEndOfEntry(w io.Writer) error
```

---

### WriteField

WriteField writes the field serialized to Journal Export format to w.

The variable name must consist only of uppercase characters, numbers and
underscores. No validation or sanitization is performed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/journald/internal/export/export.go#L30)  

```go
func WriteField(w io.Writer, variable, value string) error
```

---

## Types

This section is empty.


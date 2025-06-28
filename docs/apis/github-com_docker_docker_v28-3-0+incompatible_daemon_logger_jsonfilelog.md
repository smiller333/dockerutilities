# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/jsonfilelog

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:29:19 UTC

## Overview

Package jsonfilelog provides the default Logger implementation for
Docker logging. This logger logs to files on the host server in the
JSON format.


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonfilelog.go#L21)

```go
const Name = "json-file"
```

## Variables

This section is empty.

## Functions

### New

New creates new JSONFileLogger which writes to filename passed in
on given context.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonfilelog.go#L48)  

```go
func New(info logger.Info) (logger.Logger, error)
```

---

### ValidateLogOpt

ValidateLogOpt looks for json specific log options max-file & max-size.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonfilelog.go#L155)  
**Added in:** v1.8.0

```go
func ValidateLogOpt(cfg map[string]string) error
```

---

## Types

### JSONFileLogger

JSONFileLogger is Logger implementation for default Docker logging.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonfilelog.go#L31)  

```go
type JSONFileLogger struct {
	// contains filtered or unexported fields
}
```

#### Methods

##### JSONFileLogger.Close

Close closes underlying file and signals all the readers
that the logs producer is gone.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonfilelog.go#L175)  

```go
func (l *JSONFileLogger) Close() error
```

##### JSONFileLogger.Log

Log converts logger.Message to jsonlog.JSONLog and serializes it to file.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonfilelog.go#L120)  

```go
func (l *JSONFileLogger) Log(msg *logger.Message) error
```

##### JSONFileLogger.Name

Name returns name of this logger.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/jsonfilelog.go#L180)  

```go
func (l *JSONFileLogger) Name() string
```

##### JSONFileLogger.ReadLogs

ReadLogs implements the logger's LogReader interface for the logs
created by this driver.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/jsonfilelog/read.go#L19)  
**Added in:** v1.8.0

```go
func (l *JSONFileLogger) ReadLogs(ctx context.Context, config logger.ReadConfig) *logger.LogWatcher
```

---


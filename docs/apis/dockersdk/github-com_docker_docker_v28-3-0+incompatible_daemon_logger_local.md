# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/local

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:05:20 UTC

## Overview

Package local provides a logger implementation that stores logs on disk.

Log messages are encoded as protobufs with a header and footer for each message.
The header and footer are big-endian binary encoded uint32 values which indicate the size of the log message.
The header and footer of each message allows you to efficiently read through a file either forwards or in
backwards (such as is the case when tailing a file)

Example log message format: [22][This is a log message.][22][28][This is another log message.][28]


## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/local/local.go#L20)

```go
const (
	// Name is the name of the driver
	Name = "local"
)
```

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/local/local.go#L39)

```go
var LogOptKeys = map[string]bool{
	"max-file": true,
	"max-size": true,
	"compress": true,
}
```

## Functions

### New

New creates a new local logger
You must provide the `LogPath` in the passed in info argument, this is the file path that logs are written to.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/local/local.go#L70)  

```go
func New(info logger.Info) (logger.Logger, error)
```

---

### ValidateLogOpt

ValidateLogOpt looks for log driver specific options.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/local/local.go#L46)  

```go
func ValidateLogOpt(cfg map[string]string) error
```

---

## Types

### CreateConfig

CreateConfig is used to configure new instances of driver

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/local/config.go#L8)  

```go
type CreateConfig struct {
	DisableCompression bool
	MaxFileSize        int64
	MaxFileCount       int
}
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/logger/awslogs

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:04:54 UTC

## Overview

Package awslogs provides the logdriver for forwarding container logs to Amazon CloudWatch Logs


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### New

New creates an awslogs logger using the configuration passed in on the
context.  Supported context configuration variables are awslogs-region,
awslogs-endpoint, awslogs-group, awslogs-stream, awslogs-create-group,
awslogs-multiline-pattern and awslogs-datetime-format.
When available, configuration is also taken from environment variables
AWS_REGION, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, the shared credentials
file (~/.aws/credentials), and the EC2 Instance Metadata Service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/awslogs/cloudwatchlogs.go#L143)  

```go
func New(info logger.Info) (logger.Logger, error)
```

---

### ValidateLogOpt

ValidateLogOpt looks for awslogs-specific log options awslogs-region, awslogs-endpoint
awslogs-group, awslogs-stream, awslogs-create-group, awslogs-datetime-format,
awslogs-multiline-pattern

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/logger/awslogs/cloudwatchlogs.go#L739)  

```go
func ValidateLogOpt(cfg map[string]string) error
```

---

## Types

This section is empty.


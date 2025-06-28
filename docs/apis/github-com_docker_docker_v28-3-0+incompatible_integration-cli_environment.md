# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration-cli/environment

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:58 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/environment/environment.go#L12)

```go
var DefaultClientBinary = os.Getenv("TEST_CLIENT_BINARY")
```

## Functions

This section is empty.

## Types

### Execution

Execution contains information about the current test execution and daemon
under test

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/environment/environment.go#L22)  

```go
type Execution struct {
	environment.Execution
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New returns details about the testing environment

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/environment/environment.go#L33)  

```go
func New(ctx context.Context) (*Execution, error)
```

#### Methods

##### Execution.DockerBinary

DockerBinary returns the docker binary for this testing environment

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/environment/environment.go#L28)  

```go
func (e *Execution) DockerBinary() string
```

---


# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration-cli/cli

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:30:50 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Args

Args build an icmd.Cmd struct from the specified (command and) arguments.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L142)  

```go
func Args(commandAndArgs ...string) icmd.Cmd
```

---

### BuildCmd

BuildCmd executes the specified docker build command and expect a success

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L35)  

```go
func BuildCmd(t testing.TB, name string, cmdOperators ...CmdOperator) *icmd.Result
```

---

### Daemon

Daemon points to the specified daemon

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L147)  

```go
func Daemon(d *daemon.Daemon) func(*icmd.Cmd) func()
```

---

### Docker

Docker executes the specified docker command

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L96)  

```go
func Docker(cmd icmd.Cmd, cmdOperators ...CmdOperator) *icmd.Result
```

---

### DockerCmd

DockerCmd executes the specified docker command and expect a success

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L29)  

```go
func DockerCmd(t testing.TB, args ...string) *icmd.Result
```

---

### Format

Format sets the specified format with --format flag

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L131)  

```go
func Format(format string) func(*icmd.Cmd) func()
```

---

### InDir

InDir sets the folder in which the command should be executed

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L182)  

```go
func InDir(path string) func(*icmd.Cmd) func()
```

---

### InspectCmd

InspectCmd executes the specified docker inspect command and expect a success

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L41)  

```go
func InspectCmd(t testing.TB, name string, cmdOperators ...CmdOperator) *icmd.Result
```

---

### SetTestEnvironment

SetTestEnvironment sets a static test environment
TODO: decouple this package from environment

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L21)  

```go
func SetTestEnvironment(env *environment.Execution)
```

---

### WaitExited

WaitExited will wait for the specified container to state exit, subject
to a maximum time limit in seconds supplied by the caller

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L54)  

```go
func WaitExited(t testing.TB, name string, timeout time.Duration, cmdOperators ...CmdOperator)
```

---

### WaitRun

WaitRun will wait for the specified container to be running, maximum 5 seconds.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L47)  

```go
func WaitRun(t testing.TB, name string, cmdOperators ...CmdOperator)
```

---

### WithEnvironmentVariables

WithEnvironmentVariables sets the specified environment variables for the command to run

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L163)  

```go
func WithEnvironmentVariables(envs ...string) func(cmd *icmd.Cmd) func()
```

---

### WithFlags

WithFlags sets the specified flags for the command to run

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L174)  

```go
func WithFlags(flags ...string) func(*icmd.Cmd) func()
```

---

### WithStdin

WithStdin sets the standard input reader for the command

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L198)  

```go
func WithStdin(stdin io.Reader) func(*icmd.Cmd) func()
```

---

### WithStdout

WithStdout sets the standard output writer of the command

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L190)  

```go
func WithStdout(writer io.Writer) func(*icmd.Cmd) func()
```

---

### WithTimeout

WithTimeout sets the timeout for the command to run

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L155)  

```go
func WithTimeout(timeout time.Duration) func(cmd *icmd.Cmd) func()
```

---

## Types

### CmdOperator

CmdOperator defines functions that can modify a command

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/cli/cli.go#L26)  

```go
type CmdOperator func(*icmd.Cmd) func()
```

---


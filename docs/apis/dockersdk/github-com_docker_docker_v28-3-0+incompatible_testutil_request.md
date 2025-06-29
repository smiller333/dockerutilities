# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/testutil/request

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:16:17 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### ContentType

ContentType sets the specified Content-Type request header

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/ops.go#L53)  

```go
func ContentType(contentType string) func(*Options)
```

---

### DaemonHost

DaemonHost return the daemon host string for this test execution

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L189)  

```go
func DaemonHost() string
```

---

### DaemonTime

DaemonTime provides the current time on the daemon host

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L36)  

```go
func DaemonTime(ctx context.Context, t testing.TB, client client.APIClient, testEnv *environment.Execution) time.Time
```

---

### DaemonUnixTime

DaemonUnixTime returns the current time on the daemon host with nanoseconds precision.
It return the time formatted how the client sends timestamps to the server.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L52)  

```go
func DaemonUnixTime(ctx context.Context, t testing.TB, client client.APIClient, testEnv *environment.Execution) string
```

---

### Delete

Delete creates and execute a DELETE request on the specified host and endpoint, with the specified request modifiers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L64)  

```go
func Delete(ctx context.Context, endpoint string, modifiers ...func(*Options)) (*http.Response, io.ReadCloser, error)
```

---

### Do

Do creates and execute a request on the specified endpoint, with the specified request modifiers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L79)  

```go
func Do(ctx context.Context, endpoint string, modifiers ...func(*Options)) (*http.Response, io.ReadCloser, error)
```

---

### Get

Get creates and execute a GET request on the specified host and endpoint, with the specified request modifiers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L69)  

```go
func Get(ctx context.Context, endpoint string, modifiers ...func(*Options)) (*http.Response, io.ReadCloser, error)
```

---

### Head

Head creates and execute a HEAD request on the specified host and endpoint, with the specified request modifiers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L74)  

```go
func Head(ctx context.Context, endpoint string, modifiers ...func(*Options)) (*http.Response, io.ReadCloser, error)
```

---

### Host

Host creates a modifier that sets the specified host as the request URL host

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/ops.go#L18)  

```go
func Host(host string) func(*Options)
```

---

### JSON

JSON sets the Content-Type request header to json

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/ops.go#L61)  

```go
func JSON(o *Options)
```

---

### JSONBody

JSONBody creates a modifier that encodes the specified data to a JSON string and set it as request body. It also sets
the Content-Type header of the request.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/ops.go#L67)  

```go
func JSONBody(data interface{}) func(*Options)
```

---

### Method

Method creates a modifier that sets the specified string as the request method

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/ops.go#L32)  

```go
func Method(method string) func(*Options)
```

---

### NewAPIClient

NewAPIClient returns a docker API client configured from environment variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L27)  

```go
func NewAPIClient(t testing.TB, ops ...client.Opt) client.APIClient
```

---

### Post

Post creates and execute a POST request on the specified host and endpoint, with the specified request modifiers

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L59)  

```go
func Post(ctx context.Context, endpoint string, modifiers ...func(*Options)) (*http.Response, io.ReadCloser, error)
```

---

### RawContent

RawContent sets the specified reader as body for the request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/ops.go#L45)  

```go
func RawContent(reader io.ReadCloser) func(*Options)
```

---

### RawString

RawString sets the specified string as body for the request

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/ops.go#L40)  

```go
func RawString(content string) func(*Options)
```

---

### ReadBody

ReadBody read the specified ReadCloser content and returns it

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L109)  

```go
func ReadBody(b io.ReadCloser) ([]byte, error)
```

---

### SockConn

SockConn opens a connection on the specified socket

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/request.go#L198)  

```go
func SockConn(timeout time.Duration, daemon string) (net.Conn, error)
```

---

### With

With adds a request modifier to the options

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/ops.go#L25)  

```go
func With(f func(*http.Request) error) func(*Options)
```

---

## Types

### Options

Options defines request options, like request modifiers and which host to target

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/testutil/request/ops.go#L12)  

```go
type Options struct {
	// contains filtered or unexported fields
}
```

---


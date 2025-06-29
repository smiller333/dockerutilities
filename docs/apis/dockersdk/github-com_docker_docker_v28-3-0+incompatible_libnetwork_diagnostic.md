# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/diagnostic

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:08:56 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### DebugHTTPForm

DebugHTTPForm helper to print the form url parameters

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L202)  

```go
func DebugHTTPForm(r *http.Request)
```

---

### HTTPReply

HTTPReply helper function that takes care of sending the message out

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L226)  

```go
func HTTPReply(w http.ResponseWriter, r *HTTPResult, j *JSONOutput) (int, error)
```

---

## Types

### ErrorCmd

ErrorCmd command with error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L67)  

```go
type ErrorCmd struct {
	Error string `json:"error"`
}
```

#### Methods

##### ErrorCmd.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L71)  

```go
func (e *ErrorCmd) String() string
```

---

### HTTPResult

HTTPResult Diagnostic Server HTTP result operation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L35)  

```go
type HTTPResult struct {
	Message string          `json:"message"`
	Details StringInterface `json:"details"`
}
```

#### Functions

##### CommandSucceed

CommandSucceed creates a success message

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L11)  

```go
func CommandSucceed(result StringInterface) *HTTPResult
```

##### FailCommand

FailCommand creates a failure message with error

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L19)  

```go
func FailCommand(err error) *HTTPResult
```

##### WrongCommand

WrongCommand creates a wrong command response

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L27)  

```go
func WrongCommand(message, usage string) *HTTPResult
```

#### Methods

##### HTTPResult.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L40)  

```go
func (h *HTTPResult) String() string
```

---

### JSONOutput

JSONOutput contains details on JSON output printing

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L209)  

```go
type JSONOutput struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### ParseHTTPFormOptions

ParseHTTPFormOptions easily parse the JSON printing options

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L215)  

```go
func ParseHTTPFormOptions(r *http.Request) (bool, *JSONOutput)
```

---

### NetworkStatsResult

NetworkStatsResult network db stats related to entries and queue len for a network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L125)  

```go
type NetworkStatsResult struct {
	Entries  int `json:"entries"`
	QueueLen int `jsoin:"qlen"`
}
```

#### Methods

##### NetworkStatsResult.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L130)  

```go
func (n *NetworkStatsResult) String() string
```

---

### PeerEntryObj

PeerEntryObj entry in the networkdb peer table

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L90)  

```go
type PeerEntryObj struct {
	Index int    `json:"-"`
	Name  string `json:"-=name"`
	IP    string `json:"ip"`
}
```

#### Methods

##### PeerEntryObj.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L96)  

```go
func (p *PeerEntryObj) String() string
```

---

### Server

Server when the debug is enabled exposes a
This data structure is protected by the Agent mutex so does not require and additional mutex here

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L20)  

```go
type Server struct {
	// contains filtered or unexported fields
}
```

#### Functions

##### New

New creates a new diagnostic server

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L30)  

```go
func New() *Server
```

#### Methods

##### Server.Enable

Enable opens a TCP socket to debug the passed network DB

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L74)  

```go
func (s *Server) Enable(ip string, port int)
```

##### Server.Enabled

Enabled returns true when the debug is enabled

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L125)  

```go
func (s *Server) Enabled() bool
```

##### Server.Handle

Handle registers the handler for the given pattern,
replacing any existing handler.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L44)  

```go
func (s *Server) Handle(pattern string, handler http.Handler)
```

##### Server.HandleFunc

HandleFunc registers the handler function for the given pattern,
replacing any existing handler.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L63)  

```go
func (s *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
```

##### Server.ServeHTTP

ServeHTTP this is the method called bu the ListenAndServe, and is needed to allow us to
use our custom mux

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L69)  

```go
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

##### Server.Shutdown

Shutdown stop the debug and closes the tcp socket

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/server.go#L108)  

```go
func (s *Server) Shutdown()
```

---

### StringCmd

StringCmd command with info string

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L58)  

```go
type StringCmd struct {
	Info string `json:"info"`
}
```

#### Methods

##### StringCmd.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L62)  

```go
func (s *StringCmd) String() string
```

---

### StringInterface

StringInterface interface that has to be implemented by messages

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L6)  

```go
type StringInterface interface {
	String() string
}
```

---

### TableEndpointsResult

TableEndpointsResult fully typed message for proper unmarshaling on the client side

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L113)  

```go
type TableEndpointsResult struct {
	TableObj
	Elements []TableEntryObj `json:"entries"`
}
```

---

### TableEntryObj

TableEntryObj network db table entry object

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L101)  

```go
type TableEntryObj struct {
	Index int    `json:"-"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Owner string `json:"owner"`
}
```

#### Methods

##### TableEntryObj.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L108)  

```go
func (t *TableEntryObj) String() string
```

---

### TableObj

TableObj network db table object

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L76)  

```go
type TableObj struct {
	Length   int               `json:"size"`
	Elements []StringInterface `json:"entries"`
}
```

#### Methods

##### TableObj.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L81)  

```go
func (t *TableObj) String() string
```

---

### TablePeersResult

TablePeersResult fully typed message for proper unmarshaling on the client side

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L119)  

```go
type TablePeersResult struct {
	TableObj
	Elements []PeerEntryObj `json:"entries"`
}
```

---

### UsageCmd

UsageCmd command with usage field

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L49)  

```go
type UsageCmd struct {
	Usage string `json:"usage"`
}
```

#### Methods

##### UsageCmd.String

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/diagnostic/types.go#L53)  

```go
func (u *UsageCmd) String() string
```

---


# Docker Client Package

The `dockerclient` package provides a configurable wrapper around the Docker SDK client, enhancing the standard Docker client with additional functionality and consistent error handling.

## Features

- **Flexible Configuration**: Support for custom timeouts, hosts, API versions, and HTTP clients
- **Enhanced Error Handling**: Comprehensive error handling with descriptive messages
- **Connection Management**: Built-in connection testing and timeout management
- **Image Operations**: Pull, push, build, inspect, and save Docker images
- **Container Operations**: Create, copy from, and remove containers
- **System Information**: Retrieve Docker daemon and version information

## Usage

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    	"github.com/smiller333/dockerutilities/src/dockerclient"
)

func main() {
    // Create a default client
    client, err := dockerclient.NewDefaultClient()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    
    // Test connection
    ctx := context.Background()
    if err := client.Ping(ctx); err != nil {
        log.Fatal("Cannot connect to Docker daemon:", err)
    }
    
    // Get Docker version information
    version, err := client.GetVersion(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Docker version: %s\n", version.Version)
    fmt.Printf("API version: %s\n", version.APIVersion)
}
```

### Custom Configuration

```go
config := &dockerclient.Config{
    Host:       "unix:///var/run/docker.sock",
    APIVersion: "1.41", 
    Timeout:    30 * time.Second,
}

client, err := dockerclient.NewDockerClient(config)
if err != nil {
    log.Fatal(err)
}
defer client.Close()
```

### Image Operations

```go
// Pull an image
pullReader, err := client.PullImage(ctx, "nginx:latest", nil)
if err != nil {
    log.Fatal(err)
}
defer pullReader.Close()

// Build an image from Dockerfile
buildReader, err := client.BuildImage(ctx, "/path/to/Dockerfile", "myapp:latest")
if err != nil {
    log.Fatal(err)
}
defer buildReader.Close()

// Inspect an image
imageInfo, err := client.InspectImage(ctx, "nginx:latest")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Image ID: %s\n", imageInfo.ID)

// Save an image to tar archive
saveReader, err := client.SaveImage(ctx, []string{"nginx:latest"})
if err != nil {
    log.Fatal(err)
}
defer saveReader.Close()
```

### Container Operations

```go
// Create a container
containerConfig := &container.Config{
    Image: "nginx:latest",
    Cmd:   []string{"nginx", "-g", "daemon off;"},
}

resp, err := client.CreateContainer(ctx, containerConfig, nil, nil, nil, "my-nginx")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Container created: %s\n", resp.ID)

// Copy files from container
reader, stat, err := client.CopyFromContainer(ctx, resp.ID, "/etc/nginx")
if err != nil {
    log.Fatal(err)
}
defer reader.Close()

// Remove container
err = client.RemoveContainer(ctx, resp.ID, true)
if err != nil {
    log.Fatal(err)
}
```

### Direct API Access

For operations not covered by the wrapper, access the underlying Docker client:

```go
apiClient := client.GetClient()
// Use apiClient for any Docker SDK operations
```

## API Reference

### Types

#### Config
Configuration options for the Docker client wrapper.

```go
type Config struct {
    Host       string        // Docker daemon host (e.g., "unix:///var/run/docker.sock")
    APIVersion string        // Docker API version to use
    HTTPClient *http.Client  // Custom HTTP client configuration
    Timeout    time.Duration // Default timeout for operations
    TLSVerify  bool         // Enable TLS verification
    CertPath   string       // Path to TLS certificates
}
```

#### DockerClient
Main wrapper struct that provides enhanced Docker client functionality.

```go
type DockerClient struct {
    // Contains unexported fields
}
```

### Methods

#### Client Creation and Management
- `NewDefaultClient() (*DockerClient, error)` - Creates a client with default configuration
- `NewDockerClient(config *Config) (*DockerClient, error)` - Creates a client with custom configuration  
- `Close() error` - Closes the client connection
- `GetClient() client.APIClient` - Gets underlying Docker client for direct API access

#### Connection and System Information
- `Ping(ctx context.Context) error` - Tests connection to Docker daemon
- `IsConnected(ctx context.Context) bool` - Checks if connected to Docker daemon
- `GetInfo(ctx context.Context) (*system.Info, error)` - Gets detailed daemon information
- `GetVersion(ctx context.Context) (*types.Version, error)` - Gets daemon version information

#### Image Operations
- `PullImage(ctx context.Context, imageName string, authConfig *registry.AuthConfig) (io.ReadCloser, error)` - Pulls an image from registry
- `PushImage(ctx context.Context, imageName string, authConfig *registry.AuthConfig) (io.ReadCloser, error)` - Pushes an image to registry
- `BuildImage(ctx context.Context, dockerfilePath string, tag string) (io.ReadCloser, error)` - Builds an image from Dockerfile
- `InspectImage(ctx context.Context, nameOrID string) (*image.InspectResponse, error)` - Inspects an image
- `SaveImage(ctx context.Context, imageNames []string) (io.ReadCloser, error)` - Saves images to tar archive

#### Container Operations
- `CreateContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *ocispec.Platform, containerName string) (*container.CreateResponse, error)` - Creates a new container
- `CopyFromContainer(ctx context.Context, containerID, srcPath string) (io.ReadCloser, container.PathStat, error)` - Copies files from container
- `RemoveContainer(ctx context.Context, containerID string, force bool) error` - Removes a container

#### Configuration Management
- `GetTimeout() time.Duration` - Gets current timeout
- `SetTimeout(timeout time.Duration)` - Sets new timeout
- `GetConfig() *Config` - Gets client configuration

## Error Handling

The package provides comprehensive error handling:

- Connection errors are wrapped with descriptive messages
- All methods return errors following Go conventions
- Context timeouts are handled gracefully
- Image not found operations return `ErrImageNotFound` for consistent error detection

### Standard Errors

```go
var (
    // ErrImageNotFound is returned when an image is not found in the registry
    ErrImageNotFound = errors.New("image not found")
)
```

## Testing

The package includes comprehensive unit tests that can be run in two modes:

```bash
# Run unit tests only (no Docker daemon required)
go test -short ./src/dockerclient

# Run all tests including integration tests (requires Docker daemon)
go test ./src/dockerclient
```

Integration tests are skipped in short mode and require a running Docker daemon for full functionality testing.

## Implementation Notes

- All methods accept a `context.Context` parameter for timeout and cancellation support
- If `nil` context is passed, a default timeout context is created automatically
- The wrapper maintains the underlying Docker client for direct API access when needed
- Configuration is immutable after client creation except for timeout adjustments
- The client properly closes resources and connections when `Close()` is called

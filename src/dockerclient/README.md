# Docker Client Package

The `dockerclient` package provides a flexible wrapper around the Docker SDK client, enhancing the standard Docker client with additional functionality and configuration options.

## Features

- **Flexible Configuration**: Support for custom timeouts, hosts, API versions, and HTTP clients
- **Enhanced Error Handling**: Comprehensive error handling with descriptive messages
- **Simplified API**: High-level functions for common Docker operations
- **Connection Management**: Built-in connection testing and management
- **Container & Image Operations**: Easy listing, filtering, and inspection of containers and images
- **System Information**: Retrieve Docker daemon and system information

## Usage

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/smiller333/dockerutils/src/dockerclient"
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
    
    // Get system information
    info, err := client.GetSystemSummary(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Docker version: %s\n", info.ServerVersion)
    fmt.Printf("Running containers: %d\n", info.ContainersRunning)
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

### Listing Containers

```go
// Get all containers
containers, err := client.GetContainerSummaries(ctx, true)
if err != nil {
    log.Fatal(err)
}

for _, container := range containers {
    fmt.Printf("Container: %s (%s)\n", container.Name, container.State)
}

// Get only running containers
runningContainers, err := client.GetRunningContainers(ctx)
if err != nil {
    log.Fatal(err)
}
```

### Listing Images

```go
images, err := client.GetImageSummaries(ctx, false)
if err != nil {
    log.Fatal(err)
}

for _, image := range images {
    fmt.Printf("Image: %s (Size: %d bytes)\n", image.Tags[0], image.Size)
}
```

### Finding Specific Resources

```go
// Find container by name
container, err := client.FindContainerByName(ctx, "my-container")
if err != nil {
    log.Printf("Container not found: %v", err)
} else {
    fmt.Printf("Found container: %s\n", container.ID)
}

// Find image by tag
image, err := client.FindImageByTag(ctx, "nginx:latest")
if err != nil {
    log.Printf("Image not found: %v", err)
} else {
    fmt.Printf("Found image: %s\n", image.ID)
}

// Get containers using specific image
containers, err := client.GetContainersByImage(ctx, "nginx:latest")
if err != nil {
    log.Fatal(err)
}
```

### Direct API Access

For operations not covered by the wrapper, you can access the underlying Docker client:

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
    Host       string        // Docker daemon host
    APIVersion string        // Docker API version
    HTTPClient *http.Client  // Custom HTTP client
    Timeout    time.Duration // Default timeout for operations
    TLSVerify  bool         // Enable TLS verification
    CertPath   string       // Path to TLS certificates
}
```

#### ContainerSummary
Simplified view of container information.

```go
type ContainerSummary struct {
    ID      string    // Short container ID
    Name    string    // Container name
    Image   string    // Image name
    Status  string    // Container status
    State   string    // Container state
    Ports   []string  // Port mappings
    Created int64     // Creation timestamp
}
```

#### ImageSummary
Simplified view of image information.

```go
type ImageSummary struct {
    ID       string   // Short image ID
    Tags     []string // Image tags
    Size     int64    // Image size in bytes
    Created  int64    // Creation timestamp
    RepoTags []string // Repository tags
}
```

#### SystemInfo
Simplified view of Docker daemon information.

```go
type SystemInfo struct {
    ContainersRunning int    // Number of running containers
    ContainersPaused  int    // Number of paused containers
    ContainersStopped int    // Number of stopped containers
    Images            int    // Number of images
    ServerVersion     string // Docker server version
    DockerRootDir     string // Docker root directory
    OperatingSystem   string // Host operating system
    Architecture      string // Host architecture
    MemTotal          int64  // Total memory in bytes
    NCPU              int    // Number of CPUs
}
```

### Methods

#### Client Creation
- `NewDefaultClient() (*DockerClient, error)` - Creates a client with default configuration
- `NewDockerClient(config *Config) (*DockerClient, error)` - Creates a client with custom configuration

#### Connection Management
- `Ping(ctx context.Context) error` - Tests connection to Docker daemon
- `IsConnected(ctx context.Context) bool` - Checks if connected to Docker daemon
- `Close() error` - Closes the client connection

#### System Information
- `GetSystemSummary(ctx context.Context) (*SystemInfo, error)` - Gets system summary
- `GetInfo(ctx context.Context) (*system.Info, error)` - Gets detailed daemon info
- `GetVersion(ctx context.Context) (*types.Version, error)` - Gets daemon version

#### Container Operations
- `ListContainers(ctx context.Context, all bool, filterArgs ...string) ([]types.Container, error)`
- `GetContainerSummaries(ctx context.Context, all bool) ([]ContainerSummary, error)`
- `GetRunningContainers(ctx context.Context) ([]ContainerSummary, error)`
- `FindContainerByName(ctx context.Context, name string) (*types.Container, error)`
- `GetContainersByImage(ctx context.Context, imageName string) ([]ContainerSummary, error)`
- `ContainerExists(ctx context.Context, nameOrID string) (bool, error)`
- `GetContainerInfo(ctx context.Context, nameOrID string) (*types.ContainerJSON, error)`

#### Image Operations
- `ListImages(ctx context.Context, all bool, filterArgs ...string) ([]image.Summary, error)`
- `GetImageSummaries(ctx context.Context, all bool) ([]ImageSummary, error)`
- `FindImageByTag(ctx context.Context, tag string) (*image.Summary, error)`
- `ImageExists(ctx context.Context, nameOrID string) (bool, error)`
- `GetImageInfo(ctx context.Context, nameOrID string) (*types.ImageInspect, error)`

#### Configuration
- `GetTimeout() time.Duration` - Gets current timeout
- `SetTimeout(timeout time.Duration)` - Sets new timeout
- `GetConfig() *Config` - Gets client configuration
- `GetClient() client.APIClient` - Gets underlying Docker client

## Error Handling

The package provides comprehensive error handling:

- Connection errors are wrapped with descriptive messages
- Not found errors are properly detected using `client.IsErrNotFound()`
- All methods return errors following Go conventions
- Context timeouts are handled gracefully

## Testing

The package includes comprehensive tests:

```bash
# Run unit tests only
go test -short ./src/dockerclient

# Run all tests (requires Docker daemon)
go test ./src/dockerclient
```

Integration tests are skipped in short mode and require a running Docker daemon.

## Examples

See the `cmd/docker.go` file for a complete example of using the Docker client in a CLI application.

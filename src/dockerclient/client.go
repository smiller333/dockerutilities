// Package dockerclient provides a flexible wrapper around the Docker SDK client.
// It enhances the standard Docker client with additional functionality and configuration options.
package dockerclient

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
)

// Config holds configuration options for the Docker client wrapper
type Config struct {
	// Host specifies the Docker daemon host (e.g., "unix:///var/run/docker.sock")
	Host string
	// APIVersion specifies the Docker API version to use
	APIVersion string
	// HTTPClient allows custom HTTP client configuration
	HTTPClient *http.Client
	// Timeout sets the default timeout for operations
	Timeout time.Duration
	// TLSVerify enables TLS verification
	TLSVerify bool
	// CertPath specifies the path to TLS certificates
	CertPath string
}

// DockerClient wraps the standard Docker client with additional functionality
type DockerClient struct {
	client  client.APIClient
	config  *Config
	timeout time.Duration
}

// NewDockerClient creates a new Docker client wrapper with the provided configuration
func NewDockerClient(config *Config) (*DockerClient, error) {
	if config == nil {
		config = &Config{
			Timeout: 30 * time.Second,
		}
	}

	// Set default timeout if not specified
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	var opts []client.Opt

	// Configure from environment if no explicit host is set
	if config.Host == "" {
		opts = append(opts, client.FromEnv)
	} else {
		opts = append(opts, client.WithHost(config.Host))
	}

	// Set API version if specified
	if config.APIVersion != "" {
		opts = append(opts, client.WithVersion(config.APIVersion))
	} else {
		opts = append(opts, client.WithAPIVersionNegotiation())
	}

	// Use custom HTTP client if provided
	if config.HTTPClient != nil {
		opts = append(opts, client.WithHTTPClient(config.HTTPClient))
	}

	// Create the Docker client
	cli, err := client.NewClientWithOpts(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}

	return &DockerClient{
		client:  cli,
		config:  config,
		timeout: config.Timeout,
	}, nil
}

// NewDefaultClient creates a Docker client with default configuration
func NewDefaultClient() (*DockerClient, error) {
	return NewDockerClient(nil)
}

// Close closes the underlying Docker client connection
func (dc *DockerClient) Close() error {
	return dc.client.Close()
}

// GetClient returns the underlying Docker API client for direct access
func (dc *DockerClient) GetClient() client.APIClient {
	return dc.client
}

// Ping pings the Docker daemon to verify connectivity
func (dc *DockerClient) Ping(ctx context.Context) error {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	_, err := dc.client.Ping(ctx)
	return err
}

// GetInfo retrieves information about the Docker daemon
func (dc *DockerClient) GetInfo(ctx context.Context) (*system.Info, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	info, err := dc.client.Info(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Docker daemon info: %w", err)
	}

	return &info, nil
}

// GetVersion returns version information about the Docker daemon
func (dc *DockerClient) GetVersion(ctx context.Context) (*types.Version, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	version, err := dc.client.ServerVersion(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Docker daemon version: %w", err)
	}

	return &version, nil
}

// ListContainers lists containers with optional filtering
func (dc *DockerClient) ListContainers(ctx context.Context, all bool, filterArgs ...string) ([]types.Container, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	// Build filters
	filterSet := filters.NewArgs()
	for i := 0; i < len(filterArgs); i += 2 {
		if i+1 < len(filterArgs) {
			filterSet.Add(filterArgs[i], filterArgs[i+1])
		}
	}

	options := container.ListOptions{
		All:     all,
		Filters: filterSet,
	}

	containers, err := dc.client.ContainerList(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	return containers, nil
}

// ListImages lists images with optional filtering
func (dc *DockerClient) ListImages(ctx context.Context, all bool, filterArgs ...string) ([]image.Summary, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	// Build filters
	filterSet := filters.NewArgs()
	for i := 0; i < len(filterArgs); i += 2 {
		if i+1 < len(filterArgs) {
			filterSet.Add(filterArgs[i], filterArgs[i+1])
		}
	}

	options := image.ListOptions{
		All:     all,
		Filters: filterSet,
	}

	images, err := dc.client.ImageList(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("failed to list images: %w", err)
	}

	return images, nil
}

// ContainerExists checks if a container exists by name or ID
func (dc *DockerClient) ContainerExists(ctx context.Context, nameOrID string) (bool, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	_, err := dc.client.ContainerInspect(ctx, nameOrID)
	if err != nil {
		if client.IsErrNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check container existence: %w", err)
	}

	return true, nil
}

// ImageExists checks if an image exists by name or ID
func (dc *DockerClient) ImageExists(ctx context.Context, nameOrID string) (bool, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	_, _, err := dc.client.ImageInspectWithRaw(ctx, nameOrID)
	if err != nil {
		if client.IsErrNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check image existence: %w", err)
	}

	return true, nil
}

// GetContainerInfo retrieves detailed information about a container
func (dc *DockerClient) GetContainerInfo(ctx context.Context, nameOrID string) (*types.ContainerJSON, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	info, err := dc.client.ContainerInspect(ctx, nameOrID)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect container %s: %w", nameOrID, err)
	}

	return &info, nil
}

// GetImageInfo retrieves detailed information about an image
func (dc *DockerClient) GetImageInfo(ctx context.Context, nameOrID string) (*types.ImageInspect, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	info, _, err := dc.client.ImageInspectWithRaw(ctx, nameOrID)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect image %s: %w", nameOrID, err)
	}

	return &info, nil
}

// IsConnected checks if the client is connected to the Docker daemon
func (dc *DockerClient) IsConnected(ctx context.Context) bool {
	return dc.Ping(ctx) == nil
}

// GetTimeout returns the configured timeout for operations
func (dc *DockerClient) GetTimeout() time.Duration {
	return dc.timeout
}

// SetTimeout updates the default timeout for operations
func (dc *DockerClient) SetTimeout(timeout time.Duration) {
	dc.timeout = timeout
}

// GetConfig returns the client configuration
func (dc *DockerClient) GetConfig() *Config {
	return dc.config
}

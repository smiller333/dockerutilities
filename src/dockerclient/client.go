// Package dockerclient provides a flexible wrapper around the Docker SDK client.
// It enhances the standard Docker client with additional functionality and configuration options.
// Copyright (c) 2025 Docker Utils Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

package dockerclient

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/build"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// Standard errors returned by the Docker client wrapper
var (
	// ErrImageNotFound is returned when an image is not found in the registry
	ErrImageNotFound = errors.New("image not found")
)

var trustedRegistries = []string{
	"docker.io",
	"gcr.io",
	"quay.io",
	"registry.k8s.io",
	"ghcr.io",
}

// ValidateDockerAccess checks Docker socket access and warns about risks
func ValidateDockerAccess() error {
	var socketPath string

	if runtime.GOOS == "windows" {
		socketPath = `\\.\pipe\docker_engine`
	} else {
		socketPath = "/var/run/docker.sock"
	}

	// Check if socket exists and is accessible
	if _, err := os.Stat(socketPath); err != nil {
		return fmt.Errorf("Docker socket not accessible at %s: %w", socketPath, err)
	}

	// Log security warning
	log.Printf("⚠️  SECURITY WARNING: This application requires Docker socket access")
	log.Printf("   Socket: %s", socketPath)
	log.Printf("   Risk: Significant system privileges - only analyze trusted Docker images")
	log.Printf("   Recommendation: Only use with images from trusted sources")

	return nil
}

// ValidateImageSource checks if an image comes from a trusted source
func ValidateImageSource(imageName string) (bool, string) {
	// Parse image name to extract registry
	parts := strings.Split(imageName, "/")
	var registry string

	if len(parts) > 1 && strings.Contains(parts[0], ".") {
		registry = parts[0]
	} else {
		registry = "docker.io" // Default registry
	}

	// Check against trusted registries
	for _, trusted := range trustedRegistries {
		if registry == trusted {
			return true, "Trusted registry"
		}
	}

	return false, fmt.Sprintf("Untrusted registry: %s", registry)
}

// WarnUntrustedImage displays warnings for untrusted images
func WarnUntrustedImage(imageName string) {
	trusted, reason := ValidateImageSource(imageName)

	if !trusted {
		log.Printf("🚨 SECURITY WARNING: Analyzing potentially untrusted image")
		log.Printf("   Image: %s", imageName)
		log.Printf("   Reason: %s", reason)
		log.Printf("   Risk: Malicious images could exploit the analysis process")
		log.Printf("   Recommendation: Only proceed if you trust this image source")
	}
}

// ValidateImageName validates Docker image name format
func ValidateImageName(imageName string) error {
	if imageName == "" {
		return fmt.Errorf("image name cannot be empty")
	}

	if len(imageName) > 255 {
		return fmt.Errorf("image name too long (max 255 characters)")
	}

	// Basic validation for Docker image name format
	imageNameRegex := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*(/[a-zA-Z0-9][a-zA-Z0-9_.-]*)*(:[\w][\w.-]{0,127})?$`)
	if !imageNameRegex.MatchString(imageName) {
		return fmt.Errorf("invalid image name format: %s", imageName)
	}

	return nil
}

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

// PullImage pulls an image from a registry
func (dc *DockerClient) PullImage(ctx context.Context, imageName string, authConfig *registry.AuthConfig) (io.ReadCloser, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	// Create the options for the image pull
	options := image.PullOptions{}

	// Encode the auth config to base64 if provided
	if authConfig != nil {
		authBytes, err := json.Marshal(authConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal auth config: %w", err)
		}
		options.RegistryAuth = base64.URLEncoding.EncodeToString(authBytes)
	}

	// Pull the image
	reader, err := dc.client.ImagePull(ctx, imageName, options)
	if err != nil {
		// Check if the error is due to image not found (manifest not found)
		if strings.Contains(err.Error(), "manifest for") && strings.Contains(err.Error(), "not found") {
			return nil, ErrImageNotFound
		}
		return nil, fmt.Errorf("failed to pull image %s: %w", imageName, err)
	}

	return reader, nil
}

// PushImage pushes an image to a registry
func (dc *DockerClient) PushImage(ctx context.Context, imageName string, authConfig *registry.AuthConfig) (io.ReadCloser, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	// Create the options for the image push
	options := image.PushOptions{}

	// Encode the auth config to base64 if provided
	if authConfig != nil {
		authBytes, err := json.Marshal(authConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal auth config: %w", err)
		}
		options.RegistryAuth = base64.URLEncoding.EncodeToString(authBytes)
	}

	// Push the image
	reader, err := dc.client.ImagePush(ctx, imageName, options)
	if err != nil {
		return nil, fmt.Errorf("failed to push image %s: %w", imageName, err)
	}

	return reader, nil
}

// BuildImage builds a Docker image from a single Dockerfile
func (dc *DockerClient) BuildImage(ctx context.Context, dockerfilePath string, tag string) (io.ReadCloser, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	// Read the Dockerfile content
	dockerfileContent, err := os.ReadFile(dockerfilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read Dockerfile %s: %w", dockerfilePath, err)
	}

	// Create a tar archive with just the Dockerfile
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	// Add the Dockerfile to the archive
	header := &tar.Header{
		Name:     "Dockerfile",
		Mode:     0644,
		Size:     int64(len(dockerfileContent)),
		ModTime:  time.Now(),
		Typeflag: tar.TypeReg,
	}

	if err := tw.WriteHeader(header); err != nil {
		return nil, fmt.Errorf("failed to write Dockerfile header: %w", err)
	}

	if _, err := tw.Write(dockerfileContent); err != nil {
		return nil, fmt.Errorf("failed to write Dockerfile content: %w", err)
	}

	if err := tw.Close(); err != nil {
		return nil, fmt.Errorf("failed to close tar writer: %w", err)
	}

	// Create build options
	options := build.ImageBuildOptions{
		Tags:           []string{tag},
		Remove:         true,
		SuppressOutput: false,
		NoCache:        false,
		PullParent:     true,
	}

	// Build the image
	response, err := dc.client.ImageBuild(ctx, &buf, options)
	if err != nil {
		return nil, fmt.Errorf("failed to build image with tag %s: %w", tag, err)
	}

	return response.Body, nil
}

// InspectImage retrieves detailed information about an image
func (dc *DockerClient) InspectImage(ctx context.Context, nameOrID string) (*image.InspectResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	info, err := dc.client.ImageInspect(ctx, nameOrID)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect image %s: %w", nameOrID, err)
	}

	return &info, nil
}

// SaveImage saves one or more images to a tar archive
func (dc *DockerClient) SaveImage(ctx context.Context, imageNames []string) (io.ReadCloser, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	if len(imageNames) == 0 {
		return nil, fmt.Errorf("at least one image name must be provided")
	}

	// Save the images to a tar archive
	reader, err := dc.client.ImageSave(ctx, imageNames)
	if err != nil {
		return nil, fmt.Errorf("failed to save images %v: %w", imageNames, err)
	}

	return reader, nil
}

// CreateContainer creates a new container based on the given configuration.
// It returns the container ID and any warnings from the creation process.
func (dc *DockerClient) CreateContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *ocispec.Platform, containerName string) (*container.CreateResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	// Create the container
	resp, err := dc.client.ContainerCreate(ctx, config, hostConfig, networkingConfig, platform, containerName)
	if err != nil {
		return nil, fmt.Errorf("failed to create container %s: %w", containerName, err)
	}

	return &resp, nil
}

// CopyFromContainer gets the content from the container and returns it as a Reader
// for a TAR archive to manipulate it in the host. It's up to the caller to close the reader.
// Returns the tar archive reader, path statistics, and any error that occurred.
func (dc *DockerClient) CopyFromContainer(ctx context.Context, containerID, srcPath string) (io.ReadCloser, container.PathStat, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	// Copy content from the container
	reader, stat, err := dc.client.CopyFromContainer(ctx, containerID, srcPath)
	if err != nil {
		return nil, container.PathStat{}, fmt.Errorf("failed to copy from container %s (path: %s): %w", containerID, srcPath, err)
	}

	return reader, stat, nil
}

// RemoveContainer removes a container by ID or name
func (dc *DockerClient) RemoveContainer(ctx context.Context, containerID string, force bool) error {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	options := container.RemoveOptions{
		RemoveVolumes: true,
		Force:         force,
	}

	err := dc.client.ContainerRemove(ctx, containerID, options)
	if err != nil {
		return fmt.Errorf("failed to remove container %s: %w", containerID, err)
	}

	return nil
}

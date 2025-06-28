// Package dockerclient provides additional utility functions for common Docker operations.
package dockerclient

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
)

// ContainerSummary provides a simplified view of container information
type ContainerSummary struct {
	ID      string
	Name    string
	Image   string
	Status  string
	State   string
	Ports   []string
	Created int64
}

// ImageSummary provides a simplified view of image information
type ImageSummary struct {
	ID       string
	Tags     []string
	Size     int64
	Created  int64
	RepoTags []string
}

// SystemInfo provides a simplified view of Docker daemon information
type SystemInfo struct {
	ContainersRunning int
	ContainersPaused  int
	ContainersStopped int
	Images            int
	ServerVersion     string
	DockerRootDir     string
	OperatingSystem   string
	Architecture      string
	MemTotal          int64
	NCPU              int
}

// GetSystemSummary retrieves a summary of the Docker system status
func (dc *DockerClient) GetSystemSummary(ctx context.Context) (*SystemInfo, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	info, err := dc.GetInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get system info: %w", err)
	}

	version, err := dc.GetVersion(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get version info: %w", err)
	}

	return &SystemInfo{
		ContainersRunning: info.ContainersRunning,
		ContainersPaused:  info.ContainersPaused,
		ContainersStopped: info.ContainersStopped,
		Images:            info.Images,
		ServerVersion:     version.Version,
		DockerRootDir:     info.DockerRootDir,
		OperatingSystem:   info.OperatingSystem,
		Architecture:      info.Architecture,
		MemTotal:          info.MemTotal,
		NCPU:              info.NCPU,
	}, nil
}

// GetContainerSummaries retrieves simplified container information
func (dc *DockerClient) GetContainerSummaries(ctx context.Context, all bool) ([]ContainerSummary, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	containers, err := dc.ListContainers(ctx, all)
	if err != nil {
		return nil, err
	}

	summaries := make([]ContainerSummary, len(containers))
	for i, c := range containers {
		name := "unknown"
		if len(c.Names) > 0 {
			// Remove leading slash from container name
			name = strings.TrimPrefix(c.Names[0], "/")
		}

		// Format ports
		ports := make([]string, len(c.Ports))
		for j, port := range c.Ports {
			if port.PublicPort > 0 {
				ports[j] = fmt.Sprintf("%d:%d/%s", port.PublicPort, port.PrivatePort, port.Type)
			} else {
				ports[j] = fmt.Sprintf("%d/%s", port.PrivatePort, port.Type)
			}
		}

		summaries[i] = ContainerSummary{
			ID:      c.ID[:12], // Show short ID
			Name:    name,
			Image:   c.Image,
			Status:  c.Status,
			State:   c.State,
			Ports:   ports,
			Created: c.Created,
		}
	}

	return summaries, nil
}

// GetImageSummaries retrieves simplified image information
func (dc *DockerClient) GetImageSummaries(ctx context.Context, all bool) ([]ImageSummary, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	images, err := dc.ListImages(ctx, all)
	if err != nil {
		return nil, err
	}

	summaries := make([]ImageSummary, len(images))
	for i, img := range images {
		summaries[i] = ImageSummary{
			ID:       img.ID[:12], // Show short ID
			Tags:     img.RepoTags,
			Size:     img.Size,
			Created:  img.Created,
			RepoTags: img.RepoTags,
		}
	}

	return summaries, nil
}

// GetRunningContainers retrieves only running containers
func (dc *DockerClient) GetRunningContainers(ctx context.Context) ([]ContainerSummary, error) {
	return dc.GetContainerSummaries(ctx, false)
}

// FindContainerByName finds a container by its name
func (dc *DockerClient) FindContainerByName(ctx context.Context, name string) (*types.Container, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	containers, err := dc.ListContainers(ctx, true)
	if err != nil {
		return nil, err
	}

	for _, c := range containers {
		for _, containerName := range c.Names {
			// Remove leading slash and compare
			if strings.TrimPrefix(containerName, "/") == name {
				return &c, nil
			}
		}
	}

	return nil, fmt.Errorf("container not found: %s", name)
}

// FindImageByTag finds an image by its tag
func (dc *DockerClient) FindImageByTag(ctx context.Context, tag string) (*image.Summary, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	images, err := dc.ListImages(ctx, true)
	if err != nil {
		return nil, err
	}

	for _, img := range images {
		for _, repoTag := range img.RepoTags {
			if repoTag == tag {
				return &img, nil
			}
		}
	}

	return nil, fmt.Errorf("image not found: %s", tag)
}

// GetContainersByImage finds all containers using a specific image
func (dc *DockerClient) GetContainersByImage(ctx context.Context, imageName string) ([]ContainerSummary, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), dc.timeout)
		defer cancel()
	}

	// Use filter to find containers by image
	containers, err := dc.ListContainers(ctx, true, "ancestor", imageName)
	if err != nil {
		return nil, err
	}

	summaries := make([]ContainerSummary, len(containers))
	for i, c := range containers {
		name := "unknown"
		if len(c.Names) > 0 {
			name = strings.TrimPrefix(c.Names[0], "/")
		}

		ports := make([]string, len(c.Ports))
		for j, port := range c.Ports {
			if port.PublicPort > 0 {
				ports[j] = fmt.Sprintf("%d:%d/%s", port.PublicPort, port.PrivatePort, port.Type)
			} else {
				ports[j] = fmt.Sprintf("%d/%s", port.PrivatePort, port.Type)
			}
		}

		summaries[i] = ContainerSummary{
			ID:      c.ID[:12],
			Name:    name,
			Image:   c.Image,
			Status:  c.Status,
			State:   c.State,
			Ports:   ports,
			Created: c.Created,
		}
	}

	return summaries, nil
}

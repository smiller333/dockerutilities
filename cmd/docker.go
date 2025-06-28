// Package cmd provides command-line interface functionality for dockerutils.
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/smiller333/dockerutils/src/dockerclient"
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker client operations",
	Long: `Docker client operations using the enhanced Docker SDK wrapper.

This command provides various operations for interacting with the Docker daemon.`,
}

// statusCmd represents the docker status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show Docker daemon status",
	Long:  `Display information about the Docker daemon and system.`,
	RunE:  showDockerStatus,
}

// containersCmd represents the docker containers command
var containersCmd = &cobra.Command{
	Use:   "containers",
	Short: "List Docker containers",
	Long:  `List Docker containers with optional filtering.`,
	RunE:  listContainers,
}

// imagesCmd represents the docker images command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List Docker images",
	Long:  `List Docker images with optional filtering.`,
	RunE:  listImages,
}

var (
	// Flags
	showAll     bool
	onlyRunning bool
)

func init() {
	// Add docker command to root
	rootCmd.AddCommand(dockerCmd)

	// Add subcommands to docker command
	dockerCmd.AddCommand(statusCmd)
	dockerCmd.AddCommand(containersCmd)
	dockerCmd.AddCommand(imagesCmd)

	// Add flags
	containersCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all containers (default shows just running)")
	containersCmd.Flags().BoolVar(&onlyRunning, "running", false, "Show only running containers")

	imagesCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all images including intermediate images")
}

// showDockerStatus displays Docker daemon status information
func showDockerStatus(cmd *cobra.Command, args []string) error {
	client, err := dockerclient.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check connectivity
	fmt.Println("Checking Docker daemon connectivity...")
	if err := client.Ping(ctx); err != nil {
		return fmt.Errorf("failed to connect to Docker daemon: %w", err)
	}
	fmt.Println("✓ Connected to Docker daemon")

	// Get system summary
	summary, err := client.GetSystemSummary(ctx)
	if err != nil {
		return fmt.Errorf("failed to get system summary: %w", err)
	}

	// Display information
	fmt.Printf("\nDocker System Information:\n")
	fmt.Printf("Server Version: %s\n", summary.ServerVersion)
	fmt.Printf("Operating System: %s\n", summary.OperatingSystem)
	fmt.Printf("Architecture: %s\n", summary.Architecture)
	fmt.Printf("CPUs: %d\n", summary.NCPU)
	fmt.Printf("Total Memory: %.2f GB\n", float64(summary.MemTotal)/(1024*1024*1024))
	fmt.Printf("Docker Root Directory: %s\n", summary.DockerRootDir)

	fmt.Printf("\nContainer Statistics:\n")
	fmt.Printf("Running: %d\n", summary.ContainersRunning)
	fmt.Printf("Paused: %d\n", summary.ContainersPaused)
	fmt.Printf("Stopped: %d\n", summary.ContainersStopped)
	fmt.Printf("Total Images: %d\n", summary.Images)

	return nil
}

// listContainers lists Docker containers
func listContainers(cmd *cobra.Command, args []string) error {
	client, err := dockerclient.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Determine which containers to show
	all := showAll && !onlyRunning

	containers, err := client.GetContainerSummaries(ctx, all)
	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	if len(containers) == 0 {
		if all {
			fmt.Println("No containers found.")
		} else {
			fmt.Println("No running containers found.")
		}
		return nil
	}

	// Display header
	fmt.Printf("%-12s %-20s %-30s %-10s %-20s %s\n",
		"CONTAINER ID", "NAME", "IMAGE", "STATE", "STATUS", "PORTS")
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")

	// Display containers
	for _, container := range containers {
		ports := ""
		if len(container.Ports) > 0 {
			ports = container.Ports[0]
			if len(container.Ports) > 1 {
				ports += fmt.Sprintf(" (+%d more)", len(container.Ports)-1)
			}
		}

		fmt.Printf("%-12s %-20s %-30s %-10s %-20s %s\n",
			container.ID,
			container.Name,
			truncateString(container.Image, 30),
			container.State,
			truncateString(container.Status, 20),
			ports,
		)
	}

	fmt.Printf("\nTotal: %d containers\n", len(containers))
	return nil
}

// listImages lists Docker images
func listImages(cmd *cobra.Command, args []string) error {
	client, err := dockerclient.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	images, err := client.GetImageSummaries(ctx, showAll)
	if err != nil {
		return fmt.Errorf("failed to list images: %w", err)
	}

	if len(images) == 0 {
		fmt.Println("No images found.")
		return nil
	}

	// Display header
	fmt.Printf("%-12s %-50s %-10s %s\n",
		"IMAGE ID", "REPOSITORY:TAG", "SIZE", "CREATED")
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────")

	// Display images
	for _, image := range images {
		tag := "<none>:<none>"
		if len(image.RepoTags) > 0 && image.RepoTags[0] != "<none>:<none>" {
			tag = image.RepoTags[0]
		}

		size := formatSize(image.Size)
		created := formatTime(image.Created)

		fmt.Printf("%-12s %-50s %-10s %s\n",
			image.ID,
			truncateString(tag, 50),
			size,
			created,
		)
	}

	fmt.Printf("\nTotal: %d images\n", len(images))
	return nil
}

// Helper functions

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	now := time.Now()

	if now.Sub(t) < 24*time.Hour {
		return t.Format("15:04")
	} else if now.Sub(t) < 7*24*time.Hour {
		return t.Format("Mon 15:04")
	} else {
		return t.Format("2006-01-02")
	}
}

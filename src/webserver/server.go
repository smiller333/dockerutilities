// Package webserver provides a web server for viewing Docker image analysis results.
package webserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/smiller333/dockerutils/src/analyzer"
)

// Config holds configuration options for the web server
type Config struct {
	Host    string // Host/IP address to bind to
	Port    string // Port number to listen on
	WebRoot string // Custom web root directory (optional)
}

// Server represents the web server instance
type Server struct {
	config     *Config
	httpServer *http.Server
	webRoot    string
}

// New creates a new web server instance with the given configuration
func New(config *Config) (*Server, error) {
	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	if config.Host == "" {
		config.Host = "localhost"
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	// Determine web root directory
	webRoot := config.WebRoot
	if webRoot == "" {
		// Use embedded/default web root
		webRoot = getDefaultWebRoot()
	}

	// Validate web root exists
	if _, err := os.Stat(webRoot); os.IsNotExist(err) {
		return nil, fmt.Errorf("web root directory does not exist: %s", webRoot)
	}

	server := &Server{
		config:  config,
		webRoot: webRoot,
	}

	return server, nil
}

// Start starts the web server and begins listening for requests
func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Register routes
	s.registerRoutes(mux)

	// Create HTTP server
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.config.Host, s.config.Port),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("Web server starting on http://%s:%s\n", s.config.Host, s.config.Port)
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the web server
func (s *Server) Shutdown() error {
	if s.httpServer == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Shutting down web server...")
	return s.httpServer.Shutdown(ctx)
}

// registerRoutes sets up all HTTP routes for the server
func (s *Server) registerRoutes(mux *http.ServeMux) {
	// Static file serving for the web UI
	mux.Handle("/", http.FileServer(http.Dir(s.webRoot)))

	// API routes
	mux.HandleFunc("/api/summaries", s.handleGetSummaries)
	mux.HandleFunc("/api/summaries/", s.handleGetSummary)
	mux.HandleFunc("/api/health", s.handleHealth)
}

// handleHealth returns server health status
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"version":   "1.0.0",
	}

	json.NewEncoder(w).Encode(response)
}

// handleGetSummaries returns a list of all available image summaries
func (s *Server) handleGetSummaries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Look for summary files in the tmp directory
	summaries, err := s.findSummaryFiles()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find summary files: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summaries)
}

// handleGetSummary returns a specific image summary by ID
func (s *Server) handleGetSummary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract summary ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/summaries/")
	if path == "" {
		http.Error(w, "Summary ID required", http.StatusBadRequest)
		return
	}

	// Find and return the specific summary
	summary, err := s.getSummaryByID(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get summary: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

// findSummaryFiles searches for summary JSON files in the tmp directory
func (s *Server) findSummaryFiles() ([]analyzer.ImageSummary, error) {
	summaries := []analyzer.ImageSummary{}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: %w", err)
	}

	tmpDir := filepath.Join(cwd, "tmp")

	// Check if tmp directory exists
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		return summaries, nil // Return empty list if no tmp directory
	}

	// Walk through tmp directory to find summary files
	err = filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Look for files matching the pattern summary.*.json
		if !info.IsDir() && strings.HasPrefix(info.Name(), "summary.") && strings.HasSuffix(info.Name(), ".json") {
			summary, err := s.parseSummaryFile(path)
			if err != nil {
				// Log error but continue processing other files
				fmt.Printf("Warning: failed to parse summary file %s: %v\n", path, err)
				return nil
			}

			summaries = append(summaries, summary)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk tmp directory: %w", err)
	}

	return summaries, nil
}

// parseSummaryFile reads and parses a summary JSON file
func (s *Server) parseSummaryFile(filePath string) (analyzer.ImageSummary, error) {
	var summary analyzer.ImageSummary

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return summary, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse JSON directly into ImageSummary struct
	err = json.Unmarshal(data, &summary)
	if err != nil {
		return summary, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return summary, nil
}

// getSummaryByID retrieves a specific summary by its ID
func (s *Server) getSummaryByID(id string) (analyzer.ImageSummary, error) {
	var summary analyzer.ImageSummary

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return summary, fmt.Errorf("failed to get current working directory: %w", err)
	}

	tmpDir := filepath.Join(cwd, "tmp")

	// Look for the specific summary file
	var summaryPath string
	err = filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileName := info.Name()
		if !info.IsDir() && strings.HasPrefix(fileName, "summary.") && strings.HasSuffix(fileName, ".json") {
			// Extract ID from filename
			fileID := strings.TrimSuffix(strings.TrimPrefix(fileName, "summary."), ".json")
			if fileID == id {
				summaryPath = path
				return filepath.SkipDir // Stop walking
			}
		}

		return nil
	})

	if err != nil {
		return summary, fmt.Errorf("failed to search for summary file: %w", err)
	}

	if summaryPath == "" {
		return summary, fmt.Errorf("summary with ID %s not found", id)
	}

	// Use the parseSummaryFile method to read and parse the summary
	summary, err = s.parseSummaryFile(summaryPath)
	if err != nil {
		return summary, fmt.Errorf("failed to parse summary file: %w", err)
	}

	return summary, nil
}

// getDefaultWebRoot returns the default web root directory
func getDefaultWebRoot() string {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		// Fallback to current directory
		return "."
	}

	// Check if web directory exists in the project
	webDir := filepath.Join(cwd, "web")
	if _, err := os.Stat(webDir); err == nil {
		return webDir
	}

	// Fallback to current directory
	return "."
}

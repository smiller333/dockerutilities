// Package webserver provides a web server for viewing Docker image analysis results.
package webserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

// AnalyzeRequest represents the request body for analyzing an image
type AnalyzeRequest struct {
	ImageName     string `json:"image_name"`
	KeepTempFiles bool   `json:"keep_temp_files,omitempty"`
	ForcePull     bool   `json:"force_pull,omitempty"`
}

// AnalyzeResponse represents the response for a successful image analysis
type AnalyzeResponse struct {
	Success bool   `json:"success"`
	ImageID string `json:"image_id,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// responseWriter wraps http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

// WriteHeader captures the status code before writing it
func (rw *responseWriter) WriteHeader(statusCode int) {
	if !rw.written {
		rw.statusCode = statusCode
		rw.written = true
		rw.ResponseWriter.WriteHeader(statusCode)
	}
}

// Write ensures WriteHeader is called with 200 if not already called
func (rw *responseWriter) Write(data []byte) (int, error) {
	if !rw.written {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(data)
}

// loggingMiddleware logs all HTTP requests with method, path, response code, and duration
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap the response writer to capture status code
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Default to 200 if WriteHeader is never called
			written:        false,
		}

		// Call the next handler
		next.ServeHTTP(wrapped, r)

		// Calculate duration and log the request
		duration := time.Since(start)
		log.Printf("%s %s %d %v", r.Method, r.URL.Path, wrapped.statusCode, duration)
	})
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
	mux.Handle("/", loggingMiddleware(http.FileServer(http.Dir(s.webRoot))))

	// Add ability to serve the Chrome DevTools JSON file locally
	mux.Handle("GET /.well-known/appspecific/com.chrome.devtools.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/.well-known/appspecific/com.chrome.devtools.json")
	}))

	// API routes - wrap each handler with logging middleware
	mux.Handle("GET /api/summaries", loggingMiddleware(http.HandlerFunc(s.handleGetSummaries)))
	mux.Handle("GET /api/summaries/", loggingMiddleware(http.HandlerFunc(s.handleGetSummary)))
	mux.Handle("DELETE /api/summaries/", loggingMiddleware(http.HandlerFunc(s.handleDeleteSummary)))
	mux.Handle("GET /api/health", loggingMiddleware(http.HandlerFunc(s.handleHealth)))
	mux.Handle("POST /api/analyze", loggingMiddleware(http.HandlerFunc(s.handleAnalyzeImage)))
}

// handleHealth returns server health status
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
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

// handleDeleteSummary handles DELETE requests to remove a specific image summary
func (s *Server) handleDeleteSummary(w http.ResponseWriter, r *http.Request) {
	// Extract summary ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/summaries/")
	if path == "" {
		http.Error(w, "Summary ID required", http.StatusBadRequest)
		return
	}

	// Delete the summary file
	err := s.deleteSummaryByID(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete summary: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Summary %s deleted successfully", path),
	}
	json.NewEncoder(w).Encode(response)
}

// handleAnalyzeImage handles POST requests to analyze a Docker image
func (s *Server) handleAnalyzeImage(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req AnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := AnalyzeResponse{
			Success: false,
			Error:   "Invalid request body",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate image name
	if req.ImageName == "" {
		response := AnalyzeResponse{
			Success: false,
			Error:   "Image name is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if image has already been analyzed
	existingSummaries, err := s.findSummaryFiles()
	if err == nil {
		for _, summary := range existingSummaries {
			if strings.EqualFold(summary.ImageTag, req.ImageName) {
				// Extract short image ID for response
				shortImageID := strings.Replace(summary.ImageID, "sha256:", "", 1)
				if len(shortImageID) > 12 {
					shortImageID = shortImageID[:12]
				}

				response := AnalyzeResponse{
					Success: true,
					ImageID: shortImageID,
					Message: "Image has already been analyzed",
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
				return
			}
		}
	}

	// Perform the analysis using the existing AnalyzeImage function
	result, err := analyzer.AnalyzeImage(req.ImageName, req.KeepTempFiles, req.ForcePull)
	if err != nil {
		response := AnalyzeResponse{
			Success: false,
			Error:   fmt.Sprintf("Analysis failed: %v", err),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Extract short image ID for response (first 12 characters without sha256: prefix)
	imageID := strings.TrimPrefix(result.ImageID, "sha256:")
	if len(imageID) > 12 {
		imageID = imageID[:12]
	}

	// Return success response with image ID
	response := AnalyzeResponse{
		Success: true,
		ImageID: imageID,
		Message: "Image analysis completed successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

// deleteSummaryByID removes a specific summary file by its ID
func (s *Server) deleteSummaryByID(id string) error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	tmpDir := filepath.Join(cwd, "tmp")

	// Look for the specific summary file
	var summaryPath string
	var imageFolder string

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
				imageFolder = filepath.Dir(path)
				return filepath.SkipDir // Stop walking
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to search for summary file: %w", err)
	}

	if summaryPath == "" {
		return fmt.Errorf("summary with ID %s not found", id)
	}

	// Remove the summary file
	if err := os.Remove(summaryPath); err != nil {
		return fmt.Errorf("failed to delete summary file: %w", err)
	}

	// Check if the image folder is empty after removing the summary file
	// If so, remove the entire folder
	entries, err := os.ReadDir(imageFolder)
	if err != nil {
		// Log warning but don't fail the operation
		fmt.Printf("Warning: failed to read directory %s: %v\n", imageFolder, err)
		return nil
	}

	// If the folder is empty, remove it
	if len(entries) == 0 {
		if err := os.Remove(imageFolder); err != nil {
			// Log warning but don't fail the operation
			fmt.Printf("Warning: failed to remove empty directory %s: %v\n", imageFolder, err)
		}
	}

	return nil
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

// Package webserver provides a web server for viewing Docker image analysis results.
package webserver

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/smiller333/dockerutils/src/analyzer"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed webpages/*
var staticFS embed.FS

const (
	// summaryFileName is the name of the file that stores all image summaries
	summaryFileName = "summaries.json"
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

// ImageSummary represents the minimum fields needed for image list
// This is done to reduce the size of the JSON data when listing images
type ImageSummary struct {
	ImageID      string `json:"image_id"` // ID of the Docker image
	ImageTag     string `json:"image_tag"`
	ImageSource  string `json:"image_source,omitempty"` // Source registry for non-DockerHub images
	ImageSize    int64  `json:"image_size"`             // Size in bytes
	Architecture string `json:"architecture"`
	AnalyzedAt   string `json:"analyzed_at"`          // Timestamp when analysis was performed
	Status       string `json:"status"`               // Status: "completed", "analyzing", "failed"
	RequestID    string `json:"request_id,omitempty"` // Request ID for tracking async operations
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

// AsyncAnalyzeRequest represents the request body for analyzing an image asynchronously
type AsyncAnalyzeRequest struct {
	ImageName     string `json:"image_name"`
	KeepTempFiles bool   `json:"keep_temp_files,omitempty"`
	ForcePull     bool   `json:"force_pull,omitempty"`
}

// AsyncAnalyzeResponse represents the response for an async image analysis request
type AsyncAnalyzeResponse struct {
	Success   bool   `json:"success"`
	RequestID string `json:"request_id,omitempty"`
	Message   string `json:"message,omitempty"`
	Error     string `json:"error,omitempty"`
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

	// If a web root was specified, validate it exists
	if config.WebRoot != "" {
		// Validate web root exists
		if _, err := os.Stat(config.WebRoot); os.IsNotExist(err) {
			return nil, fmt.Errorf("web root directory does not exist: %s", config.WebRoot)
		}
	}

	server := &Server{
		config:  config,
		webRoot: config.WebRoot,
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

	// If we are not using embedded files, write a log message
	if s.webRoot != "" {
		fmt.Printf("Development - Serving pages from web root: %s\n", s.webRoot)
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
	// Server static files for the server UI.  When using the embedded
	// web files, they are compiled into the binary.  Otherwise, they
	// are served from the web root directory.
	if s.webRoot == "" {
		var staticFS = fs.FS(staticFS)
		pages, err := fs.Sub(staticFS, "webpages")
		if err != nil {
			log.Fatalf("Failed to create subdirectory for static files: %v", err)
		}
		mux.Handle("/", loggingMiddleware(http.FileServer(http.FS(pages))))
	} else {
		mux.Handle("/", loggingMiddleware(http.FileServer(http.Dir(s.webRoot))))
	}

	// API routes - wrap each handler with logging middleware
	mux.Handle("GET /api/summaries", loggingMiddleware(http.HandlerFunc(s.handleGetSummaries)))
	mux.Handle("GET /api/info/", loggingMiddleware(http.HandlerFunc(s.handleGetInfo)))
	mux.Handle("DELETE /api/info/", loggingMiddleware(http.HandlerFunc(s.handleDeleteInfo)))
	mux.Handle("GET /api/health", loggingMiddleware(http.HandlerFunc(s.handleHealth)))
	mux.Handle("POST /api/analyze", loggingMiddleware(http.HandlerFunc(s.handleAnalyzeImage)))
	mux.Handle("POST /api/analyze-async", loggingMiddleware(http.HandlerFunc(s.handleAnalyzeImageAsync)))
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
// TODO: Rework this...
func (s *Server) handleGetSummaries(w http.ResponseWriter, r *http.Request) {
	// Look for summary files in the tmp directory
	summaries, err := s.findSummaries()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find summary files: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summaries)
}

// handleGetInfo returns a specific image info by ID
func (s *Server) handleGetInfo(w http.ResponseWriter, r *http.Request) {
	// Extract info ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/info/")
	if path == "" {
		http.Error(w, "Info ID required", http.StatusBadRequest)
		return
	}

	// Find and return the specific info
	info, err := s.getInfoByID(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get info: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// handleDeleteInfo handles DELETE requests to remove a specific image info
func (s *Server) handleDeleteInfo(w http.ResponseWriter, r *http.Request) {
	// Extract info ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/info/")
	if path == "" {
		http.Error(w, "Info ID required", http.StatusBadRequest)
		return
	}

	// Delete the info file
	err := s.deleteInfoByID(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete info: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Info %s deleted successfully", path),
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
	existingInfos, err := s.findSummaries()
	if err == nil {
		for _, info := range existingInfos {
			if strings.EqualFold(info.ImageTag, req.ImageName) {
				// Extract short image ID for response
				shortImageID := strings.Replace(info.ImageID, "sha256:", "", 1)
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
			// Error:   fmt.Sprintf("Analysis failed: %v", err),
			Error: cases.Title(language.English).String(err.Error()), // Title case error message
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

	// Try to read the generated info file and add it to the summary file
	infoPath := filepath.Join(result.ExtractedPath, fmt.Sprintf("info.%s.json", imageID))
	if _, err := os.Stat(infoPath); err == nil {
		// Read the info file
		data, err := os.ReadFile(infoPath)
		if err == nil {
			var info analyzer.ImageInfo
			if json.Unmarshal(data, &info) == nil {
				// Convert to summary and add to summary file
				summary := s.imageInfoToSummary(info)
				if err := s.addSummaryToFile(summary); err != nil {
					fmt.Printf("Warning: failed to add summary to file: %v\n", err)
				}
			}
		}
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

// handleAnalyzeImageAsync handles POST requests to analyze a Docker image asynchronously
func (s *Server) handleAnalyzeImageAsync(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req AsyncAnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := AsyncAnalyzeResponse{
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
		response := AsyncAnalyzeResponse{
			Success: false,
			Error:   "Image name is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if image has already been analyzed or is currently being analyzed
	existingInfos, err := s.findSummaries()
	if err == nil {
		for _, info := range existingInfos {
			if strings.EqualFold(info.ImageTag, req.ImageName) {
				// If analysis is completed, return the existing result
				if info.Status == "completed" {
					shortImageID := strings.Replace(info.ImageID, "sha256:", "", 1)
					if len(shortImageID) > 12 {
						shortImageID = shortImageID[:12]
					}

					response := AsyncAnalyzeResponse{
						Success:   true,
						RequestID: shortImageID,
						Message:   "Image has already been analyzed",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(response)
					return
				}

				// If analysis is currently in progress, return the existing request ID
				if info.Status == "analyzing" {
					response := AsyncAnalyzeResponse{
						Success:   true,
						RequestID: info.RequestID,
						Message:   "Image analysis is already in progress",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(response)
					return
				}
			}
		}
	}

	// Generate a unique request ID for tracking
	requestID := fmt.Sprintf("req_%d", time.Now().Unix())

	// Add a pending entry to the summaries file
	pendingSummary := ImageSummary{
		ImageID:      "", // Will be filled once analysis completes
		ImageTag:     req.ImageName,
		ImageSource:  "", // Will be filled once analysis completes
		ImageSize:    0,  // Will be filled once analysis completes
		Architecture: "", // Will be filled once analysis completes
		AnalyzedAt:   time.Now().UTC().Format(time.RFC3339),
		Status:       "analyzing",
		RequestID:    requestID,
	}

	if err := s.addSummaryToFile(pendingSummary); err != nil {
		log.Printf("Warning: failed to add pending summary to file: %v", err)
		// Continue with the analysis even if we can't track it in the summary file
	}

	// Start analysis in a goroutine
	go func() {
		log.Printf("Starting async analysis for image: %s (Request ID: %s)", req.ImageName, requestID)

		result, err := analyzer.AnalyzeImage(req.ImageName, req.KeepTempFiles, req.ForcePull)
		if err != nil {
			log.Printf("Async analysis failed for image %s (Request ID: %s): %v", req.ImageName, requestID, err)

			// Update the summary to show failed status
			failedSummary := ImageSummary{
				ImageID:      "",
				ImageTag:     req.ImageName,
				ImageSource:  "",
				ImageSize:    0,
				Architecture: "",
				AnalyzedAt:   time.Now().UTC().Format(time.RFC3339),
				Status:       "failed",
				RequestID:    requestID,
			}
			if updateErr := s.updateSummaryByRequestID(requestID, failedSummary); updateErr != nil {
				log.Printf("Warning: failed to update failed summary for request %s: %v", requestID, updateErr)
			}
			return
		}

		// Extract short image ID
		imageID := strings.TrimPrefix(result.ImageID, "sha256:")
		if len(imageID) > 12 {
			imageID = imageID[:12]
		}

		// Try to read the generated info file and update the summary
		infoPath := filepath.Join(result.ExtractedPath, fmt.Sprintf("info.%s.json", imageID))
		if _, err := os.Stat(infoPath); err == nil {
			// Read the info file
			data, err := os.ReadFile(infoPath)
			if err == nil {
				var info analyzer.ImageInfo
				if json.Unmarshal(data, &info) == nil {
					// Convert to summary and update the entry
					completedSummary := s.imageInfoToSummary(info)
					completedSummary.Status = "completed"
					completedSummary.RequestID = requestID

					if err := s.updateSummaryByRequestID(requestID, completedSummary); err != nil {
						log.Printf("Warning: failed to update completed summary for request %s: %v", requestID, err)
					}
				}
			}
		}

		log.Printf("Async analysis completed for image %s (Request ID: %s, Image ID: %s)", req.ImageName, requestID, imageID)
	}()

	// Return immediate response with request ID
	response := AsyncAnalyzeResponse{
		Success:   true,
		RequestID: requestID,
		Message:   "Image analysis started successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// findSummaries reads the summaries from the summary file or rebuilds it if missing/corrupt
func (s *Server) findSummaries() ([]ImageSummary, error) {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: %w", err)
	}

	tmpDir := filepath.Join(cwd, "tmp")
	summaryFilePath := filepath.Join(tmpDir, summaryFileName)

	// Check if tmp directory exists
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		return []ImageSummary{}, nil // Return empty list if no tmp directory
	}

	// Try to read existing summary file
	summaries, err := s.readSummaryFile(summaryFilePath)
	if err == nil {
		return summaries, nil
	}

	// If summary file doesn't exist or is corrupt, rebuild it
	fmt.Printf("Summary file not found or corrupt, rebuilding: %v\n", err)
	return s.rebuildSummaryFile(tmpDir, summaryFilePath)
}

// readSummaryFile reads and parses the summary file containing all ImageSummary objects
func (s *Server) readSummaryFile(filePath string) ([]ImageSummary, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("summary file does not exist")
	}

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read summary file: %w", err)
	}

	// Parse JSON into slice of ImageSummary
	var summaries []ImageSummary
	err = json.Unmarshal(data, &summaries)
	if err != nil {
		return nil, fmt.Errorf("failed to parse summary file JSON: %w", err)
	}

	return summaries, nil
}

// writeSummaryFile writes the slice of ImageSummary objects to the summary file
func (s *Server) writeSummaryFile(filePath string, summaries []ImageSummary) error {
	// Ensure the directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Marshal summaries to JSON with pretty formatting
	data, err := json.MarshalIndent(summaries, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal summaries to JSON: %w", err)
	}

	// Write to file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write summary file: %w", err)
	}

	return nil
}

// rebuildSummaryFile scans the tmp directory and rebuilds the summary file
func (s *Server) rebuildSummaryFile(tmpDir, summaryFilePath string) ([]ImageSummary, error) {
	summaries := []ImageSummary{}

	// Walk through tmp directory to find info files
	err := filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the summary file itself
		if info.Name() == summaryFileName {
			return nil
		}

		// Look for files matching the pattern info.*.json
		if !info.IsDir() && strings.HasPrefix(info.Name(), "info.") && strings.HasSuffix(info.Name(), ".json") {
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
		return nil, fmt.Errorf("failed to walk tmp directory during rebuild: %w", err)
	}

	// Write the rebuilt summaries to the summary file
	if err := s.writeSummaryFile(summaryFilePath, summaries); err != nil {
		fmt.Printf("Warning: failed to write summary file after rebuild: %v\n", err)
		// Still return the summaries even if we couldn't write the file
	}

	return summaries, nil
}

// addSummaryToFile adds a new ImageSummary to the summary file
func (s *Server) addSummaryToFile(summary ImageSummary) error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	tmpDir := filepath.Join(cwd, "tmp")
	summaryFilePath := filepath.Join(tmpDir, summaryFileName)

	// Read existing summaries
	summaries, err := s.readSummaryFile(summaryFilePath)
	if err != nil {
		// If file doesn't exist, start with empty slice
		summaries = []ImageSummary{}
	}

	// Check if summary already exists (by ImageID for completed analyses, or by ImageTag for pending)
	for i, existing := range summaries {
		// For completed analyses, match by ImageID
		if summary.ImageID != "" && existing.ImageID == summary.ImageID {
			summaries[i] = summary
			return s.writeSummaryFile(summaryFilePath, summaries)
		}
		// For pending analyses or when matching by image tag and status
		if summary.Status == "analyzing" && existing.ImageTag == summary.ImageTag && existing.Status == "analyzing" {
			summaries[i] = summary
			return s.writeSummaryFile(summaryFilePath, summaries)
		}
	}

	// Add new summary
	summaries = append(summaries, summary)
	return s.writeSummaryFile(summaryFilePath, summaries)
}

// removeSummaryFromFile removes an ImageSummary from the summary file by ID
func (s *Server) removeSummaryFromFile(imageID string) error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	tmpDir := filepath.Join(cwd, "tmp")
	summaryFilePath := filepath.Join(tmpDir, summaryFileName)

	// Read existing summaries
	summaries, err := s.readSummaryFile(summaryFilePath)
	if err != nil {
		return err // File must exist to remove from it
	}

	// Find and remove the summary
	updatedSummaries := make([]ImageSummary, 0)
	for _, summary := range summaries {
		if summary.ImageID != imageID {
			updatedSummaries = append(updatedSummaries, summary)
		}
	}

	// Write updated summaries back to file
	return s.writeSummaryFile(summaryFilePath, updatedSummaries)
}

// parseSummaryFile reads and parses a summary JSON file
func (s *Server) parseSummaryFile(filePath string) (ImageSummary, error) {
	var summary ImageSummary

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

	// Set default status if not present (for backward compatibility with existing files)
	if summary.Status == "" {
		summary.Status = "completed"
	}

	return summary, nil
}

// parseInfoFile reads and parses an info JSON file
func (s *Server) parseInfoFile(filePath string) (analyzer.ImageInfo, error) {
	var info analyzer.ImageInfo

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return info, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse JSON directly into ImageInfo struct
	err = json.Unmarshal(data, &info)
	if err != nil {
		return info, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return info, nil
}

// getInfoByID retrieves a specific info by its ID
func (s *Server) getInfoByID(id string) (analyzer.ImageInfo, error) {
	var info analyzer.ImageInfo

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return info, fmt.Errorf("failed to get current working directory: %w", err)
	}

	tmpDir := filepath.Join(cwd, "tmp")

	// Look for the specific info file
	var infoPath string
	err = filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileName := info.Name()
		if !info.IsDir() && strings.HasPrefix(fileName, "info.") && strings.HasSuffix(fileName, ".json") {
			// Extract ID from filename
			fileID := strings.TrimSuffix(strings.TrimPrefix(fileName, "info."), ".json")
			if fileID == id {
				infoPath = path
				return filepath.SkipDir // Stop walking
			}
		}

		return nil
	})

	if err != nil {
		return info, fmt.Errorf("failed to search for info file: %w", err)
	}

	if infoPath == "" {
		return info, fmt.Errorf("info with ID %s not found", id)
	}

	// Use the parseInfoFile method to read and parse the info
	info, err = s.parseInfoFile(infoPath)
	if err != nil {
		return info, fmt.Errorf("failed to parse info file: %w", err)
	}

	return info, nil
}

// deleteInfoByID removes a specific info file by its ID
func (s *Server) deleteInfoByID(id string) error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	tmpDir := filepath.Join(cwd, "tmp")

	// Look for the specific info file
	var infoPath string
	var imageFolder string

	err = filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileName := info.Name()
		if !info.IsDir() && strings.HasPrefix(fileName, "info.") && strings.HasSuffix(fileName, ".json") {
			// Extract ID from filename
			fileID := strings.TrimSuffix(strings.TrimPrefix(fileName, "info."), ".json")
			if fileID == id {
				infoPath = path
				imageFolder = filepath.Dir(path)
				return filepath.SkipDir // Stop walking
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to search for info file: %w", err)
	}

	if infoPath == "" {
		return fmt.Errorf("info with ID %s not found", id)
	}

	// Remove the info file
	if err := os.Remove(infoPath); err != nil {
		return fmt.Errorf("failed to delete info file: %w", err)
	}

	// Extract imageID for summary removal
	imageID := id
	// If it's a short ID, we need to find the full ID from the summary file
	// since the summary file stores full image IDs
	summaries, err := s.findSummaries()
	if err == nil {
		for _, summary := range summaries {
			shortImageID := strings.TrimPrefix(summary.ImageID, "sha256:")
			if len(shortImageID) > 12 {
				shortImageID = shortImageID[:12]
			}
			if shortImageID == id {
				imageID = summary.ImageID
				break
			}
		}
	}

	// Remove from summary file
	if err := s.removeSummaryFromFile(imageID); err != nil {
		fmt.Printf("Warning: failed to remove summary from file: %v\n", err)
		// Don't fail the operation if summary removal fails
	}

	// Check if the image folder is empty after removing the info file
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

// updateSummaryByRequestID updates an existing ImageSummary in the summary file by request ID
func (s *Server) updateSummaryByRequestID(requestID string, updatedSummary ImageSummary) error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	tmpDir := filepath.Join(cwd, "tmp")
	summaryFilePath := filepath.Join(tmpDir, summaryFileName)

	// Read existing summaries
	summaries, err := s.readSummaryFile(summaryFilePath)
	if err != nil {
		return fmt.Errorf("failed to read summary file: %w", err)
	}

	// Find and update the summary with matching request ID
	found := false
	for i, summary := range summaries {
		if summary.RequestID == requestID {
			summaries[i] = updatedSummary
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("summary with request ID %s not found", requestID)
	}

	// Write updated summaries back to file
	return s.writeSummaryFile(summaryFilePath, summaries)
}

// imageInfoToSummary converts an ImageInfo to an ImageSummary
func (s *Server) imageInfoToSummary(info analyzer.ImageInfo) ImageSummary {
	return ImageSummary{
		ImageID:      info.ImageID,
		ImageTag:     info.ImageTag,
		ImageSource:  info.ImageSource,
		ImageSize:    info.ImageSize,
		Architecture: info.Architecture,
		AnalyzedAt:   info.AnalyzedAt,
		Status:       "completed", // Default status for completed analysis
		RequestID:    "",          // Will be set by caller if needed
	}
}

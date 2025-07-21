# Summary File Implementation

## Overview
The `findSummaries` function in `server.go` has been improved to use a centralized summary file instead of scanning through all individual info files in the tmp directory each time. This significantly improves performance by avoiding filesystem scanning on every request.

## Changes Made

### 1. Added Summary File Constant
- Added `summaryFileName = "summaries.json"` constant to define the summary file name

### 2. Replaced `findSummaries` Function
The original function that walked through the tmp directory has been replaced with a more efficient implementation that:
- Reads from a centralized `summaries.json` file
- Falls back to rebuilding the summary file if it's missing or corrupt
- Returns cached summary data instead of parsing individual files

### 3. Added Helper Methods
- `readSummaryFile()` - Reads and parses the summary JSON file
- `writeSummaryFile()` - Writes the summary data to the JSON file
- `rebuildSummaryFile()` - Rebuilds the summary file by scanning all info files
- `addSummaryToFile()` - Adds a new ImageSummary to the summary file
- `removeSummaryFromFile()` - Removes an ImageSummary from the summary file
- `imageInfoToSummary()` - Converts analyzer.ImageInfo to webserver.ImageSummary

### 4. Updated Image Analysis Handlers
Both `handleAnalyzeImage` and `handleAnalyzeImageAsync` functions now:
- Check for existing images using the summary file
- After successful analysis, read the generated info file and add it to the summary file
- Handle both synchronous and asynchronous analysis workflows

### 5. Updated Delete Handler
The `deleteInfoByID` function now:
- Removes the individual info file as before
- Additionally removes the entry from the summary file
- Maps short image IDs to full image IDs for proper removal

## File Structure
The summary file `summaries.json` is stored in the `tmp/` directory and contains an array of `ImageSummary` objects:

```json
[
  {
    "image_id": "sha256:abc123...",
    "image_tag": "nginx:latest",
    "image_source": "",
    "image_size": 123456789,
    "architecture": "amd64",
    "analyzed_at": "2025-07-20T10:00:00Z"
  }
]
```

## Performance Benefits
- **Before**: O(n) filesystem scan for every API request to `/api/summaries`
- **After**: O(1) file read for cached summary data
- Automatic rebuilding when summary file is missing or corrupt ensures data consistency
- Reduced I/O operations and improved response times for the web interface

## Backward Compatibility
- Existing individual info files are preserved and continue to work
- The system automatically rebuilds the summary file from existing info files if needed
- No breaking changes to the API endpoints or response formats

## Error Handling
- Graceful fallback to rebuilding summary file if it's missing or corrupt
- Warning messages for non-critical failures (summary file operations)
- Continued operation even if summary file operations fail

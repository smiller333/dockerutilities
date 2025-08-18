# Changelog

All notable changes to Docker Utils will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive documentation suite including architecture, deployment, and security guides
- Enhanced CLI command structure with improved help text and examples
- Web interface for interactive Docker image analysis
- REST API for programmatic access to analysis functionality

### Changed
- Improved error handling and validation throughout the application
- Enhanced Docker client wrapper with security features
- Modernized build process with Go 1.24.2 support

### Security
- Added image source validation and trusted registry checks
- Implemented Docker socket security warnings and access validation
- Enhanced container security with non-root user execution

## [1.0.0] - 2025-01-21

### Added
- Initial release of Docker Utils
- CLI tool for Docker image analysis
- Web server with embedded static files
- Docker image inspection and metadata extraction
- Build context analysis with .dockerignore support
- Comprehensive test suite
- Docker container support with multi-stage builds
- Cross-platform binary builds (Linux, macOS, Windows)

### Features
- **CLI Commands**:
  - `dockerutils server` - Start web server for analysis
  - `dockerutils version` - Display version information
  - `dockerutils completion` - Generate shell completion scripts

- **Web Interface**:
  - Interactive image analysis form
  - Real-time analysis progress tracking
  - Filesystem browser for extracted image contents
  - Analysis result management and cleanup
  - Modern responsive UI with embedded assets

- **API Endpoints**:
  - `POST /api/analyze` - Start image analysis
  - `GET /api/summaries` - List analysis results
  - `GET /api/buildcontext` - Build context analysis
  - `DELETE /api/analysis/{id}` - Cleanup analysis data

- **Docker Integration**:
  - Official Docker SDK v28.3.0+ support
  - Image pulling and inspection
  - Container creation and filesystem extraction
  - Layer analysis and metadata extraction
  - Security warnings for untrusted images

- **Build Context Support**:
  - .dockerignore pattern matching using moby/patternmatcher
  - Recursive directory analysis
  - File size calculations and exclusion reporting

### Technical Details
- **Language**: Go 1.24.2
- **Dependencies**:
  - github.com/docker/docker v28.3.0+incompatible
  - github.com/spf13/cobra v1.9.1
  - github.com/moby/patternmatcher v0.6.0
  - github.com/opencontainers/image-spec v1.1.1
  - golang.org/x/text v0.26.0

- **Architecture**:
  - Modular package structure with clear separation of concerns
  - Interface-based design for testability
  - Embedded static assets for self-contained deployment
  - Asynchronous operations with progress tracking

### Security
- Docker socket access validation
- Image source verification against trusted registries
- Input validation and sanitization
- Non-root container execution
- Temporary file cleanup and resource management

### Documentation
- Comprehensive README with quick start guide
- Installation instructions for multiple platforms
- User guide with examples and troubleshooting
- API reference documentation
- Contributing guidelines and code of conduct
- Security policy and vulnerability reporting procedures

---

## Version History Format

This changelog follows the format:

### Version Format
- **[Major.Minor.Patch]** - Release date (YYYY-MM-DD)
- **[Unreleased]** - Changes not yet released

### Change Categories
- **Added** - New features
- **Changed** - Changes to existing functionality  
- **Deprecated** - Soon-to-be removed features
- **Removed** - Removed features
- **Fixed** - Bug fixes
- **Security** - Security improvements

### Semantic Versioning
- **Major** (X.0.0) - Breaking changes
- **Minor** (0.X.0) - New features (backward compatible)
- **Patch** (0.0.X) - Bug fixes (backward compatible)

### Example Entry Template

```markdown
## [1.1.0] - 2025-02-15

### Added
- New feature description with details about functionality
- Another feature with usage examples
- Integration with external service

### Changed
- Modified existing behavior with migration notes
- Updated dependency versions
- Improved performance in specific areas

### Fixed
- Bug fix description with issue reference (#123)
- Another bug fix with details

### Security
- Security improvement description
- Vulnerability patch details

### Deprecated
- Feature scheduled for removal in next major version
- Alternative recommendations provided
```

---

**Links:**
- [Unreleased]: https://github.com/smiller333/dockerutils/compare/v1.0.0...HEAD
- [1.0.0]: https://github.com/smiller333/dockerutils/releases/tag/v1.0.0

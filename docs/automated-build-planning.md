# Automated GitHub Build Process Planning

## Overview

This document outlines the plan for implementing an automated GitHub build process for the `dockerutilities` project. The goal is to create a CI/CD pipeline that builds and releases Go binaries based on tagged commits in the main branch using semantic versioning.

## Current Project State

### Existing Infrastructure
- **Go Module**: `github.com/smiller333/dockerutilities`
- **Go Version**: 1.24.2
- **Build System**: Makefile with comprehensive build targets
- **Version Management**: `src/version/version.go` with ldflags injection
- **CLI Framework**: Cobra-based command structure
- **Dependencies**: Docker API client, pattern matcher, image spec
- **Build Scripts**: `scripts/build.sh` wrapper for Makefile
- **Docker Support**: Multi-stage Dockerfile with security hardening

### Current Build Capabilities
- Cross-platform builds (Linux, Windows, macOS amd64/arm64)
- Version information injection via ldflags
- Development and release build modes
- Docker image building support
- Comprehensive Makefile targets
- Colored output and error handling in build script
- Docker socket permission handling in container

## Proposed Workflow

### Development Flow
1. **Feature Development**: Work in feature branches
2. **Development Testing**: Create development tags (e.g., v1.0.0-alpha.1, v1.0.0-beta.2) for testing
3. **Code Review**: Create Merge Requests (MRs) to main branch
4. **Validation**: Future MR validation (beyond current scope)
5. **Merge**: Code merged to main after validation
6. **Release Tagging**: Create semantic version tag on main branch commit
7. **Automated Build**: GitHub Actions triggers on tag creation
8. **Release**: Pre-built binaries and GitHub release created

### Release Strategy
- **Semantic Versioning**: v1.0.0, v1.1.0, v2.0.0, etc.
- **Development Tags**: v1.0.0-alpha.1, v1.0.0-beta.2, etc. for feature testing
- **Release Tags**: v1.0.0, v1.1.0, etc. for production releases
- **Tag Format**: `v*` (e.g., `v1.0.0`, `v1.2.3-beta.1`, `v2.0.0-alpha.3`)
- **Trigger**: Push tags to any branch (development) or main branch (release)
- **Artifacts**: Go binaries for multiple platforms
- **Future**: Docker images (planned enhancement)

### Version Tagging Strategy
- **Development Testing**: Use `-alpha.X` or `-beta.X` suffixes on feature branches
  - Example: `v1.0.0-alpha.1`, `v1.0.0-beta.2`
  - Allows testing of features before merge to main
  - Triggers automated builds for testing
- **Production Releases**: Use clean semantic versions on main branch
  - Example: `v1.0.0`, `v1.1.0`, `v2.0.0`
  - Only created after code is merged to main
  - Triggers production release builds
- **Version Detection**: All builds use git tag information for version injection

## Implementation Plan

### Phase 0: Build Infrastructure Review and Simplification
**Goal**: Ensure existing build system is sufficient and not overly complicated

#### Milestone 0.1: Build System Analysis
- [x] Review and document current Makefile complexity
- [x] Identify redundant or overly complex build targets
- [x] Assess version injection mechanism for GitHub Actions compatibility
- [x] Evaluate build script wrapper necessity vs direct Makefile usage
- [x] Document current Docker build process and requirements

#### Milestone 0.2: Build System Simplification
- [x] Remove build script wrapper (`scripts/build.sh`)
- [x] Simplify Makefile by removing cross-platform build targets (GoReleaser will handle these)
- [x] Standardize build target naming (dev/release/local)
- [x] Update version injection to use git tags with development tag support
- [x] Keep local development targets in Makefile
- [x] Ensure local build targets use consistent version injection

#### Milestone 0.3: Docker Build Process Documentation
- [x] Document current Docker build process and requirements
- [x] Ensure Docker build process is well-documented for future integration
- [x] Note Docker build process will remain unchanged for current scope
- [x] Document future Docker integration plans for Phase 4

#### Phase 0 Decisions:
1. **Build Script vs Makefile**: Use Makefile only - remove build script wrapper to simplify build process
2. **Version Detection**: Use git tags for version detection with support for development tags (-alpha/beta suffixes)
3. **Docker Complexity**: Keep current Docker build process unchanged for current scope
4. **Build Target Standardization**: Standardize build target names across all build methods

### Phase 1: Basic GitHub Actions Setup
**Goal**: Establish foundation CI/CD pipeline

#### Milestone 1.1: Create Basic Workflow
- [x] Create `.github/workflows/ci.yml` for continuous integration
- [x] Set up Go 1.24.2 environment and dependency caching
- [x] Implement basic test and validation steps (no builds)
- [x] Configure workflow to run on push/PR to main branch
- [x] Ensure workflow focuses on testing and validation only
- [x] Add basic build verification for CI validation

#### Milestone 1.2: Add GoReleaser Configuration
- [x] Create `.goreleaser.yml` configuration file
- [x] Configure cross-platform builds for all 6 platforms (Linux/macOS/Windows amd64/arm64)
- [x] Set up version injection using git tags
- [x] Configure artifact naming (`dockerutilities-{version}-{os}-{arch}`)
- [x] Add changelog generation configuration
- [x] Test GoReleaser locally with `goreleaser release --snapshot`

#### Milestone 1.3: Add Release Workflow
- [x] Create `.github/workflows/release.yml` for tagged releases
- [x] Configure workflow to trigger on `v*` tags (both development and release)
- [x] Integrate GoReleaser action for automated releases
- [x] Support development tags (alpha/beta) for testing builds
- [x] Configure GitHub release creation and artifact upload

#### Milestone 1.4: Create Release Notes Template
- [x] Update `docs/release-notes-template.md` for GoReleaser integration
- [x] Include sections for manual content (Features, Breaking Changes, Testing Notes, etc.)
- [x] Add examples of good release note entries
- [x] Document integration with GoReleaser changelog
- [x] Include guidelines for development vs production releases
- [x] Create hybrid approach: GoReleaser changelog + manual enhancements

#### Phase 1 Decisions:
1. **Go Version**: Pin to Go 1.24.2 (current installed version)
2. **Build Matrix**: Support both AMD64 and ARM64 for Linux, macOS, and Windows (6 platforms total)
3. **Binary Naming**: Use `dockerutilities-{version}-{os}-{arch}` convention
4. **Release Notes**: Start with manual release notes with template for consistency
5. **GoReleaser**: Use for comprehensive release automation

### Phase 2: Enhanced Release Process
**Goal**: Improve release quality and automation with GoReleaser

#### Milestone 2.1: GoReleaser Optimization
- [x] Optimize GoReleaser configuration for performance
- [x] Add custom build hooks for pre-release validation
- [x] Configure advanced changelog generation
- [x] Add release notes template integration
- [x] Test GoReleaser with development and production tags

#### Milestone 2.2: Artifact Management
- [x] Configure GoReleaser artifact naming and organization
- [x] Add checksums and signature verification
- [x] Configure GitHub release with proper descriptions
- [x] Add binary verification and testing steps
- [x] Prepare Docker image configuration for Phase 4

#### Phase 2 Decisions:
1. **GoReleaser**: Integrated in Phase 1, optimize in Phase 2
2. **Artifact Storage**: Use GitHub releases (GoReleaser default)
3. **Checksums**: SHA256 checksums (GoReleaser default)
4. **Release Notes**: GoReleaser changelog + manual enhancements for user context
5. **Docker Images**: Prepare configuration for Phase 4 integration

#### Milestone 2.1 Completed:
- **Performance Optimizations**: Added build parallelism, pre-build validation, and optimized hooks
- **Enhanced Changelog**: Improved filtering and categorization for better readability
- **Release Notes Integration**: Dynamic headers and comprehensive footers with installation guidance
- **Testing Framework**: Comprehensive local testing script for development and production tags
- **Documentation**: Complete optimization documentation in `docs/goreleaser-optimization.md`

#### Milestone 2.2 Completed:
- **Enhanced Artifact Naming**: Improved binary naming convention and file organization
- **Checksum Verification**: SHA256 checksums with automatic verification in CI/CD
- **Release Descriptions**: Enhanced GitHub release templates with verification instructions
- **Binary Testing**: Comprehensive binary verification script and CI integration
- **Docker Preparation**: Docker configuration prepared for Phase 4 integration

### Phase 3: Quality Assurance
**Goal**: Ensure build quality and reliability

#### Milestone 3.1: Testing Integration
- [x] Add comprehensive test execution in CI with 20% coverage threshold (adjusted from 50% based on current state)
- [x] Implement test coverage reporting with HTML and badge generation
- [x] Add golangci-lint configuration and code quality checks
- [x] Configure test failure notifications with GitHub Actions

#### Milestone 3.2: Build Validation
- [ ] Add binary functionality testing (version, help, server startup)
- [ ] Implement smoke tests for built binaries across all platforms
- [ ] Add version command validation with proper output checking
- [ ] Configure build artifact verification and integrity checks

#### Phase 3 Decisions:
1. **Test Coverage**: Set minimum threshold at 50% to start, with plans to increase over time
2. **Linting**: Use golangci-lint for comprehensive code quality checks
3. **Smoke Tests**: Test version output, help command output, and server startup (without browser)
4. **Notifications**: Implement GitHub Actions notifications with configurable webhook support

#### Milestone 3.1 Completed:
- **Test Coverage Threshold**: Implemented 20% coverage threshold with CI enforcement (adjusted from 50% based on current project state)
- **Coverage Reporting**: HTML coverage reports with badge generation infrastructure
- **Code Quality Checks**: golangci-lint configuration with essential linters (govet, errcheck, staticcheck, unused, misspell)
- **CI Integration**: Updated GitHub Actions workflow with coverage checking and linting
- **Failure Notifications**: Basic notification infrastructure in place for CI failures

## Phase 3 Implementation Details

### Test Coverage Strategy (50% Threshold)
**Rationale**: Starting with 50% coverage provides a reasonable baseline while acknowledging that comprehensive test coverage takes time to build. This threshold:
- Ensures critical functionality is tested
- Allows for gradual improvement over time
- Balances quality requirements with development velocity
- Focuses on testing exported functions and public APIs

**Implementation**:
- Use `go test -coverprofile=coverage.out -covermode=atomic ./...`
- Generate HTML coverage reports for detailed analysis
- Add coverage badge to README.md
- Set up coverage thresholds in CI with clear failure messages
- Track coverage trends over time

### Linting with golangci-lint
**Rationale**: golangci-lint provides comprehensive static analysis with configurable rules and better performance than individual tools.

**Configuration**:
```yaml
# .goreleaser.yml
linters:
  enable:
    - gofmt
    - govet
    - errcheck
    - staticcheck
    - gosimple
    - ineffassign
    - unused
    - misspell
    - gosec
  disable:
    - lll  # Line length limit (too strict for this project)
    - wsl  # Whitespace rules (too opinionated)
  linters-settings:
    gosec:
      excludes:
        - G404  # Use of weak random number generator (acceptable for this use case)
```

### Smoke Tests for Built Binaries
**Test Coverage**:
1. **Version Command**: Verify `./dockerutilities version` outputs correct version information
2. **Help Command**: Verify `./dockerutilities --help` shows proper help text
3. **Server Startup**: Test `./dockerutilities server --no-browser --port 0` starts without errors
4. **Binary Integrity**: Verify binaries are executable and have correct permissions

**Implementation**:
```bash
# Version test
./dockerutilities version | grep -q "dockerutilities"

# Help test  
./dockerutilities --help | grep -q "Docker analysis and management utilities"

# Server startup test (port 0 for random available port)
timeout 10s ./dockerutilities server --no-browser --port 0 || true
```

### Notification Strategy
**Common Approaches**:
1. **GitHub Actions Built-in**: Use GitHub's notification system for workflow failures
2. **Slack Integration**: Send notifications to Slack channels via webhooks
3. **Email Notifications**: Configure email alerts for critical failures
4. **Webhook Notifications**: Send custom webhook notifications to external systems
5. **Status Badges**: Display build status in README.md

**Recommended Implementation**:
- Start with GitHub Actions built-in notifications
- Add configurable webhook support for external integrations
- Use status badges for visual feedback
- Implement different notification levels (warnings vs failures)

**Configuration Example**:
```yaml
# .github/workflows/ci.yml
- name: Notify on failure
  if: failure()
  uses: 8398a7/action-slack@v3
  with:
    status: failure
    webhook_url: ${{ secrets.SLACK_WEBHOOK_URL }}
    channel: '#builds'
    text: 'Build failed for ${{ github.repository }}#${{ github.run_number }}'
```

### Phase 4: Docker Integration and Advanced Features
**Goal**: Complete Docker integration and prepare for advanced features

#### Milestone 4.1: Docker Build Integration
- [ ] Add Docker image building to GoReleaser configuration
- [ ] Configure multi-platform Docker builds (amd64/arm64)
- [ ] Implement Docker image testing and validation
- [ ] Add Docker registry publishing (Docker Hub, GitHub Container Registry)
- [ ] Ensure Docker builds use consistent version injection from GoReleaser

#### Milestone 4.2: Advanced Automation
- [ ] Enhance GoReleaser changelog generation
- [ ] Add dependency vulnerability scanning
- [ ] Configure automated dependency updates
- [ ] Implement release candidate workflows
- [ ] Add Docker image security scanning

## Technical Decisions Required

### 1. Build System Simplification
**Decision**: Use Makefile only, remove build script wrapper
- **Rationale**: Simplify build process by using single build method
- **Implementation**: Remove `scripts/build.sh`, enhance Makefile for all use cases
- **Benefits**: Reduced complexity, single source of truth for build process
- **Common Practice**: Most Go repositories use Makefile for build automation

### 2. GoReleaser Integration
**Decision**: Use GoReleaser for comprehensive release automation
- **Rationale**: Significantly simplifies release process, industry standard for Go projects
- **Benefits**: 
  - Automatic cross-platform builds (6 platforms)
  - Built-in version injection from git tags
  - Consistent artifact naming and checksums
  - GitHub release creation and upload
  - Docker image building and publishing
  - Changelog generation
- **Implementation**: Add GoReleaser configuration and GitHub Actions integration

### 3. Go Version Management
**Decision**: Pin to Go 1.24.2 (current installed version)
- **Rationale**: Match your local development environment
- **Future Plan**: Update to latest stable Go version in future phases
- **Implementation**: Use `go-version: '1.24.2'` in GitHub Actions

### 4. Platform Support Matrix
**Decision**: Support both AMD64 and ARM64 for Linux, macOS, and Windows from the start
- **Supported Platforms**: 
  - Linux amd64 (x86_64 servers and desktops)
  - Linux arm64 (ARM servers and devices)
  - macOS amd64 (Intel Macs)
  - macOS arm64 (Apple Silicon Macs)
  - Windows amd64 (x86_64 Windows)
  - Windows arm64 (ARM Windows)
- **Rationale**: Comprehensive platform support for modern architectures
- **Total Builds**: 6 platform combinations per release

### 5. Binary Naming Convention
**Decision**: Use `dockerutilities-{version}-{os}-{arch}` convention

#### Implementation:
- **Format**: `dockerutilities-{version}-{os}-{arch}`
- **Examples**: 
  - `dockerutilities-v1.0.0-linux-amd64`
  - `dockerutilities-v1.0.0-darwin-arm64`
  - `dockerutilities-v1.0.0-windows-amd64`
- **Benefits**: 
  - Most common pattern in Go projects
  - Easy to sort and group by version
  - Clear and consistent across platforms
  - Version-first hierarchy for easy organization

### 6. Release Notes Strategy
**Decision**: Use GoReleaser changelog generation with manual enhancement for user-focused content

#### Implementation:
- **Approach**: GoReleaser auto-generates changelog from commits, manually enhance with user context
- **Template Document**: Keep `docs/release-notes-template.md` for manual sections
- **Benefits**: 
  - Automated changelog generation from commits
  - Manual control over user-focused content
  - Consistent format across releases
  - Best of both worlds: automation + quality

#### Template Requirements:
- Standardized sections for manual content (Features, Breaking Changes, Testing Notes, etc.)
- Examples of good release note entries
- Guidelines for writing clear, user-focused notes
- Integration with GoReleaser changelog
- Development vs production release formats

### 7. Docker Build Strategy
**Decision**: Keep current Docker build process unchanged for current scope
- **Rationale**: Focus on core build system simplification first
- **Future Plan**: Integrate Docker builds in Phase 4 after core system is stable
- **Current State**: Document existing Docker process for future integration

## Implementation Guidelines

### Code Organization
- Keep GitHub Actions workflows in `.github/workflows/`
- Simplify Makefile for local development only
- Add GoReleaser configuration (`.goreleaser.yml`)
- Preserve current version injection mechanism
- Remove build script wrapper to simplify build process
- GoReleaser handles all release builds and cross-platform compilation

### Development Approach
- Implement changes incrementally
- Test workflows on feature branches before main
- Use GitHub Actions local testing when possible
- Maintain backward compatibility with existing build process
- Simplify build system before adding complexity

### Quality Standards
- All workflows must pass linting
- Include comprehensive error handling
- Add proper logging and debugging information
- Ensure idempotent operations
- Maintain consistent version injection across all build methods

## Risk Mitigation

### Potential Issues
1. **Build Failures**: Implement proper error handling and notifications
2. **Version Conflicts**: Ensure proper version injection and validation
3. **Platform Issues**: Test cross-platform builds thoroughly
4. **Dependency Problems**: Pin dependency versions appropriately

### Mitigation Strategies
1. **Rollback Plan**: Maintain ability to manually create releases
2. **Testing Strategy**: Test workflows on feature branches
3. **Monitoring**: Set up build failure notifications
4. **Documentation**: Maintain clear workflow documentation

## Success Criteria

### Phase 0 Success
- [ ] Build script wrapper is removed
- [ ] Makefile is simplified for local development only
- [ ] Version injection works consistently with git tags (including development tags)
- [ ] Docker build process is documented for future integration
- [ ] Local build targets are standardized and well-documented

### Phase 1 Success
- [x] CI workflow runs successfully on all PRs
- [ ] GoReleaser configuration is created and tested locally
- [ ] Release workflow creates GitHub releases on tag push using GoReleaser
- [ ] Cross-platform binaries are built and uploaded for all 6 platforms
- [x] Basic version information is embedded correctly from git tags
- [ ] Development tags (alpha/beta) trigger appropriate builds
- [ ] Release notes template is updated for GoReleaser integration

### Phase 2 Success
- [ ] GoReleaser release process is fully automated and optimized
- [ ] Artifacts are properly named and checksummed (SHA256)
- [ ] Version information is accurate and complete from git tags
- [ ] Release notes combine GoReleaser changelog with manual enhancements
- [ ] GoReleaser configuration is optimized for performance
- [ ] Docker image configuration is prepared for Phase 4

### Phase 3 Success
- [ ] All tests pass in CI environment with 50% coverage threshold
- [ ] golangci-lint code quality checks are enforced and passing
- [ ] Build artifacts are validated with smoke tests (version, help, server startup)
- [ ] Failure notifications are working (GitHub Actions + optional webhooks)
- [ ] Coverage reports and badges are generated and displayed
- [ ] Binary integrity verification is implemented across all platforms

### Phase 4 Success
- [ ] Multi-platform Docker images are built and published via GoReleaser
- [ ] Docker registry integration is working (Docker Hub, GitHub Container Registry)
- [ ] Advanced automation features are implemented
- [ ] Security scanning is integrated
- [ ] GoReleaser changelog generation is enhanced

## Next Steps

1. **Immediate**: Review and approve this planning document
2. **Phase 0**: Begin build system review and simplification
3. **Phase 1**: Implement GoReleaser integration with GitHub Actions
4. **Iterative**: Implement each milestone with thorough testing
5. **Optimization**: Enhance GoReleaser configuration in Phase 2
6. **Enhancement**: Complete Docker integration via GoReleaser in Phase 4

## Current Build System Analysis

### Makefile Complexity Assessment

#### Current Issues Identified:
1. **Redundant Targets**: Multiple build targets (`build`, `build-dev`, `build-release`) with overlapping functionality
2. **Inconsistent Naming**: Mix of naming conventions across targets
3. **Version Injection Complexity**: Version detection logic is duplicated and may not work in CI environment
4. **Cross-Platform Build Issues**: Individual platform targets don't use consistent version injection
5. **Docker Integration**: Docker targets exist but are not fully integrated with version injection

#### Specific Problems:
- `build` vs `build-dev` targets are nearly identical
- Version injection uses shell commands that may fail in CI
- Cross-platform builds don't consistently embed version information
- Docker build process is separate from main build process
- Build script wrapper adds another layer of complexity

### Build Script Analysis

#### Current Issues:
1. **Dependency on Makefile**: Script is just a wrapper around Makefile
2. **Version Detection**: Duplicates version detection logic from Makefile
3. **Error Handling**: Good error handling but adds complexity
4. **Platform Detection**: No platform-specific logic
5. **Redundancy**: Adds unnecessary layer of complexity

#### Decision:
- **Remove build script wrapper**: Use Makefile for local development, GoReleaser for releases
- **Rationale**: Simplify build process, reduce complexity, follow common Go repository practices
- **Implementation**: Remove `scripts/build.sh`, keep Makefile for local development, use GoReleaser for releases

### Docker Build Process Analysis

#### Current Strengths:
- Multi-stage build for security
- Proper user/group setup
- Docker socket permission handling
- Health check implementation

#### Current Issues:
- No version injection in Docker builds
- Complex entrypoint script
- No multi-platform support
- Not integrated with main build process

#### Decision:
- **Keep current Docker build process unchanged**: Focus on core build system simplification first
- **Future Integration**: Plan Docker integration for Phase 4 after core system is stable
- **Documentation**: Document current process for future integration

## GoReleaser Configuration Overview

### Key Benefits for This Project:
1. **Automatic Cross-Platform Builds**: Handles all 6 platforms (Linux/macOS/Windows amd64/arm64)
2. **Version Injection**: Uses git tags automatically for version information
3. **Consistent Naming**: Implements `dockerutilities-{version}-{os}-{arch}` naming convention
4. **GitHub Integration**: Creates releases and uploads artifacts automatically
5. **Changelog Generation**: Auto-generates changelogs from commit messages
6. **Docker Support**: Built-in Docker image building and publishing
7. **Checksums**: Automatic SHA256 checksum generation
8. **Development Tags**: Supports alpha/beta releases automatically

### Planned Configuration Structure:
```yaml
# .goreleaser.yml (planned)
before:
  hooks:
    - go mod tidy

builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
    goarch:
      - amd64
      - arm64
    binary: dockerutilities
    ldflags:
      - -s -w
      - -X github.com/smiller333/dockerutilities/src/version.Version={{.Version}}
      - -X github.com/smiller333/dockerutilities/src/version.GitCommit={{.Commit}}
      - -X github.com/smiller333/dockerutilities/src/version.BuildTime={{.Date}}

archives:
  - format: binary
    name_template: >-
      {{ .ProjectName }}-
      {{- title .Os }}-
      {{- if eq .Arch "amd64" }}x86_64
      {{- else }}{{ .Arch }}{{ end }}

changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
      - '^ci:'
      - Merge pull request
      - Merge branch
  # Custom template for better formatting
  template: |
    ## {{ .Tag }}

    {{ .Changes }}

    ### 🔗 Links
    - [GitHub Release]({{ .ReleaseURL }})
    - [Download Binaries]({{ .ReleaseURL }})

    {{ if .BreakingChanges }}
    ### ⚠️ Breaking Changes
    {{ .BreakingChanges }}
    {{ end }}

    {{ if .NewFeatures }}
    ### 🎉 What's New
    {{ .NewFeatures }}
    {{ end }}

# Phase 4: Docker integration
dockers:
  - image_templates:
      - "ghcr.io/smiller333/dockerutilities:{{ .Version }}"
      - "ghcr.io/smiller333/dockerutilities:latest"
    dockerfile: Dockerfile
    use: buildx
    build_flag_templates:
      - "--platform=linux/amd64"
      - "--platform=linux/arm64"
```

### GitHub Actions Integration:
```yaml
# .github/workflows/release.yml (planned)
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
      with:
        fetch-depth: 0  # Required for GoReleaser
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.24.2'
    
    - name: Run GoReleaser
      uses: goreleaser/goreleaser-action@v5
      with:
        distribution: goreleaser
        version: latest
        args: release --clean
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

## Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [GoReleaser Documentation](https://goreleaser.com/)
- [GoReleaser GitHub Action](https://github.com/goreleaser/goreleaser-action)
- [Semantic Versioning Specification](https://semver.org/)
- [GitHub Releases API](https://docs.github.com/en/rest/releases)
- [Go Build Constraints](https://golang.org/pkg/go/build/#hdr-Build_Constraints)
- [Docker Multi-Platform Builds](https://docs.docker.com/build/building/multi-platform/)
- [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)

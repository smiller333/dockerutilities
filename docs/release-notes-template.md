# Release Notes Template

This document provides a template and guidelines for writing consistent release notes for the `dockerutilities` project, with integration for GoReleaser automated changelog generation.

## GoReleaser Integration

### Overview
This project uses GoReleaser for automated release management, which includes:
- **Automated Changelog Generation**: Creates changelogs from git commit messages
- **Cross-Platform Builds**: Builds binaries for 6 platforms (Linux/macOS/Windows amd64/arm64)
- **GitHub Release Creation**: Automatically creates GitHub releases with artifacts
- **Version Injection**: Embeds version information from git tags

### Hybrid Approach: GoReleaser + Manual Enhancement
We use a hybrid approach that combines:
1. **GoReleaser Automated Changelog**: Auto-generated from commit messages
2. **Manual Enhancement**: User-focused content added manually for clarity and context

### GoReleaser Changelog Configuration
The project's `.goreleaser.yml` includes changelog configuration:
```yaml
changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
      - '^ci:'
      - Merge pull request
      - Merge branch
```

This configuration:
- Sorts commits chronologically (ascending)
- Excludes documentation, test, and CI commits
- Excludes merge commits for cleaner changelogs

## Template Structure

### Production Release Template (GoReleaser Enhanced)

```markdown
## [Version] - [Date]

### 🎉 What's New
- [New feature description with user benefit]
- [Another new feature with context]

### 🐛 Bug Fixes
- [Bug fix description with impact]
- [Another bug fix with issue number if applicable]

### 🔧 Improvements
- [Improvement description with user benefit]
- [Another improvement with context]

### ⚠️ Breaking Changes
- [Breaking change description with migration steps]

### 📦 Dependencies
- Updated [dependency] to version [X.Y.Z] for [reason]

### 🔗 Links
- [GitHub Release](link-to-github-release)
- [Download Binaries](link-to-binaries)

---

## 📋 Detailed Changelog

[GoReleaser-generated changelog content goes here]

### Commit Categories
- **Features**: New functionality added
- **Fixes**: Bug fixes and issue resolutions
- **Improvements**: Enhancements to existing features
- **Documentation**: Documentation updates and clarifications
- **CI/CD**: Build system and automation improvements
```

### Development Release Template (Alpha/Beta) - GoReleaser Enhanced

```markdown
## [Version] - [Date] - [Alpha/Beta Release]

### 🧪 Development Release
This is a development release for testing purposes. Not recommended for production use.

### 🎯 Release Focus
- [Primary features being tested]
- [Key improvements in this release]

### 🎉 What's New
- [New feature description with testing context]
- [Another new feature with usage notes]

### 🐛 Bug Fixes
- [Bug fix description with testing impact]
- [Another bug fix with verification steps]

### 🔧 Improvements
- [Improvement description with testing benefit]
- [Another improvement with context]

### ⚠️ Known Issues
- [Known issue description with workarounds]
- [Another known issue with impact]

### 📝 Testing Notes
- [Specific testing instructions]
- [Expected behavior and limitations]
- [How to report issues]

### 🔗 Links
- [GitHub Release](link-to-github-release)
- [Download Binaries](link-to-binaries)

---

## 📋 Detailed Changelog

[GoReleaser-generated changelog content goes here]

### Development Commit Categories
- **Features**: New functionality being tested
- **Fixes**: Bug fixes and stability improvements
- **Improvements**: Enhancements and optimizations
- **Testing**: Test-related changes and improvements
```

## Writing Guidelines

### General Principles

1. **User-Focused**: Write from the user's perspective
2. **Clear and Concise**: Use simple, direct language
3. **Consistent Format**: Follow the template structure
4. **Actionable**: Tell users what they need to know or do
5. **GoReleaser Integration**: Leverage automated changelog while adding user context

### Section Guidelines

#### 🎉 What's New
- Focus on user-visible features and improvements
- Explain the benefit to users
- Use present tense
- Be specific about what changed
- **For Development Releases**: Include testing context and usage notes

**Good Examples:**
- Added Docker image layer analysis with detailed file listings
- Implemented web interface for interactive image exploration
- Added support for multi-platform Docker image builds
- **Development**: Added Docker Compose analysis (alpha feature - test with simple compose files)

**Avoid:**
- "Fixed various bugs" (too vague)
- "Improved performance" (not specific enough)

#### 🐛 Bug Fixes
- Describe the problem that was fixed
- Mention the impact on users
- Include issue numbers if applicable
- **For Development Releases**: Include verification steps

**Good Examples:**
- Fixed crash when analyzing images with empty layers
- Resolved memory leak in large image processing
- Fixed incorrect file permissions in extracted layers
- **Development**: Fixed circular dependency detection (verify with complex compose files)

#### 🔧 Improvements
- Describe enhancements to existing features
- Focus on user benefits
- Be specific about improvements
- **For Development Releases**: Include testing context

**Good Examples:**
- Improved error messages for better debugging
- Enhanced performance for large image analysis
- Added progress indicators for long-running operations
- **Development**: Enhanced compose file parsing (test with various formats)

#### ⚠️ Breaking Changes
- Clearly explain what changed
- Provide migration steps if needed
- Use warning emoji to draw attention
- Explain why the change was necessary
- **For Development Releases**: Emphasize this is expected in alpha/beta releases

**Good Examples:**
- Changed default port from 8080 to 9090 (use `--port` flag to override)
- Renamed `--image` flag to `--target` for clarity
- Removed deprecated `--legacy` option
- **Development**: Changed compose analysis API (breaking change expected in alpha)

#### 📦 Dependencies
- List significant dependency updates
- Mention security updates
- Note any compatibility changes
- **For Development Releases**: Note any experimental dependencies

#### 📝 Testing Notes (Development Releases Only)
- Provide specific testing instructions
- List expected behavior and limitations
- Include how to report issues
- Mention any special setup requirements

**Good Examples:**
- Test with various Docker Compose file formats (v1, v2, v3)
- Verify environment variable resolution works correctly
- Check performance with large compose projects (>50 services)
- Report issues with sample compose files attached

### Version Numbering

#### Semantic Versioning
- **Major.Minor.Patch** (e.g., 1.0.0, 1.2.3)
- **Major**: Breaking changes
- **Minor**: New features, backward compatible
- **Patch**: Bug fixes, backward compatible

#### Development Versions
- **Alpha**: Early development, major features incomplete
  - Format: `v1.0.0-alpha.1`, `v1.0.0-alpha.2`
  - Use for testing incomplete features
- **Beta**: Feature complete, testing phase
  - Format: `v1.0.0-beta.1`, `v1.0.0-beta.2`
  - Use for final testing before release

### GoReleaser Integration Guidelines

#### Commit Message Standards
To ensure good GoReleaser changelog generation, follow these commit message patterns:

**Feature Commits:**
```
feat: add Docker Compose analysis support
feat(analyzer): implement layer-by-layer file analysis
```

**Bug Fix Commits:**
```
fix: resolve memory leak in large image processing
fix(server): handle empty response from Docker API
```

**Improvement Commits:**
```
improve: enhance error messages for better debugging
improve(ui): add progress indicators for long operations
```

**Documentation Commits:**
```
docs: update installation instructions
docs(api): add examples for new endpoints
```

**CI/CD Commits:**
```
ci: add automated testing workflow
ci: update GoReleaser configuration
```

#### Changelog Enhancement Process
1. **Review GoReleaser Output**: Check the auto-generated changelog
2. **Add User Context**: Enhance with user-focused descriptions
3. **Categorize Changes**: Group related changes logically
4. **Add Migration Steps**: For breaking changes
5. **Include Testing Notes**: For development releases

### Examples

#### Production Release Example (GoReleaser Enhanced)

```markdown
## v1.2.0 - 2024-01-15

### 🎉 What's New
- Added support for Docker image vulnerability scanning with detailed CVE reporting
- Implemented image comparison tool for detecting differences between image versions
- Added export functionality for analysis reports in JSON and CSV formats

### 🐛 Bug Fixes
- Fixed memory leak when processing large images (#123) - significantly reduced memory usage
- Resolved crash on Windows when accessing Docker socket (#124) - improved Windows compatibility
- Fixed incorrect file size reporting in layer analysis - now shows accurate sizes

### 🔧 Improvements
- Improved error handling for network timeouts with retry logic
- Enhanced progress reporting for long operations with percentage indicators
- Added better validation for Docker registry credentials with helpful error messages

### 📦 Dependencies
- Updated Docker API client to v28.3.0 for latest Docker Engine support
- Updated Cobra CLI framework to v1.9.1 for improved command handling

### 🔗 Links
- [GitHub Release](https://github.com/smiller333/dockerutilities/releases/tag/v1.2.0)
- [Download Binaries](https://github.com/smiller333/dockerutilities/releases/tag/v1.2.0)

---

## 📋 Detailed Changelog

### Features
- feat: add Docker image vulnerability scanning support
- feat(analyzer): implement image comparison functionality
- feat(export): add JSON and CSV export options

### Fixes
- fix: resolve memory leak in large image processing
- fix(server): handle Windows Docker socket access
- fix(analyzer): correct file size reporting in layers

### Improvements
- improve: enhance error handling with retry logic
- improve(ui): add progress indicators for long operations
- improve(auth): better Docker registry credential validation

### Documentation
- docs: update installation and usage instructions
- docs(api): add examples for new vulnerability scanning endpoints
```

#### Development Release Example (GoReleaser Enhanced)

```markdown
## v1.3.0-alpha.1 - 2024-01-10 - Alpha Release

### 🧪 Development Release
This is an alpha release for testing the new Docker Compose analysis features. Not recommended for production use.

### 🎯 Release Focus
- Docker Compose file analysis and validation
- Multi-service dependency mapping
- Environment variable resolution

### 🎉 What's New
- Added Docker Compose file analysis and validation (test with v1, v2, v3 formats)
- Implemented multi-service dependency mapping (verify dependency graphs)
- Added support for environment variable resolution (test with complex expressions)

### 🐛 Bug Fixes
- Fixed issue with circular dependency detection (verify with complex compose files)
- Resolved memory usage in large compose files (test with >50 services)

### ⚠️ Known Issues
- Environment variable substitution may fail with complex expressions (workaround: use simple env vars)
- Large compose files (>100 services) may cause performance issues (monitor memory usage)
- Windows path handling needs additional testing (test on Windows systems)

### 📝 Testing Notes
- Test with various Docker Compose file formats (v1, v2, v3)
- Verify environment variable resolution works correctly with simple and complex expressions
- Check performance with large compose projects (>50 services)
- Test circular dependency detection with complex service relationships
- Report issues with sample compose files attached

### 🔗 Links
- [GitHub Release](https://github.com/smiller333/dockerutilities/releases/tag/v1.3.0-alpha.1)
- [Download Binaries](https://github.com/smiller333/dockerutilities/releases/tag/v1.3.0-alpha.1)

---

## 📋 Detailed Changelog

### Features
- feat: add Docker Compose file analysis and validation
- feat(analyzer): implement multi-service dependency mapping
- feat(env): add environment variable resolution support

### Fixes
- fix: resolve circular dependency detection issues
- fix(performance): optimize memory usage for large compose files

### Improvements
- improve(parser): enhance compose file parsing performance
- improve(ui): add dependency graph visualization

### Testing
- test: add comprehensive compose file test cases
- test: add performance benchmarks for large files
```

## Checklist for Release Notes

Before publishing release notes, ensure:

### General Checklist
- [ ] All sections are filled out appropriately
- [ ] User-focused language is used throughout
- [ ] Breaking changes are clearly explained with migration steps
- [ ] Links to GitHub release and binaries are included
- [ ] Version number and date are correct

### GoReleaser Integration Checklist
- [ ] GoReleaser changelog is reviewed and enhanced
- [ ] Manual sections provide user context beyond commit messages
- [ ] Commit categories are properly organized
- [ ] Breaking changes are highlighted in both manual and changelog sections

### Development Release Checklist
- [ ] Development releases are clearly marked as such
- [ ] Known issues are documented with workarounds
- [ ] Testing notes are provided with specific instructions
- [ ] Release focus is clearly stated
- [ ] Expected behavior and limitations are documented

### Production Release Checklist
- [ ] All features are properly described with user benefits
- [ ] Bug fixes include impact descriptions
- [ ] Improvements focus on user value
- [ ] Dependencies are listed with reasons for updates

## Tips for Good Release Notes

1. **Start Early**: Begin collecting notes during development
2. **Be Consistent**: Use the same format and style every time
3. **Think Like a User**: Focus on what users need to know
4. **Keep It Simple**: Avoid technical jargon when possible
5. **Be Honest**: Don't hide known issues or limitations
6. **Provide Context**: Explain why changes were made when relevant
7. **Use Emojis**: They help organize and make notes more readable
8. **Link Everything**: Provide links to issues, PRs, and downloads
9. **Leverage GoReleaser**: Use automated changelog as foundation, enhance with user context
10. **Test Your Notes**: Have someone unfamiliar with the project read them

## GoReleaser Workflow Integration

### Release Process
1. **Create Git Tag**: `git tag v1.2.0 && git push origin v1.2.0`
2. **GitHub Actions Trigger**: Release workflow runs automatically
3. **GoReleaser Execution**: Builds binaries and generates changelog
4. **Manual Enhancement**: Add user-focused content to release notes
5. **Publish**: Release is published with both automated and manual content

### Development Release Process
1. **Create Development Tag**: `git tag v1.3.0-alpha.1 && git push origin v1.3.0-alpha.1`
2. **Automated Build**: GoReleaser creates development release
3. **Testing Focus**: Emphasize testing instructions and known issues
4. **Feedback Collection**: Use development releases to gather user feedback

### Best Practices
- **Consistent Tagging**: Use semantic versioning for all releases
- **Clear Commit Messages**: Follow conventional commit format for better changelog generation
- **Regular Releases**: Maintain consistent release cadence
- **User Feedback**: Incorporate feedback from development releases into production releases

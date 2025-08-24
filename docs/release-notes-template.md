# Release Notes Template

This document provides a template and guidelines for writing consistent release notes for the `dockerutilities` project.

## Template Structure

### Production Release Template

```markdown
## [Version] - [Date]

### 🎉 What's New
- [New feature description]
- [Another new feature]

### 🐛 Bug Fixes
- [Bug fix description]
- [Another bug fix]

### 🔧 Improvements
- [Improvement description]
- [Another improvement]

### ⚠️ Breaking Changes
- [Breaking change description and migration guide]

### 📦 Dependencies
- Updated [dependency] to version [X.Y.Z]

### 🔗 Links
- [GitHub Release](link-to-github-release)
- [Download Binaries](link-to-binaries)
```

### Development Release Template (Alpha/Beta)

```markdown
## [Version] - [Date] - [Alpha/Beta Release]

### 🧪 Development Release
This is a development release for testing purposes. Not recommended for production use.

### 🎉 What's New
- [New feature description]
- [Another new feature]

### 🐛 Bug Fixes
- [Bug fix description]
- [Another bug fix]

### 🔧 Improvements
- [Improvement description]
- [Another improvement]

### ⚠️ Known Issues
- [Known issue description]
- [Another known issue]

### 📝 Testing Notes
- [Instructions for testing specific features]
- [Known limitations]

### 🔗 Links
- [GitHub Release](link-to-github-release)
- [Download Binaries](link-to-binaries)
```

## Writing Guidelines

### General Principles

1. **User-Focused**: Write from the user's perspective
2. **Clear and Concise**: Use simple, direct language
3. **Consistent Format**: Follow the template structure
4. **Actionable**: Tell users what they need to know or do

### Section Guidelines

#### 🎉 What's New
- Focus on user-visible features and improvements
- Explain the benefit to users
- Use present tense
- Be specific about what changed

**Good Examples:**
- Added Docker image layer analysis with detailed file listings
- Implemented web interface for interactive image exploration
- Added support for multi-platform Docker image builds

**Avoid:**
- "Fixed various bugs" (too vague)
- "Improved performance" (not specific enough)

#### 🐛 Bug Fixes
- Describe the problem that was fixed
- Mention the impact on users
- Include issue numbers if applicable

**Good Examples:**
- Fixed crash when analyzing images with empty layers
- Resolved memory leak in large image processing
- Fixed incorrect file permissions in extracted layers

#### 🔧 Improvements
- Describe enhancements to existing features
- Focus on user benefits
- Be specific about improvements

**Good Examples:**
- Improved error messages for better debugging
- Enhanced performance for large image analysis
- Added progress indicators for long-running operations

#### ⚠️ Breaking Changes
- Clearly explain what changed
- Provide migration steps if needed
- Use warning emoji to draw attention
- Explain why the change was necessary

**Good Examples:**
- Changed default port from 8080 to 9090 (use `--port` flag to override)
- Renamed `--image` flag to `--target` for clarity
- Removed deprecated `--legacy` option

#### 📦 Dependencies
- List significant dependency updates
- Mention security updates
- Note any compatibility changes

### Version Numbering

#### Semantic Versioning
- **Major.Minor.Patch** (e.g., 1.0.0, 1.2.3)
- **Major**: Breaking changes
- **Minor**: New features, backward compatible
- **Patch**: Bug fixes, backward compatible

#### Development Versions
- **Alpha**: Early development, major features incomplete
  - Format: `v1.0.0-alpha.1`, `v1.0.0-alpha.2`
- **Beta**: Feature complete, testing phase
  - Format: `v1.0.0-beta.1`, `v1.0.0-beta.2`

### Examples

#### Production Release Example

```markdown
## v1.2.0 - 2024-01-15

### 🎉 What's New
- Added support for Docker image vulnerability scanning
- Implemented image comparison tool for detecting differences
- Added export functionality for analysis reports

### 🐛 Bug Fixes
- Fixed memory leak when processing large images (#123)
- Resolved crash on Windows when accessing Docker socket (#124)
- Fixed incorrect file size reporting in layer analysis

### 🔧 Improvements
- Improved error handling for network timeouts
- Enhanced progress reporting for long operations
- Added better validation for Docker registry credentials

### 📦 Dependencies
- Updated Docker API client to v28.3.0
- Updated Cobra CLI framework to v1.9.1

### 🔗 Links
- [GitHub Release](https://github.com/smiller333/dockerutilities/releases/tag/v1.2.0)
- [Download Binaries](https://github.com/smiller333/dockerutilities/releases/tag/v1.2.0)
```

#### Development Release Example

```markdown
## v1.3.0-alpha.1 - 2024-01-10 - Alpha Release

### 🧪 Development Release
This is an alpha release for testing the new Docker Compose analysis features. Not recommended for production use.

### 🎉 What's New
- Added Docker Compose file analysis and validation
- Implemented multi-service dependency mapping
- Added support for environment variable resolution

### 🐛 Bug Fixes
- Fixed issue with circular dependency detection
- Resolved memory usage in large compose files

### ⚠️ Known Issues
- Environment variable substitution may fail with complex expressions
- Large compose files (>100 services) may cause performance issues
- Windows path handling needs additional testing

### 📝 Testing Notes
- Test with various Docker Compose file formats
- Verify environment variable resolution works correctly
- Check performance with large compose projects

### 🔗 Links
- [GitHub Release](https://github.com/smiller333/dockerutilities/releases/tag/v1.3.0-alpha.1)
- [Download Binaries](https://github.com/smiller333/dockerutilities/releases/tag/v1.3.0-alpha.1)
```

## Checklist for Release Notes

Before publishing release notes, ensure:

- [ ] All sections are filled out appropriately
- [ ] User-focused language is used throughout
- [ ] Breaking changes are clearly explained with migration steps
- [ ] Links to GitHub release and binaries are included
- [ ] Version number and date are correct
- [ ] Development releases are clearly marked as such
- [ ] Known issues are documented for development releases
- [ ] Testing notes are provided for development releases

## Tips for Good Release Notes

1. **Start Early**: Begin collecting notes during development
2. **Be Consistent**: Use the same format and style every time
3. **Think Like a User**: Focus on what users need to know
4. **Keep It Simple**: Avoid technical jargon when possible
5. **Be Honest**: Don't hide known issues or limitations
6. **Provide Context**: Explain why changes were made when relevant
7. **Use Emojis**: They help organize and make notes more readable
8. **Link Everything**: Provide links to issues, PRs, and downloads

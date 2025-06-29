# Copilot Instructions for VSCode

This document configures GitHub Copilot for a Go-based command-line utilities project that interacts with the Docker Engine API using the `github.com/docker/docker/client` package (version v28.3.0+incompatible). The project enforces separation of concerns, with all implementation logic in the `src/` directory organized into unique packages, adhering to Go best practices. Documentation for the main client is in `github-com_docker_docker_v28-3-0+incompatible_client.md`, with related type definitions (e.g., `container.Config`, `swarm.ServiceSpec`) in separate files within the same directory or module.

Reference: [GitHub Copilot Documentation](https://docs.github.com/en/copilot/configuring-github-copilot/configuring-github-copilot-in-your-ide)

## Copilot Settings

### Enabled Languages
Configure which languages Copilot is enabled for:

- **Go**: Enabled
- **JavaScript**: Disabled
- **TypeScript**: Disabled
- **Python**: Disabled
- **Java**: Disabled
- **HTML**: Disabled
- **CSS**: Disabled
- **Markdown**: Enabled (for documentation)
- **YAML**: Enabled (for configuration files)

*Note*: Copilot is focused on Go and documentation-related files to align with the project's scope.

### Suggestions
Control how Copilot provides suggestions:

- **Inline Suggestions**: Enabled
- **Suggestion Delay**: 150ms (slightly longer for Go's static typing)
- **Minimum Characters for Suggestions**: 3
- **Suggestions in Comments**: Enabled (useful for TODOs and documentation)
- **Suggestions in Strings**: Disabled (avoids suggestions in string literals)

### Autocompletion
Configure autocompletion behavior:

- **Enabled**: True
- **Line Completions**: Enabled
- **Block Completions**: Enabled
- **Confidence Threshold**: 0.8 (higher threshold for Go's strict conventions)

### Contextual Awareness
Enhance Copilot's understanding of the project:

- **Use Context**: Enabled
- **Context Lines**: 50 (sufficient for Go's concise code)
- **Include Imports**: Enabled (critical for Go's explicit imports)
- **Include Nearby Files**: Enabled (includes related `src/` package files)
- **Context Directories**:
  - `src/**` (implementation logic)
  - `cmd/**` (CLI command definitions)
  - `docs/apis/dockersdk` (root directory for Docker API documentation files, e.g., `github-com_docker_docker_v28-3-0+incompatible_client.md`)

### File Exclusions
Exclude specific file patterns from Copilot suggestions:

- `*.log` (log files)
- `vendor/**` (vendored dependencies)
- `testdata/**` (test data files)
- `*.pb.go` (generated Protobuf files)
- `*.gen.go` (other generated Go files)
- `go.sum` (Go module checksums)

### Docker API Type References
When generating code, documentation, or tests involving the Docker Engine API (`github.com/docker/docker/client`):

- Use `github-com_docker_docker_v28-3-0+incompatible_client.md` as the primary reference for client methods and types.
- Assume related type definitions (e.g., `container.Config`, `swarm.ServiceSpec`, `network.CreateOptions`) are in separate markdown files within the same directory.
- Infer type details from import paths (e.g., `github.com/docker/docker/container` for `container.Config`).
- Infer documentation file name from the import paths (e.g., `github.com/docker/docker/api/types` will be found in `github-com_docker_docker_v28-3-0+incompatible_api_types.md`).
- If type details are unavailable, make reasonable assumptions based on Go conventions and note them in comments (e.g., `// Assuming container.Config requires Image and Cmd fields`).
- Ensure code uses the correct package imports and follows Go best practices (e.g., explicit error handling, godoc comments).

### Custom Prompts
Define custom prompts for Go-specific tasks:

#### Go
- **Command Implementation**:
  ```
  Write a Go function for a CLI command in the cmd/ directory using the Cobra library. The function should:
  - Be defined in a package under cmd/
  - Follow Go best practices (error handling, idiomatic naming)
  - Include a clear command description and usage
  - Delegate logic to a package in src/
  - Use dependency injection for modularity
  - If interacting with the Docker API, reference github-com_docker_docker_v28-3-0+incompatible_client.md and related type files in the same directory
  Format the code with gofmt and include necessary imports.
  ```

- **Package Logic**:
  ```
  Write a Go function for a package in the src/ directory. The function should:
  - Be part of a specific package (e.g., src/dockerutil)
  - Follow Go best practices (explicit error returns, clear naming)
  - Be exported only if intended for external use
  - Include comprehensive godoc comments
  - Avoid dependencies outside stdlib or explicitly defined packages (e.g., github.com/docker/docker/client)
  - If using Docker API types, reference related files in the same directory for type definitions (e.g., container.Config in container/ package)
  Format the code with gofmt and include necessary imports.
  ```

- **Test Generation**:
  ```
  Write a unit test for the following Go code in a _test.go file. The test should:
  - Use Go's testing package and testify for assertions
  - Cover all exported functions in the package
  - Include table-driven tests where applicable
  - Follow Go naming conventions (e.g., TestFunctionName)
  - Mock dependencies using interfaces
  - If testing Docker API interactions, reference github-com_docker_docker_v28-3-0+incompatible_client.md and related type files for accurate type usage
  Format the code with gofmt and include necessary imports.
  ```

- **Documentation**:
  ```
  Generate godoc-style documentation for the following Go code. Include:
  - Package-level overview describing purpose and usage
  - Detailed comments for all exported types, functions, and constants
  - Example usage code where applicable
  - Clear parameter and return value descriptions
  - For Docker API-related code, reference github-com_docker_docker_v28-3-0+incompatible_client.md and related type files in the same directory
  Ensure comments are formatted per Go conventions and placed above declarations.
  ```

### Debugging and Logging
Configure logging for troubleshooting:

- **Enable Logging**: Disabled
- **Log Level**: warn (Options: `error`, `warn`, `info`, `debug`)
- **Log Path**: `./copilot.log`

### Experimental Features
Control experimental features:

- **Enable Experimental Features**: Disabled
- **Code Explanation**:
  ```
  Provide a detailed explanation of the following Go code snippet. Include:
  - Functionality and purpose
  - Explanation of package structure and imports
  - Highlight adherence to Go best practices
  - Suggest optimizations or idiomatic improvements
  - Identify potential error handling issues
  - For Docker API-related code, reference github-com_docker_docker_v28-3-0+incompatible_client.md and related type files in the same directory for type details
  ```

## Notes
- Place this file in the project root or `.copilot` directory.
- The configuration prioritizes Go-specific conventions and the project structure (`cmd/`, `src/`, and Docker API documentation files).
- Update prompts and settings as the project evolves or Copilot features change.
- Ensure sensitive data (e.g., API keys) is excluded from prompts and logs.
- Use with Go modules and ensure `go.mod` is properly configured.
- When referencing Docker API types, Copilot should infer details from related files in the same directory, using `github-com_docker_docker_v28-3-0+incompatible_client.md` as the primary client reference.
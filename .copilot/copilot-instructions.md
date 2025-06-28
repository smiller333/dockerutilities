# Copilot Instructions for VSCode

This document configures GitHub Copilot for a Go-based command-line utilities project. The project enforces separation of concerns, with all implementation logic in the `src/` directory organized into unique packages, adhering to Go best practices.

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

### File Exclusions
Exclude specific file patterns from Copilot suggestions:

- `*.log` (log files)
- `vendor/**` (vendored dependencies)
- `testdata/**` (test data files)
- `*.pb.go` (generated Protobuf files)
- `*.gen.go` (other generated Go files)
- `go.sum` (Go module checksums)

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
  Format the code with gofmt and include necessary imports.
  ```

- **Package Logic**:
  ```
  Write a Go function for a package in the src/ directory. The function should:
  - Be part of a specific package (e.g., src/stringutil)
  - Follow Go best practices (explicit error returns, clear naming)
  - Be exported only if intended for external use
  - Include comprehensive godoc comments
  - Avoid dependencies outside stdlib or explicitly defined packages
  Format the code with gofmt and include necessary imports.
  ```

- **Test Generation**:
  ```
  Write a unit test for the following Go code in a _test.go file. The test should:
  - Use only Go's standard testing package (no external assert libraries)
  - Use table-driven tests with t.Run() for all test cases
  - Use t.Errorf() and t.Fatalf() for assertions and error reporting
  - Cover all exported functions in the package
  - Follow Go naming conventions (e.g., TestFunctionName)
  - Mock dependencies using interfaces
  - Include clear test case names and expected vs actual value comparisons
  Format the code with gofmt and include necessary imports.
  ```

- **Documentation**:
  ```
  Generate godoc-style documentation for the following Go code. Include:
  - Package-level overview describing purpose and usage
  - Detailed comments for all exported types, functions, and constants
  - Example usage code where applicable
  - Clear parameter and return value descriptions
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
  ```

## Notes
- Place this file in the project root or `.copilot` directory.
- The configuration prioritizes Go-specific conventions and the project structure (`cmd/`, `src/`).
- Update prompts and settings as the project evolves or Copilot features change.
- Ensure sensitive data (e.g., API keys) is excluded from prompts and logs.
- Use with Go modules and ensure `go.mod` is properly configured.
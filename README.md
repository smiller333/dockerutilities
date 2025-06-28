# Docker Utilities

A collection of Docker utilities providing a command-line interface for various Docker-related tasks.

## Overview

`dockerutils` is a CLI tool built in Go that aims to simplify common Docker operations and provide additional functionality for Docker workflows.

## Installation

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/smiller333/dockerutils.git
   cd dockerutils
   ```

2. Build the binary:
   ```bash
   go build -o dockerutils
   ```

3. (Optional) Move to your PATH:
   ```bash
   mv dockerutils /usr/local/bin/
   ```

## Usage

### Basic Commands

```bash
# Display help
dockerutils --help

# Show version
dockerutils version
```

### Available Commands

- `version` - Print the version number of dockerutils

## Development

### Prerequisites

- Go 1.24.2 or later
- Git

### Running Tests

```bash
go test ./...
```

### Project Structure

```
dockerutils/
├── cmd/                 # Command definitions
│   └── root.go         # Root command and CLI setup
├── src/
│   └── version/        # Version management
│       ├── version.go  # Version constants and functions
│       └── version_test.go
├── main.go             # Application entry point
├── go.mod              # Go module definition
└── README.md
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework for Go

## License

This project is open source.

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## Version

Current version: v0.0.1
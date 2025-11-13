# Copilot Instructions for Core Element Template

## Project Overview

This repository is a template for creating custom HTML elements for the core web3 framework. It consists of two main components:

1. **Go Backend (CLI)**: A command-line interface built with Go and Cobra framework
2. **Angular Frontend**: A custom HTML element built with Angular Elements

## Technology Stack

- **Backend**: Go 1.24.3 with Cobra CLI framework
- **Frontend**: Angular 20.3+ with Angular Elements
- **Build Tools**: GoReleaser for Go builds, Angular CLI for frontend builds
- **Package Management**: Go modules for Go, npm for Node.js/Angular

## Project Structure

```
.
├── cmd/demo-cli/          # Go CLI application
│   ├── cmd/               # Cobra command implementations
│   └── main.go            # Application entry point
├── ui/                    # Angular custom element
│   ├── src/               # Angular source code
│   ├── public/            # Static assets
│   └── package.json       # Node.js dependencies
├── .github/               # GitHub configurations and workflows
└── .goreleaser.yaml       # Release configuration
```

## Development Setup

### Prerequisites

- Go 1.24.3 or later
- Node.js and npm
- Git

### Initial Setup

1. Install Go dependencies:
   ```bash
   go mod tidy
   ```

2. Install Node.js dependencies:
   ```bash
   cd ui
   npm install
   ```

### Running the Application

Start the development server:
```bash
go run ./cmd/demo-cli serve
```

This starts the Go backend and serves the Angular custom element.

## Building

### Build Go CLI

The Go CLI is built using GoReleaser in the release workflow. For local development:
```bash
go build -o demo-cli ./cmd/demo-cli
```

### Build Angular Custom Element

```bash
cd ui
npm run build
```

This creates a single JavaScript file in the `dist` directory for use in any HTML page.

## Testing

### Go Tests

Run Go tests:
```bash
go test ./...
```

Note: Currently, there are no test files in the Go codebase.

### Angular Tests

Run Angular tests:
```bash
cd ui
npm test
```

## Code Style and Conventions

### Go Code

- Follow standard Go conventions and formatting (use `gofmt`)
- Use meaningful variable and function names
- Keep functions focused and single-purpose
- Use Cobra framework patterns for CLI commands

### Angular/TypeScript Code

- Follow Angular style guide
- Use Prettier for code formatting (configuration in `ui/package.json`)
- Settings:
  - Print width: 100 characters
  - Single quotes preferred
  - Angular parser for HTML files

### General Guidelines

- Write clear, self-documenting code
- Add comments for complex logic or non-obvious decisions
- Keep commits atomic and well-described
- Follow the existing code patterns in the repository

## Important Notes for AI Assistants

1. **Minimal Changes**: Make the smallest possible changes to achieve the goal
2. **No Breaking Changes**: Don't modify working code unless necessary
3. **Testing**: Run tests before and after changes to ensure nothing breaks
4. **Build Verification**: Always verify builds succeed after making changes
5. **Dependencies**: Check for security vulnerabilities before adding new dependencies
6. **Module Names**: Note that the Go module name uses a placeholder `github.com/your-username/core-element-template` which should be updated when forking

## Workflows

- **Release**: Automated release workflow using GoReleaser (`.github/workflows/release.yml`)

## Contributing

When contributing to this repository:

1. Ensure all tests pass
2. Follow the existing code style
3. Keep changes minimal and focused
4. Update documentation if needed
5. Test the build locally before submitting

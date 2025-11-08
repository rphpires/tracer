# Contributing to Tracer

Thank you for your interest in contributing to Tracer! This document provides guidelines and instructions for contributing.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/yourusername/tracer.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Run tests: `go test -v`
6. Commit your changes: `git commit -am "Add your feature"`
7. Push to your fork: `git push origin feature/your-feature-name`
8. Create a Pull Request

## Development Guidelines

### Code Style

- Follow the [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use `gofmt` to format your code: `go fmt ./...`
- Run `go vet` to check for common mistakes: `go vet ./...`
- Add comments for all exported functions, types, and constants
- Keep functions focused and concise

### Testing

- Write tests for new features
- Ensure all tests pass: `go test -v`
- Aim for good test coverage: `go test -cover`
- Test files should be named `*_test.go`
- Use table-driven tests when appropriate

### Commit Messages

- Use clear and descriptive commit messages
- Start with a verb in the present tense (e.g., "Add", "Fix", "Update")
- Keep the first line under 50 characters
- Add detailed description if needed after a blank line

Example:
```
Add TraceWithColorf function

Implements a new function that allows formatted logging with
custom colors, similar to fmt.Printf() functionality.
```

## Running Tests

```bash
# Run all tests
go test -v

# Run tests with coverage
go test -v -cover

# Run specific test
go test -v -run TestTracef
```

## Building

```bash
# Build the example
cd example
go build

# Run the example
./example
```

## Documentation

- All exported functions must have godoc comments
- Comments should start with the function name
- Include examples in comments when helpful
- Update README.md if adding new features

## Pull Request Process

1. Ensure your code follows the style guidelines
2. Add or update tests for your changes
3. Update documentation (README.md, godoc comments)
4. Ensure all tests pass
5. Create a descriptive Pull Request
6. Link any related issues

## Reporting Bugs

When reporting bugs, please include:
- Go version (`go version`)
- Operating system
- Steps to reproduce
- Expected behavior
- Actual behavior
- Any error messages or logs

## Feature Requests

Feature requests are welcome! Please:
- Describe the feature clearly
- Explain why it would be useful
- Provide examples of how it would be used
- Consider if it fits the project's scope

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers
- Focus on constructive feedback
- Help others learn and grow

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

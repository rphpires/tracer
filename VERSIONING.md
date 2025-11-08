# Versioning Guide

This document explains how to version and release the Tracer package.

## Semantic Versioning

This project follows [Semantic Versioning 2.0.0](https://semver.org/):

- **MAJOR** version (v1.0.0): Incompatible API changes
- **MINOR** version (v1.1.0): New functionality, backwards compatible
- **PATCH** version (v1.0.1): Bug fixes, backwards compatible

## Creating a Release

### 1. Ensure Code Quality

Before creating a release:

```bash
# Format code
go fmt ./...

# Check for issues
go vet ./...

# Run all tests
go test -v

# Check test coverage
go test -cover
```

### 2. Update CHANGELOG (if you have one)

Document all changes in the release.

### 3. Create a Git Tag

```bash
# For the first release
git tag v1.0.0

# For subsequent releases
git tag v1.1.0  # Minor version
git tag v1.0.1  # Patch version
git tag v2.0.0  # Major version
```

### 4. Push the Tag to GitHub

```bash
# Push specific tag
git push origin v1.0.0

# Push all tags
git push --tags
```

### 5. Create a GitHub Release (Optional but Recommended)

1. Go to your repository on GitHub
2. Click on "Releases"
3. Click "Create a new release"
4. Select the tag you just created
5. Add release notes describing the changes
6. Publish the release

## Version History Recommendations

### v1.0.0 - First Stable Release
When to release:
- All core functionality is implemented and tested
- API is stable and unlikely to change
- Documentation is complete
- No known critical bugs

### v1.x.0 - Minor Releases
When to release:
- New features added
- Improvements to existing functionality
- Backwards compatible changes
- New optional parameters

Examples:
- Adding new logging functions
- Adding configuration options
- Performance improvements

### v1.0.x - Patch Releases
When to release:
- Bug fixes
- Documentation updates
- Security patches
- Performance optimizations (without API changes)

### v2.0.0 - Major Releases
When to release:
- Breaking API changes
- Major refactoring
- Removing deprecated features
- Significant architecture changes

## How Users Install Specific Versions

```bash
# Install latest version
go get github.com/rphpires/tracer@latest

# Install specific version
go get github.com/rphpires/tracer@v1.0.0

# Install specific commit
go get github.com/rphpires/tracer@abc1234

# Install from specific branch
go get github.com/rphpires/tracer@main
```

## Checking Current Version

Users can check which version they're using:

```bash
# In their project
go list -m github.com/rphpires/tracer
```

## Suggested First Release

For your first release, I recommend:

```bash
# Commit all changes
git add .
git commit -m "Prepare for v1.0.0 release"
git push

# Create and push tag
git tag v1.0.0 -m "First stable release

Features:
- HTML-based logging with colors
- Automatic log rotation
- Formatted logging (Tracef, TraceWithColorf)
- Panic recovery
- Thread-safe operations
- Configurable settings
"

git push origin v1.0.0
```

## Version Numbering Examples

| Change | Old Version | New Version |
|--------|-------------|-------------|
| Add new function | v1.0.0 | v1.1.0 |
| Fix bug | v1.0.0 | v1.0.1 |
| Change function signature | v1.0.0 | v2.0.0 |
| Add optional parameter | v1.0.0 | v1.1.0 |
| Remove deprecated function | v1.9.0 | v2.0.0 |
| Improve performance | v1.0.0 | v1.0.1 |

## Pre-release Versions

For testing before official release:

```bash
git tag v1.0.0-beta.1
git tag v1.0.0-rc.1  # Release candidate
```

Users can install pre-release versions:

```bash
go get github.com/rphpires/tracer@v1.0.0-beta.1
```

## Best Practices

1. **Never delete tags** - Once published, tags should be permanent
2. **Never modify tags** - Create a new version instead
3. **Document breaking changes** - Clearly communicate API changes
4. **Test before tagging** - Ensure all tests pass
5. **Use annotated tags** - Include release notes in tag message

## Automated Versioning (Advanced)

You can use GitHub Actions or other CI/CD tools to automate:
- Running tests before allowing tags
- Automatically creating GitHub releases
- Generating changelogs
- Publishing documentation

## Summary

For your project right now:

1. Make final commits
2. Run tests: `go test -v`
3. Create tag: `git tag v1.0.0`
4. Push tag: `git push origin v1.0.0`
5. Users can then install: `go get github.com/rphpires/tracer@v1.0.0`

That's it! No version files needed - Git tags handle everything.

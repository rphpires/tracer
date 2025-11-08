# Quick Release Guide

This is a quick reference for creating your first release of the Tracer package.

## Pre-Release Checklist

✅ All tests pass: `go test -v`
✅ Code is formatted: `go fmt ./...`
✅ No vet issues: `go vet ./...`
✅ Documentation is up to date
✅ Examples work correctly

## Creating Your First Release (v1.0.0)

### Step 1: Final Commit

```bash
# Make sure all changes are committed
git add .
git commit -m "Prepare for v1.0.0 release"
git push origin main
```

### Step 2: Create Version Tag

```bash
# Create annotated tag with release notes
git tag -a v1.0.0 -m "Release v1.0.0 - First stable release

Features:
- HTML-based logging with colors
- Automatic log rotation (5MB default)
- Formatted logging (Tracef, TraceWithColorf)
- Panic recovery and logging
- Thread-safe operations
- Configurable settings
- Interactive JavaScript log filtering
- Enable/disable via configuration files

Full documentation: https://github.com/rphpires/tracer
"
```

### Step 3: Push Tag to GitHub

```bash
git push origin v1.0.0
```

### Step 4: Verify Tag

```bash
# List all tags
git tag

# Show tag details
git show v1.0.0
```

### Step 5: Create GitHub Release (Optional but Recommended)

1. Go to: https://github.com/rphpires/tracer/releases
2. Click "Create a new release"
3. Select tag: `v1.0.0`
4. Title: `v1.0.0 - First Stable Release`
5. Description: Copy the release notes from the tag
6. Click "Publish release"

## Users Can Now Install Your Package

```bash
# Latest version
go get github.com/rphpires/tracer@latest

# Specific version
go get github.com/rphpires/tracer@v1.0.0
```

## Future Releases

### Bug Fix (Patch) - v1.0.1

```bash
git tag -a v1.0.1 -m "Release v1.0.1 - Bug fixes

- Fixed issue with log rotation
- Improved error handling
"
git push origin v1.0.1
```

### New Feature (Minor) - v1.1.0

```bash
git tag -a v1.1.0 -m "Release v1.1.0 - New features

- Added new TraceDebug function
- Added configuration for custom folder names
"
git push origin v1.1.0
```

### Breaking Change (Major) - v2.0.0

```bash
git tag -a v2.0.0 -m "Release v2.0.0 - Major update

Breaking Changes:
- Changed Config struct fields
- Renamed some functions for clarity

New Features:
- JSON logging support
- Custom formatters
"
git push origin v2.0.0
```

## Verifying Your Release

After creating a tag, test installation in a new project:

```bash
# Create test project
mkdir /tmp/test-tracer
cd /tmp/test-tracer
go mod init test
go get github.com/rphpires/tracer@v1.0.0

# Verify version
go list -m github.com/rphpires/tracer
```

## Documentation Updates

After your first release, your package will appear on:
- https://pkg.go.dev/github.com/rphpires/tracer
- https://goreportcard.com/report/github.com/rphpires/tracer

These sites automatically index new Go packages from GitHub.

## Quick Commands Reference

```bash
# View all tags
git tag -l

# Delete local tag (if needed before pushing)
git tag -d v1.0.0

# Delete remote tag (use with caution!)
git push origin :refs/tags/v1.0.0

# View tag information
git show v1.0.0

# List all releases
gh release list  # requires GitHub CLI
```

## Summary

The complete release process is:

1. ✅ Run tests: `go test -v`
2. ✅ Commit changes: `git commit -am "Your message"`
3. ✅ Push changes: `git push`
4. ✅ Create tag: `git tag -a v1.0.0 -m "Release notes"`
5. ✅ Push tag: `git push origin v1.0.0`
6. ✅ Create GitHub release (optional)

That's it! 🚀

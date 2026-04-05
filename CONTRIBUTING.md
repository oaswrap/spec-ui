# Contributing

Thanks for your interest in contributing to spec-ui.

## Development Setup

1. Fork and clone the repository.
2. Ensure you have Go installed (the project uses Go modules).
3. Install dependencies:

```bash
go mod tidy
```

## Run Checks Locally

Before opening a pull request, run:

```bash
go test ./...
```

If you have golangci-lint installed:

```bash
golangci-lint run
```

## Embedded Assets Notes

For library users, embedded UI assets are already included in this module.

For maintainers only: if you intentionally update bundled provider assets (CSS/JS/favicon files), regenerate them with:

```bash
make download-assets
```

Then run tests again:

```bash
go test ./...
```

## Pull Requests

- Keep changes focused and scoped to one concern when possible.
- Add or update tests for behavior changes.
- Update documentation (README or this file) when behavior/config changes.
- Use clear commit messages.

## Issues and Discussions

For major changes, please open an issue first so we can discuss approach and compatibility.

## Code of Conduct

Please be respectful and constructive in all interactions.

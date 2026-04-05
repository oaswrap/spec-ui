# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run all tests
go test ./...

# Run a single test
go test -run TestHandler/Docs/SwaggerUI ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...

# Lint (requires golangci-lint)
golangci-lint run

# Tidy dependencies
go mod tidy
```

## Architecture

`spec-ui` is a Go library that serves OpenAPI documentation UIs as HTTP handlers. It supports five UI providers: SwaggerUI, StoplightElements, ReDoc, Scalar, and RapiDoc.

### Data flow

1. The user calls `specui.NewHandler(opts...)` with functional options from `option.go`
2. Options populate a `config.SpecUI` struct (defined in `config/config.go`)
3. `Handler.Docs()` dispatches to the appropriate provider package under `internal/` based on `cfg.Provider`
4. `Handler.Spec()` always dispatches to `internal/spec`, which serves the raw OpenAPI file

### Key structural patterns

- **Provider packages** (`internal/swaggerui`, `internal/stoplightelements`, `internal/redoc`, `internal/scalar`, `internal/rapidoc`): each has a `handler.go` that renders an HTML template, and an `index.tpl.go` that holds the template string. All UI assets are loaded from CDN (URLs pinned in `internal/constant/constant.go`).

- **Spec serving** (`internal/spec/spec.go`): uses `sync.Once` to read the spec file once and cache it in memory. Supports four source modes: `SpecGenerator` interface, `embed.FS`, `fs.FS`, or plain OS file path.

- **Config** (`config/config.go`): a single `SpecUI` struct holds all configuration. Each UI provider has its own nested config struct with typed constants for enum-like fields (e.g., `ElementLayout`, `RapiDocTheme`).

- **Default provider** is `ProviderStoplightElements`; default paths are `/docs` (docs) and `/docs/openapi.json` (spec).

### Adding a new UI provider

1. Add a new `Provider` constant to `config/config.go`
2. Add the provider config struct to `config/config.go`
3. Create `internal/<provider>/handler.go` and `internal/<provider>/index.tpl.go`
4. Add CDN asset URLs to `internal/constant/constant.go`
5. In `option.go`: import the new internal package, add a `WithXxx` option that sets `c.Provider`, `c.DocsHandlerFactory` (closure wrapping `<pkg>.NewHandler(c)`), and any provider-specific config defaults. `handler.go:Docs()` needs no changes — it dispatches via `DocsHandlerFactory`.

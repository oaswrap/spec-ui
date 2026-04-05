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

`spec-ui` is a Go library that serves OpenAPI documentation UIs as HTTP handlers. It supports five UI providers: SwaggerUI, StoplightElements, ReDoc, Scalar, and RapiDoc. Each provider comes in two variants: CDN mode and embedded mode.

### Data flow

1. The user calls `specui.NewHandler(opts...)` with functional options from `option.go` and a provider-specific `WithUI()` option
2. Options populate a `config.SpecUI` struct (defined in `config/config.go`)
3. `Handler.Docs()` calls `cfg.DocsHandlerFactory` — set by the provider's `WithUI()` — to create the handler once (via `sync.Once`)
4. `Handler.Assets()` calls `cfg.AssetsHandlerFactory` for embedded asset serving; returns `nil` in CDN mode
5. `Handler.Spec()` dispatches to `internal/spec`, which serves the raw OpenAPI file

### Provider packages

Each provider lives in two top-level packages:

- **`<provider>/`** (e.g., `swaggerui/`, `rapidoc/`): CDN mode — loads JS/CSS from pinned CDN URLs in `internal/constant/constant.go`. Contains `handler.go`, `index.tpl.go`, and `option.go`.
- **`<provider>emb/`** (e.g., `swaggeruiemb/`, `rapidocemb/`): Embedded mode — sets `cfg.EmbedAssets = true` and delegates to the CDN handler, but also registers `AssetsHandlerFactory` to serve files from the package's `assets/` directory via `//go:embed assets`.

Each provider package exposes `WithUI(cfg ...config.<Provider>) specui.Option`. Only the imported provider's code is linked into the binary (tree-shaking via factory pattern).

### Key structural patterns

- **`option.go` (root)**: Shared options only — `WithTitle`, `WithDocsPath`, `WithSpecPath`, `WithAssetsPath`, `WithSpecFile`, `WithSpecEmbedFS`, `WithSpecIOFS`, `WithSpecGenerator`, `WithCacheAge`. Provider selection is done by importing the provider package.

- **`config/config.go`**: A single `SpecUI` struct holds all configuration. `DocsHandlerFactory` and `AssetsHandlerFactory` are function fields set by `WithUI()`. Each provider has its own typed config struct with enum-like constants (e.g., `SwaggerLayout`, `RapiDocTheme`).

- **`internal/spec/spec.go`**: Uses `sync.Once` to read the spec file once and cache it. Supports four source modes: `SpecGenerator` interface, `embed.FS`, `fs.FS`, or plain OS file path.

- **`handler.go`**: `Docs()`, `Spec()`, `Assets()` — the three HTTP handlers. Call `handler.AssetsEnabled()` to check if embedded mode is active; if so, register `handler.Assets()` at `handler.AssetsPath() + "/"`.

### Adding a new UI provider

1. Add a new `Provider` constant to `config/config.go`
2. Add the provider config struct to `config/config.go`
3. Create `<provider>/handler.go`, `<provider>/index.tpl.go`, `<provider>/option.go` — `option.go` sets `DocsHandlerFactory`, `AssetsHandlerFactory = nil`, and any config defaults
4. Add CDN asset URLs to `internal/constant/constant.go`
5. Create `<provider>emb/handler.go`, `<provider>emb/assets.go`, `<provider>emb/option.go` — `assets.go` uses `//go:embed assets` and serves files; `option.go` sets `EmbedAssets = true` and registers both factories

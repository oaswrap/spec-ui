# Spec UI

[![Go Reference](https://pkg.go.dev/badge/github.com/oaswrap/spec-ui.svg)](https://pkg.go.dev/github.com/oaswrap/spec-ui)
[![codecov](https://codecov.io/gh/oaswrap/spec-ui/graph/badge.svg?token=RqYnxmf6mW)](https://codecov.io/gh/oaswrap/spec-ui)
[![Go Report Card](https://goreportcard.com/badge/github.com/oaswrap/spec-ui)](https://goreportcard.com/report/github.com/oaswrap/spec-ui)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/oaswrap/spec-ui/blob/main/LICENSE)

A Go library that provides multiple OpenAPI documentation UIs. Serve beautiful, interactive API documentation for your OpenAPI specifications.

## Features

- 🚀 **Multiple UI Options**: Support for Swagger UI, Stoplight Elements, ReDoc, Scalar and RapiDoc
- ⚡ **Easy Integration**: Simple HTTP handler integration with Go's standard library
- 🎨 **Customizable**: Configure titles, base paths, and OpenAPI spec locations
- 🔧 **Flexible**: Works with any Go HTTP router or framework

## Installation

```bash
go get github.com/oaswrap/spec-ui
```

## Quick Start

```go
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

func main() {
	r := chi.NewRouter()

	// Stoplight Elements
	handler := specui.NewHandler(
		specui.WithTitle("Pet Store API"),
		specui.WithDocsPath("/docs"),
		specui.WithSpecPath("/docs/openapi.yaml"),
		specui.WithSpecFile("openapi.yaml"),
		specui.WithStoplightElements(),
	)

	r.Get(handler.DocsPath(), handler.DocsFunc())
	r.Get(handler.SpecPath(), handler.SpecFunc())

	log.Printf("OpenAPI Documentation available at http://localhost:3000/docs")
	log.Printf("OpenAPI YAML available at http://localhost:3000/docs/openapi.yaml")

	http.ListenAndServe(":3000", r)
}
```

## UI Options

### Stoplight Elements
**Beautiful Three-Column Design** - Modern API documentation with a "Stripe-esque" three-column layout, powered by OpenAPI and Markdown for an elegant developer experience.  
[View Demo](https://elements-demo.stoplight.io/?spec=https://petstore3.swagger.io/api/v3/openapi.json)

```go
handler := specui.NewHandler(
	specui.WithTitle("My API"),
	specui.WithSpecFile("openapi.yaml"),
	specui.WithStoplightElements(config.StoplightElements{
		HideExport:  false,
		HideSchemas: false,
		HideTryIt:   false,
		Layout:      "sidebar",
		Logo:        "/assets/logo.png",
		Router:      "hash",
	}),
)
```

### Swagger UI
**Interactive API Explorer** - Your interactive coding buddy for AI-driven API testing workflows, widely used by developers with extensive framework integrations.  
[View Demo](https://petstore3.swagger.io/)

```go
handler := specui.NewHandler(
	specui.WithTitle("My API"),
	specui.WithSpecFile("openapi.yaml"),
	specui.WithSwaggerUI(),
)
```

### ReDoc
**Modern Three-Column Layout** - Similar to Swagger UI but renders documentation in a modern three-column format, perfect for polished executive summaries and presenting API schemas.  
[View Demo](https://redocly.github.io/redoc/?url=https://petstore3.swagger.io/api/v3/openapi.json)

```go
handler := specui.NewHandler(
	specui.WithTitle("My API"),
	specui.WithSpecFile("openapi.yaml"),
	specui.WithReDoc(),
)
```

### Scalar
**Feature-Rich Modern Interface** - Provides the most feature-rich interface compared to Swagger UI and ReDoc, with built-in themes, search function, and code examples.  
[View Demo](https://docs.scalar.com/swagger-editor)

```go
handler := specui.NewHandler(
	specui.WithTitle("My API"),
	specui.WithSpecFile("openapi.yaml"),
	specui.WithScalar(),
)
```

### RapiDoc
**Flexible Rendering Styles** - Web component-based viewer with multiple themes, styling options, and distinctive tabular/tree model representations perfect for large schemas.  
[View Demo](https://rapidocweb.com/examples/petstore-extended.html)

```go
handler := specui.NewHandler(
	specui.WithTitle("My API"),
	specui.WithSpecFile("openapi.yaml"),
	specui.WithRapiDoc(),
)
```

## Handler Methods

The handler provides convenient methods for integration:

- `handler.Docs()` - Returns HTTP handler for the documentation UI
- `handler.DocsFunc()` - Returns the HTTP handler function for the documentation UI
- `handler.DocsPath()` - Returns the documentation path (e.g., `/docs`)
- `handler.Spec()` - Returns HTTP handler for the OpenAPI specification
- `handler.SpecFunc()` - Returns the HTTP handler function for serving the OpenAPI specification
- `handler.SpecPath()` - Returns the OpenAPI spec path (e.g., `/docs/openapi.yaml`)

## Basic Usage

The API uses a builder pattern with functional options for flexible configuration:

```go
// Complete example with all available options
handler := specui.NewHandler(
	specui.WithTitle("My API"),                    // Set documentation title
	specui.WithDocsPath("/docs"),                  // Set docs URL path
	specui.WithSpecPath("/docs/openapi.yaml"),     // Set spec URL path
	specui.WithSpecFile("openapi.yaml"),           // Set the spec file location
	specui.WithStoplightElements(config.StoplightElements{  // Choose UI with config
		HideExport: false,
		HideTryIt:  false,
	}),
)

// Minimal setup (uses sensible defaults)
handler := specui.NewHandler(
	specui.WithSpecFile("openapi.yaml"),
	specui.WithSwaggerUI(), // No config needed, uses defaults
)

// Register with any HTTP router
r.Get(handler.DocsPath(), handler.DocsFunc())   // Documentation UI
r.Get(handler.SpecPath(), handler.SpecFunc())   // OpenAPI spec file
```

## Configuration Options

The library uses functional options for flexible configuration:

### Core Options

| Option | Description | Example |
|--------|-------------|---------|
| `WithTitle` | Set documentation title | `specui.WithTitle("My API")` |
| `WithDocsPath` | Set documentation URL path | `specui.WithDocsPath("/docs")` |
| `WithSpecPath` | Set OpenAPI spec URL path | `specui.WithSpecPath("/docs/openapi.yaml")` |
| `WithSpecFile` | Set the spec file location | `specui.WithSpecFile("openapi.yaml")` |
| `WithSpecEmbedFS` | Set spec file location with embedded filesystem | `specui.WithSpecEmbedFS("openapi.yaml", embedFS)` |
| `WithSpecIOFS` | Set spec file location with OS filesystem | `specui.WithSpecIOFS("openapi.yaml", os.DirFS("docs"))` |
| `WithCacheAge` | Set cache age for the documentation | `specui.WithCacheAge(3600)` |

### UI Selection with Configuration

#### Stoplight Elements Configuration

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `HideExport` | `bool` | `false` | Hide the "Export" button |
| `HideSchemas` | `bool` | `false` | Hide schemas in the Table of Contents |
| `HideTryIt` | `bool` | `false` | Hide the "Try it" interactive feature |
| `HideTryItPanel` | `bool` | `false` | Hide the "Try it" panel |
| `Layout` | `string` | `"sidebar"` | Layout: "sidebar" or "responsive" |
| `Logo` | `string` | `""` | URL to logo image |
| `Router` | `string` | `"history"` | Router type: "history", "hash", "memory", or "static" |

**Usage:**
```go
specui.WithStoplightElements(config.StoplightElements{
	HideExport:		false,
	HideSchemas:	false,
	HideTryIt:		false,
	HideTryItPanel:	false,
	Layout:			"sidebar",
	Logo:			"/assets/logo.png",
	Router:			"hash",
})
```

#### Swagger UI Configuration

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `HideCurl` | `bool` | `false` | Hide curl code snippets |
| `JsonEditor` | `bool` | `false` | Enable visual JSON editor (experimental) |
| `Layout` | `string` | `"StandaloneLayout"` | Layout type: "StandaloneLayout" or "BaseLayout" |
| `DefaultModelsExpandDepth` | `int` | `1` | Default depth for model expansion in the UI |

**Usage:**
```go
specui.WithSwaggerUI(config.SwaggerUI{
	HideCurl:   				false,
	JsonEditor: 				true,
	Layout:     				"StandaloneLayout",
	DefaultModelsExpandDepth:	1,
})
```

#### ReDoc Configuration

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `DisableSearch` | `bool` | `false` | Disable search functionality |
| `HideDownloadButtons` | `bool` | `false` | Hide the "Download" button for saving the API definition source file |
| `HideSchemaTitles` | `bool` | `false` | Hide the schema titles in the documentation |

**Usage:**
```go
specui.WithReDoc(config.ReDoc{
	DisableSearch:       true,
	HideDownloadButtons: true,
	HideSchemaTitles:    true,
})
```

#### Scalar Configuration

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `ProxyURL` | `string` | `""` | Set Proxy URL for making API requests |
| `HideSidebar` | `bool` | `false` | Hide sidebar navigation |
| `HideModels` | `bool` | `false` | Hide models in the sidebar |
| `DocumentDownloadType` | `string` | `"both"` | Document download type: "json", "yaml", "both", or "none" |
| `HideTestRequestButton` | `bool` | `false` | Hide the "Test Request" button |
| `HideSearch` | `bool` | `false` | Hide search bar |
| `DarkMode` | `bool` | `false` | Enable dark mode |
| `Layout` | `string` | `"modern"` | Layout type: "modern" or "classic" |
| `Theme` | `string` | `"default"` | Theme name, see [Scalar themes](https://guides.scalar.com/scalar/scalar-api-references/themes) for available options |

**Usage:**
```go
specui.WithScalar(config.Scalar{
	ProxyURL:                "https://proxy.scalar.com",
	HideSidebar:             false,
	HideModels:              false,
	DocumentDownloadType:    "both",
	HideTestRequestButton:   false,
	HideSearch:              false,
	DarkMode:                true,
	Layout:                  "modern",
	Theme:                   "moon",
})
```

#### RapiDoc Configuration

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Theme` | `string` | `"light"` | Theme style: "light" or "dark" |
| `Layout` | `string` | `"row"` | Layout type: "row" or "column" |
| `RenderStyle` | `string` | `"read"` | Render style: "read", "view", or "focused" |
| `SchemaStyle` | `string` | `"table"` | Schema style: "table" or "tree" |
| `BgColor` | `string` | `"#fff"` | Background color |
| `TextColor` | `string` | `"#444"` | Text color |
| `HeaderColor` | `string` | `"#444444"` | Header color |
| `PrimaryColor` | `string` | `"#FF791A"` | Primary color |
| `HideInfo` | `bool` | `false` | Hide the info section |
| `HideHeader` | `bool` | `false` | Hide the header section |
| `HideSearch` | `bool` | `false` | Hide the search bar |
| `HideAdvancedSearch` | `bool` | `false` | Hide the advanced search bar |
| `HideTryIt` | `bool` | `false` | Hide the "Try" feature |
| `Logo` | `string` | `""` | Logo URL |

**Usage:**
```go
specui.WithRapiDoc(config.RapiDoc{
	Theme:               "light",
	Layout:              "row",
	RenderStyle:         "read",
	SchemaStyle:         "table",
	BgColor:             "#fff",
	TextColor:           "#444",
	HeaderColor:         "#444444",
	PrimaryColor:        "#FF791A",
	HideInfo:            false,
	HideHeader:          false,
	HideSearch:          false,
	HideAdvancedSearch:  false,
	HideTryIt:           false,
	Logo:                "/assets/logo.png",
})
```

## Examples

Check out the [`examples`](/examples) directory for more examples.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License—see the [LICENSE](/LICENSE) file for details.

Made with ❤️ for the Go community
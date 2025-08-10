# Spec UI

[![Go Reference](https://pkg.go.dev/badge/github.com/oaswrap/spec-ui.svg)](https://pkg.go.dev/github.com/oaswrap/spec-ui)
[![Go Report Card](https://goreportcard.com/badge/github.com/oaswrap/spec-ui)](https://goreportcard.com/report/github.com/oaswrap/spec-ui)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/oaswrap/spec-ui/blob/main/LICENSE)

A Go library that provides multiple OpenAPI documentation UIs including Swagger UI, Redoc, and Stoplight Elements. Easily serve beautiful, interactive API documentation for your OpenAPI specifications.

## Features

- 🚀 **Multiple UI Options**: Support for Swagger UI, Redoc, and Stoplight Elements
- 📱 **Responsive Design**: All UIs are mobile-friendly and responsive
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
		specui.WithTitle("Petstore API"),
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
Modern, customizable API documentation with excellent developer experience.

```go
handler := specui.NewHandler(
	specui.WithTitle("My API"),
	specui.WithDocsPath("/docs"),
	specui.WithSpecPath("/docs/openapi.yaml"),
	specui.WithSpecFile("openapi.yaml"),
	specui.WithStoplightElements(config.Elements{
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
The classic, feature-rich OpenAPI documentation interface with interactive API exploration.

```go
handler := specui.NewHandler(
	specui.WithTitle("My API"),
	specui.WithDocsPath("/docs"),
	specui.WithSpecPath("/docs/openapi.yaml"),
	specui.WithSpecFile("openapi.yaml"),
	specui.WithSwaggerUI(config.Swagger{
		ShowTopBar:         true,
		HideCurl:           false,
		JsonEditor:         true,
	}),
)
```

### Redoc
A clean, responsive documentation interface optimized for readability.

```go
handler := specui.NewHandler(
	specui.WithTitle("My API"),
	specui.WithDocsPath("/docs"),
	specui.WithSpecPath("/docs/openapi.yaml"),
	specui.WithSpecFile("openapi.yaml"),
	specui.WithRedoc(config.Redoc{
		HideDownload: false,
	}),
)
```

## Handler Methods

The handler provides convenient methods for integration:

- `handler.DocsPath()` - Returns the documentation path (e.g., `/docs`)
- `handler.DocsFunc()` - Returns the HTTP handler function for the documentation UI
- `handler.SpecPath()` - Returns the OpenAPI spec path (e.g., `/docs/openapi.yaml`)  
- `handler.SpecFunc()` - Returns the HTTP handler function for serving the OpenAPI specification

## Basic Usage

The API uses a builder pattern with functional options for flexible configuration:

```go
// Complete example with all available options
handler := specui.NewHandler(
	specui.WithTitle("My API"),                    // Set documentation title
	specui.WithDocsPath("/docs"),                  // Set docs URL path
	specui.WithSpecPath("/docs/openapi.yaml"),     // Set spec URL path
	specui.WithSpecFile("openapi.yaml"),           // Set spec file location
	specui.WithStoplightElements(config.Elements{  // Choose UI with config
		HideExport: false,
		HideTryIt:  false,
	}),
)

// Minimal setup (uses sensible defaults)
handler := specui.NewHandler(
	specui.WithSpecFile("openapi.yaml"),
	specui.WithSwaggerUI(config.Swagger{}),
)

// Register with any HTTP router
r.Get(handler.DocsPath(), handler.DocsFunc())   // Documentation UI
r.Get(handler.SpecPath(), handler.SpecFunc())   // OpenAPI spec file
```

## Configuration Options

The library uses functional options for flexible configuration:

### Core Options

```go
// Basic configuration
specui.WithTitle("My API")                     // Set documentation title
specui.WithDocsPath("/docs")                   // Set documentation URL path
specui.WithSpecPath("/docs/openapi.yaml")      // Set OpenAPI spec URL path
specui.WithSpecFile("openapi.yaml")            // Specify OpenAPI specification file path
```

### UI Selection with Configuration

#### Stoplight Elements Configuration
```go
specui.WithStoplightElements(config.Elements{
	HideExport:  false,                           // Hide the "Export" button
	HideSchemas: false,                           // Hide schemas in Table of Contents
	HideTryIt:   false,                           // Hide "Try it" interactive feature
	Layout:      "sidebar",                       // Layout: "sidebar" or "responsive"
	Logo:        "/assets/logo.png",              // URL to logo image
	Router:      "hash",                          // Router type: "hash", "memory"
})
```

#### Swagger UI Configuration
```go
specui.WithSwaggerUI(config.Swagger{
	ShowTopBar:         true,                     // Show navigation top bar
	HideCurl:           false,                    // Hide curl code snippets
	JsonEditor:         true,                     // Enable visual JSON editor (experimental)
	PreAuthorizeApiKey: map[string]string{        // Pre-authorize API keys
		"api_key": "your-api-key-here",
		"bearer":  "your-bearer-token",
	},
	SettingsUI: map[string]string{                // Advanced SwaggerUI configuration
		"deepLinking":            "true",
		"displayRequestDuration": "true",
		"filter":                 "true",
		"showExtensions":         "true",
	},
})
```

#### Redoc Configuration
```go
specui.WithRedoc(config.Redoc{
	HideDownload: false,                          // Hide download button for OpenAPI spec
})
```

## Configuration Examples

### Complete Example with All Options
```go
handler := specui.NewHandler(
	// Core configuration
	specui.WithTitle("Pet Store API"),
	specui.WithDocsPath("/documentation"),
	specui.WithSpecPath("/documentation/openapi.yaml"),
	specui.WithSpecFile("specs/petstore.yaml"),
	
	// Swagger UI with full configuration
	specui.WithSwaggerUI(config.Swagger{
		ShowTopBar:  true,
		JsonEditor:  true,
		PreAuthorizeApiKey: map[string]string{
			"api_key": "demo-key",
			"bearer":  "demo-token",
		},
		SettingsUI: map[string]string{
			"deepLinking":            "true",
			"displayRequestDuration": "true",
			"filter":                 "true",
		},
	}),
)
```

### Minimal Configuration
```go
// Uses sensible defaults
handler := specui.NewHandler(
	specui.WithSpecFile("openapi.yaml"),
	specui.WithSwaggerUI(config.Swagger{}),
)
// Default paths: /docs and /docs/openapi.yaml
```

The library automatically serves your OpenAPI specification file. Simply ensure your `openapi.yaml` (or `openapi.json`) file is in the correct location:

```go
// The file path is relative to your application's working directory
handler := specui.NewHandler(
	specui.WithSpecFile("openapi.yaml"),        // ./openapi.yaml
	specui.WithSpecFile("docs/openapi.yaml"),   // ./docs/openapi.yaml
	specui.WithSpecFile("/path/to/spec.json"),  // absolute path
	specui.WithStoplightElements(),
)
```

The handler automatically:
- Serves the documentation UI at the docs path
- Serves the OpenAPI specification at the spec path
- Handles both YAML and JSON format specifications

## Examples

Check out the [`examples`](/examples) directory for more usage patterns:

- Basic HTTP server setup
- Integration with popular Go frameworks
- Custom configuration options
- Dynamic OpenAPI spec generation

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

- 📖 Documentation: Check the examples and configuration options above
- 🐛 Issues: Report bugs or request features via GitHub Issues
- 💬 Discussions: Join the community discussions for questions and ideas

---

Made with ❤️ for the Go community
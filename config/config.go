package config

import (
	"embed"
	"io/fs"
)

type Provider uint8

const (
	ProviderSwaggerUI Provider = iota
	ProviderStoplightElements
	ProviderRedoc
)

// SpecGenerator is an interface for types that can generate OpenAPI specifications.
type SpecGenerator interface {
	MarshalYAML() ([]byte, error)
	MarshalJSON() ([]byte, error)
}

// SpecUI holds the configuration for the OpenAPI UI.
type SpecUI struct {
	Title         string        // Title of the OpenAPI UI
	DocsPath      string        // Path to the OpenAPI UI documentation
	SpecFile      string        // Path to the OpenAPI specification file
	SpecPath      string        // Path to the OpenAPI specification, default is openapi.yaml
	SpecIOFS      fs.FS         // Filesystem for the OpenAPI specification
	SpecEmbedFS   *embed.FS     // Embedded file system for the OpenAPI specification
	SpecGenerator SpecGenerator // OpenAPI specification generator

	Provider          Provider          // Provider type
	SwaggerUI         SwaggerUI         // Swagger UI configuration
	StoplightElements StoplightElements // Stoplight Elements configuration
	Redoc             Redoc             // Redoc configuration
}

// SwaggerUI holds the configuration for the Swagger UI.
type SwaggerUI struct {
	ShowTopBar         bool              // Show navigation top bar, hidden by default.
	HideCurl           bool              // Hide curl code snippet.
	JsonEditor         bool              // Enable visual json editor support (experimental, can fail with complex schemas).
	PreAuthorizeApiKey map[string]string // Map of security name to key value.

	// SettingsUI contains keys and plain javascript values of SwaggerUIBundle configuration.
	// Overrides default values.
	// See https://swagger.io/docs/open-source-tools/swagger-ui/usage/configuration/ for available options.
	SettingsUI map[string]string
}

// StoplightElements holds the configuration for the Stoplight Elements.
type StoplightElements struct {
	HideExport  bool   // Hide the "Export" button on overview section of the documentation.
	HideSchemas bool   // Hide the schemas in the Table of Contents, when using the sidebar layout.
	HideTryIt   bool   // Hide "Try it" feature.
	Layout      string // Layout type, e.g. "sidebar" or "responsive".
	Logo        string // Logo URL to an image that displays as a small square logo next to the title, above the table of contents.
	Router      string // Router type.
}

// Redoc holds the configuration for the Redoc.
type Redoc struct {
	DisableSearch    bool // Disable search functionality.
	HideDownload     bool // Hides the "Download" button for saving the API definition source file.
	HideSchemaTitles bool // Hides the schema titles in the documentation.
}

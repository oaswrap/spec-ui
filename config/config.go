package config

import (
	"embed"
	"io/fs"
)

type Provider uint8

const (
	ProviderSwaggerUI Provider = iota
	ProviderStoplightElements
	ProviderReDoc
	ProviderScalar
)

// SpecGenerator is an interface for types that can generate OpenAPI specifications.
type SpecGenerator interface {
	MarshalYAML() ([]byte, error)
	MarshalJSON() ([]byte, error)
}

// SpecUI holds the configuration for the OpenAPI UI.
type SpecUI struct {
	Title         string        // Title of the OpenAPI UI
	DocsPath      string        // Path to the OpenAPI UI documentation, defaults are "/docs"
	SpecFile      string        // Path to the OpenAPI specification file
	SpecPath      string        // Path to the OpenAPI specification, defaults are "/docs/openapi.json"
	SpecIOFS      fs.FS         // Filesystem for the OpenAPI specification
	SpecEmbedFS   *embed.FS     // Embedded file system for the OpenAPI specification
	SpecGenerator SpecGenerator // OpenAPI specification generator

	Provider          Provider           // Provider type
	SwaggerUI         *SwaggerUI         // Swagger UI configuration
	StoplightElements *StoplightElements // Stoplight Elements configuration
	ReDoc             *ReDoc             // ReDoc configuration
	Scalar            *Scalar            // Scalar configuration
}

// SwaggerUI holds the configuration for the Swagger UI.
type SwaggerUI struct {
	ShowTopBar         bool              // Show navigation top bar, hidden by default.
	HideCurl           bool              // Hide the curl code snippet.
	JsonEditor         bool              // Enable visual JSON editor support (experimental, can fail with complex schemas).
	PreAuthorizeApiKey map[string]string // Map of security name to key value.

	// SettingsUI contains keys and plain javascript values of SwaggerUIBundle configuration.
	// Overrides default values.
	// See https://swagger.io/docs/open-source-tools/swagger-ui/usage/configuration/ for available options.
	SettingsUI map[string]string
}

// StoplightElements holds the configuration for the Stoplight Elements.
type StoplightElements struct {
	HideExport  bool   // Hide the "Export" button on an overview section of the documentation.
	HideSchemas bool   // Hide the schemas in the Table of Contents when using the sidebar layout.
	HideTryIt   bool   // Hide the "Try it" feature.
	Layout      string // Layout type, e.g. "sidebar" or "responsive".
	Logo        string // Logo URL to an image that displays as a small square logo next to the title, above the table of contents.
	Router      string // Router type.
}

// ReDoc holds the configuration for the ReDoc.
type ReDoc struct {
	DisableSearch       bool // Disable search functionality.
	HideDownloadButtons bool // Hides the "Download" button for saving the API definition source file.
	HideSchemaTitles    bool // Hides the schema titles in the documentation.
}

// Scalar holds the configuration for the Scalar.
type Scalar struct {
	ProxyURL              string // Set Proxy URL to making API requests
	HideSidebar           bool   // Hide sidebar navigation
	HideModels            bool   // Hide models in the sidebar
	DocumentDownloadType  string // Document download type e.g. "json", "yaml", "both", or "none"
	HideTestRequestButton bool   // Hide the "Test Request" button
	HideSearch            bool   // Hide search bar
	DarkMode              bool   // Enable dark mode
	Layout                string // Layout type e.g. "modern" or "classic"
	Theme                 string // Theme name, see https://guides.scalar.com/scalar/scalar-api-references/themes for available themes
}

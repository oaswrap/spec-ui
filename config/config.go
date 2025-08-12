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
	ProviderRapiDoc
)

// SpecGenerator is an interface for types that can generate OpenAPI specifications.
type SpecGenerator interface {
	MarshalYAML() ([]byte, error)
	MarshalJSON() ([]byte, error)
}

// SpecUI holds the configuration for the OpenAPI UI.
type SpecUI struct {
	Title         string        // Title of the OpenAPI UI
	CacheAge      int           // Cache age for the OpenAPI specification, defaults is 1 hour
	DocsPath      string        // Path to the OpenAPI UI documentation, defaults are "/docs"
	SpecPath      string        // Path to the OpenAPI specification, defaults are "/docs/openapi.json"
	SpecFile      string        // Path to the OpenAPI specification file
	SpecIOFS      fs.FS         // Filesystem for the OpenAPI specification
	SpecEmbedFS   *embed.FS     // Embedded file system for the OpenAPI specification
	SpecGenerator SpecGenerator // OpenAPI specification generator

	Provider          Provider           // Provider type
	SwaggerUI         *SwaggerUI         // Swagger UI configuration
	StoplightElements *StoplightElements // Stoplight Elements configuration
	ReDoc             *ReDoc             // ReDoc configuration
	Scalar            *Scalar            // Scalar configuration
	RapiDoc           *RapiDoc           // RapiDoc configuration
}

type SwaggerLayout string

const (
	SwaggerLayoutStandalone SwaggerLayout = "StandaloneLayout"
	SwaggerLayoutBase       SwaggerLayout = "BaseLayout"
)

// SwaggerUI holds the configuration for the Swagger UI.
type SwaggerUI struct {
	HideCurl                 bool          // Hide the curl code snippet.
	JsonEditor               bool          // Enable visual JSON editor support (experimental, can fail with complex schemas).
	Layout                   SwaggerLayout // Layout type, e.g. "StandaloneLayout" or "BaseLayout".
	DefaultModelsExpandDepth int           // Default models expand depth, -1 means hide.

	// UIConfig specifies additional SwaggerUIBundle config object properties.
	// See https://swagger.io/docs/open-source-tools/swagger-ui/usage/configuration/ for available options.
	UIConfig map[string]string
}

type ElementLayout string

const (
	ElementLayoutSidebar    ElementLayout = "sidebar"
	ElementLayoutResponsive ElementLayout = "responsive"
	ElementLayoutStacked    ElementLayout = "stacked"
)

type ElementRouter string

const (
	ElementRouterHash    ElementRouter = "hash"
	ElementRouterHistory ElementRouter = "history"
	ElementRouterMemory  ElementRouter = "memory"
	ElementRouterStatic  ElementRouter = "static"
)

// StoplightElements holds the configuration for the Stoplight Elements.
type StoplightElements struct {
	HideExport     bool          // Hide the "Export" button on an overview section of the documentation.
	HideSchemas    bool          // Hide the schemas in the Table of Contents when using the sidebar layout.
	HideTryIt      bool          // Hide the "Try it" feature.
	HideTryItPanel bool          // Hide the "Try it" panel.
	Layout         ElementLayout // Layout type, e.g. "sidebar" or "responsive".
	Router         ElementRouter // Router type.
	Logo           string        // Logo URL to an image that displays as a small square logo next to the title, above the table of contents.
}

// ReDoc holds the configuration for the ReDoc.
type ReDoc struct {
	DisableSearch       bool // Disable search functionality.
	HideDownloadButtons bool // Hides the "Download" button for saving the API definition source file.
	HideSchemaTitles    bool // Hides the schema titles in the documentation.
}

type ScalarLayout string

const (
	ScalarLayoutModern  ScalarLayout = "modern"
	ScalarLayoutClassic ScalarLayout = "classic"
)

// Scalar holds the configuration for the Scalar.
type Scalar struct {
	ProxyURL              string       // Set Proxy URL to making API requests
	HideSidebar           bool         // Hide sidebar navigation
	HideModels            bool         // Hide models in the sidebar
	DocumentDownloadType  string       // Document download type e.g. "json", "yaml", "both", or "none"
	HideTestRequestButton bool         // Hide the "Test Request" button
	HideSearch            bool         // Hide search bar
	DarkMode              bool         // Enable dark mode
	Layout                ScalarLayout // Layout type e.g. "modern" or "classic"
	Theme                 string       // Theme name, see https://guides.scalar.com/scalar/scalar-api-references/themes for available themes
}

type RapiDocLayout string

const (
	RapiDocLayoutRow    RapiDocLayout = "row"
	RapiDocLayoutColumn RapiDocLayout = "column"
)

type RapiDocTheme string

const (
	RapiDocThemeLight RapiDocTheme = "light"
	RapiDocThemeDark  RapiDocTheme = "dark"
)

type RapiDocRenderStyle string

const (
	RapiDocRenderStyleRead    RapiDocRenderStyle = "read"
	RapiDocRenderStyleView    RapiDocRenderStyle = "view"
	RapiDocRenderStyleFocused RapiDocRenderStyle = "focused"
)

type RapiDocSchemaStyle string

const (
	RapiDocSchemaStyleTable RapiDocSchemaStyle = "table"
	RapiDocSchemaStyleTree  RapiDocSchemaStyle = "tree"
)

type RapiDoc struct {
	Theme              RapiDocTheme       // Theme style, "light" or "dark"
	Layout             RapiDocLayout      // Layout type, "row" or "column"
	RenderStyle        RapiDocRenderStyle // Render style, "read", "view", or "focused"
	SchemaStyle        RapiDocSchemaStyle // Schema style, "table" or "tree"
	BgColor            string             // Background color, e.g. "#fff"
	TextColor          string             // Text color, e.g. "#444"
	HeaderColor        string             // Header color, e.g. "#444444"
	PrimaryColor       string             // Primary color, e.g. "#FF791A"
	HideInfo           bool               // Hide the info section
	HideHeader         bool               // Hide the header section
	HideSearch         bool               // Hide the search bar
	HideAdvancedSearch bool               // Hide the advanced search bar
	HideTryIt          bool               // Hide the "Try" feature
	Logo               string             // Logo URL
}

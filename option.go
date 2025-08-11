package specui

import (
	"embed"
	"io/fs"

	"github.com/oaswrap/spec-ui/config"
)

func newConfig(opts ...Option) *config.SpecUI {
	cfg := &config.SpecUI{
		Title:    "OpenAPI Documentation",
		DocsPath: "/docs",
		SpecPath: "/docs/openapi.yaml",
		StoplightElements: config.StoplightElements{
			Router: "hash",
			Layout: "sidebar",
		},
		Provider: config.ProviderStoplightElements,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

// Option is a function that configures the OpenAPI UI.
type Option func(*config.SpecUI)

// WithTitle sets the title of the documentation.
func WithTitle(title string) Option {
	return func(c *config.SpecUI) {
		c.Title = title
	}
}

// WithDocsPath sets the path to the documentation.
func WithDocsPath(path string) Option {
	return func(c *config.SpecUI) {
		c.DocsPath = path
	}
}

// WithSpecPath sets the path to the specification.
func WithSpecPath(path string) Option {
	return func(c *config.SpecUI) {
		c.SpecPath = path
	}
}

// WithSpecFile sets the path to the specification file.
func WithSpecFile(filepath string) Option {
	return func(c *config.SpecUI) {
		c.SpecFile = filepath
	}
}

// WithSpecIOFS sets the generic I/O filesystem for the specification.
func WithSpecIOFS(filepath string, iofs fs.FS) Option {
	return func(c *config.SpecUI) {
		c.SpecFile = filepath
		c.SpecIOFS = iofs
	}
}

// WithSpecEmbedFS sets the embedded file system for the specification.
func WithSpecEmbedFS(filepath string, fs *embed.FS) Option {
	return func(c *config.SpecUI) {
		c.SpecFile = filepath
		c.SpecEmbedFS = fs
	}
}

// WithSpecGenerator sets up the OpenAPI specification generator.
func WithSpecGenerator(cfg config.SpecGenerator) Option {
	return func(c *config.SpecUI) {
		c.SpecGenerator = cfg
	}
}

// WithSwaggerUI sets up the Swagger UI configuration.
func WithSwaggerUI(cfg ...config.SwaggerUI) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderSwaggerUI
		if len(cfg) > 0 {
			c.SwaggerUI = cfg[0]
		}
	}
}

// WithStoplightElements sets up the Stoplight Elements configuration.
func WithStoplightElements(cfg ...config.StoplightElements) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderStoplightElements
		if len(cfg) > 0 {
			c.StoplightElements = cfg[0]
		}
		if c.StoplightElements.Router == "" {
			c.StoplightElements.Router = "hash"
		}
		if c.StoplightElements.Layout == "" {
			c.StoplightElements.Layout = "sidebar"
		}
	}
}

// WithReDoc sets up the ReDoc configuration.
func WithReDoc(cfg ...config.ReDoc) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderReDoc
		if len(cfg) > 0 {
			c.ReDoc = cfg[0]
		}
	}
}

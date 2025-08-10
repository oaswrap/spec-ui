package specui

import (
	"embed"

	"github.com/oaswrap/spec-ui/config"
)

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

// WithSpecFS sets the file system for the specification.
func WithSpecFS(fs *embed.FS) Option {
	return func(c *config.SpecUI) {
		c.SpecFS = fs
	}
}

// WithSpecGenerator sets up the OpenAPI specification generator.
func WithSpecGenerator(cfg config.SpecGenerator) Option {
	return func(c *config.SpecUI) {
		c.SpecGenerator = cfg
	}
}

// WithSwaggerUI sets up the Swagger UI configuration.
func WithSwaggerUI(cfg ...config.Swagger) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderSwaggerUI
		if len(cfg) > 0 {
			c.Swagger = cfg[0]
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

// WithRedoc sets up the Redoc configuration.
func WithRedoc(cfg ...config.Redoc) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderRedoc
		if len(cfg) > 0 {
			c.Redoc = cfg[0]
		}
	}
}

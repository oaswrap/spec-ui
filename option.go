package specui

import (
	"embed"
	"io/fs"

	"github.com/oaswrap/spec-ui/config"
)

func newConfig(opts ...Option) *config.SpecUI {
	cfg := &config.SpecUI{
		Title:             "OpenAPI Documentation",
		CacheAge:          3600, // Default cache age is 3600 seconds (1 hour)
		DocsPath:          "/docs",
		SpecPath:          "/docs/openapi.json",
		StoplightElements: &config.StoplightElements{},
		Provider:          config.ProviderStoplightElements,
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

// WithCacheAge sets the cache age for the documentation.
func WithCacheAge(age int) Option {
	return func(c *config.SpecUI) {
		if age >= 0 {
			c.CacheAge = age
		}
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

// WithSwaggerUI set ui documentation to use Swagger UI.
// It can be used to override the default Swagger UI configuration.
func WithSwaggerUI(cfg ...config.SwaggerUI) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderSwaggerUI
		if len(cfg) > 0 {
			c.SwaggerUI = &cfg[0]
		}
		if c.SwaggerUI == nil {
			c.SwaggerUI = &config.SwaggerUI{}
		}
		if c.SwaggerUI.Layout == "" {
			c.SwaggerUI.Layout = config.SwaggerLayoutStandalone
		}
		if c.SwaggerUI.DefaultModelsExpandDepth == 0 {
			c.SwaggerUI.DefaultModelsExpandDepth = 1
		}
	}
}

// WithStoplightElements set ui documentation to use Stoplight Elements.
// It can be used to override the default Stoplight Elements configuration.
// It sets the default router to "hash" and layout to "sidebar" if not specified.
func WithStoplightElements(cfg ...config.StoplightElements) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderStoplightElements
		if len(cfg) > 0 {
			c.StoplightElements = &cfg[0]
		}
		if c.StoplightElements == nil {
			c.StoplightElements = &config.StoplightElements{}
		}
	}
}

// WithReDoc set ui documentation to use ReDoc.
// It can be used to override the default ReDoc configuration.
func WithReDoc(cfg ...config.ReDoc) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderReDoc
		if len(cfg) > 0 {
			c.ReDoc = &cfg[0]
		}
		if c.ReDoc == nil {
			c.ReDoc = &config.ReDoc{}
		}
	}
}

// WithScalar set ui documentation to use Scalar.
// It can be used to override the default Scalar configuration.
func WithScalar(cfg ...config.Scalar) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderScalar
		if len(cfg) > 0 {
			c.Scalar = &cfg[0]
		}
		if c.Scalar == nil {
			c.Scalar = &config.Scalar{}
		}
	}
}

// WithRapiDoc set ui documentation to use RapiDoc.
// It can be used to override the default RapiDoc configuration.
func WithRapiDoc(cfg ...config.RapiDoc) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderRapiDoc
		if len(cfg) > 0 {
			c.RapiDoc = &cfg[0]
		}
		if c.RapiDoc == nil {
			c.RapiDoc = &config.RapiDoc{}
		}
	}
}

package specui

import (
	"embed"
	"io/fs"

	"github.com/oaswrap/spec-ui/config"
)

func newConfig(opts ...Option) *config.SpecUI {
	cfg := &config.SpecUI{
		Title:      "OpenAPI Documentation",
		CacheAge:   3600, // Default cache age is 3600 seconds (1 hour)
		DocsPath:   "/docs",
		SpecPath:   "/docs/openapi.json",
		AssetsPath: "/docs/_assets",
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

// WithAssetsPath overrides the URL prefix where embedded assets are served.
// It is meaningful only when embed mode is enabled.
func WithAssetsPath(path string) Option {
	return func(c *config.SpecUI) {
		c.AssetsPath = path
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

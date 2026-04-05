package specui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/rapidoc"
	"github.com/oaswrap/spec-ui/internal/rapidocemb"
	"github.com/oaswrap/spec-ui/internal/redoc"
	"github.com/oaswrap/spec-ui/internal/redocemb"
	"github.com/oaswrap/spec-ui/internal/scalar"
	"github.com/oaswrap/spec-ui/internal/scalaremb"
	"github.com/oaswrap/spec-ui/internal/stoplightelements"
	"github.com/oaswrap/spec-ui/internal/stoplightelementsemb"
	"github.com/oaswrap/spec-ui/internal/swaggerui"
	"github.com/oaswrap/spec-ui/internal/swaggeruiemb"
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

// WithEmbedAssets enables serving UI CSS/JS assets from the local binary.
// Without this option, assets are loaded from CDN.
func WithEmbedAssets() Option {
	return func(c *config.SpecUI) {
		c.EmbedAssets = true
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
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if c.EmbedAssets {
				return swaggeruiemb.NewHandler(c)
			}
			return swaggerui.NewHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if !c.EmbedAssets {
				return nil
			}
			return swaggeruiemb.NewAssetsHandler(c)
		}
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
func WithStoplightElements(cfg ...config.StoplightElements) Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderStoplightElements
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if c.EmbedAssets {
				return stoplightelementsemb.NewHandler(c)
			}
			return stoplightelements.NewHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if !c.EmbedAssets {
				return nil
			}
			return stoplightelementsemb.NewAssetsHandler(c)
		}
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
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if c.EmbedAssets {
				return redocemb.NewHandler(c)
			}
			return redoc.NewHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if !c.EmbedAssets {
				return nil
			}
			return redocemb.NewAssetsHandler(c)
		}
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
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if c.EmbedAssets {
				return scalaremb.NewHandler(c)
			}
			return scalar.NewHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if !c.EmbedAssets {
				return nil
			}
			return scalaremb.NewAssetsHandler(c)
		}
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
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if c.EmbedAssets {
				return rapidocemb.NewHandler(c)
			}
			return rapidoc.NewHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			if !c.EmbedAssets {
				return nil
			}
			return rapidocemb.NewAssetsHandler(c)
		}
		if len(cfg) > 0 {
			c.RapiDoc = &cfg[0]
		}
		if c.RapiDoc == nil {
			c.RapiDoc = &config.RapiDoc{}
		}
	}
}

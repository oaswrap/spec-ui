package scalar

import (
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

// WithUI configures the handler to use Scalar, loading assets from CDN.
// An optional config.Scalar value may be passed to customise the UI behaviour.
func WithUI(cfg ...config.Scalar) specui.Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderScalar
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return NewHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return nil
		}
		if len(cfg) > 0 {
			c.Scalar = &cfg[0]
		}
		if c.Scalar == nil {
			c.Scalar = &config.Scalar{}
		}
	}
}

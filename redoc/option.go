package redoc

import (
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

// WithUI configures the handler to use ReDoc, loading assets from CDN.
// An optional config.ReDoc value may be passed to customise the UI behaviour.
func WithUI(cfg ...config.ReDoc) specui.Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderReDoc
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return NewHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return nil
		}
		if len(cfg) > 0 {
			c.ReDoc = &cfg[0]
		}
		if c.ReDoc == nil {
			c.ReDoc = &config.ReDoc{}
		}
	}
}

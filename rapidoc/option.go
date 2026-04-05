package rapidoc

import (
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

// WithUI configures the handler to use RapiDoc, loading assets from CDN.
// An optional config.RapiDoc value may be passed to customise the UI behaviour.
func WithUI(cfg ...config.RapiDoc) specui.Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderRapiDoc
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return NewHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return nil
		}
		if len(cfg) > 0 {
			c.RapiDoc = &cfg[0]
		}
		if c.RapiDoc == nil {
			c.RapiDoc = &config.RapiDoc{}
		}
	}
}

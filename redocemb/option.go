package redocemb

import (
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

// WithUI configures the handler to use ReDoc, serving assets from embedded files.
// An optional config.ReDoc value may be passed to customise the UI behaviour.
func WithUI(cfg ...config.ReDoc) specui.Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderReDoc
		c.EmbedAssets = true
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return newHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return newAssetsHandler(c)
		}
		if len(cfg) > 0 {
			c.ReDoc = &cfg[0]
		}
		if c.ReDoc == nil {
			c.ReDoc = &config.ReDoc{}
		}
	}
}

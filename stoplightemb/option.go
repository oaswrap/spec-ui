package stoplightemb

import (
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

// WithUI configures the handler to use Stoplight Elements, serving assets from embedded files.
// An optional config.StoplightElements value may be passed to customise the UI behaviour.
func WithUI(cfg ...config.StoplightElements) specui.Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderStoplightElements
		c.EmbedAssets = true
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return newHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return newAssetsHandler(c)
		}
		if len(cfg) > 0 {
			c.StoplightElements = &cfg[0]
		}
		if c.StoplightElements == nil {
			c.StoplightElements = &config.StoplightElements{}
		}
	}
}

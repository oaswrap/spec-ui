package rapidocemb

import (
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

func WithUI(cfg ...config.RapiDoc) specui.Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderRapiDoc
		c.EmbedAssets = true
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return newHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return newAssetsHandler(c)
		}
		if len(cfg) > 0 {
			c.RapiDoc = &cfg[0]
		}
		if c.RapiDoc == nil {
			c.RapiDoc = &config.RapiDoc{}
		}
	}
}

package swaggeruiemb

import (
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

// WithUI configures the handler to use Swagger UI, serving assets from embedded files.
// An optional config.SwaggerUI value may be passed to customise the UI behaviour.
func WithUI(cfg ...config.SwaggerUI) specui.Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderSwaggerUI
		c.EmbedAssets = true
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return newHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return newAssetsHandler(c)
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

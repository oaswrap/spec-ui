package swaggerui

import (
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

func WithUI(cfg ...config.SwaggerUI) specui.Option {
	return func(c *config.SpecUI) {
		c.Provider = config.ProviderSwaggerUI
		c.DocsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return NewHandler(c)
		}
		c.AssetsHandlerFactory = func(c *config.SpecUI) http.Handler {
			return nil
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

package scalar

import (
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
)

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

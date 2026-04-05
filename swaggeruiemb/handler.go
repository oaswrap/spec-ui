package swaggeruiemb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	swaggerui "github.com/oaswrap/spec-ui/swaggerui"
)

func newHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return swaggerui.NewHandler(cfg)
}

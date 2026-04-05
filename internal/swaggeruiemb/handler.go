package swaggeruiemb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/swaggerui"
)

func NewHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return swaggerui.NewHandler(cfg)
}

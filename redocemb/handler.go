package redocemb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	redoc "github.com/oaswrap/spec-ui/redoc"
)

func newHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return redoc.NewHandler(cfg)
}

package redocemb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/redoc"
)

func NewHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return redoc.NewHandler(cfg)
}

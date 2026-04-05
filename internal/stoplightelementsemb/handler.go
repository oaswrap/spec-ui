package stoplightelementsemb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/stoplightelements"
)

func NewHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return stoplightelements.NewHandler(cfg)
}

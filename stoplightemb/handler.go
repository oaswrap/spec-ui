package stoplightemb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	stoplight "github.com/oaswrap/spec-ui/stoplight"
)

func newHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return stoplight.NewHandler(cfg)
}

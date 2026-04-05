package scalaremb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	scalar "github.com/oaswrap/spec-ui/scalar"
)

func newHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return scalar.NewHandler(cfg)
}

package scalaremb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/scalar"
)

func NewHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return scalar.NewHandler(cfg)
}

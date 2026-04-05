package rapidocemb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/rapidoc"
)

func NewHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return rapidoc.NewHandler(cfg)
}

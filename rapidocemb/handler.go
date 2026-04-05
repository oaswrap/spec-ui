package rapidocemb

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	rapidoc "github.com/oaswrap/spec-ui/rapidoc"
)

func newHandler(cfg *config.SpecUI) http.Handler {
	cfg.EmbedAssets = true
	return rapidoc.NewHandler(cfg)
}

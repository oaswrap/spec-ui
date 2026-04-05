package stoplightelementsemb

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/oaswrap/spec-ui/config"
)

//go:embed assets
var embeddedAssets embed.FS

func NewAssetsHandler(cfg *config.SpecUI) http.Handler {
	sub, err := fs.Sub(embeddedAssets, "assets")
	if err != nil {
		panic(err)
	}

	return http.StripPrefix(cfg.AssetsPath, http.FileServer(http.FS(sub)))
}

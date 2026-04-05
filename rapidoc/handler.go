package rapidoc

import (
	"html/template"
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/constant"
)

type Handler struct {
	Data

	tpl *template.Template
}

type Data struct {
	Title      string `json:"title"`
	OpenAPIURL string `json:"openapiURL"`
	Logo       string `json:"logo"`
}

// NewHandler returns a HTTP handler for swagger UI.
func NewHandler(cfg *config.SpecUI) *Handler {
	h := &Handler{
		Data: Data{
			Title:      cfg.Title,
			OpenAPIURL: cfg.SpecPath,
			Logo:       cfg.RapiDoc.Logo,
		},
	}
	var err error

	assetsBase := constant.RapiDocAssetBase
	faviconBase := constant.RapiDocFaviconBase
	if cfg.EmbedAssets {
		assetsBase = cfg.AssetsPath
		faviconBase = cfg.AssetsPath
	}

	h.tpl, err = template.New("index").Parse(IndexTpl(assetsBase, faviconBase, cfg.RapiDoc))
	if err != nil {
		panic(err)
	}

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := h.tpl.Execute(w, h); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

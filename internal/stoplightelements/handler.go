package stoplightelements

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/constant"
)

// Handler handles swagger UI request.
type Handler struct {
	Data

	ConfigJson template.JS

	tpl *template.Template
}

type Data struct {
	Title          string               `json:"title"`
	OpenAPIURL     string               `json:"openapiURL"`
	HideExport     bool                 `json:"hideExport"`
	HideSchemas    bool                 `json:"hideSchemas"`
	HideTryIt      bool                 `json:"hideTryIt"`
	HideTryItPanel bool                 `json:"hideTryItPanel"`
	Layout         config.ElementLayout `json:"layout"`
	Logo           string               `json:"logo"`
	Router         config.ElementRouter `json:"router"`
}

// NewHandler returns a HTTP handler for swagger UI.
func NewHandler(cfg *config.SpecUI) *Handler {
	h := &Handler{
		Data: Data{
			Title:          cfg.Title,
			OpenAPIURL:     cfg.SpecPath,
			HideExport:     cfg.StoplightElements.HideExport,
			HideSchemas:    cfg.StoplightElements.HideSchemas,
			HideTryIt:      cfg.StoplightElements.HideTryIt,
			HideTryItPanel: cfg.StoplightElements.HideTryItPanel,
			Layout:         cfg.StoplightElements.Layout,
			Logo:           cfg.StoplightElements.Logo,
			Router:         cfg.StoplightElements.Router,
		},
	}

	j, err := json.Marshal(h.Data)
	if err != nil {
		panic(err)
	}

	h.ConfigJson = template.JS(j) //nolint:gosec // Data is well formed.

	assetsBase := constant.StoplightElementsAssetsBase
	faviconBase := constant.StoplightElementFaviconBase
	if cfg.EmbedAssets {
		assetsBase = cfg.AssetsPath
		faviconBase = cfg.AssetsPath
	}

	h.tpl, err = template.New("index").Parse(IndexTpl(assetsBase, faviconBase, cfg))
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

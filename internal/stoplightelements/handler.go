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
func NewHandler(config *config.SpecUI) *Handler {
	h := &Handler{
		Data: Data{
			Title:          config.Title,
			OpenAPIURL:     config.SpecPath,
			HideExport:     config.StoplightElements.HideExport,
			HideSchemas:    config.StoplightElements.HideSchemas,
			HideTryIt:      config.StoplightElements.HideTryIt,
			HideTryItPanel: config.StoplightElements.HideTryItPanel,
			Layout:         config.StoplightElements.Layout,
			Logo:           config.StoplightElements.Logo,
			Router:         config.StoplightElements.Router,
		},
	}

	j, err := json.Marshal(h.Data)
	if err != nil {
		panic(err)
	}

	h.ConfigJson = template.JS(j) //nolint:gosec // Data is well formed.

	h.tpl, err = template.New("index").Parse(IndexTpl(constant.StoplightElementsAssetsBase, constant.StoplightElementFaviconBase, config))
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

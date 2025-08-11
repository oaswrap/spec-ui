package swaggerui

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
	Title              string            `json:"title"`
	OpenAPIURL         string            `json:"openapiURL"`
	ShowTopBar         bool              `json:"showTopBar"`
	HideCurl           bool              `json:"hideCurl"`
	JsonEditor         bool              `json:"jsonEditor"`
	PreAuthorizeApiKey map[string]string `json:"-"`
	SettingsUI         map[string]string `json:"-"`
}

// NewHandler returns a HTTP handler for swagger UI.
func NewHandler(config *config.SpecUI) *Handler {
	h := &Handler{
		Data: Data{
			Title:              config.Title,
			OpenAPIURL:         config.SpecPath,
			ShowTopBar:         config.SwaggerUI.ShowTopBar,
			HideCurl:           config.SwaggerUI.HideCurl,
			JsonEditor:         config.SwaggerUI.JsonEditor,
			PreAuthorizeApiKey: config.SwaggerUI.PreAuthorizeApiKey,
			SettingsUI:         config.SwaggerUI.SettingsUI,
		},
	}

	j, err := json.Marshal(h.Data)
	if err != nil {
		panic(err)
	}

	h.ConfigJson = template.JS(j) //nolint:gosec // Data is well formed.

	h.tpl, err = template.New("index").Parse(IndexTpl(constant.SwaggerUIAssetsBase, constant.SwaggerUIFaviconBase, config.SwaggerUI))
	if err != nil {
		panic(err)
	}

	return h
}

// ServeHTTP implements http.Handler interface to handle swagger UI request.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := h.tpl.Execute(w, h); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

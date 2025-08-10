package redoc

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/oaswrap/spec-ui/config"
)

// Handler handles swagger UI request.
type Handler struct {
	Data

	ConfigJson template.JS

	tpl *template.Template
}

type Data struct {
	Title          string `json:"title"`
	OpenAPIYAMLURL string `json:"openapiYamlUrl"`
	HideDownload   bool   `json:"hideDownload"`
}

// NewHandler returns a HTTP handler for swagger UI.
func NewHandler(config *config.SpecUI) *Handler {
	h := &Handler{
		Data: Data{
			Title:          config.Title,
			OpenAPIYAMLURL: config.SpecPath,
			HideDownload:   config.Redoc.HideDownload,
		},
	}

	j, err := json.Marshal(h.Data)
	if err != nil {
		panic(err)
	}

	h.ConfigJson = template.JS(j) //nolint:gosec // Data is well formed.

	h.tpl, err = template.New("index").Parse(IndexTpl(AssetsBase))
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

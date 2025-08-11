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
func NewHandler(config *config.SpecUI) *Handler {
	h := &Handler{
		Data: Data{
			Title:      config.Title,
			OpenAPIURL: config.SpecPath,
			Logo:       config.RapiDoc.Logo,
		},
	}
	var err error

	h.tpl, err = template.New("index").Parse(IndexTpl(constant.RapiDocAssetBase, config.RapiDoc))
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

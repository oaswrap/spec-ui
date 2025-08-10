package specui

import (
	"net/http"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/redoc"
	"github.com/oaswrap/spec-ui/internal/spec"
	"github.com/oaswrap/spec-ui/internal/stoplightelements"
	"github.com/oaswrap/spec-ui/internal/swaggerui"
)

// NewHandler creates a new HTTP handler for the OpenAPI UI.
//
// It applies the provided options to configure the OpenAPI UI.
func NewHandler(opts ...Option) *Handler {
	cfg := newConfig(opts...)

	return &Handler{cfg: cfg}
}

// Handler handles HTTP requests for the OpenAPI UI.
type Handler struct {
	cfg *config.SpecUI
}

// DocsPath returns the path to the API documentation.
func (h *Handler) DocsPath() string {
	return h.cfg.DocsPath
}

// SpecPath returns the path to the OpenAPI specification.
func (h *Handler) SpecPath() string {
	return h.cfg.SpecPath
}

// Docs returns the HTTP handler for the API documentation.
func (h *Handler) Docs() http.Handler {
	switch h.cfg.Provider {
	case config.ProviderSwaggerUI:
		return swaggerui.NewHandler(h.cfg)
	case config.ProviderStoplightElements:
		return stoplightelements.NewHandler(h.cfg)
	case config.ProviderRedoc:
		return redoc.NewHandler(h.cfg)
	default:
		return nil
	}
}

// DocsFunc returns the HTTP handler function for the API documentation.
func (h *Handler) DocsFunc() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Docs().ServeHTTP(w, r)
	})
}

// Spec returns the HTTP handler for the OpenAPI specification.
func (h *Handler) Spec() http.Handler {
	return spec.NewHandler(h.cfg)
}

// SpecFunc returns the HTTP handler function for the OpenAPI specification.
func (h *Handler) SpecFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Spec().ServeHTTP(w, r)
	}
}

package specui

import (
	"errors"
	"net/http"
	"sync"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/spec"
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
	cfg         *config.SpecUI
	docsOnce    sync.Once
	docsHandler http.Handler
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
// The handler is created once and cached for subsequent calls.
func (h *Handler) Docs() http.Handler {
	if h.cfg.DocsHandlerFactory == nil {
		panic(errors.New("no UI provider configured: use WithSwaggerUI, WithStoplightElements, WithReDoc, WithScalar, or WithRapiDoc"))
	}
	h.docsOnce.Do(func() {
		h.docsHandler = h.cfg.DocsHandlerFactory(h.cfg)
	})
	return h.docsHandler
}

// DocsFunc returns the HTTP handler function for the API documentation.
func (h *Handler) DocsFunc() http.HandlerFunc {
	return h.Docs().ServeHTTP
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

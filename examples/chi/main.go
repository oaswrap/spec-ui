package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/stoplight"
)

func main() {
	r := chi.NewRouter()

	// Stoplight Elements
	handler := specui.NewHandler(
		specui.WithTitle("Pet Store API"),
		specui.WithDocsPath("/docs"),
		specui.WithSpecPath("/docs/openapi.yaml"),
		specui.WithSpecFile("openapi.yaml"),
		stoplight.WithUI(),
	)

	r.Get(handler.DocsPath(), handler.DocsFunc())
	r.Get(handler.SpecPath(), handler.SpecFunc())
	if handler.AssetsEnabled() {
		r.Handle(handler.AssetsPath()+"/*", handler.Assets())
	}

	log.Printf("OpenAPI Documentation available at http://localhost:3000/docs")
	log.Printf("OpenAPI YAML available at http://localhost:3000/docs/openapi.yaml")

	http.ListenAndServe(":3000", r)
}

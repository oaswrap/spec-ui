package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	specui "github.com/oaswrap/spec-ui"
)

func main() {
	mux := mux.NewRouter()

	// Stoplight Elements
	handler := specui.NewHandler(
		specui.WithTitle("To-dos API"),
		specui.WithDocsPath("/docs"),
		specui.WithSpecPath("/docs/openapi.yaml"),
		specui.WithSpecFile("openapi.yaml"),
		specui.WithStoplightElements(),
	)

	mux.Handle(handler.DocsPath(), handler.Spec()).Methods("GET")
	mux.Handle(handler.SpecPath(), handler.Spec()).Methods("GET")

	log.Printf("OpenAPI Documentation available at http://localhost:3000/docs")
	log.Printf("OpenAPI YAML available at http://localhost:3000/docs/openapi.yaml")

	http.ListenAndServe(":3000", mux)
}

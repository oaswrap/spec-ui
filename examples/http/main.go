package main

import (
	"log"
	"net/http"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/rapidoc"
)

func main() {
	mux := http.NewServeMux()

	// RapiDoc
	handler := specui.NewHandler(
		specui.WithTitle("To-dos API"),
		specui.WithDocsPath("/docs"),
		specui.WithSpecPath("/docs/openapi.yaml"),
		specui.WithSpecFile("openapi.yaml"),
		rapidoc.WithUI(),
	)

	mux.Handle("GET "+handler.DocsPath(), handler.Docs())
	mux.Handle("GET "+handler.SpecPath(), handler.Spec())
	if handler.AssetsEnabled() {
		mux.Handle("GET "+handler.AssetsPath()+"/", handler.Assets())
	}

	log.Printf("OpenAPI Documentation available at http://localhost:3000/docs")
	log.Printf("OpenAPI YAML available at http://localhost:3000/docs/openapi.yaml")

	http.ListenAndServe(":3000", mux)
}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	specui "github.com/oaswrap/spec-ui"
)

func main() {
	r := gin.Default()

	// Stoplight Elements
	handler := specui.NewHandler(
		specui.WithTitle("To-dos API"),
		specui.WithDocsPath("/docs"),
		specui.WithSpecPath("/docs/openapi.yaml"),
		specui.WithSpecFile("openapi.yaml"),
		specui.WithStoplightElements(),
	)

	r.GET(handler.DocsPath(), gin.WrapH(handler.Docs()))
	r.GET(handler.SpecPath(), gin.WrapH(handler.Spec()))

	log.Printf("OpenAPI Documentation available at http://localhost:3000/docs")
	log.Printf("OpenAPI YAML available at http://localhost:3000/docs/openapi.yaml")

	http.ListenAndServe(":3000", r)
}

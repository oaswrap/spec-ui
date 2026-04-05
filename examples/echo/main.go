package main

import (
	"log"

	"github.com/labstack/echo/v4"
	specui "github.com/oaswrap/spec-ui"
)

func main() {
	e := echo.New()

	// Stoplight Elements
	handler := specui.NewHandler(
		specui.WithTitle("To-dos API"),
		specui.WithDocsPath("/docs"),
		specui.WithSpecPath("/docs/openapi.yaml"),
		specui.WithSpecFile("openapi.yaml"),
		specui.WithStoplightElements(),
	)

	e.GET(handler.DocsPath(), echo.WrapHandler(handler.Docs()))
	e.GET(handler.SpecPath(), echo.WrapHandler(handler.Spec()))

	log.Printf("OpenAPI Documentation available at http://localhost:3000/docs")
	log.Printf("OpenAPI YAML available at http://localhost:3000/docs/openapi.yaml")

	e.Logger.Fatal(e.Start(":3000"))
}

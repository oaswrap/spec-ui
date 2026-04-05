package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/scalar"
)

func main() {
	app := fiber.New()

	// Scalar
	handler := specui.NewHandler(
		specui.WithTitle("To-dos API"),
		specui.WithDocsPath("/docs"),
		specui.WithSpecPath("/docs/openapi.yaml"),
		specui.WithSpecFile("openapi.yaml"),
		scalar.WithUI(),
	)

	app.Get(handler.DocsPath(), adaptor.HTTPHandler(handler.Docs()))
	app.Get(handler.SpecPath(), adaptor.HTTPHandler(handler.Spec()))
	if handler.AssetsEnabled() {
		app.Get(handler.AssetsPath()+"/*", adaptor.HTTPHandler(handler.Assets()))
	}

	log.Printf("OpenAPI Documentation available at http://localhost:3000/docs")
	log.Printf("OpenAPI YAML available at http://localhost:3000/docs/openapi.yaml")

	log.Fatal(app.Listen(":3000"))
}

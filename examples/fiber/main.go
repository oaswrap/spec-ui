package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	specui "github.com/oaswrap/spec-ui"
)

func main() {
	app := fiber.New()

	// Stoplight Elements
	handler := specui.NewHandler(
		specui.WithTitle("To-dos API"),
		specui.WithDocsPath("/docs"),
		specui.WithSpecPath("/docs/openapi.yaml"),
		specui.WithSpecFile("openapi.yaml"),
		specui.WithStoplightElements(),
	)

	app.Get(handler.DocsPath(), adaptor.HTTPHandler(handler.DocsFunc()))
	app.Get(handler.SpecPath(), adaptor.HTTPHandler(handler.SpecFunc()))

	log.Printf("OpenAPI Documentation available at http://localhost:3000/docs")
	log.Printf("OpenAPI YAML available at http://localhost:3000/docs/openapi.yaml")

	log.Fatal(app.Listen(":3000"))
}

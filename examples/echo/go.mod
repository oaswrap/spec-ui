module github.com/oaswrap/spec-ui/examples/echo

go 1.26.0

require (
	github.com/labstack/echo/v5 v5.4.4
	github.com/oaswrap/spec-ui v0.0.0
)

require golang.org/x/net v0.60.0 // indirect

replace github.com/oaswrap/spec-ui => ../..

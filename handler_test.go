package specui_test

import (
	"testing"

	specui "github.com/oaswrap/spec-ui"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	t.Run("SwaggerUI", func(t *testing.T) {
		handler := specui.NewHandler(specui.WithSwaggerUI())
		assert.NotNil(t, handler)
	})
	t.Run("StoplightElements", func(t *testing.T) {
		handler := specui.NewHandler(specui.WithStoplightElements())
		assert.NotNil(t, handler)
	})
	t.Run("Redoc", func(t *testing.T) {
		handler := specui.NewHandler(specui.WithRedoc())
		assert.NotNil(t, handler)
	})
}

package specui_test

import (
	"testing"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/swaggerui"
	"github.com/oaswrap/spec-ui/swaggeruiemb"
	"github.com/stretchr/testify/assert"
)

func TestHandlerAssetsEnabled(t *testing.T) {
	t.Run("false for CDN provider", func(t *testing.T) {
		h := specui.NewHandler(swaggerui.WithUI())
		assert.False(t, h.AssetsEnabled())
	})
	t.Run("true for embed provider", func(t *testing.T) {
		h := specui.NewHandler(swaggeruiemb.WithUI())
		assert.True(t, h.AssetsEnabled())
	})
}

func TestHandlerAssetsPath(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		h := specui.NewHandler()
		assert.Equal(t, "/docs/_assets", h.AssetsPath())
	})
	t.Run("custom via WithAssetsPath", func(t *testing.T) {
		h := specui.NewHandler(specui.WithAssetsPath("/static/ui"))
		assert.Equal(t, "/static/ui", h.AssetsPath())
	})
}

func TestHandlerAssets(t *testing.T) {
	t.Run("nil for CDN provider", func(t *testing.T) {
		h := specui.NewHandler(swaggerui.WithUI())
		assert.Nil(t, h.Assets())
	})
	t.Run("non-nil for embed provider", func(t *testing.T) {
		h := specui.NewHandler(swaggeruiemb.WithUI())
		assert.NotNil(t, h.Assets())
	})
}

func TestHandlerDocsPanic(t *testing.T) {
	h := specui.NewHandler()
	assert.Panics(t, func() { h.Docs() })
}

package specui_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	specui "github.com/oaswrap/spec-ui"
	"github.com/stretchr/testify/assert"
)

func TestHandlerAssetsInCDNMode(t *testing.T) {
	h := specui.NewHandler(
		specui.WithSwaggerUI(),
		specui.WithSpecFile("testdata/petstore.yaml"),
	)

	assert.False(t, h.AssetsEnabled())
	assert.Equal(t, "/docs/_assets", h.AssetsPath())
	assert.Nil(t, h.Assets())
}

func TestHandlerAssetsInEmbedMode(t *testing.T) {
	h := specui.NewHandler(
		specui.WithEmbedAssets(),
		specui.WithSwaggerUI(),
		specui.WithSpecFile("testdata/petstore.yaml"),
	)

	assert.True(t, h.AssetsEnabled())
	assert.Equal(t, "/docs/_assets", h.AssetsPath())
	assert.NotNil(t, h.Assets())

	req := httptest.NewRequest(http.MethodGet, "/docs/_assets/swagger-ui.min.css", nil)
	rec := httptest.NewRecorder()
	h.Assets().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

func TestWithAssetsPathInCDNMode(t *testing.T) {
	h := specui.NewHandler(
		specui.WithSwaggerUI(),
		specui.WithAssetsPath("/custom/assets"),
		specui.WithSpecFile("testdata/petstore.yaml"),
	)

	assert.Equal(t, "/custom/assets", h.AssetsPath())
	assert.False(t, h.AssetsEnabled())
	assert.Nil(t, h.Assets())
}

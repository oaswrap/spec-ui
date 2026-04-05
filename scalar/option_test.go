package scalar

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/stretchr/testify/assert"
)

func TestWithUI(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI()(cfg)

	assert.Equal(t, config.ProviderScalar, cfg.Provider)
	assert.NotNil(t, cfg.DocsHandlerFactory)
	assert.NotNil(t, cfg.AssetsHandlerFactory)
	assert.NotNil(t, cfg.Scalar)
	assert.NotNil(t, cfg.DocsHandlerFactory(cfg))
	assert.Nil(t, cfg.AssetsHandlerFactory(cfg))
}

func TestWithUICustomConfig(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI(config.Scalar{DarkMode: true})(cfg)

	assert.True(t, cfg.Scalar.DarkMode)
}

func TestNewHandlerEmbedAssets(t *testing.T) {
	handler := NewHandler(&config.SpecUI{
		Title:       "My API",
		SpecPath:    "/openapi.json",
		AssetsPath:  "/docs/_assets",
		EmbedAssets: true,
		Scalar:      &config.Scalar{},
	})
	assert.NotNil(t, handler)

	req := httptest.NewRequest("GET", "/docs", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "/docs/_assets/style.min.css")
}

func TestIndexTplHideDeveloperTools(t *testing.T) {
	handler := NewHandler(&config.SpecUI{
		Title:    "My API",
		SpecPath: "/openapi.json",
		Scalar:   &config.Scalar{HideDeveloperTools: true},
	})
	assert.NotNil(t, handler)

	req := httptest.NewRequest("GET", "/docs", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), `showDeveloperTools: 'never'`)
}

package rapidoc

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/stretchr/testify/assert"
)

func TestWithUI(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI()(cfg)

	assert.Equal(t, config.ProviderRapiDoc, cfg.Provider)
	assert.NotNil(t, cfg.DocsHandlerFactory)
	assert.NotNil(t, cfg.AssetsHandlerFactory)
	assert.NotNil(t, cfg.RapiDoc)
	assert.NotNil(t, cfg.DocsHandlerFactory(cfg))
	assert.Nil(t, cfg.AssetsHandlerFactory(cfg))
}

func TestWithUICustomConfig(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI(config.RapiDoc{HideInfo: true})(cfg)

	assert.True(t, cfg.RapiDoc.HideInfo)
}

func TestNewHandlerEmbedAssets(t *testing.T) {
	handler := NewHandler(&config.SpecUI{
		Title:       "My API",
		SpecPath:    "/openapi.json",
		AssetsPath:  "/docs/_assets",
		EmbedAssets: true,
		RapiDoc:     &config.RapiDoc{},
	})
	assert.NotNil(t, handler)

	req := httptest.NewRequest("GET", "/docs", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "/docs/_assets/rapidoc-min.js")
}

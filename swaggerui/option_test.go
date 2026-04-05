package swaggerui

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/stretchr/testify/assert"
)

func TestWithUI(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI()(cfg)

	assert.Equal(t, config.ProviderSwaggerUI, cfg.Provider)
	assert.NotNil(t, cfg.DocsHandlerFactory)
	assert.NotNil(t, cfg.AssetsHandlerFactory)
	assert.Equal(t, config.SwaggerLayoutStandalone, cfg.SwaggerUI.Layout)
	assert.Equal(t, 1, cfg.SwaggerUI.DefaultModelsExpandDepth)
	assert.NotNil(t, cfg.DocsHandlerFactory(cfg))
	assert.Nil(t, cfg.AssetsHandlerFactory(cfg))
}

func TestWithUICustomConfig(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI(config.SwaggerUI{HideCurl: true, DefaultModelsExpandDepth: 3})(cfg)

	assert.True(t, cfg.SwaggerUI.HideCurl)
	assert.Equal(t, 3, cfg.SwaggerUI.DefaultModelsExpandDepth)
}

func TestNewHandlerEmbedAssets(t *testing.T) {
	handler := NewHandler(&config.SpecUI{
		Title:       "My API",
		SpecPath:    "/openapi.json",
		AssetsPath:  "/docs/_assets",
		EmbedAssets: true,
		SwaggerUI:   &config.SwaggerUI{Layout: config.SwaggerLayoutStandalone},
	})
	assert.NotNil(t, handler)

	req := httptest.NewRequest("GET", "/docs", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "/docs/_assets/swagger-ui.min.css")
}

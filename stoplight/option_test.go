package stoplight

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/stretchr/testify/assert"
)

func TestWithUI(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI()(cfg)

	assert.Equal(t, config.ProviderStoplightElements, cfg.Provider)
	assert.NotNil(t, cfg.DocsHandlerFactory)
	assert.NotNil(t, cfg.AssetsHandlerFactory)
	assert.NotNil(t, cfg.StoplightElements)
	assert.NotNil(t, cfg.DocsHandlerFactory(cfg))
	assert.Nil(t, cfg.AssetsHandlerFactory(cfg))
}

func TestWithUICustomConfig(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI(config.StoplightElements{HideExport: true})(cfg)

	assert.True(t, cfg.StoplightElements.HideExport)
}

func TestNewHandlerEmbedAssets(t *testing.T) {
	handler := NewHandler(&config.SpecUI{
		Title:             "My API",
		SpecPath:          "/openapi.json",
		AssetsPath:        "/docs/_assets",
		EmbedAssets:       true,
		StoplightElements: &config.StoplightElements{},
	})
	assert.NotNil(t, handler)

	req := httptest.NewRequest("GET", "/docs", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "/docs/_assets/styles.min.css")
}

func TestIndexTplHistoryRouter(t *testing.T) {
	handler := NewHandler(&config.SpecUI{
		Title:    "My API",
		DocsPath: "/docs",
		SpecPath: "/openapi.json",
		StoplightElements: &config.StoplightElements{
			Router: config.ElementRouterHistory,
		},
	})
	assert.NotNil(t, handler)

	req := httptest.NewRequest("GET", "/docs", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), `basePath="/docs"`)
}

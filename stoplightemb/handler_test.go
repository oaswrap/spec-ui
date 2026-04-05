package stoplightemb

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
	assert.True(t, cfg.EmbedAssets)
	assert.NotNil(t, cfg.DocsHandlerFactory)
	assert.NotNil(t, cfg.AssetsHandlerFactory)
	assert.NotNil(t, cfg.StoplightElements)
}

func TestWithUICustomConfig(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI(config.StoplightElements{HideTryIt: true})(cfg)

	assert.True(t, cfg.StoplightElements.HideTryIt)
}

func TestHandlerAndAssets(t *testing.T) {
	cfg := &config.SpecUI{
		Title:             "My API",
		DocsPath:          "/docs",
		AssetsPath:        "/docs/_assets",
		StoplightElements: &config.StoplightElements{},
	}

	handler := newHandler(cfg)
	assert.NotNil(t, handler)
	assert.True(t, cfg.EmbedAssets)

	docsReq := httptest.NewRequest("GET", "/docs", nil)
	docsRec := httptest.NewRecorder()
	handler.ServeHTTP(docsRec, docsReq)
	assert.Equal(t, 200, docsRec.Code)
	assert.Contains(t, docsRec.Body.String(), `/docs/_assets/favicons/favicon.ico`)

	assets := newAssetsHandler(cfg)
	assert.NotNil(t, assets)

	assetsReq := httptest.NewRequest("GET", "/docs/_assets/styles.min.css", nil)
	assetsRec := httptest.NewRecorder()
	assets.ServeHTTP(assetsRec, assetsReq)

	assert.Equal(t, 200, assetsRec.Code)
	assert.NotEmpty(t, assetsRec.Body.String())

	faviconReq := httptest.NewRequest("GET", "/docs/_assets/favicons/favicon.ico", nil)
	faviconRec := httptest.NewRecorder()
	assets.ServeHTTP(faviconRec, faviconReq)

	assert.Equal(t, 200, faviconRec.Code)
}

package scalaremb

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
	assert.True(t, cfg.EmbedAssets)
	assert.NotNil(t, cfg.DocsHandlerFactory)
	assert.NotNil(t, cfg.AssetsHandlerFactory)
	assert.NotNil(t, cfg.Scalar)
}

func TestWithUICustomConfig(t *testing.T) {
	cfg := &config.SpecUI{Title: "T", SpecPath: "/s", AssetsPath: "/a"}
	WithUI(config.Scalar{HideSearch: true})(cfg)

	assert.True(t, cfg.Scalar.HideSearch)
}

func TestHandlerAndAssets(t *testing.T) {
	cfg := &config.SpecUI{
		Title:      "My API",
		DocsPath:   "/docs",
		AssetsPath: "/docs/_assets",
		Scalar:     &config.Scalar{},
	}

	handler := newHandler(cfg)
	assert.NotNil(t, handler)
	assert.True(t, cfg.EmbedAssets)

	docsReq := httptest.NewRequest("GET", "/docs", nil)
	docsRec := httptest.NewRecorder()
	handler.ServeHTTP(docsRec, docsReq)
	assert.Equal(t, 200, docsRec.Code)
	assert.Contains(t, docsRec.Body.String(), `/docs/_assets/favicon.png`)

	assets := newAssetsHandler(cfg)
	assert.NotNil(t, assets)

	assetsReq := httptest.NewRequest("GET", "/docs/_assets/style.min.css", nil)
	assetsRec := httptest.NewRecorder()
	assets.ServeHTTP(assetsRec, assetsReq)

	assert.Equal(t, 200, assetsRec.Code)
	assert.NotEmpty(t, assetsRec.Body.String())

	faviconReq := httptest.NewRequest("GET", "/docs/_assets/favicon.png", nil)
	faviconRec := httptest.NewRecorder()
	assets.ServeHTTP(faviconRec, faviconReq)

	assert.Equal(t, 200, faviconRec.Code)
}

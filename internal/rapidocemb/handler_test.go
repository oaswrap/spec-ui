package rapidocemb_test

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/rapidocemb"
	"github.com/stretchr/testify/assert"
)

func TestHandlerAndAssets(t *testing.T) {
	cfg := &config.SpecUI{
		Title:      "My API",
		DocsPath:   "/docs",
		AssetsPath: "/docs/_assets",
		RapiDoc:    &config.RapiDoc{},
	}

	handler := rapidocemb.NewHandler(cfg)
	assert.NotNil(t, handler)
	assert.True(t, cfg.EmbedAssets)

	docsReq := httptest.NewRequest("GET", "/docs", nil)
	docsRec := httptest.NewRecorder()
	handler.ServeHTTP(docsRec, docsReq)
	assert.Equal(t, 200, docsRec.Code)
	assert.Contains(t, docsRec.Body.String(), `/docs/_assets/images/logo.png`)

	assets := rapidocemb.NewAssetsHandler(cfg)
	assert.NotNil(t, assets)

	assetsReq := httptest.NewRequest("GET", "/docs/_assets/rapidoc-min.js", nil)
	assetsRec := httptest.NewRecorder()
	assets.ServeHTTP(assetsRec, assetsReq)

	assert.Equal(t, 200, assetsRec.Code)
	assert.NotEmpty(t, assetsRec.Body.String())

	faviconReq := httptest.NewRequest("GET", "/docs/_assets/images/logo.png", nil)
	faviconRec := httptest.NewRecorder()
	assets.ServeHTTP(faviconRec, faviconReq)

	assert.Equal(t, 200, faviconRec.Code)
}

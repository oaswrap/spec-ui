package redocemb_test

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/redocemb"
	"github.com/stretchr/testify/assert"
)

func TestHandlerAndAssets(t *testing.T) {
	cfg := &config.SpecUI{
		Title:      "My API",
		DocsPath:   "/docs",
		AssetsPath: "/docs/_assets",
		ReDoc:      &config.ReDoc{},
	}

	handler := redocemb.NewHandler(cfg)
	assert.NotNil(t, handler)
	assert.True(t, cfg.EmbedAssets)

	docsReq := httptest.NewRequest("GET", "/docs", nil)
	docsRec := httptest.NewRecorder()
	handler.ServeHTTP(docsRec, docsReq)
	assert.Equal(t, 200, docsRec.Code)

	assets := redocemb.NewAssetsHandler(cfg)
	assert.NotNil(t, assets)

	assetsReq := httptest.NewRequest("GET", "/docs/_assets/redoc.standalone.js", nil)
	assetsRec := httptest.NewRecorder()
	assets.ServeHTTP(assetsRec, assetsReq)

	assert.Equal(t, 200, assetsRec.Code)
	assert.NotEmpty(t, assetsRec.Body.String())
}

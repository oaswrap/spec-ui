package stoplightelementsemb_test

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/stoplightelementsemb"
	"github.com/stretchr/testify/assert"
)

func TestHandlerAndAssets(t *testing.T) {
	cfg := &config.SpecUI{
		Title:             "My API",
		DocsPath:          "/docs",
		AssetsPath:        "/docs/_assets",
		StoplightElements: &config.StoplightElements{},
	}

	handler := stoplightelementsemb.NewHandler(cfg)
	assert.NotNil(t, handler)
	assert.True(t, cfg.EmbedAssets)

	docsReq := httptest.NewRequest("GET", "/docs", nil)
	docsRec := httptest.NewRecorder()
	handler.ServeHTTP(docsRec, docsReq)
	assert.Equal(t, 200, docsRec.Code)
	assert.Contains(t, docsRec.Body.String(), `/docs/_assets/favicons/favicon.ico`)

	assets := stoplightelementsemb.NewAssetsHandler(cfg)
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

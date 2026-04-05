package swaggeruiemb_test

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/swaggeruiemb"
	"github.com/stretchr/testify/assert"
)

func TestHandlerAndAssets(t *testing.T) {
	cfg := &config.SpecUI{
		Title:      "My API",
		DocsPath:   "/docs",
		AssetsPath: "/docs/_assets",
		SwaggerUI:  &config.SwaggerUI{},
	}

	handler := swaggeruiemb.NewHandler(cfg)
	assert.NotNil(t, handler)
	assert.True(t, cfg.EmbedAssets)

	docsReq := httptest.NewRequest("GET", "/docs", nil)
	docsRec := httptest.NewRecorder()
	handler.ServeHTTP(docsRec, docsReq)
	assert.Equal(t, 200, docsRec.Code)
	assert.Contains(t, docsRec.Body.String(), `/docs/_assets/favicon-16x16.png`)
	assert.Contains(t, docsRec.Body.String(), `/docs/_assets/favicon-32x32.png`)

	assets := swaggeruiemb.NewAssetsHandler(cfg)
	assert.NotNil(t, assets)

	assetsReq := httptest.NewRequest("GET", "/docs/_assets/swagger-ui.min.css", nil)
	assetsRec := httptest.NewRecorder()
	assets.ServeHTTP(assetsRec, assetsReq)

	assert.Equal(t, 200, assetsRec.Code)
	assert.NotEmpty(t, assetsRec.Body.String())
}

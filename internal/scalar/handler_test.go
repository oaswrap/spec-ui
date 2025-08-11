package scalar_test

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/scalar"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	handler := scalar.NewHandler(&config.SpecUI{
		Title:    "My API",
		DocsPath: "/docs",
		Scalar: &config.Scalar{
			ProxyURL:             "https://proxy.scalar.com",
			Layout:               "modern",
			DocumentDownloadType: "both",
			Theme:                "moon",
		},
	})
	assert.NotNil(t, handler)

	req := httptest.NewRequest("GET", "/docs", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "My API")
	assert.Contains(t, rec.Body.String(), "Scalar")
	assert.Contains(t, rec.Body.String(), "proxyUrl: 'https://proxy.scalar.com'")
	assert.Contains(t, rec.Body.String(), "layout: 'modern'")
	assert.Contains(t, rec.Body.String(), "documentDownloadType: 'both'")
	assert.Contains(t, rec.Body.String(), "theme: 'moon'")
}

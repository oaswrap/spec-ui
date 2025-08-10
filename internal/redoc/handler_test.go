package redoc_test

import (
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/redoc"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	handler := redoc.NewHandler(&config.SpecUI{
		Title:    "My API",
		DocsPath: "/docs",
	})
	assert.NotNil(t, handler)

	req := httptest.NewRequest("GET", "/docs", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "My API")
	assert.Contains(t, rec.Body.String(), "Redoc")
}

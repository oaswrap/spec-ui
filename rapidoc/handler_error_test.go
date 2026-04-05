package rapidoc

import (
	"errors"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type failFirstWriteResponseWriter struct {
	header http.Header
	code   int
	writes int
}

func newFailFirstWriteResponseWriter() *failFirstWriteResponseWriter {
	return &failFirstWriteResponseWriter{header: make(http.Header)}
}

func (w *failFirstWriteResponseWriter) Header() http.Header { return w.header }

func (w *failFirstWriteResponseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
}

func (w *failFirstWriteResponseWriter) Write(p []byte) (int, error) {
	w.writes++
	if w.writes == 1 {
		return 0, errors.New("forced write failure")
	}
	return len(p), nil
}

func TestServeHTTP_TemplateExecuteError(t *testing.T) {
	h := &Handler{tpl: template.Must(template.New("index").Parse("ok"))}
	w := newFailFirstWriteResponseWriter()
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)

	h.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.code)
	assert.GreaterOrEqual(t, w.writes, 2)
}

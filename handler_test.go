package specui_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/testdata"
	"github.com/stretchr/testify/assert"
)

type mockGenerator struct {
	shouldFail bool
}

func (m *mockGenerator) MarshalJSON() ([]byte, error) {
	if m.shouldFail {
		return nil, fmt.Errorf("failed to generate JSON")
	}
	return testdata.FS.ReadFile("petstore.json")
}

func (m *mockGenerator) MarshalYAML() ([]byte, error) {
	if m.shouldFail {
		return nil, fmt.Errorf("failed to generate YAML")
	}
	return testdata.FS.ReadFile("petstore.yaml")
}

func TestHandler(t *testing.T) {
	t.Run("DocsPath", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			handler := specui.NewHandler()
			assert.NotNil(t, handler)
			assert.Equal(t, "/docs", handler.DocsPath())
		})
		t.Run("custom", func(t *testing.T) {
			handler := specui.NewHandler(
				specui.WithDocsPath("/custom/docs"),
			)
			assert.NotNil(t, handler)
			assert.Equal(t, "/custom/docs", handler.DocsPath())
		})
	})
	t.Run("SpecPath", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			handler := specui.NewHandler()
			assert.NotNil(t, handler)
			assert.Equal(t, "/docs/openapi.yaml", handler.SpecPath())
		})
		t.Run("custom", func(t *testing.T) {
			handler := specui.NewHandler(
				specui.WithSpecPath("/custom/docs/openapi.yaml"),
			)
			assert.NotNil(t, handler)
			assert.Equal(t, "/custom/docs/openapi.yaml", handler.SpecPath())
		})
	})
	t.Run("Docs", func(t *testing.T) {
		t.Run("SwaggerUI", func(t *testing.T) {
			handler := specui.NewHandler(
				specui.WithTitle("Petstore API"),
				specui.WithSpecFile("testdata/petstore.yaml"),
				specui.WithSwaggerUI(config.SwaggerUI{}),
			)
			assert.NotNil(t, handler)

			req := httptest.NewRequest("GET", "/docs", nil)
			rec := httptest.NewRecorder()
			handler.DocsFunc()(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotNil(t, rec.Body)
			assert.Contains(t, rec.Body.String(), "Petstore API")
			assert.Contains(t, rec.Body.String(), "Swagger UI")
		})
		t.Run("StoplightElements", func(t *testing.T) {
			handler := specui.NewHandler(
				specui.WithTitle("Petstore API"),
				specui.WithSpecFile("testdata/petstore.yaml"),
				specui.WithStoplightElements(config.StoplightElements{}),
			)
			assert.NotNil(t, handler)

			req := httptest.NewRequest("GET", "/docs", nil)
			rec := httptest.NewRecorder()
			handler.DocsFunc()(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotNil(t, rec.Body)
			assert.Contains(t, rec.Body.String(), "Petstore API")
			assert.Contains(t, rec.Body.String(), "Stoplight Elements")
		})
		t.Run("Redoc", func(t *testing.T) {
			handler := specui.NewHandler(
				specui.WithTitle("Petstore API"),
				specui.WithSpecFile("testdata/petstore.yaml"),
				specui.WithReDoc(config.ReDoc{}),
			)
			assert.NotNil(t, handler)

			req := httptest.NewRequest("GET", "/docs", nil)
			rec := httptest.NewRecorder()
			handler.DocsFunc()(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotNil(t, rec.Body)
			assert.Contains(t, rec.Body.String(), "Petstore API")
			assert.Contains(t, rec.Body.String(), "Redoc")
		})
	})
	t.Run("Spec", func(t *testing.T) {
		t.Run("os file system", func(t *testing.T) {
			handler := specui.NewHandler(
				specui.WithTitle("Petstore API"),
				specui.WithSpecFile("testdata/petstore.yaml"),
			)
			assert.NotNil(t, handler)

			req := httptest.NewRequest("GET", "/docs/openapi.yaml", nil)
			rec := httptest.NewRecorder()
			handler.SpecFunc()(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotNil(t, rec.Body)
		})
		t.Run("io file system", func(t *testing.T) {
			handler := specui.NewHandler(
				specui.WithTitle("Petstore API"),
				specui.WithSpecIOFS("petstore.yaml", os.DirFS("testdata")),
			)
			assert.NotNil(t, handler)

			req := httptest.NewRequest("GET", "/docs/openapi.yaml", nil)
			rec := httptest.NewRecorder()
			handler.SpecFunc()(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotNil(t, rec.Body)
		})
		t.Run("embed file system", func(t *testing.T) {
			handler := specui.NewHandler(
				specui.WithTitle("Petstore API"),
				specui.WithSpecEmbedFS("petstore.yaml", &testdata.FS),
			)
			assert.NotNil(t, handler)

			req := httptest.NewRequest("GET", "/docs/openapi.yaml", nil)
			rec := httptest.NewRecorder()
			handler.SpecFunc()(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotNil(t, rec.Body)
		})
		t.Run("generator", func(t *testing.T) {
			handler := specui.NewHandler(
				specui.WithTitle("Petstore API"),
				specui.WithSpecGenerator(&mockGenerator{}),
			)
			assert.NotNil(t, handler)

			req := httptest.NewRequest("GET", "/docs/openapi.yaml", nil)
			rec := httptest.NewRecorder()
			handler.SpecFunc()(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotNil(t, rec.Body)
		})
	})
}

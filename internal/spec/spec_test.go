package spec_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/spec"
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
	t.Run("when use os filesystem", func(t *testing.T) {
		t.Run("yaml", func(t *testing.T) {
			cfg := &config.SpecUI{
				SpecPath: "/docs/openapi.yaml",
				SpecFile: "../../testdata/petstore.yaml",
			}
			handler := spec.NewHandler(cfg)

			req := httptest.NewRequest("GET", "/docs/openapi.yaml", nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			assert.Equal(t, 200, rec.Code)
			assert.Equal(t, "application/x-yaml", rec.Header().Get("Content-Type"))
			assert.Contains(t, rec.Body.String(), "Swagger Petstore")
		})
		t.Run("json", func(t *testing.T) {
			cfg := &config.SpecUI{
				SpecPath: "/docs/openapi.json",
				SpecFile: "../../testdata/petstore.json",
			}
			handler := spec.NewHandler(cfg)

			req := httptest.NewRequest("GET", "/docs/openapi.json", nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			assert.Equal(t, 200, rec.Code)
			assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
			assert.Contains(t, rec.Body.String(), "Swagger Petstore")
		})
		t.Run("not found", func(t *testing.T) {
			cfg := &config.SpecUI{
				SpecPath: "/docs/openapi.json",
				SpecFile: "../../testdata/nonexistent.json",
			}
			handler := spec.NewHandler(cfg)

			req := httptest.NewRequest("GET", "/docs/openapi.json", nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			assert.Equal(t, 404, rec.Code)
			assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
			assert.Contains(t, rec.Body.String(), "OpenAPI specification file is not found")
		})
	})
	t.Run("when use embed filesystem", func(t *testing.T) {
		t.Run("yaml", func(t *testing.T) {
			cfg := &config.SpecUI{
				SpecPath: "/docs/openapi.yaml",
				SpecFile: "petstore.yaml",
				SpecFS:   &testdata.FS,
			}
			handler := spec.NewHandler(cfg)

			req := httptest.NewRequest("GET", "/docs/openapi.yaml", nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			assert.Equal(t, 200, rec.Code)
			assert.Equal(t, "application/x-yaml", rec.Header().Get("Content-Type"))
			assert.Contains(t, rec.Body.String(), "Swagger Petstore")
		})
		t.Run("json", func(t *testing.T) {
			cfg := &config.SpecUI{
				SpecPath: "/docs/openapi.json",
				SpecFile: "petstore.json",
				SpecFS:   &testdata.FS,
			}
			handler := spec.NewHandler(cfg)

			req := httptest.NewRequest("GET", "/docs/openapi.json", nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			assert.Equal(t, 200, rec.Code)
			assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
			assert.Contains(t, rec.Body.String(), "Swagger Petstore")
		})
		t.Run("not found", func(t *testing.T) {
			cfg := &config.SpecUI{
				SpecPath: "/docs/openapi.json",
				SpecFile: "nonexistent.json",
				SpecFS:   &testdata.FS,
			}
			handler := spec.NewHandler(cfg)

			req := httptest.NewRequest("GET", "/docs/openapi.json", nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			assert.Equal(t, 404, rec.Code)
			assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
			assert.Contains(t, rec.Body.String(), "OpenAPI specification file is not found")
		})
	})
	t.Run("when use spec generator", func(t *testing.T) {
		t.Run("yaml", func(t *testing.T) {
			cfg := &config.SpecUI{
				SpecPath:      "/docs/openapi.yaml",
				SpecGenerator: &mockGenerator{},
			}
			handler := spec.NewHandler(cfg)

			req := httptest.NewRequest("GET", "/docs/openapi.yaml", nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			assert.Equal(t, 200, rec.Code)
			assert.Equal(t, "application/x-yaml", rec.Header().Get("Content-Type"))
			assert.Contains(t, rec.Body.String(), "Swagger Petstore")
		})
		t.Run("json", func(t *testing.T) {
			cfg := &config.SpecUI{
				SpecPath:      "/docs/openapi.json",
				SpecGenerator: &mockGenerator{},
			}
			handler := spec.NewHandler(cfg)

			req := httptest.NewRequest("GET", "/docs/openapi.json", nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			assert.Equal(t, 200, rec.Code)
			assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
			assert.Contains(t, rec.Body.String(), "Swagger Petstore")
		})
		t.Run("failed", func(t *testing.T) {
			cfg := &config.SpecUI{
				SpecPath:      "/docs/openapi.json",
				SpecGenerator: &mockGenerator{shouldFail: true},
			}
			handler := spec.NewHandler(cfg)

			req := httptest.NewRequest("GET", "/docs/openapi.json", nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			assert.Equal(t, 500, rec.Code)
			assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
			assert.Contains(t, rec.Body.String(), "failed to generate OpenAPI schema")
		})
	})
	t.Run("when spec file not set", func(t *testing.T) {
		cfg := &config.SpecUI{
			SpecPath: "/docs/openapi.json",
		}
		handler := spec.NewHandler(cfg)

		req := httptest.NewRequest("GET", "/docs/openapi.json", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		assert.Contains(t, rec.Body.String(), "OpenAPI specification file is not set")
	})
}

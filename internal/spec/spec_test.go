package spec_test

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec-ui/internal/spec"
	"github.com/oaswrap/spec-ui/testdata"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name        string
		config      *config.SpecUI
		contentType string
		shouldError bool
		errorStatus int
	}{
		{
			name: "when serving OpenAPI YAML from embed FS",
			config: &config.SpecUI{
				SpecPath:    "/docs/openapi.yaml",
				SpecFile:    "petstore.yaml",
				SpecEmbedFS: &testdata.FS,
			},
			contentType: "application/x-yaml",
		},
		{
			name: "when serving OpenAPI JSON from embed FS",
			config: &config.SpecUI{
				SpecPath:    "/docs/openapi.json",
				SpecFile:    "petstore.json",
				SpecEmbedFS: &testdata.FS,
			},
			contentType: "application/json",
		},
		{
			name: "when serving OpenAPI YAML from embed FS and not found",
			config: &config.SpecUI{
				SpecPath:    "/docs/openapi.yaml",
				SpecFile:    "notexists.yaml",
				SpecEmbedFS: &testdata.FS,
			},
			shouldError: true,
			errorStatus: 404,
		},
		{
			name: "when serving OpenAPI YAML from OS FS",
			config: &config.SpecUI{
				SpecPath: "/docs/openapi.yaml",
				SpecFile: "../../testdata/petstore.yaml",
			},
			contentType: "application/x-yaml",
		},
		{
			name: "when serving OpenAPI JSON from OS FS",
			config: &config.SpecUI{
				SpecPath: "/docs/openapi.json",
				SpecFile: "../../testdata/petstore.json",
			},
			contentType: "application/json",
		},
		{
			name: "when serving OpenAPI YAML from OS FS and not found",
			config: &config.SpecUI{
				SpecPath: "/docs/openapi.yaml",
				SpecFile: "../../testdata/notexists.yaml",
			},
			shouldError: true,
			errorStatus: 404,
		},
		{
			name: "when serving OpenAPI YAML from IOFS",
			config: &config.SpecUI{
				SpecPath: "/docs/openapi.yaml",
				SpecFile: "petstore.yaml",
				SpecIOFS: os.DirFS("../../testdata"),
			},
			contentType: "application/x-yaml",
		},
		{
			name: "when serving OpenAPI JSON from IOFS",
			config: &config.SpecUI{
				SpecPath: "/docs/openapi.json",
				SpecFile: "petstore.json",
				SpecIOFS: os.DirFS("../../testdata"),
			},
			contentType: "application/json",
		},
		{
			name: "when serving OpenAPI JSON from IOFS and not found",
			config: &config.SpecUI{
				SpecPath: "/docs/openapi.json",
				SpecFile: "notexists.json",
				SpecIOFS: os.DirFS("../../testdata"),
			},
			shouldError: true,
			errorStatus: 404,
		},
		{
			name: "when serving OpenAPI YAML from SpecGenerator",
			config: &config.SpecUI{
				SpecPath:      "/docs/openapi.yaml",
				SpecFile:      "petstore.yaml",
				SpecGenerator: &mockGenerator{},
			},
			contentType: "application/x-yaml",
		},
		{
			name: "when serving OpenAPI JSON from SpecGenerator",
			config: &config.SpecUI{
				SpecPath:      "/docs/openapi.json",
				SpecFile:      "petstore.json",
				SpecGenerator: &mockGenerator{},
			},
			contentType: "application/json",
		},
		{
			name: "when serving OpenAPI YAML from SpecGenerator and failure",
			config: &config.SpecUI{
				SpecPath:      "/docs/openapi.yaml",
				SpecFile:      "petstore.yaml",
				SpecGenerator: &mockGenerator{shouldFail: true},
			},
			shouldError: true,
			errorStatus: 500,
		},
		{
			name: "when config not set",
			config: &config.SpecUI{
				SpecPath: "/docs/openapi.yaml",
			},
			shouldError: true,
			errorStatus: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := spec.NewHandler(tt.config)

			req := httptest.NewRequest("GET", tt.config.SpecPath, nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			if tt.shouldError {
				assert.Equal(t, tt.errorStatus, rec.Code)
			} else {
				assert.Equal(t, 200, rec.Code)
				assert.Equal(t, tt.contentType, rec.Header().Get("Content-Type"))
				assert.Contains(t, rec.Body.String(), "Swagger Petstore")
			}
		})
	}
}

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

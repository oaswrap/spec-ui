package specui_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	specui "github.com/oaswrap/spec-ui"
	"github.com/oaswrap/spec-ui/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmbeddedProviders(t *testing.T) {
	tests := []struct {
		name          string
		opts          []specui.Option
		docsContains  []string
		assetRequests []string
	}{
		{
			name: "SwaggerUI",
			opts: []specui.Option{
				specui.WithEmbedAssets(),
				specui.WithSwaggerUI(),
				specui.WithSpecFile("testdata/petstore.yaml"),
			},
			docsContains: []string{
				"/docs/_assets/swagger-ui.min.css",
				"/docs/_assets/favicon-16x16.png",
				"/docs/_assets/favicon-32x32.png",
			},
			assetRequests: []string{
				"/docs/_assets/swagger-ui.min.css",
			},
		},
		{
			name: "StoplightElements",
			opts: []specui.Option{
				specui.WithEmbedAssets(),
				specui.WithStoplightElements(),
				specui.WithSpecFile("testdata/petstore.yaml"),
			},
			docsContains: []string{
				"/docs/_assets/styles.min.css",
				"/docs/_assets/web-components.min.js",
				"/docs/_assets/favicons/favicon.ico",
			},
			assetRequests: []string{
				"/docs/_assets/styles.min.css",
				"/docs/_assets/web-components.min.js",
			},
		},
		{
			name: "ReDoc",
			opts: []specui.Option{
				specui.WithEmbedAssets(),
				specui.WithReDoc(),
				specui.WithSpecFile("testdata/petstore.yaml"),
			},
			docsContains: []string{
				"/docs/_assets/redoc.standalone.js",
			},
			assetRequests: []string{
				"/docs/_assets/redoc.standalone.js",
			},
		},
		{
			name: "Scalar",
			opts: []specui.Option{
				specui.WithEmbedAssets(),
				specui.WithScalar(),
				specui.WithSpecFile("testdata/petstore.yaml"),
			},
			docsContains: []string{
				"/docs/_assets/style.min.css",
				"/docs/_assets/browser/standalone.min.js",
				"/docs/_assets/favicon.png",
			},
			assetRequests: []string{
				"/docs/_assets/style.min.css",
				"/docs/_assets/browser/standalone.min.js",
			},
		},
		{
			name: "RapiDoc",
			opts: []specui.Option{
				specui.WithEmbedAssets(),
				specui.WithRapiDoc(),
				specui.WithSpecFile("testdata/petstore.yaml"),
			},
			docsContains: []string{
				"/docs/_assets/rapidoc-min.js",
				"/docs/_assets/images/logo.png",
			},
			assetRequests: []string{
				"/docs/_assets/rapidoc-min.js",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := specui.NewHandler(tt.opts...)
			require.True(t, h.AssetsEnabled())
			require.NotNil(t, h.Assets())

			req := httptest.NewRequest(http.MethodGet, "/docs", nil)
			rec := httptest.NewRecorder()
			h.DocsFunc()(rec, req)

			require.Equal(t, http.StatusOK, rec.Code)
			for _, expected := range tt.docsContains {
				assert.Contains(t, rec.Body.String(), expected)
			}

			for _, assetPath := range tt.assetRequests {
				assetReq := httptest.NewRequest(http.MethodGet, assetPath, nil)
				assetRec := httptest.NewRecorder()
				h.Assets().ServeHTTP(assetRec, assetReq)

				require.Equal(t, http.StatusOK, assetRec.Code)
				require.NotEmpty(t, assetRec.Body.String())
			}
		})
	}
}

func TestDocsPanicsWithoutProvider(t *testing.T) {
	h := specui.NewHandler(specui.WithSpecFile("testdata/petstore.yaml"))

	require.Panics(t, func() {
		h.Docs()
	})
}

func TestProviderAssetsInCDNMode(t *testing.T) {
	tests := []struct {
		name string
		opts []specui.Option
	}{
		{
			name: "SwaggerUI",
			opts: []specui.Option{specui.WithSwaggerUI(), specui.WithSpecFile("testdata/petstore.yaml")},
		},
		{
			name: "StoplightElements",
			opts: []specui.Option{specui.WithStoplightElements(), specui.WithSpecFile("testdata/petstore.yaml")},
		},
		{
			name: "ReDoc",
			opts: []specui.Option{specui.WithReDoc(), specui.WithSpecFile("testdata/petstore.yaml")},
		},
		{
			name: "Scalar",
			opts: []specui.Option{specui.WithScalar(), specui.WithSpecFile("testdata/petstore.yaml")},
		},
		{
			name: "RapiDoc",
			opts: []specui.Option{specui.WithRapiDoc(), specui.WithSpecFile("testdata/petstore.yaml")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := specui.NewHandler(tt.opts...)
			require.False(t, h.AssetsEnabled())
			assert.Nil(t, h.Assets())
		})
	}
}

func TestWithEmbedAssetsAndAssetsPath(t *testing.T) {
	h := specui.NewHandler(
		specui.WithEmbedAssets(),
		specui.WithAssetsPath("/custom/assets"),
		specui.WithSwaggerUI(config.SwaggerUI{}),
		specui.WithSpecFile("testdata/petstore.yaml"),
	)

	require.True(t, h.AssetsEnabled())
	assert.Equal(t, "/custom/assets", h.AssetsPath())
	require.NotNil(t, h.Assets())

	req := httptest.NewRequest(http.MethodGet, "/custom/assets/swagger-ui.min.css", nil)
	rec := httptest.NewRecorder()
	h.Assets().ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	require.NotEmpty(t, rec.Body.String())
}

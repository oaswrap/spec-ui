package spec

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/oaswrap/spec-ui/config"
)

type Handler struct {
	cfg      *config.SpecUI
	fileType string
	once     sync.Once
	schema   []byte
	err      error
}

func NewHandler(cfg *config.SpecUI) *Handler {
	fileType := "yaml"
	if strings.HasSuffix(cfg.SpecPath, ".json") {
		fileType = "json"
	}
	return &Handler{cfg: cfg, fileType: fileType}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.cfg.SpecGenerator != nil {
		h.once.Do(func() {
			if h.fileType == "json" {
				h.schema, h.err = h.cfg.SpecGenerator.MarshalJSON()
				return
			}
			h.schema, h.err = h.cfg.SpecGenerator.MarshalYAML()
		})

		if h.err != nil {
			h.renderError(w, 500, errors.New("failed to generate OpenAPI schema"))
			return
		}
	} else if h.cfg.SpecFS != nil && h.cfg.SpecFile != "" {
		h.once.Do(func() {
			h.schema, h.err = h.cfg.SpecFS.ReadFile(h.cfg.SpecFile)
		})

		if h.err != nil {
			h.renderError(w, 404, errors.New("OpenAPI specification file is not found"))
			return
		}
	} else if h.cfg.SpecFile != "" {
		h.once.Do(func() {
			h.schema, h.err = os.ReadFile(h.cfg.SpecFile)
		})

		if h.err != nil {
			h.renderError(w, 404, errors.New("OpenAPI specification file is not found"))
			return
		}
	} else {
		h.renderError(w, 500, errors.New("OpenAPI specification file is not set"))
		return
	}

	if h.fileType == "json" {
		w.Header().Set("Content-Type", "application/json")
	} else {
		w.Header().Set("Content-Type", "application/x-yaml")
	}
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(h.schema)
	if err != nil {
		log.Printf("failed to write OpenAPI schema: %v", err)
		return
	}
}

func (h *Handler) renderError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status":  status,
		"message": err.Error(),
	})
}

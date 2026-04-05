# Variables
PKG := ./...
COVERAGE_FILE := coverage.out
SWAGGERUI_VER := 5.32.1
STOPLIGHT_VER := 9.0.16
REDOC_VER := 2.5.2
SCALAR_VER := 1.51.0
RAPIDOC_VER := 9.3.8
CDN := https://cdn.jsdelivr.net/npm

# Default target
.PHONY: all
all: test

# Run tests with gotestsum
.PHONY: test
test:
	@echo "Running tests with gotestsum..."
	@gotestsum --format standard-quiet -- $(PKG)

# Run tests and update golden files
.PHONY: test-update
test-update:
	@echo "Running tests with gotestsum and updating golden files..."
	@gotestsum --format standard-quiet -- -update $(PKG)

# Run tests with coverage and generate report
.PHONY: testcov
testcov:
	@echo "Running tests with gotestsum and coverage..."
	@gotestsum --format standard-quiet -- -coverprofile=$(COVERAGE_FILE) $(PKG)
	@go tool cover -func=$(COVERAGE_FILE)
	@echo "Open HTML coverage report: make testcov-html"

# Open HTML coverage report
.PHONY: testcov-html
testcov-html:
	@go tool cover -html=$(COVERAGE_FILE)

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting..."
	@go fmt $(PKG)

# Tidy go.mod and go.sum
.PHONY: tidy
tidy:
	@echo "Tidying modules..."
	@go mod tidy

# Lint code (requires golangci-lint)
.PHONY: lint
lint:
	@echo "Linting..."
	@golangci-lint run

# Clean up generated files
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f $(COVERAGE_FILE)

# Update dependencies
.PHONY: update
update:
	@echo "Updating dependencies..."
	@go get -u ./...

.PHONY: download-assets
download-assets:
	@mkdir -p internal/swaggeruiemb/assets
	@curl -fsSL $(CDN)/swagger-ui@$(SWAGGERUI_VER)/dist/swagger-ui.min.css -o internal/swaggeruiemb/assets/swagger-ui.min.css
	@curl -fsSL $(CDN)/swagger-ui@$(SWAGGERUI_VER)/dist/swagger-ui-bundle.js -o internal/swaggeruiemb/assets/swagger-ui-bundle.js
	@curl -fsSL $(CDN)/swagger-ui@$(SWAGGERUI_VER)/dist/swagger-ui-standalone-preset.js -o internal/swaggeruiemb/assets/swagger-ui-standalone-preset.js
	@curl -fsSL https://petstore.swagger.io/favicon-16x16.png -o internal/swaggeruiemb/assets/favicon-16x16.png
	@curl -fsSL https://petstore.swagger.io/favicon-32x32.png -o internal/swaggeruiemb/assets/favicon-32x32.png
	@mkdir -p internal/stoplightelementsemb/assets
	@curl -fsSL $(CDN)/@stoplight/elements@$(STOPLIGHT_VER)/styles.min.css -o internal/stoplightelementsemb/assets/styles.min.css
	@curl -fsSL $(CDN)/@stoplight/elements@$(STOPLIGHT_VER)/web-components.min.js -o internal/stoplightelementsemb/assets/web-components.min.js
	@mkdir -p internal/stoplightelementsemb/assets/favicons
	@curl -fsSL https://docs.stoplight.io/favicons/favicon.ico -o internal/stoplightelementsemb/assets/favicons/favicon.ico
	@mkdir -p internal/redocemb/assets
	@curl -fsSL $(CDN)/redoc@$(REDOC_VER)/bundles/redoc.standalone.js -o internal/redocemb/assets/redoc.standalone.js
	@mkdir -p internal/scalaremb/assets/browser
	@curl -fsSL $(CDN)/@scalar/api-reference@$(SCALAR_VER)/dist/style.min.css -o internal/scalaremb/assets/style.min.css
	@curl -fsSL $(CDN)/@scalar/api-reference@$(SCALAR_VER)/dist/browser/standalone.min.js -o internal/scalaremb/assets/browser/standalone.min.js
	@curl -fsSL https://scalar.com/favicon.png -o internal/scalaremb/assets/favicon.png
	@mkdir -p internal/rapidocemb/assets
	@curl -fsSL $(CDN)/rapidoc@$(RAPIDOC_VER)/dist/rapidoc-min.js -o internal/rapidocemb/assets/rapidoc-min.js
	@mkdir -p internal/rapidocemb/assets/images
	@curl -fsSL https://rapidocweb.com/images/logo.png -o internal/rapidocemb/assets/images/logo.png
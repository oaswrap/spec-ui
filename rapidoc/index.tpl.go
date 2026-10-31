package rapidoc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/oaswrap/spec-ui/config"
)

func IndexTpl(assetBase, faviconBase string, cfg *config.RapiDoc) string {
	settings := map[string]string{
		"spec-url":              `"{{ .OpenAPIURL }}"`,
		"show-info":             fmt.Sprintf(`"%t"`, !cfg.HideInfo),
		"show-header":           fmt.Sprintf(`"%t"`, !cfg.HideHeader),
		"allow-search":          fmt.Sprintf(`"%t"`, !cfg.HideSearch),
		"allow-advanced-search": fmt.Sprintf(`"%t"`, !cfg.HideAdvancedSearch),
		"allow-try":             fmt.Sprintf(`"%t"`, !cfg.HideTryIt),
	}
	// Helper to add a quoted string if not empty
	addSetting := func(key, val string) {
		if val != "" {
			settings[key] = fmt.Sprintf(`"%s"`, val)
		}
	}

	// RapiDoc v11+ replaced theme="light|dark" with color-scheme.
	// Default to "light" to keep the pre-v11 look (v12 defaults to dark).
	colorScheme := string(cfg.Theme)
	if colorScheme == "" {
		colorScheme = string(config.RapiDocThemeLight)
	}
	addSetting("color-scheme", colorScheme)
	addSetting("layout", string(cfg.Layout))
	addSetting("render-style", string(cfg.RenderStyle))
	addSetting("schema-style", string(cfg.SchemaStyle))
	addSetting("primary-color", cfg.PrimaryColor)

	// bg-color / text-color attributes were removed in v11; use CSS variables.
	var styles []string
	if cfg.BgColor != "" {
		styles = append(styles, "--background:"+cfg.BgColor)
	}
	if cfg.TextColor != "" {
		styles = append(styles, "--foreground:"+cfg.TextColor)
	}
	addSetting("style", strings.Join(styles, ";"))

	settingsStr := make([]string, 0, len(settings))
	for k, v := range settings {
		settingsStr = append(settingsStr, "\t"+k+"="+v)
	}

	sort.Strings(settingsStr)

	faviconLink := ""
	if faviconBase != "" {
		faviconLink = `
	<link rel="shortcut icon" type="image/png" href="` + faviconBase + `/images/logo.png"/>`
	}

	return `
<!doctype html>
<html>
<head>
	<title>{{.Title}} - RapiDoc</title>
	<meta charset="utf-8">
	<script type="module" src="` + assetBase + `/rapidoc-min.js"></script>
` + faviconLink + `
</head>
<body>
<rapi-doc
` + strings.Join(settingsStr, "\n") + `
>
{{ if .Logo }}
	<img slot="nav-logo" src="{{ .Logo }}" />
{{ end }}
</rapi-doc>
</body>
</html>
`
}

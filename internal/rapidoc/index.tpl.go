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
		"show-info":             fmt.Sprintf("'%t'", !cfg.HideInfo),
		"show-header":           fmt.Sprintf("'%t'", !cfg.HideHeader),
		"allow-search":          fmt.Sprintf("'%t'", !cfg.HideSearch),
		"allow-advanced-search": fmt.Sprintf("'%t'", !cfg.HideAdvancedSearch),
		"allow-try":             fmt.Sprintf("'%t'", !cfg.HideTryIt),
	}
	// Helper to add a quoted string if not empty
	addSetting := func(key, val string) {
		if val != "" {
			settings[key] = fmt.Sprintf(`"%s"`, val)
		}
	}

	addSetting("theme", string(cfg.Theme))
	addSetting("layout", string(cfg.Layout))
	addSetting("render-style", string(cfg.RenderStyle))
	addSetting("schema-style", string(cfg.SchemaStyle))
	addSetting("bg-color", cfg.BgColor)
	addSetting("text-color", cfg.TextColor)
	addSetting("header-color", cfg.HeaderColor)
	addSetting("primary-color", cfg.PrimaryColor)

	settingsStr := make([]string, 0, len(settings))
	for k, v := range settings {
		settingsStr = append(settingsStr, "\t"+k+"="+v)
	}

	sort.Strings(settingsStr)

	return `
<!doctype html>
<html>
<head>
	<title>{{.Title}} - RapiDoc</title>
	<meta charset="utf-8">
	<script type="module" src="` + assetBase + `/rapidoc-min.js"></script>
	<link rel="shortcut icon" type="image/png" href="` + faviconBase + `/images/logo.png"/>
</head>
<body>
<rapi-doc
` + strings.Join(settingsStr, ",\n") + `
>
{{ if .Logo }}
	<img slot="nav-logo" src="{{ .Logo }}" />
{{ end }}
</rapi-doc>
</body>
</html>
`
}

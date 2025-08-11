package scalar

import (
	"fmt"
	"sort"
	"strings"

	"github.com/oaswrap/spec-ui/config"
)

func IndexTpl(assetBase string, cfg *config.Scalar) string {
	settings := map[string]string{
		"url":                   "url",
		"hideModels":            fmt.Sprintf("%t", cfg.HideModels),
		"showSidebar":           fmt.Sprintf("%t", !cfg.HideSidebar),
		"hideTestRequestButton": fmt.Sprintf("%t", cfg.HideTestRequestButton),
		"hideSearch":            fmt.Sprintf("%t", cfg.HideSearch),
		"darkMode":              fmt.Sprintf("%t", cfg.DarkMode),
	}
	if cfg.ProxyURL != "" {
		settings["proxyUrl"] = fmt.Sprintf("'%s'", cfg.ProxyURL)
	}
	if cfg.Layout != "" {
		settings["layout"] = fmt.Sprintf("'%s'", cfg.Layout)
	}
	if cfg.DocumentDownloadType != "" {
		settings["documentDownloadType"] = fmt.Sprintf("'%s'", cfg.DocumentDownloadType)
	}
	if cfg.Theme != "" {
		settings["theme"] = fmt.Sprintf("'%s'", cfg.Theme)
	}

	settingsStr := make([]string, 0, len(settings))
	for k, v := range settings {
		settingsStr = append(settingsStr, "\t\t\t\t"+k+": "+v)
	}

	sort.Strings(settingsStr)

	return `
<!doctype html>
<html>
<head>
	<title>{{.Title}} - Scalar</title>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<link href="` + assetBase + `/style.min.css" rel="stylesheet">
</head>
<body>
	<div id="app"></div>
	<script src="` + assetBase + `/browser/standalone.min.js"></script>
	<script>
		window.onload = function () {
			var url = "{{ .OpenAPIURL }}";
			if (!url.startsWith("https://") && !url.startsWith("http://")) {
				if (url.startsWith(".")) {
				var path = window.location.pathname;
				path = path.endsWith("/") ? path : path + "/";
					url = window.location.protocol + "//" + window.location.host + path + url;
				} else {
					url = window.location.protocol + "//" + window.location.host + url;
				}
			}
			Scalar.createApiReference('#app', {
` + strings.Join(settingsStr, ",\n") + `
			})
		}
	</script>
</body>
</html>
`
}

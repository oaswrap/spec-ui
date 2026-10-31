package redoc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/oaswrap/spec-ui/config"
)

func IndexTpl(assetBase string, cfg *config.ReDoc) string {
	settings := map[string]string{
		"disableTelemetry":    "true", // Redoc 3 sends usage telemetry to redocly.com by default.
		"hideDownloadButtons": fmt.Sprintf("%t", cfg.HideDownloadButtons),
		"hideSchemaTitles":    fmt.Sprintf("%t", cfg.HideSchemaTitles),
	}

	settingsStr := make([]string, 0, len(settings))
	for k, v := range settings {
		settingsStr = append(settingsStr, "\t\t\t"+k+": "+v)
	}

	sort.Strings(settingsStr)

	return `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}} - ReDoc</title>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body {
			margin: 0;
			padding: 0;
		}
	</style>
</head>
<body>
<div id="redoc-container"></div>
<script type="module">
	import { init } from "` + assetBase + `/redoc.standalone.js";
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
	const options = {
` + strings.Join(settingsStr, ",\n") + `
	}
	init(url, options, document.getElementById('redoc-container'))
</script>
</body>
</html>
`
}

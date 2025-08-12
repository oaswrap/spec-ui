package stoplightelements

import (
	"fmt"
	"sort"
	"strings"

	"github.com/oaswrap/spec-ui/config"
)

func IndexTpl(assetBase, faviconBase string, specCfg *config.SpecUI) string {
	settings := map[string]string{
		"hideExport":  fmt.Sprintf("%t", specCfg.StoplightElements.HideExport),
		"hideSchemas": fmt.Sprintf("%t", specCfg.StoplightElements.HideSchemas),
		"hideTryIt":   fmt.Sprintf("%t", specCfg.StoplightElements.HideTryIt),
	}
	cfg := specCfg.StoplightElements
	addSetting := func(key, val string) {
		if val != "" {
			settings[key] = fmt.Sprintf(`"%s"`, val)
		}
	}
	addSetting("router", string(cfg.Router))
	addSetting("layout", string(cfg.Layout))
	addSetting("logo", cfg.Logo)
	if cfg.Router == config.ElementRouterHistory || cfg.Router == "" {
		addSetting("basePath", specCfg.DocsPath)
	}

	settingsStr := make([]string, 0, len(settings))
	for k, v := range settings {
		settingsStr = append(settingsStr, "\t"+k+"="+v)
	}

	sort.Strings(settingsStr)

	return `
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>{{ .Title }} - Stoplight Elements</title>
    <link rel="stylesheet" href="` + assetBase + `/styles.min.css">
    <link rel="shortcut icon" type="image/x-icon" href="` + faviconBase + `/favicons/favicon.ico"/>
    <style>
        html, body {
        height: 100%;
        margin: 0;
    }
    </style>
</head>
<body>
<elements-api
    id="docs"
` + strings.Join(settingsStr, ",\n") + `
></elements-api>
<script src="` + assetBase + `/web-components.min.js"></script>
<script>
    window.onload = function () {
        (async () => {
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

            const docs = document.getElementById('docs');
            const text = await fetch(url).then(res => res.text())

            docs.apiDescriptionDocument = text;
        })();
    }
</script>
</body>
</html>
`
}

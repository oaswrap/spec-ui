package elements

func IndexTpl(assetBase string) string {
	return `
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>{{ .Title }} - Elements</title>

    <script src="` + assetBase + `web-components.min.js"></script>
    <link rel="stylesheet" href="` + assetBase + `styles.min.css">
    <style>
        html, body {
        height: 100%;
        margin: 0;
    }
    </style>
</head>
<body>
    <elements-api id="docs" router="{{ .Router }}" layout="{{ .Layout }}"></elements-api>
</body>
<script>
    window.onload = function () {
        (async () => {
            const cfg = {{ .ConfigJson }};
            var url = cfg.openapiYamlUrl;
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
            docs.hideTryIt = cfg.hideTryIt;
            docs.hideSchemas = cfg.hideSchemas;
            docs.logo = cfg.logo;
        })();
    }
</script>
</html>
`
}

package redoc

func IndexTpl(assetBase string) string {
	return `
<!DOCTYPE html>
<html>
	<head>
		<title>{{.Title}} - Redoc</title>
		<!-- needed for adaptive design -->
		<meta charset="utf-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
		<style>
		body {
			margin: 0;
			padding: 0;
		}
		</style>
	</head>
	<body>
		<div id="redoc-container"></div>
		<script src="` + assetBase + `redoc.standalone.js"> </script>
		<script>
			window.onload = function () {
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
				Redoc.init(url, {
					"expandResponses": "200,400",
					"hideDownloadButtons": cfg.hideDownload,
					"disableSearch": cfg.disableSearch,
					"hideSchemaTitles": cfg.hideSchemaTitles
				}, document.getElementById('redoc-container'))
			}
		</script>
	</body>
</html>
`
}

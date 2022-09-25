package client

import _ "embed"

//go:embed static/index.html
var indexHtml []byte

//go:embed static/bundle.js
var bundleJs []byte

//go:embed static/style.css
var styleCss []byte

func GetStaticFiles() map[string][]byte {
	return map[string][]byte{
		"/":          indexHtml,
		"/bundle.js": bundleJs,
		"/style.css": styleCss,
	}
}

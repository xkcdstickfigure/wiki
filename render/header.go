package render

import (
	"html"
	"os"
)

func renderHeader(title string, pctx PageContext) string {
	return `<div class="article-header">` +
		`<h1 class="title">` +
		html.EscapeString(title) +
		`</h1><h2 class="subtitle">` + html.EscapeString(pctx.Site+"."+os.Getenv("DOMAIN")+"/"+pctx.PageSlug) + `</h2></div>`
}

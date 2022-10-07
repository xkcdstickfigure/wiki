package render

import (
	"html"
)

func renderHeader(title string, pctx PageContext) string {
	output := `<div class="header">`
	output += `<h1 class="title">`
	output += html.EscapeString(title)
	output += `</h1><h2 class="subtitle">`
	output += html.EscapeString(pctx.Site + "." + pctx.Domain + "/" + pctx.PageSlug)
	output += `</h2></div>`
	return output
}

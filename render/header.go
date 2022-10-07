package render

import (
	"html"
	"os"
)

func renderHeader(title string, pctx PageContext) string {
	output := `<div class="header">`
	output += `<h1 class="title">`
	output += html.EscapeString(title)
	output += `</h1><h2 class="subtitle">`
	output += html.EscapeString(pctx.Site + "." + os.Getenv("DOMAIN") + "/" + pctx.PageSlug)
	output += `</h2></div>`
	return output
}

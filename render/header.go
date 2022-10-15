package render

import (
	"html"
)

func renderHeader(pctx PageContext) string {
	output := `<div>`
	output += `<h1 class="text-4xl font-semibold">`
	output += html.EscapeString(pctx.Title)
	output += `</h1>`
	output += `<h2 class="text-sm text-gray-600">`
	output += html.EscapeString(pctx.Site + "." + pctx.Domain + "/" + pctx.PageSlug)
	output += `</h2>`
	output += `</div>`
	return output
}

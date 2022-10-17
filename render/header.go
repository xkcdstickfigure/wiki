package render

import (
	"html"
)

func renderHeader(pctx PageContext) string {
	output := `<h1 class="text-4xl font-semibold">`
	output += html.EscapeString(pctx.Title)
	output += `</h1>`
	output += `<h2 class="text-sm flex space-x-1">`
	output += `<span class="text-gray-600 whitespace-pre overflow-hidden text-ellipsis">`
	output += html.EscapeString(pctx.Site + "." + pctx.Domain + "/" + pctx.Slug)
	output += `</span>`
	output += `<a class="text-blue-700 hover:underline" href="/` + html.EscapeString(pctx.Slug) + `/edit">[edit]</a>`
	output += `</h2>`
	return output
}

package render

import (
	"html"
)

func renderHeader(pctx PageContext) string {
	output := `<div class="flex justify-between space-x-2">`
	output += `<div class="min-w-0">`
	output += `<h1 class="text-4xl font-semibold">`
	output += html.EscapeString(pctx.Title)
	output += `</h1>`
	output += `<h2 class="text-sm text-gray-600 whitespace-pre overflow-hidden text-ellipsis">`
	output += html.EscapeString(pctx.Site + "." + pctx.Domain + "/" + pctx.Slug)
	output += `</h2>`
	output += `</div>`
	output += `<div>`
	output += `<a class="block py-1 px-2 uppercase text-xs text-gray-600 font-semibold bg-gray-100 border border-gray-200" href="/` + html.EscapeString(pctx.Slug) + `/edit">Edit</a>`
	output += `</div>`
	output += `</div>`
	return output
}

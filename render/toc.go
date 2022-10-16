package render

import (
	"html"
	"net/url"
	"strconv"
	"strings"

	"alles/wiki/markup"
)

func renderToc(sections []markup.Section) string {
	if len(sections) == 0 {
		return ""
	}

	output := `<div class="mt-8 bg-gray-100 border border-gray-200 p-3 table">`
	output += `<h2 class="text-center font-semibold">Contents</h2>`
	output += renderTocList(sections, []string{})
	output += `</div>`

	return output
}

func renderTocList(sections []markup.Section, path []string) string {
	output := `<ul`
	if len(path) > 1 {
		output += ` class="pl-4"`
	}
	output += `>`
	for i, section := range sections {
		output += `<li>`

		slug := url.QueryEscape(strings.ReplaceAll(strings.ToLower(section.Title), " ", "_"))
		path2 := append(path, strconv.Itoa(i+1))

		output += `<span class="text-gray-800">` + strings.Join(path2, ".") + `</span>`
		output += `<a href="#` + slug + `" class="text-blue-700 hover:underline ml-2">` + html.EscapeString(section.Title) + `</a>`

		output += renderTocList(section.Sections, path2)

		output += `</li>`
	}
	output += `</ul>`
	return output
}

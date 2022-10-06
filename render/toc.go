package render

import (
	"fmt"
	"html"
	"net/url"
	"strings"

	"alles/wiki/markup"
)

func renderToc(sections []markup.Section) string {
	if len(sections) == 0 {
		return ""
	}

	output := `<div class="toc">`
	output += `<h2>Contents</h2>`
	output += renderTocList(sections, []string{})
	output += `</div>`

	return output
}

func renderTocList(sections []markup.Section, path []string) string {
	output := `<ul>`
	for i, section := range sections {
		output += `<li>`

		slug := url.QueryEscape(strings.ReplaceAll(strings.ToLower(section.Title), " ", "_"))
		path2 := append(path, fmt.Sprintf("%v", i+1))

		output += `<a href="#` + slug + `">`
		output += `<span class="number">` + strings.Join(path2, ".") + `</span> `
		output += `<span class="name">` + html.EscapeString(section.Title) + `</span>`
		output += `</a>`

		output += renderTocList(section.Sections, path2)

		output += `</li>`
	}
	output += `</ul>`
	return output
}

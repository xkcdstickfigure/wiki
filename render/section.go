package render

import (
	"fmt"
	"html"
	"strings"

	"alles/wiki/markup"
)

func renderSections(sections []markup.Section, depth int, pctx PageContext) (string, error) {
	var output string

	for _, section := range sections {
		// start
		slug := strings.ReplaceAll(strings.ToLower(section.Title), " ", "_")
		output += `<div id="` + html.EscapeString(slug) + `" class="section section-depth-` + fmt.Sprintf("%v", depth+1) + `">`

		// title
		titleDepth := depth + 1
		if titleDepth > 6 {
			titleDepth = 6
		}
		output += `<h` + fmt.Sprintf("%v", titleDepth) + ` class="title">` + html.EscapeString(section.Title) + `</h1>`

		// elements
		elements, err := renderElements(section.Elements, pctx)
		if err != nil {
			return "", err
		}
		output += elements

		// sections
		sections, err := renderSections(section.Sections, depth+1, pctx)
		if err != nil {
			return "", err
		}
		output += sections

		// end
		output += `</div>`
	}

	return output, nil
}

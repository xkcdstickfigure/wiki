package render

import (
	"fmt"
	"html"
	"net/url"
	"strings"

	"alles/wiki/markup"
)

func renderSections(sections []markup.Section, depth int, pctx PageContext) (string, error) {
	var output string

	for _, section := range sections {
		// start
		slug := strings.ReplaceAll(strings.ToLower(section.Title), " ", "_")
		output += `<section id="` + html.EscapeString(slug) + `" class="section">`

		// title
		titleDepth := depth + 2
		if titleDepth > 6 {
			titleDepth = 6
		}
		titleElem := "h" + fmt.Sprintf("%v", titleDepth)
		output += `<` + titleElem + ` class="section-title">` + html.EscapeString(section.Title) + `</` + titleElem + `>`

		// media
		if len(section.Images) > 0 {
			output += `<aside class="media">`
			for _, image := range section.Images {

				text, err := renderText(image.Text, pctx)
				if err != nil {
					return output, err
				}

				output += `<div class="image-container">`
				output += `<img class="image" alt="` + html.EscapeString(image.Source) + `" src="` + pctx.StorageOrigin + `/sites/` + pctx.Site + `/images/` + url.QueryEscape(image.Source) + `/image.png" />`
				output += `<p class="caption">` + text + `</p>`
				output += `</div>`

			}
			output += `</aside>`
		}

		// elements
		elements, err := renderElements(section.Elements, pctx)
		if err != nil {
			return output, err
		}
		output += elements

		// sections
		sections, err := renderSections(section.Sections, depth+1, pctx)
		if err != nil {
			return output, err
		}
		output += sections

		// end
		output += `</section>`
	}

	return output, nil
}

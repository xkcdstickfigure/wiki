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
		output += `<section id="` + html.EscapeString(slug) + `" class="clear-both mt-12`
		if depth > 0 {
			output += ` pl-8`
		}
		output += `">`

		// title
		titleDepth := depth + 2
		if titleDepth > 6 {
			titleDepth = 6
		}
		titleElem := "h" + fmt.Sprintf("%v", titleDepth)

		var titleSize string
		if titleDepth == 2 {
			titleSize = "text-2xl"
		} else if titleDepth == 3 {
			titleSize = "text-xl"
		} else {
			titleSize = "text-lg"
		}

		output += `<` + titleElem + ` class="font-semibold ` + titleSize + `">` + html.EscapeString(section.Title) + `</` + titleElem + `>`

		// media
		if len(section.Images) > 0 {
			output += `<aside class="md:float-right md:ml-4 space-y-6">`
			for _, image := range section.Images {

				text, err := renderText(image.Text, pctx)
				if err != nil {
					return output, err
				}

				output += `<div class="p-2 max-w-xs table md:ml-auto bg-gray-100 border-gray-200">`
				output += `<div class="flex justifiy-center">`
				output += `<img alt="` + html.EscapeString(image.Source) + `" src="` + pctx.StorageOrigin + `/sites/` + pctx.Site + `/images/` + url.QueryEscape(image.Source) + `/image.png" />`
				output += `</div>`
				if text != "" {
					output += `<p class="mt-2 text-sm text-center text-gray-800">` + text + `</p>`
				}
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

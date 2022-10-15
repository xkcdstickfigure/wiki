package render

import (
	"errors"

	"alles/wiki/markup"
)

func renderElements(elements []markup.Element, pctx PageContext) (string, error) {
	var output string

	for _, element := range elements {
		if element.Type == "text" {

			// text
			text, err := renderText(element.Content[0], pctx)
			if err != nil {
				return output, err
			}

			output += `<p class="mt-4">` + text + `</p>`

		} else if element.Type == "list" {

			// list
			output += `<ul class="mt-4 pl-8 list-disc">`
			for _, item := range element.Content {
				text, err := renderText(item, pctx)
				if err != nil {
					return output, err
				}

				output += `<li>` + text + `</li>`
			}
			output += `</ul>`

		} else {

			return output, errors.New("invalid element type: " + element.Type)

		}
	}

	return output, nil
}

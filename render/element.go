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

			output += `<div class="element"><p>` + text + `</p></div>`

		} else if element.Type == "list" {

			// list
			output += `<ul class="element">`
			for _, item := range element.Content {
				text, err := renderText(item, pctx)
				if err != nil {
					return output, err
				}

				output += "<li>" + text + "</li>"
			}
			output += "</ul>"

		} else {

			return output, errors.New("invalid element type: " + element.Type)

		}
	}

	return output, nil
}

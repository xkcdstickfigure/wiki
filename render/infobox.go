package render

import (
	"alles/wiki/markup"
	"html"
	"net/url"
)

func renderInfobox(article markup.Article, pctx PageContext) (string, error) {
	if article.Image.Source == "" && len(article.Infobox) == 0 {
		return "", nil
	}

	output := `<aside class="md:float-right w-80 ml-auto md:ml-4 mr-auto md:mr-0 mb-12 md:mb-0 p-3 max-w-full bg-gray-100 border border-200 space-y-4">`

	if article.Image.Source != "" {
		caption := html.EscapeString(pctx.Title)
		if len(article.Image.Text) > 0 {
			var err error
			caption, err = renderText(article.Image.Text, pctx)
			if err != nil {
				return output, err
			}
		}

		output += `<div>`
		output += `<div class="flex justify-center">`
		output += `<img alt="` + html.EscapeString(article.Image.Source) + `" src="` + pctx.StorageOrigin + `/sites/` + pctx.Site + `/images/` + url.QueryEscape(article.Image.Source) + `/image.png" class="w-auto max-h-48" />`
		output += `</div>`
		output += `<p class="mt-2 text-xs text-center text-gray-800">` + caption + `</p>`
		output += `</div>`
	}

	for _, section := range article.Infobox {
		output += `<section class="space-y-2">`

		if section.Name != "" {
			output += `<h2 class="bg-gray-200 py-0.5 px-2 text-center font-semibold">` + html.EscapeString(section.Name) + `</h2>`
		}

		for _, field := range section.Fields {
			if len(field.Value) > 0 {
				// double (key = value)
				output += `<div class="text-base flex justify-between">`

				keyText, err := renderText(field.Key, pctx)
				if err != nil {
					return output, err
				}

				output += `<p>` + keyText + `</p>`
				output += `<div>`

				for _, value := range field.Value {
					valueText, err := renderText(value, pctx)
					if err != nil {
						return output, err
					}
					output += `<p>` + valueText + `</p>`
				}

				output += `</div>`
				output += `</div>`
			} else {
				// single (key only)
				keyText, err := renderText(field.Key, pctx)
				if err != nil {
					return output, err
				}

				output += `<p class="text-center">` + keyText + `</p>`
			}
		}

		output += `</section>`
	}

	output += `</aside>`

	return output, nil
}

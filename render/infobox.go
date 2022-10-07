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

	output := `<aside class="infobox">`

	if article.Image.Source != "" {
		caption := html.EscapeString(pctx.Title)
		if len(article.Image.Text) > 0 {
			var err error
			caption, err = renderText(article.Image.Text, pctx)
			if err != nil {
				return output, err
			}
		}

		output += `<div class="image-container">`
		output += `<img class="image" alt="` + html.EscapeString(article.Image.Source) + `" src="` + pctx.StorageOrigin + `/sites/` + pctx.Site + `/images/` + url.QueryEscape(article.Image.Source) + `/image.png" />`
		output += `<p class="caption">` + caption + `</p>`
		output += `</div>`
	}

	for _, section := range article.Infobox {
		output += `<section class="infobox-section">`

		if section.Name != "" {
			output += `<h2 class="infobox-section-title">` + html.EscapeString(section.Name) + `</h2>`
		}

		for _, field := range section.Fields {
			if len(field.Value) > 0 {
				// key = value
				output += `<div class="infobox-field infobox-field-double">`

				keyText, err := renderText(field.Key, pctx)
				if err != nil {
					return output, err
				}

				output += `<div class="key"><p>` + keyText + `</p></div>`
				output += `<div class="value">`

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
				// key only
				keyText, err := renderText(field.Key, pctx)
				if err != nil {
					return output, err
				}

				output += `<div class="infobox-field infobox-field-single"><p>` + keyText + `</p></div>`
			}
		}

		output += `</section>`
	}

	output += `</aside>`

	return output, nil
}

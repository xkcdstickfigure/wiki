package render

import (
	"errors"
	"html"
	"net/url"
	"strings"

	"alles/wiki/markup"
)

func renderText(text markup.Text, pctx PageContext) (string, error) {
	var output string

	for _, component := range text {
		// eek, not much type safety here
		if component.Type == "plain" {
			output += html.EscapeString(component.Value[0])
		} else if component.Type == "link internal" {
			output += renderLinkInternal(component.Value)
		} else if component.Type == "link external" {
			output += renderLinkExternal(component.Value)
		} else if component.Type == "icon" {
			output += renderIcon(component.Value, pctx)
		} else {
			return output, errors.New("invalid text component type: " + component.Type)
		}
	}

	return output, nil
}

func renderLinkInternal(data []string) string {
	display := data[0]
	page := data[1]
	section := data[2]

	pageUrl := ""
	if page != "" {
		pageUrl = "/" + url.QueryEscape(strings.ReplaceAll(strings.ToLower(page), " ", "_"))
	}

	sectionUrl := ""
	if section != "" {
		sectionUrl = "#" + url.QueryEscape(strings.ReplaceAll(strings.ToLower(section), " ", "_"))
	}

	return `<a class="text-blue-700 hover:underline" href="` + pageUrl + sectionUrl + `">` + html.EscapeString(display) + "</a>"
}

func renderLinkExternal(data []string) string {
	return `<a class="text-blue-700 hover:underline" target="_blank" href="` + data[1] + `">` + html.EscapeString(data[0]) + "</a>"
}

func renderIcon(data []string, pctx PageContext) string {
	return `<img class="inline h-4" alt="` + html.EscapeString(data[0]) + `" src="` + pctx.StorageOrigin + `/sites/` + pctx.Site + `/icons/` + url.QueryEscape(data[0]) + `/icon.png" />`
}

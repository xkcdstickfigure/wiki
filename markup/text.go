package markup

import (
	"strings"
)

type Text []TextComponent
type TextComponent struct {
	Type  string
	Value []string
}

func ParseText(str string) (Text, error) {
	text := Text{}

	componentType := "plain"
	value := ""

	for _, char := range strings.Split(strings.TrimSpace(str), "") {
		if componentType == "plain" {
			// plain
			if char == "{" {

				// start icon
				if value != "" {
					text = append(text, TextComponent{
						Type:  "plain",
						Value: []string{value},
					})
				}
				componentType = "icon"
				value = ""

			} else if char == "[" {

				// start link
				if value != "" {
					text = append(text, TextComponent{
						Type:  "plain",
						Value: []string{value},
					})
				}
				componentType = "link"
				value = ""

			} else {

				// continue
				value += char

			}

		} else if componentType == "icon" {

			// icon
			if char == "}" {
				text = append(text, TextComponent{
					Type:  "icon",
					Value: []string{value},
				})
				componentType = "plain"
				value = ""
			} else {
				value += char
			}

		} else if componentType == "link" {

			// link
			if char == "]" {
				external, data, err := parseLink(value)
				if err != nil {
					return text, err
				}

				if external {
					text = append(text, TextComponent{
						Type:  "external link",
						Value: data,
					})
				} else {
					text = append(text, TextComponent{
						Type:  "link internal",
						Value: data,
					})
				}

				componentType = "plain"
				value = ""
			} else {
				value += char
			}

		}
	}

	// end plain text
	if componentType == "plain" {
		if value != "" {
			text = append(text, TextComponent{
				Type:  "plain",
				Value: []string{value},
			})
		}
	}

	return text, nil
}

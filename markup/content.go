package markup

import (
	"errors"
	"strings"
)

type Content struct {
	Elements []Element
	Sections []Section
}

type Section struct {
	Title    string
	Elements []Element
	Images   []Image
	Sections []Section
}

type Element struct {
	Type    string
	Content []Text
}

type Image struct {
	Source string
	Text   Text
}

func parseContent(lines []string) (Content, error) {
	topLines, rawSections := organizeContentLines(lines, 0)
	elements, images, err := parseContentLines(topLines)
	if err != nil {
		return Content{}, err
	}

	if len(images) > 0 {
		return Content{}, errors.New("top level content cannot contain images")
	}

	sections, err := parseContentSections(rawSections)
	if err != nil {
		return Content{}, err
	}

	return Content{
		Elements: elements,
		Sections: sections,
	}, nil
}

func parseContentSections(rawSections []RawSection) ([]Section, error) {
	sections := []Section{}
	for _, rawSection := range rawSections {
		elements, images, err := parseContentLines(rawSection.Lines)
		if err != nil {
			return sections, err
		}

		subsections, err := parseContentSections(rawSection.Sections)
		if err != nil {
			return sections, err
		}

		sections = append(sections, Section{
			Title:    rawSection.Title,
			Elements: elements,
			Images:   images,
			Sections: subsections,
		})
	}

	return sections, nil
}

func parseContentLines(lines []string) ([]Element, []Image, error) {
	elements := []Element{}
	element := Element{}
	images := []Image{}

	for _, line := range lines {
		if strings.HasPrefix(line, ":img ") {

			// image
			img := strings.TrimPrefix(line, ":img ")
			source := strings.Split(img, " ")[0]
			text, err := parseText(strings.TrimPrefix(img, source))
			if err != nil {
				return elements, images, err
			}

			images = append(images, Image{
				Source: source,
				Text:   text,
			})

		} else if strings.HasPrefix(line, "- ") {

			// list
			text, err := parseText(strings.TrimPrefix(line, "- "))
			if err != nil {
				return elements, images, err
			}

			if element.Type == "list" {
				element.Content = append(element.Content, text)
			} else {
				if len(element.Content) > 0 {
					elements = append(elements, element)
				}
				element = Element{
					Type:    "list",
					Content: []Text{text},
				}
			}

		} else {

			// text
			text, err := parseText(line)
			if err != nil {
				return elements, images, err
			}

			if len(element.Content) > 0 {
				elements = append(elements, element)
			}
			element = Element{
				Type:    "text",
				Content: []Text{text},
			}

		}
	}

	if len(element.Content) > 0 {
		elements = append(elements, element)
	}

	return elements, images, nil
}

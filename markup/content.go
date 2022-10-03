package markup

import (
	"errors"
	"strings"
)

type Content struct {
	Elements []ContentElement
	Sections []ContentSection
}

type ContentSection struct {
	Title    string
	Elements []ContentElement
	Images   []ContentImage
	Sections []ContentSection
}

type ContentElement struct {
	Type    string
	Content []Text
}

type ContentImage struct {
	Source string
	Text   Text
}

func ParseContent(lines []string) (Content, error) {
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

func parseContentSections(rawSections []RawSection) ([]ContentSection, error) {
	sections := []ContentSection{}
	for _, rawSection := range rawSections {
		elements, images, err := parseContentLines(rawSection.Lines)
		if err != nil {
			return sections, err
		}

		subsections, err := parseContentSections(rawSection.Sections)
		if err != nil {
			return sections, err
		}

		sections = append(sections, ContentSection{
			Title:    rawSection.Title,
			Elements: elements,
			Images:   images,
			Sections: subsections,
		})
	}

	return sections, nil
}

func parseContentLines(lines []string) ([]ContentElement, []ContentImage, error) {
	elements := []ContentElement{}
	element := ContentElement{}
	images := []ContentImage{}

	for _, line := range lines {
		if strings.HasPrefix(line, ":img ") {

			// image
			img := strings.TrimPrefix(line, ":img ")
			source := strings.Split(img, " ")[0]
			text, err := ParseText(strings.TrimPrefix(img, source))
			if err != nil {
				return elements, images, err
			}

			images = append(images, ContentImage{
				Source: source,
				Text:   text,
			})

		} else if strings.HasPrefix(line, "- ") {

			// list
			text, err := ParseText(strings.TrimPrefix(line, "- "))
			if err != nil {
				return elements, images, err
			}

			if element.Type == "list" {
				element.Content = append(element.Content, text)
			} else {
				if len(element.Content) > 0 {
					elements = append(elements, element)
				}
				element = ContentElement{
					Type:    "list",
					Content: []Text{text},
				}
			}

		} else {

			// text
			text, err := ParseText(line)
			if err != nil {
				return elements, images, err
			}

			if len(element.Content) > 0 {
				elements = append(elements, element)
			}
			element = ContentElement{
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

package markup

import (
	"errors"
	"strings"
)

type Content struct {
	Intro    []string
	Sections []RawSection
}

func ParseContent(lines []string) (Content, error) {
	intro, rawSections, err := parseContentSection(lines, 0)
	if err != nil {
		return Content{}, err
	}

	return Content{intro, rawSections}, nil
}

type RawSection struct {
	Title    string
	Text     []string
	Sections []RawSection
}

func parseContentSection(lines []string, depth int) ([]string, []RawSection, error) {
	text := []string{}
	sections := []RawSection{}
	sectionTitle := ""
	sectionLines := []string{}
	titlePrefix := strings.Repeat("#", depth+1) + " "

	// go through lines
	for _, line := range lines {
		if strings.HasPrefix(line, titlePrefix) {
			// current section (recursion!)
			if sectionTitle != "" {
				sText, sSections, err := parseContentSection(sectionLines, depth+1)
				if err != nil {
					return text, sections, err
				}

				sections = append(sections, RawSection{
					Title:    sectionTitle,
					Text:     sText,
					Sections: sSections,
				})
			}

			// start new section
			sectionTitle = strings.TrimSpace(strings.TrimPrefix(line, titlePrefix))
			sectionLines = []string{}
			if sectionTitle == "" {
				return text, sections, errors.New("content section must have title")
			}
		} else {
			// add lines
			if sectionTitle == "" {
				text = append(text, line)
			} else {
				sectionLines = append(sectionLines, line)
			}
		}
	}

	// add final section
	if sectionTitle != "" {
		sText, sSections, err := parseContentSection(sectionLines, depth+1)
		if err != nil {
			return text, sections, err
		}

		sections = append(sections, RawSection{
			Title:    sectionTitle,
			Text:     sText,
			Sections: sSections,
		})
	}

	return text, sections, nil
}

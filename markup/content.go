package markup

import (
	"strings"
)

type Content struct {
	Intro    []string
	Sections []RawSection
}

func ParseContent(lines []string) (Content, error) {
	intro, rawSections := parseContentSection(lines, 0)
	return Content{intro, rawSections}, nil
}

type RawSection struct {
	Title    string
	Text     []string
	Sections []RawSection
}

func parseContentSection(lines []string, depth int) ([]string, []RawSection) {
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
				sText, sSections := parseContentSection(sectionLines, depth+1)
				sections = append(sections, RawSection{
					Title:    sectionTitle,
					Text:     sText,
					Sections: sSections,
				})
			}

			// start new section
			sectionTitle = strings.TrimSpace(strings.TrimPrefix(line, titlePrefix))
			sectionLines = []string{}
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
		sText, sSections := parseContentSection(sectionLines, depth+1)
		sections = append(sections, RawSection{
			Title:    sectionTitle,
			Text:     sText,
			Sections: sSections,
		})
	}

	return text, sections
}

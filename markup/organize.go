package markup

import (
	"strings"
)

type RawSection struct {
	Title    string
	Lines    []string
	Sections []RawSection
}

func organizeContentLines(lines []string, depth int) ([]string, []RawSection) {
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
				sLines, sSections := organizeContentLines(sectionLines, depth+1)
				sections = append(sections, RawSection{
					Title:    sectionTitle,
					Lines:    sLines,
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
		sLines, sSections := organizeContentLines(sectionLines, depth+1)
		sections = append(sections, RawSection{
			Title:    sectionTitle,
			Lines:    sLines,
			Sections: sSections,
		})
	}

	return text, sections
}

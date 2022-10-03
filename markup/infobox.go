package markup

import (
	"errors"
	"strings"
)

type Infobox []InfoboxSection
type InfoboxSection struct {
	Name   string
	Fields []InfoboxField
}
type InfoboxField struct {
	Key   string
	Value []string
}

func ParseInfobox(lines []string) (Infobox, error) {
	infobox := Infobox{}
	sectionName := ""
	fields := []InfoboxField{}

	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			// section
			if len(fields) > 0 {
				infobox = append(infobox, InfoboxSection{
					Name:   sectionName,
					Fields: fields,
				})
			}

			sectionName = strings.TrimSpace(strings.TrimPrefix(line, "# "))
			fields = []InfoboxField{}
			if sectionName == "" {
				return infobox, errors.New("infobox sections must have a title")
			}
		} else {
			// field
			fieldSplit := strings.Split(line, " = ")
			if len(fieldSplit) == 1 {
				// key only
				fields = append(fields, InfoboxField{
					Key:   line,
					Value: []string{},
				})
			} else if len(fieldSplit) == 2 {
				// key = value
				values := []string{}
				for _, value := range strings.Split(fieldSplit[1], " // ") {
					values = append(values, strings.TrimSpace(value))
				}

				fields = append(fields, InfoboxField{
					Key:   strings.TrimSpace(fieldSplit[0]),
					Value: values,
				})
			} else {
				return infobox, errors.New("infobox fields may not contain more than one equal sign")
			}
		}
	}

	if len(fields) > 0 {
		infobox = append(infobox, InfoboxSection{
			Name:   sectionName,
			Fields: fields,
		})
	}

	return infobox, nil
}

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
	Key   Text
	Value []Text
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
			split := strings.Split(line, " = ")
			if len(split) == 1 {
				// key only
				keyText, err := ParseText(line)
				if err != nil {
					return infobox, err
				}

				fields = append(fields, InfoboxField{
					Key:   keyText,
					Value: []Text{},
				})
			} else if len(split) == 2 {
				// key = value
				keyText, err := ParseText(strings.TrimSpace(split[0]))
				if err != nil {
					return infobox, err
				}

				values := []Text{}
				for _, value := range strings.Split(split[1], " // ") {
					valueText, err := ParseText(strings.TrimSpace(value))
					if err != nil {
						return infobox, err
					}

					values = append(values, valueText)
				}

				fields = append(fields, InfoboxField{
					Key:   keyText,
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

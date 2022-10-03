package markup

import (
	"errors"
	"strings"
)

type Meta map[string]string

func ParseMeta(lines []string) (Meta, error) {
	meta := Meta{}

	for _, line := range lines {
		split := strings.Split(line, " = ")
		if len(split) != 2 {
			return meta, errors.New("meta fields must contain a single equals sign between the key and value")
		}

		meta[strings.ToLower(strings.TrimSpace(split[0]))] = strings.TrimSpace(split[1])
	}

	return meta, nil
}

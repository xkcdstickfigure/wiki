package markup

import (
	"errors"
	"strings"
)

func ParseMeta(lines []string) (map[string]string, error) {
	meta := map[string]string{}

	for _, line := range lines {
		split := strings.Split(line, "=")
		if len(split) != 2 {
			return meta, errors.New("meta lines must contain a single equals sign between the key and values")
		}

		meta[strings.ToLower(strings.TrimSpace(split[0]))] = strings.TrimSpace(split[1])
	}

	return meta, nil
}

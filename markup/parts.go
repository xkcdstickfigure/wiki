package markup

import (
	"errors"
	"strings"
)

func SplitParts(str string) (map[string][]string, error) {
	var name string
	var content []string
	parts := map[string][]string{}

	lines := strings.Split(str, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "=== ") && strings.HasSuffix(line, " ===") {
			if i > 0 {
				parts[name] = content
			}
			name = strings.TrimPrefix(strings.TrimSuffix(line, " ==="), "=== ")
			content = []string{}
		} else {
			if i > 0 {
				l := strings.TrimSpace(line)
				if l != "" {
					content = append(content, l)
				}
			} else {
				return parts, errors.New("a part must be declared on the first line")
			}
		}
	}
	parts[name] = content

	return parts, nil
}

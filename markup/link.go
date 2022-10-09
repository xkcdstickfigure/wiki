package markup

import (
	"errors"
	"strings"
)

func parseLink(link string) (bool, []string, error) {
	split1 := strings.Split(link, "|")
	var display, target string

	if len(split1) == 1 {
		display = strings.TrimSpace(link)
		target = display
	} else if len(split1) == 2 {
		display = strings.TrimSpace(split1[0])
		target = strings.TrimSpace(split1[1])
	} else {
		return false, []string{}, errors.New("link must not have more than 2 parts (display and target)")
	}

	if strings.Contains(target, "/") {
		// external link
		return true, []string{display, target}, nil
	} else {
		// internal link
		var page string
		var section string

		split2 := strings.Split(target, "#")
		if len(split2) == 1 {
			page = target
		} else if len(split2) == 2 {
			page = split2[0]
			section = split2[1]
		} else {
			return false, []string{}, errors.New("internal link target must not contain multiple hashes")
		}

		return false, []string{display, page, section}, nil
	}
}

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

	if strings.Contains(target, "://") {
		// external url
		return true, []string{display, target}, nil
	} else {
		// internal link
		var site string
		var page string
		var section string

		// ...#section
		split2 := strings.Split(target, "#")
		if len(split2) == 2 {
			section = split2[1]
		} else if len(split2) > 2 {
			return false, []string{}, errors.New("internal link target must not contain multiple hashes")
		}

		// site:page
		split3 := strings.Split(split2[0], ":")
		if len(split3) == 1 {
			page = split3[0]
		} else if len(split3) == 2 {
			site = split3[0]
			page = split3[1]
		} else {
			return false, []string{}, errors.New("internal link target must not contain multiple colons")
		}

		site = strings.ToLower(site)
		page = strings.ReplaceAll(strings.ToLower(page), " ", "_")
		section = strings.ReplaceAll(strings.ToLower(section), " ", "_")
		return false, []string{display, site, page, section}, nil
	}
}

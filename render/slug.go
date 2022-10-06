package render

import (
	"net/url"
	"strings"
)

func slug(str string) string {
	return url.QueryEscape(strings.ReplaceAll(strings.ToLower(str), " ", "_"))
}

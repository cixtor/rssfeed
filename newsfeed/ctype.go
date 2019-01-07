package newsfeed

import (
	"strings"
)

func isValidContentType(ctype string) bool {
	allowed := map[string]bool{
		"text/html":                          true,
		"application/rss+xml":                true,
		"text/html;charset=utf-8":            true,
		"text/html; charset=utf-8":           true,
		"text/html; charset=iso-8859-1":      true,
		"text/html; charset=\"UTF-8\"":       true,
		"application/rss+xml; charset=utf-8": true,
	}

	ctype = strings.ToLower(ctype)

	if _, ok := allowed[ctype]; ok {
		return true
	}

	return false
}

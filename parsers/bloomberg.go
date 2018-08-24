package parsers

import (
	"strings"
)

func init() {
	register("www.bloomberg.com", bloomberg)
}

func bloomberg(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="body-copy fence-body">`)
	body = body[mark+34:]
	mark = strings.Index(body, `<ol class="noscript-footnotes">`)
	body = body[0:mark]

	return body
}

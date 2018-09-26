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

	if mark == -1 {
		mark = strings.Index(body, `<div class="body-copy-v2 fence-body">`)
	}

	body = body[mark:]
	mark = strings.Index(body, `<ol class="noscript-footnotes"></ol>`)
	body = body[0:mark]

	return body
}

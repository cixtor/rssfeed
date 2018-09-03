package parsers

import (
	"strings"
)

func init() {
	register("www.zdnet.com", zdnet)
}

func zdnet(body string) string {
	var mark int

	mark = strings.Index(body, `class="storyBody"`)
	body = body[mark:]
	mark = strings.Index(body, `>`)
	body = body[mark+1:]
	mark = strings.Index(body, `<section class="related-topics">`)
	body = body[0:mark]

	return body
}

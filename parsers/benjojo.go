package parsers

import (
	"strings"
)

func init() {
	register("blog.benjojo.co.uk", benjojo)
}

func benjojo(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="contain">`)
	body = body[mark+21:]
	mark = strings.Index(body, `<p>`)
	body = body[mark:]
	mark = strings.LastIndex(body, `</p>`)
	body = body[0 : mark+4]

	return body
}

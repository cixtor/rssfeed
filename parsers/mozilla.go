package parsers

import (
	"strings"
)

func init() {
	register("blog.mozilla.org", mozilla)
}

func mozilla(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="entry-content">`)
	body = body[mark:]
	mark = strings.Index(body, `</article>`)
	body = body[0:mark]

	return body
}

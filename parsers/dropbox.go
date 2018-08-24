package parsers

import (
	"strings"
)

func init() {
	register("blogs.dropbox.com", dropbox)
}

func dropbox(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="entry-content">`)
	body = body[mark:]
	mark = strings.Index(body, `<!-- .entry-content -->`)
	body = body[0:mark]

	return body
}

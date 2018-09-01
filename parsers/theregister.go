package parsers

import (
	"strings"
)

func init() {
	register("www.theregister.co.uk", theregister)
}

func theregister(body string) string {
	var mark int

	mark = strings.Index(body, `<div id="body">`)
	body = body[mark:]
	mark = strings.Index(body, `<div id=article_body_btm>`)
	body = body[0:mark]

	return body
}

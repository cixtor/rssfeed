package parsers

import (
	"strings"
)

func init() {
	register("www.theverge.com", theverge)
}

func theverge(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="article__body">`)
	body = body[mark:]
	mark = strings.Index(body, `<section class="c-nextclick">`)
	body = body[0:mark] + "</div>"

	return body
}

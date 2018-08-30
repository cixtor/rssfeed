package parsers

import (
	"strings"
)

func init() {
	register("www.theguardian.com", theguardian)
}

func theguardian(body string) string {
	var mark int

	mark = strings.Index(body, `<article`)
	body = body[mark:]
	mark = strings.Index(body, `</article>`)
	body = body[0:mark]
	mark = strings.Index(body, `<div class="content__article-body`)
	body = body[mark:]
	mark = strings.Index(body, `<div class="after-article js-after-article"></div>`)
	body = body[0:mark] + "</div>"

	return body
}

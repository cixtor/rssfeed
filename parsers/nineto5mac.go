package parsers

import (
	"strings"
)

func init() {
	register("9to5mac.com", nineto5mac)
}

func nineto5mac(body string) string {
	var mark int

	mark = strings.Index(body, `<article class="post-content"`)
	body = body[mark:]
	mark = strings.Index(body, `</article>`)
	body = body[0 : mark+10]
	mark = strings.Index(body, `<div class="post-body" itemprop="articleBody">`)
	body = body[mark:]
	mark = strings.Index(body, `<!-- .post-body -->`)
	body = body[0:mark]

	return body
}

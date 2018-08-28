package parsers

import (
	"strings"
)

func init() {
	register("www.reuters.com", reuters)
}

func reuters(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="StandardArticleBody_body"><p>`)
	body = body[mark:]
	mark = strings.Index(body, `<div class="Attribution_container">`)
	body = body[0:mark] + "</div>"

	return body
}

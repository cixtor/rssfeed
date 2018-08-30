package parsers

import (
	"strings"
)

func init() {
	register("techcrunch.com", techcrunch)
}

func techcrunch(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="article-content">`)
	body = body[mark:]
	mark = strings.Index(body, `<footer class="article-footer">`)
	body = body[0:mark]

	return body
}

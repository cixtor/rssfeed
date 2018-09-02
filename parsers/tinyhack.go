package parsers

import (
	"strings"
)

func init() {
	register("tinyhack.com", tinyhack)
}

func tinyhack(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="entry-content">`)
	body = body[mark:]
	mark = strings.Index(body, `<footer class="entry-footer">`)
	body = body[0:mark]

	return body
}

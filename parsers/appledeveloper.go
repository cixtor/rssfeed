package parsers

import (
	"strings"
)

func init() {
	register("developer.apple.com", appledeveloper)
}

func appledeveloper(body string) string {
	var mark int

	mark = strings.Index(body, `<article>`)
	body = body[mark:]
	mark = strings.Index(body, `</article>`)
	body = body[0 : mark+10]

	return body
}

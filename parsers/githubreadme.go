package parsers

import (
	"strings"
)

func init() {
	register("github.com", githubreadme)
}

func githubreadme(body string) string {
	var mark int

	mark = strings.Index(body, `<article`)

	if mark == -1 {
		return ""
	}

	body = body[mark:]
	mark = strings.Index(body, `</article>`)
	body = body[0 : mark+10]

	return body
}

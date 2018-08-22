package parsers

import (
	"strings"
)

func init() {
	register("medium.com", medium)
}

func medium(body string) string {
	var mark int

	if strings.Index(body, `content="medium://`) == -1 {
		return ""
	}

	mark = strings.Index(body, `</header>`)
	body = body[mark+9:]
	mark = strings.Index(body, `<footer`)
	body = body[0:mark]

	return body
}

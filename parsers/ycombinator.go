package parsers

import (
	"strings"
)

func init() {
	register("news.ycombinator.com", ycombinator)
}

func ycombinator(body string) string {
	var mark int

	mark = strings.Index(body, `<tr style="height:2px">`)

	if mark == -1 {
		return ""
	}

	body = body[mark:]
	mark = strings.Index(body, `<tr style="height:10px">`)
	body = body[0:mark]

	return body
}

package parsers

import (
	"strings"
)

func init() {
	register("lwn.net", lwnnet)
}

func lwnnet(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="GAByline">`)
	body = body[mark+22:]
	mark = strings.Index(body, `<hr width="`)
	body = body[0:mark]

	return body
}

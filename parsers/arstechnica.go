package parsers

import (
	"strings"
)

func init() {
	register("arstechnica.com", arstechnica)
}

func arstechnica(body string) string {
	var mark int

	mark = strings.Index(body, `<article`)
	body = body[mark:]
	mark = strings.Index(body, `itemprop="articleBody"`)
	body = "<div\x20" + body[mark:]
	mark = strings.Index(body, `<nav class="page-numbers">`)
	body = body[0:mark]

	return body
}

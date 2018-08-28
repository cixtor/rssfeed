package parsers

import (
	"strings"
)

func init() {
	register("phys.org", physorg)
}

func physorg(body string) string {
	var mark int

	mark = strings.Index(body, `<!--begin:article-block-->`)
	body = body[mark+26:]
	mark = strings.Index(body, `<!--end:article-block-->`)
	body = body[0:mark]
	mark = strings.Index(body, `<div class="first-block">`)
	body = body[mark:]
	mark = strings.Index(body, `<footer`)
	body = body[:mark]

	return body
}

package parsers

import (
	"strings"
)

func init() {
	register("www.bbc.co.uk", bbcnews)
}

func bbcnews(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="story-body__inner" property="articleBody">`)

	if mark == -1 {
		return body
	}

	body = body[mark:]
	mark = strings.Index(body, `<div id="topic-tags">`)
	body = body[0:mark]

	return body
}

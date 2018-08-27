package parsers

import (
	"strings"
)

func init() {
	register("www.nytimes.com", nytimes)
}

func nytimes(body string) string {
	var mark int

	mark = strings.Index(body, `<article`)
	body = body[mark:]
	mark = strings.Index(body, `</article>`)
	body = body[0:mark]
	mark = strings.Index(body, `StoryBodyCompanionColumn`)
	body = `<div class="css-18sbwfn ` + body[mark:]
	mark = strings.Index(body, `<div class="bottom-of-article">`)
	body = body[0:mark]

	return body
}

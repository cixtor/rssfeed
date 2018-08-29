package parsers

import (
	"strings"
)

func init() {
	register("spectrum.ieee.org", spectrumieee)
}

func spectrumieee(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="articleBody entry-content">`)
	body = body[mark:]
	mark = strings.Index(body, `<article class="sml_article_static automaton">`)
	body = body[0:mark]

	return body
}

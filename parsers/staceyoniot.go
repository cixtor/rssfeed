package parsers

import (
	"strings"
)

func init() {
	register("staceyoniot.com", staceyoniot)
}

func staceyoniot(body string) string {
	var mark int

	mark = strings.Index(body, `<main class="content"><article`)
	body = body[mark+22:]
	mark = strings.Index(body, `</article>`)
	body = body[0 : mark+10]
	mark = strings.Index(body, `</header>`)
	body = body[mark+9:]
	mark = strings.Index(body, "<div class=`mailmunch-forms-after-post`")
	body = body[0:mark]

	return body
}

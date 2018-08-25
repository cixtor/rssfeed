package parsers

import (
	"strings"
)

func init() {
	register("www.blog.google", googleblog)
}

func googleblog(body string) string {
	var mark int

	mark = strings.Index(body, `<div class="uni-content uni-blog-article-container uni-tombstone">`)
	body = body[mark:]
	mark = strings.Index(body, `<div class="share-sticky-end"></div>`)
	body = body[0:mark]

	return body
}

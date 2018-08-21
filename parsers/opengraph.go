package parsers

import (
	"strings"
)

// parseOpenGraph parses the HTML code looking for OpenGraph tags.
//
// ```
// <meta property="og:title" content="foobar">
// <meta property="og:image" content="6fc0c0a.jpg">
// <meta property="og:description" content="foobar">
// ```
func parseOpenGraph(host string, body string) string {
	var content string

	// title := parseOpenGraphProperty(body, "title")
	image := parseOpenGraphProperty(body, "image")
	description := parseOpenGraphProperty(body, "description")

	if image != "" {
		content += "<div><img src=\"" + image + "\"></div>"
	}

	if description != "" {
		content += "<p>" + description + "</p>"
	}

	content += `<hr><small><em>&mdash;Missing "` + host + `" parser&mdash;</em></small>`

	return content
}

// parseOpenGraphImage parses the HTML code looking for OpenGraph tags.
func parseOpenGraphProperty(body string, property string) string {
	mark := strings.Index(body, `property="og:`+property+`"`)

	if mark == -1 {
		return ""
	}

	body = body[mark+19:]
	mark = strings.Index(body, `content="`)

	if mark == -1 {
		return ""
	}

	body = body[mark+9:]
	mark = strings.Index(body, `"`)

	return body[0:mark]
}

package parsers

var supported = make(map[string]Parser)

type Parser func(string) string

func register(host string, callback Parser) {
	supported[host] = callback
}

// Article parses the HTML and returns the relevant content and the name of the
// parser to allow for an easy identification of the code in case that it fails
// or causes a panic.
func Article(host string, body string) (string, string) {
	if callback, ok := supported[host]; ok {
		return callback(body), host
	}

	if content := medium(body); content != "" {
		return content, "medium"
	}

	return parseOpenGraph(host, body), "opengraph"
}

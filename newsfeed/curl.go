package newsfeed

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const timeout time.Duration = 10 * time.Second

func Curl(target string) (io.Reader, error) {
	var err error
	var uri *url.URL
	var req *http.Request
	var res *http.Response

	if uri, err = url.Parse(target); err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: timeout}

	if req, err = http.NewRequest("GET", target, nil); err != nil {
		return nil, err
	}

	req.Header.Set("host", uri.Host)
	req.Header.Set("authority", uri.Host)
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("connection", "keep-alive")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("cookie", "__cfduid=d73722d8bb11742b3676371f1c97f19d11517447820")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("accept-language", "en-US,en;q=0.9")

	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if mime := res.Header.Get("Content-Type"); !isValidContentType(mime) {
		return nil, fmt.Errorf("invalid content-type `%s`", mime)
	}

	var buf bytes.Buffer
	(&buf).ReadFrom(res.Body)
	return &buf, nil
}

func isValidContentType(mime string) bool {
	allowed := map[string]bool{
		"text/html":                          true,
		"application/rss+xml":                true,
		"text/html;charset=utf-8":            true,
		"text/html;charset=UTF-8":            true,
		"text/html; charset=utf-8":           true,
		"text/html; charset=UTF-8":           true,
		"text/html; charset=iso-8859-1":      true,
		"text/html; charset=\"UTF-8\"":       true,
		"application/rss+xml; charset=utf-8": true,
	}

	if value, ok := allowed[mime]; ok {
		return value
	}

	return false
}

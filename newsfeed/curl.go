package newsfeed

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const timeout time.Duration = 6 * time.Second

func Curl(target string) (io.Reader, error) {
	info, err := url.Parse(target)

	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", target, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("host", info.Host)
	req.Header.Set("authority", info.Host)
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("connection", "keep-alive")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("cookie", "__cfduid=d73722d8bb11742b3676371f1c97f19d11517447820")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("accept-language", "en-US,en;q=0.9")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	ctype := res.Header.Get("Content-Type")

	if !isValidContentType(ctype) {
		return nil, fmt.Errorf("invalid content-type `%s`", ctype)
	}

	var buf bytes.Buffer
	(&buf).ReadFrom(res.Body)
	return &buf, nil
}

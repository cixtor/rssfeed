package newsfeed

import (
	"encoding/xml"
	"io"
)

// New returns an instance of the RSS feed from HackerNews.
//
// Parameter `batch` defines the number of concurrent web crawlers.
func New(batch int) (*Feed, error) {
	var err error
	var reader io.Reader

	rss := &Feed{batch: batch}

	if reader, err = Curl("https://news.ycombinator.com/rss"); err != nil {
		return nil, err
	}

	if err = xml.NewDecoder(reader).Decode(&rss.data); err != nil {
		return nil, err
	}

	return rss, nil
}

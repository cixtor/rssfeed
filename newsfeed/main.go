package newsfeed

import (
	"encoding/xml"

	"github.com/cixtor/rssfeed/mercury"
)

// New returns an instance of the RSS feed from HackerNews.
//
// Parameter `hand` defines the number of concurrent web crawlers.
func New(client *mercury.Mercury, hand int) (*Feed, error) {
	rss := new(Feed)

	rss.hand = hand
	rss.client = client

	reader, err := Curl("https://news.ycombinator.com/rss")

	if err != nil {
		return nil, err
	}

	if err := xml.NewDecoder(reader).Decode(&rss.data); err != nil {
		return nil, err
	}

	return rss, nil
}

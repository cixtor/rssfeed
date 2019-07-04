package main

import (
	"encoding/xml"
	"io"
	"log"
	"sync"
	"time"
)

type Feed struct {
	sync.WaitGroup
	batch int
	data  News
	items []Item
}

type News struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	BuildDate   string `xml:"lastBuildDate"`
	Items       []Item `xml:"item"`
}

// NewFeed returns an instance of the RSS feed from HackerNews.
//
// Parameter `batch` defines the number of concurrent web crawlers.
func NewFeed(batch int) (*Feed, error) {
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

func (v *Feed) TotalItems() int {
	return len(v.data.Channel.Items)
}

func (v *Feed) DownloadItems() {
	var wg sync.WaitGroup

	wg.Add(v.TotalItems())

	sem := make(chan bool, v.batch)

	for _, item := range v.data.Channel.Items {
		go func(wg *sync.WaitGroup, sem chan bool, v *Feed, item Item) {
			sem <- true
			defer func() { <-sem }()
			defer func() { wg.Done() }()

			if err := item.Curate(); err != nil {
				log.Printf("%s [%s]", err, item.Link)
				return
			}

			v.items = append(v.items, item)
		}(&wg, sem, v, item)
	}

	wg.Wait()
}

func (v *Feed) Encode(w io.Writer) error {
	w.Write([]byte(xml.Header))

	v.data.Channel.Items = v.items
	v.data.Channel.BuildDate = time.Now().Format(time.RFC1123Z)

	writer := xml.NewEncoder(w)
	writer.Indent("", "\t")
	return writer.Encode(v.data)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/cixtor/rssfeed/newsfeed"
)

const concurrency int = 30

func init() {
	router.GET("/news.rss", webNews)
}

func webNews(w http.ResponseWriter, r *http.Request) {
	rss, err := newsfeed.New(concurrency)

	if err != nil {
		fmt.Fprintf(w, "newsfeed.New %s", err)
		return
	}

	rss.DownloadItems()

	w.Header().Set("content-type", "application/rss+xml")

	if err := rss.Encode(w); err != nil {
		fmt.Fprintf(w, "rss.Encode %s", err)
		return
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/cixtor/rssfeed/newsfeed"
)

const concurrency int = 30

func init() {
	router.GET("/news.rss", webNews)
}

func webNews(w http.ResponseWriter, r *http.Request) {
	rss, err := newsfeed.New(client, concurrency)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	rss.DownloadItems()

	w.Header().Set("content-type", "application/rss+xml")

	if err := rss.Encode(w); err != nil {
		log.Println(err)
		return
	}
}

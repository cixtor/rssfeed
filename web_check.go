package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cixtor/rssfeed/newsfeed"
)

func init() {
	router.GET("/check", webCheck)
}

func webCheck(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("url")

	if link == "" {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	item := newsfeed.Item{Link: link}
	out, origin := item.Download()
	_ = os.Remove("/tmp/.txt")

	w.Header().Set("content-type", "text/plain")
	w.Header().Set("x-html-parser", origin)

	if _, err := w.Write([]byte(out)); err != nil {
		log.Println(err)
		return
	}
}

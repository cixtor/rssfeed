package main

import (
	"io"
	"net/http"

	"github.com/cixtor/readability"
	"github.com/cixtor/rssfeed/newsfeed"
)

func init() {
	router.GET("/test", webCheck)
	router.GET("/check", webCheck)
}

func webCheck(w http.ResponseWriter, r *http.Request) {
	var err error
	var uri string
	var rdr io.Reader
	var doc readability.Article

	if uri = r.URL.Query().Get("url"); uri == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	if rdr, err = newsfeed.Curl(uri); err != nil {
		http.Error(w, "newsfeed.Curl: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if doc, err = readability.FromReader(rdr, uri); err != nil {
		http.Error(w, "readability.FromReader: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if _, err = w.Write([]byte(doc.Content)); err != nil {
		panic(err)
	}
}

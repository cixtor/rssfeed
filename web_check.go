package main

import (
	"log"
	"net/http"
	"time"

	"github.com/cixtor/readability"
)

func init() {
	router.GET("/test", webCheck)
	router.GET("/check", webCheck)
}

func webCheck(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("url")

	if link == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	doc, err := readability.FromURL(link, 10*time.Second)

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("content-type", "text/plain")

	if _, err := w.Write([]byte(doc.Content)); err != nil {
		log.Println(err)
		return
	}
}

package main

import (
	"net/http"
)

func init() {
	router.GET("/", webIndex)
}

func webIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("rssfeed.index\n"))
}

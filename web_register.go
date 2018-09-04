package main

import (
	"net/http"
)

func init() {
	router.GET("/register", webRegister)
}

func webRegister(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	if token == "" {
		w.Write([]byte("Pass the Mercury API token to register the app\n"))
		w.Write([]byte("http://localhost:9628/register?token=API_TOKEN\n"))
		return
	}

	client.SetToken(token)

	http.Redirect(w, r, "/news.rss", 302)
}

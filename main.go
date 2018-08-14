package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cixtor/middleware"
)

var router = middleware.New()

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	router.Port = "9628"
	router.IdleTimeout = 5
	router.ReadTimeout = 5
	router.WriteTimeout = 10

	go func() {
		<-shutdown
		router.Shutdown()
	}()

	router.ListenAndServe()
}

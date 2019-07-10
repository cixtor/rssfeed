package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cixtor/middleware"
)

var router = middleware.New()

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	router.Port = "9628"
	router.IdleTimeout = 5 * time.Second
	router.ReadTimeout = 5 * time.Second
	router.WriteTimeout = 40 * time.Second
	router.ShutdownTimeout = 10 * time.Second
	router.ReadHeaderTimeout = 5 * time.Second

	go func() {
		<-shutdown
		router.Shutdown()
	}()

	router.ListenAndServe()
}

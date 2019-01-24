package main

import (
	"gin-gonic/config"
	"gin-gonic/server"
	"gin-gonic/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// @title Host Agent API
// @version 1.0
// @description This is the gateway to access internal serving components
func main() {
	log.Printf("Starting application")

	appConfig := config.Get()
	log.Printf("Using application configurations\n %#v", appConfig)

	log.Printf("Creating services")

	service.Add(service.IExampleServiceType, service.NewExampleService())

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	log.Printf("Creating servers")

	httpServer := server.NewHttpServer(appConfig, &waitGroup)
	registerShutdownHook(httpServer)

	waitGroup.Wait()
}

func registerShutdownHook(server *http.Server) {
	log.Printf("Registering shutdown signals")

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	go func() {
		sig := <-signalChan

		log.Printf("Caught signal: %+v", sig)

		defer func() {
			log.Printf("Stopping HTTP server")

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			server.Shutdown(ctx)
		}()
	}()
}

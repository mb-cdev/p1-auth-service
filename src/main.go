package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/mb-cdev/p1-auth-service/infrastructure"
	"github.com/mb-cdev/p1-auth-service/registry"
	"github.com/mb-cdev/p1-auth-service/settings"
)

func main() {
	settings.InstanceId = rand.Int()

	registry.ControllerRegistry.Init()
	r := infrastructure.GetRoutes()

	go func(mx *chi.Mux) {
		err := http.ListenAndServe("0.0.0.0:8080", mx)
		if err != nil {
			panic(err)
		}
	}(r)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt)

	log.Default().Printf("Started... [%d]\n", os.Getpid())
	<-sigChan
	log.Default().Println("BYE!")
}

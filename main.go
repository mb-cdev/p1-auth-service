package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/mb-cdev/p1-auth-service/controller"
	"github.com/mb-cdev/p1-auth-service/gateway"
	"github.com/mb-cdev/p1-auth-service/usecase/user/usecase"
)

func main() {
	db := gateway.NewUserMemoryDb()
	userUsecase := usecase.NewUser(db)
	_ = controller.NewUser(userUsecase)

	r := chi.NewMux()
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.Write(
			[]byte(
				fmt.Sprintln(
					r.PostFormValue("name"),
					r.PostFormValue("password")),
			),
		)
	})

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

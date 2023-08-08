package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mb-cdev/p1-auth-service/registry"
)

func GetRoutes() *chi.Mux {
	r := chi.NewMux()

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		out := registry.ControllerRegistry.User.RegisterAction(r.Form)
		outJson, err := json.Marshal(out)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(out.HttpStatusCode)
		w.Write(outJson)
	})

	return r
}

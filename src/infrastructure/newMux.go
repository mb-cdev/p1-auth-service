package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mb-cdev/jwt"
	"github.com/mb-cdev/p1-auth-service/registry"
	"github.com/mb-cdev/p1-auth-service/settings"
)

func GetRoutes() *chi.Mux {
	r := chi.NewMux()

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("InstanceId", fmt.Sprint(settings.InstanceId))
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

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("InstanceId", fmt.Sprint(settings.InstanceId))
		r.ParseForm()

		out := registry.ControllerRegistry.User.LoginAction(r.Form)
		outJson, err := json.Marshal(out)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Add("InstanceId", fmt.Sprint(settings.InstanceId))

		w.WriteHeader(out.HttpStatusCode)
		w.Write(outJson)
	})

	r.Post("/validate-jwt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("InstanceId", fmt.Sprint(settings.InstanceId))
		r.ParseForm()

		token := r.Form.Get("token")
		tokenValid, _ := jwt.IsTokenValidFromString("", token, settings.JwtSecret)
		if !tokenValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	return r
}

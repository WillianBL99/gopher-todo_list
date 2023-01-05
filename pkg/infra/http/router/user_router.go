package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/", GetUsers)
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}
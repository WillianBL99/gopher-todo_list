package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func TaskRouter(r chi.Router) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", GetTasks)
	})
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all tasks"))
}